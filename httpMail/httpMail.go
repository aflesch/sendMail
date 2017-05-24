package httpMail

import (
	"bytes"
	"html/template"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type SendMailConfig struct {
	Subject   string
	Body      string
	Region    string
	Key       string
	Secret    string
	Sender    string
	Recipient string
}

func CreateMessage(templateFile, name string) (string, error) {
	type Person struct {
		Name string //exported field since it begins with a capital letter
	}

	t, err := template.ParseFiles(templateFile)
	if err != nil {
		return "", err
	}
	p := Person{Name: name}
	buf := new(bytes.Buffer)
	t.Execute(buf, p)
	return buf.String(), nil
}

func SendMail(config SendMailConfig) error {
	awsSession, err := session.NewSession(&aws.Config{
		Region:      aws.String(config.Region),
		Credentials: credentials.NewStaticCredentials(config.Key, config.Secret, ""),
	})
	if err != nil {
		return err
	}

	sesSession := ses.New(awsSession)
	sesEmailInput := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{aws.String(config.Recipient)},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Data: aws.String(config.Body)},
			},
			Subject: &ses.Content{
				Data: aws.String(config.Subject),
			},
		},
		Source: aws.String(config.Sender),
		ReplyToAddresses: []*string{
			aws.String(config.Sender),
		},
	}

	_, err = sesSession.SendEmail(sesEmailInput)
	return err
}
