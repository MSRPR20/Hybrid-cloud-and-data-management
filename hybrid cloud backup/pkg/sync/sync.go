package sync

import (
    "database/sql"
    "fmt"
    "hybrid-data-management-backup/pkg/database"
    "hybrid-data-management-backup/pkg/cloud"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/aws"
)

func SyncData(db *sql.DB, dynamoSvc *dynamodb.DynamoDB, tableName string) error {
    users, err := database.GetUsers(db)
    if err != nil {
        return err
    }

    for _, user := range users {
        item := map[string]*dynamodb.AttributeValue{
            "id": {
                N: aws.String(fmt.Sprintf("%d", user["id"])),
            },
            "name": {
                S: aws.String(user["name"].(string)),
            },
            "email": {
                S: aws.String(user["email"].(string)),
            },
        }

        err := cloud.PutItem(dynamoSvc, tableName, item)
        if err != nil {
            return err
        }
    }

    return nil
}
