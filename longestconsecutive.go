package main 
import "fmt"

func longestConsecutive(nums [] int)int{
	set:=make(map[int]bool)

	//step 1:add all numbers to the set

	for _,num:=range nums{
		set[num]=true
	}
	count :=0

	//step 2 : check each number

	for _,num:=range nums{
		//start only if previous number is absent
		if !set[num-1]{
			dummy :=num
			streak :=1

			// step 3 : find consecutive numbers

			for set[dummy+1]{
				dummy++
				streak++
			}
			//step 4:update maximum length

			if streak>count{
				count = streak
			}
		}
	}
	return count
}

func main(){
	nums:=[]int{100,4,200,1,3,2}
	result:=longestConsecutive(nums)
	fmt.Println(result)
}