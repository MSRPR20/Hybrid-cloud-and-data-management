package config

import (
    "os"
)

type Config struct {
    BackupSource      string
    BackupDestination string
    BackupSchedule    string
    SQLitePath        string
    AWSRegion         string
    AWSAccessKey      string
    AWSSecretKey      string
    DynamoDBTable     string
    S3Bucket          string
}

func LoadConfig() *Config {
    return &Config{
        BackupSource:      getEnv("BACKUP_SOURCE", "/path/to/source"),
        BackupDestination: getEnv("BACKUP_DESTINATION", "/path/to/backup.zip"),
        BackupSchedule:    getEnv("BACKUP_SCHEDULE", "0 0 * * *"),
        SQLitePath:        getEnv("SQLITE_PATH", "./local.db"),
        AWSRegion:         getEnv("AWS_REGION", "us-west-2"),
        AWSAccessKey:      getEnv("AWS_ACCESS_KEY", ""),
        AWSSecretKey:      getEnv("AWS_SECRET_KEY", ""),
        DynamoDBTable:     getEnv("DYNAMO_DB_TABLE", "MyTable"),
        S3Bucket:          getEnv("S3_BUCKET", "my-backup-bucket"),
    }
}

func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}
    