package s3bucket

import (
	"cdk.tf/go/stack/generated/hashicorp/aws/s3bucket"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func CreateS3Bucket(stack cdktf.TerraformStack, name string) {
	s3bucket := s3bucket.NewS3Bucket(stack, jsii.String("bucket"), &s3bucket.S3BucketConfig{
		Bucket: jsii.String(name),
	})

	cdktf.NewTerraformOutput(stack, jsii.String("Bucket name"), &cdktf.TerraformOutputConfig{
		Value: s3bucket.Id(),
	})
}
