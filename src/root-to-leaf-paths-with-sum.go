/**
 * Definition for binary tree
 * type treeNode struct {
 *     left *treeNode
 *     value int
 *     right *treeNode
 * }
 *
 * func treeNode_new(val int) *treeNode{
 *     var node *treeNode = new(treeNode)
 *     node.value=val
 *     node.left=nil
 *     node.right=nil
 *     return node
 * }
 */
/**
 * @input A : Root pointer of the tree
 * @input B : Integer
 *
 * @Output 2D int array.
 */

func sumLeaves(node *treeNode, nodes []int, sum int, givenSum int, paths *[][]int){
    nodes = append(nodes,node.value);
    sum += node.value;
    if node.left != nil { sumLeaves(node.left, nodes, sum, givenSum, paths); }
    if node.right != nil {sumLeaves(node.right, nodes, sum, givenSum, paths);}
    if node.left == nil && node.right == nil{
        if sum == givenSum {
            tmp := make([]int, len(nodes))
            copy(tmp, nodes)
            *paths = append(*paths,tmp);
        }
    }
}

func pathSum(A *treeNode , B int )  ([][]int) {
    var nodes [][]int;
    sumLeaves(A,make([]int,0),0, B, &nodes);
    return nodes;
}
