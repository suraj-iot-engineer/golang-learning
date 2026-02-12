package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	bucketName := os.Getenv("BUCKET_NAME")
	if bucketName == "" {
		fmt.Println("⚠️  BUCKET_NAME env var not set. Skipping real S3 call.")
		fmt.Println("Usage: set BUCKET_NAME=my-bucket && go run main.go")
		return
	}

	// 1. Load AWS Config (credentials from env/file)
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// 2. Create S3 Client
	client := s3.NewFromConfig(cfg)

	// 3. List Objects
	output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		log.Fatalf("failed to list objects, %v", err)
	}

	fmt.Printf("📦 Objects in %s:\n", bucketName)
	for _, object := range output.Contents {
		fmt.Printf(" - %s (%d bytes)\n", *object.Key, object.Size)
	}
}
