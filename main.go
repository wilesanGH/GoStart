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


	testReader()



	/*fmt.Println(
		testLongestString("abcabcdef"))
	fmt.Println(
		testLongestString("bbbbb"))
	fmt.Println(
		testLongestString("pwwkew"))
	fmt.Println(
		testLongestString(""))
	fmt.Println(
		testLongestString("b"))
	fmt.Println(
		testLongestString("abcdef"))
	fmt.Println(
		testLongestString("一二三二一"))
	fmt.Println(
		testLongestString(
			"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"))*/

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