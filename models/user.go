package models

import "time"

// User represents any user of the site
type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	Name     string `gorm:"type:varchar(100); not null" json:"name"`
	Username string `gorm:"type:varchar(30); unique_index; not null" json:"username"`
	Email    string `gorm:"type:varchar(100); unique_index; not null" json:"email"`
	Country  string `gorm:"type:varchar(30); not null" json:"country,omitempty"`

	CohortID uint `json:"cohort_id,omitempty"`
}
