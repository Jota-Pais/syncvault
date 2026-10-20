# config: configuração da aplicação

O que não pertence a nenhuma entidade. O `main.go` chama estas funções, nesta ordem, antes de a API começar a responder:

| Arquivo | Função | O que faz | Em Java seria |
|---|---|---|---|
| `database.go` | `ConectarBanco` | Abre a conexão com o PostgreSQL do `docker-compose.yml` | `spring.datasource.*` no `application.properties` |
| `migrations.go` | `RodarMigrations` | Aplica as migrations de `migrations/` que ainda não rodaram | O Flyway subindo junto com a aplicação |
| `validar_entidades.go` | `ValidarEntidades` | Confere se cada struct de `model` bate com a sua tabela; se não bater, a API não sobe e o erro diz qual coluna falta | `spring.jpa.hibernate.ddl-auto=validate` |
| `swagger.go` | `RegistrarSwagger` | Publica o Swagger UI em `/swagger/index.html` | `OpenApiConfig` e o springdoc |

Os dados do banco (porta 5439, banco, usuário e senha `projeto`) são os mesmos do `docker-compose.yml`: mudou lá, muda no `database.go`.
