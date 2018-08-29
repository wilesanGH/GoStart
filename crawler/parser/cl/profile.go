package cl

import (
	"GoStart/crawler/engine"
	"regexp"
	"strconv"
	"GoStart/crawler/model"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var MarriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
func ParseProfile(contents []byte,m string) engine.ParseResult{

	profile := model.Profile{}
	r := extractString(contents,ageRe)
	age, err := strconv.Atoi(r)
	if err == nil {
		profile.Age = age
	}

	profile.Marriage = extractString(contents,MarriageRe)
	profile.Name = m
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}


func extractString(contents []byte, re *regexp.Regexp) string{
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	}else{
		return ""
	}
}