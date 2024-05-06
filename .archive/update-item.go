package model

import (
	"encoding/json"
	"fmt"
	"stashtape/db"
	"stashtape/types"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func UpdateItem(table string, data types.CollectionItem) string {

	session := db.SessionAWS()

	service := dynamodb.New(session)
	tableName := aws.String(table)
	item := data
	fmt.Println(item.CollectionId)

	newItem := map[string]*dynamodb.AttributeValue{
		"Title": {
			S: aws.String(item.List[0].Title),
		},
		"Description": {
			S: aws.String(item.List[0].Description),
		},
	}

	newItemWrap := []*dynamodb.AttributeValue{{
		M: newItem,
	}}

	fmt.Println(newItemWrap)

	input := &dynamodb.UpdateItemInput{
		TableName: tableName,
		Key: map[string]*dynamodb.AttributeValue{
			"CollectionId": {
				S: aws.String(item.CollectionId),
			},
			"Timestamp": {
				S: aws.String(item.Timestamp),
			},
		},
		UpdateExpression: aws.String("SET Entries = list_append(Entries, :newItem)"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":newItem": {
				L: newItemWrap,
			},
		},
	}

	_, err := service.UpdateItem(input)
	if err != nil {
		fmt.Println("Error adding item:", err)

		responseFailed := types.Response{
			Status:  "failed",
			Message: "Failed to update item.",
		}

		responseJSON, err := json.Marshal(responseFailed)
		if err != nil {
			fmt.Println(err)
		}

		return string(responseJSON)
	}

	responseOK := types.Response{
		Status:  "ok",
		Message: "Collection item successfully updated.",
	}

	responseJSON, err := json.Marshal(responseOK)
	if err != nil {
		fmt.Println(err)
	}

	return string(responseJSON)
}
