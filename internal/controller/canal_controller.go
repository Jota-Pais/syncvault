// CAMADA: Controller
// RESPONSABILIDADE: receber as requisições de /canais e devolver status e corpo.
// ROTAS ATENDIDAS: GET /canais.
// SITUAÇÃO: implementada.

package controller

import (
	"net/http"

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
	r.GET("/canais", c.Listar)
}

// Listar atende GET /canais.
func (c *CanalController) Listar(ctx *gin.Context) {
	resposta, err := c.service.Listar()
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, resposta)
}
