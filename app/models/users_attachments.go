package model

type UsersAttachments struct {
	Model      `xorm:"extends"`
	UserId     int64  `xorm:"int(4) notnull 'user_id'"`
	AvatarUrl  string `xorm:"varchar(200) notnull 'avatar_url'"`
	NickName   string `xorm:"varchar(100) notnull 'nickname'"`
	Gender     int    `xorm:"tinyint(4) notnull 'gender'"`
	IsVIP      int    `xorm:"int(11) notnull 'is_vip'"`
	SessionKey string `xorm:"varchar(200) notnull 'session_key'"`
}
