package email

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/dasper/listrak"
)

const Email = "/email"

//Request required to call Email api
type Request struct {
	listrak.BaseRequest
}

//NewClient returns new Email Request client
func NewClient() Request {
	r := Request{}
	r.SetBasePath(listrak.HOST + Email)
	return r
}

func (r Request) GetCampaignCollection()                       {}
func (r Request) PostCampaignResource()                        {}
func (r Request) GetCampaignResource()                         {}
func (r Request) PutCampaignResource()                         {}
func (r Request) DeleteCampaignResource()                      {}
func (r Request) GetContactCollection()                        {}
func (r Request) PostContactResource()                         {}
func (r Request) GetContactResourceByIdentifier()              {}
func (r Request) GetConversationCollection()                   {}
func (r Request) GetConversationResource()                     {}
func (r Request) GetConversationMessageCollection()            {}
func (r Request) GetConversationMessageResource()              {}
func (r Request) GetConversationMessageActivityCollection()    {}
func (r Request) GetConversationMessageLinkCollection()        {}
func (r Request) GetConversationMessageLinkResource()          {}
func (r Request) GetConversationMessageLinkClickerCollection() {}
func (r Request) GetConversationMessageSummaryResource()       {}
func (r Request) GetConversationSummaryResource()              {}
func (r Request) GetEventCollection()                          {}
func (r Request) PostEventResource()                           {}
func (r Request) GetEventResource()                            {}
func (r Request) PutEventResource()                            {}
func (r Request) GetEventGroupCollection()                     {}
func (r Request) PostEventGroupResource()                      {}
func (r Request) GetEventGroupResource()                       {}
func (r Request) PutEventGroupResource()                       {}
func (r Request) DeleteEventGroupResource()                    {}
func (r Request) GetFolderCollection()                         {}
func (r Request) PostFolderResource()                          {}
func (r Request) GetFolderResource()                           {}
func (r Request) PutFolderResource()                           {}
func (r Request) GetIpPoolCollection()                         {}
func (r Request) GetListCollection()                           {}
func (r Request) PostListResource()                            {}
func (r Request) GetListResourceById()                         {}
func (r Request) PutListResource()                             {}
func (r Request) DeleteListResource()                          {}
func (r Request) GetListImportCollection()                     {}
func (r Request) PostImportFileResource()                      {}
func (r Request) GetListImportResource()                       {}
func (r Request) GetListImportContactCollection()              {}
func (r Request) GetListImportStatusResource()                 {}
func (r Request) GetListImportSummaryResource()                {}
func (r Request) GetMessageCollection()                        {}
func (r Request) PostMessageResource()                         {}
func (r Request) GetMessageResource()                          {}
func (r Request) GetMessageActivityCollection()                {}
func (r Request) GetMessageLinkCollection()                    {}
func (r Request) GetMessageLinkResource()                      {}
func (r Request) GetMessageLinkClickerCollection()             {}
func (r Request) GetMessageStatusResource()                    {}
func (r Request) GetMessageSummaryResource()                   {}
func (r Request) GetSavedAudienceResourceCollection()          {}
func (r Request) PostSavedMessageSend()                        {}
func (r Request) GetSavedMessageCollection()                   {}
func (r Request) GetSavedMessageResource()                     {}
func (r Request) GetSegmentationFieldCollection()              {}
func (r Request) PostSegmentationFieldResource()               {}
func (r Request) GetSegmentationFieldResource()                {}
func (r Request) PutSegmentationFieldResource()                {}
func (r Request) DeleteSegmentationFieldResource()             {}
func (r Request) GetSegmentationFieldGroupCollection()         {}
func (r Request) PostSegmentationFieldGroupResource()          {}
func (r Request) GetSegmentationFieldGroupResource()           {}
func (r Request) PutSegmentationFieldGroupResource()           {}
func (r Request) DeleteSegmentationFieldGroupResource()        {}
func (r Request) PostTransactionalMessageSend()                {}

func (r Request) GetTransactionalMessageCollection(listID int) (data TransactionalMessageCollectionResponse, err error) {
	r.NewRequest()
	path := fmt.Sprintf("/v1/List/%v/TransactionalMessage", listID)
	response, err := r.SendRequest("GET", path)
	if err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			return
		}
	}(response.Body)
	dec := json.NewDecoder(response.Body)

	switch response.StatusCode {
	case 200:
		err = dec.Decode(&data)
	case 400:
		fallthrough
	case 401:
		fallthrough
	case 404:
		err = listrak.HandleErrorResponse(dec)
	default:
		err = listrak.ErrUnhandledStatusCode
	}

	return
}
func (r Request) GetTransactionalMessageResource()           {}
func (r Request) GetTransactionalMessageActivityCollection() {}
