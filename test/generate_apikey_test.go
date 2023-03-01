package test

import (
	"aws_make_api_key/controller"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateAPIKey(t *testing.T) {
	generateApiKey := controller.NewApiKeyGenerator()
	apiKey, err := generateApiKey.Generate("test_email")

	_ = apiKey

	assert.NoError(t, err)
}
