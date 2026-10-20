// CAMADA: Service
// RESPONSABILIDADE: registrar vendas de qualquer canal no estoque único.
// REGRAS: R1 (a venda não passa do estoque do produto),
//         R2 (o preço, descontada a taxa do canal, não fica abaixo do custo) e
//         R3 (a venda não deixa o estoque abaixo do estoque de segurança do canal).
//         Conferidas nesta ordem: R2, R1 e R3.
// ROTAS ATENDIDAS: POST /vendas.
// SITUAÇÃO: implementada.

package service

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"syncvault/internal/dto"
	"syncvault/internal/mapper"
	"syncvault/internal/model"
	"syncvault/internal/repository"
)

type VendaService struct {
	// Acesso à tabela vendas.
	repository *repository.VendaRepository
	// Usado para buscar o produto pelo id (R1 e R3).
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

// Registrar confere as regras, grava a venda e desconta o estoque único.
// Produto ou canal que não existe vira 404; regra violada vira 400.
func (s *VendaService) Registrar(requisicao dto.VendaRequestDTO) (dto.VendaResponseDTO, error) {
	produto, err := s.produtoRepository.BuscarPorID(requisicao.ProdutoID)
	if err != nil {
		return dto.VendaResponseDTO{}, err
	}
	if produto == nil {
		return dto.VendaResponseDTO{}, ProdutoNaoEncontrado(requisicao.ProdutoID)
	}

	canal, err := s.canalRepository.BuscarPorID(requisicao.CanalID)
	if err != nil {
		return dto.VendaResponseDTO{}, err
	}
	if canal == nil {
		return dto.VendaResponseDTO{}, CanalNaoEncontrado(requisicao.CanalID)
	}

	err = conferirR2(*produto, *canal, requisicao.PrecoUnitarioCentavos)
	if err != nil {
		return dto.VendaResponseDTO{}, err
	}
	err = conferirEstoque(*produto, *canal, requisicao.Quantidade)
	if err != nil {
		return dto.VendaResponseDTO{}, err
	}

	venda := s.mapper.ParaEntidade(requisicao)
	venda.VendidaEm = time.Now().Truncate(time.Second)

	// R1 e R3 de novo, agora dentro do banco: o desconto só acontece se sobrar
	// pelo menos o estoque de segurança do canal.
	descontou, err := s.repository.SalvarDescontandoEstoque(&venda, canal.EstoqueSeguranca)
	if err != nil {
		return dto.VendaResponseDTO{}, err
	}
	if !descontou {
		// Outra venda levou o estoque entre a conferência acima e o desconto.
		// Confere de novo com o estoque de agora, para a mensagem dizer a regra certa.
		atual, err := s.produtoRepository.BuscarPorID(produto.ID)
		if err != nil {
			return dto.VendaResponseDTO{}, err
		}
		err = conferirEstoque(*atual, *canal, requisicao.Quantidade)
		if err != nil {
			return dto.VendaResponseDTO{}, err
		}
		return dto.VendaResponseDTO{}, errors.New("o estoque mudou durante a venda; tente de novo")
	}

	// Lê o produto de novo para a resposta mostrar o estoque depois da venda.
	produto, err = s.produtoRepository.BuscarPorID(produto.ID)
	if err != nil {
		return dto.VendaResponseDTO{}, err
	}

	recebido := recebidoCentavos(venda.PrecoUnitarioCentavos, venda.Quantidade, canal.TaxaPercentual)
	return s.mapper.ParaResposta(venda, *produto, *canal, recebido), nil
}

// conferirR2: descontada a taxa do canal, o preço não pode ficar abaixo do custo,
// senão a loja vende no prejuízo. A conta é feita em inteiros, sem arredondar:
// preço × (100 − taxa) precisa ser pelo menos custo × 100.
func conferirR2(produto model.Produto, canal model.Canal, precoCentavos int64) error {
	if precoCentavos*int64(100-canal.TaxaPercentual) >= produto.CustoCentavos*100 {
		return nil
	}
	return &RegraDeNegocioError{Mensagem: fmt.Sprintf(
		"R2: vender a %s no canal %s, com taxa de %d%%, rende %s por unidade, abaixo do custo de %s",
		reais(precoCentavos), canal.Nome, canal.TaxaPercentual,
		reais(recebidoCentavos(precoCentavos, 1, canal.TaxaPercentual)), reais(produto.CustoCentavos))}
}

// conferirEstoque aplica as duas regras de estoque, nesta ordem:
// R1: a quantidade não pode passar do estoque do produto.
// R3: a venda não pode deixar o estoque abaixo do estoque de segurança do canal.
func conferirEstoque(produto model.Produto, canal model.Canal, quantidade int) error {
	if quantidade > produto.Estoque {
		return &RegraDeNegocioError{Mensagem: fmt.Sprintf(
			"R1: estoque insuficiente: há %s do produto %d e a venda pede %d",
			unidades(produto.Estoque), produto.ID, quantidade)}
	}

	if produto.Estoque-quantidade < canal.EstoqueSeguranca {
		maximo := produto.Estoque - canal.EstoqueSeguranca
		podeVender := "então ele não pode vender nenhuma"
		if maximo > 0 {
			podeVender = fmt.Sprintf("então ele vende no máximo %d", maximo)
		}
		return &RegraDeNegocioError{Mensagem: fmt.Sprintf(
			"R3: o canal %s não vende as últimas %s: há %d em estoque, %s",
			canal.Nome, unidades(canal.EstoqueSeguranca), produto.Estoque, podeVender)}
	}
	return nil
}

// recebidoCentavos é quanto a loja recebe depois da comissão do canal,
// arredondado para baixo.
func recebidoCentavos(precoCentavos int64, quantidade int, taxaPercentual int) int64 {
	return precoCentavos * int64(quantidade) * int64(100-taxaPercentual) / 100
}

// reais formata centavos como dinheiro: 1234567 vira "R$ 12.345,67".
func reais(centavos int64) string {
	inteiro := strconv.FormatInt(centavos/100, 10)
	for i := len(inteiro) - 3; i > 0; i -= 3 {
		inteiro = inteiro[:i] + "." + inteiro[i:]
	}
	return fmt.Sprintf("R$ %s,%02d", inteiro, centavos%100)
}

// unidades escreve a quantidade com a palavra no singular ou no plural.
func unidades(n int) string {
	if n == 1 {
		return "1 unidade"
	}
	return fmt.Sprintf("%d unidades", n)
}
