package store

import (
	"fmt"
	"stashtape/db"
	"stashtape/types"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func GetItem(tableName string, valCond string) []types.CollectionItem {

	session := db.SessionAWS()

	service := dynamodb.New(session)

	tn := aws.String(tableName)
	vc := aws.String(valCond)

	keyCondExp := "#CollectionId = :collectionId"
	expAttrName := map[string]*string{
		"#CollectionId": aws.String("CollectionId"),
	}

	expAttrVal := map[string]*dynamodb.AttributeValue{
		":collectionId": {
			S: vc,
		},
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

		list := []types.CollectionItemList{}

		for _, colItem := range item["Entries"].L {
			fmt.Println(colItem)

			collectionItemList := types.CollectionItemList{
				Title:       *colItem.M["Title"].S,
				Description: *colItem.M["Description"].S,
			}

			list = append(list, collectionItemList)
		}

		item := types.CollectionItem{
			CollectionId: *item["CollectionId"].S,
			Timestamp:    *item["Timestamp"].S,
			List:         list,
		}

		resStruct = append(resStruct, item)
	}

	return resStruct
}
