# dto: o JSON que entra e o que sai

Um DTO descreve o formato do JSON de uma rota. É o que o cliente da API enxerga; a entidade fica escondida atrás dele.

Cada entidade tem um arquivo com os seus dois DTOs. Em Java seriam dois `record`; em Go são duas `struct`.

| DTO | Para que serve | Exemplo |
|---|---|---|
| `...RequestDTO` | O que o cliente pode enviar. As validações ficam nas tags `binding` (`required`, `gt=0`, `max=255`...), conferidas pelo Gin | `ProdutoRequestDTO` |
| `...ResponseDTO` | O que a API devolve. Inclui o `id` e os campos que a API calcula, como o `recebidoCentavos` da venda | `ProdutoResponseDTO` |

- A tag `json` dá o nome do campo no JSON (`custoCentavos`); a tag `example` aparece no Swagger.
- Nenhuma tag `gorm` aqui, e nenhuma entidade como campo de um DTO.
- Relacionamento entra como id (`"produtoId": 1`) e sai como objeto (`"produto": {...}`).
- Campo numérico em que zero é válido, como o estoque, é ponteiro (`*int`): assim dá para separar "não veio" de "veio zero".

| Arquivo | DTOs |
|---|---|
| `produto_dto.go` | `ProdutoRequestDTO` e `ProdutoResponseDTO` |
| `canal_dto.go` | `CanalRequestDTO` (nenhuma rota usa nesta etapa) e `CanalResponseDTO` |
| `venda_dto.go` | `VendaRequestDTO` e `VendaResponseDTO` |
| `erro.go` | Em construção (issue #1). `ErroDTO` e `CampoErroDTO`, o formato único de erro |
