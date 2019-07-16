package model

type WeiBo struct {
	Id string `xorm:"varchar(36) notnull unique pk 'id'"`
	Url string `xorm:"varchar(36) notnull  'url'"`
	User string `xorm:"varchar(100) notnull  'user'"`
	Content string `xorm:"longtext  notnull 'content'"`
	CommentCount string `xorm:"varchar(100) notnull  'commentCount'"`
	RedirectCount string `xorm:"varchar(100) notnull  'redirectCount'"`
	AgreeCount string `xorm:"varchar(100) notnull  'agreeCount'"`

}
