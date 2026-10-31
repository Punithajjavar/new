package main 
import "fmt"

func longestIncreasingSubsequence(arr []int)int{

n:=len(arr)
dp:=make([]int,n)
hash:=make([]int,n)

maxi:=1

lastIndex:=0

//step1 calculate DP & Hash

for i:=0;i<n;i++{
	dp[i]=1
	hash[i]=i

	for prev:=0;prev<i;prev++{
		//compare privious and current element
		if arr[prev]<arr[i] && 1+dp[prev]>dp[i]{
			dp[i]=1+dp[prev]
			hash[i]=prev
		}

	}

		//compare maximum length
		
	if dp[i]>maxi{
		maxi=dp[i]
		lastIndex=i
	}
}
//step 2: Recunstruct the LIS

temp:=[] int {}
temp =append(temp,arr[lastIndex])
for hash[lastIndex]!=lastIndex{
	lastIndex=hash[lastIndex]
	temp=append(temp,arr[lastIndex])
}
//step3 reverse array

for i,j:=0,len(temp)-1;i<j;i,j=i+1,j-1{
	temp[i],temp[j]=temp[j],temp[i]
}

//step 4 print LIS
fmt.Println("LIS :", temp)
return maxi
}

func main(){
	arr:=[]int{5,4,11,1,16,9}
	result:=longestIncreasingSubsequence(arr)
	fmt.Println("length:",result)
}