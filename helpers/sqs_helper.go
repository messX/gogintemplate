package helpers

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/messx/gogintemplate/infra/logger"
)

type SqsHelper struct {
	QueueUrl string
	Session  *session.Session
}

func (sqsHelper *SqsHelper) SendMessage(messageBody string) error {
	sqsClient := sqs.New(sqsHelper.Session)
	_, err := sqsClient.SendMessage(&sqs.SendMessageInput{
		QueueUrl:    &sqsHelper.QueueUrl,
		MessageBody: aws.String(messageBody),
	})

	return err
}

func (sqsHelper *SqsHelper) ReadMessage(maxMessages int) (*sqs.ReceiveMessageOutput, error) {
	sqsClient := sqs.New(sqsHelper.Session)
	if maxMessages < 0 {
		maxMessages = 10
	}
	msgResult, err := sqsClient.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:            &sqsHelper.QueueUrl,
		MaxNumberOfMessages: aws.Int64(int64(maxMessages)),
	})

	if err != nil {
		return nil, err
	}

	return msgResult, nil
}

func (sqsHelper *SqsHelper) ReceiveMessageBulkAsync(chn chan<- *sqs.Message, maxMessages int) {
	sqsClient := sqs.New(sqsHelper.Session)
	if maxMessages < 0 {
		maxMessages = 10
	}

	for {
		msgResults, err := sqsClient.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl:            &sqsHelper.QueueUrl,
			MaxNumberOfMessages: aws.Int64(int64(maxMessages)),
		})
		if err != nil {
			logger.Errorf("Unable to read message from queue %s", sqsHelper.QueueUrl)
		}
		for _, message := range msgResults.Messages {
			chn <- message
		}
	}

}

func (sqsHelper *SqsHelper) DeleteMessage(message *sqs.Message) error {
	sqsClient := sqs.New(sqsHelper.Session)
	_, err := sqsClient.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      &sqsHelper.QueueUrl,
		ReceiptHandle: message.ReceiptHandle,
	})
	return err
}
