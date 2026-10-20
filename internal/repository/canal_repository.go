// CAMADA: Repository
// RESPONSABILIDADE: ler e gravar canais no banco. Só o Service chama.
// SITUAÇÃO: implementada.

package repository

import (
	"syncvault/internal/model"

	"gorm.io/gorm"
)

// CanalRepository é a única porta para a tabela canais.
type CanalRepository struct {
	db *gorm.DB
}

func NovoCanalRepository(db *gorm.DB) *CanalRepository {
	return &CanalRepository{db: db}
}

// Salvar grava um canal novo e preenche o ID que o banco gerou. Nenhuma rota
// usa nesta etapa: os canais nascem na migration 4.
func (r *CanalRepository) Salvar(canal *model.Canal) error {
	return r.db.Create(canal).Error
}

// ListarTodos devolve todos os canais, na ordem em que foram cadastrados.
func (r *CanalRepository) ListarTodos() ([]model.Canal, error) {
	var canais []model.Canal
	err := r.db.Order("id").Find(&canais).Error
	return canais, err
}

// BuscarPorID devolve o canal com esse id, ou nil se ele não existir.
func (r *CanalRepository) BuscarPorID(id int64) (*model.Canal, error) {
	var canal model.Canal
	resultado := r.db.Limit(1).Find(&canal, id)
	if resultado.Error != nil {
		return nil, resultado.Error
	}
	if resultado.RowsAffected == 0 {
		return nil, nil
	}
	return &canal, nil
}
