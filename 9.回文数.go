package main

import "strconv"

// 解题思路 https://www.hbwai.com/index.php/%E7%AE%97%E6%B3%95/3.html
func isPalindrome1(x int) bool {
	chars:=[]byte(strconv.Itoa(x))
	head:=0
	tail:=len(chars)-1
	for head<tail{
		if chars[head]!=chars[tail]{
			return false
		}
		head++
		tail--
	}
	return true
}

func isPalindrome2(x int) bool  {
	chars:=[]byte(strconv.Itoa(x))
	n:=len(chars)
	m:=n>>1 // 取中间值
	first:=m-1
	second:=n-m
	for first>=0{
		if chars[first]!=chars[second]{
			return false
		}
		first--
		second++
	}
	return true
}