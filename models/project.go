package models

// Project represents a project created by the users
type Project struct {
	BaseModel

	Name        string `gorm:"not_null" json:"name"`
	Description string `gorm:"not_null" json:"description"`

	Users []User `gorm:"many2many:user_projects" json:"users,omitempty"`
}
