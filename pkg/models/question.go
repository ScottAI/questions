package models

import (
	"time"
)

type Question struct {
	//gorm.Model
	Id             int `gorm:"primary_key" json:"id"`
	Title          string
	Body           string `sql:"type:text;" json:"body"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Views          int
	Likes          int
	AnswerCount    int
	AcceptedAnswer bool
	UserID         int `gorm:"size:10"`
	User           User
	Tags           []Tag `gorm:"many2many:taggings;"`
	User_id        int   `sql:"type:integer REFERENCES users(id)"`
}
