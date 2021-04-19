package main

import (
	"fmt"
	"math"
	"strconv"
)

func ReverseString(str string,from,to int) string {
	chars:=[]byte(str)
	for from<to{
		temp:=chars[from]
		chars[from]=chars[to]
		chars[to]=temp
		from++
		to--
	}
	return string(chars)
}

func reverse(x int) int {
	str:=strconv.Itoa(x)
	str=ReverseString(str,0,len(str)-1)
	if x<0{
		str="-"+str[0:len(str)-1]
	}
	i,err:=strconv.Atoi(str)
	if err!=nil || i < int(-math.Pow(2,31)) || i> int(math.Pow(2,31)){
		return 0
	}
	return i
}

func main()  {
	fmt.Println(reverse(-123))
}