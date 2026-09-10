package main

import (
	"context"
	"log"
	"subscriptionTracker/internal/config"
	"subscriptionTracker/internal/db"
	"subscriptionTracker/internal/server"
	"subscriptionTracker/internal/subscriptions"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"

	_ "github.com/lib/pq"
)

var ginLambda *ginadapter.GinLambdaV2

func init() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	dbConn, err := db.Connect(context.Background(), cfg.DatabaseUrl)
	if err != nil {
		log.Fatalf("error connect to database: %v", err)
	}

	repo := subscriptions.NewRespository(dbConn)
	handler := subscriptions.NewHandler(repo)
	r := server.NewRouter(handler)

	ginLambda = ginadapter.NewV2(r)
}

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
