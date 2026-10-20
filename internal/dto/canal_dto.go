// CAMADA: DTO
// RESPONSABILIDADE: o JSON que entra e o que sai nas rotas de /canais.
// SITUAÇÃO: implementada.

package dto

// CanalRequestDTO é o que o cliente enviaria para cadastrar um canal. Nenhuma
// rota usa nesta etapa: os canais nascem na migration 4.
type CanalRequestDTO struct {
	Nome string `json:"nome" binding:"required,max=100"`
	// Ponteiros para separar "não veio" (nil) de "veio zero": 0 é válido nos dois.
	TaxaPercentual   *int `json:"taxaPercentual" binding:"required,gte=0,lte=100"`
	EstoqueSeguranca *int `json:"estoqueSeguranca" binding:"required,gte=0"`
}

// CanalResponseDTO é o canal como a API devolve.
type CanalResponseDTO struct {
	ID               int64  `json:"id"`
	Nome             string `json:"nome"`
	TaxaPercentual   int    `json:"taxaPercentual"`
	EstoqueSeguranca int    `json:"estoqueSeguranca"`
}
