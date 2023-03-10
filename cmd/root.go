package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   `awstui`,
	Short: `awstui visulize contents of services EC2, S3, VPC, Disks, RDS..`,
	Long:  `awstui visulize contents of services EC2, S3, VPC, Disks, RDS`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}