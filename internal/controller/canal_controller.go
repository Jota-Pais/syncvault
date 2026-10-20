// CAMADA: Controller
// RESPONSABILIDADE: receber as requisições de /canais e devolver status e corpo.
// ROTAS ATENDIDAS: GET /canais.
// SITUAÇÃO: montada. A rota entra quando GET /canais for implementada.

package controller

import (
	"syncvault/internal/service"

	"github.com/gin-gonic/gin"
)

type CanalController struct {
	service *service.CanalService
}

func NovoCanalController(service *service.CanalService) *CanalController {
	return &CanalController{service: service}
}

// RegistrarRotas liga cada verbo e caminho de /canais a um método deste controller.
func (c *CanalController) RegistrarRotas(r *gin.Engine) {
	// Montada: GET /canais entra aqui quando for implementada.
}
