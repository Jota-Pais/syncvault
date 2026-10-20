// CAMADA: Mapper
// RESPONSABILIDADE: converter Canal em CanalResponseDTO. Não consulta o banco e
//                   não tem regra.
// SITUAÇÃO: implementada na parte que as rotas usam. A conversão de
//           CanalRequestDTO em Canal entra quando houver rota para cadastrar canal.

package mapper

import (
	"syncvault/internal/dto"
	"syncvault/internal/model"
)

// CanalMapper não guarda estado: só converte.
type CanalMapper struct{}

func NovoCanalMapper() *CanalMapper {
	return &CanalMapper{}
}

// ParaResposta converte o canal no JSON que a API devolve.
func (m *CanalMapper) ParaResposta(canal model.Canal) dto.CanalResponseDTO {
	return dto.CanalResponseDTO{
		ID:               canal.ID,
		Nome:             canal.Nome,
		TaxaPercentual:   canal.TaxaPercentual,
		EstoqueSeguranca: canal.EstoqueSeguranca,
	}
}

// ParaRespostaLista converte uma lista de canais; sem canal, o JSON sai [].
func (m *CanalMapper) ParaRespostaLista(canais []model.Canal) []dto.CanalResponseDTO {
	respostas := make([]dto.CanalResponseDTO, 0, len(canais))
	for _, canal := range canais {
		respostas = append(respostas, m.ParaResposta(canal))
	}
	return respostas
}
