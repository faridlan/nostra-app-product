package mysql

import (
	"database/sql"
	"encoding/json"
)

type SqlNull interface {
	MarshalJSON() ([]byte, error)
}

//Int

type NullInt struct {
	sql.NullInt64
}

func (v *NullInt) MarshalJSON() ([]byte, error) {
	if !v.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(v.Int64)
}

//String

type NullString struct {
	sql.NullString
}

func (v *NullString) MarshalJSON() ([]byte, error) {
	if !v.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(v.String)
}
