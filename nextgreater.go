package main

import "fmt"

func nextGreaterElement(nums []int)[]int{
	n:=len(nums)
	ans:=make([]int,n)
	var stack []int
     for i:=n-1;i>=0;i--{
		for len(stack)>0 && stack[len(stack)-1]<=nums[i]{
			stack=stack[:len(stack)-1]
		}
		if len(stack)==0{
			ans[i]=-1
		}else{
			ans[i]=stack[len(stack)-1]
		}
		stack=append(stack,nums[i])
	 }
	 return ans
}

func main(){
	nums :=[]int{4,5,2,10,8}

	result :=nextGreaterElement(nums)

	fmt.Println("input:",nums)
	fmt.Println("output:",result)
}