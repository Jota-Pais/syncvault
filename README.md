# SyncVault

Estoque único para uma loja que vende ao mesmo tempo no balcão, no Mercado Livre e na Shopee: toda venda, de qualquer canal, desconta do mesmo estoque, e o sistema barra a venda que daria prejuízo ou que deixaria um marketplace vender a última peça.

Projeto Final da disciplina de Backend (Engenharia de Software, SATC). **Etapa 1: arquitetura do backend.** Feito em **Go**, com a autorização do professor para usar outra linguagem. A seção [Go no lugar de Java](#go-no-lugar-de-java) mostra onde cada critério da revisão aparece no projeto.

## Índice

- [Integrantes](#integrantes)
- [O problema](#o-problema)
- [Tecnologias](#tecnologias)
- [Go no lugar de Java](#go-no-lugar-de-java)
- [Arquitetura](#arquitetura)
- [Entidades](#entidades)
- [Rotas](#rotas)
- [Como executar](#como-executar)
- [Situação do projeto](#situação-do-projeto)
- [Documentação](#documentação)

## Integrantes

| Nome | GitHub |
|---|---|
| João Carlos Pais | [@Jota-Pais](https://github.com/Jota-Pais) |
| Gabriel Ugione | [@tpswrds1345ggzc-svg](https://github.com/tpswrds1345ggzc-svg) |
| Arthur Santin | [@ArthurSSantin](https://github.com/ArthurSSantin) |

## O problema

**Tema sorteado:** Varejo e E-commerce.

Lojas pequenas que vendem no balcão e em marketplaces controlam o estoque de cada canal separado, muitas vezes numa planilha. Quando uma peça sai no balcão, o anúncio do Mercado Livre continua no ar, e a loja vende o que não tem: precisa cancelar a venda e perde reputação no marketplace. E como cada marketplace fica com uma comissão diferente, é fácil vender abaixo do custo sem perceber. O SyncVault mantém um estoque só para todos os canais e confere cada venda antes de aceitá-la.

**Quem usa o sistema:**

| Quem | O que faz no sistema |
|---|---|
| Lojista | Cadastra os produtos com o custo e o estoque inicial, e acompanha o estoque único |
| Vendedor | Registra cada venda, do balcão ou dos marketplaces, com o canal e o preço praticado |

**Regras de negócio.** O grupo criou as três regras olhando para os prejuízos do estoque separado: vender o que já saiu, vender abaixo do custo por causa da comissão e vender pelo marketplace a peça que, na verdade, já foi vendida no balcão e o anúncio ainda não sabe.

| # | Regra | Onde está no código | Rota que usa |
|---|---|---|---|
| R1 | Uma venda não pode passar do estoque do produto, mesmo com duas vendas chegando ao mesmo tempo pela última peça | `VendaService` (`conferirEstoque`) e `VendaRepository.SalvarDescontandoEstoque` | `POST /vendas` |
| R2 | O preço, descontada a comissão do canal, não pode ficar abaixo do custo do produto: a loja não vende no prejuízo | `VendaService` (`conferirR2`) | `POST /vendas` |
| R3 | Um marketplace não pode vender as últimas unidades do produto: cada canal tem um estoque de segurança, porque o marketplace atualiza o estoque com atraso. O balcão vende até a última | `VendaService` (`conferirEstoque`) e `VendaRepository.SalvarDescontandoEstoque` | `POST /vendas` |

As regras são conferidas na ordem R2, R1 e R3. Quando uma delas barra a venda, a API devolve `400` com a regra explicada na mensagem.

**Fora do escopo:** a conexão automática com as APIs do Mercado Livre e da Shopee (nesta etapa, as vendas de lá são registradas pela API do SyncVault), o cadastro e a edição de canais (os três já nascem cadastrados), a reposição de estoque, a edição e a remoção de produtos, login e tela de frontend.

## Tecnologias

| O quê | Qual |
|---|---|
| Linguagem | Go 1.27 |
| Framework web | Gin |
| Validação | go-playground/validator, pelas tags `binding` do Gin |
| Acesso ao banco | GORM |
| Banco de dados | PostgreSQL 16, em Docker |
| Migrations | golang-migrate |
| Documentação da API | swaggo (Swagger UI) |
| Execução | Docker Compose: o banco e a API, compilada pelo `Dockerfile` |

## Go no lugar de Java

**Por que Go.** O ponto mais delicado do problema é a disputa pela última peça: o balcão e um marketplace vendendo o mesmo produto no mesmo segundo. Em Go, cada requisição roda numa *goroutine*, que custa poucos KB, então a API aguenta muitas vendas simultâneas com pouca memória. A garantia de que só uma leva a peça é do banco: a conferência do estoque e o desconto são o mesmo `UPDATE`, numa transação com a gravação da venda. No teste do grupo, 10 vendas simultâneas disputaram um produto com 5 unidades no Mercado Livre (estoque de segurança 2): passaram exatamente 3, e o estoque parou em 2.

**Cada peça da disciplina, em Go:**

| Na disciplina (Java) | No SyncVault (Go) |
|---|---|
| `@RestController` e `@GetMapping` | Controller com o método `RegistrarRotas`, que liga cada verbo e caminho a um método |
| `@Valid` com `@NotBlank`, `@Positive`... | `ShouldBindJSON` no Controller, com as tags `binding` dos DTOs (`required`, `gt=0`...) |
| `record` como DTO | `struct` em `internal/dto`, separada da entidade |
| `@Entity` e `@Table` | `struct` em `internal/model`, com as tags `gorm` (coluna de cada campo) e o método `TableName` |
| `JpaRepository` | `struct` em `internal/repository` com `Salvar`, `ListarTodos` e `BuscarPorID` |
| Flyway (`V1__criar_livros.sql`) | golang-migrate (`000001_criar_produtos.up.sql`), aplicado quando a API sobe |
| `ddl-auto=validate` | `config.ValidarEntidades`: se uma struct não bater com a tabela, a API não sobe e diz qual coluna falta |
| Injeção de dependência do Spring | Construtores `NovoXxx`, chamados em `main.go` de baixo para cima: Repository e Mapper, Service, Controller |
| `RecursoNaoEncontradoException` e `LivroNaoEncontradoException` | `service.RecursoNaoEncontradoError`, criado por `ProdutoNaoEncontrado`, `CanalNaoEncontrado` e `VendaNaoEncontrada` |
| `RegraDeNegocioException` | `service.RegraDeNegocioError` |
| `@RestControllerAdvice` (`ApiExceptionHandler`) | `controller.TratarErros`, que transforma cada erro no `ErroDTO` |
| springdoc-openapi | swaggo: comentários `@Summary`, `@Router`... nos Controllers, gerados na pasta `swagger/` |

**Onde cada critério da revisão aparece:**

| Critério | No SyncVault |
|---|---|
| P1 Sobe pelo README | [Como executar](#como-executar): só Docker, com `docker compose up -d --build`, que sobe o banco e a API sem precisar instalar Go |
| P2 Camadas | Uma pasta por camada em `internal/`: `controller`, `service`, `mapper`, `repository`, `model` e `dto`, com os nomes da convenção (`ProdutoController`, `ProdutoService`...) |
| P3 Banco e entidades | As 4 migrations em `migrations/` criam as tabelas e as chaves estrangeiras; as structs de `internal/model` batem com elas, e `config.ValidarEntidades` confere isso a cada subida. O GORM nunca cria tabela (não usamos `AutoMigrate`) |
| P4 DTOs | `internal/dto` tem o DTO de entrada, com as validações, e o de saída das três entidades. Nenhum arquivo de `internal/controller` importa `internal/model` |
| P5 Estrutura completa e documentada | As três entidades têm Repository, Mapper, Service e Controller, ligados pelo construtor em `main.go`; todo arquivo começa com o cabeçalho `CAMADA`, `RESPONSABILIDADE` e `SITUAÇÃO`; o [mapa da estrutura](#arquitetura) lista todos |
| P6 API simples | `POST /produtos`, `GET /produtos` e `GET /produtos/{id}`: campo inválido devolve `400`, e id que não existe devolve `404` pelo `ProdutoNaoEncontrado` do Service |
| P7 Regras implementadas | R1, R2 e R3 no `VendaService`; regra violada lança `RegraDeNegocioError`, que vira `400` |
| P8 Regras que atendem ao tema | Ver [O problema](#o-problema): as três regras nascem dos prejuízos de vender em vários canais com estoque separado |
| P9 Código igual ao desenho | As colunas de `docs/modelo-de-dominio.md` são as das migrations e das structs; as rotas de `docs/contrato-da-api.md` são as dos Controllers |
| P10 README | Este arquivo |

## Arquitetura

A API é organizada em camadas. Uma requisição entra pelo Controller, passa pelo Service, que aplica as regras de negócio, e chega ao banco pelo Repository; a resposta faz o caminho de volta. O Mapper converte DTO em entidade e entidade em DTO.

**Mapa da estrutura:**

| Entidade | Migration | Model | Repository | DTOs | Mapper | Service | Controller |
|---|---|---|---|---|---|---|---|
| Produto | `000001_criar_produtos.up.sql`, pronta | `Produto`, pronta | `ProdutoRepository`, pronto | `ProdutoRequestDTO`, `ProdutoResponseDTO`, prontos | `ProdutoMapper`, implementado | `ProdutoService`, implementado | `ProdutoController`, implementado |
| Canal | `000002_criar_canais.up.sql` e `000004_cadastrar_canais.up.sql`, prontas | `Canal`, pronta | `CanalRepository`, pronto | `CanalRequestDTO`, `CanalResponseDTO`, prontos | `CanalMapper`, em parte: só a conversão para a resposta | `CanalService`, em parte: só a listagem | `CanalController`, em parte: só `GET /canais` |
| Venda | `000003_criar_vendas.up.sql`, pronta | `Venda`, pronta | `VendaRepository`, pronto | `VendaRequestDTO`, `VendaResponseDTO`, prontos | `VendaMapper`, implementado | `VendaService`, implementado (R1, R2 e R3) | `VendaController`, em parte: só `POST /vendas` |

Os arquivos de cada entidade ficam em `internal/<camada>/`, com o nome da entidade: `internal/service/venda_service.go`, por exemplo. Também fazem parte do projeto, sem ser de uma entidade só:

| Arquivo | O que faz |
|---|---|
| `Dockerfile` | Monta a imagem da API: compila o Go numa etapa e deixa só o executável e as migrations na imagem final |
| `docker-compose.yml` | Sobe o banco e a API; a API só liga depois que o banco está pronto |
| `main.go` | Liga tudo: conecta no banco, roda as migrations, confere as entidades, monta as camadas e sobe a API na porta 8080 |
| `internal/config/database.go` | Abre a conexão com o PostgreSQL |
| `internal/config/migrations.go` | Aplica as migrations que ainda não rodaram |
| `internal/config/validar_entidades.go` | Confere se cada struct de `model` bate com a sua tabela |
| `internal/config/swagger.go` | Publica o Swagger UI |
| `internal/service/erros.go` | Os erros do domínio: não encontrado (`404`) e regra violada (`400`) |
| `internal/controller/parametros.go` | Lê o `{id}` do caminho; um id que não é número vira `400` |
| `internal/controller/tratamento_de_erros.go` | `TratarErros`: transforma cada erro das rotas em `400`, `404` ou `500` no formato `ErroDTO` |
| `internal/dto/erro.go` | `ErroDTO` e `CampoErroDTO`, o formato único de erro |

Canal não tem rota de cadastro porque os três canais nascem com a migration 4, e canal quase nunca muda. O `CanalRequestDTO` existe para quando houver esse cadastro.

## Entidades

| Entidade | O que representa | Como se relaciona |
|---|---|---|
| Produto | Um item à venda, com um estoque só para todos os canais | Aparece em várias vendas |
| Canal | Onde a loja vende: Balcão, Mercado Livre ou Shopee, com a comissão e o estoque de segurança de cada um | Registra várias vendas |
| Venda | A venda de um produto por um canal, com a quantidade e o preço praticado | Pertence a um produto e a um canal |

Atributos, tipos e o diagrama das tabelas estão em [docs/modelo-de-dominio.md](docs/modelo-de-dominio.md). Dinheiro é sempre um número inteiro em centavos: R$ 109,90 é `10990`.

## Rotas

### As que já funcionam

| Verbo | Caminho | O que faz | Regra | Sucesso | Erros |
|---|---|---|---|---|---|
| `POST` | `/produtos` | Cadastra um produto com o estoque inicial | Nenhuma | `201` | `400` |
| `GET` | `/produtos` | Lista os produtos, com o estoque atual | Nenhuma | `200` | Nenhum |
| `GET` | `/produtos/{id}` | Busca um produto pelo id | Nenhuma | `200` | `400`, `404` |
| `GET` | `/canais` | Lista os canais, com a taxa e o estoque de segurança | Nenhuma | `200` | Nenhum |
| `POST` | `/vendas` | Registra uma venda e desconta do estoque único | R1, R2 e R3 | `201` | `400`, `404` |

Exemplo de requisição:

```http
POST /vendas
Content-Type: application/json

{
  "produtoId": 1,
  "canalId": 2,
  "quantidade": 1,
  "precoUnitarioCentavos": 10990
}
```

Resposta, `201 Created`. Entram os ids; produto e canal voltam completos, o produto já com o estoque depois da venda, e `recebidoCentavos` mostra quanto a loja recebe depois da comissão do Mercado Livre:

```json
{
  "id": 1,
  "produto": { "id": 1, "nome": "Casaco corta-vento", "custoCentavos": 6000, "estoque": 4 },
  "canal": { "id": 2, "nome": "Mercado Livre", "taxaPercentual": 16, "estoqueSeguranca": 2 },
  "quantidade": 1,
  "precoUnitarioCentavos": 10990,
  "recebidoCentavos": 9231,
  "vendidaEm": "2026-10-20T19:30:00-03:00"
}
```

### Erros

Todo erro sai no mesmo formato. Este é o `400` da regra R2, quando alguém tenta vender o casaco a R$ 65,00 no Mercado Livre:

```json
{
  "status": 400,
  "erro": "Bad Request",
  "mensagem": "R2: vender a R$ 65,00 no canal Mercado Livre, com taxa de 16%, rende R$ 54,60 por unidade, abaixo do custo de R$ 60,00",
  "caminho": "/vendas",
  "timestamp": "2026-10-20T19:30:00-03:00",
  "campos": []
}
```

Um campo inválido também devolve `400`, com a lista `campos` dizendo qual campo corrigir. Um produto ou canal que não existe devolve `404`.

### As planejadas

Estão no contrato, mas ainda não respondem.

| Recurso | O que vai fazer |
|---|---|
| `GET /canais/{id}` | Buscar um canal |
| `GET /vendas` e `GET /vendas/{id}` | Listar e buscar as vendas de todos os canais |

O contrato completo, com todas as rotas e os DTOs, está em [docs/contrato-da-api.md](docs/contrato-da-api.md).

## Como executar

### Só com Docker: o jeito para rodar e corrigir

Não precisa instalar Go nem nada além do Docker: a API é compilada dentro do próprio Docker.

**Pré-requisitos:** Git e o [Docker Desktop](https://www.docker.com/products/docker-desktop/) instalado e **aberto**. Ele está pronto quando aparece **Engine running** no canto inferior esquerdo da janela dele.

```bash
git clone https://github.com/Jota-Pais/syncvault.git
cd syncvault
docker compose up -d --build
```

O último comando faz tudo:

1. monta a imagem da API pelo `Dockerfile`, compilando o código Go dentro do Docker;
2. sobe o PostgreSQL, na porta 5439;
3. espera o banco ficar pronto e sobe a API, na porta 8080;
4. a API aplica as migrations sozinha: cria as tabelas e cadastra os três canais.

Na primeira vez, ele demora de 2 a 5 minutos, porque baixa a imagem do Go e as bibliotecas. Nas próximas, leva segundos.

Para conferir se está tudo no ar:

```bash
docker compose ps
```

Tem que aparecer o `syncvault-postgres` como `healthy` e o `syncvault-api` como `Up`. Aí é só abrir no navegador:

| Endereço | O que é |
|---|---|
| <http://localhost:8080/swagger/index.html> | Swagger UI: as rotas, com **Try it out** para enviar requisições |
| <http://localhost:8080/swagger/doc.json> | O contrato OpenAPI em JSON |

Para testar as rotas, use o Swagger UI ou o **Postman**. No Postman, importe todas as rotas de uma vez: **Import** → **Link** → `http://localhost:8080/swagger/doc.json`.

### Para ver as regras funcionando

No Swagger UI, cadastre um produto com custo de R$ 60,00 e 5 unidades (`POST /produtos` com `"custoCentavos": 6000` e `"estoque": 5`) e veja os canais em `GET /canais`. Depois, em `POST /vendas`:

1. Venda 1 unidade pelo Mercado Livre (canal 2) a R$ 65,00: a R2 barra, porque depois da comissão a loja receberia menos que o custo.
2. Venda 1 unidade pelo Mercado Livre a R$ 109,90: passa, e o estoque cai para 4.
3. Venda 1 unidade pelo Balcão (canal 1) a R$ 99,90: o estoque cai para 3.
4. Venda 2 unidades pela Shopee (canal 3) a R$ 99,90: a R3 barra, porque a Shopee não vende as últimas 2.
5. Venda 4 unidades pelo Balcão a R$ 99,90: a R1 barra, porque só há 3.

### Comandos do dia a dia

Todos rodam na pasta do projeto.

| Para | Comando |
|---|---|
| Ver se está tudo no ar | `docker compose ps` |
| Acompanhar o log da API (cada requisição aparece aqui) | `docker compose logs -f api` (sai com `Ctrl+C`) |
| Parar tudo, guardando os dados | `docker compose stop` |
| Ligar de novo | `docker compose start` |
| Apagar o banco e começar do zero | `docker compose down -v` e depois `docker compose up -d --build` |

### Se algo der errado

| O que aparece | O que fazer |
|---|---|
| `error during connect` ou `cannot find the file specified` com `dockerDesktopLinuxEngine` | O Docker Desktop está fechado. Abra, espere o **Engine running** e rode o comando de novo |
| `port is already allocated` ou `address already in use` | Outro programa está usando a porta 8080 ou a 5439. Veja quem é com `docker ps` e pare o outro container, ou feche o programa |
| O Swagger UI não abre logo depois do comando | A API só sobe depois que o banco fica pronto. Espere alguns segundos e confira com `docker compose ps`; se o `syncvault-api` não estiver `Up`, `docker compose logs api` mostra o motivo |

### Para desenvolver: a API direto na máquina

Quem vai mexer no código roda a API com o Go, e só o banco no Docker. Pré-requisito a mais: [Go](https://go.dev/dl/) 1.21 ou mais novo (o projeto pede o 1.27; com uma versão mais antiga, o próprio Go baixa a 1.27 sozinho na primeira vez).

```bash
docker compose up -d postgres   # só o banco, na porta 5439
go run .                        # aplica as migrations e sobe a API na porta 8080
```

Se a API do Docker estiver ligada, ela ocupa a porta 8080: pare antes com `docker compose stop api`. No Windows, o firewall pode perguntar se permite o acesso do `go run`: pode permitir na rede privada.

Quem mudar os comentários `@` de uma rota precisa gerar o Swagger de novo:

```bash
go install github.com/swaggo/swag/cmd/swag@v1.16.4
swag init -g internal/config/swagger.go -o swagger --parseInternal --outputTypes go,json
```

## Situação do projeto

| Parte | Situação |
|---|---|
| Desenho | Completo: problema e regras neste README, modelo e contrato em `docs/` |
| Estrutura | As 3 entidades têm migration, entidade, Repository, DTOs, Mapper, Service e Controller |
| API e regras | Produtos com as 3 rotas da API simples; listagem de canais; venda com R1, R2 e R3 implementadas e testadas, inclusive com vendas simultâneas |
| Erros | Todo `400` e `404` sai no formato `ErroDTO`, montado em um lugar só pelo `TratarErros` |

**Ainda não funciona:** as rotas planejadas `GET /canais/{id}`, `GET /vendas` e `GET /vendas/{id}`, que nenhuma regra usa nesta etapa.

## Documentação

| Documento | O que tem |
|---|---|
| [Modelo de domínio](docs/modelo-de-dominio.md) | Entidades, relacionamentos, diagrama das tabelas, migrations |
| [Contrato da API](docs/contrato-da-api.md) | Todas as rotas, os DTOs, exemplos de JSON e erros |
| [O caminho de uma venda](docs/caminho-de-uma-venda.md) | `POST /vendas` classe por classe, em diagrama, e onde cada pedido é barrado |
| [Autoavaliação](docs/autoavaliacao.md) | Os critérios da revisão do código, marcados pelo grupo |
