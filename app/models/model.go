package model

import "time"

const (
	WX_APP = iota
	WX_Web
)

type Model struct {
	Id        int64     `xorm:"pk autoincr"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	//DeletedAt time.Time `xorm:"deleted"`
}
