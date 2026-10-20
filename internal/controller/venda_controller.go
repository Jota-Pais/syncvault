// CAMADA: Controller
// RESPONSABILIDADE: receber as requisições de /vendas e devolver status e corpo.
//                   As regras R1, R2 e R3 não moram aqui: moram no VendaService.
// ROTAS ATENDIDAS: POST /vendas.
// SITUAÇÃO: montada. A rota entra quando POST /vendas for implementada.

package controller

import (
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
	// Montada: POST /vendas entra aqui quando for implementada.
}
