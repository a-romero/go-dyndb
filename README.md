# go-dyndb

Package records has got the following functions:
## Generic record writer
The WriteRecord function can write any type of record with any structure (not just the ones defined above) on a given DynamoDB table. This is achieved by the use of interfaces{}.
```
// WriteRecord ...
func WriteRecord(object interface{}, tableName string, svc *dynamodb.DynamoDB) error
```
## Create DynamoDB Service
Returns a DynamoDB service ready to be used with DynamoDB.
```
// CreateDynDBSvc ...
func CreateDynDBSvc() *dynamodb.DynamoDB
```
## Record Lookup
Provides with an easy way of doing a lookup on any given table.
```
// RecordLookup ...
func RecordLookup(table string, attrToGet string, conditionLeft string, conditionRight string) (*dynamodb.QueryOutput, error)
```
# go-dyndb
