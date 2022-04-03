package email

// ResourceCreatedResponse #/definitions/ResourceCreated
type ResourceCreatedResponse struct {
	ResourceId string `json:"resourceId"`
	Status     int    `json:"status"`
}

// TransactionalMessage #/definitions/TransactionalMessage
type TransactionalMessage struct {
	TransactionalMessageID   string `json:"transactionalMessageId"`
	TransactionalMessageName string `json:"transactionalMessageName"`
	CampaignId               string `json:"campaignId"`
	ExternalCampaignID       string `json:"externalCampaignId"`
	Subject                  string `json:"subject"`
}

// MessageActivity #/definitions/MessageActivity
type MessageActivity struct {
	Abuse             string `json:"abuse"`
	ActivityDate      string `json:"activityDate"`
	Bounce            string `json:"bounce"`
	BounceReason      string `json:"bounceReason"`
	Click             string `json:"click"`
	ClickCount        string `json:"clickCount"`
	EmailAddress      string `json:"emailAddress"`
	Open              string `json:"open"`
	OrderTotal        string `json:"orderTotal"`
	Read              string `json:"read"`
	SendDate          string `json:"sendDate"`
	Unsubscribe       string `json:"unsubscribe"`
	VisitDate         string `json:"visitDate"`
	ExternalContactID string `json:"externalContactID"`
}

// TransactionalMessageContact #/definitions/TransactionalMessageContact
type TransactionalMessageContact struct {
	Items []SegmentationFieldValue `json:"items"`
}

// SegmentationFieldValue #/definitions/SegmentationFieldValue
type SegmentationFieldValue struct {
	Value                    string `json:"value"`
	SegmentationFieldGroupID int    `json:"segmentationFieldGroupId"`
	SegmentationFieldID      int    `json:"segmentationFieldId"`
}
