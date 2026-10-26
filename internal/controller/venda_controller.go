// CAMADA: Controller
// RESPONSABILIDADE: receber as requisições de /vendas e devolver status e corpo.
//                   As regras R1, R2 e R3 não moram aqui: moram no VendaService.
// ROTAS ATENDIDAS: POST /vendas.
// SITUAÇÃO: implementada.

package controller

import (
	"net/http"

	"syncvault/internal/dto"
	"syncvault/internal/service"

	"github.com/gin-gonic/gin"
)

type VendaController struct {
	service *service.VendaService
}

func NovoVendaController(service *service.VendaService) *VendaController {
	return &VendaController{service: service}
}

// RegistrarRotas liga cada verbo e caminho de /vendas a um método deste controller.
func (c *VendaController) RegistrarRotas(r *gin.Engine) {
	r.POST("/vendas", c.Registrar)
}

// Registrar atende POST /vendas.
//
// @Summary     Registra uma venda
// @Description Registra uma venda de qualquer canal e desconta do estoque único. As regras são conferidas nesta ordem: R2 (o preço, descontada a taxa do canal, não fica abaixo do custo), R1 (a quantidade não passa do estoque) e R3 (a venda não deixa o estoque abaixo do estoque de segurança do canal). Regra violada devolve 400, com a regra na mensagem.
// @Tags        vendas
// @Accept      json
// @Produce     json
// @Param       venda body     dto.VendaRequestDTO true "Venda a registrar"
// @Success     201   {object} dto.VendaResponseDTO
// @Failure     400   {object} dto.ErroDTO
// @Failure     404   {object} dto.ErroDTO
// @Router      /vendas [post]
func (c *VendaController) Registrar(ctx *gin.Context) {
	var requisicao dto.VendaRequestDTO
	err := ctx.ShouldBindJSON(&requisicao)
	if err != nil {
		ctx.Error(err)
		return
	}

	resposta, err := c.service.Registrar(requisicao)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, resposta)
}
