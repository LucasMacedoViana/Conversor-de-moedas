package models

import (
	"Conversor-de-moedas/models/config"
	"gorm.io/gorm"
)

// SaveResponse é a estrutura da tabela save_responses
type SaveResponse struct {
	gorm.Model
	Amount float64 `json:"amount"`
	From   string  `json:"from"`
	To     string  `json:"to"`
	Result float64 `json:"result"`
	Simbol string  `json:"simbol"`
	Rate   float64 `json:"rate"`
}

// Create salva o resultado da conversão no banco de dados
func (s *SaveResponse) Create() error {
	connection := config.OpenConnection()
	defer config.CloseConnection(connection)

	err := connection.Table("save_responses").Create(&s).Error
	if err != nil {
		return err
	}
	return err
}
