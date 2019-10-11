package records

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)

// RecordLookup ...
func RecordLookup(table string, attrToGet string, conditionLeft string, conditionRight string) (*dynamodb.QueryOutput, error) {

	svc := CreateDynDBSvc()

	params := &dynamodb.QueryInput{
		TableName: aws.String(table),
		AttributesToGet: []*string{
			aws.String(attrToGet),
		},
		KeyConditions: map[string]*dynamodb.Condition{
			conditionLeft: &dynamodb.Condition{
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					&dynamodb.AttributeValue{
						S: aws.String(conditionRight),
					},
				},
			},
		},
	}

	result, err := svc.Query(params)
	if err != nil {
		return result, recordReadError(err)
	}

	return result, nil
}

func recordReadError(err error) error {
	errorLogger.Printf("DynamoDB error reading record: %v", err.Error())

	return err
}
