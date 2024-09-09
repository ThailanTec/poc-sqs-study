package usecases

import (
	"github.com/ThailanTec/go-sqs-sns/domain"
)

type SendMessageUseCase struct {
	SQSRepository domain.SQSRepository
}

func NewSendMessageUseCase(SQSRepository domain.SQSRepository) *SendMessageUseCase {
	return &SendMessageUseCase{
		SQSRepository: SQSRepository,
	}
}

func (uc *SendMessageUseCase) SendMessage(body string) error {
	message := domain.Message{Body: body}
	return uc.SQSRepository.SendMessage(message)
}

func (uc *SendMessageUseCase) Recive() ([]domain.Message, error) {
	return uc.SQSRepository.ReceiveMessages()
}

func (uc *SendMessageUseCase) DeleteAll(queueUrl string, messages []domain.Message) error {
	return uc.SQSRepository.DeleteMessages(queueUrl, messages)
}
