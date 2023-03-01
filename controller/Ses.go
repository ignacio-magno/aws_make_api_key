package controller

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"os"
	"strings"
)

type Ses struct {
	client *ses.SES
}

func NewSes() Ses {
	mySession := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	}))

	client := ses.New(mySession)
	return Ses{client: client}
}

func (s *Ses) NotifyByEmail(email string, token string) error {
	dataString, err := s.makeEmailHtmlBody(token)
	if err != nil {
		return err
	}

	sendEmail, err := s.client.SendEmail(&ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(email),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Data: aws.String(dataString),
				},
			},
			Subject: &ses.Content{
				Data: aws.String(os.Getenv("EMAIL_SUBJECT")),
			},
		},
		Source: aws.String(os.Getenv("EMAIL_SOURCE")),
	})

	_ = sendEmail
	return err
}

func (s *Ses) makeEmailHtmlBody(token string) (string, error) {
	data, err := os.ReadFile("../emailTemplate/build_production/notify_token.html")
	if err != nil {
		return "", err
	}

	dataString := string(data)
	dataString = strings.Replace(dataString, "{token}", token, 1)
	dataString = strings.Replace(dataString, "{emailResponse}", os.Getenv("EmailToNotifyErrorsApi"), 1)

	return dataString, nil
}
