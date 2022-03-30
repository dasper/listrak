package listrak

type ResourceCreatedResponse struct {
	ResourceId string `json:"resourceId"`
	Status     int    `json:"status"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Status  int    `json:"status"`
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
