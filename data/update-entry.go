package data

import (
	"encoding/json"
	"fmt"
	"stashtape/config"
	"stashtape/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func UpdateEntry(table string, data model.User) string {

	session := config.SessionAWS()

	service := dynamodb.New(session)
	tableName := aws.String(table)
	item := data

	newItem := map[string]*dynamodb.AttributeValue{
		"UID": {
			S: aws.String(item.ENTRIES[0].UID),
		},
		"Title": {
			S: aws.String(item.ENTRIES[0].Title),
		},
		"Description": {
			S: aws.String(item.ENTRIES[0].Description),
		},
	}

	newItemWrap := []*dynamodb.AttributeValue{{
		M: newItem,
	}}

	input := &dynamodb.UpdateItemInput{
		TableName: tableName,
		Key: map[string]*dynamodb.AttributeValue{
			"USER_ID": {
				S: aws.String(item.USER_ID),
			},
		},
		UpdateExpression: aws.String("SET ENTRIES = list_append(ENTRIES, :newItem)"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":newItem": {
				L: newItemWrap,
			},
		},
	}

	_, err := service.UpdateItem(input)
	if err != nil {
		fmt.Println("Error adding item:", err)

		responseFailed := model.Response{
			Status:  "failed",
			Message: "Failed to update item.",
		}

		responseJSON, err := json.Marshal(responseFailed)
		if err != nil {
			fmt.Println(err)
		}

		return string(responseJSON)
	}

	responseOK := model.Response{
		Status:  "ok",
		Message: "Collection item successfully updated.",
	}

	responseJSON, err := json.Marshal(responseOK)
	if err != nil {
		fmt.Println(err)
	}

	return string(responseJSON)
}
