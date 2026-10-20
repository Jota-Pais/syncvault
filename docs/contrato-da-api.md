# Contrato da API

## Convenções

| Item | Valor |
|---|---|
| Endereço base | `http://localhost:8080` |
| Formato | JSON, na requisição e na resposta |
| Dinheiro | Número inteiro em centavos: R$ 109,90 é `10990` |
| Percentual | Número inteiro: `16` é 16% |
| Data e hora | ISO 8601 com fuso horário: `2026-10-20T19:30:00-03:00` |
| Erros | Sempre no formato `ErroDTO` (ver [Erros](#erros)) |
| Documentação viva | `http://localhost:8080/swagger/index.html` |

## Rotas

Coluna **Funciona**: `sim` para as rotas que já respondem nesta etapa: as três da API simples (produtos), a que as regras de negócio usam (`POST /vendas`) e a que mostra os ids dos canais.

### Produtos

| Verbo | Caminho | O que faz | Entrada | Saída | Sucesso | Erros | Funciona |
|---|---|---|---|---|---|---|---|
| `POST` | `/produtos` | Cadastra um produto com o estoque inicial | `ProdutoRequestDTO` | `ProdutoResponseDTO` | `201` | `400` | sim |
| `GET` | `/produtos` | Lista os produtos, com o estoque atual | Nenhuma | lista de `ProdutoResponseDTO` | `200` | Nenhum | sim |
| `GET` | `/produtos/{id}` | Busca um produto | Nenhuma | `ProdutoResponseDTO` | `200` | `400`, `404` | sim |

### Canais

Os canais já nascem cadastrados pela migration 4 (Balcão, Mercado Livre e Shopee), então não existe rota para criar canal. A listagem serve para quem usa a API descobrir o id de cada canal antes de registrar uma venda.

| Verbo | Caminho | O que faz | Entrada | Saída | Sucesso | Erros | Funciona |
|---|---|---|---|---|---|---|---|
| `GET` | `/canais` | Lista os canais, com a taxa e o estoque de segurança de cada um | Nenhuma | lista de `CanalResponseDTO` | `200` | Nenhum | sim |
| `GET` | `/canais/{id}` | Busca um canal | Nenhuma | `CanalResponseDTO` | `200` | `400`, `404` | não |

### Vendas

`POST /vendas` é onde moram as três regras de negócio do README.

| Verbo | Caminho | O que faz | Regra | Entrada | Saída | Sucesso | Erros | Funciona |
|---|---|---|---|---|---|---|---|---|
| `POST` | `/vendas` | Registra uma venda e desconta a quantidade do estoque único | R1, R2 e R3 | `VendaRequestDTO` | `VendaResponseDTO` | `201` | `400`, `404` | sim |
| `GET` | `/vendas` | Lista as vendas, de todos os canais | Nenhuma | Nenhuma | lista de `VendaResponseDTO` | `200` | Nenhum | não |
| `GET` | `/vendas/{id}` | Busca uma venda | Nenhuma | Nenhuma | `VendaResponseDTO` | `200` | `400`, `404` | não |

As regras são conferidas nesta ordem, e a primeira que falhar barra a venda:

1. **R2**: o preço, descontada a taxa do canal, não pode ficar abaixo do custo do produto.
2. **R1**: a quantidade não pode passar do estoque do produto.
3. **R3**: a venda não pode deixar o estoque abaixo do estoque de segurança do canal.

## DTOs

O que entra não é igual ao que sai: o cliente não envia `id` nem os campos que a API decide ou calcula.

| DTO | Direção | Campos | Validações |
|---|---|---|---|
| `ProdutoRequestDTO` | entrada | `nome`, `custoCentavos`, `estoque` | os três obrigatórios; `nome` com até 255 caracteres; `custoCentavos` maior que zero; `estoque` zero ou mais |
| `ProdutoResponseDTO` | saída | `id`, `nome`, `custoCentavos`, `estoque` | Nenhuma |
| `CanalRequestDTO` | entrada | `nome`, `taxaPercentual`, `estoqueSeguranca` | os três obrigatórios; `nome` com até 100 caracteres; `taxaPercentual` de 0 a 100; `estoqueSeguranca` zero ou mais. Nenhuma rota usa nesta etapa: existe para quando houver cadastro de canal |
| `CanalResponseDTO` | saída | `id`, `nome`, `taxaPercentual`, `estoqueSeguranca` | Nenhuma |
| `VendaRequestDTO` | entrada | `produtoId`, `canalId`, `quantidade`, `precoUnitarioCentavos` | os quatro obrigatórios; `quantidade` e `precoUnitarioCentavos` maiores que zero |
| `VendaResponseDTO` | saída | `id`, `produto`, `canal`, `quantidade`, `precoUnitarioCentavos`, `recebidoCentavos`, `vendidaEm` | Nenhuma. `produto` e `canal` voltam completos, e o `produto` já com o estoque depois da venda. `recebidoCentavos` é calculado pela API: o total da venda menos a taxa do canal, arredondado para baixo. `vendidaEm` é o momento em que a API registrou a venda |

## Exemplos

Os exemplos seguem uma história só: a loja cadastra um casaco e vende um pelo Mercado Livre.

### `POST /produtos` (funciona)

Requisição:

```json
{
  "nome": "Casaco corta-vento",
  "custoCentavos": 6000,
  "estoque": 5
}
```

Resposta, `201 Created`:

```json
{
  "id": 1,
  "nome": "Casaco corta-vento",
  "custoCentavos": 6000,
  "estoque": 5
}
```

### `GET /canais` (funciona)

Resposta, `200 OK`:

```json
[
  { "id": 1, "nome": "Balcão", "taxaPercentual": 0, "estoqueSeguranca": 0 },
  { "id": 2, "nome": "Mercado Livre", "taxaPercentual": 16, "estoqueSeguranca": 2 },
  { "id": 3, "nome": "Shopee", "taxaPercentual": 20, "estoqueSeguranca": 2 }
]
```

### `POST /vendas` (funciona)

Entram só os ids; produto e canal voltam completos.

Requisição:

```json
{
  "produtoId": 1,
  "canalId": 2,
  "quantidade": 1,
  "precoUnitarioCentavos": 10990
}
```

Resposta, `201 Created`:

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

O Mercado Livre fica com 16% de R$ 109,90, e a loja recebe R$ 92,31. O estoque do casaco caiu de 5 para 4, e esse é o mesmo estoque que o Balcão e a Shopee enxergam.

## Erros

Todo erro sai no mesmo formato. Este é o `404` de `GET /produtos/999`:

```json
{
  "status": 404,
  "erro": "Not Found",
  "mensagem": "Produto não encontrado: 999",
  "caminho": "/produtos/999",
  "timestamp": "2026-10-20T19:30:00-03:00",
  "campos": []
}
```

No `400` de validação, a lista `campos` traz um item por campo inválido, por exemplo `{"campo": "nome", "mensagem": "nome é obrigatório"}`.

Quando uma regra de negócio barra o pedido, o `400` vem com a regra na `mensagem` e `campos` vazio. Este é o de `POST /vendas` vendendo o casaco a R$ 65,00 no Mercado Livre:

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

As mensagens das outras duas regras, com o mesmo casaco:

| Regra | Situação | `mensagem` |
|---|---|---|
| R1 | Há 4 casacos e o Balcão tenta vender 5 | `R1: estoque insuficiente: há 4 unidades do produto 1 e a venda pede 5` |
| R3 | Há 3 casacos e o Mercado Livre tenta vender 2 | `R3: o canal Mercado Livre não vende as últimas 2 unidades: há 3 em estoque, então ele vende no máximo 1` |

| Status | Quando acontece neste projeto |
|---|---|
| `400` | JSON ilegível, campo inválido no DTO de entrada ou id do caminho que não é número (`/produtos/abc`) |
| `400` | Regra R1, R2 ou R3 violada. A `mensagem` explica a regra |
| `404` | Id de produto ou de canal que não existe |
