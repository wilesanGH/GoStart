package goroutine

import (
	"fmt"
	"time"
)

func GoLoop()  {
	for i := 0;i<10;i++{
		go func(){
			for{
				fmt.Printf("Hello from"+ "goroutine %d\n",i)
			}
		}()
	}
	time.Sleep(time.Microsecond)
}
