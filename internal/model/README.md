# model: as entidades

Cada struct daqui representa uma tabela do banco. São as entidades de `docs/modelo-de-dominio.md`.

O que toda entidade tem, e o equivalente em Java:

| Em Go | Para quê | Em Java |
|---|---|---|
| Uma `struct` com um campo por coluna | A linha da tabela | Classe com `@Entity` |
| O método `TableName()` | O nome da tabela que a migration criou (`produtos`, `canais`, `vendas`) | `@Table(name = "...")` |
| A tag `gorm:"column:..."` em cada campo | A coluna que guarda o campo | `@Column` |
| `ID int64` com `gorm:"column:id;primaryKey"` | A chave, gerada pelo banco | `@Id` e `@GeneratedValue` |
| `ProdutoID` e `CanalID` em `Venda` | As chaves estrangeiras (muitos para um) | `@ManyToOne` |

Os campos começam com maiúscula porque o GORM precisa enxergá-los de fora do pacote; em Go não existe getter. Quem protege a entidade é a regra de que ela nunca sai do Service: o Controller e o JSON só veem o DTO.

Quem cria as tabelas é a migration, nunca o GORM. Quando a API sobe, `config.ValidarEntidades` confere se cada struct bate com a sua tabela.
