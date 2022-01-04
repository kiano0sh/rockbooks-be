package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
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

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowedOrigins:     []string{"*"},
		AllowCredentials:   true,
		AllowedHeaders:     []string{"*"},
		OptionsPassthrough: true,
		Debug:              true,
	}).Handler)

	router.Use(auth.Middleware())

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	router.Handle("/graphql", srv)
	fs := http.FileServer(http.Dir("media"))
	router.Handle("/media/*", http.StripPrefix("/media/", fs))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground 🚀", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
