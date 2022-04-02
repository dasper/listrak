package listrak

import (
	"bytes"
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

func (r *SmsRequest) newRequest() {
	r.payload = bytes.NewBuffer(nil)
	r.client = http.Client{}
}

func (r SmsRequest) sendRequest(method string, path string) (response *http.Response, err error) {
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

//PostImmediateBroadcast Immediately send an SMS broadcast message to all subscribed contacts on an SMS List
func (r SmsRequest) PostImmediateBroadcast(shortCodeID int, broadcastMessage BroadcastMessageRequest) (data ResourceCreatedResponse, err error) {
	r.newRequest()
	path := fmt.Sprintf("/v1/ShortCode/%v/Broadcast/Immediate", shortCodeID)

	jsonData, err := json.Marshal(broadcastMessage)
	if err != nil {
		return
	}
	r.payload = bytes.NewBuffer(jsonData)
	response, err := r.sendRequest("POST", path)
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
		err = handleErrorResponse(dec)
	default:
		err = ErrUnhandledStatusCode
	}

	return
}

//GetContactCollection Returns all contacts that exist on a specific SMS list by opt-out status
func (r SmsRequest) GetContactCollection(shortCodeID int, phoneListId int) (data SMSContactCollectionResponse, err error) {
	r.newRequest()
	path := fmt.Sprintf("/v1/ShortCode/%v/PhoneList/%v/Contact", shortCodeID, phoneListId)
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
		err = handleErrorResponse(dec)
	default:
		err = ErrUnhandledStatusCode
	}

	return
}

//PostContactListResource Creates and subscribes a new contact for a phone number if the contact does not already exist on the short code
func (r SmsRequest) PostContactListResource(shortCodeID int, phoneListId int) {
	path := "/v1/ShortCode/{shortCodeID}/PhoneList/{phoneListId}/Contact"
	fmt.Println(path)
}

//GetContactResource Returns a contact for a phone number
func (r SmsRequest) GetContactResource(shortCodeID int, phoneNumber int) {
	path := "/v1/ShortCode/{shortCodeID}/Contact/{phoneNumber}"
	fmt.Println(path)
}

//PutContactResource Updates a contact's information for a phone number
func (r SmsRequest) PutContactResource(shortCodeID int) {
	path := "/v1/ShortCode/{shortCodeID}/Contact"
	fmt.Println(path)
}

//GetContactListCollection Retrieve a collection of SMS Lists that a contact belongs to along with subscription status
func (r SmsRequest) GetContactListCollection(shortCodeID, phoneNumber int, cursor string, count int) (data ContactListCollectionResponse, err error) {
	path := fmt.Sprintf("/v1/ShortCode/%v/Contact/%v/PhoneList", shortCodeID, phoneNumber)
	if cursor == "" {
		cursor = "Start"
	}
	if count == 0 {
		count = 1000
	}
	if count < 5000 {
		count = 5000
	}
	params := url.Values{}
	params.Set("cursor", cursor)
	params.Set("count", strconv.Itoa(count))
	r.payload = strings.NewReader(params.Encode())
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
		err = handleErrorResponse(dec)
	default:
		err = ErrUnhandledStatusCode
	}

	return
}

//PostContactListSubscription Subscribes a contact to an SMS list. This will only subscribe contacts that already exist on the short code
func (r SmsRequest) PostContactListSubscription(shortCodeID int, phoneNumber string, phoneListId int) {
	path := fmt.Sprintf("/v1/ShortCode/%v/Contact/%v/PhoneList/%v", shortCodeID, phoneNumber, phoneListId)
	fmt.Println(path)
}

//DeleteUnsubscribeContactListSubscription Unsubscribes a contact from an SMS List
func (r SmsRequest) DeleteUnsubscribeContactListSubscription(shortCodeID int, phoneNumber int, phoneListId int) {
	path := "/v1/ShortCode/{shortCodeID}/ContactUnsubscribe/{phoneNumber}/PhoneList/{phoneListId}"
	fmt.Println(path)
}

//GetListCollection Retrieve a collection of SMS Lists for a shortcode by list status
func (r SmsRequest) GetListCollection(shortCodeID int) {
	path := "/v1/ShortCode/{shortCodeID}/PhoneList"
	fmt.Println(path)
}

//GetListResource Retrieve a single SMS List by List ID
func (r SmsRequest) GetListResource(shortCodeID int, phoneListId int) {
	path := "/v1/ShortCode/{shortCodeID}/PhoneList/{phoneListId}"
	fmt.Println(path)
}

//GetPhoneAttribute Returns the specified profile field
func (r SmsRequest) GetPhoneAttribute(shortCodeID int) {
	path := "/v1/ShortCode/{shortCodeID}/SegmentationField/{segmentationFieldId}"
	fmt.Println(path)
}

//GetPhoneAttributeCollection Returns the collection of profile fields that exist for the company associated with the specified Short Code
func (r SmsRequest) GetPhoneAttributeCollection(shortCodeID int) {
	path := "/v1/ShortCode/{shortCodeID}/SegmentationField"
	fmt.Println(path)
}

//GetShortCodeCollection Retrieve a collection of Short Code objects for a Company
func (r SmsRequest) GetShortCodeCollection() (data ShortCodeCollectionResponse, err error) {
	r.newRequest()
	path := "/v1/ShortCode"
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
		err = handleErrorResponse(dec)
	default:
		err = ErrUnhandledStatusCode
	}

	return
}

//GetShortCodeResource Retrieve a single SMS Short Code by Short Code ID
func (r SmsRequest) GetShortCodeResource(shortCodeID int) {
	path := "/v1/ShortCode/{shortCodeID}"
	fmt.Println(path)
}

//GetTransactionalMessageCollection Retrieve a Collection of Transactional SMS Message for an SMS List
func (r SmsRequest) GetTransactionalMessageCollection(shortCodeID int, phoneListId int) {
	path := fmt.Sprintf(
		"/v1/ShortCode/%v/PhoneList/%v/TransactionalMessage",
		shortCodeID, phoneListId,
	)
	fmt.Println(path)
}

//GetTransactionalMessageResource Retrieve a single Transactional SMS Message for an SMS List by ID
func (r SmsRequest) GetTransactionalMessageResource(shortCodeID int, phoneListId int, transactionalMessageId int) {
	path := fmt.Sprintf(
		"/v1/ShortCode/%v/PhoneList/%v/TransactionalMessage/%v",
		shortCodeID, phoneListId, transactionalMessageId,
	)
	fmt.Println(path)
}

//PostTransactionalMessageSend Send a single Transactional SMS Message for an SMS List by ID
func (r SmsRequest) PostTransactionalMessageSend(shortCodeID int, phoneListId int, transactionalMessageId int) {
	path := fmt.Sprintf(
		"/v1/ShortCode/%v/PhoneList/%v/TransactionalMessage/%v/Message",
		shortCodeID, phoneListId, transactionalMessageId,
	)
	fmt.Println(path)
}
