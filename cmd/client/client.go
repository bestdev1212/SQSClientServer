package main

import (
	"github.com/topTalent1212/SQSClientServer/internal/sqs"
)

func main() {
	sqs.SendMessage("hello")
}