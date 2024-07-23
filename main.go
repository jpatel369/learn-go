package main

import (
    "context"
    "fmt"
    "log"

    "github.com/minio/minio-go/v7"
    "github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
    // Replace these values with your MinIO server details
    endpoint := "your-minio-endpoint"
    accessKeyID := "your-access-key"
    secretAccessKey := "your-secret-key"
    useSSL := false // Change to true if your MinIO server uses SSL

    // Initialize minio client object.
    minioClient, err := minio.New(endpoint, &minio.Options{
        Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
        Secure: useSSL,
    })
    if err != nil {
        log.Fatalln(err)
    }

    // List all buckets.
    buckets, err := minioClient.ListBuckets(context.Background())
    if err != nil {
        log.Fatalln(err)
    }
    fmt.Println("Buckets:")
    for _, bucket := range buckets {
        fmt.Println(bucket.Name)
    }

    // Create a new bucket.
    newBucketName := "your-new-bucket"
    err = minioClient.MakeBucket(context.Background(), newBucketName, minio.MakeBucketOptions{Region: "us-east-1"})
    if err != nil {
        // Check to see if the bucket already exists.
        exists, errBucketExists := minioClient.BucketExists(context.Background(), newBucketName)
        if errBucketExists == nil && exists {
            fmt.Printf("Bucket %s already exists\n", newBucketName)
        } else {
            log.Fatalln(err)
        }
    } else {
        fmt.Printf("Successfully created %s\n", newBucketName)
    }
}
