package listrak

import (
	"encoding/json"
	"fmt"
)

func HandleErrorResponse(dec *json.Decoder) (err error) {
	errResponse := ErrorResponse{}
	if err = dec.Decode(&errResponse); err != nil {
		err = fmt.Errorf("improper error response: %v", err)
	} else {
		err = errResponse.ToError()
	}
	return
}

//ErrorResponse #/definitions/Error
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

//ToError formatting
func (e ErrorResponse) ToError() error {
	err, found := errorMap[e.Error]
	if !found {
		err = fmt.Errorf(e.Message)
	}
	return err
}

//ToError formatting
func (e ErrorResponse) String() error {
	return fmt.Errorf("%v (%v): %v", e.Status, e.Error, e.Message)
}
