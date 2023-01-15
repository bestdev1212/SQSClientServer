package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/joho/godotenv"
)

func SendMessage(sess *session.Session, queueUrl string, messageBody string) error {
	sqsClient := sqs.New(sess)

	_, err := sqsClient.SendMessage(&sqs.SendMessageInput{
		QueueUrl:    &queueUrl,
		MessageBody: aws.String(messageBody),
	})

	return err
}

func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}

func main() {

	err := godotenv.Load("./.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(1)
	}
	awsAccessKeyID := GetEnvWithKey("AWS_ACCESS_KEY_ID")
	awsRegion := GetEnvWithKey("AWS_REGION")
	awsSecretKey := GetEnvWithKey("AWS_SECRET_ACCESS_KEY")
	sqlUrl := GetEnvWithKey("SQS_URL")
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(awsRegion),
			Credentials: credentials.NewStaticCredentials(
				awsAccessKeyID,
				awsSecretKey,
				"", // a token will be created when the session it's used.
			),
		})
	if err != nil {
		panic(err)
	}

	

	messageBody := "This is a test message"
	err = SendMessage(sess, sqlUrl, messageBody)
	if err != nil {
		fmt.Printf("Got an error while trying to send message to queue: %v", err)
		return
	}

	fmt.Println("Message sent successfully")
}