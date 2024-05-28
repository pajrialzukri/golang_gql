package main

import (
	"backend/graph/generated"
	graph "backend/graph/resolvers"
	"backend/pkg/config"
	"backend/pkg/database"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
)

const defaultPort = "8080"

func main() {
	config.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	if len(os.Args) == 2 {
		port = os.Args[1]
	}
	router := chi.NewRouter()
	database.InitDB()
	server := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
