package tree

type ElemType string

/**
 * 二叉链
 */
type BTNode struct {
	data string
	lchild *BTNode
	rchild *BTNode
}

/*
 * 孩子兄弟链储存结构
 */
type ChildBroTree struct {
	data ElemType
	hp *ChildBroTree
	vp *ChildBroTree 	//孩子节点
}

