// CAMADA: Service
// RESPONSABILIDADE: cadastrar, listar e buscar produtos.
// REGRAS: nenhuma. O estoque só muda pela venda, no VendaService.
// ROTAS ATENDIDAS: POST /produtos, GET /produtos e GET /produtos/{id}.
// SITUAÇÃO: montada. A lógica entra quando a API simples for implementada.

package service

import (
	"syncvault/internal/mapper"
	"syncvault/internal/repository"
)

type ProdutoService struct {
	// Acesso à tabela produtos.
	repository *repository.ProdutoRepository
	// Converte entre DTO e entidade.
	mapper *mapper.ProdutoMapper
}

func NovoProdutoService(repository *repository.ProdutoRepository, mapper *mapper.ProdutoMapper) *ProdutoService {
	return &ProdutoService{repository: repository, mapper: mapper}
}
