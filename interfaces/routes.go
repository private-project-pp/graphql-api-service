package interfaces

import (
	"fmt"
	"log"
	"net/http"
)

func (c container) setupRoutesAndListen() (err error) {

	fmt.Println("Setup route and server")
	http.Handle("/", c.playground)
	http.Handle("/query", c.query)

	fmt.Println("Server started and listening")
	log.Printf("connect to PORT %s/ for GraphQL playground", c.config.Service.Port)
	log.Printf("connect to PORT %s/query for GraphQL Query", c.config.Service.Port)
	err = http.ListenAndServe(c.config.Service.Port, nil)
	if err != nil {
		return err
	}

	return nil
}
