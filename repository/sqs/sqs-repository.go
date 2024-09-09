package sqs

import (
	"context"
	"fmt"
	"log"

	"github.com/ThailanTec/go-sqs-sns/domain"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

func (r *SQSRepository) SendMessage(message domain.Message) error {
	input := &sqs.SendMessageInput{
		QueueUrl:    &r.queueURL,
		MessageBody: aws.String(message.Body),
	}

	_, err := r.client.SendMessage(context.TODO(), input)
	if err != nil {
		log.Printf("failed to send message: %v", err)
		return err
	}

	log.Println("Message sent successfully")
	return nil
}

func (r *SQSRepository) ReceiveMessages() ([]domain.Message, error) {
	input := &sqs.ReceiveMessageInput{
		QueueUrl:            &r.queueURL,
		MaxNumberOfMessages: 10, // Numero máximo de mensagens. Alterar depois THAILAN!!!
		WaitTimeSeconds:     5,  // Tempo de espera para long polling
		VisibilityTimeout:   10,
	}

	result, err := r.client.ReceiveMessage(context.TODO(), input)
	if err != nil {
		log.Printf("failed to receive messages: %v", err)
		return nil, err
	}

	var messages []domain.Message
	for _, m := range result.Messages {
		messages = append(messages, domain.Message{Body: *m.Body})
	}

	return messages, nil
}

func (actor *SQSRepository) DeleteMessages(queueUrl string, messages []domain.Message) error {
	entries := make([]types.DeleteMessageBatchRequestEntry, len(messages))
	for msgIndex, msg := range messages {
		entries[msgIndex].Id = aws.String(fmt.Sprintf("%v", msgIndex))
		entries[msgIndex].ReceiptHandle = aws.String(fmt.Sprintf("%v", msg))
	}

	_, err := actor.client.DeleteMessageBatch(context.TODO(), &sqs.DeleteMessageBatchInput{
		Entries:  entries,
		QueueUrl: aws.String(queueUrl),
	})

	if err != nil {
		log.Printf("Couldn't delete messages from queue %v. Here's why: %v\n", queueUrl, err)
	}
	return err
}
