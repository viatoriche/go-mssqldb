package mssql

import (
	"database/sql/driver"
)

type NullDateTime1 struct {
	Time  DateTime1
	Valid bool // Valid is true if String is not NULL
}

// Scan implements the Scanner interface.
func (nd *NullDateTime1) Scan(value interface{}) error {
	if value == nil {
		nd.Time, nd.Valid = DateTime1{}, false
		return nil
	}
	nd.Valid = true
	return convertAssign(&nd.Time, value)
}

// Value implements the driver Valuer interface.
func (nd NullDateTime1) Value() (driver.Value, error) {
	if !nd.Valid {
		return nil, nil
	}
	return nd.Time, nil
}
