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

type TransactionalMessageResponse struct {
	Data   TransactionalMessage `json:"data"`
	Status int                  `json:"status"`
}

type ContactListCollectionResponse struct {
	NextPageCursor string        `json:"nextPageCursor"`
	Data           []ContactList `json:"data"`
	Status         int           `json:"status"`
}

type ContactList struct {
	SubscriptionState string `json:"subscriptionState"`
	SubscribeDate     string `json:"subscribeDate"`
	PhoneListId       int32  `json:"phoneListId"`
}

type TransactionalMessage struct {
	TransactionalMessageName   string `json:"transactionalMessageName"`
	TransactionalMessageStatus string `json:"transactionalMessageStatus"`
	TransactionalMessageID     int32  `json:"transactionalMessageId"`
}

//ShortCodeCollectionResponse #/definitions/Collection[ShortCode]
type ShortCodeCollectionResponse struct {
	NextPageCursor string             `json:"nextPageCursor"`
	Data           []ShortCodeMessage `json:"data"`
	Status         int                `json:"status"`
}

type ShortCodeMessage struct {
	PhoneNumberSegmentationFieldGroupName       string `json:"phoneNumberSegmentationFieldGroupName"`
	Code                                        string `json:"code"`
	Country                                     string `json:"country"`
	MerchantName                                string `json:"merchantName"`
	AcquisitionSourceSegmentationFieldGroupName string `json:"acquisitionSourceSegmentationFieldGroupName"`
	EmailListName                               string `json:"emailListName"`
	PhoneNumberSegmentationFieldName            string `json:"phoneNumberSegmentationFieldName"`
	ExternalEventName                           string `json:"externalEventName"`
	AcquisitionSourceSegmentationFieldName      string `json:"acquisitionSourceSegmentationFieldName"`
	PhoneNumberSegmentationFieldGroupId         int    `json:"phoneNumberSegmentationFieldGroupId"`
	PhoneNumberSegmentationFieldId              int    `json:"phoneNumberSegmentationFieldId"`
	ExternalEventId                             int    `json:"externalEventId"`
	AcquisitionSourceSegmentationFieldGroupId   int    `json:"acquisitionSourceSegmentationFieldGroupId"`
	EmailListId                                 int    `json:"emailListId"`
	AcquisitionSourceSegmentationFieldId        int    `json:"acquisitionSourceSegmentationFieldId"`
	ShortCodeId                                 int    `json:"shortCodeId"`
}
