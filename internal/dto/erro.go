// CAMADA: DTO
// RESPONSABILIDADE: o formato único de erro da API, o mesmo no 400, no 404 e no 500.
// SITUAÇÃO: implementada.

package dto

import (
	"net/http"
	"time"
)

// CampoErroDTO é um campo inválido: qual é e o que há de errado nele.
type CampoErroDTO struct {
	Campo    string `json:"campo" example:"nome"`
	Mensagem string `json:"mensagem" example:"nome é obrigatório"`
}

// ErroDTO é o formato de todo erro da API. "campos" só tem itens no 400 de validação.
type ErroDTO struct {
	Status    int            `json:"status" example:"404"`
	Erro      string         `json:"erro" example:"Not Found"`
	Mensagem  string         `json:"mensagem" example:"Produto não encontrado: 999"`
	Caminho   string         `json:"caminho" example:"/produtos/999"`
	Timestamp time.Time      `json:"timestamp" example:"2026-10-20T19:30:00-03:00"`
	Campos    []CampoErroDTO `json:"campos"`
}

// NovoErro monta um ErroDTO com o status, a mensagem e o caminho da requisição.
func NovoErro(status int, mensagem, caminho string) ErroDTO {
	return ErroDTO{
		Status:    status,
		Erro:      http.StatusText(status),
		Mensagem:  mensagem,
		Caminho:   caminho,
		Timestamp: time.Now().Truncate(time.Second),
		// Lista vazia em vez de nil, para o JSON sair "campos": [] e não null.
		Campos: []CampoErroDTO{},
	}
}
