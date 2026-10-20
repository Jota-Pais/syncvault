// CAMADA: Repository
// RESPONSABILIDADE: ler e gravar vendas no banco. Gravar uma venda também
//                   desconta o estoque do produto, na mesma transação. Só o
//                   Service chama.
// SITUAÇÃO: implementada.

package repository

import (
	"syncvault/internal/model"

	"gorm.io/gorm"
)

// VendaRepository é a única porta para a tabela vendas.
type VendaRepository struct {
	db *gorm.DB
}

func NovoVendaRepository(db *gorm.DB) *VendaRepository {
	return &VendaRepository{db: db}
}

// Salvar grava uma venda nova e preenche o ID que o banco gerou, sem mexer no estoque.
func (r *VendaRepository) Salvar(venda *model.Venda) error {
	return r.db.Create(venda).Error
}

// SalvarDescontandoEstoque grava a venda e desconta a quantidade do estoque do
// produto numa transação: ou as duas coisas acontecem, ou nenhuma.
//
// O desconto só acontece se, depois dele, sobrarem pelo menos minimoQueSobra
// unidades. A conferência e o desconto são o mesmo UPDATE, então duas vendas
// chegando ao mesmo tempo não conseguem levar a mesma última peça: o banco
// executa uma de cada vez. Devolve false, sem gravar nada, se não havia estoque.
func (r *VendaRepository) SalvarDescontandoEstoque(venda *model.Venda, minimoQueSobra int) (bool, error) {
	descontou := false
	err := r.db.Transaction(func(tx *gorm.DB) error {
		resultado := tx.Model(&model.Produto{}).
			Where("id = ? AND estoque - ? >= ?", venda.ProdutoID, venda.Quantidade, minimoQueSobra).
			Update("estoque", gorm.Expr("estoque - ?", venda.Quantidade))
		if resultado.Error != nil {
			return resultado.Error
		}
		if resultado.RowsAffected == 0 {
			return nil
		}
		descontou = true
		return tx.Create(venda).Error
	})
	return descontou, err
}

// ListarTodos devolve todas as vendas, de todos os canais, da mais antiga para a mais nova.
func (r *VendaRepository) ListarTodos() ([]model.Venda, error) {
	var vendas []model.Venda
	err := r.db.Order("id").Find(&vendas).Error
	return vendas, err
}

// BuscarPorID devolve a venda com esse id, ou nil se ela não existir.
func (r *VendaRepository) BuscarPorID(id int64) (*model.Venda, error) {
	var venda model.Venda
	resultado := r.db.Limit(1).Find(&venda, id)
	if resultado.Error != nil {
		return nil, resultado.Error
	}
	if resultado.RowsAffected == 0 {
		return nil, nil
	}
	return &venda, nil
}
