package test

import (
	"aws_make_api_key/controller"
	"aws_make_api_key/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPingDatabase(t *testing.T) {
	cP := controller.NewPersistance()
	err := cP.Save(domain.EmailPersistance{
		Email:     "this is testing email",
		CreatedAt: time.Now(),
	})

	assert.NoError(t, err)
}
