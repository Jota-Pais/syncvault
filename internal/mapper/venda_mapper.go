// CAMADA: Mapper
// RESPONSABILIDADE: converter VendaRequestDTO em Venda, e Venda em
//                   VendaResponseDTO, montando o produto e o canal de dentro da
//                   resposta com os outros dois mappers. O recebidoCentavos chega
//                   pronto: quem calcula é o VendaService, porque depende da taxa
//                   do canal (R2), e o mapper não tem regra de negócio.
// SITUAÇÃO: implementada.

package mapper

import (
	"syncvault/internal/dto"
	"syncvault/internal/model"
)

// VendaMapper usa os mappers de produto e de canal para montar a resposta.
type VendaMapper struct {
	produtoMapper *ProdutoMapper
	canalMapper   *CanalMapper
}

func NovoVendaMapper(produtoMapper *ProdutoMapper, canalMapper *CanalMapper) *VendaMapper {
	return &VendaMapper{produtoMapper: produtoMapper, canalMapper: canalMapper}
}

// ParaEntidade cria uma venda nova, ainda sem id e sem o momento da venda,
// que quem preenche é o Service.
func (m *VendaMapper) ParaEntidade(requisicao dto.VendaRequestDTO) model.Venda {
	return model.Venda{
		ProdutoID:             requisicao.ProdutoID,
		CanalID:               requisicao.CanalID,
		Quantidade:            requisicao.Quantidade,
		PrecoUnitarioCentavos: requisicao.PrecoUnitarioCentavos,
	}
}

// ParaResposta monta o JSON da venda, com o produto e o canal completos.
func (m *VendaMapper) ParaResposta(venda model.Venda, produto model.Produto, canal model.Canal,
	recebidoCentavos int64) dto.VendaResponseDTO {
	return dto.VendaResponseDTO{
		ID:                    venda.ID,
		Produto:               m.produtoMapper.ParaResposta(produto),
		Canal:                 m.canalMapper.ParaResposta(canal),
		Quantidade:            venda.Quantidade,
		PrecoUnitarioCentavos: venda.PrecoUnitarioCentavos,
		RecebidoCentavos:      recebidoCentavos,
		VendidaEm:             venda.VendidaEm,
	}
}
