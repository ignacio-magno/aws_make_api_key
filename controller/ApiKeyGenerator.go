package controller

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"os"
)

type ApiKeyGenerator struct {
	client *apigateway.APIGateway
}

func NewApiKeyGenerator() ApiKeyGenerator {
	// create new client apigateway and assign to var client
	mySession := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	}))

	client := apigateway.New(mySession)

	return ApiKeyGenerator{client: client}
}

func (a *ApiKeyGenerator) Generate(email string) (*apigateway.ApiKey, error) {
	// create api key with usage plan
	key, err := a.createApiKey(email)
	if err != nil {
		return nil, err
	}

	// assign api key to usage plan
	err = a.assignToUsagePlan(err, key)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func (a *ApiKeyGenerator) assignToUsagePlan(err error, key *apigateway.ApiKey) error {
	_, err = a.client.CreateUsagePlanKey(&apigateway.CreateUsagePlanKeyInput{
		KeyId:       key.Id,
		KeyType:     aws.String("API_KEY"),
		UsagePlanId: aws.String(os.Getenv("USAGE_PLAN_ID")),
	})
	return err
}

func (a *ApiKeyGenerator) createApiKey(email string) (*apigateway.ApiKey, error) {
	key, err := a.client.CreateApiKey(&apigateway.CreateApiKeyInput{
		Name:        aws.String(email),
		Description: aws.String("api key to clients"),
		Enabled:     aws.Bool(true),
	})
	return key, err
}
