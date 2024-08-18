package server

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/private-project-pp/graphql-api-service/interfaces/graph_server"
	"github.com/private-project-pp/graphql-api-service/resolver"
	"github.com/private-project-pp/graphql-api-service/shared/config"
)

func StartServer() (err error) {
	var resolvers *resolver.Resolver
	configs := config.SetupConfig()

	srv := handler.NewDefaultServer(graph_server.NewExecutableSchema(graph_server.Config{Resolvers: resolvers}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost%s/ for GraphQL playground", config.Service.Port)
	err = http.ListenAndServe(configs.Service.Port, nil)
	if err != nil {
		panic(err)
	}
	return nil
}
