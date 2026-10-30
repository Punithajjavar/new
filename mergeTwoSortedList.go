package main
import "fmt"

type ListNode struct{
	Val int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode,list2 *ListNode) *ListNode{
	//base case
	if list1==nil || list2==nil{
		if list1==nil{
			return list2
		}
		return list1
	}

	//case 1 
	if list1.Val<=list2.Val{
		list1.Next=mergeTwoLists(list1.Next,list2)
		return list1
	}else{
		//case 2
		list2.Next=mergeTwoLists(list1,list2.Next)
		return list2
	}

}

func main(){
	list1:=&ListNode{Val:1}
	list1.Next=&ListNode{Val:2}
	list1.Next.Next=&ListNode{Val:4}

	list2:=&ListNode{Val:1}
	list2.Next=&ListNode{Val:3}
	list2.Next.Next=&ListNode{Val:4}

	result :=mergeTwoLists(list1,list2)
	
	for result !=nil{
		fmt.Print(result.Val," ")
		result =result.Next
	}
}