package email_services

import (
	"bytes"
	"html/template"
	"log"
	"path"

	"github.com/zapscloud/golib-email/email_common"
	"github.com/zapscloud/golib-email/email_repository"
	"github.com/zapscloud/golib-utils/utils"
)

// EMailService - Email Service
type EMailService struct {
	emailClient reposEMailService
}

// reposEMailService - Email Service Repositories Interface
type reposEMailService interface {
	InitializeService(props utils.Map) error
	SendEMail(strSender string, strRecipient string, strSubject string, strBody string) error
	SendEMail2(strSender string, arrRecipient []string, arrCCAddresses []string, strSubject string, strBody string) error
}

// NewEMailService - Contruct EMail Service
func NewEMailService(props utils.Map) (EMailService, error) {

	// Instantiate the EMail Service
	emailService := EMailService{
		emailClient: nil,
	}

	// Get EMailType from the Parameter
	emailType, err := email_common.GetEMailType(props)
	if err != nil {
		return emailService, err
	}

	// Get the EMail's Object based on EMailType
	switch emailType {
	case email_common.EMAIL_TYPE_AWS_SES_SDK:
		emailService.emailClient = &email_repository.AWS_SES_SDKEMailServices{}

	case email_common.EMAIL_TYPE_AWS_SES_SMTP:
		// *Not Implemented yet*
		emailService.emailClient = nil

	case email_common.EMAIL_TYPE_MS_AZURE:
		// *Not Implemented yet*
		emailService.emailClient = nil

	case email_common.EMAIL_TYPE_GOOGLE:
		// *Not Implemented yet*
		emailService.emailClient = nil
	}

	if emailService.emailClient != nil {
		// Initialize the Dao
		err = emailService.initialize(props)
		if err != nil {
			return emailService, err
		}
	}

	return emailService, nil
}

func (p *EMailService) initialize(props utils.Map) error {
	var err error = nil

	if p.emailClient == nil {
		err = &utils.AppError{ErrorStatus: 412, ErrorMsg: "Initialize Error", ErrorDetail: "EMail Service is not created"}
	} else {
		err = p.emailClient.InitializeService(props)
	}

	return err
}

func (p *EMailService) SendEMail(strSender string, strRecipient string, strSubject string, strBody string) error {

	var err error = nil

	if p.emailClient == nil {
		err = &utils.AppError{ErrorStatus: 412, ErrorMsg: "Initialize Error", ErrorDetail: "EMail Service is not created"}
	} else {
		err = p.emailClient.SendEMail(strSender, strRecipient, strSubject, strBody)
	}

	return err
}

func (p *EMailService) SendEMail2(strSender string, strRecipient []string, strCCAddresses []string, strSubject string, strBody string) error {

	var err error = nil

	if p.emailClient == nil {
		err = &utils.AppError{ErrorStatus: 412, ErrorMsg: "Initialize Error", ErrorDetail: "EMail Service is not created"}
	} else {
		err = p.emailClient.SendEMail2(strSender, strRecipient, strCCAddresses, strSubject, strBody)
	}

	return err
}

func (p *EMailService) SendEMailWithTemplate(
	strSender string,
	strRecipient string,
	strSubject string,
	templateFileName string,
	templateData utils.Map) error {

	log.Println("SendEMailWithTemplate Enter=> ", strSender, strRecipient, strSubject, templateFileName, path.Base(templateFileName))

	htmlBody, err := p.convertTemplateToHTML(templateFileName, templateData)
	if err != nil {
		return err
	}

	return p.SendEMail(strSender, strRecipient, strSubject, htmlBody)
}

func (p *EMailService) SendEMail2WithTemplate(
	strSender string,
	arrToAddresses []string,
	arrCCAddresses []string,
	strSubject string,
	templateFileName string,
	templateData utils.Map) error {

	log.Println("SendEMailWithTemplate Enter=> ", strSender, arrToAddresses, arrCCAddresses, strSubject, templateFileName, path.Base(templateFileName))

	htmlBody, err := p.convertTemplateToHTML(templateFileName, templateData)
	if err != nil {
		return err
	}

	return p.emailClient.SendEMail2(strSender, arrToAddresses, arrCCAddresses, strSubject, htmlBody)
}

func (p *EMailService) convertTemplateToHTML(templateFileName string, templateData utils.Map) (string, error) {
	// Add function maps to the Template
	funcMap := template.FuncMap{
		"add": func(a, b int) int { return a + b },
		"sub": func(a, b int) int { return a - b },
		"mul": func(a, b int) int { return a * b },
		"div": func(a, b int) float32 { return float32(a) / float32(b) },
	}

	t, err := template.New(path.Base(templateFileName)).Funcs(funcMap).ParseFiles(templateFileName)
	if err != nil {
		log.Println(err)
		return "", err
	}

	log.Println("SendEMailWithTemplate ParseFiles Success")

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, templateData); err != nil {
		log.Println(err)
		return "", err
	}
	log.Println("SendEMailWithTemplate Execute Success")

	htmlBody := buf.String()

	return htmlBody, nil
}
