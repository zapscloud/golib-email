package email_repository

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/zapscloud/golib-email/email_common"
	"github.com/zapscloud/golib-utils/utils"
)

// AWSStorageServices - AWS Storage Service structure
type AWS_SES_SDKEmailServices struct {
	awsSESSdkRegion    string
	awsSESSdkAccessKey string
	awsSESSdkSecretKey string
}

func (p *AWS_SES_SDKEmailServices) InitializeService(props utils.Map) error {

	var err error = nil

	if dataVal, dataOk := props[email_common.EMAIL_AWS_SES_SDK_REGION]; !dataOk || len(dataVal.(string)) == 0 {
		err = &utils.AppError{ErrorStatus: 400, ErrorMsg: "Bad Request", ErrorDetail: "Parameter Region is not received"}
	} else if dataVal, dataOk := props[email_common.EMAIL_AWS_SES_SDK_ACCESSKEY]; !dataOk || len(dataVal.(string)) == 0 {
		err = &utils.AppError{ErrorStatus: 400, ErrorMsg: "Bad Request", ErrorDetail: "Parameter AccessKey is not received"}
	} else if dataVal, dataOk := props[email_common.EMAIL_AWS_SES_SDK_SECRETKEY]; !dataOk || len(dataVal.(string)) == 0 {
		err = &utils.AppError{ErrorStatus: 400, ErrorMsg: "Bad Request", ErrorDetail: "Parameter SecretKey is not received"}
	}

	if err == nil {
		// Store the Parameter to member variable
		p.awsSESSdkRegion = props[email_common.EMAIL_AWS_SES_SDK_REGION].(string)
		p.awsSESSdkAccessKey = props[email_common.EMAIL_AWS_SES_SDK_ACCESSKEY].(string)
		p.awsSESSdkSecretKey = props[email_common.EMAIL_AWS_SES_SDK_SECRETKEY].(string)

		//log.Println("At Initialise:", p.awsSESSdkRegion, p.awsSESSdkAccessKey, p.awsSESSdkSecretKey)
	}

	return err
}

// Send EMail to Single Recipient
func (p *AWS_SES_SDKEmailServices) SendEMail(strSender string, strRecipient string, strSubject string, strBody string) error {

	// Convert strReceipt as Array
	toAddresses := []*string{
		aws.String(strRecipient)}

	// Empty CCAdddress
	ccAddress := []*string{}

	return p.sendEMail(strSender, toAddresses, ccAddress, strSubject, strBody)

}

// Send Email to Multiple Recipient
func (p *AWS_SES_SDKEmailServices) SendEMail2(strSender string, arrRecipients []string, arrCCAddresses []string, strSubject string, strBody string) error {

	// Create the list of ToAddresses
	var toAddresses []*string
	for _, recipient := range arrRecipients {
		toAddresses = append(toAddresses, aws.String(recipient))
	}

	// Create the list of CCAddresses
	var ccAddresses []*string
	for _, cc := range arrCCAddresses {
		ccAddresses = append(ccAddresses, aws.String(cc))
	}

	return p.sendEMail(strSender, toAddresses, ccAddresses, strSubject, strBody)

}

func (p *AWS_SES_SDKEmailServices) sendEMail(strSender string, toAddresses []*string, ccAddresses []*string, strSubject string, strBody string) error {

	log.Println("SDKEMailServices.sendEMail Enter:", strSender, toAddresses, ccAddresses, strSubject)

	// Create new Session
	sess, _ := session.NewSession(
		&aws.Config{
			Region:      aws.String(p.awsSESSdkRegion),
			Credentials: credentials.NewStaticCredentials(p.awsSESSdkAccessKey, p.awsSESSdkSecretKey, "")},
	)

	// The character encoding for the email.
	CharSet := "UTF-8"

	// Create an SES Service.
	svc := ses.New(sess)

	// Assemble the email.
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: ccAddresses,
			ToAddresses: toAddresses,
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(strBody),
				},
				// Text: &ses.Content{
				// 	Charset: aws.String(CharSet),
				// 	Data:    aws.String(TextBody),
				// },
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(strSubject),
			},
		},
		Source: aws.String(strSender),
		// Uncomment to use a configuration set
		//ConfigurationSetName: aws.String(ConfigurationSet),
	}

	log.Println("SDKEMailServices.sendEMail Before SendEmail:")

	// Attempt to send the email.
	result, err := svc.SendEmail(input)

	// Display error messages if they occur.
	log.Println("SDKEMailServices.sendEMail, After SendEmail:", err)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {

			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				log.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				log.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				log.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				log.Println(aerr.Error())
			}

			err := &utils.AppError{
				ErrorCode: aerr.Code(),
				ErrorMsg:  aerr.Error(),
			}
			return err
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			log.Println(err.Error())
			err := &utils.AppError{
				ErrorMsg: err.Error(),
			}
			return err
		}
	}

	log.Println("Email Sent")
	log.Println(result)

	return nil
}
