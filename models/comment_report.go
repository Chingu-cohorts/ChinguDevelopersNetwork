package models

// CommentReport represents the infractions that a comment may have
type CommentReport struct {
	BaseModel

	Description string `gorm:"not null" json:"description"`
	IsClosed    bool   `gorm:"default:false" json:"is_closed"`

	Comment   Comment `json:"comment,omitempty"`
	CommentID uint    `gorm:"not null" json:"comment_id"`
}
