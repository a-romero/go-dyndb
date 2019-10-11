package records

import (
	"fmt"
	"log"
	"reflect"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// GenericRecordWriter ...
func GenericRecordWriter(object interface{}, objectType reflect.Type, tableName string, svc *dynamodb.DynamoDB) error {
	var err error

	items := reflect.ValueOf(object)

	if items.Kind() == reflect.Slice {
		for i := 0; i < items.Len(); i++ {
			item := items.Index(i)
			if item.Kind() == reflect.Struct {

				errProcessDynDBRecord := processDynDBRecord(item.Interface(), tableName, svc)
				if errProcessDynDBRecord != nil {
					log.Fatalf("error calling PutItem: %v", errProcessDynDBRecord)
					return errProcessDynDBRecord
				}
			} else {
				fmt.Println("Not found a struct!")
			}
		}
	} else {
		fmt.Println("Not found a slice")
	}
	return err
}

func processDynDBRecord(object interface{}, tableName string, svc *dynamodb.DynamoDB) error {
	av, errMarshal := dynamodbattribute.MarshalMap(object)
	if errMarshal != nil {
		log.Fatalf("marshalling record: %v", errMarshal)
		return errMarshal
	}

	record := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}
	_, errPutItem := svc.PutItem(record)
	if errPutItem != nil {
		log.Fatalf("error writing record: %v", errPutItem)
		return errPutItem
	}
	return nil
}

// CreateDynDBSvc ...
func CreateDynDBSvc() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := dynamodb.New(sess)
	return svc
}

// WriteRecord ...
func WriteRecord(object interface{}, tableName string, svc *dynamodb.DynamoDB) error {

	objectType := reflect.TypeOf(object)
	errWriteRecords := GenericRecordWriter(object, objectType, tableName, svc)

	return errWriteRecords
}
