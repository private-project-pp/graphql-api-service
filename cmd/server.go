package cmd

import "github.com/private-project-pp/graphql-api-service/interfaces"

func StartServer() {
	err := interfaces.Container()
	if err != nil {
		panic(err)
	}
}
