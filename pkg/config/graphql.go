package config

import (
	"strconv"

	"github.com/joho/godotenv"
)

type GraphQLConfig struct {
	PlaygroundEnabled bool
	QueryCacheSize    int
}

func LoadGraphQLConfig() *GraphQLConfig {
	_ = godotenv.Load()

	enabled := getEnv("GRAPHQL_PLAYGROUND_ENABLED", "true") == "true"
	cacheSize, _ := strconv.Atoi(getEnv("GRAPHQL_QUERY_CACHE_SIZE", "1000"))

	return &GraphQLConfig{
		PlaygroundEnabled: enabled,
		QueryCacheSize:    cacheSize,
	}
}
