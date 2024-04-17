package handler

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"os"
	"stashtape/types"
)

func GetItem(tableName string, valCond string) []byte {

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

	var resStruct []types.CollectionItem

	for _, item := range result.Items {

		item := types.CollectionItem{
			CollectionId: *item["CollectionId"].S,
			Timestamp:    *item["Timestamp"].S,
		}

		resStruct = append(resStruct, item)
	}

	res, err := json.Marshal(resStruct)
	if err != nil {
		fmt.Println(err)
	}

	return res
}
