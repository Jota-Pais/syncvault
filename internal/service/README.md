# service: as regras do sistema

É aqui que o sistema decide o que pode e o que não pode acontecer. O Service recebe o pedido do Controller, aplica as regras de negócio e usa o Repository para ler e gravar.

**As três regras do README moram no `VendaService`**, em `venda_service.go`: `conferirR2` e `conferirEstoque` (R1 e R3), chamadas pelo método `Registrar`.

| Faz | Não faz |
|---|---|
| Aplica as regras de negócio que estão no README | Conhecer HTTP: nada de status nem de `gin` aqui |
| Recebe DTO, pede a conversão ao Mapper e devolve DTO | Devolver entidade para o Controller |
| Busca as outras entidades de que precisa, pelo id | Devolver vazio quando não acha: devolve o erro de "não encontrado" |

Em Go não há exceção: toda função que pode falhar devolve um `error` como último valor. Os erros do domínio estão em `erros.go`:

| Erro | Quando | Vira | Em Java |
|---|---|---|---|
| `RecursoNaoEncontradoError` | Um id que não existe, criado por `ProdutoNaoEncontrado`, `CanalNaoEncontrado` ou `VendaNaoEncontrada` | `404` | `RecursoNaoEncontradoException` e as filhas de cada entidade |
| `RegraDeNegocioError` | Uma regra do README violada; a mensagem começa pela regra (`R1:`, `R2:`, `R3:`) | `400` | `RegraDeNegocioException` |
