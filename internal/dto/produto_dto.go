// CAMADA: DTO
// RESPONSABILIDADE: o JSON que entra e o que sai nas rotas de /produtos.
// SITUAÇÃO: implementada.

package dto

// ProdutoRequestDTO é o que o cliente envia para cadastrar um produto. As tags
// binding são as validações: o Gin confere o JSON antes de o Controller usá-lo.
// As tags example aparecem no Swagger UI.
type ProdutoRequestDTO struct {
	Nome string `json:"nome" binding:"required,max=255" example:"Casaco corta-vento"`
	// Quanto a loja pagou por uma unidade, em centavos (R$ 60,00 = 6000).
	CustoCentavos int64 `json:"custoCentavos" binding:"required,gt=0" example:"6000"`
	// Ponteiro para separar "não veio" (nil) de "veio zero": estoque 0 é válido.
	Estoque *int `json:"estoque" binding:"required,gte=0" example:"5"`
}

// ProdutoResponseDTO é o produto como a API devolve.
type ProdutoResponseDTO struct {
	ID            int64  `json:"id" example:"1"`
	Nome          string `json:"nome" example:"Casaco corta-vento"`
	CustoCentavos int64  `json:"custoCentavos" example:"6000"`
	// O estoque único: o mesmo para todos os canais.
	Estoque int `json:"estoque" example:"5"`
}
