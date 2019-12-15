package main

import (
	// "time"
)

// TODO: add time property test.
type SampleModel struct {
	ID int64 `sql:"type:bigserial" gorm:"primary_key"`
	Name string
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt *time.Time
}