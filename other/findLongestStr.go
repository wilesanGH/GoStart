package other

import (
	"fmt"
)

//
func LengthOfNonRepeatingSubStr(s string) int{
	lastOccurred := make(map[rune]int)
	start :=0
	maxLength:=0

	for i,ch:=range []rune(s){

		if lastI,ok:=lastOccurred[ch];
		ok&&lastI >=start{
			start = lastI+1
		}
		if i -start +1>maxLength{
			maxLength =i-start+1
		}
		lastOccurred[ch]=i
	}

	return maxLength
}

func LengthOfStrings(){
	s:="我爱中国"
	for _,b:=range []byte(s) {
		fmt.Printf("%X", b)
	}
}

func FindLongestRepeatString(s string)  {
	var suffixArray []string
	for i:=0;i<len(s);i++{
		suffixArray[i] = s[i:]
	}
	//partition(suffixArray,0,len(suffixArray)-1)
	fmt.Println(suffixArray)

}

func partition(suffix_array []string,start,end int)  {
	if end<=start{
		return
	}
	index1,index2:=start,end
	base := suffix_array[start]
	for ;index1<index2 && suffix_array[index2] >= base;index2--{

		suffix_array[index1] = suffix_array[index2]
	}
	for ;index1<index2 && suffix_array[index2] <= base;index1++{
		suffix_array[index1] = suffix_array[index2]
	}
	suffix_array[index1] = base
	partition(suffix_array,start,index1-1)
	partition(suffix_array,index1+1,end)
}

func findCommonString(str1,str2 string)  {

}