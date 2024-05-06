package model

import (
	"fmt"
	"stashtape/db"
	"stashtape/types"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func GetEntry(tableName string, valCond string) []types.User {

	session := db.SessionAWS()

	service := dynamodb.New(session)

	tn := aws.String(tableName)
	vc := aws.String(valCond)

	keyCondExp := "#USER_ID = :userId"
	expAttrName := map[string]*string{
		"#USER_ID": aws.String("USER_ID"),
	}

	expAttrVal := map[string]*dynamodb.AttributeValue{
		":userId": {
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

	var resStruct []types.User

	for _, item := range result.Items {

		entries := []types.Entries{}

		for _, colItem := range item["ENTRIES"].L {

			entryItem := types.Entries{
				UID:         *colItem.M["UID"].S,
				Title:       *colItem.M["Title"].S,
				Description: *colItem.M["Description"].S,
			}

			entries = append(entries, entryItem)
		}

		resItem := types.User{
			USER_ID: *item["USER_ID"].S,
			ENTRIES: entries,
		}

		resStruct = append(resStruct, resItem)
	}

	return resStruct
}
