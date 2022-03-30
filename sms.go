package listrak

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

const SMS = "/sms"

type SmsRequest struct {
	baseRequest
	basePath string
}

func NewSMSRequest() SmsRequest {
	r := SmsRequest{}
	r.basePath = SMS
	return r
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
	path := fmt.Sprintf("/v1/ShortCode/%v/Contact/%v/PhoneList", shortCodeID, phoneNumber)
	get, err := http.NewRequest("GET", path, strings.NewReader(r.params.Encode()))
	if err != nil {
		return
	}
	if cursor == "" {
		cursor = "Start"
	}
	if count == 0 {
		count = 1000
	}
	if count < 5000 {
		count = 5000
	}
	q := get.URL.Query()
	q.Add("cursor", cursor)
	q.Add("count", strconv.Itoa(count))
	get.URL.RawQuery = q.Encode()
	response, err := r.client.Do(get)
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			return
		}
	}(response.Body)
	dec := json.NewDecoder(get.Body)

	switch response.StatusCode {
	case 200:
		err = dec.Decode(&data)
	case 400:
	case 401:
	case 404:
		errResponse := ErrorResponse{}
		if err = dec.Decode(&errResponse); err != nil {
			err = fmt.Errorf("improper error response: %v", err)
		} else {
			err = fmt.Errorf("%v: %v", errResponse.Error, errResponse.Message)
		}
	default:
		err = fmt.Errorf("unhandled StatusCode: %v", response.StatusCode)
	}

	return
}

//PostContactListSubscription Subscribes a contact to an SMS list. This will only subscribe contacts that already exist on the short code
func (r SmsRequest) PostContactListSubscription() {}

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
func (r SmsRequest) GetShortCodeCollection() {}

//GetShortCodeResource Retrieve a single SMS Short Code by Short Code ID
func (r SmsRequest) GetShortCodeResource() {}

//GetTransactionalMessageCollection Retrieve a Collection of Transactional SMS Message for an SMS List
func (r SmsRequest) GetTransactionalMessageCollection() {}

//GetTransactionalMessageResource Retrieve a single Transactional SMS Message for an SMS List by ID
func (r SmsRequest) GetTransactionalMessageResource() {}

//PostTransactionalMessageSend Send a single Transactional SMS Message for an SMS List by ID
func (r SmsRequest) PostTransactionalMessageSend() {}
