// CAMADA: DTO
// RESPONSABILIDADE: o JSON que entra e o que sai nas rotas de /produtos.
// SITUAÇÃO: implementada.

package dto

// ProdutoRequestDTO é o que o cliente envia para cadastrar um produto. As tags
// binding são as validações: o Gin confere o JSON antes de o Controller usá-lo.
type ProdutoRequestDTO struct {
	Nome          string `json:"nome" binding:"required,max=255"`
	CustoCentavos int64  `json:"custoCentavos" binding:"required,gt=0"`
	// Ponteiro para separar "não veio" (nil) de "veio zero": estoque 0 é válido.
	Estoque *int `json:"estoque" binding:"required,gte=0"`
}

// ProdutoResponseDTO é o produto como a API devolve.
type ProdutoResponseDTO struct {
	ID            int64  `json:"id"`
	Nome          string `json:"nome"`
	CustoCentavos int64  `json:"custoCentavos"`
	Estoque       int    `json:"estoque"`
}
