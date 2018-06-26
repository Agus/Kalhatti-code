package main

import "fmt"
import "os"
import "strconv"

type TreeNode struct {
	Value int
	Left, Right *TreeNode
}

func Node(val int) *TreeNode{
	return &TreeNode{val,nil,nil}
}

// From left to right, bottom up
func populateTree() *TreeNode{
	root := Node(5);
	root.Left = Node(4);
	root.Left.Left = Node(11);
	root.Left.Left.Left = Node(7);
	root.Left.Left.Right = Node(2);
	root.Right = Node(8);
	root.Right.Left = Node(13);
	root.Right.Right = Node(4);
	root.Right.Right.Left = Node(5);
	root.Right.Right.Right = Node(1);
	return root;
}

func sumLeaves(node *TreeNode, nodes []int, sum int, givenSum int, paths *[][]int){
	nodes = append(nodes,node.Value);
	sum += node.Value;
	if node.Left != nil { sumLeaves(node.Left, nodes, sum, givenSum, paths); }
	if node.Right != nil {sumLeaves(node.Right, nodes, sum, givenSum, paths);}
	if node.Left == nil && node.Right == nil{
		if sum == givenSum {
			tmp := make([]int, len(nodes))
			copy(tmp, nodes)
			*paths = append(*paths,tmp);
		}
	}
}

func main() {
	if len(os.Args) != 2 {fmt.Println("Please provide an integer"); return;};
	givenSum, err := strconv.Atoi(os.Args[1]);
	if err != nil {fmt.Println("Please provide an integer")};
	tree := populateTree();
	var nodes [][]int;
	sumLeaves(tree,make([]int,0),0, givenSum, &nodes);
	fmt.Println(nodes);
}
