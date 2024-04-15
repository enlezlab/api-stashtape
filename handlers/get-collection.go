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

func GetCollection() string {

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
		fmt.Println("Error creating sessoin: ", err)
	}

	service := dynamodb.New(session)

	input := &dynamodb.ScanInput{
		TableName: aws.String("collection"),
	}

	var allItems []map[string]*dynamodb.AttributeValue

	for {
		result, err := service.Scan(input)
		if err != nil {
			fmt.Println("Error scaning table", err)
		}

		for _, item := range result.Items {
			allItems = append(allItems, item)
		}

		if len(result.LastEvaluatedKey) == 0 {
			break
		}
	}

	var collection []types.CollectionItem

	for _, item := range allItems {

		c := types.CollectionItem{
			CollectionId: *item["CollectionId"].S,
			Timestamp:    *item["Timestamp"].S,
		}

		collection = append(collection, c)
	}

	res, err := json.Marshal(collection)
	if err != nil {
		fmt.Println(err)
	}

	return string(res)
}

func GetCollectionItemById(id string) {
	fmt.Println(id)
	fmt.Println("TODO: implement get collection by id")
}
