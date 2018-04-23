package model

type ScanRecords struct {
	Model  `xorm:"extends"`
	UserId int    `xorm:"int(11) notnull 'user_id'"`
	Name   string `xorm:"varchar(100) notnull 'name'"`
	Isbn   string `xorm:"varchar(13) notnull 'isbn'"`
	Status int    `xorm:"tinyint(1) notnull 'status'"`
}
