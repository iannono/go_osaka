package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"what is your name"`
	Age  int    `json:"how old are you"`
}

type MyResponse struct {
	Message string `json:"Answer:"`
}

func osakaHandler(event MyEvent) (MyResponse, error) {
	return MyResponse{Message: fmt.Sprintf("%s is %d old!", event.Name, event.Age)}, nil
}

//How to build go osaka
//GOOS=linux GOARCH=amd64 go build -o main main.go
//zip main.zip main
func main() {
	lambda.Start(osakaHandler)
}
