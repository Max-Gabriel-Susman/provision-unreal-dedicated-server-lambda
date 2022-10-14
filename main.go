package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	log.Println("Goliath Online")

	response := events.APIGatewayProxyResponse{
		StatusCode: 209,
	}

	return &response, nil
}
