// CAMADA: Model
// RESPONSABILIDADE: o produto à venda, com um estoque só para todos os canais.
//                   Bate com a tabela produtos da migration 1.
// SITUAÇÃO: implementada.

package model

// Produto é uma linha da tabela produtos. As tags gorm dizem em qual coluna
// fica cada campo; quem cria a tabela é a migration, nunca o GORM.
type Produto struct {
	ID            int64  `gorm:"column:id;primaryKey"`
	Nome          string `gorm:"column:nome"`
	CustoCentavos int64  `gorm:"column:custo_centavos"`
	Estoque       int    `gorm:"column:estoque"`
}

// TableName diz ao GORM o nome da tabela, como o @Table do Java.
func (Produto) TableName() string {
	return "produtos"
}
