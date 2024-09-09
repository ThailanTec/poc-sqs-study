package domain

type Message struct {
	Body string
}

type SQSRepository interface {
	SendMessage(message Message) error
	ReceiveMessages() ([]Message, error)
	DeleteMessages(queueUrl string, messages []Message) error
}
