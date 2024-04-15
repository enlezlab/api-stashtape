package handler

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"os"
	"stashtape/types"
)

func AddItem(data types.CollectionItem) {

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
		fmt.Println("Error creating session:", err)
		return
	}

	service := dynamodb.New(session)
	item := data

	tableName := aws.String("collection")
	collectionItem, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		fmt.Println("Error marshalling data:", err)
		return
	}

	input := &dynamodb.PutItemInput{
		TableName: tableName,
		Item:      collectionItem,
	}

	_, err = service.PutItem(input)
	if err != nil {
		fmt.Println("Error adding item:", err)
		return
	}

	fmt.Println("Item added successfully!")

}
