package models

import "time"

// User represents an user of the application
type User struct {
	ID         uint   `gorm:"primary_key" json:"id"`
	Username   string `gorm:"not_null;unique_index" json:"username"`
	Email      string `gorm:"not_null;unique_index" json:"email"`
	FirstName  string `json:"first_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
	About      string `json:"about,omitempty"`
	Experience uint   `json:"experience,omitempty"`

	GithubUsername   string `json:"github_username,omitempty"`
	MediumUsername   string `json:"medium_username,omitempty"`
	TwitterUsername  string `json:"twitter_username,omitempty"`
	LinkedinUsername string `json:"linkedin_username,omitempty"`

	Cohort   Cohort `json:"cohort,omitempty"`
	CohortID uint   `gorm:"not null" json:"cohort_id"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
