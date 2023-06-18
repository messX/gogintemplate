package config

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/spf13/viper"
)

type AwsConfig struct {
}

func (awsConf AwsConfig) GetSession() *session.Session {
	SetupConfig()
	log.Printf("Region %s", viper.GetString("AWS_REGION"))
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "default",
		Config: aws.Config{
			Region:      aws.String(viper.GetString("AWS_REGION")),
			Credentials: credentials.NewStaticCredentials(viper.GetString("AWS_ACCESS_KEY_ID"), viper.GetString("AWS_SECRET_ACCESS_KEY"), ""),
		},
	})

	if err != nil {
		log.Fatalf("Failed to initialize new session: %v", err)
	}

	return sess
}
