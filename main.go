package main

import (
	"flag"
	"log"
	"os"

	"github.com/aflesch/sendMail/httpMail"
)

const senderAddress = "avner@minutelab.io"
const mailSubject = "Welcome MinuteLab user"

func main() {
	awsRegion := flag.String("region", os.Getenv("AWS_REGION"), "aws region working with")
	awsKey := flag.String("key", os.Getenv("AWS_KEY_ID"), "aws key id")
	awsSecret := flag.String("secret", os.Getenv("AWS_SECRET_KEY"), "aws secret")

	template := flag.String("template", "template.html", "html template file")
	name := flag.String("name", "", "user name")
	recipientMail := flag.String("email", "", "recipient email address")

	flag.Parse()

	// Send mail from (const - hard coded)
	// get source and dest email
	if body, err := httpMail.CreateMessage(*template, *name); err != nil {
		log.Fatalf("Fatal: %s\n", err)
	} else {
		err = httpMail.SendMail(httpMail.SendMailConfig{
			Subject:   mailSubject,
			Body:      body,
			Region:    *awsRegion,
			Key:       *awsKey,
			Secret:    *awsSecret,
			Sender:    senderAddress,
			Recipient: *recipientMail,
		})
		log.Printf("sendMail [error=%s]\n", err)
	}
}
