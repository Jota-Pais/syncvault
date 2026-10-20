// CAMADA: Mapper
// RESPONSABILIDADE: converter ProdutoRequestDTO em Produto, e Produto em
//                   ProdutoResponseDTO. Não consulta o banco e não tem regra.
// SITUAÇÃO: montada. Os métodos entram quando as rotas de /produtos forem implementadas.

package mapper

// ProdutoMapper não guarda estado: só converte.
type ProdutoMapper struct{}

func NovoProdutoMapper() *ProdutoMapper {
	return &ProdutoMapper{}
}
