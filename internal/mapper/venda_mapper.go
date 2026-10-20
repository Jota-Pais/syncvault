// CAMADA: Mapper
// RESPONSABILIDADE: converter VendaRequestDTO em Venda, e Venda em
//                   VendaResponseDTO, montando o produto e o canal de dentro da
//                   resposta com os outros dois mappers. O recebidoCentavos chega
//                   pronto: quem calcula é o VendaService, porque depende da taxa
//                   do canal (R2), e o mapper não tem regra de negócio.
// SITUAÇÃO: montada. Os métodos entram quando a rota POST /vendas for implementada.

package mapper

// VendaMapper usa os mappers de produto e de canal para montar a resposta.
type VendaMapper struct {
	produtoMapper *ProdutoMapper
	canalMapper   *CanalMapper
}

func NovoVendaMapper(produtoMapper *ProdutoMapper, canalMapper *CanalMapper) *VendaMapper {
	return &VendaMapper{produtoMapper: produtoMapper, canalMapper: canalMapper}
}
