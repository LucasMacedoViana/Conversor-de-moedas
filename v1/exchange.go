package v1

import (
	"Conversor-de-moedas/models"
	"Conversor-de-moedas/utils"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// GetCounting - Rota para conversão de moedas
func GetCounting(c *fiber.Ctx) error {
	//URL da API
	url := os.Getenv("url_api")
	//Receber os parâmetros da URL
	amount := c.Params("amount")
	from := strings.ToUpper(c.Params("from"))
	url += from
	url += "-"
	to := strings.ToUpper(c.Params("to"))
	url += to
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Erro na requisição. Status: %d", resp.StatusCode)
	}
	//Ler o corpo da resposta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	//Converter o JSON para struct
	var apiResponse models.APIResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return err
	}
	//Converter a string para float64
	amountFloat, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		fmt.Println("Erro ao converter a string para float64:", err)
	}
	//Identifica qual moeda está sendo convertida e pega o valor do Bid
	var bidFloat float64
	switch fmt.Sprintf("%s%s", from, to) {
	case "BRLUSD":
		bidFloat, _ = strconv.ParseFloat(apiResponse.BRLUSD.Bid, 64)
	case "USDBRL":
		bidFloat, _ = strconv.ParseFloat(apiResponse.USDBRL.Bid, 64)
	case "BRLEUR":
		bidFloat, _ = strconv.ParseFloat(apiResponse.BRLEUR.Bid, 64)
	case "EURBRL":
		bidFloat, _ = strconv.ParseFloat(apiResponse.EURBRL.Bid, 64)
	case "BTCUSD":
		bidFloat, _ = strconv.ParseFloat(apiResponse.BTCUSD.Bid, 64)
	case "BTCBRL":
		bidFloat, _ = strconv.ParseFloat(apiResponse.BTCBRL.Bid, 64)
	default:
		fmt.Println("Erro ao converter a string para float64:", err)
	}
	//Faz o calculo da conversão
	total := amountFloat * bidFloat

	//Cria um map para retornar o simbolo da moeda
	symbols := map[string]string{
		"BRL": "R$",
		"USD": "$",
		"BTC": "₿",
		"EUR": "€",
	}
	//Pega o simbolo da moeda
	symbol, _ := symbols[to]
	//Cria um map para retornar o resultado
	resultMap := map[string]interface{}{
		"valorConvertido": utils.Round(total, 2),
		"simboloMoeda":    symbol,
	}
	//Salva o resultado no banco de dados
	save := models.SaveResponse{
		Amount: utils.Round(amountFloat, 2),
		From:   from,
		To:     to,
		Result: utils.Round(total, 2),
		Simbol: symbol,
		Rate:   bidFloat,
	}
	//Chama a função para salvar no banco de dados
	save.Create()

	//Retorna o resultado
	return c.JSON(resultMap)
}

// GetCountingCustom - Rota para conversão de moedas com a cotação customizada
func GetCountingCustom(c *fiber.Ctx) error {
	//Receber os parâmetros da URL
	amount := c.Params("amount")
	from := strings.ToUpper(c.Params("from"))
	to := strings.ToUpper(c.Params("to"))
	rate := c.Params("rate")

	//Converter a string para float64
	amountFloat, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		fmt.Println("Erro ao converter a string para float64:", err)
	}
	//Converter a string para float64
	rateFloat, err := strconv.ParseFloat(rate, 64)
	if err != nil {
		fmt.Println("Erro ao converter a string para float64:", err)
	}
	//Faz o calculo da conversão
	total := amountFloat * rateFloat

	//Cria um map para retornar o simbolo da moeda
	symbols := map[string]string{
		"BRL": "R$",
		"USD": "$",
		"BTC": "₿",
		"EUR": "€",
	}
	//Pega o simbolo da moeda
	symbol, _ := symbols[to]
	//Cria um map para retornar o resultado
	resultMap := map[string]interface{}{
		"valorConvertido": utils.Round(total, 2),
		"simboloMoeda":    symbol,
	}
	//Salva o resultado no banco de dados
	save := models.SaveResponse{
		Amount: utils.Round(amountFloat, 2),
		From:   from,
		To:     to,
		Result: utils.Round(total, 2),
		Simbol: symbol,
		Rate:   rateFloat,
	}
	//Chama a função para salvar no banco de dados
	save.Create()

	//Retorna o resultado
	return c.JSON(resultMap)
}
