package parser

import (
	"testing"
	"fmt"
	"io/ioutil"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")

	if err != nil{
		panic(err)
	}

	fmt.Printf("%s\n",contents)
}
