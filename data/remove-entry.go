package data

import (
	"fmt"
	"stashtape/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func RemoveEntry(userId string, uid string) string {

	session := config.SessionAWS()
	service := dynamodb.New(session)

	updateExpression := "REMOVE entries[" + *aws.String(uid) + "]"
	conditionExpression := "attribute_exists(USER_ID)"

	updateInput := &dynamodb.UpdateItemInput{
		TableName: aws.String("USER_CONTENT"),
		Key: map[string]*dynamodb.AttributeValue{
			"USER_ID": {
				S: aws.String(userId),
			},
		},
		UpdateExpression:    aws.String(updateExpression),
		ConditionExpression: aws.String(conditionExpression),
	}

	_, err := service.UpdateItem(updateInput)
	if err != nil {
		fmt.Println("Error updating item:", err)
		return "error removing..."
	}

	fmt.Println("Successfully removed entry with UID:", uid)
	return "remove success"

}
