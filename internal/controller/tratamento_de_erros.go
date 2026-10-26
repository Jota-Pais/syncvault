// CAMADA: Controller
// RESPONSABILIDADE: transformar cada erro das rotas em 400, 404 ou 500 no formato ErroDTO.
// SITUAÇÃO: implementada.

package controller

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"reflect"
	"strings"

	"syncvault/internal/dto"
	"syncvault/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func TratarErros() gin.HandlerFunc {
	// Faz o validador chamar cada campo pelo nome do JSON (canalId), e não pelo da struct (CanalID).
	if validacao, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validacao.RegisterTagNameFunc(func(campo reflect.StructField) string {
			return strings.Split(campo.Tag.Get("json"), ",")[0]
		})
	}

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
			nome := campo.Field()
			resposta.Campos = append(resposta.Campos, dto.CampoErroDTO{Campo: nome, Mensagem: mensagemDoCampo(nome, campo)})
		}
		ctx.JSON(status, resposta)
	}
}

// mensagemDoCampo escreve em português a validação (tag binding do DTO) que o campo não passou.
func mensagemDoCampo(nome string, campo validator.FieldError) string {
	switch campo.Tag() {
	case "required":
		// Número que não é ponteiro chega como 0 quando falta, e todos eles
		// também pedem gt=0: a mensagem já diz as duas coisas.
		if campo.Kind() >= reflect.Int && campo.Kind() <= reflect.Float64 {
			return nome + " é obrigatório e deve ser maior que 0"
		}
		return nome + " é obrigatório"
	case "gt":
		return nome + " deve ser maior que " + campo.Param()
	case "gte":
		return nome + " deve ser maior ou igual a " + campo.Param()
	case "lte":
		return nome + " deve ser menor ou igual a " + campo.Param()
	case "max":
		return nome + " deve ter no máximo " + campo.Param() + " caracteres"
	}
	return nome + " inválido"
}
