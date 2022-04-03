package listrak

import (
	"fmt"
	"strings"
	"time"
)

var dateTimeStringLayout = "2006-01-02T15:04:05"

// DateTime to quickly format dates coming in and out of the api
type DateTime struct {
	time.Time
}

func (t DateTime) MarshalJSON() ([]byte, error) {
	date := t.Time.Format(dateTimeStringLayout)
	date = fmt.Sprintf(`"%s"`, date)
	return []byte(date), nil
}

func (t *DateTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	date, err := time.Parse(dateTimeStringLayout, s)
	if err != nil {
		return err
	}
	t.Time = date
	return
}
