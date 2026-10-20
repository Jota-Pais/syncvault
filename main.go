package main

import (
	"log"

	"syncvault/internal/config"
	"syncvault/internal/controller"
	"syncvault/internal/mapper"
	"syncvault/internal/model"
	"syncvault/internal/repository"
	"syncvault/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.ConectarBanco()
	if err != nil {
		log.Fatal("não conectou no banco: ", err)
	}

	err = config.RodarMigrations()
	if err != nil {
		log.Fatal("não rodou as migrations: ", err)
	}

	err = config.ValidarEntidades(db, &model.Produto{}, &model.Canal{}, &model.Venda{})
	if err != nil {
		log.Fatal("entidade não bate com o banco: ", err)
	}

	// Liga as camadas de cada entidade pelo construtor, de baixo para cima:
	// Repository e Mapper, depois Service, depois Controller.
	produtoRepository := repository.NovoProdutoRepository(db)
	canalRepository := repository.NovoCanalRepository(db)
	vendaRepository := repository.NovoVendaRepository(db)

	produtoMapper := mapper.NovoProdutoMapper()
	canalMapper := mapper.NovoCanalMapper()
	vendaMapper := mapper.NovoVendaMapper(produtoMapper, canalMapper)

	produtoService := service.NovoProdutoService(produtoRepository, produtoMapper)
	canalService := service.NovoCanalService(canalRepository, canalMapper)
	vendaService := service.NovoVendaService(vendaRepository, produtoRepository, canalRepository, vendaMapper)

	produtoController := controller.NovoProdutoController(produtoService)
	canalController := controller.NovoCanalController(canalService)
	vendaController := controller.NovoVendaController(vendaService)

	r := gin.Default()
	config.RegistrarSwagger(r)

	produtoController.RegistrarRotas(r)
	canalController.RegistrarRotas(r)
	vendaController.RegistrarRotas(r)

	log.Println("SyncVault no ar: http://localhost:8080/swagger/index.html")
	err = r.Run(":8080")
	if err != nil {
		log.Fatal("a API não subiu na porta 8080 (ela já está em uso?): ", err)
	}
}
