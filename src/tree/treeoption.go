package tree

import "fmt"

/**
 * 输出从根节点到每个叶子节点的路径的逆
 */
func AllPath1(btnode *BTNode){
	var bts = make([]*BTNode, 50)
	var bt *BTNode
	position := 0
	bts[position] = btnode
	bt = btnode.lchild
	for{
		//这里不能加判断条件if bt== nil，因为最后bt是有值的
		if position == -1{
			break
		}
		//处理左子树
		for {
			if bt == nil {
				break
			}
			position++
			bts[position] = bt
			bt = bt.lchild
		}
		//到这里左子树就处理完了
		var lastBt *BTNode
		lastBt = nil
		flag := false
		for{
			if position == -1 || flag {
				break
			}
			bt = bts[position]
			//像bt.rchild != nil这样会造成死循环，仔细理解一下，所以不能用判空来做这样的操作
			if bt.rchild == lastBt {
				lastBt = bt
				if bt.lchild == nil && bt.rchild == nil {
					tmpPosition := position
					for {
						if tmpPosition < 0{
							break
						}
						fmt.Print(bts[tmpPosition].data)
						tmpPosition--
					}
					fmt.Println("")
				}
				position--
			}else{
				bt = bt.rchild
				flag = true
			}
		}

	}
}

/**
* 层次遍历非递归算法
*/
func LevelOrderV2(btnode *BTNode) {
	var bts = make([]*BTNode, 50)
	rear, front := -1,-1
	rear++
	bts[rear] = btnode
	for{
		if rear == front {
			break
		}
		front++
		item := bts[front]
		fmt.Print(item.data)
		if item.lchild != nil {
			rear++
			bts[rear] = item.lchild
		}

		if item.rchild != nil {
			rear++
			bts[rear] = item.rchild
		}
	}
}

/**
 * 后序遍历非递归算法
 */
func  PostOrderV2(btnode *BTNode) {
	var bts = make([]*BTNode, 50)
	var bt *BTNode
	position := 0
	bts[position] = btnode
	bt = btnode.lchild
	for{
		//这里不能加判断条件if bt== nil，因为最后bt是有值的
		if position == -1{
				break
		}
		//处理左子树
		for {
			if bt == nil {
				break
			}
			position++
			bts[position] = bt
			bt = bt.lchild
		}
		//到这里左子树就处理完了
		var lastBt *BTNode
		lastBt = nil
		flag := false
		for{
			if position == -1 || flag {
				break
			}
			bt = bts[position]
			//像bt.rchild != nil这样会造成死循环，仔细理解一下，所以不能用判空来做这样的操作
			if bt.rchild == lastBt {
				lastBt = bt
				fmt.Print(bt.data)
				position--
			}else{
				bt = bt.rchild
				flag = true
			}
		}

	}
}

/**
 * 中序遍历非递归算法
 */
func InOrderV2(btnode *BTNode) {
	var bts = make([]*BTNode, 50)
	var bt *BTNode
	position := 0
	bts[position] = btnode
	bt = btnode.lchild
	for {
		if position == -1{
			if bt == nil{
				break
			}
		}
		//处理左子树
		for {
			if bt == nil {
				break
			}
			position++
			bts[position] = bt
			bt = bt.lchild
		}
		bt = bts[position]
		position--
		fmt.Print(bt.data)
		bt = bt.rchild
	}
}

/**
 * 先序遍历非递归算法
 */
func PreOrderV2(btnode *BTNode) {
	var bts = make([]*BTNode, 50)
	var bt *BTNode
	position := 0
	bts[position] = btnode
	for {
		if position == -1 {
			break
		}
		bt = bts[position]
		position--
		print(bt.data)
		if bt.rchild != nil {
			position++
			bts[position] = bt.rchild
		}
		if bt.lchild != nil {
			position++
			bts[position] = bt.lchild
		}

	}
	fmt.Println("")
}

/**
 * 输出值为x的节点的所有祖宗节点
 */
func ancestor(btnode *BTNode, x string) bool {
	if btnode == nil {
		return false
	} else if (btnode.lchild != nil && btnode.lchild.data == x) || (btnode.rchild != nil && btnode.rchild.data == x) {
		fmt.Print(btnode.data)
		return true
	} else if ancestor(btnode.lchild, x) || ancestor(btnode.rchild, x) {
		fmt.Print(btnode.data)
		return true
	}
	return false
}

func Level(btnode *BTNode, h int, str string) int {
	if btnode == nil {
		return 0
	} else if btnode.data == str {
		return h
	} else {
		tmp := Level(btnode.lchild, h+1, str)
		if tmp == 0 {
			return Level(btnode.rchild, h+1, str)
		} else {
			return tmp
		}
	}
}

/**
 * 求一棵二叉树的叶子节点
 */
func DispLeaf(btnode *BTNode) {
	if btnode.lchild == nil && btnode.rchild == nil {
		fmt.Print(btnode.data)
	} else {
		if btnode.lchild != nil {
			DispLeaf(btnode.lchild)
		}
		if btnode.rchild != nil {
			DispLeaf(btnode.rchild)
		}
	}
}

/**
 * 后序遍历二叉树
 */
func PostOrder(btnode *BTNode) {

	if btnode.lchild != nil {
		PostOrder(btnode.lchild)
	}

	if btnode.rchild != nil {
		PostOrder(btnode.rchild)
	}
	if btnode != nil {
		fmt.Print(btnode.data)
	}
}

/**
 * 中序遍历二叉树
 */
func InOrder(btnode *BTNode) {

	if btnode != nil {
		InOrder(btnode.lchild)
		fmt.Print(btnode.data)
		InOrder(btnode.rchild)
	}



	//if btnode.lchild != nil {
	//	InOrder(btnode.lchild)
	//}
	//if btnode != nil {
	//	fmt.Print(btnode.data)
	//}
	//if btnode.rchild != nil {
	//	InOrder(btnode.rchild)
	//}
}

/**
 * 先序遍历二叉树
 */
func PreOrder(btnode *BTNode) {
	if btnode != nil {
		fmt.Print(btnode.data)
	}
	if btnode.lchild != nil {
		PreOrder(btnode.lchild)
	}

	if btnode.rchild != nil {
		PreOrder(btnode.rchild)
	}

}

/**
 * 求二叉树的深度
 */
func BtNodeDepth(btnode *BTNode) int {
	if btnode == nil {
		return 0
	} else if btnode.lchild == nil && btnode.rchild == nil {
		return 1
	} else {
		tmp := BtNodeDepth(btnode.lchild)
		tmp1 := BtNodeDepth(btnode.rchild)
		if tmp > tmp1 {
			return tmp + 1
		} else {
			return tmp1 + 1
		}
	}

}

/**
 * 查找二叉树的某个节点
 */
func FiindBTNode(btnode *BTNode, str string) *BTNode {
	if btnode == nil {
		return nil
	} else if btnode.data == str {
		return btnode
	} else {
		var b *BTNode
		b = FiindBTNode(btnode.lchild, str)
		if b != nil {
			return b
		} else {
			return FiindBTNode(btnode.rchild, str)
		}
	}
}

/**
 * 创建一颗二叉树
 */
func CreateBTNode(str []string) *BTNode {
	btStack := make([]*BTNode, 20)
	strPosi, positin, state := 0, -1, 0
	var st *BTNode
	for {
		//栈为空，退出
		if strPosi == len(str)-1 {
			break
		}
		item := str[strPosi]

		switch item {
		case "(":
			//遇到左括号，把上一个节点进栈
			state = 1
			positin++
			btStack[positin] = st
			break
		case ")":
			positin--
			break
		case ",":
			state = 2
			break
		default:
			if state == 0 {
				st = new(BTNode)
				st.data = item
			} else if state == 1 {
				//上一个char是左括号
				st = new(BTNode)
				st.data = item
				if btStack[positin] != nil {
					btStack[positin].lchild = st
				} else {
					fmt.Println("btStack[positin]为空")
				}

			} else if state == 2 {
				//上一个char是逗号
				st = new(BTNode)
				st.data = item
				if btStack[positin] != nil {
					btStack[positin].rchild = st
				} else {
					fmt.Println("btStack[positin]为空")
				}

			}
			state = 0
			break
		}
		strPosi++
	}
	return btStack[0]
}

/**
 * 输出一颗二叉树
 */
func DispBTNode(btnode *BTNode) {
	var node = btnode

	if node != nil {
		fmt.Print(node.data)
	}

	if node.lchild != nil {
		DispBTNode(node.lchild)
	}

	if node.rchild != nil {
		DispBTNode(node.rchild)
	}
}

/**
* 获取一颗孩子兄弟链树
*/
func GetChildBroTree() *ChildBroTree {
	a := new(ChildBroTree)
	a.data = "A"
	b := new(ChildBroTree)
	b.data = "B"
	c := new(ChildBroTree)
	c.data = "C"
	d := new(ChildBroTree)
	d.data = "D"
	e := new(ChildBroTree)
	e.data = "E"
	f := new(ChildBroTree)
	f.data = "F"
	g := new(ChildBroTree)
	g.data = "G"
	a.vp = b
	b.vp = d
	b.hp = c
	d.hp = e
	e.vp = g
	e.hp = f
	return a
}
