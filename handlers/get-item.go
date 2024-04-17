package handler

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"os"
)

func GetItem(tableName string, valCond string) string {

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

	service := dynamodb.New(session)

	tn := aws.String(tableName)
	vc := aws.String(valCond)

	keyCondExp := "#CollectionId = :collectionId"
	expAttrName := map[string]*string{
		"#CollectionId": aws.String("CollectionId"),
	}

	expAttrVal := map[string]*dynamodb.AttributeValue{
		":collectionId": {S: vc},
	}

	input := &dynamodb.QueryInput{
		TableName:                 tn,
		KeyConditionExpression:    aws.String(keyCondExp),
		ExpressionAttributeNames:  expAttrName,
		ExpressionAttributeValues: expAttrVal,
	}

	result, err := service.Query(input)
	if err != nil {
		fmt.Println("Query Error: ", err)
	}

	type ResponseItem struct {
		Id   string `json: "id"`
		Hash string `json: "hash"`
	}

	var res []ResponseItem

	for _, item := range result.Items {

		item := ResponseItem{
			Id:   *item["CollectionId"].S,
			Hash: *item["CollectionHash"].S,
		}

		res = append(res, item)
	}

	resJSON, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
	}

	return string(resJSON)

}
