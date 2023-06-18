package config

import (
	"github.com/messx/gogintemplate/executors"
	"github.com/messx/gogintemplate/infra/logger"
	"github.com/spf13/viper"
)

type SqsWatcher struct {
	/* this obj opens all channels to read from multiple sqs queue */
	Handlers map[string]executors.MessageExecutorInterface
}

func (sqsWatcher *SqsWatcher) Init() {
	awsConfig := new(AwsConfig)
	sqsWatcher.Handlers = make(map[string]executors.MessageExecutorInterface)
	sqsWatcher.Handlers["TEST_SQS_HANDLER"] = new(executors.SampleTestMessageExecutor)
	sqsWatcher.Handlers["TEST_SQS_HANDLER"].Init(viper.GetString("TEST_SQS_QUEUE_URL"), 10, awsConfig.GetSession())
	logger.Debugf("Successfully initialised SQS handlers")
}

func (sqsWatcher *SqsWatcher) Process() {
	logger.Debugf("Process sqs messages using handlers")
	for name, handler := range sqsWatcher.Handlers {
		err := handler.Process()
		if err != nil {
			logger.Errorf("Unable to process %s", name)
		}
	}
}

func InitAndProcessSQS() {
	sqsWatcher := new(SqsWatcher)
	sqsWatcher.Init()
	sqsWatcher.Process()
}
