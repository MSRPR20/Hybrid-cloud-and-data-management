package cloud

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/asws-sdk-go/service/dynamodb"
)

func ConnectDynamoDB(region, accessKey, secretKey string) (*dynamodb.DynamoDB, error) {
    sess, err := session.NewSession(&aws.Config{
        Region:      aws.String(region),
        Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
    })
    if err != nil {
        return nil, err
    }

    svc := dynamodb.New(sess)
    return svc, nil
}

func PutItem(svc *dynamodb.DynamoDB, table string, item map[string]*dynamodb.AttributeValue) error {
    input := &dynamodb.PutItemInput{
        Item:      item,
        TableName: aws.String(table),
    }

    _, err := svc.PutItem(input)
    return err
}
