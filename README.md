﻿#Conversor de moedas

## Descrição do Projeto
#### O projeto consiste em um conversor de moedas, onde o usuário pode converter o valor de uma moeda para outra

## Funcionalidades
- [x] Conversão de moedas
- [x] Mostrar o valor convertido
- [x] Mostrar o simbolo da moeda a qual foi convertida

## Como rodar a aplicação
Configure o .env com suas credenciais do postgres.

Crie uma database no postgresSql com o nome de conversor_moedas.

No terminal, clone o projeto:
```
git@github.com:LucasMacedoViana/Conversor-de-moedas.git
```
Entre na pasta do projeto:
```
cd Conversor-de-moedas
```
No terminal, rode o codigo:
```
go run main.go
```
## Tecnologias utilizadas
- [x] Golang
- [x] API
- [x] JSON
- [x] HTTP
- [x] PostgresSql

## Frameworks
- [x] Fiber
- [x] Gorm

## API utilizada
- [x] [AwesomeAPI](https://docs.awesomeapi.com.br/api-de-moedas)

## Endpoints (onde tem {Valor} voce deve colococar o valor que deseja converter, lembre-se de tirar as chaves)
- [x] Para converter Dolar em Real: http://localhost:8000/exchange/{valor}/USD/BRL
- [x] Para converter Real em Dolar: http://localhost:8000/exchange/{valor}/BRL/USD
- [x] Para converter Euro em Real: http://localhost:8000/exchange/{valor}/EUR/BRL
- [x] Para converter Real em Euro: http://localhost:8000/exchange/{valor}/BRL/EUR
- [x] Para converter Bitcoin em Real: http://localhost:8000/exchange/{valor}/BTC/BRL
- [x] Para converter Bitcoin em Dolar: http://localhost:8000/exchange/{valor}/BTC/USD

## Caso queira colocar uma cotação usar esses endpoints (onde tem {Valor} voce deve colococar o valor que deseja converter e onde tem {cotação} voce deve colocar a cotação que deseja usar, lembre-se de tirar as chaves)
- [x] Para converter Dolar em Real: http://localhost:8000/exchange/{valor}/USD/BRL/{cotação}
- [x] Para converter Real em Dolar: http://localhost:8000/exchange/{valor}/BRL/USD/{cotação}
- [x] Para converter Euro em Real: http://localhost:8000/exchange/{valor}/EUR/BRL/{cotação}
- [x] Para converter Real em Euro: http://localhost:8000/exchange/{valor}/BRL/EUR/{cotação}
- [x] Para converter Bitcoin em Real: http://localhost:8000/exchange/{valor}/BTC/BRL/{cotação}
- [x] Para converter Bitcoin em Dolar: http://localhost:8000/exchange/{valor}/BTC/USD/{cotação}
## Exemplo de retorno
```
{
    "simboloMoeda": "R$"
    "valorConvertido": 5.5,
}
```



