package db

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func SessionAWS() *session.Session {

	awsCreds := os.Getenv("AWS_CREDS")
	awsCredsSecret := os.Getenv("AWS_CREDS_SECRET")

	session, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2"),
		Credentials: credentials.NewStaticCredentials(
			awsCreds,
			awsCredsSecret,
			"",
		),
	})

	if err != nil {
		fmt.Println("Error creating session: ", err)
	}

	return session
}
