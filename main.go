package main

import (
	"cdk.tf/go/stack/ec2"
	provider "cdk.tf/go/stack/provider"
	s3 "cdk.tf/go/stack/s3"
	"cdk.tf/go/stack/vpc"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func Dev(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	cdktf.NewS3Backend(stack, &cdktf.S3BackendConfig{
		Bucket:        jsii.String("cdktf-go-backend"),
		Key:           jsii.String("dev"),
		Region:        jsii.String("us-west-1"),
		DynamodbTable: jsii.String("cdktf-go-lock"),
	})

	// Init AWS provider
	provider.Initprovider(stack, "us-west-1")

	// Create an Ec2 instance
	ec2.CreateEC2Instance(stack, "t3.micro", "ami-01456a894f71116f2")

	// Create an S3 bucket
	s3.CreateS3Bucket(stack, "hh-cdktf-go-bucket")

	// Create VPC
	vpc.CreateVPC(stack, "mydevVPC")

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	Dev(app, "dev")

	app.Synth()
}
