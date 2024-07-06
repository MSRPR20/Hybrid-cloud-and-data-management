package backup

import (
    "archive/zip"
    "io/ioutil"
    "os"
    "path/filepath"
)

func PerformBackup(sourceDir, destinationFile string) error {
    zipFile, err := os.Create(destinationFile)
    if err != nil {
        return err
    }
    defer zipFile.Close()

    archive := zip.NewWriter(zipFile)
    defer archive.Close()

    err = filepath.Walk(sourceDir, func(filePath string, fileInfo os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if fileInfo.IsDir() {
            return nil
        }

        fileBytes, err := ioutil.ReadFile(filePath)
        if err != nil {
            return err
        }

        zipFile, err := archive.Create(filePath)
        if err != nil {
            return err
        }

        _, err = zipFile.Write(fileBytes)
        return err
    })

    return err
}
