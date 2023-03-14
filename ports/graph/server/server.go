package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/slaff-bg/stockroom/adapters"
	"github.com/slaff-bg/stockroom/ports/graph"
)

const defaultHTTPPort = "8080"

type httpServer struct {
	db       adapters.DBRepo
	httpPort string
}

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	// Set httpServer properties.
	var httpSet httpServer
	httpSet.httpPort = os.Getenv("GRAPHQL_PORT")
	if httpSet.httpPort == "" {
		httpSet.httpPort = defaultHTTPPort
	}
	httpSet.db = adapters.DBRepo{}
	httpSet.db.Connect()

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{DBConn: httpSet.db.DB.INST}}))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", httpSet.httpPort)
	log.Fatal(http.ListenAndServe(":"+httpSet.httpPort, nil))
}
