package main

import "fmt"

func convert(s string, numRows int) string {
	chars:=[]byte(s)
	res := make([][]string, numRows)
	for i := 0; i < numRows; i++ {
		res[i] = make([]string, len(chars))
	}

	index:=0
	col:=0
	for true{
		// 向下走
		for row:=0;row<numRows && index<len(chars);row++{
			res[row][col]=string(chars[index])
			index++
		}
		if index>=len(chars){
			break
		}
		// Z型走
		col++
		for row:=numRows-2;row>0 && index<len(chars);row--{
			res[row][col]=string(chars[index])
			index++
			col++
		}
		if index>=len(chars){
			break
		}
	}
	str:=""
	for _,line:=range res{
		for _,l:=range line{
			if l!=""{
				str+=l
			}
		}
	}
	return str
}

func main() {
	s := "A"
	numRows := 1
	fmt.Println(convert(s,numRows))
}