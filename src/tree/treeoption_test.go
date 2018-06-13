package tree

import (
	"fmt"
	"testing"
)

func TestAllPath1(t *testing.T) {
	str := []string{"A","(","B","(","D","(",",","G",")",")",",","C","(","E",",","F",")",")"}
	node := CreateBTNode(str)
	AllPath1(node)
}

func TestLevelOrderV2(t *testing.T) {
	str := []string{"A","(","B","(","D","(",",","G",")",")",",","C","(","E",",","F",")",")"}
	node := CreateBTNode(str)
	LevelOrderV2(node)
}

func TestPostOrderV2(t *testing.T) {
	str := []string{"A","(","B","(","D","(",",","G",")",")",",","C","(","E",",","F",")",")"}
	node := CreateBTNode(str)
	fmt.Println("递归中序遍历和非递归中序遍历：")
	fmt.Print("递归：")
	PostOrder(node)
	fmt.Print("非递归：")
	PostOrderV2(node)

}

func TestInOrderV2(t *testing.T) {
	str := []string{"A","(","B","(","D","(",",","G",")",")",",","C","(","E",",","F",")",")"}
	node := CreateBTNode(str)
	fmt.Println("递归中序遍历和非递归中序遍历：")
	fmt.Print("递归：")
	InOrder(node)
	fmt.Print("非递归：")
	InOrderV2(node)

}

/**
 * 对比递归遍历和非递归遍历
 */
func TestPreOrderV2(t *testing.T) {
	str := []string{"A","(","B","(","D","(",",","G",")",")",",","C","(","E",",","F",")",")"}
	node := CreateBTNode(str)
	fmt.Println("递归先序遍历和非递归先序遍历：")
	fmt.Print("递归：")
	PreOrder(node)
	fmt.Print("非递归：")
	PreOrderV2(node)

}

/**
* 输出值为x的节点的所有祖宗节点
*/
func TestAncestor(t *testing.T) {
	str := []string{"A","(","B","(","D","(",",","G",")",")",",","C","(","E",",","F",")",")"}
	node := CreateBTNode(str)
	ancestor(node, "F")
}

func TestLevel(t *testing.T) {
	str := []string{"A","(","B","(","D","(",",","G",")",")",",","C","(","E",",","F",")",")"}
	node := CreateBTNode(str)
	println(Level(node,1,"D"))
}

/**
* 求叶子节点
*/
func TestDispLeaf(t *testing.T) {
	str := []string{"A","(","B","(","D","(",",","G",")",")",",","C","(","E",",","F",")",")"}
	node := CreateBTNode(str)
	DispLeaf(node)
}

/**
 * 以孩子兄弟链作为树的存储结构，编写一个求树高度的递归算法
 */
//func TestTreeHeight(tree *ChildBroTree) int{
//	max := 0
//	var nextTree *ChildBroTree
//	if tree == nil{
//		return 0
//	}else if tree.vp == nil{
//		return 1
//	}else{
//		//拿到孩子节点
//		nextTree = tree.vp
//		for{
//			if nextTree == nil {
//				break
//			}
//			tmpCount := GetTreeHeight(nextTree)
//			if max<tmpCount{
//				max = tmpCount
//			}
//			//这里不能是nextTree = tree.hp，因为这里要求得是所有兄弟中高度最大的，然后赋值给max
//			nextTree = nextTree.hp
//		}
//		return max+1
//	}
//}
//
/**
* 生成一棵二叉树
*/
func TestBTNode(t *testing.T) {
	str := []string{"A","(","B","(","D","(",",","G",")",")",",","C","(","E",",","F",")",")"}
	node := CreateBTNode(str)
	DispBTNode(node)
}

/**
* 获取二叉树中某一个值对应的节点
*/
func TestFindBTNode(t *testing.T) {
	str := []string{"A","(","B","(","D","(",",","G",")",")",",","C","(","E",",","F",")",")"}
	node := CreateBTNode(str)
	fmt.Println("生成的tree如下")
	DispBTNode(node)
	fmt.Println("")
	fmt.Println("查找值为K的node")
	btnode := FiindBTNode(node,"K")
	if btnode != nil {
		fmt.Println("找到的node:"+btnode.data)
	}else{
		fmt.Println("没有找到对应的node")
	}

}

/**
* 求二叉树的深度
*/
func TestGetBTNodeDepth(t *testing.T) {
	str := []string{"A","(","B","(","D","(",",","G",")",")",",","C","(","E",",","F",")",")"}
	node := CreateBTNode(str)
	depth := BtNodeDepth(node)
	fmt.Println("树的深度是%d：",depth)
}

/**
* 二叉树的先根遍历，中序遍历，和后序遍历
*/
func TestProOrder(t *testing.T) {
	str := []string{"A","(","B","(","D","(",",","G",")",")",",","C","(","E",",","F",")",")"}
	node := CreateBTNode(str)
	fmt.Print("树结构：")
	fmt.Println(str)
	fmt.Print("先序遍历结果如下:")
	PreOrder(node)
	fmt.Println("")
	fmt.Print("中序遍历结果如下:")
	InOrder(node)
	fmt.Println("")
	fmt.Print("后序遍历结果如下:")
	PostOrder(node)
}
