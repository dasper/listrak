package listrak

import (
	"bytes"
	"net/http"
	"strings"
)

const Email = "/email"

//EmailRequest required to call Email api
type EmailRequest struct {
	baseRequest
	basePath string
}

//NewEmailClient return SmsRequest
func NewEmailClient() EmailRequest {
	r := EmailRequest{}
	r.basePath = HOST + Email
	return r
}

func (r *EmailRequest) newRequest() {
	r.payload = bytes.NewBuffer(nil)
	r.client = http.Client{}
}

func (r EmailRequest) sendRequest(method string, path string) (response *http.Response, err error) {
	method = strings.ToUpper(method)
	firstCharacter := path[0:1]
	if firstCharacter != "/" {
		path = "/" + path
	}
	fullPath := r.basePath + path
	r.request, err = http.NewRequest(method, fullPath, r.payload)
	if err != nil {
		return
	}
	err = r.setHeaders()
	if err != nil {
		return
	}
	response, err = r.client.Do(r.request)

	return
}

func GetCampaignCollection()                       {}
func PostCampaignResource()                        {}
func GetCampaignResource()                         {}
func PutCampaignResource()                         {}
func DeleteCampaignResource()                      {}
func GetContactCollection()                        {}
func PostContactResource()                         {}
func GetContactResourceByIdentifier()              {}
func GetConversationCollection()                   {}
func GetConversationResource()                     {}
func GetConversationMessageCollection()            {}
func GetConversationMessageResource()              {}
func GetConversationMessageActivityCollection()    {}
func GetConversationMessageLinkCollection()        {}
func GetConversationMessageLinkResource()          {}
func GetConversationMessageLinkClickerCollection() {}
func GetConversationMessageSummaryResource()       {}
func GetConversationSummaryResource()              {}
func GetEventCollection()                          {}
func PostEventResource()                           {}
func GetEventResource()                            {}
func PutEventResource()                            {}
func GetEventGroupCollection()                     {}
func PostEventGroupResource()                      {}
func GetEventGroupResource()                       {}
func PutEventGroupResource()                       {}
func DeleteEventGroupResource()                    {}
func GetFolderCollection()                         {}
func PostFolderResource()                          {}
func GetFolderResource()                           {}
func PutFolderResource()                           {}
func GetIpPoolCollection()                         {}
func GetListCollection()                           {}
func PostListResource()                            {}
func GetListResourceById()                         {}
func PutListResource()                             {}
func DeleteListResource()                          {}
func GetListImportCollection()                     {}
func PostImportFileResource()                      {}
func GetListImportResource()                       {}
func GetListImportContactCollection()              {}
func GetListImportStatusResource()                 {}
func GetListImportSummaryResource()                {}
func GetMessageCollection()                        {}
func PostMessageResource()                         {}
func GetMessageResource()                          {}
func GetMessageActivityCollection()                {}
func GetMessageLinkCollection()                    {}
func GetMessageLinkResource()                      {}
func GetMessageLinkClickerCollection()             {}
func GetMessageStatusResource()                    {}
func GetMessageSummaryResource()                   {}
func GetSavedAudienceResourceCollection()          {}
func PostSavedMessageSend()                        {}
func GetSavedMessageCollection()                   {}
func GetSavedMessageResource()                     {}
func GetSegmentationFieldCollection()              {}
func PostSegmentationFieldResource()               {}
func GetSegmentationFieldResource()                {}
func PutSegmentationFieldResource()                {}
func DeleteSegmentationFieldResource()             {}
func GetSegmentationFieldGroupCollection()         {}
func PostSegmentationFieldGroupResource()          {}
func GetSegmentationFieldGroupResource()           {}
func PutSegmentationFieldGroupResource()           {}
func DeleteSegmentationFieldGroupResource()        {}
func PostTransactionalMessageSend()                {}
func GetTransactionalMessageCollection()           {}
func GetTransactionalMessageResource()             {}
func GetTransactionalMessageActivityCollection()   {}
