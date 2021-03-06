package sms

import (
	"github.com/dasper/listrak"
)

// BroadcastMessage #/definitions/BroadcastMessage
type BroadcastMessage struct {
	MessageContent      string `json:"messageContent"`
	MessageTitle        string `json:"messageTitle"`
	SegmentationFieldID string `json:"segmentationFieldId"`
	PhoneListID         int    `json:"phoneListId"`
}

// ContactListSubscription #/definitions/ContactListSubscription
type ContactListSubscription struct {
	SubscriptionState string           `json:"subscriptionState"`
	SubscribeDate     listrak.DateTime `json:"subscribeDate"`
	PhoneListID       int32            `json:"phoneListId"`
}

// TransactionalMessage #/definitions/TransactionalMessage
type TransactionalMessage struct {
	TransactionalMessageName   string `json:"transactionalMessageName"`
	TransactionalMessageStatus string `json:"transactionalMessageStatus"`
	TransactionalMessageID     int32  `json:"transactionalMessageId"`
}

// Contact #/definitions/SMSContact
type Contact struct {
	PhoneNumber             string                  `json:"phoneNumber,omitempty"`
	EmailAddress            string                  `json:"emailAddress,omitempty"`
	FirstName               string                  `json:"firstName,omitempty"`
	LastName                string                  `json:"lastName,omitempty"`
	Birthday                string                  `json:"birthday,omitempty"`
	PostalCode              string                  `json:"postalCode,omitempty"`
	SegmentationFieldValues []PhoneContactAttribute `json:"segmentationFieldValues,omitempty"`
	OptedOut                bool                    `json:"optedOut,omitempty"`
}

// ShortCodeMessage #/definitions/ShortCode
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
	ExternalEventID                             int    `json:"externalEventId"`
	AcquisitionSourceSegmentationFieldGroupId   int    `json:"acquisitionSourceSegmentationFieldGroupId"`
	EmailListID                                 int    `json:"emailListId"`
	AcquisitionSourceSegmentationFieldId        int    `json:"acquisitionSourceSegmentationFieldId"`
	ShortCodeID                                 int    `json:"shortCodeId"`
}

// ContactSubscriptionDetails #/definitions/SMSContactSubscriptionDetails
type ContactSubscriptionDetails struct {
	SubscribeDate           listrak.DateTime        `json:"subscribeDate"`
	UnsubscribeDate         listrak.DateTime        `json:"unsubscribeDate"`
	PhoneNumber             string                  `json:"phoneNumber"`
	EmailAddress            string                  `json:"emailAddress"`
	FirstName               string                  `json:"firstName"`
	LastName                string                  `json:"lastName"`
	Birthday                string                  `json:"birthday"`
	PostalCode              string                  `json:"postalCode"`
	SegmentationFieldValues []PhoneContactAttribute `json:"segmentationFieldValues"`
	OptedOut                bool                    `json:"optedOut"`
}

// PhoneContactAttribute #/definitions/SMSPhoneContactAttribute
type PhoneContactAttribute struct {
	Value               string `json:"value,omitempty"`
	SegmentationFieldID int    `json:"segmentationFieldId,omitempty"`
}

// PhoneList #/definitions/PhoneList
type PhoneList struct {
	MessageLimitTimeFrame string           `json:"messageLimitTimeFrame"`
	PhoneListName         string           `json:"phoneListName"`
	CreateDate            listrak.DateTime `json:"createDate"`
	Status                string           `json:"status"`
	SmsListType           string           `json:"smsListType"`
	MessageLimit          int              `json:"messageLimit"`
	PhoneListID           int              `json:"phoneListId"`
	RequireDoubleOptIn    bool             `json:"requireDoubleOptIn"`
}

// PhoneAttribute #/definitions/PhoneAttribute
type PhoneAttribute struct {
	SegmentationFieldName string `json:"segmentationFieldName"`
	DataType              string `json:"dataType"`
	StatusFlag            string `json:"statusFlag"`
	SegmentationFieldID   int    `json:"segmentationFieldId"`
	MaxLength             int    `json:"maxLength"`
}
