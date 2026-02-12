# 11. Cloud Native Go

Go is the language of the cloud. This module covers authenticating with Cloud Providers (AWS/GCP) and using their SDKs.

## 📚 Topics Covered

1.  **AWS SDK for Go v2**: Connecting to S3, DynamoDB.
2.  **Google Cloud Client Libraries**: Connecting to Cloud Storage, Pub/Sub.
3.  **12-Factor App**: Config, Backing Services, Disposability.
4.  **Serverless**: Writing AWS Lambda functions in Go.

## 🛠️ Prerequisites

- valid AWS/GCP credentials configured locally (`~/.aws/credentials` or `GOOGLE_APPLICATION_CREDENTIALS`).

## 🚀 Example
The example demonstrates how to list data from an S3 bucket using the AWS SDK v2.

```bash
# Set bucket name via env var
export BUCKET_NAME=my-test-bucket
go run main.go
```
