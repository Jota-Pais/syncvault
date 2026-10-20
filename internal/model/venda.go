// CAMADA: Model
// RESPONSABILIDADE: a venda de um produto, registrada por um canal. Bate com a
//                   tabela vendas da migration 3.
// SITUAÇÃO: implementada.

package model

import "time"

// Venda é uma linha da tabela vendas. ProdutoID e CanalID são as chaves
// estrangeiras: cada venda é de um produto e de um canal (muitos para um).
type Venda struct {
	ID                    int64     `gorm:"column:id;primaryKey"`
	ProdutoID             int64     `gorm:"column:produto_id"`
	CanalID               int64     `gorm:"column:canal_id"`
	Quantidade            int       `gorm:"column:quantidade"`
	PrecoUnitarioCentavos int64     `gorm:"column:preco_unitario_centavos"`
	VendidaEm             time.Time `gorm:"column:vendida_em"`
}

// TableName diz ao GORM o nome da tabela, como o @Table do Java.
func (Venda) TableName() string {
	return "vendas"
}
