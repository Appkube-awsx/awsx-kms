package controllers

import (
	"github.com/Appkube-awsx/awsx-kms/command"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/aws/aws-sdk-go/service/kms"
	"log"
)

func GetKmsByAccountNo(vaultUrl string, vaultToken string, accountNo string, region string) (*kms.ListKeysOutput, error) {
	authFlag, clientAuth, err := authenticate.AuthenticateData(vaultUrl, vaultToken, accountNo, region, "", "", "", "")
	return GetKmsByFlagAndClientAuth(authFlag, clientAuth, err)
}

func GetKmsByUserCreds(region string, accessKey string, secretKey string, crossAccountRoleArn string, externalId string) (*kms.ListKeysOutput, error) {
	authFlag, clientAuth, err := authenticate.AuthenticateData("", "", "", region, accessKey, secretKey, crossAccountRoleArn, externalId)
	return GetKmsByFlagAndClientAuth(authFlag, clientAuth, err)
}

func GetKmsByFlagAndClientAuth(authFlag bool, clientAuth *client.Auth, err error) (*kms.ListKeysOutput, error) {
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if !authFlag {
		log.Println(err.Error())
		return nil, err
	}
	response, err := command.ListKeys(*clientAuth)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return response, nil
}

func GetKms(clientAuth *client.Auth) (*kms.ListKeysOutput, error) {
	response, err := command.ListKeys(*clientAuth)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return response, nil
}
