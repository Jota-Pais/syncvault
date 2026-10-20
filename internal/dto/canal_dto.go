// CAMADA: DTO
// RESPONSABILIDADE: o JSON que entra e o que sai nas rotas de /canais.
// SITUAÇÃO: implementada.

package dto

// CanalRequestDTO é o que o cliente enviaria para cadastrar um canal. Nenhuma
// rota usa nesta etapa: os canais nascem na migration 4.
type CanalRequestDTO struct {
	Nome string `json:"nome" binding:"required,max=100" example:"Mercado Livre"`
	// Ponteiros para separar "não veio" (nil) de "veio zero": 0 é válido nos dois.
	TaxaPercentual   *int `json:"taxaPercentual" binding:"required,gte=0,lte=100" example:"16"`
	EstoqueSeguranca *int `json:"estoqueSeguranca" binding:"required,gte=0" example:"2"`
}

// CanalResponseDTO é o canal como a API devolve.
type CanalResponseDTO struct {
	ID   int64  `json:"id" example:"2"`
	Nome string `json:"nome" example:"Mercado Livre"`
	// A comissão que o canal fica de cada venda, em percentual inteiro (16 = 16%).
	TaxaPercentual int `json:"taxaPercentual" example:"16"`
	// Quantas unidades este canal não pode vender (R3).
	EstoqueSeguranca int `json:"estoqueSeguranca" example:"2"`
}
