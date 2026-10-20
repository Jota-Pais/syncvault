// CAMADA: Service
// RESPONSABILIDADE: os erros do domínio que viram resposta HTTP. Um id que não
//                   existe vira 404; uma regra de negócio violada vira 400. Quem
//                   transforma cada um no ErroDTO é o tratamento central de erros.
// SITUAÇÃO: implementada.

package service

import "fmt"

// RecursoNaoEncontradoError é o erro de um id que não existe: vira 404.
// Cada entidade tem a sua função que cria este erro com a mensagem certa.
type RecursoNaoEncontradoError struct {
	Mensagem string
}

func (e *RecursoNaoEncontradoError) Error() string {
	return e.Mensagem
}

// RegraDeNegocioError é uma regra do README violada: vira 400, e a mensagem
// começa pela regra (R1, R2 ou R3).
type RegraDeNegocioError struct {
	Mensagem string
}

func (e *RegraDeNegocioError) Error() string {
	return e.Mensagem
}

// ProdutoNaoEncontrado é o "não encontrado" próprio de produto.
func ProdutoNaoEncontrado(id int64) *RecursoNaoEncontradoError {
	return &RecursoNaoEncontradoError{Mensagem: fmt.Sprintf("Produto não encontrado: %d", id)}
}

// CanalNaoEncontrado é o "não encontrado" próprio de canal.
func CanalNaoEncontrado(id int64) *RecursoNaoEncontradoError {
	return &RecursoNaoEncontradoError{Mensagem: fmt.Sprintf("Canal não encontrado: %d", id)}
}

// VendaNaoEncontrada é o "não encontrado" próprio de venda.
func VendaNaoEncontrada(id int64) *RecursoNaoEncontradoError {
	return &RecursoNaoEncontradoError{Mensagem: fmt.Sprintf("Venda não encontrada: %d", id)}
}
