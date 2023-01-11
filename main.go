package main

import (
	"fmt"
	"log"

	"github.com/aws-tui/cmd"
	"github.com/aws-tui/internal/config"
	"github.com/aws-tui/internal/ec2"
)

func main() {
	cmd.Execute()

	//Remove section
	cfg, err := config.Get()
	if err != nil {
		log.Fatal(err)
	}

	sess, err := config.GetSession("test", "ap-south-1", cfg.AwsConfig)
	if err != nil {
		fmt.Println("Error creating session:", err)
		return
	}

	ec2Service := ec2.NewEc2Service(*sess)
	ec2Service.GetInstances()

	// fmt.Println("config: ", cfg)
	// // Create an Amazon S3 service client
	// s3Service := s3.NewS3Service(cfg)
	// s3Service.ListBuckets()

	// sess, err := session.NewSession(&aws.Config{
	// 	Region:   aws.String("ap-south-1"),
	// 	Endpoint: aws.String("http://localhost:4566"),
	// })

	// sess, err := session.NewSessionWithOptions(session.Options{Config: aws.Config{
	// 	Region: aws.String("ap-south-1"),
	// 	//Endpoint: aws.String("http://localhost:4566"),
	// }, Profile: "dev"})
	// if err != nil {
	// 	fmt.Println("Error creating session:", err)
	// 	return
	// }

	// Create a new RDS client
	// Define the EC2 instance to launch
	// runResult, err := svc.RunInstances(&ec2.RunInstancesInput{
	// 	ImageId:      aws.String("1234"),
	// 	InstanceType: aws.String("t2.nano"),
	// 	MinCount:     aws.Int64(1),
	// 	MaxCount:     aws.Int64(1),
	// })

	// if err != nil {
	// 	fmt.Println("Could not create instance", err)
	// 	return
	// }

	// fmt.Println("Successfully created EC2 instance")
	// fmt.Println("Instance ID: ", *runResult.Instances[0].InstanceId)
}
