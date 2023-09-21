/*
Copyright © 2023 Afreen khan <afreen.khan@synectiks.com>
*/

package kmscmd

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/spf13/cobra"
)

// GetConfigDataCmd represents the getConfigData command
var GetConfigDataCmd = &cobra.Command{
	Use:   "getConfigData",
	Short: "A brief description of your command",
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {

		authFlag, clientAuth, err  := authenticate.SubCommandAuth(cmd)

		if err != nil {
			cmd.Help()
			return
		}

		if authFlag {
			keyId, _ := cmd.Flags().GetString("keyId")

			if keyId != "" {
				GetKeysData(keyId, *clientAuth)
			} else {
				log.Fatalln("keyId not provided. program exit")
			}
		}
	},
}


func GetKeysData(keyId string, auth client.Auth) (*kms.DescribeKeyOutput, error) {

	log.Println("Getting kms key data")

	// Kms client

	kmsClient := client.GetClient(auth, client.KMS_CLIENT).(*kms.KMS)

	// Prepare the DescribeKeyInput request

	kmsRequest := &kms.DescribeKeyInput{
		KeyId: &keyId,
	}

	// Send the DescribeKey request to AWS KMS

	kmsResponse, err := kmsClient.DescribeKey(kmsRequest)
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	log.Println(kmsResponse)

	return kmsResponse,err
}

func init() {
	GetConfigDataCmd.Flags().StringP("keyId", "t", "", "Key Id")

	if err := GetConfigDataCmd.MarkFlagRequired("keyId"); err != nil {
		fmt.Println("--keyId or -t is required", err)
	}
}
