package test

import (
	"aws_make_api_key/domain"
	"encoding/base64"
	"github.com/aws/aws-lambda-go/events"
	"github.com/badoux/checkmail"
	"github.com/stretchr/testify/assert"
	"testing"
)

const payload = `{
		"email": "ignacio@gmail.com"
	}`

func TestUnmarshalOnIsBase64(t *testing.T) {

	payloadBase64 := base64.StdEncoding.EncodeToString([]byte(payload))

	request, err := domain.NewRequest(events.APIGatewayProxyRequest{
		Body:            payloadBase64,
		IsBase64Encoded: true,
	})

	assert.Nil(t, err)
	assert.Equal(t, "ignacio@gmail.com", request.Email)
}

func TestUnmarshalOnNotIsBase64(t *testing.T) {
	request, err := domain.NewRequest(events.APIGatewayProxyRequest{
		Body:            payload,
		IsBase64Encoded: false,
	})

	assert.Nil(t, err)
	assert.Equal(t, "ignacio@gmail.com", request.Email)
}

func TestOnInvalidEmail(t *testing.T) {
	_, err := domain.NewRequest(events.APIGatewayProxyRequest{
		Body:            `{"email": "ignacio"}`,
		IsBase64Encoded: false,
	})

	assert.Equal(t, checkmail.ErrBadFormat, err)
}
