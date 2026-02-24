package main

import (
	"beyond/internal/controllers"
	"beyond/internal/infra/database"
	"beyond/internal/middlewares"
	"beyond/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	database.Connect()
	database.InitRedis()

	go services.StartProcessor()

	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/cadastro", controllers.SignUp)
		auth.POST("/login", controllers.Login)
	}

	usuario := router.Group("/usuario")
	usuario.Use(middlewares.AuthMiddleware())
	{
		usuario.GET("/perfil", func(c *gin.Context) {
			userID, _ := c.Get("user_id")
			c.JSON(http.StatusOK, gin.H{"message": "Acesso permitido!", "user_id": userID})
		})
	}

	produtos := router.Group("/produtos")
	produtos.Use(middlewares.AuthMiddleware())
	{
		produtos.POST("/monitorar", controllers.CreateProductMonitor)
	}
	router.Run(":8080")
}
