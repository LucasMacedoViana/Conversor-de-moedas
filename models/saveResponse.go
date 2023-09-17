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

// FindAll busca todos os registros da tabela save_responses
func (s *SaveResponse) FindAll() ([]SaveResponse, error) {
	connection := config.OpenConnection()
	defer config.CloseConnection(connection)

	var saveResponses []SaveResponse
	err := connection.Table("save_responses").Find(&saveResponses).Error
	if err != nil {
		return nil, err
	}
	return saveResponses, nil
}

// FindByID busca um registro da tabela save_responses pelo ID
func (s *SaveResponse) FindByID(id string) (SaveResponse, error) {
	connection := config.OpenConnection()
	defer config.CloseConnection(connection)

	var saveResponse SaveResponse
	err := connection.Table("save_responses").Where("id = ?", id).First(&saveResponse).Error
	if err != nil {
		return SaveResponse{}, err
	}
	return saveResponse, nil
}

// DeleteByID deleta um registro da tabela save_responses pelo ID
func (s *SaveResponse) DeleteByID(id string) error {
	connection := config.OpenConnection()
	defer config.CloseConnection(connection)

	err := connection.Table("save_responses").Where("id = ?", id).Delete(&s).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateByID atualiza um registro da tabela save_responses pelo ID
func (s *SaveResponse) UpdateByID(id string) error {
	connection := config.OpenConnection()
	defer config.CloseConnection(connection)

	err := connection.Table("save_responses").Where("id = ?", id).Updates(&s).Error
	if err != nil {
		return err
	}
	return nil
}
