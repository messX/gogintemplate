package tests

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/messx/gogintemplate/config"
	"github.com/messx/gogintemplate/helpers"
	"github.com/messx/gogintemplate/infra/logger"
)

type testSQS struct {
	queueUrl string
	session  *session.Session
}

func (testSQS) Initialise() testSQS {
	awsConf := config.AwsConfig{}
	testObj := testSQS{
		queueUrl: "https://sqs.ap-south-1.amazonaws.com/871308545741/testQueue",
		session:  awsConf.GetSession(),
	}
	return testObj
}

func TestWriteToSQS(t *testing.T) {
	mainObj := testSQS{}
	testObj := mainObj.Initialise()
	sqsHelper := helpers.SqsHelper{
		QueueUrl: testObj.queueUrl,
		Session:  testObj.session,
	}
	err := sqsHelper.SendMessage("test body")
	if err != nil {
		fmt.Println(err)
		t.Error("Unable to send message sent test failed")
	}
}

func TestReadFromSQS(t *testing.T) {
	mainObj := testSQS{}
	testObj := mainObj.Initialise()
	sqsHelper := helpers.SqsHelper{
		QueueUrl: testObj.queueUrl,
		Session:  testObj.session,
	}
	response, err := sqsHelper.ReadMessage(10)
	if err != nil {
		t.Error("Unable to send message sent test failed")
	}
	logger.Infof("Respose from queue %s", response.Messages)

}
