package core

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

/*
S3Images allows interaction with s3 buckets having images
*/
type S3Images struct {
	Bucket string
	Region string
	Prefix string
	Limit  int
}

/*
Get gets the string urls for s3 buckets.
The access is for 10 mins
*/
func (s3Images S3Images) Get() []string {
	// see https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html
	// for setting the env key
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(s3Images.Region),
		// Credentials: credentials.NewStaticCredentials("", "", "")
	},
	)
	bucket := aws.String(s3Images.Bucket)
	// Create S3 service client
	svc := s3.New(sess)
	var items []string
	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: bucket, Prefix: aws.String(s3Images.Prefix)})
	if err != nil {
		exitErrorf("Unable to list items in bucket %q, %v", bucket, err)
	}
	for _, item := range resp.Contents {
		if *item.Size == 0 {
			continue
		}
		req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
			Bucket: bucket,
			Key:    aws.String(*item.Key),
		})
		urlStr, err := req.Presign(10 * time.Minute)

		if err != nil {
			fmt.Println("Failed to sign request", err)
		}
		items = append(items, urlStr)
	}
	return items
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
