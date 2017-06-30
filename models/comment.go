package models

// Comment represents a comment on a post
type Comment struct {
	BaseModel

	Content string `gorm:"not null" json:"content"`

	Post   Post `json:"post,omitempty"`
	PostID uint `gorm:"not null" json:"post_id"`

	User   User `json:"user,omitempty"`
	UserID uint `gorm:"not null" json:"user_id"`
}
