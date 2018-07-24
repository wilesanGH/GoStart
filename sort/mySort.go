package sort

import "fmt"

//插入排序
/*
插入排序（Insertion Sort）的算法描述是一种简单直观的排序算法。
它的工作原理是通过构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入。
插入排序在实现上，通常采用in-place排序（即只需用到O(1)的额外空间的排序），因而在从后向前扫描过程中，
需要反复把已排序元素逐步向后挪位，为最新元素提供插入空间。
 */
func InsertionSort(s []int32)  {
	for i:=1;i< len(s);i++{
		for j := i-1;j>=0&&s[i]<s[j];j--{
			s[j+1], s[j] = s[j],s[j+1]
		}
	}
	fmt.Println(s)

}
//冒泡排序
/**
冒泡排序（Bubble Sort，台湾译为：泡沫排序或气泡排序）是一种简单的排序算法。
它重复地走访过要排序的数列，一次比较两个元素，如果他们的顺序错误就把他们交换过来。
走访数列的工作是重复地进行直到没有再需要交换，也就是说该数列已经排序完成。
这个算法的名字由来是因为越小的元素会经由交换慢慢“浮”到数列的顶端。
 */
func BubbleSort(s []int32){
	for i:=0;i<len(s)-1;i++{
		for j:=1;j<len(s)-i;j++{
			if s[j-1]>s[j]{
				s[j],s[j-1]=s[j-1],s[j]
				//fmt.Println(s)
			}

		}
	}
	fmt.Println(s)

	}

//选择排序
/**
选择排序(Selection sort)是一种简单直观的排序算法。它的工作原理如下。
首先在未排序序列中找到最小元素，存放到排序序列的起始位置，然后，再从剩余未排序元素中继续寻找最小元素，然后放到排序序列末尾。
以此类推，直到所有元素均排序完毕。
 */
func SelectionSort(s []int32){

 	for i:=0;i<len(s);i++{
 		c:=i
 		for j:=i+1;j<len(s);j++{
 			if s[j]<s[c]{
 				//s[j-1],s[j]=s[j],s[j-1]
 				c=j
			}
		}
		if c!=0 {
			s[i],s[c]=s[c],s[i]
			fmt.Println(s)
		}
		c=0
	}
	fmt.Println(s)

}

//快速排序
func QuickSort(s []int32,first,last int)  {


	flag:= first
	left:= first
	right:=last

	if first>=last{
		return
	}

	for first < last{
		for first <last{
			if s[last] >= s[flag]{
				last--
				continue
			}else{
				s[last],s[flag] = s[flag],s[last]
				fmt.Println(s)
				flag=last
				break
			}
		}

		for first < last{
			if s[first] <= s[flag]{
				first ++
				continue
			}else{
				s[first],s[flag]=s[flag],s[first]
				fmt.Println(s)
				flag=first
				break
			}

		}
	}
	QuickSort(s,left,flag-1)

	QuickSort(s,flag+1,right)
}

// 第二种写法
func Quick2Sort(values []int32) {
	if len(values) <= 1 {
		return
	}
	mid, i := values[0], 1
	head, tail := 0, len(values)-1
	for head < tail {
		fmt.Println(values)
		if values[i] > mid {
			values[i], values[tail] = values[tail], values[i]
			tail--
		} else {
			values[i], values[head] = values[head], values[i]
			head++
			i++
		}
	}
	values[head] = mid
	Quick2Sort(values[:head])
	Quick2Sort(values[head+1:])
}