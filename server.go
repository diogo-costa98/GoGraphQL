package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/diogo-costa98/GoGraphQL/auth"
	"github.com/diogo-costa98/GoGraphQL/graph"
	sqlite "github.com/diogo-costa98/GoGraphQL/internal/db/sqlite"
	"github.com/go-chi/chi"
)

const (
	host        = "localhost"
	port        = 5432
	user        = "postgres"
	password    = "DiAnMaEd22"
	dbname      = "TogglHire-BackendHomework"
	defaultPort = "3000"
)

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	router.Use(auth.Middleware())

	sqlite.InitDB(psqlconn)
	defer sqlite.CloseDB()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
