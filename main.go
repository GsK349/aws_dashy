package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"os"
)

var (
	bucketname      string
	objectPrefix    string
	objectDelimiter string
	maxKeys         int
)

func main() {
	sess, err := session.NewSession(&aws.Config{})
	if err != nil {
		fmt.Println("Error creating a session: %v", err)
		return
	}
}
