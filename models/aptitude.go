package models

// Aptitude represents the skills of an user
type Aptitude struct {
	BaseModel

	Name  string `gorm:"unique_index;not_null" json:"name"`
	Users []User `gorm:"many2many:user_aptitudes" json:"users,omitempty"`
}
