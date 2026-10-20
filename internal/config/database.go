// CAMADA: Config
// RESPONSABILIDADE: abrir a conexão com o PostgreSQL do docker-compose.yml.
// SITUAÇÃO: implementada.

package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const dsn = "host=localhost port=5437 user=projeto password=projeto dbname=projeto sslmode=disable"

func ConectarBanco() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
