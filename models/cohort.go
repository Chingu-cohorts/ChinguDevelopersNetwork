package models

import "time"

// Cohort represents any given cohort in the site
type Cohort struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	Name        string `gorm:"type:varchar(40); unique_index; not null" json:"name"`
	Description string `gorm:"type:varchar(255); not null" json:"description"`
	Users       []User `json:"users,omitempty"`
}
