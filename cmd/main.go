package main

import (
	"fmt"
	"log"

	"github.com/ThailanTec/go-sqs-sns/repository/sqs"
	usecases "github.com/ThailanTec/go-sqs-sns/src/usecase"
)

func main() {
	queueURL := "http://sqs.us-east-1.localhost.localstack.cloud:4566/000000000000/email"
	if queueURL == "" {
		log.Fatal("SQS_QUEUE_URL is not set")
	}

	sqsRepo, err := sqs.NewSQSRepository(queueURL)
	if err != nil {
		log.Fatalf("Failed to create SQS repository: %v", err)
	}

	sendMessageUseCase := usecases.NewSendMessageUseCase(sqsRepo)
	for i := 0; i < 10; i++ {
		/* err = sendMessageUseCase.SendMessage("thailandev09@gmail.com")
		if err != nil {
			log.Fatalf("Failed to send message: %v", err)
		} */
	}

	messages, err := sendMessageUseCase.Recive()
	if err != nil {
		log.Fatalf("Failed to receive messages: %v", err)
	}

	// Imprimir mensagens recebidas.
	for _, msg := range messages {
		fmt.Printf("Received message: %s\n", msg.Body)
	}
}
