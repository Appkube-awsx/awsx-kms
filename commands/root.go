package commands

import (
	"log"

	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/Appkube-awsx/awsx-kms/commands/kmscmd"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/spf13/cobra"
)

// AwsxKmsCmd represents the base command when called without any subcommands
var AwsxKmsCmd = &cobra.Command{
	Use:   "kms",
	Short: "get kms Details command gets resource counts",
	Long:  `get kms Details command gets resource counts details of an AWS account`,

	Run: func(cmd *cobra.Command, args []string) {

		log.Println("Command kms started")

		authFlag, clientAuth, err := authenticate.CommandAuth(cmd)
		if err != nil {
			cmd.Help()
			return
		}
		if authFlag {
			ListKeys(*clientAuth)
		} else {
			cmd.Help()
			return
		}
	},
}

func ListKeys(auth client.Auth) (*kms.ListKeysOutput,error) {

	log.Println("Getting kms key list summary")

	// this is Api auth and compulsory for every controller
	kmsClient := client.GetClient(auth, client.KMS_CLIENT).(*kms.KMS)

	input := &kms.ListKeysInput{}

	keyList, err := kmsClient.ListKeys(input)

	if err != nil {
		log.Fatalln("Error: in getting kms list", err)
	}

	log.Println(keyList)

	return keyList, err
}

// Execute runs the command
// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := AwsxKmsCmd.Execute()

	if err != nil {
		log.Fatal("There was some error while executing the CLI: ", err)
		return
	}
}

func init() {
	AwsxKmsCmd.AddCommand(kmscmd.GetConfigDataCmd)

	// Define persistent flags for the command

	AwsxKmsCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxKmsCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxKmsCmd.PersistentFlags().String("accountId", "", "aws account number")
	AwsxKmsCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxKmsCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxKmsCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxKmsCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxKmsCmd.PersistentFlags().String("externalId", "", "aws external id auth")
}
