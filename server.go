package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/yuzuy/gqlgen-todos/db"
	"github.com/yuzuy/gqlgen-todos/graph"
	"github.com/yuzuy/gqlgen-todos/graph/generated"
)

const defaultPort = "8080"

func main() {
	rand.Seed(time.Now().UnixNano())
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		DB: db.New(),
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
