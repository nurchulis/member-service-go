package libs

import (
	"bytes"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// UploadFile uploads a file to an S3 bucket
func UploadFile(bucket, key string, contents []byte) error {

	sess, err := session.NewSession(&aws.Config{
		Endpoint:    aws.String("https://is3.cloudhost.id"),
		Region:      aws.String("ap-southeast-3"),
		Credentials: credentials.NewStaticCredentials("N62BOTLXZ23X1TA0P35H", "c2H89J4eDYImF9OwAjdDQHQnDKwVwkKOmB5nKXOa", ""),
	})

	// Call the NewAwsS3 function and assign the return value to a variable
	send := s3.New(sess)
	// Create an S3 object and set its content type
	params := &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(contents),
		ACL:    aws.String("public-read"),
	}

	// Upload the file to S3
	_, err = send.PutObject(params)
	if err != nil {
		fmt.Println("failed to upload file to S3", err)
		return err
	}

	fmt.Println("file uploaded successfully")

	return nil
}
