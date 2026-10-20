// CAMADA: Repository
// RESPONSABILIDADE: ler e gravar produtos no banco. Só o Service chama.
// SITUAÇÃO: implementada.

package repository

import (
	"syncvault/internal/model"

	"gorm.io/gorm"
)

// ProdutoRepository é a única porta para a tabela produtos.
type ProdutoRepository struct {
	db *gorm.DB
}

func NovoProdutoRepository(db *gorm.DB) *ProdutoRepository {
	return &ProdutoRepository{db: db}
}

// Salvar grava um produto novo e preenche o ID que o banco gerou.
func (r *ProdutoRepository) Salvar(produto *model.Produto) error {
	return r.db.Create(produto).Error
}

// ListarTodos devolve todos os produtos, na ordem em que foram cadastrados.
func (r *ProdutoRepository) ListarTodos() ([]model.Produto, error) {
	var produtos []model.Produto
	err := r.db.Order("id").Find(&produtos).Error
	return produtos, err
}

// BuscarPorID devolve o produto com esse id, ou nil se ele não existir.
// Quem decide o que fazer quando não existe é o Service.
func (r *ProdutoRepository) BuscarPorID(id int64) (*model.Produto, error) {
	var produto model.Produto
	resultado := r.db.Limit(1).Find(&produto, id)
	if resultado.Error != nil {
		return nil, resultado.Error
	}
	if resultado.RowsAffected == 0 {
		return nil, nil
	}
	return &produto, nil
}
