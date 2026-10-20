# internal: o código do projeto

É o equivalente ao `src/main/java/api` do template da disciplina. Em Go, o que está dentro de `internal/` só pode ser importado pelo próprio projeto.

## As pastas

Uma pasta por camada, e cada pasta é um pacote Go. Cada uma tem um `README.md` explicando o que mora ali.

| Pasta | O que mora ali | Um arquivo por entidade |
|---|---|---|
| [controller/](controller) | A porta de entrada da API | `produto_controller.go` → `ProdutoController` |
| [service/](service) | As regras do sistema | `produto_service.go` → `ProdutoService` |
| [mapper/](mapper) | A conversão entre DTO e entidade | `produto_mapper.go` → `ProdutoMapper` |
| [repository/](repository) | A conversa com o banco | `produto_repository.go` → `ProdutoRepository` |
| [model/](model) | As entidades | `produto.go` → `Produto` |
| [dto/](dto) | O JSON que entra e o que sai | `produto_dto.go` → `ProdutoRequestDTO` e `ProdutoResponseDTO` |
| [config/](config) | Conexão, migrations, conferência das entidades e Swagger | Não se aplica |

As migrations ficam fora daqui, em [`migrations/`](../migrations), na raiz.

## Como as camadas se ligam

Em Java, o Spring cria cada classe e entrega as dependências pelo construtor. Em Go, quem faz isso é o `main.go`, com as funções `NovoXxx`, de baixo para cima:

1. Repository e Mapper, que só dependem do banco ou de nada;
2. Service, que recebe o seu Repository e o seu Mapper;
3. Controller, que recebe o seu Service e registra as rotas com `RegistrarRotas`.

Se faltar uma peça, o projeto nem compila. Se uma entidade não bater com a sua tabela, a API não sobe (`config.ValidarEntidades`).

## O comentário no topo de cada arquivo

Todo arquivo que o grupo cria começa com um comentário dizendo o que ele é e o que faz. É a documentação da estrutura.

```go
// CAMADA: Service
// RESPONSABILIDADE: registrar vendas de qualquer canal no estoque único.
// REGRAS: R1 (a venda não passa do estoque do produto),
//         R2 (o preço, descontada a taxa do canal, não fica abaixo do custo) e
//         R3 (a venda não deixa o estoque abaixo do estoque de segurança do canal).
// ROTAS ATENDIDAS: POST /vendas.
// SITUAÇÃO: implementada.

package service
```

| Linha | O que escrever | Em quais arquivos |
|---|---|---|
| `CAMADA` | Controller, Service, Mapper, Repository, Model, DTO ou Config | Todos |
| `RESPONSABILIDADE` | Uma frase: o que este arquivo faz no projeto | Todos |
| `REGRAS` | As regras do README que ele aplica | Service |
| `ROTAS ATENDIDAS` | As rotas de `docs/contrato-da-api.md` que ele atende | Controller e Service |
| `SITUAÇÃO` | `montada` ou `implementada` | Todos |

A linha em branco entre o comentário e o `package` é proposital: sem ela, o Go trataria o cabeçalho como a documentação do pacote.
