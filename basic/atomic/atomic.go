package main

import (
	"fmt"
	"time"
	"sync"
)

type atomicInt struct{
	value int
	lock sync.Mutex
}

func (a *atomicInt) increament(){
	fmt.Println("safe increament")
	func(){
		a.lock.Lock()
		defer
			a.lock.Unlock()
		a.value++
	}()
}

func (a *atomicInt) get() int{
	a.lock.Lock()
	defer
		a.lock.Unlock()
	return a.value
}

func main(){
	var a atomicInt
	a.increament()
	go func(){
		a.increament()
	}()
	time.Sleep(time.Microsecond)
	fmt.Println(a)
}
