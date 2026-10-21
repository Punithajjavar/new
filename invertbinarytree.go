package main
import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode{
	if root==nil{
		return root
	}
//swap left and right children
temp:=root.Left
root.Left=root.Right
root.Right=temp

//Recursively invert children
invertTree(root.Left)
invertTree(root.Right)

return root
}

//print tree in preorder

func printTree(root *TreeNode){
	if root==nil{
		return
	}
	fmt.Print(root.Val," ")
	printTree(root.Left)
	printTree(root.Right)
}
func main(){
	//create the binary tree
	root:=&TreeNode{Val:4}
	root.Left=&TreeNode{Val:2}
	root.Right=&TreeNode{Val:7}

	root.Left.Left=&TreeNode{Val:3}
	root.Right.Right=&TreeNode{Val:9}

	fmt.Print("Before inversion:")
	printTree(root)

	invertTree(root)
	fmt.Print("\nAfter inversion:")
	printTree(root)
	fmt.Println()
}