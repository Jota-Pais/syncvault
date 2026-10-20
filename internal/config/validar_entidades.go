// CAMADA: Config
// RESPONSABILIDADE: conferir, quando a API sobe, se cada entidade bate com a
//                   tabela que a migration criou: a tabela existe e cada campo
//                   tem a sua coluna. É o papel do ddl-auto=validate do Spring:
//                   se não bater, a API não sobe e o erro diz o que falta.
// SITUAÇÃO: implementada.

package config

import (
	"fmt"

	"gorm.io/gorm"
)

func ValidarEntidades(db *gorm.DB, entidades ...any) error {
	for _, entidade := range entidades {
		// Lê as tags gorm da struct: nome da tabela e coluna de cada campo.
		stmt := &gorm.Statement{DB: db}
		err := stmt.Parse(entidade)
		if err != nil {
			return err
		}

		tabela := stmt.Schema.Table
		if !db.Migrator().HasTable(tabela) {
			return fmt.Errorf("a tabela %s da entidade %s não existe: falta a migration", tabela, stmt.Schema.Name)
		}

		for _, campo := range stmt.Schema.Fields {
			if !db.Migrator().HasColumn(entidade, campo.DBName) {
				return fmt.Errorf("o campo %s.%s espera a coluna %s.%s, que não existe no banco",
					stmt.Schema.Name, campo.Name, tabela, campo.DBName)
			}
		}
	}
	return nil
}
