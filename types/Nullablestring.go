package types

import (
	"database/sql/driver"
	"errors"
)

// "database/sql"
// "encoding/json"

// type NullableString sql.NullString
//
// func (ns *NullableString) Value() sql.NullString {
// 	s := ns.String
// 	if len(s) == 0 {
// 		return sql.NullString{}
// 	}
// 	return sql.NullString{
// 		String: s,
// 		Valid:  true,
// 	}
// }
//
// func (ns *NullableString) UnmarshalJSON(data []byte) error {
// 	if string(data) == "null" {
// 		ns.Valid = false
// 		ns.String = ""
// 		return nil
// 	}
//
// 	var s string
// 	if err := json.Unmarshal(data, &s); err != nil {
// 		return err
// 	}
//
// 	ns.Valid = true
// 	ns.String = s
// 	return nil
// }

type NullableString string

func (s *NullableString) Scan(value interface{}) error {
	if value == nil {
		*s = ""
		return nil
	}
	strVal, ok := value.(string)
	if !ok {
		return errors.New("Column is not a string")
	}
	*s = NullableString(strVal)
	return nil
}
func (s NullableString) Value() (driver.Value, error) {
	if len(s) == 0 { // if nil or empty string
		return nil, nil
	}
	return string(s), nil
}
