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
//
// @Summary     Cadastra um produto
// @Description Cadastra um produto com o estoque inicial. O estoque é um só para todos os canais.
// @Tags        produtos
// @Accept      json
// @Produce     json
// @Param       produto body     dto.ProdutoRequestDTO true "Produto a cadastrar"
// @Success     201     {object} dto.ProdutoResponseDTO
// @Router      /produtos [post]
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
//
// @Summary     Lista os produtos
// @Description Lista todos os produtos, com o estoque atual. Sem produto, devolve lista vazia.
// @Tags        produtos
// @Produce     json
// @Success     200 {array} dto.ProdutoResponseDTO
// @Router      /produtos [get]
func (c *ProdutoController) Listar(ctx *gin.Context) {
	resposta, err := c.service.Listar()
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, resposta)
}

// BuscarPorID atende GET /produtos/{id}.
//
// @Summary     Busca um produto
// @Tags        produtos
// @Produce     json
// @Param       id  path     int true "Id do produto"
// @Success     200 {object} dto.ProdutoResponseDTO
// @Router      /produtos/{id} [get]
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
