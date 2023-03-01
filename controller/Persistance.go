package controller

import (
	"aws_make_api_key/domain"
	"github.com/ignacio-magno/database/dynamo"
	"os"
)

type Persistance struct {
	repo *dynamo.Repository[domain.EmailPersistance]
}

func NewPersistance() *Persistance {
	return &Persistance{repo: dynamo.NewRepositoryDynamo[domain.EmailPersistance](os.Getenv("TABLE_NAME"), false)}
}

func () Save(emailPersistance domain.EmailPersistance) error {
	return p.repo.Save(emailPersistance)
}
