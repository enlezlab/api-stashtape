package handler

import (
	"fmt"
	"log"
	"stashtape/db"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func DeleteItem(tableName string, id string) {
	session := db.SessionAWS()
	service := dynamodb.New(session)

	collectionId := aws.String(id)
	tn := aws.String(tableName)

	// TODO: remove test code, replace with the steps below
	// - make a get item request
	// - if item not exist, exit. if exist, continute steps below
	// - pass the result of get item to DeleteInput method
	// - make a delete request
	input := &dynamodb.DeleteItemInput{
		TableName: tn,
		Key: map[string]*dynamodb.AttributeValue{
			"CollectionId": {S: collectionId},
			"Timestamp":    {S: aws.String("1713122179")},
		},
	}

	_, err := service.DeleteItem(input)
	if err != nil {
		log.Fatalf("Got error calling DeleteItem: %s", err)
	}

	fmt.Printf("Item %s deleted", collectionId)

}
