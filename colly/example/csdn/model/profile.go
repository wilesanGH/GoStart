package model

type CSDN_BLOG struct {
	Title string
	Date string
	Keywords []string
	ReadCount int
	CommentCount int
	CommentContent []string
	CsdnBase CSDN_BASE

}

type CSDN_BASE struct {
	Id string
	Url string
	Body string
}
