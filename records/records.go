package records

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

//Configuration file structure
type Configuration struct {
	Region     string
	AWSProfile string
	TableName  string
}

//GenericRecord ...
type GenericRecord struct {
	Data AnyRecord `json:"data"`
}

//AnyRecord ...
type AnyRecord interface {
	writeRecords(tableName string, svc *dynamodb.DynamoDB)
}
