package main
import "fmt"

func getRow(n int)[]int{
	ans:=make([]int,n+1)
	ans[0]=1
	
	for i:=1;i<=n;i++{
		for j:=i;j>0;j--{
			ans[j]+=ans[j-1]
		}
		}
		return ans
	}
func main(){
	n:=3
	result:=getRow(n)
	fmt.Println(result)
}
