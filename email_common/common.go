package email_common

import "github.com/zapscloud/golib-utils/utils"

// Enums
type EMailType byte

const (
	EMAIL_TYPE_NONE EMailType = iota
	EMAIL_TYPE_AWS_SES_SDK
	EMAIL_TYPE_AWS_SES_SMTP
	EMAIL_TYPE_MS_AZURE
	EMAIL_TYPE_GOOGLE
	EMAIL_TYPE_PLACEHOLDER_LAST // Only a place holder
)

const (
	EMAIL_TYPE = "email_type"

	// Param for AWS_S3
	EMAIL_AWS_SES_SDK_REGION    = "aws_ses_sdk_region"
	EMAIL_AWS_SES_SDK_ACCESSKEY = "aws_ses_sdk_accesskey"
	EMAIL_AWS_SES_SDK_SECRETKEY = "aws_ses_sdk_secretkey"

	EMAIL_SENDER    = "sender"
	EMAIL_RECIPIENT = "recipient"
	EMAIL_SUBJECT   = "subject"
	EMAIL_BODY      = "body"
)

func GetEMailType(props utils.Map) (EMailType, error) {

	dataVal, dataOk := props[EMAIL_TYPE]

	if !dataOk {
		err := &utils.AppError{ErrorStatus: 401, ErrorCode: "401", ErrorMsg: "EMailType not found", ErrorDetail: "EMailType value is not received"}
		return EMAIL_TYPE_NONE, err
	}

	// Convert it to String type
	storageType := dataVal.(EMailType)

	if storageType <= EMAIL_TYPE_NONE || storageType >= EMAIL_TYPE_PLACEHOLDER_LAST {

		err := &utils.AppError{ErrorStatus: 401, ErrorCode: "401", ErrorMsg: "Invalid EMailType", ErrorDetail: "StorageType value is Invalid"}
		return EMAIL_TYPE_NONE, err
	}

	return storageType, nil
}
