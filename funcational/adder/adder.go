package adder

import "fmt"

func Adder() func(int) int{
	sum := 0
	return func(v int) int{
		sum += v
		return sum
	}
}


type iAdder func(int) (int,iAdder)

func Adder2(base int) iAdder {
	return func(v int) (int,iAdder){
		return base+v,Adder2(base+v)
	}
}


func PrintAdder(){
	a := Adder2(0)
	for i:=0;i<10;i++{
		var s int
		s,a = a(i)
		fmt.Printf("0+1+...+ %d=%d\n",i,s)
	}
}