package sms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/dasper/listrak"
)

//SMS base url
const SMS = "/sms"

//Request required to call SMS api
type Request struct {
	listrak.BaseRequest
}

//NewClient //NewClient returns new SMS Request client
func NewClient() Request {
	r := Request{}
	r.SetBasePath(listrak.HOST + SMS)
	return r
}

//PostImmediateBroadcast Immediately send an SMS broadcast message to all subscribed contacts on an SMS List
func (r Request) PostImmediateBroadcast(shortCodeID int, broadcastMessage BroadcastMessage) (data ResourceCreatedResponse, err error) {
	r.NewRequest()
	path := fmt.Sprintf("/v1/ShortCode/%v/Broadcast/Immediate", shortCodeID)

	jsonData, err := json.Marshal(broadcastMessage)
	if err != nil {
		return
	}
	r.Payload = bytes.NewBuffer(jsonData)
	response, err := r.SendRequest("POST", path)
	if err != nil {
		return
	}
	defer func() {
		err = response.Body.Close()
		if err != nil {
			return
		}
	}()
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

//GetContactCollection Returns all contacts that exist on a specific SMS list by opt-out status
func (r Request) GetContactCollection(shortCodeID int, phoneListID int) (data ContactCollectionResponse, err error) {
	r.NewRequest()
	path := fmt.Sprintf("/v1/ShortCode/%v/PhoneList/%v/Contact", shortCodeID, phoneListID)
	response, err := r.SendRequest("GET", path)
	if err != nil {
		return
	}
	defer func() {
		err = response.Body.Close()
		if err != nil {
			return
		}
	}()
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

//PostContactListResource Creates and subscribes a new contact for a phone number if the contact does not already exist on the short code
func (r Request) PostContactListResource(shortCodeID int, phoneListID int, SMSContact Contact) (data ResourceCreatedResponse, err error) {
	r.NewRequest()
	path := fmt.Sprintf("/v1/ShortCode/%v/PhoneList/%v/Contact", shortCodeID, phoneListID)
	jsonData, err := json.Marshal(SMSContact)
	if err != nil {
		return
	}
	r.Payload = bytes.NewBuffer(jsonData)
	response, err := r.SendRequest("POST", path)
	if err != nil {
		return
	}
	defer func() {
		err = response.Body.Close()
		if err != nil {
			return
		}
	}()
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

//GetContactResource Returns a contact for a phone number
func (r Request) GetContactResource(shortCodeID int, phoneNumber int) (data ContactSubscriptionDetailsResponse, err error) {
	r.NewRequest()
	path := fmt.Sprintf("/v1/ShortCode/%v/Contact/%v", shortCodeID, phoneNumber)
	response, err := r.SendRequest("GET", path)
	if err != nil {
		return
	}
	defer func() {
		err = response.Body.Close()
		if err != nil {
			return
		}
	}()
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

//PutContactResource Updates a contact's information for a phone number
func (r Request) PutContactResource(shortCodeID int) (data ResourceUpdatedResponse, err error) {
	r.NewRequest()
	path := fmt.Sprintf("/v1/ShortCode/%v/Contact", shortCodeID)
	response, err := r.SendRequest("PUT", path)
	if err != nil {
		return
	}
	defer func() {
		err = response.Body.Close()
		if err != nil {
			return
		}
	}()
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

//GetContactListCollection Retrieve a collection of SMS Lists that a contact belongs to along with subscription status
func (r Request) GetContactListCollection(shortCodeID, phoneNumber int, cursor string, count int) (data ContactListSubscriptionResponse, err error) {
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
	r.Payload = strings.NewReader(params.Encode())
	response, err := r.SendRequest("GET", path)
	if err != nil {
		return
	}

	defer func() {
		err = response.Body.Close()
		if err != nil {
			return
		}
	}()
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

//PostContactListSubscription Subscribes a contact to an SMS list. This will only subscribe contacts that already exist on the short code
func (r Request) PostContactListSubscription(shortCodeID int, phoneNumber string, phoneListID int) (data ResourceCreatedResponse, err error) {
	path := fmt.Sprintf("/v1/ShortCode/%v/Contact/%v/PhoneList/%v", shortCodeID, phoneNumber, phoneListID)
	response, err := r.SendRequest("POST", path)
	if err != nil {
		return
	}
	defer func() {
		err = response.Body.Close()
		if err != nil {
			return
		}
	}()
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

//DeleteUnsubscribeContactListSubscription Unsubscribes a contact from an SMS List
func (r Request) DeleteUnsubscribeContactListSubscription(shortCodeID int, phoneNumber int, phoneListID int) (data ResourceDeletedResponse, err error) {
	path := fmt.Sprintf("/v1/ShortCode/%v/ContactUnsubscribe/%v/PhoneList/%v", shortCodeID, phoneNumber, phoneListID)
	response, err := r.SendRequest("DELETE", path)
	if err != nil {
		return
	}
	defer func() {
		err = response.Body.Close()
		if err != nil {
			return
		}
	}()
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

//GetListCollection Retrieve a collection of SMS Lists for a shortcode by list status
func (r Request) GetListCollection(shortCodeID int) (data PhoneListCollectionResponse, err error) {
	path := fmt.Sprintf("/v1/ShortCode/%v/PhoneList", shortCodeID)
	response, err := r.SendRequest("GET", path)
	if err != nil {
		return
	}
	defer func() {
		err = response.Body.Close()
		if err != nil {
			return
		}
	}()
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

//GetListResource Retrieve a single SMS List by List ID
func (r Request) GetListResource(shortCodeID int, phoneListID int) (data PhoneListResponse, err error) {
	path := fmt.Sprintf("/v1/ShortCode/%v/PhoneList/%v", shortCodeID, phoneListID)
	response, err := r.SendRequest("GET", path)
	if err != nil {
		return
	}
	defer func() {
		err = response.Body.Close()
		if err != nil {
			return
		}
	}()
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

//GetPhoneAttribute Returns the specified profile field
func (r Request) GetPhoneAttribute(shortCodeID int, segmentationFieldID int) (data PhoneAttributeResponse, err error) {
	path := fmt.Sprintf("/v1/ShortCode/%v/SegmentationField/%v", shortCodeID, segmentationFieldID)
	response, err := r.SendRequest("GET", path)
	if err != nil {
		return
	}
	defer func() {
		err = response.Body.Close()
		if err != nil {
			return
		}
	}()
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

//GetPhoneAttributeCollection Returns the collection of profile fields that exist for the company associated with the specified Short Code
func (r Request) GetPhoneAttributeCollection(shortCodeID int) (data PhoneAttributeCollectionResponse, err error) {
	path := fmt.Sprintf("/v1/ShortCode/%v/SegmentationField", shortCodeID)
	response, err := r.SendRequest("GET", path)
	if err != nil {
		return
	}
	defer func() {
		err = response.Body.Close()
		if err != nil {
			return
		}
	}()
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

//GetShortCodeCollection Retrieve a collection of Short Code objects for a Company
func (r Request) GetShortCodeCollection() (data ShortCodeCollectionResponse, err error) {
	r.NewRequest()
	path := "/v1/ShortCode"
	response, err := r.SendRequest("get", path)
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
		err = listrak.HandleErrorResponse(dec)
	default:
		err = listrak.ErrUnhandledStatusCode
	}

	return
}

//GetShortCodeResource Retrieve a single SMS Short Code by Short Code ID
func (r Request) GetShortCodeResource(shortCodeID int) (data ShortCodeResponse, err error) {
	path := fmt.Sprintf("/v1/ShortCode/%v", shortCodeID)
	response, err := r.SendRequest("get", path)
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
		err = listrak.HandleErrorResponse(dec)
	default:
		err = listrak.ErrUnhandledStatusCode
	}

	return
}

//GetTransactionalMessageCollection Retrieve a Collection of Transactional SMS Message for an SMS List
func (r Request) GetTransactionalMessageCollection(shortCodeID int, phoneListID int) (data TransactionalMessageCollectionResponse, err error) {
	path := fmt.Sprintf(
		"/v1/ShortCode/%v/PhoneList/%v/TransactionalMessage",
		shortCodeID, phoneListID,
	)
	response, err := r.SendRequest("get", path)
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
		err = listrak.HandleErrorResponse(dec)
	default:
		err = listrak.ErrUnhandledStatusCode
	}

	return
}

//GetTransactionalMessageResource Retrieve a single Transactional SMS Message for an SMS List by ID
func (r Request) GetTransactionalMessageResource(shortCodeID int, phoneListID int, transactionalMessageID int) (data TransactionalMessageResponse, err error) {
	path := fmt.Sprintf(
		"/v1/ShortCode/%v/PhoneList/%v/TransactionalMessage/%v",
		shortCodeID, phoneListID, transactionalMessageID,
	)
	response, err := r.SendRequest("get", path)
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
		err = listrak.HandleErrorResponse(dec)
	default:
		err = listrak.ErrUnhandledStatusCode
	}

	return
}

//PostTransactionalMessageSend Send a single Transactional SMS Message for an SMS List by ID
func (r Request) PostTransactionalMessageSend(shortCodeID int, phoneListID int, transactionalMessageID int) (data ResourceCreatedResponse, err error) {
	path := fmt.Sprintf(
		"/v1/ShortCode/%v/PhoneList/%v/TransactionalMessage/%v/Message",
		shortCodeID, phoneListID, transactionalMessageID,
	)

	response, err := r.SendRequest("POST", path)
	if err != nil {
		return
	}
	defer func() {
		err = response.Body.Close()
		if err != nil {
			return
		}
	}()
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
