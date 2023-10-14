// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX - License - Identifier: Apache - 2.0
// snippet-start:[ec2.go-v2.CreateInstance]
package ec2

import (
	"context"
	"fmt"
	"mci/common/helper"
	uaws "mci/common/util/aws"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/urfave/cli/v2"
)

func CreateEc2Instance(cCtx *cli.Context) error {
	name := "testinstance"
	value := "value"
	ec2model := uaws.Init()

	// Create separate values if required.
	minMaxCount := int32(1)

	input := &ec2.RunInstancesInput{
		ImageId:      aws.String("ami-e7527ed7"),
		InstanceType: types.InstanceTypeT2Micro,
		MinCount:     &minMaxCount,
		MaxCount:     &minMaxCount,
	}

	result, err := uaws.MakeInstance(context.TODO(), ec2model.Ec2Client(), input)
	helper.CheckError(err, "Cant Create Instance", true)

	tagInput := &ec2.CreateTagsInput{
		Resources: []string{*result.Instances[0].InstanceId},
		Tags: []types.Tag{
			{
				Key:   &name,
				Value: &value,
			},
		},
	}
	_, err = uaws.MakeTags(context.TODO(), ec2model.Ec2Client(), tagInput)
	helper.CheckError(err, "Got an error tagging the instance:")
	fmt.Println("Created tagged instance with ID " + *result.Instances[0].InstanceId)
	return nil
}

// snippet-end:[ec2.go-v2.CreateInstance]
