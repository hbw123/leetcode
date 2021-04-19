package main

type ListNode struct {
    Val int
   Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head,cur *ListNode
	var addFlag,num int
	for true{
		if l1==nil && l2==nil && addFlag==0{
			break
		}else if l1==nil && l2!=nil{
			num=l2.Val+addFlag
			l2=l2.Next
		}else if l1!=nil && l2==nil{
			num=l1.Val+addFlag
			l1=l1.Next
		}else if l1!=nil && l2!=nil{
			num=l1.Val+l2.Val+addFlag
			l1=l1.Next
			l2=l2.Next
		}else {
			num=1
		}
		if num>=10{
			num-=10
			addFlag=1
		}else {
			addFlag=0
		}
		node:=&ListNode{
			Val:  num,
			Next: nil,
		}
		if cur==nil{
			head=node
			cur=head
		}else {
			cur.Next=node
			cur=cur.Next
		}
	}
	return head
}