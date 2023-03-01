package main

import (
	"aws_make_api_key/domain"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	payload, err := domain.NewRequest(e)
	if err != nil {
		return returnError(err)
	}

	_ = payload

	return events.APIGatewayProxyResponse{
		Body:       "Hello, World!",
		StatusCode: 200,
	}, nil
}

func returnError(err error) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       err.Error(),
		StatusCode: 400,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
