package email

// TransactionalMessage #/definitions/TransactionalMessage
type TransactionalMessage struct {
	TransactionalMessageID   string `json:"transactionalMessageId"`
	TransactionalMessageName string `json:"transactionalMessageName"`
	CampaignId               string `json:"campaignId"`
	ExternalCampaignId       string `json:"externalCampaignId"`
	Subject                  string `json:"subject"`
}
