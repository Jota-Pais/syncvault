# mapper: a conversão entre DTO e entidade

O Mapper transforma o DTO que chegou em entidade, e a entidade em DTO para devolver.

| Método | O que converte | Usado em |
|---|---|---|
| `ParaEntidade(requisicao)` | DTO de entrada em entidade nova, ainda sem id | `POST` |
| `ParaResposta(entidade)` | Entidade em DTO de saída | Toda resposta |
| `ParaRespostaLista(entidades)` | Lista de entidades em lista de DTOs; sem item, sai `[]`, nunca `null` | `GET` de listagem |

O Mapper não consulta o banco e não tem regra de negócio. Quando a conversão precisa de outra entidade, o Service busca e entrega pronta: o `VendaMapper` recebe o produto e o canal já buscados, e usa o `ProdutoMapper` e o `CanalMapper` para montá-los dentro da resposta. O `recebidoCentavos` também chega pronto, porque depende da taxa do canal (R2), e quem calcula é o `VendaService`.

| Arquivo | Situação |
|---|---|
| `produto_mapper.go` | Implementado |
| `canal_mapper.go` | Em parte: só a conversão para a resposta, porque não há rota de cadastro de canal |
| `venda_mapper.go` | Implementado |
