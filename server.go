package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"gitlab.com/kian00sh/rockbooks-be/graph"
	"gitlab.com/kian00sh/rockbooks-be/graph/generated"
	"gitlab.com/kian00sh/rockbooks-be/src/database/migrations"
	database "gitlab.com/kian00sh/rockbooks-be/src/database/postgresql"
	"gitlab.com/kian00sh/rockbooks-be/src/middlewares/auth"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Initialize database
	db := database.InitDB()
	// Running migrations
	migrations.InitMigrations(db)

	// Using chi in order to be able to consume our middleware
	router := chi.NewRouter()
	router.Use(auth.Middleware())

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground 🚀", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
