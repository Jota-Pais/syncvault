// CAMADA: Controller
// RESPONSABILIDADE: ler o {id} do caminho das rotas. Um id que não é número
//                   inteiro positivo vira 400.
// SITUAÇÃO: implementada.

package controller

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ParametroInvalidoError é um valor do caminho que não serve, como /produtos/abc: vira 400.
type ParametroInvalidoError struct {
	Mensagem string
}

func (e *ParametroInvalidoError) Error() string {
	return e.Mensagem
}

// lerID converte o {id} do caminho em número.
func lerID(ctx *gin.Context) (int64, error) {
	texto := ctx.Param("id")
	id, err := strconv.ParseInt(texto, 10, 64)
	if err != nil || id <= 0 {
		return 0, &ParametroInvalidoError{Mensagem: fmt.Sprintf("id inválido: %s", texto)}
	}
	return id, nil
}
