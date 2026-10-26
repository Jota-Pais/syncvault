# Autoavaliação

Os 10 critérios que serão avaliados na revisão do código, descritos em `Aula 10-11/AVALIACAO.md`, no repositório da disciplina. Para cada um, a marca do grupo e onde está a evidência. O projeto é em Go: a seção [Go no lugar de Java](../README.md#go-no-lugar-de-java) do README mostra o equivalente de cada nome do Java.

- ✅ atende
- ◐ atende em parte (diga o que falta)
- ❌ não atende

| Critério | Marca | Evidência |
|---|---|---|
| P1 A API sobe seguindo o README | ✅ | `README.md`, seção "Como executar": `git clone` e `docker compose up -d --build` sobem o banco e a API, sem instalar Go. Conferido num clone em pasta vazia |
| P2 Cada classe no pacote da sua camada | ✅ | Uma pasta por camada em `internal/`: `controller`, `service`, `mapper`, `repository`, `model` e `dto`, com os nomes da convenção (`ProdutoController`, `ProdutoService`...) |
| P3 Migrations e `@Entity` de todas as tabelas, e a API sobe | ✅ | `000001` a `000004` em `migrations/` (as chaves estrangeiras estão na `000003_criar_vendas`); `Produto`, `Canal` e `Venda` em `internal/model/`; `config.ValidarEntidades` confere structs e tabelas a cada subida, sem `AutoMigrate` |
| P4 DTOs de entrada e de saída de todas as entidades | ✅ | `produto_dto.go`, `canal_dto.go` e `venda_dto.go` em `internal/dto/`, com as validações nas tags `binding`; nenhum arquivo de `internal/controller` importa `internal/model` |
| P5 Repository, Mapper, Service e Controller de todas, com o comentário no topo, e o mapa da estrutura | ✅ | Os arquivos de `internal/repository`, `mapper`, `service` e `controller`, ligados pelos construtores `NovoXxx` em `main.go`; todo arquivo começa com `CAMADA`, `RESPONSABILIDADE` e `SITUAÇÃO`; mapa no `README.md`, seção "Arquitetura" |
| P6 API simples: `POST`, `GET` e `GET /{id}`, com `400` e `404` | ✅ | `ProdutoController`: `POST /produtos`, `GET /produtos` e `GET /produtos/{id}`; o `404` nasce em `ProdutoNaoEncontrado` (`internal/service/erros.go`) e sai pelo `TratarErros` |
| P7 No mínimo 3 regras de negócio implementadas, com `400` quando violadas | ✅ | R1, R2 e R3 em `VendaService` (`conferirR2` e `conferirEstoque`), usadas por `POST /vendas`; regra violada vira `RegraDeNegocioError` e sai `400` com a regra na mensagem |
| P8 As regras atendem ao problema e ao tema | ✅ | `README.md`, seção "O problema": as três regras vêm dos prejuízos de vender no balcão e nos marketplaces com o estoque separado |
| P9 Código igual ao desenho | ✅ | As colunas de `docs/modelo-de-dominio.md` são as das migrations e das structs; as rotas, os JSON e as mensagens de erro de `docs/contrato-da-api.md` são os da API |
| P10 README completo e bem escrito | ✅ | `README.md`, conferido no GitHub |

**Pontos que o grupo espera no projeto:** 5,0

A apresentação vale os outros 5,0 pontos: são 3 perguntas, feitas na Aula 11.
