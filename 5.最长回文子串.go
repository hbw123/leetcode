package main

import (
	"fmt"
	"math"
)

func longestPalindrome(str string)string  {
	var (
		strat,end int
	)
	chars:=[]byte(str)
	for i:=0;i<len(chars);i++{  // 遍历所有点依次作为中心点
		len1:=expandAroundCenter(chars,i,i) // 从一个字符进行扩展
		len2:=expandAroundCenter(chars,i,i+1) // 从两个字符之间进行扩展
		length:=int(math.Max(float64(len1),float64(len2)))
		if length>end-strat{
			strat=i-(length-1)>>1
			end=i+length>>1+1
		}
	}
	return str[strat:end]
}

func expandAroundCenter(chars []byte,left,right int) int {
	for left>=0 && right<len(chars) && chars[left]==chars[right]{
		left--
		right++
	}
	return right-left-1
}


func main()  {
	str:="-121"
	fmt.Println(longestPalindrome(str))
}