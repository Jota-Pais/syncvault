// CAMADA: Controller
// RESPONSABILIDADE: receber as requisições de /produtos e devolver status e corpo.
// ROTAS ATENDIDAS: POST /produtos, GET /produtos e GET /produtos/{id}.
// SITUAÇÃO: montada. As rotas entram quando a API simples for implementada.

package controller

import (
	"syncvault/internal/service"

	"github.com/gin-gonic/gin"
)

type ProdutoController struct {
	service *service.ProdutoService
}

func NovoProdutoController(service *service.ProdutoService) *ProdutoController {
	return &ProdutoController{service: service}
}

// RegistrarRotas liga cada verbo e caminho de /produtos a um método deste
// controller, como o @RequestMapping e o @GetMapping do Java.
func (c *ProdutoController) RegistrarRotas(r *gin.Engine) {
	// Montada: as rotas de /produtos entram aqui quando forem implementadas.
}
