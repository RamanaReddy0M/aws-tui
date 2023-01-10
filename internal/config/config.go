package config

import (
	"context"
	"reflect"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
)

var awsCfg aws.Config

func Get() (aws.Config, error) {
	emptyCfg := aws.Config{}
	if reflect.DeepEqual(emptyCfg, awsCfg) {
		// Load the Shared AWS Configuration (~/.aws/config)
		cfg, err := awsConfig.LoadDefaultConfig(context.TODO())
		if err != nil {
			return aws.Config{}, err
		}
		awsCfg = cfg
	}
	return awsCfg, nil
}
