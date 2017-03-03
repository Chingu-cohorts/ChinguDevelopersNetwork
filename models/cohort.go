package models

import "gopkg.in/mgo.v2/bson"

// Cohort represents any cohort of the site
type Cohort struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
}
