package model

type Users struct {
	Model   `xorm:"extends"`
	OpenId  string `xorm:"varchar(100) notnull unique 'openid'"`
	UnionId string `xorm:"varchar(100) notnull unique 'unionid'"`
	Channel uint8  `xorm:"tinyint(4) notnull unique 'channel'"`
}
