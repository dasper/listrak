package listrak

// BroadcastMessageRequest #/definitions/BroadcastMessage
type BroadcastMessageRequest struct {
	PhoneListID         int    `json:"phoneListId"`
	MessageContent      string `json:"messageContent"`
	MessageTitle        string `json:"messageTitle"`
	SegmentationFieldID string `json:"segmentationFieldId"`
}
