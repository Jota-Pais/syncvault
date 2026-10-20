# controller: a porta de entrada da API

É aqui que a requisição HTTP chega. O Controller recebe o pedido, repassa para o Service e devolve a resposta.

Cada entidade tem o seu Controller, com o método `RegistrarRotas`, que liga cada verbo e caminho a um método, e um método por rota.

| Faz | Não faz |
|---|---|
| Liga verbo e caminho a um método (`RegistrarRotas`, o papel do `@GetMapping` e do `@PostMapping`) | Regra de negócio |
| Lê o JSON e dispara a validação do DTO de entrada (`ShouldBindJSON`, o papel do `@Valid`) | Acessar o Repository |
| Escolhe o status HTTP de sucesso (`200`, `201`) | Importar qualquer coisa de `internal/model` |
| Descreve a rota para o Swagger (comentários `@Summary`, `@Router`...) | Montar a resposta de erro: entrega o erro com `ctx.Error(err)` e quem responde é o `TratarErros` |

| Arquivo | Rotas |
|---|---|
| `produto_controller.go` | `POST /produtos`, `GET /produtos` e `GET /produtos/{id}` |
| `canal_controller.go` | `GET /canais` |
| `venda_controller.go` | `POST /vendas` |
| `parametros.go` | Lê o `{id}` do caminho; um id que não é número vira `400` |
| `tratamento_de_erros.go` | Em construção. `TratarErros`, que transforma cada erro em `400` ou `404` no formato `ErroDTO` (o papel do `ApiExceptionHandler`) |
