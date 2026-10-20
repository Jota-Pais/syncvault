// CAMADA: Service
// RESPONSABILIDADE: listar os canais de venda, para quem usa a API descobrir o
//                   id de cada um.
// REGRAS: nenhuma. Taxa e estoque de segurança são lidos pelo VendaService.
// ROTAS ATENDIDAS: GET /canais.
// SITUAÇÃO: montada. A lógica entra quando GET /canais for implementada.

package service

import (
	"syncvault/internal/mapper"
	"syncvault/internal/repository"
)

type CanalService struct {
	// Acesso à tabela canais.
	repository *repository.CanalRepository
	// Converte a entidade em DTO.
	mapper *mapper.CanalMapper
}

func NovoCanalService(repository *repository.CanalRepository, mapper *mapper.CanalMapper) *CanalService {
	return &CanalService{repository: repository, mapper: mapper}
}
