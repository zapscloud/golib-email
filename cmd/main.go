package main

import (
	"log"
	"os"

	"github.com/zapscloud/golib-email/email_common"
	"github.com/zapscloud/golib-email/email_services"
	"github.com/zapscloud/golib-utils/utils"
)

func main() {

	testAWS_SES_SMTP_Mail()

}

func testAWS_SES_SMTP_Mail() error {
	emailConf := getAWSSES_SMTPConfig()

	svcEmail, err := email_services.NewEMailService(emailConf)
	if err != nil {
		log.Println("Error: ", err)
		return err
	}

	// Parse the Parameters
	strSender := "no-reply@clm.zapscloud.com"
	strRecipient := "karthikeyan.n@inforios.com"
	strSubject := "This is test Email with attachment"
	//strBody := "This is Test mail body"
	arrToAddresses := []string{
		strRecipient,
		"anburajan.inforios@gmail.com",
	}
	arrCCAddresses := []string{}
	strTemplateName := "./cmd/business_signup.html"
	mapTemplateData := utils.Map{
		"business_name": "Test Business",
		"email_id":      "abc@xyz.com",
		"strBody":       "1234",
	}

	// Send Test Email
	//err = svcEmail.SendEMail(strSender, strRecipient, strSubject, strBody)

	// Send Test Email
	// err = svcEmail.SendEMail2WithTemplate(
	// 	strSender,
	// 	arrToAddresses, arrCCAddresses,
	// 	strSubject,
	// 	strTemplateName, mapTemplateData)

	// Send Test Email with Attachment
	err = svcEmail.SendEMailWithTemplateAndAttachment(
		strSender,
		arrToAddresses, arrCCAddresses,
		strSubject,
		strTemplateName, mapTemplateData,
		"./cmd/cat.jpg")
	if err != nil {
		log.Println("Error: ", err)
		return err
	}
	return nil
}

func testAWS_SES_SDK_Mail() error {

	emailConf := getAWSSES_SDKConfig()

	svcEmail, err := email_services.NewEMailService(emailConf)
	if err != nil {
		log.Println("Error: ", err)
		return err
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

	return nil
}

func getAWSSES_SMTPConfig() utils.Map {
	emailConf := utils.Map{
		email_common.EMAIL_TYPE:                  email_common.EMAIL_TYPE_AWS_SES_SMTP,
		email_common.EMAIL_AWS_SES_SMTP_HOST:     os.Getenv("AWS_SES_SMTP_HOST"),
		email_common.EMAIL_AWS_SES_SMTP_PORT:     os.Getenv("AWS_SES_SMTP_PORT"),
		email_common.EMAIL_AWS_SES_SMTP_USERNAME: os.Getenv("AWS_SES_SMTP_USERNAME"),
		email_common.EMAIL_AWS_SES_SMTP_PASSWORD: os.Getenv("AWS_SES_SMTP_PASSWORD"),
	}

	return emailConf
}

func getAWSSES_SDKConfig() utils.Map {
	emailConf := utils.Map{
		email_common.EMAIL_TYPE:                  email_common.EMAIL_TYPE_AWS_SES_SDK,
		email_common.EMAIL_AWS_SES_SDK_REGION:    os.Getenv("AWS_SES_SDK_REGION"),
		email_common.EMAIL_AWS_SES_SDK_ACCESSKEY: os.Getenv("AWS_SES_SDK_ACCESSKEY"),
		email_common.EMAIL_AWS_SES_SDK_SECRETKEY: os.Getenv("AWS_SES_SDK_SECRETKEY"),
	}

	return emailConf
}
