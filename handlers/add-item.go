package handler

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"os"
	"stashtape/types"
)

func AddItem(table string, data types.CollectionItem) string {

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
	}

	service := dynamodb.New(session)
	tableName := aws.String(table)
	item := data

	collectionItem, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		fmt.Println("Error marshalling data:", err)

		responseFailed := types.Response{
			Status:  "failed",
			Message: "Failed marshalling data.",
		}

		responseJSON, err := json.Marshal(responseFailed)

		if err != nil {
			fmt.Println(err)
		}

		return string(responseJSON)
	}

	input := &dynamodb.PutItemInput{
		TableName: tableName,
		Item:      collectionItem,
	}

	_, err = service.PutItem(input)
	if err != nil {
		fmt.Println("Error adding item:", err)

		responseFailed := types.Response{
			Status:  "failed",
			Message: "Failed to item.",
		}

		responseJSON, err := json.Marshal(responseFailed)
		if err != nil {
			fmt.Println(err)
		}

		return string(responseJSON)
	}

	responseOK := types.Response{
		Status:  "ok",
		Message: "Collection item successfully added.",
	}

	responseJSON, err := json.Marshal(responseOK)
	if err != nil {
		fmt.Println(err)
	}

	return string(responseJSON)
}
