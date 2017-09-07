package models

// PostReport represents the infractions that a post may have
type PostReport struct {
	BaseModel

	Description string `gorm:"not null" json:"description"`
	IsClosed    bool   `gorm:"default:false" json:"is_closed"`

	Post   Post `json:"post,omitempty"`
	PostID uint `gorm:"not null" json:"post_id"`
}
