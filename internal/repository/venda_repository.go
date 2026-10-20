// CAMADA: Repository
// RESPONSABILIDADE: ler e gravar vendas no banco. Só o Service chama.
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

// Salvar grava uma venda nova e preenche o ID que o banco gerou.
func (r *VendaRepository) Salvar(venda *model.Venda) error {
	return r.db.Create(venda).Error
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
