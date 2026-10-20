// CAMADA: Config
// RESPONSABILIDADE: publicar a documentação viva da API: o Swagger UI em
//                   /swagger/index.html e o contrato OpenAPI em JSON, que o
//                   Postman importa, em /swagger/doc.json. Os dois são gerados
//                   pelo swag a partir dos comentários @ dos controllers, e o
//                   resultado fica na pasta swagger/ da raiz.
// SITUAÇÃO: implementada.

package config

import (
	// Registra o contrato gerado pelo swag; não é usado pelo nome.
	_ "syncvault/swagger"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Informações gerais que aparecem no topo do Swagger UI.
//
// @title       SyncVault API
// @version     1.0
// @description Estoque único para uma loja que vende no balcão, no Mercado Livre e na Shopee: toda venda, de qualquer canal, desconta do mesmo estoque. Dinheiro sempre em centavos (R$ 109,90 = 10990).
// @BasePath    /

// RegistrarSwagger liga a rota que serve o Swagger UI e o contrato em JSON.
func RegistrarSwagger(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
