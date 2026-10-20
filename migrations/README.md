# migrations: a criação das tabelas

Cada arquivo SQL daqui cria ou altera tabelas. Quando a API sobe, o golang-migrate (o Flyway do Go) aplica, em ordem, os que ainda não rodaram, e anota a versão na tabela `schema_migrations`.

- O nome é `<versão>_<o_que_faz>.up.sql`, com a versão em 6 dígitos para os arquivos ficarem em ordem na pasta: `000001_criar_produtos.up.sql`.
- A ordem é a da versão: a tabela que as outras referenciam nasce primeiro.
- Tabela no plural e em `snake_case` (`produtos`, `preco_unitario_centavos`); chave estrangeira com `REFERENCES`.
- **Nunca edite uma migration que já rodou.** Para corrigir, crie a próxima (`000005_...`). Se o banco ainda está só na sua máquina, recrie do zero com `docker compose down -v`.

| Versão | Arquivo | O que faz |
|---|---|---|
| 1 | `000001_criar_produtos.up.sql` | Cria a tabela `produtos` |
| 2 | `000002_criar_canais.up.sql` | Cria a tabela `canais` |
| 3 | `000003_criar_vendas.up.sql` | Cria a tabela `vendas`, com as chaves estrangeiras para `produtos` e `canais` |
| 4 | `000004_cadastrar_canais.up.sql` | Cadastra os três canais: Balcão, Mercado Livre e Shopee |
