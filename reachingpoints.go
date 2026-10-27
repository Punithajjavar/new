package main
import "fmt"
func reachingPoints(sx int,sy int,tx int,ty int)bool{
	for tx>=sx && ty>=sy{
		if tx==sx && ty==sy{
			return true
		}
		if tx>ty{
			if ty==sy{
				return(tx-sx)%ty==0
			}
			tx%=ty
		}else{
			if tx==sx{
				return(ty-sx)%tx==0
			}
			ty%=tx
	}
}
return false
}

func main(){
	sx,sy:=1,1
	tx,ty:=3,5
	result:=reachingPoints(sx,sy,tx,ty)
	fmt.Println(result)
}