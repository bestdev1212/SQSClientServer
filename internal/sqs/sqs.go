package sqs

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/topTalent1212/SQSClientServer/internal/env"
)

var (
	accessKeyID string;
	region string;
	secretKey string;
	url string;
	sess *session.Session;
	err error;
)

func init() {
	accessKeyID = env.GetEnvWithKey("AWS_ACCESS_KEY_ID")
	region = env.GetEnvWithKey("AWS_REGION")
	secretKey = env.GetEnvWithKey("AWS_SECRET_ACCESS_KEY")
	url = env.GetEnvWithKey("SQS_URL")
	sess, err = session.NewSession(
		&aws.Config{
			Region: aws.String(region),
			Credentials: credentials.NewStaticCredentials(
				accessKeyID,
				secretKey,
				"", // a token will be created when the session it's used.
			),
		})
	if err != nil {
		panic(err)
	}
}

func SendMessage(messageBody string) error {
	sqsClient := sqs.New(sess)

	_, err := sqsClient.SendMessage(&sqs.SendMessageInput{
		QueueUrl:    &url,
		MessageBody: aws.String(messageBody),
	})

	if(err == nil){
		fmt.Println("Message was sent successfuly")
	}

	return err
}