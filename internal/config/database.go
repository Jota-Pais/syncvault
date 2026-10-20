// CAMADA: Config
// RESPONSABILIDADE: abrir a conexão com o PostgreSQL do docker-compose.yml.
// SITUAÇÃO: implementada.

package config

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Endereço do banco quando a API roda direto na máquina (go run .): o PostgreSQL
// do docker-compose.yml, publicado na porta 5439. Mudou lá, muda aqui.
const dsnLocal = "postgres://projeto:projeto@localhost:5439/projeto?sslmode=disable"

// enderecoDoBanco usa a variável de ambiente DATABASE_URL quando ela existe. É o
// caso da API rodando no Docker: lá dentro, o banco se chama postgres e fica na
// porta 5432, e quem define a variável é o docker-compose.yml.
func enderecoDoBanco() string {
	url := os.Getenv("DATABASE_URL")
	if url != "" {
		return url
	}
	return dsnLocal
}

func ConectarBanco() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(enderecoDoBanco()), &gorm.Config{})
}
