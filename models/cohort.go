package models

// Cohort represents a cohort of the site
type Cohort struct {
	BaseModel

	Name        string `gorm:"not_null;unique_index" json:"name"`
	Description string `gorm:"not_null" json:"description"`

	Users []User `json:"users,omitempty"`
}
