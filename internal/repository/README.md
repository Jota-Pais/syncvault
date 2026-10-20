# repository: a conversa com o banco

É a única camada que lê e grava no banco. Só o Service chama o Repository.

Em Java, o Spring escreve o Repository a partir de uma interface. Em Go, cada Repository é uma struct que guarda a conexão do GORM e tem um método por operação:

| Método | O que faz | Em Java |
|---|---|---|
| `Salvar(&entidade)` | Grava uma linha nova e preenche o id que o banco gerou | `save` |
| `ListarTodos()` | Devolve todas as linhas, em ordem de id | `findAll` |
| `BuscarPorID(id)` | Devolve a linha, ou `nil` se o id não existe. Quem decide o que fazer com o `nil` é o Service | `findById` |

O `VendaRepository` tem mais um: `SalvarDescontandoEstoque`, que grava a venda e desconta o estoque do produto numa transação. A conferência do estoque e o desconto são o mesmo `UPDATE`, então duas vendas ao mesmo tempo não levam a mesma última peça. O Repository só executa; quem decide quanto precisa sobrar (o estoque de segurança do canal, R3) é o `VendaService`.
