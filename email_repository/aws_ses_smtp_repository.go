package email_repository

import (
	"log"
	"os"
	"strconv"

	gomail "gopkg.in/gomail.v2"

	"github.com/zapscloud/golib-email/email_common"
	"github.com/zapscloud/golib-utils/utils"
)

// AWSStorageServices - AWS Storage Service structure
type AWS_SES_SMTPEmailServices struct {
	smtpHost     string
	smtpPort     string
	smtpUsername string
	smtpPassword string
}

func (p *AWS_SES_SMTPEmailServices) InitializeService(props utils.Map) error {

	var err error = nil

	if dataVal, dataOk := props[email_common.EMAIL_AWS_SES_SMTP_HOST]; !dataOk || len(dataVal.(string)) == 0 {
		err = &utils.AppError{ErrorStatus: 400, ErrorMsg: "Bad Request", ErrorDetail: "Parameter SMTP Host is not received"}
	} else if dataVal, dataOk := props[email_common.EMAIL_AWS_SES_SMTP_PORT]; !dataOk || len(dataVal.(string)) == 0 {
		err = &utils.AppError{ErrorStatus: 400, ErrorMsg: "Bad Request", ErrorDetail: "Parameter SMTP Port is not received"}
	} else if dataVal, dataOk := props[email_common.EMAIL_AWS_SES_SMTP_USERNAME]; !dataOk || len(dataVal.(string)) == 0 {
		err = &utils.AppError{ErrorStatus: 400, ErrorMsg: "Bad Request", ErrorDetail: "Parameter SMTP Username is not received"}
	} else if dataVal, dataOk := props[email_common.EMAIL_AWS_SES_SMTP_PASSWORD]; !dataOk || len(dataVal.(string)) == 0 {
		err = &utils.AppError{ErrorStatus: 400, ErrorMsg: "Bad Request", ErrorDetail: "Parameter SMTP Password is not received"}
	}

	if err == nil {
		// Store the Parameter to member variable
		p.smtpHost = props[email_common.EMAIL_AWS_SES_SMTP_HOST].(string)
		p.smtpPort = props[email_common.EMAIL_AWS_SES_SMTP_PORT].(string)
		p.smtpUsername = props[email_common.EMAIL_AWS_SES_SMTP_USERNAME].(string)
		p.smtpPassword = props[email_common.EMAIL_AWS_SES_SMTP_PASSWORD].(string)

		//log.Println("At Initialise:", p.smtpHost, p.smtpPort, p.smtpUsername, p.smtpPassword)
	}

	return err
}

// Send EMail to Single Recipient
func (p *AWS_SES_SMTPEmailServices) SendEMail(strSender string, strRecipient string, strSubject string, strBody string) error {

	// Receiver email address.
	toAddresses := []string{
		strRecipient,
	}

	// CC email address.
	ccAddresses := []string{}

	return p.sendEmailWithGomail(strSender, toAddresses, ccAddresses, strSubject, strBody, "")
}

// Send Email to Multiple Recipient
func (p *AWS_SES_SMTPEmailServices) SendEMail2(strSender string, arrRecipients []string, arrCCAddresses []string, strSubject string, strBody string) error {
	return p.sendEmailWithGomail(strSender, arrRecipients, arrCCAddresses, strSubject, strBody, "")
}

func (p *AWS_SES_SMTPEmailServices) SendEMailWithAttachment(
	strSender string,
	arrRecipient []string,
	arrCCAddresses []string,
	strSubject string,
	strBody string,
	strAttachmentFile string) error {

	return p.sendEmailWithGomail(strSender, arrRecipient, arrCCAddresses, strSubject, strBody, strAttachmentFile)
}

// func (p *AWS_SES_SMTPEmailServices) sendEMail(strSender string, toAddresses []string, ccAddresses []string, strSubject string, strBody string) error {

// 	log.Println("AWS_SES_SMTPEmailServices.sendEMail Enter:", strSender, toAddresses, ccAddresses, strSubject)

// 	// Authentication.
// 	auth := smtp.PlainAuth("", p.smtpUsername, p.smtpPassword, p.smtpHost)

// 	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

// 	// Message.
// 	message := []byte(
// 		"Subject: " + strSubject + "\r\n" + mimeHeaders + strBody + "\r\n")

// 	// Sending email.
// 	err := smtp.SendMail(p.smtpHost+":"+p.smtpPort, auth, strSender, toAddresses, message)
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}

// 	log.Println("Email Sent Successfully!")

// 	return nil

// }

func (p *AWS_SES_SMTPEmailServices) sendEmailWithGomail(
	strSender string,
	toAddresses []string,
	ccAddresses []string,
	strSubject string,
	strBody string,
	strAttachment string) error {

	msg := gomail.NewMessage()
	msg.SetHeader("From", strSender)
	msg.SetHeader("To", toAddresses...)
	msg.SetHeader("Cc", ccAddresses...)
	msg.SetHeader("Subject", strSubject)
	msg.SetBody("text/html", strBody)

	// Attachment
	if !utils.IsEmpty(strAttachment) {
		if _, err := os.Stat(strAttachment); err == nil {
			msg.Attach(strAttachment)
			log.Println("Attachment file found ", strAttachment)
		} else {
			log.Println("Attachment file **not found** ", strAttachment)
		}
	}

	intPort, err := strconv.Atoi(p.smtpPort)
	if err != nil {
		return err
	}

	n := gomail.NewDialer(p.smtpHost, intPort, p.smtpUsername, p.smtpPassword)

	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		return err
	}

	log.Println("Email Sent Successfully!")

	return nil
}
