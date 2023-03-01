package domain

import "time"

type EmailPersistance struct {
	Email     string    `dynamodbav:"email"`
	CreatedAt time.Time `dynamodbav:"created_at"`
}
