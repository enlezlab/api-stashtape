package data

import (
	"encoding/json"
	"fmt"
	"stashtape/config"
	"stashtape/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func NewEntry(table string, data model.User) string {

	session := config.SessionAWS()

	service := dynamodb.New(session)
	tableName := aws.String(table)
	item := data

	collectionItem, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		fmt.Println("Error marshalling data:", err)

		responseFailed := model.Response{
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

		responseFailed := model.Response{
			Status:  "failed",
			Message: "Failed to item.",
		}

		responseJSON, err := json.Marshal(responseFailed)
		if err != nil {
			fmt.Println(err)
		}

		return string(responseJSON)
	}

	responseOK := model.Response{
		Status:  "ok",
		Message: "Collection item successfully added.",
	}

	responseJSON, err := json.Marshal(responseOK)
	if err != nil {
		fmt.Println(err)
	}

	return string(responseJSON)
}
