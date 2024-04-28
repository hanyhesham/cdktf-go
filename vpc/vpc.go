package vpc

import (
	"cdk.tf/go/stack/generated/hashicorp/aws/vpc"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func CreateVPC(stack cdktf.TerraformStack, name string) {
	vpc := vpc.NewVpc(stack, jsii.String("vpc"), &vpc.VpcConfig{
		Tags: &map[string]*string{
			"Name": jsii.String(name),
		},
		CidrBlock: jsii.String("10.0.0.0/24"),
	})

	cdktf.NewTerraformOutput(stack, jsii.String("VPC ARN"), &cdktf.TerraformOutputConfig{
		Value: vpc.Arn(),
	})
}
