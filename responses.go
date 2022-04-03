package listrak

import (
	"fmt"
)

type ResourceCreatedResponse struct {
	ResourceId string `json:"resourceId"`
	Status     int    `json:"status"`
}

//ErrorResponse #/definitions/Error
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

//ToError formatting
func (e ErrorResponse) ToError() error {
	return fmt.Errorf("%v (%v): %v", e.Status, e.Error, e.Message)
}

// TransactionalMessageResponse #/definitions/TransactionalMessage
type TransactionalMessageResponse struct {
	TransactionalMessage
}

// ContactListSubscriptionResponse #/definitions/Collection[ContactListSubscription]
type ContactListSubscriptionResponse struct {
	NextPageCursor string                    `json:"nextPageCursor"`
	Data           []ContactListSubscription `json:"data"`
	Status         int                       `json:"status"`
}

//ShortCodeCollectionResponse #/definitions/Collection[ShortCode]
type ShortCodeCollectionResponse struct {
	NextPageCursor string             `json:"nextPageCursor"`
	Data           []ShortCodeMessage `json:"data"`
	Status         int                `json:"status"`
}

// SMSContactSubscriptionDetailsResponse #/definitions/Resource[SMSContactSubscriptionDetails]
type SMSContactSubscriptionDetailsResponse struct {
	Data   SMSContactSubscriptionDetails `json:"data"`
	Status int                           `json:"status"`
}

// ResourceUpdatedResponse #/definitions/ResourceUpdated
type ResourceUpdatedResponse struct {
	Status     int    `json:"status"`
	ResourceID string `json:"resourceId"`
}

//SMSContactCollectionResponse #/definitions/Collection[SMSContact]
type SMSContactCollectionResponse struct {
	Data   []ShortCodeMessage `json:"data"`
	Status int                `json:"status"`
}

// ResourceDeletedResponse #/definitions/ResourceDeleted
type ResourceDeletedResponse struct {
	Status int `json:"status"`
}

// PhoneListCollectionResponse #/definitions/Collection[PhoneList]
type PhoneListCollectionResponse struct {
	Status int         `json:"status"`
	Data   []PhoneList `json:"data"`
}

// PhoneAttributeCollectionResponse #/definitions/Collection[PhoneAttribute]
type PhoneAttributeCollectionResponse struct {
	Status int              `json:"status"`
	Data   []PhoneAttribute `json:"data"`
}

// PhoneListResponse #/definitions/PhoneList
type PhoneListResponse struct {
	PhoneList
}

// PhoneAttributeResponse #/definitions/Resource[PhoneAttribute]
type PhoneAttributeResponse struct {
	Status int            `json:"status"`
	Data   PhoneAttribute `json:"data"`
}

// ShortCodeResponse #/definitions/Resource[ShortCode]
type ShortCodeResponse struct {
	Status int              `json:"status"`
	Data   ShortCodeMessage `json:"data"`
}

// TransactionalMessageCollectionResponse #/definitions/Collection[TransactionalMessage]
type TransactionalMessageCollectionResponse struct {
	Status int                    `json:"status"`
	Data   []TransactionalMessage `json:"data"`
}
