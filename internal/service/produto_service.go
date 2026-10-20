// CAMADA: Service
// RESPONSABILIDADE: cadastrar, listar e buscar produtos.
// REGRAS: nenhuma. O estoque só muda pela venda, no VendaService.
// ROTAS ATENDIDAS: POST /produtos, GET /produtos e GET /produtos/{id}.
// SITUAÇÃO: implementada.

package service

import (
	"syncvault/internal/dto"
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

// Criar cadastra um produto com o estoque inicial e devolve o produto com o id gerado.
func (s *ProdutoService) Criar(requisicao dto.ProdutoRequestDTO) (dto.ProdutoResponseDTO, error) {
	produto := s.mapper.ParaEntidade(requisicao)
	err := s.repository.Salvar(&produto)
	if err != nil {
		return dto.ProdutoResponseDTO{}, err
	}
	return s.mapper.ParaResposta(produto), nil
}

// Listar devolve todos os produtos, com o estoque atual.
func (s *ProdutoService) Listar() ([]dto.ProdutoResponseDTO, error) {
	produtos, err := s.repository.ListarTodos()
	if err != nil {
		return nil, err
	}
	return s.mapper.ParaRespostaLista(produtos), nil
}

// BuscarPorID devolve um produto, ou o erro ProdutoNaoEncontrado (404) se o id não existe.
func (s *ProdutoService) BuscarPorID(id int64) (dto.ProdutoResponseDTO, error) {
	produto, err := s.repository.BuscarPorID(id)
	if err != nil {
		return dto.ProdutoResponseDTO{}, err
	}
	if produto == nil {
		return dto.ProdutoResponseDTO{}, ProdutoNaoEncontrado(id)
	}
	return s.mapper.ParaResposta(*produto), nil
}
