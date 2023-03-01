package controller

import (
	"aws_make_api_key/domain"
	"fmt"
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

func (p *Persistance) Exist(email string) bool {
	doc, err := p.repo.FindOne([]interface{}{"email"})
	if err == fmt.Errorf("not found") {
		return false
	}

	if err != nil {
		panic(err)
	}

	return doc.Email == email
}
