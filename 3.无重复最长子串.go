package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	chars:=[]byte(s)
	if len(chars)==1{
		return 1
	}
	var start,end int
	for i:=range chars{
		dict:=make(map[byte]bool)
		t:=i
		for j:=i;j<len(chars);j++{
			if _,ok:=dict[chars[j]];ok{
				t=j
				break
			}
			dict[chars[j]]=true
			t++
		}
		if  t-i>end-start{
			start=i
			end=t
		}
		if t==len(chars){
			break
		}
	}
	return end-start
}

func main()  {
	s:="aa"
	fmt.Println(lengthOfLongestSubstring(s))
}