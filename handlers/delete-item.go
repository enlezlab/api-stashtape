package handler

import (
	"fmt"
	"log"
	"stashtape/db"
	"stashtape/store"
	"stashtape/types"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func DeleteItem(tableName string, id string) types.Response {
	session := db.SessionAWS()
	service := dynamodb.New(session)

	item := store.GetItem(tableName, id)

	if len(item) == 0 {
		fmt.Println("No matching item")

		return types.Response{
			Status:  "Failed",
			Message: "No matching item.",
		}
	}

	collectionId := item[0].CollectionId
	timestamp := item[0].Timestamp

	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"CollectionId": {S: aws.String(collectionId)},
			"Timestamp":    {S: aws.String(timestamp)},
		},
	}

	_, err := service.DeleteItem(input)
	if err != nil {
		log.Fatalf("Got error calling DeleteItem: %s", err)
	}

	fmt.Printf("Item %s deleted", id)

	return types.Response{
		Status:  "OK",
		Message: "Item deleted",
	}

}
