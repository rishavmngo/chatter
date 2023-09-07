package types

import (
	"database/sql"
	"encoding/json"
)

type NullableString sql.NullString

func (ns *NullableString) Value() sql.NullString {
	s := ns.String
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func (ns *NullableString) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		ns.Valid = false
		ns.String = ""
		return nil
	}

	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	ns.Valid = true
	ns.String = s
	return nil
}
