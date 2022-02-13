package entity

import "time"

type Menu struct {
	Id         string
	Name       string
	Price      int64
	Quantity   int32
	CreateDate time.Time
	DeleteDate time.Time
}