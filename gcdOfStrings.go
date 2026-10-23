package main

import "fmt"

func gcdOfStrings(str1 string,str2 string) string{
	if len(str1)<len(str2){
		return gcdOfStrings(str2,str1)
	}else if len(str1)>=len(str2) && str1[:len(str2)] !=str2{
		return " "
	}else if len(str2)==0{
		return str1
	}else {
		return gcdOfStrings(str1[len(str2):],str2)
	}
}

func main(){
	//test case 1

	str1,str2 :="ABCABC","ABC"
	fmt.Println(gcdOfStrings(str1,str2))

	// test case 2
	str3, str4 :="ABABAB","ABAB"
	fmt.Println(gcdOfStrings(str3,str4))

	//test case 3
	str5,str6 :="LEET", "CODE"
	fmt.Println(gcdOfStrings(str5,str6))
}