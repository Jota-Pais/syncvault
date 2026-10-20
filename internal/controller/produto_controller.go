// CAMADA: Controller
// RESPONSABILIDADE: receber as requisições de /produtos e devolver status e corpo.
// ROTAS ATENDIDAS: POST /produtos, GET /produtos e GET /produtos/{id}.
// SITUAÇÃO: implementada.

package controller

import (
	"net/http"

	"syncvault/internal/dto"
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
	produtos := r.Group("/produtos")
	produtos.POST("", c.Criar)
	produtos.GET("", c.Listar)
	produtos.GET("/:id", c.BuscarPorID)
}

// Criar atende POST /produtos. O ShouldBindJSON lê o JSON e aplica as
// validações das tags binding do DTO, como o @Valid do Java.
func (c *ProdutoController) Criar(ctx *gin.Context) {
	var requisicao dto.ProdutoRequestDTO
	err := ctx.ShouldBindJSON(&requisicao)
	if err != nil {
		ctx.Error(err)
		return
	}

	resposta, err := c.service.Criar(requisicao)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, resposta)
}

// Listar atende GET /produtos. Sem produto, devolve 200 com lista vazia.
func (c *ProdutoController) Listar(ctx *gin.Context) {
	resposta, err := c.service.Listar()
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, resposta)
}

// BuscarPorID atende GET /produtos/{id}.
func (c *ProdutoController) BuscarPorID(ctx *gin.Context) {
	id, err := lerID(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	resposta, err := c.service.BuscarPorID(id)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, resposta)
}
