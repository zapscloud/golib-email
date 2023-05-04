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

	svcEmail, err := email_services.NewStorageService(emailConf)
	if err != nil {
		log.Println("Error: ", err)
		return
	}

	emailData := utils.Map{
		email_common.EMAIL_SENDER:    "no-reply@test.com",
		email_common.EMAIL_RECIPIENT: "abc@xyz.com",
		email_common.EMAIL_SUBJECT:   "This is test Email",
		email_common.EMAIL_BODY:      "This is Test mail body",
	}
	svcEmail.SendEMail(emailData)

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
