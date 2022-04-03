package email

// TransactionalMessageCollectionResponse #/definitions/Collection[TransactionalMessage]
type TransactionalMessageCollectionResponse struct {
	Data   []TransactionalMessage `json:"data"`
	Status int                    `json:"status"`
}

// TransactionalMessageResponse #/definitions/Resource[TransactionalMessage]
type TransactionalMessageResponse struct {
	Data   TransactionalMessage `json:"data"`
	Status int                  `json:"status"`
}

//MessageActivityCollectionResponse #/definitions/CollectionPaged[MessageActivity]
type MessageActivityCollectionResponse struct {
	NextPageCursor string            `json:"nextPageCursor"`
	Data           []MessageActivity `json:"data"`
	Status         int               `json:"status"`
}
