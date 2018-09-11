package util

import (
	"strings"
	"fmt"
)

var urlList = make(map[string]bool)
var subUrlList = make(map[string]bool)

func CheckUrlReapt(url string) bool {
	if urlList[url] == true {
		return false
	}
	urlList[url] = true
	return true

}

func CheckSubUrlReapt(url string) bool {
	if subUrlList[url] == true {
		return false
	}
	subUrlList[url] = true
	return true

}

func DeleteMoreTab(s string) string {
	flag := false
	var s2 []byte

	for _, v := range []byte(s) {
		if v == 9 && flag {
			continue
		}
		if v== 9 {
			flag = true
		}else{
			flag = false
		}
		s2 = append(s2,v)

	}
	s3 := string(s2)
	fmt.Println(s3)
	return s3
}

//删除字符串中的「空格」「换行」「tab键」
func DeleteSpaceNTab(str string) string {
	//str = strings.Replace(str, " ", "", -1)
	// 去除换行符
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, "\t", "", -1)

	return str
}

func DeleteSpaceNTabForSlice(s []string) []string {
	ss := make([]string, 0)
	for _, sub := range s {
		if StringIsSpace(sub) {
			continue
		}
		ss = append(ss, DeleteSpaceNTab(sub))
	}
	return ss
}

/**

 */
func StringIsSpace(str string) bool {
	switch str {
	case "\t", "\n", "\v", "\f", "\r", " ", "":
		return true
	}
	return false
}
