package domain

import (
	"encoding/base64"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/badoux/checkmail"
)

type Request struct {
	Email string `json:"email"`
}

func NewRequest(e events.APIGatewayProxyRequest) (*Request, error) {
	var (
		body string
		r    Request
		err  error
	)

	body, err = transformPayload(e)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(body), &r)
	if err != nil {
		return nil, err
	}

	return &r, checkmail.ValidateFormat(r.Email)
}

func transformPayload(e events.APIGatewayProxyRequest) (string, error) {
	var body string
	if e.IsBase64Encoded {
		bodyByte, err := base64.StdEncoding.DecodeString(e.Body)
		if err != nil {
			return "", err
		}

		body = string(bodyByte)
	} else {
		body = e.Body
	}
	return body, nil
}
