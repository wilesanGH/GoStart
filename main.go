package main

import (
	"GoStart/tree"
	"GoStart/extend"
	"fmt"
	"GoStart/queue"
	"GoStart/other"
	real2 "GoStart/real"
	"time"
	"GoStart/funcational/adder"
	"GoStart/funcational"
	"strings"
	"GoStart/channel"
)

func main() {
	//var s  =[]int32{2,34,6,32,1,5,7,234,14,44,22}
	//start :=time.Now()



	//elapsed:=time.Since(start)
	//fmt.Print("生成1000个数要的时间：")

	//start = time.Now()
	//sort.InsertionSort(s[:])
	//sort.BubbleSort(s[:])
	//sort.SelectionSort(s[:])
	//sort.QuickSort(s[:],0,len(s)-1)
	//sort.Quick2Sort(s)
	//fmt.Println(other.LengthOfNonRepeatingSubStr("abcabcbb"))
	//elapsed = time.Since(start)
	//fmt.Print("执行时间：",elapsed)


	//fmt.Println(s)


	//testChannel()

	//fmt.Println(f())


	const (
		a = iota   //0
		b          //1
		c          //2
		d = "ha"   //独立值，iota += 1
		e          //"ha"   iota += 1
		f = 100    //iota +=1
		g          //100  iota +=1
		h = iota   //7,恢复计数
		i          //8
	)
	fmt.Println(a,b,c,d,e,f,g,h,i)






}

func f() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func testLongestString(s string) int{
	//other.FindLongestRepeatString(s)
	return other.LengthOfNonRepeatingSubStr(s)
}


func testTree()  {
	var root tree.TreeNode
	root = tree.TreeNode{Value:3}
	root.Left = &tree.TreeNode{}
	root.Rigth = &tree.TreeNode{5,nil,nil}
	root.Rigth.Left = new(tree.TreeNode)
	root.Left.Rigth = tree.CreateNode(8)
	root.Rigth.Left.SetValue(66)

	root.Traverse()
	fmt.Println()
	myRoot:=extend.MyTreeNode{&root}
	myRoot.PostOrder()

	//fmt.Println(root)

}

func testQueue()  {
	q :=queue.Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

}

func testRetriever(){
	var r real2.Retriever
	r = & real2.MockRetriever{UserAgent:"Mozilla/5.0",Timeout:time.Minute}

	//fmt.Println(real2.Download(r))

	fmt.Printf("%T %v\n",r,r)

	fmt.Println(r)

}

func testAdder(){
	adder.PrintAdder()
}

func testFibonacci(){
	f := funcational.Fibonacci()
	funcational.PrintFileContents(f)
}


func testReader(){
	other.PrintFile("README.md")

	other.PrintcFileContents(strings.NewReader("adsf"))
}

func testChannel()  {
	channel.ChanDemo()
}