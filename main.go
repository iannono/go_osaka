package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// type MyEvent struct {
// 	Name string `json:"what is your name"`
// 	Age  int    `json:"how old are you"`
// }

// type MyResponse struct {
// 	Message string `json:"Result:"`
// }

func osakaHandler(ctx context.Context, event events.S3Event) {
	for _, record := range event.Records {
		s3 := record.S3
		fmt.Printf("[%s - %s] Bucket = %s, Key = %s \n", record.EventSource, record.EventTime, s3.Bucket.Name, s3.Object.Key)
	}
}

//How to build go osaka
//GOOS=linux GOARCH=amd64 go build -o main main.go
//zip main.zip main
func main() {
	lambda.Start(osakaHandler)
}
