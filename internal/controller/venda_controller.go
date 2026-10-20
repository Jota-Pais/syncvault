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
