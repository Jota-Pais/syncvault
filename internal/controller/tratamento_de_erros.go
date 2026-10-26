// CAMADA: Controller
// RESPONSABILIDADE: transformar cada erro das rotas em 400, 404 ou 500 no formato ErroDTO.
// SITUAÇÃO: implementada.

package controller

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"syncvault/internal/dto"
	"syncvault/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func TratarErros() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		if len(ctx.Errors) == 0 || ctx.Writer.Written() {
			return
		}
		err := ctx.Errors.Last().Err

		var naoEncontrado *service.RecursoNaoEncontradoError
		var regra *service.RegraDeNegocioError
		var parametro *ParametroInvalidoError
		var validacao validator.ValidationErrors
		var sintaxe *json.SyntaxError
		var tipo *json.UnmarshalTypeError

		status, mensagem := http.StatusInternalServerError, "Erro interno no servidor"
		switch {
		case errors.As(err, &naoEncontrado):
			status, mensagem = http.StatusNotFound, err.Error()
		case errors.As(err, &regra), errors.As(err, &parametro):
			status, mensagem = http.StatusBadRequest, err.Error()
		case errors.As(err, &validacao):
			status, mensagem = http.StatusBadRequest, "Dados inválidos"
		case errors.As(err, &sintaxe), errors.As(err, &tipo), errors.Is(err, io.EOF), errors.Is(err, io.ErrUnexpectedEOF):
			status, mensagem = http.StatusBadRequest, "JSON inválido"
		}

		resposta := dto.NovoErro(status, mensagem, ctx.Request.URL.Path)
		for _, campo := range validacao {
			nome := strings.ToLower(campo.Field()[:1]) + campo.Field()[1:]
			resposta.Campos = append(resposta.Campos, dto.CampoErroDTO{Campo: nome, Mensagem: nome + " inválido (" + campo.Tag() + ")"})
		}
		ctx.JSON(status, resposta)
	}
}
