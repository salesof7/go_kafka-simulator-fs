package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	route "github.com/salesof7/go_kafka-simulator-fs/app/routes"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	route := route.Route{
		ID:       "1",
		ClientID: "1",
	}
	route.LoadPositions()
	stringJson, _ := route.ExportJsonPositions()
	fmt.Println(stringJson[1])
}