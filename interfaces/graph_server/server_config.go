package graph_server

func SetupServerConfig(resolver ResolverRoot) Config {
	return Config{
		Resolvers: resolver,
	}
}
