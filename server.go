package main

import (
	"log"

	"github.com/private-project-pp/graphql-api-service/interfaces/server"
	_ "github.com/spf13/viper"
)

func main() {
	err := server.StartServer()
	if err != nil {
		log.Println(err)
	}
}
