package models

// PostReport represents the infractions that a post may have
type PostReport struct {
	BaseModel

	Description string `gorm:"not null" json:"description"`

	Post   Post `json:"post,omitempty"`
	PostID uint `gorm:"not null" json:"post_id"`
}
