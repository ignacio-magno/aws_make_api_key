package test

import (
	"aws_make_api_key/controller"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNotifyByEmail(t *testing.T) {
	controllerSES := controller.NewSes()
	email := os.Getenv("EMAIL_TEST")
	keyToken := "testKeyToken"

	err := controllerSES.NotifyByEmail(email, keyToken)
	assert.NoError(t, err)
}
