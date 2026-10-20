// CAMADA: Mapper
// RESPONSABILIDADE: converter ProdutoRequestDTO em Produto, e Produto em
//                   ProdutoResponseDTO. Não consulta o banco e não tem regra.
// SITUAÇÃO: implementada.

package mapper

import (
	"syncvault/internal/dto"
	"syncvault/internal/model"
)

// ProdutoMapper não guarda estado: só converte.
type ProdutoMapper struct{}

func NovoProdutoMapper() *ProdutoMapper {
	return &ProdutoMapper{}
}

// ParaEntidade cria um produto novo, ainda sem id, a partir do que o cliente enviou.
func (m *ProdutoMapper) ParaEntidade(requisicao dto.ProdutoRequestDTO) model.Produto {
	return model.Produto{
		Nome:          requisicao.Nome,
		CustoCentavos: requisicao.CustoCentavos,
		Estoque:       *requisicao.Estoque,
	}
}

// ParaResposta converte o produto no JSON que a API devolve.
func (m *ProdutoMapper) ParaResposta(produto model.Produto) dto.ProdutoResponseDTO {
	return dto.ProdutoResponseDTO{
		ID:            produto.ID,
		Nome:          produto.Nome,
		CustoCentavos: produto.CustoCentavos,
		Estoque:       produto.Estoque,
	}
}

// ParaRespostaLista converte uma lista de produtos. Começa com uma lista vazia,
// e não com nil, para o JSON sair [] quando não há produto, e não null.
func (m *ProdutoMapper) ParaRespostaLista(produtos []model.Produto) []dto.ProdutoResponseDTO {
	respostas := make([]dto.ProdutoResponseDTO, 0, len(produtos))
	for _, produto := range produtos {
		respostas = append(respostas, m.ParaResposta(produto))
	}
	return respostas
}
