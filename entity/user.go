package entity

import "time"

type User struct {
	Id         string
	Role       string
	CreateDate time.Time
	DeleteDate time.Time
}