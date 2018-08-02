package other

import (
	"io"
	"bufio"
	"fmt"
	"os"
)

func PrintFile(filename string)  {
	file,err := os.Open(filename)
	if err!=nil{
		panic(err)
	}
	PrintcFileContents(file)

}

func PrintcFileContents(reader io.Reader)  {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}
