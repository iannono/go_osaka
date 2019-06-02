package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// type MyEvent struct {
// 	Name string `json:"what is your name"`
// 	Age  int    `json:"how old are you"`
// }

// type MyResponse struct {
// 	Message string `json:"Result:"`
// }

func sendEmail(fileName string) {
	from := mail.NewEmail("Send Grid Example", "test@example.com")
	subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail("Example User", "ian@feedmob.com")
	plainTextContent := "Easy come, easy go" + fileName
	htmlContent := "<strong>Easy come, easy go" + fileName + "</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(fileName)
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

func osakaHandler(ctx context.Context, event events.S3Event) {
	for _, record := range event.Records {
		s3 := record.S3
		fmt.Printf("[%s - %s] Bucket = %s, Key = %s \n", record.EventSource, record.EventTime, s3.Bucket.Name, s3.Object.Key)
		fileName := s3.Object.Key
		sendEmail(fileName)
	}
}

//How to build go osaka
//GOOS=linux GOARCH=amd64 go build -o main main.go
//zip main.zip main
func main() {
	lambda.Start(osakaHandler)
}
