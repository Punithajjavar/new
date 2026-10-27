package main
import "fmt"

func superEggDrop(k int,n int)int{
	dp:=make([]int ,k+1)
	moves:=0
	 
	for dp[k]<n{
		moves++
		for eggs:=k;eggs>=1;eggs--{
			dp[eggs]=dp[eggs]+dp[eggs-1]+1

		}
	}
return moves
}

func main(){
	k:=2
	n:=100
	result:=superEggDrop(k,n)
	fmt.Println(result)
}