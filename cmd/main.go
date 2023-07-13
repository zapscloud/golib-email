package main

import (
	"log"
	"os"

	"github.com/zapscloud/golib-email/email_common"
	"github.com/zapscloud/golib-email/email_services"
	"github.com/zapscloud/golib-utils/utils"
)

func main() {
	emailConf := getAWSSESSDKConfig()

	svcEmail, err := email_services.NewEMailService(emailConf)
	if err != nil {
		log.Println("Error: ", err)
		return
	}

	// Parse the Parameters
	strSender := "no-reply@test.com"
	strRecipient := "abc@xyz.com"
	strSubject := "This is test Email"
	strBody := "This is Test mail body"

	svcEmail.SendEMail(strSender, strRecipient, strSubject, strBody)

}

func getAWSSESSDKConfig() utils.Map {
	emailConf := utils.Map{
		email_common.EMAIL_TYPE:                  email_common.EMAIL_TYPE_AWS_SES_SDK,
		email_common.EMAIL_AWS_SES_SDK_REGION:    os.Getenv("AWS_SES_SDK_REGION"),
		email_common.EMAIL_AWS_SES_SDK_ACCESSKEY: os.Getenv("AWS_SES_SDK_ACCESSKEY"),
		email_common.EMAIL_AWS_SES_SDK_SECRETKEY: os.Getenv("AWS_SES_SDK_SECRETKEY"),
	}

	return emailConf
}
