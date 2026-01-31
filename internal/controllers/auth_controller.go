package controllers

import (
	"beyond/internal/infra/database"
	"beyond/internal/models"
	"beyond/internal/services"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
)

func SignUp(c *gin.Context) {
	var body struct {
		Nome  string `json:"nome" binding:"required"`
		Email string `json:"email" binding:"required"`
		Senha string `json:"senha" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Campos obrigatórios: nome, email, senha"})
		return
	}

	user := models.User{
		Nome:  body.Nome,
		Email: body.Email,
		Senha: body.Senha,
	}
	result := database.DB.Create(&user)

	if result.Error != nil {
		var pgErr *pgconn.PgError
		if errors.As(result.Error, &pgErr) && pgErr.Code == "23505" {
			c.JSON(http.StatusConflict, gin.H{
				"error": "Já existe um usuário cadastrado com este e-mail, verifique!",
			})
			return
		}
		c.JSON(http.StatusConflict, gin.H{
			"error": result.Error,
		})

	}

	c.JSON(http.StatusCreated, gin.H{
		"message": " Usuário criado com sucesso!",
		"id":      user.ID,
		"email":   user.Email,
	})
}

func Login(c *gin.Context) {

	var body struct {
		Email string `json:"email" binding:"required"`
		Senha string `json:"senha" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "E-mail e senha são obrigatórios"})
		return
	}

	var user models.User

	if err := database.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}

	token, err := services.GenerateJWT(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Falha ao gerar JWT" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"access_token": token,
		"email":        user.Email,
		"nome":         user.Nome,
	})
}
