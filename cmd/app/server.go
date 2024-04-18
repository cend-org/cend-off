package main

import (
	"duval/internal/authentication"
	"duval/internal/graph"
	"duval/pkg/database"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {

	router := chi.NewRouter()
	router.Use(authentication.AuthMiddleware())

	defer func(Client *sqlx.DB) {
		err := Client.Close()
		if err != nil {
			return
		}
	}(database.Client)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            false,
	})

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", c.Handler(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
