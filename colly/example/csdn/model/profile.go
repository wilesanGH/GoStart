package model

type CSDN_BLOG struct {
	Title          string
	Time           string
	Keywords       []string
	ReadCount      string
	CommentCount   string
	CommentContent []string
	CsdnBase       CSDN_BASE
}

type CSDN_BASE struct {
	Id 		string
	Number string
	Url    string
	Body   string
}

type CSDN_DETAIL struct{
	Id string `xorm:"varchar(36) notnull unique pk 'id'"`
	Number string `xorm:"varchar(20) notnull 'number'"`
	Title string `xorm:"varchar(100) notnull 'title'"`
	Url string `xorm:"varchar(100) notnull  'url'"`
	Keywords string `xorm:"varchar(200)   'keywords'"`
	Time string `xorm:"varchar(50) notnull  'time'"`
	ReadCount string `xorm:"varchar(10)   'readCount'"`
	CommentCount string `xorm:"varchar(500)   'commentCount'"`
	Body string `xorm:"longtext   'body'"`
}