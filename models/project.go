package models

import "time"

// Project represents a project created by the users
type Project struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	Name        string `gorm:"not_null" json:"name"`
	Description string `gorm:"not_null" json:"description"`

	Users []User `gorm:"many2many:user_projects"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
