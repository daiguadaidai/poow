package types

import (
	"encoding/json"
	"github.com/go-sql-driver/mysql"
	"time"
)

type NullTime struct {
	mysql.NullTime
}

func NewNullTime(data time.Time, force bool) NullTime {
	return NullTime{mysql.NullTime{data, true}}
}

func (v *NullTime) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Time)
	} else {
		return json.Marshal(nil)
	}
}

func (v *NullTime) UnmarshalJSON(data []byte) error {
	var s *time.Time
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s != nil {
		v.Valid = true
		v.Time = *s
	} else {
		v.Valid = false
	}
	return nil
}
