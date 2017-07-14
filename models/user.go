package models

// User represents an user of the application
type User struct {
	BaseModel

	Username          string `gorm:"not_null;unique_index" json:"username"`
	Email             string `gorm:"not_null;unique_index" json:"email"`
	EncryptedPassword string `gorm:"not_null" json:"-"`
	Password          string `sql:"-" json:"password,omitempty"`
	FirstName         string `json:"first_name,omitempty"`
	LastName          string `json:"last_name,omitempty"`
	About             string `json:"about,omitempty"`
	Experience        uint   `gorm:"default:0" json:"experience"`

	GithubUsername   string `json:"github_username,omitempty"`
	MediumUsername   string `json:"medium_username,omitempty"`
	TwitterUsername  string `json:"twitter_username,omitempty"`
	LinkedinUsername string `json:"linkedin_username,omitempty"`

	Cohort   Cohort `json:"cohort,omitempty"`
	CohortID uint   `json:"cohort_id"`

	Projects []Project `gorm:"many2many:user_projects" json:"projects,omitempty"`

	Posts []Post `json:"posts,omitempty"`

	Comments []Comment `json:"comments,omitempty"`
}
