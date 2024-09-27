package dynamodb

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func ListDynamoDBTables() ([]string, error) {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	// create the input configuration instance
	input := &dynamodb.ListTablesInput{}

	var tableNames []string
	for {
			// Get the list of tables
			result, err := svc.ListTables(input)
			if err != nil {
					if aerr, ok := err.(awserr.Error); ok {
							switch aerr.Code() {
							case dynamodb.ErrCodeInternalServerError:
									fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
							default:
									fmt.Println(aerr.Error())
							}
					} else {
							// Print the error, cast err to awserr.Error to get the Code and
							// Message from an error.
							log.Print(err)
							return nil, err
					}
					log.Print(err)
					return nil, err
			}

			for _, tableName := range result.TableNames {
				tableNames = append(tableNames, *tableName)
			}

			// assign the last read tablename as the start for our next call to the ListTables function
			// the maximum number of table names returned in a call is 100 (default), which requires us to make
			// multiple calls to the ListTables function to retrieve all table names
			input.ExclusiveStartTableName = result.LastEvaluatedTableName

			if result.LastEvaluatedTableName == nil {
					break
			}
	}

	return tableNames, nil
}