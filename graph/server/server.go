package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	generated "github.com/AI1411/go-grpc-praphql/graph"
	grpcUser "github.com/AI1411/go-grpc-praphql/graph/grpc"
	"github.com/AI1411/go-grpc-praphql/graph/resolver"
)

const defaultPort = "8081"

func main() {
	port := os.Getenv("STAR_GRAPHQL_PORT")
	if port == "" {
		port = defaultPort
	}

	userClient, err := grpcUser.ConnectUserServiceClient()
	if err != nil {
		log.Fatalf("failed to connect to user server: %v", err)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{
		UserClient: userClient,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
