package ec2

import (
	"cdk.tf/go/stack/generated/hashicorp/aws/instance"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func CreateEC2Instance(stack cdktf.TerraformStack, name string) {
	instance := instance.NewInstance(stack, jsii.String("compute"), &instance.InstanceConfig{
		Ami:          jsii.String("ami-01456a894f71116f2"),
		InstanceType: jsii.String(name),
	})

	cdktf.NewTerraformOutput(stack, jsii.String("EC2 public_ip"), &cdktf.TerraformOutputConfig{
		Value: instance.PublicIp(),
	})
}
