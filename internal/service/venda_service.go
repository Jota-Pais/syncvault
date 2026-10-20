// CAMADA: Service
// RESPONSABILIDADE: registrar vendas de qualquer canal no estoque único.
// REGRAS: R1 (a venda não passa do estoque do produto),
//         R2 (o preço, descontada a taxa do canal, não fica abaixo do custo) e
//         R3 (a venda não deixa o estoque abaixo do estoque de segurança do canal).
// ROTAS ATENDIDAS: POST /vendas.
// SITUAÇÃO: montada. A lógica entra quando POST /vendas for implementada.

package service

import (
	"syncvault/internal/mapper"
	"syncvault/internal/repository"
)

type VendaService struct {
	// Acesso à tabela vendas.
	repository *repository.VendaRepository
	// Usado para buscar o produto pelo id e descontar o estoque (R1 e R3).
	produtoRepository *repository.ProdutoRepository
	// Usado para buscar o canal pelo id e ler a taxa e o estoque de segurança (R2 e R3).
	canalRepository *repository.CanalRepository
	// Monta a resposta, com produto e canal completos.
	mapper *mapper.VendaMapper
}

func NovoVendaService(repository *repository.VendaRepository, produtoRepository *repository.ProdutoRepository,
	canalRepository *repository.CanalRepository, mapper *mapper.VendaMapper) *VendaService {
	return &VendaService{
		repository:        repository,
		produtoRepository: produtoRepository,
		canalRepository:   canalRepository,
		mapper:            mapper,
	}
}
