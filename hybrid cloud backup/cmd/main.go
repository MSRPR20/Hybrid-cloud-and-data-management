package main

import (
    "log"
    "hybrid-data-management-backup/config"
    "hybrid-data-management-backup/pkg/backup"
    "hybrid-data-management-backup/pkg/cloud"
    "hybrid-data-management-backup/pkg/database"
    "hybrid-data-management-backup/pkg/scheduler"
    "hybrid-data-management-backup/pkg/sync"
)

func main() {
    // Load configuration
    cfg := config.LoadConfig()

    // Connect to local SQLite database
    db, err := database.ConnectSQLite(cfg.SQLitePath)
    if err != nil {
        log.Fatal("Failed to connect to SQLite:", err)
    }
    defer db.Close()

    // Connect to AWS services
    dynamoSvc, err := cloud.ConnectDynamoDB(cfg.AWSRegion, cfg.AWSAccessKey, cfg.AWSSecretKey)
    if err != nil {
        log.Fatal("Failed to connect to DynamoDB:", err)
    }
    s3Svc, err := cloud.ConnectS3(cfg.AWSRegion, cfg.AWSAccessKey, cfg.AWSSecretKey)
    if err != nil {
        log.Fatal("Failed to connect to S3:", err)
    }

    // Schedule backup and synchronization
    scheduler.Schedule(cfg.BackupSchedule, func() {
        // Perform data synchronization
        err := sync.SyncData(db, dynamoSvc, cfg.DynamoDBTable)
        if err != nil {
            log.Println("Data synchronization failed:", err)
        } else {
            log.Println("Data synchronization completed successfully.")
        }

        // Perform backup
        err = backup.PerformBackup(cfg.BackupSource, cfg.BackupDestination)
        if err != nil {
            log.Println("Backup failed:", err)
        } else {
            log.Println("Backup completed successfully.")
            // Upload to cloud
            err := cloud.UploadToS3(s3Svc, cfg.BackupDestination, cfg.S3Bucket)
            if err != nil {
                log.Println("Upload to cloud failed:", err)
            } else {
                log.Println("Upload to cloud completed successfully.")
            }
        }
    })
}

    