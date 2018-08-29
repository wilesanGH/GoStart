package engine

type Request struct {
	Url string
	ParserFunc func([]byte) ParseResult
	NeedVPN bool
}

type ParseResult struct{
	Requests []Request
	Items []interface{}
}

func NilParser([]byte) ParseResult{
	return ParseResult{}
}