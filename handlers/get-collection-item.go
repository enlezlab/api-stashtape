package handler

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"os"
)

func GetCollectionItem() {

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

	tableName := aws.String("collection")
	collectionId := aws.String("ST007")

	keyCondExp := "#CollectionId = :collectionId"
	expAttrName := map[string]*string{
		"#CollectionId": aws.String("CollectionId"),
	}

	expAttrVal := map[string]*dynamodb.AttributeValue{
		":collectionId": {S: collectionId},
	}

	input := &dynamodb.QueryInput{
		TableName:                 tableName,
		KeyConditionExpression:    aws.String(keyCondExp),
		ExpressionAttributeNames:  expAttrName,
		ExpressionAttributeValues: expAttrVal,
	}

	result, err := service.Query(input)
	if err != nil {
		fmt.Println("Query Error: ", err)
		return
	}

	for _, item := range result.Items {
		fmt.Println("===Query Result===")
		fmt.Println(item)
		fmt.Println("===End of Query Result===")
	}

}
