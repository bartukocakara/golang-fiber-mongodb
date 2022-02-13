package entity

import "time"

type Customer struct {
	Id          string
	UserId      string
	FirstName   string
	LastName    string
	Email       string
	Password    string
	PhoneNumber string
	Birthday    string
	CreateDate  time.Time
	DeleteDate  time.Time
}
