package utils

import (
	"database/sql"
	"gorm.io/gorm"
)

// IsEmpty ..
func IsEmpty(value string) bool {
	return value == ""
}

// IsNotEmpty ..
func IsNotEmpty(value string) bool {
	return !IsEmpty(value)
}

func DbError(err error) error {
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}
