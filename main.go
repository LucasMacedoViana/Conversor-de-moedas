package main

import (
	"Conversor-de-moedas/DB"
	v1 "Conversor-de-moedas/v1"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	//Carregar arquivo .env
	godotenv.Load(".env")
	//Carregar Rotas
	app := v1.Routes()
	//Criar tabela
	DB.CreateTableSaveResponse()
	//Iniciar aplicação
	app.Listen(fmt.Sprint(":", os.Getenv("port_application")))
}
