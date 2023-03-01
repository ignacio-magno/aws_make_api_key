package domain

import "time"

type EmailPersistance struct {
	Email     string    `dynamodbav:"email"`
	CreatedAt time.Time `dynamodbav:"created_at"`
}

func NewEmailPersistance(email string) EmailPersistance {
	return EmailPersistance{
		Email:     email,
		CreatedAt: time.Now(),
	}
}
