package models

import "time"

// Cohort represents a cohort of the site
type Cohort struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	Name        string `gorm:"not_null;unique" json:"name"`
	Description string `gorm:"not_null" json:"description"`

	Users []User `json:"users,omitempty"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
