package controller

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"os"
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
	templatedEmail, err := s.client.SendTemplatedEmail(&ses.SendTemplatedEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(email),
			},
		},
		Source:   aws.String(os.Getenv("EMAIL_SOURCE")),
		Template: aws.String(os.Getenv("TEMPLATE_EMAIL_SES_NAME")),
		TemplateData: aws.String(`{
				"token": "` + token + `",
				"emailResponse": "` + os.Getenv("EmailToNotifyErrorsApi") + `"
			}`),
	})
	if err != nil {
		return err
	}

	_ = templatedEmail
	return err
}
