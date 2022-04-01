package listrak

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

//SMS base url
const SMS = "/sms"

//SmsRequest required to call SMS api
type SmsRequest struct {
	baseRequest
	basePath string
}

//NewSMSClient return SmsRequest
func NewSMSClient() SmsRequest {
	r := SmsRequest{}
	r.basePath = HOST + SMS
	return r
}

func (r *SmsRequest) new() {
	r.params = make(url.Values)
	r.client = http.Client{}
}

func (r SmsRequest) sendRequest(method string, path string) (response *http.Response, err error) {
	r.request, err = http.NewRequest(method, r.basePath+path, strings.NewReader(r.params.Encode()))
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

//PostImmediateBroadcast Immediately send an SMS broadcast message to all subscribed contacts on an SMS List
func (r SmsRequest) PostImmediateBroadcast() {}

//GetContactCollection Returns all contacts that exist on a specific SMS list by opt-out status
func (r SmsRequest) GetContactCollection() {}

//PostContactListResource Creates and subscribes a new contact for a phone number if the contact does not already exist on the short code
func (r SmsRequest) PostContactListResource() {}

//GetContactResource Returns a contact for a phone number
func (r SmsRequest) GetContactResource() {}

//PutContactResource Updates a contact's information for a phone number
func (r SmsRequest) PutContactResource() {}

//GetContactListCollection Retrieve a collection of SMS Lists that a contact belongs to along with subscription status
func (r SmsRequest) GetContactListCollection(shortCodeID, phoneNumber int, cursor string, count int) (data ContactListCollectionResponse, err error) {
	r.new()
	path := fmt.Sprintf("/v1/ShortCode/%v/Contact/%v/PhoneList", shortCodeID, phoneNumber)
	//r.params = url.Values{}
	if cursor == "" {
		cursor = "Start"
	}
	if count == 0 {
		count = 1000
	}
	if count < 5000 {
		count = 5000
	}
	r.params.Set("cursor", cursor)
	r.params.Set("count", strconv.Itoa(count))
	response, err := r.sendRequest("GET", path)
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
		errResponse := ErrorResponse{}
		if err = dec.Decode(&errResponse); err != nil {
			err = fmt.Errorf("improper error response: %v", err)
		} else {
			err = errResponse.ToError()
		}
	default:
		err = ErrUnhandledStatusCode
	}

	return
}

//PostContactListSubscription Subscribes a contact to an SMS list. This will only subscribe contacts that already exist on the short code
func (r SmsRequest) PostContactListSubscription(shortCodeId int, phoneNumber string, phoneListId int) {
	path := fmt.Sprintf("v1/ShortCode/%v/Contact/%v/PhoneList/%v", shortCodeId, phoneNumber, phoneListId)
	fmt.Println(path)
}

//DeleteUnsubscribeContactListSubscription Unsubscribes a contact from an SMS List
func (r SmsRequest) DeleteUnsubscribeContactListSubscription() {}

//GetListCollection Retrieve a collection of SMS Lists for a shortcode by list status
func (r SmsRequest) GetListCollection() {}

//GetListResource Retrieve a single SMS List by List ID
func (r SmsRequest) GetListResource() {}

//GetPhoneAttribute Returns the specified profile field
func (r SmsRequest) GetPhoneAttribute() {}

//GetPhoneAttributeCollection Returns the collection of profile fields that exist for the company associated with the specified Short Code
func (r SmsRequest) GetPhoneAttributeCollection() {}

//GetShortCodeCollection Retrieve a collection of Short Code objects for a Company
func (r SmsRequest) GetShortCodeCollection() (data ShortCodeCollectionResponse, err error) {
	r.new()
	path := "v1/ShortCode"
	response, err := r.sendRequest("get", path)
	if err != nil {
		return
	}
	dec := json.NewDecoder(response.Body)
	switch response.StatusCode {
	case 200:
		err = dec.Decode(&data)
	case 400:
		fallthrough
	case 401:
		errResponse := ErrorResponse{}
		if err = dec.Decode(&errResponse); err != nil {
			err = fmt.Errorf("improper error response: %v", err)
		} else {
			err = fmt.Errorf("%v: %v", errResponse.Error, errResponse.Message)
		}
	default:
		err = ErrUnhandledStatusCode
	}

	return
}

//GetShortCodeResource Retrieve a single SMS Short Code by Short Code ID
func (r SmsRequest) GetShortCodeResource() {}

//GetTransactionalMessageCollection Retrieve a Collection of Transactional SMS Message for an SMS List
func (r SmsRequest) GetTransactionalMessageCollection() {}

//GetTransactionalMessageResource Retrieve a single Transactional SMS Message for an SMS List by ID
func (r SmsRequest) GetTransactionalMessageResource() {}

//PostTransactionalMessageSend Send a single Transactional SMS Message for an SMS List by ID
func (r SmsRequest) PostTransactionalMessageSend() {}
