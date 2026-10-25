package main
import "fmt"
func largestRectangleArea(heights [] int )int{
	n:=len(heights)

	if n==0{
		return 0
	}

	right:=make([]int,n)
	left:=make([]int,n)
	var s[] int

	//find the first smaller bar on the right

	for i:=n-1;i>=0;i--{
		for len(s)>0 && heights[s[len(s)-1]]>=heights[i]{
			s=s[:len(s)-1]
		}
		if len(s)==0{
			right[i]=n
		}else{
			right[i]=s[len(s)-1]
		}
		s=append(s,i)
	}

	//clear the stack 
	s=s[:0]
	//find the first smaller bar on the left 

	for i:=0;i<n;i++{
		for len(s)>0 && heights[s[len(s)-1]]>=heights[i]{
			s=s[:len(s)-1]
		}

		if len(s)==0{
			left[i]=-1
		}else{
			left[i]=s[len(s)-1]
		}
		s=append(s,i)
	}

	//calculate the maximum rectangle area

	ans:=0

	for i:=0;i<n;i++{
		width:=right[i]-left[i]-1
		currArea:=heights[i]*width
		if currArea>ans{
			ans=currArea
		}
	}
	return ans

}

func main(){
	heights :=[]int{2,1,5,6,2,3}

	result:=largestRectangleArea(heights)

	fmt.Println("heights",heights)
	fmt.Println("largest reactangle Area:" ,result)
}