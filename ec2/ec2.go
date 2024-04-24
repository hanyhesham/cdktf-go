package ec2

import (
	"cdk.tf/go/stack/generated/hashicorp/aws/instance"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func CreateEC2Instance(stack cdktf.TerraformStack, instanceType string, AMI_ID string) {
	instance := instance.NewInstance(stack, jsii.String("compute"), &instance.InstanceConfig{
		Ami:          jsii.String(AMI_ID),
		InstanceType: jsii.String(instanceType),
	})

	cdktf.NewTerraformOutput(stack, jsii.String("EC2 public_ip"), &cdktf.TerraformOutputConfig{
		Value: instance.PublicIp(),
	})
}
