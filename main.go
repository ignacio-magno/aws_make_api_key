package main

import (
	"aws_make_api_key/controller"
	"aws_make_api_key/domain"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

func Handler(e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	payload, err := domain.NewRequest(e)
	if err != nil {
		return returnError(err)
	}

	log.Default().Println("finding email in register")
	// exist email in register
	controllerPersistance := controller.NewPersistance()
	if controllerPersistance.Exist(payload.Email) {
		return returnError(domain.ErrEmailAlreadyExist)
	}

	log.Default().Println("saving email in register")
	// save email in register
	err = controllerPersistance.Save(domain.NewEmailPersistance(payload.Email))
	if err != nil {
		return returnError(err)
	}

	log.Default().Println("generating token")
	// generate token in api gateway
	controllerApiKey := controller.NewApiKeyGenerator()
	apiKey, err := controllerApiKey.Generate(payload.Email)
	if err != nil {
		return returnError(err)
	}

	log.Default().Println("notifying by sns")
	// notify by sns
	controllerSns := controller.NewSes()
	err = controllerSns.NotifyByEmail(payload.Email, *apiKey.Value)
	if err != nil {
		return returnError(err)
	}

	return events.APIGatewayProxyResponse{
		Body:       "ok",
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
