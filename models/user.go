package models

import "gopkg.in/mgo.v2/bson"

// User represents any user of the site
type User struct {
	ID    bson.ObjectId `json:"id" bson:"_id"`
	Name  string        `json:"name" bson:"name"`
	Email string        `json:"email" bson:"email"`
}
