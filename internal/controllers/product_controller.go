package controllers

import (
	"beyond/internal/models"
	"beyond/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProductMonitor(c *gin.Context) {
	var input models.Product

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.AddToQueue(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao enfileirar job"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Job enfileirado com sucesso"})
}
