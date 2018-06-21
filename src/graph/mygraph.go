package graph

type ElemType int

const MAXV  = 50

type VertexType struct {
	no int
	data ElemType
}

/**
 * 邻接矩阵
 */
type MGraph struct {
	edges [MAXV][MAXV]int
	vexs [MAXV]VertexType
	n,e int
}

type ArcNode struct {
	adjvex int
	info ElemType
	nextarc ArcNode
}

type VNode struct {
	data VertexType
	firstarc ArcNode
}

type ALGraph struct {
	adjlist []VNode
	n,e int
}