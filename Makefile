.PHONY: run create-queue create-dlq

# Variáveis
QUEUE_NAME=email
DLQ_NAME=email-dlq
AWS_ENDPOINT=http://localhost:4566
REGION=us-east-1

# all
all-queue: create-queue create-dlq run

# Iniciar a aplicação
run:
	@echo "Iniciando a aplicação..."
	go run cmd/main.go


# Criar a fila principal
create-queue:
	@echo "Criando a fila $(QUEUE_NAME)..."
	aws --endpoint-url=$(AWS_ENDPOINT) sqs create-queue \
		--queue-name $(QUEUE_NAME) \
		--region $(REGION)

# Criar a fila DLQ e associá-la à fila principal
create-dlq:
	@echo "Criando a DLQ $(DLQ_NAME)..."
	aws --endpoint-url=$(AWS_ENDPOINT) sqs create-queue \
		--queue-name $(DLQ_NAME) \
		--region $(REGION)

	@echo "Criando a fila principal com a DLQ associada..."
	aws --endpoint-url=$(AWS_ENDPOINT) sqs create-queue \
		--queue-name $(QUEUE_NAME) \
		--attributes '{"RedrivePolicy":"{\"deadLetterTargetArn\":\"arn:aws:sqs:us-east-1:000000000000:$(DLQ_NAME)\",\"maxReceiveCount\":\"5\"}"}' \
		--region $(REGION)
