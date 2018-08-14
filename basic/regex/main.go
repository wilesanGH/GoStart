package main

import (
	"regexp"
	"fmt"
)

const text = `my email is wangs_han14@126.com
				mail1@usst.edu
				main3@usst.dee.nn
`




func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9_]+)(@)([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := re.FindAllStringSubmatch(text,-1)
	fmt.Println(match)
}
