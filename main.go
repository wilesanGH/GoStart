package main

import (
	"fmt"
	"time"
	"algorithmStd001/sort"
)

func main() {
	var s  =[]int32{2,34,6,32,1,5,7,234,14,44,22}
	start :=time.Now()

	fmt.Println(s)

	elapsed:=time.Since(start)
	fmt.Print("生成1000个数要的时间：")
	fmt.Println(elapsed)
	start = time.Now()
	//sort.InsertionSort(s[:])
	//sort.BubbleSort(s[:])
	//sort.SelectionSort(s[:])
	//sort.QuickSort(s[:],0,len(s)-1)
	sort.Quick2Sort(s)
	elapsed = time.Since(start)
	fmt.Print("排序1000个数要的时间：")
	fmt.Println(elapsed)

	fmt.Println(s)

}
