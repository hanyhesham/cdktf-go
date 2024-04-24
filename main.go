package main

import (
	ec2 "cdk.tf/go/stack/ec2"
	s3 "cdk.tf/go/stack/s3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	awsprovider "github.com/cdktf/cdktf-provider-aws-go/aws/v19/provider"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	awsprovider.NewAwsProvider(stack, jsii.String("AWS"), &awsprovider.AwsProviderConfig{
		Region: jsii.String("us-west-1"),
	})

	// Create an Ec2 instance from separate file
	ec2.CreateEC2Instance(stack, "t2-micro")

	// Create an S3 bucket from separate file
	s3.CreateS3Bucket(stack, "my-unique-bucket-name")

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	NewMyStack(app, "cdktf-go")

	app.Synth()
}
