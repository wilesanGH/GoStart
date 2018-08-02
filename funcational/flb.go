package funcational

import (
	"fmt"
	"strings"
	"io"
	"bufio"
)

type IntGen func() int

func Fibonacci() IntGen  {
	a, b := 0,1
	return func() int{
		a, b = b,a+b
		return a
	}
}

func (g IntGen) Read(p []byte) (n int,err error){
	next := g()
	if next >10000{
		return 0,io.EOF
	}
	s := fmt.Sprintf("%d\n",next)
	return strings.NewReader(s).Read(p)
}

func PrintFileContents(reader io.Reader){
	scanner := bufio.NewScanner(reader)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}






