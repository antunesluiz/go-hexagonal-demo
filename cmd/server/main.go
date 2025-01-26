package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	database "github.com/antunesluiz/go-hexagonal-demo/internal/adapters/database/config"
	"github.com/antunesluiz/go-hexagonal-demo/internal/adapters/graphql/generated"
	"github.com/antunesluiz/go-hexagonal-demo/internal/adapters/graphql/resolvers"
	"github.com/antunesluiz/go-hexagonal-demo/pkg/config"
	"github.com/vektah/gqlparser/v2/ast"
)

func main() {
	appConfig := config.LoadAppConfig()
	dbConfig := config.LoadDatabaseConfig()
	graphqlConfig := config.LoadGraphQLConfig()

	config := database.PostgresConfig{
		Host:     dbConfig.Host,
		Port:     dbConfig.Port,
		User:     dbConfig.User,
		Password: dbConfig.Password,
		DBName:   dbConfig.DBName,
		SSLMode:  dbConfig.SSLMode,
	}

	db, err := database.NewPostgresDB(config)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	srv := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](graphqlConfig.QueryCacheSize))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	if graphqlConfig.PlaygroundEnabled {
		http.Handle("/", playground.Handler("GraphQL playground", "/query"))

		log.Printf("Conecte-se a http://localhost:%s/ para acessar o playground GraphQL", appConfig.Port)
	}

	http.Handle("/query", srv)

	log.Fatal(http.ListenAndServe(":"+appConfig.Port, nil))
}
