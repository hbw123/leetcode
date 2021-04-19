package main

import "fmt"

func twoSum(nums []int, target int) []int {
	dict:=make(map[int][]int)
	for i,num:=range nums{
		dict[num] = append(dict[num], i)
	}
	var res []int
	for key,val:=range dict{
		if v,ok:=dict[target-key];ok{
			if key==target-key {
				if len(v)==1{
					continue
				}else {
					res = append(res, val...)
				}
			}else{
				res = append(res, val...)
				res = append(res, v...)
			}
		break
		}
	}
	return res
}

func main()  {
	nums := []int{3,3}
	target := 6
	fmt.Println(twoSum(nums,target))
}