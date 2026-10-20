// CAMADA: Model
// RESPONSABILIDADE: um lugar onde a loja vende (balcão ou marketplace), com a
//                   comissão e o estoque de segurança dele. Bate com a tabela
//                   canais da migration 2.
// SITUAÇÃO: implementada.

package model

// Canal é uma linha da tabela canais. Os três canais nascem na migration 4.
type Canal struct {
	ID               int64  `gorm:"column:id;primaryKey"`
	Nome             string `gorm:"column:nome"`
	TaxaPercentual   int    `gorm:"column:taxa_percentual"`
	EstoqueSeguranca int    `gorm:"column:estoque_seguranca"`
}

// TableName diz ao GORM o nome da tabela. Aqui é obrigatório: sozinho, o GORM
// pluralizaria Canal como "canals".
func (Canal) TableName() string {
	return "canais"
}
