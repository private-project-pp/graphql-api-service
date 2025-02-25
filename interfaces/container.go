package interfaces

import (
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/private-project-pp/graphql-api-service/interfaces/graph_server"
	"github.com/private-project-pp/graphql-api-service/resolver"
	"github.com/private-project-pp/graphql-api-service/shared/config"
)

type container struct {
	query      *handler.Server
	playground http.HandlerFunc
	config     config.ConfigApp
}

func Container() (err error) {

	fmt.Println("Start container")
	configs := config.SetupConfig()

	// setup infrastructure(s) here

	resolvers := resolver.SetupResolver()

	server_config := graph_server.SetupServerConfig(resolvers)

	gs := graph_server.NewExecutableSchema(server_config)

	srv := handler.NewDefaultServer(gs)

	serverInstance := container{
		config:     configs,
		query:      srv,
		playground: playground.Handler("GraphQL playground", "/query"),
	}

	err = serverInstance.setupRoutesAndListen()
	if err != nil {
		return err
	}
	return nil
}
