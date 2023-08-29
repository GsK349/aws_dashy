package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String("ap-south-1"),
		},
	)
	if err != nil {
		fmt.Println("Error creating a session: %v", err)
		return
	}

	//Creating a s3 client
	s3Client := s3.New(sess)

	//getting all the buckets in the account
	results, err := s3Client.ListBuckets(nil)
	if err != nil {
		fmt.Println("Error Listing Buckets: %v", err)
		return
	}

	//Listing buckets
	fmt.Println("S3 buckets:")
	for _, bucket := range results.Buckets {
		fmt.Println(*bucket.Name)
	}

}
