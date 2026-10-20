// CAMADA: Mapper
// RESPONSABILIDADE: converter CanalRequestDTO em Canal, e Canal em
//                   CanalResponseDTO. Não consulta o banco e não tem regra.
// SITUAÇÃO: montada. Os métodos entram quando a rota GET /canais for implementada.

package mapper

// CanalMapper não guarda estado: só converte.
type CanalMapper struct{}

func NovoCanalMapper() *CanalMapper {
	return &CanalMapper{}
}
