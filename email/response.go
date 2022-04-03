package email

// TransactionalMessageCollectionResponse #/definitions/Collection[TransactionalMessage]
type TransactionalMessageCollectionResponse struct {
	Status int                    `json:"status"`
	Data   []TransactionalMessage `json:"data"`
}

// TransactionalMessageResponse #/definitions/Resource[TransactionalMessage]
type TransactionalMessageResponse struct {
	Status int                  `json:"status"`
	Data   TransactionalMessage `json:"data"`
}
