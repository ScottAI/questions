package models

import (
	"time"
)

type Answer struct {
	//gorm.Model
	Id               int    `gorm:"primary_key" json:"id"`
	Body             string `sql:"type:text;" json:"body"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Views            int
	Likes            int
	DisLikes         int
	IsAcceptedAnswer bool
	UserID           int `gorm:"size:10"`
	User             User
	QuestionID       int `gorm:"size:10"`
	Question         Question
	Question_id      int `sql:"type:integer REFERENCES questions(id)"`
	User_id          int `sql:"type:integer REFERENCES users(id)"`
}
