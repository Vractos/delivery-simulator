package main

import (
	"fmt"

	"github.com/Vractos/delivery-simulator/application/route"
)

func main() {
	route := route.Route{
		ID:       "1",
		ClientID: "1",
	}

	route.LoadPositions()
	stringJson, _ := route.ExportJsonPosition()
	fmt.Println(stringJson[1])
}
