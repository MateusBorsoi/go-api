package services

import (
	"beyond/internal/infra/database"
	"beyond/internal/models"
	"encoding/json"
	"fmt"
)

const QueueName = "scraping_tasks"

func AddToQueue(product models.Product) error {
	// Aqui você implementaria a lógica para adicionar a tarefa à fila.
	data, err := json.Marshal(product)
	if err != nil {
		return err
	}
	return database.Rdb.RPush(database.Ctx, QueueName, data).Err()
}

func StartProcessor() {
	fmt.Printf("Worker de Scraping iniciado e aguardando tarefas...")
	for {
		result, err := database.Rdb.BLPop(database.Ctx, 0, QueueName).Result()
		if err != nil {
			fmt.Printf("Erro ao processar a fila Redis: %v\n", err)
			continue
		}

		rawJSON := result[1]
		var product models.Product
		json.Unmarshal([]byte(rawJSON), &product)

		fmt.Printf("Processando: %s\n", product.URL)

		price, title, err := ScrapeProduct(product.URL)
		if err != nil {
			fmt.Printf("Erro ao ler site: %v\n", err)
			continue
		}
		fmt.Printf("Sucesso! Produto: %s | Preço Atual: %.2f | Alvo: %.2f\n", title, price, product.TargetPrice)

		if price <= product.TargetPrice {
			fmt.Println("🚨 ALERTA DE PREÇO! O PRODUTO BAIXOU!")
			//Enviar notificacao aqui

		}
	}
}
