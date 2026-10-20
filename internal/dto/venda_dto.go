// CAMADA: DTO
// RESPONSABILIDADE: o JSON que entra e o que sai nas rotas de /vendas.
// SITUAÇÃO: implementada.

package dto

import "time"

// VendaRequestDTO é o que o cliente envia para registrar uma venda. Produto e
// canal entram só como id.
type VendaRequestDTO struct {
	ProdutoID  int64 `json:"produtoId" binding:"required,gt=0" example:"1"`
	CanalID    int64 `json:"canalId" binding:"required,gt=0" example:"2"`
	Quantidade int   `json:"quantidade" binding:"required,gt=0" example:"1"`
	// O preço praticado neste canal, por unidade, em centavos (R$ 109,90 = 10990).
	PrecoUnitarioCentavos int64 `json:"precoUnitarioCentavos" binding:"required,gt=0" example:"10990"`
}

// VendaResponseDTO é a venda como a API devolve: produto e canal saem
// completos, e o produto já com o estoque depois da venda.
type VendaResponseDTO struct {
	ID                    int64              `json:"id" example:"1"`
	Produto               ProdutoResponseDTO `json:"produto"`
	Canal                 CanalResponseDTO   `json:"canal"`
	Quantidade            int                `json:"quantidade" example:"1"`
	PrecoUnitarioCentavos int64              `json:"precoUnitarioCentavos" example:"10990"`
	// Quanto a loja recebe depois da comissão do canal. Quem calcula é o Service.
	RecebidoCentavos int64     `json:"recebidoCentavos" example:"9231"`
	VendidaEm        time.Time `json:"vendidaEm" example:"2026-10-20T19:30:00-03:00"`
}
