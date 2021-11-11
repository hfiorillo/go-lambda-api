package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Declare constants
const (
	REGION     = "eu-north-1"
	TABLE_NAME = "Books"
)

// Declare new Dynamodb instance - safe for concurrent use

var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion(REGION))

// New method of declearing dynamodb instnace
// var mySession = session.Must(session.NewSession())
// db := dynamodb.New(mySession, aws.NewConfig().WithRegion(REGION)))

func getItem(isbn string) (*book, error) {
	// Prepare input for the query.
	input := &dynamodb.GetItemInput{
		TableName: aws.String(TABLE_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"ISBN": {
				S: aws.String(isbn),
			},
		},
	}

	// Retrive the item from Dynamodb. If no item is found return nil.
	result, err := db.GetItem(input)
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, nil
	}

	// The result.Item object returned has the underlying type
	// map[string]*AttributeValue. We can use the UnmarshallMap
	// helper to parse this straight into the fields of a struct!
	// UnmarshallListOfMaps exists if working with multiple items.

	bk := new(book)
	err = dynamodbattribute.UnmarshalMap(result.Item, bk)
	if err != nil {
		return nil, err
	}

	return bk, nil
}

// Add a record to DynamoDB
func putItem(bk *book) error {
	input := &dynamodb.PutItemInput{
		TableName: aws.String("Books"),
		Item: map[string]*dynamodb.AttributeValue{
			"ISBN": {
				S: aws.String(bk.ISBN),
			},
			"Title": {
				S: aws.String(bk.Title),
			},
			"Author": {
				S: aws.String(bk.Author),
			},
		},
	}
	_, err := db.PutItem(input)
	return err
}
