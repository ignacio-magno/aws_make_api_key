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
	return &Persistance{repo: dynamo.NewRepositoryDynamo[domain.EmailPersistance](os.Getenv("DynamoDBTable"), false)}
}

func (p *Persistance) Save(emailPersistance domain.EmailPersistance) error {
	return p.repo.SaveOrReplace(emailPersistance)
}

func (p *Persistance) Get(email string) (domain.EmailPersistance, error) {
	doc, err := p.repo.FindOne([]interface{}{"email"})
	return doc, err
}
