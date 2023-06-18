package executors

import "github.com/aws/aws-sdk-go/aws/session"

type MessageExecutorInterface interface {
	Init(queueUrl string, defaultLimit int, session *session.Session)
	Process() error
}
