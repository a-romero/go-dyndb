package records

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/stretchr/testify/assert"
)

func TestRecordLookup(t *testing.T) {

	// VinItem ...
	type IMEIItem struct {
		IMEI string `dynamo:"imei"`
	}
	var imei string

	tests := []struct {
		requestTableName      string
		requestAttrToGet      string
		requestConditionLeft  string
		requestConditionRight string
		expectedIMEI          string
		err                   error
	}{
		{
			requestTableName:      "TestVin2IMEIMapping",
			requestAttrToGet:      "imei",
			requestConditionLeft:  "vin",
			requestConditionRight: "JTDKB3FU203501579",
			expectedIMEI:          "353816051214205",
			err:                   nil,
		},
	}

	for _, test := range tests {
		queryOutput, errRecLookup := RecordLookup(test.requestTableName, test.requestAttrToGet, test.requestConditionLeft, test.requestConditionRight)
		assert.IsType(t, test.err, errRecLookup)

		for _, i := range queryOutput.Items {
			item := IMEIItem{}

			err := dynamodbattribute.UnmarshalMap(i, &item)
			if err != nil {
				return
			}
			imei = item.IMEI
		}
		assert.EqualValues(t, test.expectedIMEI, imei)
	}
}
