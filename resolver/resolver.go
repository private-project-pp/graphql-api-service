package resolver

import "github.com/private-project-pp/graphql-api-service/infrastructure/user_service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	userInfra user_service.UserService
}

func SetupResolver(userInfra user_service.UserService) *Resolver {
	return &Resolver{
		userInfra: userInfra,
	}
}
