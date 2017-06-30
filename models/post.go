package models

// Post represents a post of the forum
type Post struct {
	BaseModel

	Title    string `json:"title"`
	Content  string `json:"content"`
	IsPinned bool   `gorm:"default:false" json:"is_pinned"`
	IsLocked bool   `gorm:"default:false" json:"is_closed"`

	User   User `json:"user,omitempty"`
	UserID uint `gorm:"not null" json:"user_id"`

	Comments []Comment `json:"comments,omitempty"`
}
