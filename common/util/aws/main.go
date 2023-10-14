package aws

import (
	"context"
	"mci/common/helper"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func Init() (m Model) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	helper.CheckError(err, "Load Default Config error")
	m = Model{config: cfg}

	setEc2Client(&m)

	return
}

func setEc2Client(m *Model) {
	client := ec2.NewFromConfig(m.config)
	*m = Model{ec2Client: client}
}

func (m Model) DefaultConfig() (cfg aws.Config) {
	return m.config
}
func (m Model) Ec2Client() *ec2.Client {
	return m.ec2Client
}
