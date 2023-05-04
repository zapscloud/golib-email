package email_services

import (
	"github.com/zapscloud/golib-email/email_common"
	"github.com/zapscloud/golib-email/email_repository"
	"github.com/zapscloud/golib-utils/utils"
)

// EMailService - Email Service
type EMailService interface {
	InitializeEMailService(props utils.Map) error
	SendEMail(props utils.Map) error
}

// NewEMailService - Contruct EMail Service
func NewStorageService(props utils.Map) (EMailService, error) {
	var emailClient EMailService = nil

	// Get StorageType from the Parameter
	storageType, err := email_common.GetEMailType(props)
	if err != nil {
		return nil, err
	}

	// Get the Storage's Object based on StorageType
	switch storageType {
	case email_common.EMAIL_TYPE_AWS_SES_SDK:
		emailClient = &email_repository.AWS_SES_SDKEMailServices{}

	case email_common.EMAIL_TYPE_AWS_SES_SMTP:
		emailClient = &email_repository.AWS_SES_SDKEMailServices{}

	case email_common.EMAIL_TYPE_MS_AZURE:
		// *Not Implemented yet*
		emailClient = nil

	case email_common.EMAIL_TYPE_GOOGLE:
		// *Not Implemented yet*
		emailClient = nil
	}

	if emailClient != nil {
		// Initialize the Dao
		err = emailClient.InitializeEMailService(props)
		if err != nil {
			return nil, err
		}
	}

	return emailClient, nil
}
