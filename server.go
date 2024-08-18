package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"

	"github.com/MukuFlash03/hackernews/graph"
	// "github.com/MukuFlash03/hackernews/graph/generated"
	// "github.com/MukuFlash03/hackernews/internal/auth"
	// _ "github.com/MukuFlash03/hackernews/internal/auth"
	database "github.com/MukuFlash03/hackernews/internal/pkg/db/postgres"

)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	database.InitDB()
	defer database.CloseDB()
	database.Migrate()

	server := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))	
}
