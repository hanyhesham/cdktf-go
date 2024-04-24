package main

import (
	ec2 "cdk.tf/go/stack/ec2"
	"cdk.tf/go/stack/provider"
	s3 "cdk.tf/go/stack/s3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func Dev(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	// Init AWS provider
	provider.Initprovider(stack, "us-west-1")

	// Create an Ec2 instance
	ec2.CreateEC2Instance(stack, "t2-micro", "ami-01456a894f71116f2")

	// Create an S3 bucket
	s3.CreateS3Bucket(stack, "my-unique-bucket-name")

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	Dev(app, "dev")

	app.Synth()
}
