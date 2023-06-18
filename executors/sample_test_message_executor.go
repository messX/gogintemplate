package executors

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/messx/gogintemplate/helpers"
	"github.com/messx/gogintemplate/infra/logger"
)

type SampleTestMessageExecutor struct {
	QueueUrl            string
	DefaultMessageLimit int
	sqsHelper           helpers.SqsHelper
}

func (sampleMsgExecutor *SampleTestMessageExecutor) Init(queueUrl string, defaultLimit int, session *session.Session) {
	sampleMsgExecutor.QueueUrl = queueUrl
	sampleMsgExecutor.DefaultMessageLimit = defaultLimit
	sampleMsgExecutor.sqsHelper = helpers.SqsHelper{
		QueueUrl: sampleMsgExecutor.QueueUrl,
		Session:  session,
	}
}

/*
This process should be singleton
*/
func (sampleMsgExecutor *SampleTestMessageExecutor) Process() error {
	chnMessages := make(chan *sqs.Message)
	logger.Infof("Reading messages from queue %s", sampleMsgExecutor.QueueUrl)
	go sampleMsgExecutor.sqsHelper.ReceiveMessageBulkAsync(chnMessages, sampleMsgExecutor.DefaultMessageLimit)
	for message := range chnMessages {
		err := sampleMsgExecutor.handle(message)
		if err != nil {
			logger.Errorf("Unable to process message", err)
			return err
		}
		delErr := sampleMsgExecutor.delete(message)
		if delErr != nil {
			logger.Errorf("Unable to del message", delErr)
			return delErr
		}
	}
	return nil
}

func (sampleMsgExecutor *SampleTestMessageExecutor) handle(message *sqs.Message) error {
	logger.Debugf("Handling message")
	logger.Debugf("Message body %v", *message.Body)
	fmt.Println(*message.ReceiptHandle)
	return nil
}

func (sampleMsgExecutor *SampleTestMessageExecutor) delete(message *sqs.Message) error {
	return sampleMsgExecutor.sqsHelper.DeleteMessage(message)
}
