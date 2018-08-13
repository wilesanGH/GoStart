package channel

import (
	"fmt"
	"time"
)

func worker(id int,c chan int)  {
	for{
		fmt.Printf("Worker %d receivec %c \n",id, <-c)
	}

}

func createWorker(id int ) chan<- int{
	c:=make(chan int)
	go func(){
		for{
			fmt.Printf("Worker %d receivec %c \n",id, <-c)
		}
	}()
	return c
}

func bufferedChannel()  {
	c:=make(chan int,4)

	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'e'

	fmt.Println(<-c)

}

func channelClose()  {
	c:= make(chan int)
	go worker(0,c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'e'
	close(c)
	time.Sleep(time.Microsecond)
}

func ChanDemo(){
	bufferedChannel()
	var channels [10] chan<- int
	for i:=0;i<10;i++{
		channels[i] = createWorker(i)
	}

	for i:=0;i<10;i++{
		channels[i] <- 'a'+i
	}

	for i:= 0;i<10;i++{
		channels[i] <- 'A'+i
	}
	time.Sleep(time.Microsecond)

}
