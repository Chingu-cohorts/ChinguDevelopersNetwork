package models

import "gopkg.in/mgo.v2/bson"

// Cohort represents any cohort of the site
type Cohort struct {
	ID          bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	CreatedAt   *time.Time    `json:"created_at" bson:"created_at"`
}
