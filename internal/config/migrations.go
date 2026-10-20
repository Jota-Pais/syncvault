// CAMADA: Config
// RESPONSABILIDADE: aplicar no banco os arquivos .sql da pasta migrations quando a API sobe.
// SITUAÇÃO: implementada.

package config

import (
	"errors"
	"io/fs"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RodarMigrations() error {
	m, err := migrate.New("file://migrations", dsn)
	if err != nil {
		return err
	}
	defer m.Close()

	err = m.Up()
	if errors.Is(err, migrate.ErrNoChange) || errors.Is(err, fs.ErrNotExist) {
		return nil
	}
	return err
}
