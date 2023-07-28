package main

import (
	"encoding/json"
	"fmt"
	"go-currency-converter/db"
	"go-currency-converter/model"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type Response struct {
	ValorConvertido float64
	SimboloMoeda    string
}

func main() {
	db := db.GetDbConnection()
	db.AutoMigrate(&model.Conversion{})
	currencies := []string{"USD", "BRL", "EUR", "BTC"}
	currencySymbols := map[string]string{"USD": "$", "BRL": "R$", "EUR": "€", "BTC": "₿"}
	server := echo.New()
	server.GET("/exchange/:amount/:from/:to/:rate", func(c echo.Context) error {
		response := &Response{}
		amount, err := strconv.ParseFloat(c.Param("amount"), 8)
		if err != nil {
			panic(err)
		}
		rate, err := strconv.ParseFloat(c.Param("rate"), 8)
		if err != nil {
			panic(err)
		}
		response.ValorConvertido = amount * rate
		convertTo := strings.ToUpper(c.Param("to"))
		if !implContains(currencies, convertTo) {
			return c.String(http.StatusBadRequest, "Currency "+c.Param("to")+" is invalid")
		}
		response.SimboloMoeda = currencySymbols[convertTo]
		jsonRespnse, err := json.Marshal(response)
		if err != nil {
			panic(err)
		}

		conversion := model.Conversion{}
		conversion.Amount = amount
		conversion.Rate = rate
		conversion.From = strings.ToUpper(c.Param("from"))
		conversion.To = strings.ToUpper(c.Param("to"))
		conversion.ConvertedValue = response.ValorConvertido

		db.Create(&conversion)

		return c.String(http.StatusOK, fmt.Sprint(string(jsonRespnse)))
	})
	server.Logger.Fatal(server.Start(":8000"))
}

func implContains(sl []string, name string) bool {
	for _, value := range sl {
		if value == name {
			return true
		}
	}
	return false
}
