package util

var urlList = make(map[string]bool)
func CheckUrlRe(url string) bool {
	if urlList[url] == true{
		return false
	}
	urlList[url] = true
	return true

}