package main
import "fmt"

func firstMissingPositiveInteger(nums []int)int{
	n:=len(nums)
	//step1 clean the array
	for i:=0;i<n;i++{
		if nums[i]>n || nums[i]<0{
			nums[i]=n+1
		}
	}

	//step2 initilize the signs

	for i:=0;i<n;i++{
		 num:=nums[i]
		if num<0{
			num=-num
		}
		if num<=n{
			if nums[num-1]>0{
				nums[num-1]=-nums[num-1]
			}
		}
	}
	//find the first missing positive number
	for i:=0;i<n;i++{
		if nums[i]>0{
			return i+1
		}
	}
	return n+1
}

func main(){
	nums:=[]int{3,4,-1,1}
	fmt.Println(firstMissingPositiveInteger(nums))
}