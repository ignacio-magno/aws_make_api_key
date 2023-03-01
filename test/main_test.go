package test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	LoadEnv()
	m.Run()
}

func TestLoadEnv(t *testing.T) {
	env := os.Getenv("DynamoDBTable")
	assert.NotEmpty(t, env)
}

func LoadEnv() {
	var keys map[string]string
	file, err := os.ReadFile("../env.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(file, &keys)
	if err != nil {
		panic(err)
	}

	for key, value := range keys {
		err = os.Setenv(key, value)
		if err != nil {
			panic(err)
		}
	}
}
