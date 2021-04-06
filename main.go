package main

import (
	"fmt"
	"log"

	"github.com/Vractos/delivery-simulator/application/route"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	route := route.Route{
		ID:       "1",
		ClientID: "1",
	}

	route.LoadPositions()
	stringJson, _ := route.ExportJsonPosition()
	fmt.Println(stringJson[1])
}
