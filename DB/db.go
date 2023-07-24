package DB

import (
	"Conversor-de-moedas/models"
	"Conversor-de-moedas/models/config"
	"fmt"
)

// CreateTableSaveResponse cria a tabela save_responses no banco de dados
func CreateTableSaveResponse() {
	DB := config.OpenConnection()
	err := DB.AutoMigrate(&models.SaveResponse{})
	if err != nil {
		fmt.Println("Erro ao criar a tabela:", err)
		return
	}
	fmt.Println("Tabela criada com sucesso")
}
