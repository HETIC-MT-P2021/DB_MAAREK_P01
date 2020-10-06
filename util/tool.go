package util

import (
	"database/sql"
	"reflect"
	"strconv"
)

type NullString sql.NullString

// Scan implements the Scanner interface for NullString
func (ns *NullString) Scan(value interface{}) error {
	var s sql.NullString
	if err := s.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*ns = NullString{s.String, false}
	} else {
		*ns = NullString{s.String, true}
	}

	return nil
}

func ParseStringToUint64(str string) uint64 {
	var err error
	var id uint64
	id, err = strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0
	}
	return id
}
