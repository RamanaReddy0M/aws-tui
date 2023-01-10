package main

import (
	"log"

	"github.com/aws-tui/cmd"
	"github.com/aws-tui/internal/config"
	"github.com/aws-tui/internal/s3"
)

func main() {
	cmd.Execute()
	cfg, err := config.Get()
	if err != nil {
		log.Fatal(err)
	}
	// Create an Amazon S3 service client
	s3Service := s3.NewS3Service(cfg)
	s3Service.ListBuckets()
}
