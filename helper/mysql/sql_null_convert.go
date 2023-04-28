package mysql

import "database/sql"

func NewNullInt64(i int64) *NullInt {
	if i == 0 {
		return &NullInt{}
	}

	return &NullInt{
		NullInt64: sql.NullInt64{
			Int64: i,
			Valid: true,
		},
	}
}

func NewNullString(value string) *NullString {
	if len(value) == 0 {
		return &NullString{}
	}

	return &NullString{
		NullString: sql.NullString{
			String: value,
			Valid:  true,
		},
	}
}
