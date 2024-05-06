package handler

import (
	"encoding/json"
	"fmt"
	"stashtape/db"
	"stashtape/types"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// TODO:
// Refactor the dynamodb call to store package
func GetCollection(tableName string) []byte {

	session := db.SessionAWS()

	service := dynamodb.New(session)
	tn := tableName

	input := &dynamodb.ScanInput{
		TableName: aws.String(tn),
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

	return res
}
