package s3

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsS3 "github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Service interface {
	ListBuckets() error
}

type s3Service struct {
	client awsS3.Client
}

func NewS3Service(config aws.Config) S3Service {
	// Create an Amazon S3 service client
	return s3Service{client: *awsS3.NewFromConfig(config)}
}

func (s s3Service) ListBuckets() error {

	output, err := s.client.ListBuckets(context.Background(), &awsS3.ListBucketsInput{})
	if err != nil {
		return err
	}

	for i, b := range output.Buckets {
		fmt.Printf("%v: %v %v\n", i+1, *b.Name, b.CreationDate.Format("2006-01-02 15:04:05"))
	}

	// Get the first page of results for ListObjectsV2 for a bucket
	// output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
	// 	Bucket: aws.String("practo-images"),
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println("first page results:")
	// for _, object := range output.Contents {
	// 	log.Printf("key=%s size=%d", aws.ToString(object.Key), object.Size)
	// }
	return nil
}
