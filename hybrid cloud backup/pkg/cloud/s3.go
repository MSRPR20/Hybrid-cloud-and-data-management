   package cloud

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
    "os"
    "path/filepath"
)

func ConnectS3(region, accessKey, secretKey string) (*s3.S3, error) {
    sess, err := session.NewSession(&aws.Config{
        Region:      aws.String(region),
        Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
    })
    if err != nil {
        return nil, err
    }

    svc := s3.New(sess)
    return svc, nil
}

func UploadToS3(svc *s3.S3, filePath, bucket string) error {
    file, err := os.Open(filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    _, err = svc.PutObject(&s3.PutObjectInput{
        Bucket: aws.String(bucket),
        Key:    aws.String(filepath.Base(filePath)),
        Body:   file,
    })
    return err
}
