package sms

type ResourceCreatedResponse struct {
	ResourceId string `json:"resourceId"`
	Status     int    `json:"status"`
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

// ContactSubscriptionDetailsResponse #/definitions/Resource[SMSContactSubscriptionDetails]
type ContactSubscriptionDetailsResponse struct {
	Data   ContactSubscriptionDetails `json:"data"`
	Status int                        `json:"status"`
}

// ResourceUpdatedResponse #/definitions/ResourceUpdated
type ResourceUpdatedResponse struct {
	ResourceID string `json:"resourceId"`
	Status     int    `json:"status"`
}

// ContactCollectionResponse #/definitions/Collection[SMSContact]
type ContactCollectionResponse struct {
	Data   []ShortCodeMessage `json:"data"`
	Status int                `json:"status"`
}

// ResourceDeletedResponse #/definitions/ResourceDeleted
type ResourceDeletedResponse struct {
	Status int `json:"status"`
}

// PhoneListCollectionResponse #/definitions/Collection[PhoneList]
type PhoneListCollectionResponse struct {
	Data   []PhoneList `json:"data"`
	Status int         `json:"status"`
}

// PhoneAttributeCollectionResponse #/definitions/Collection[PhoneAttribute]
type PhoneAttributeCollectionResponse struct {
	Data   []PhoneAttribute `json:"data"`
	Status int              `json:"status"`
}

// PhoneListResponse #/definitions/PhoneList
type PhoneListResponse struct {
	Data   PhoneList `json:"data"`
	Status int       `json:"status"`
}

// PhoneAttributeResponse #/definitions/Resource[PhoneAttribute]
type PhoneAttributeResponse struct {
	Data   PhoneAttribute `json:"data"`
	Status int            `json:"status"`
}

// ShortCodeResponse #/definitions/Resource[ShortCode]
type ShortCodeResponse struct {
	Data   ShortCodeMessage `json:"data"`
	Status int              `json:"status"`
}

// TransactionalMessageCollectionResponse #/definitions/Collection[TransactionalMessage]
type TransactionalMessageCollectionResponse struct {
	Data   []TransactionalMessage `json:"data"`
	Status int                    `json:"status"`
}
