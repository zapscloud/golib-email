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
	strSender := "no-reply@abc.com"
	strRecipient := "abc@xyz.com"
	arrToAddresses := []string{"abc@xyz.com", "abc1@xyz.com"}
	arrCCAddresses := []string{}
	strSubject := "This is test Email"
	strBody := "This is Test mail body"

	log.Println("Recipient EMail ", strRecipient)
	log.Println("List of ToAddresses and CCAddresses ", arrToAddresses, arrCCAddresses)

	// Send to single EMail
	//svcEmail.SendEMail(strSender, strRecipient, strSubject, strBody)

	// Send to multiple EMail
	svcEmail.SendEMail2(strSender, arrToAddresses, arrCCAddresses, strSubject, strBody)

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
