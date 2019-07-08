package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

//Email struct for send
type Email struct {
	From    string // From source email
	To      string // To destination email(s)
	Subject string // Subject text to send
	Text    string // Text is the text body representation
	HTML    string // HTMLBody is the HTML body representation
	ReplyTo string // Reply-To email(s)
}

// SetConfiguration - aws configuration for SES
func SetConfiguration(awsKeyID string, awsSecretKey string, awsRegion string) {
	os.Setenv("AWS_REGION", awsRegion)
	os.Setenv("AWS_ACCESS_KEY_ID", awsKeyID)
	os.Setenv("AWS_SECRET_ACCESS_KEY", awsSecretKey)
}

//SendEmail - send mail throught aws ses
func SendEmail(emailData Email) *ses.SendEmailOutput {

	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
	}
	svc := ses.New(sess)

	params := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{aws.String(emailData.To)},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Data:    aws.String(emailData.Text),
					Charset: aws.String("UTF-8"),
				},
			},
			Subject: &ses.Content{
				Data:    aws.String(emailData.Subject),
				Charset: aws.String("UTF-8"),
			},
		},
		Source:           aws.String(emailData.From),
		ReplyToAddresses: []*string{aws.String(emailData.ReplyTo)},
	}

	// send email
	resp, err := svc.SendEmail(params)

	if err != nil {
		fmt.Println(err.Error())
	}
	return resp
}
