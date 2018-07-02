package graph

import "testing"

func TestDFS(t *testing.T) {
	var alGraph = &ALGraph{}
	m := CreateMGraph()
	MatToList(m,alGraph)
	DFS(alGraph,1)
}

func TestBFS(t *testing.T) {
	var alGraph = &ALGraph{}
	m := CreateMGraph()
	MatToList(m,alGraph)
	BFS(alGraph,0)
}

//假设图G采用邻接表存储，设计一个算法输出图G中从顶点u到v的一条简单路径
func TestFindPath1(t *testing.T) {
	var alGraph = &ALGraph{}
	m := CreateMGraph()
	MatToList(m,alGraph)
	//把头结点的俩个结点的位置倒换一下
	//vNode := alGraph.adjlist[0].firstarc
	//alGraph.adjlist[0].firstarc = vNode.nextarc
	//alGraph.adjlist[0].firstarc.nextarc = vNode
	//vNode.nextarc = nil

	path := make([]int,alGraph.e)
	tmp := 0
	for{
		if tmp > alGraph.e -1 {
			break
		}
		path[tmp] = -1
		tmp++
	}
	FindPath1(alGraph,0,5,0, path)
}

//假设图G采用邻接表存储，设计一个算法输出图G中从顶点u到v的所有简单路径
func TestFindPath2(t *testing.T) {
	var alGraph = &ALGraph{}
	m := CreateMGraph()
	MatToList(m,alGraph)
	path := make([]int,alGraph.e)
	tmp := 0
	for{
		if tmp > alGraph.e -1 {
			break
		}
		path[tmp] = -1
		tmp++
	}
	FindPath2(alGraph,0,5,0, path)
}

//假设图G采用邻接表存储，设计一个算法输出图G中从顶点u到v的所有长度为L的所有简单路径
func TestFindPath3(t *testing.T) {
	var alGraph = &ALGraph{}
	m := CreateMGraph()
	MatToList(m,alGraph)
	path := make([]int,alGraph.e)
	tmp := 0
	for{
		if tmp > alGraph.e -1 {
			break
		}
		path[tmp] = -1
		tmp++
	}
	FindPath3(alGraph,0,5,0, 4, path)
}


//假设图G采用邻接表存储，设计一个算法求图中通过某顶点k的所有简单回路
func TestFindPath4(t *testing.T) {
	var alGraph = &ALGraph{}
	m := CreateMGraph()
	MatToList(m,alGraph)
	path := make([]int,alGraph.e)
	tmp := 0
	for{
		if tmp > alGraph.e -1 {
			break
		}
		path[tmp] = -1
		tmp++
	}
	FindPath4(alGraph,2,2,-1, path)
}

func TestShortPath(t *testing.T) {
	var alGraph = &ALGraph{}
	m := CreateMGraph()
	MatToList(m,alGraph)
	path := make([]int,alGraph.e)
	tmp := 0
	for{
		if tmp > alGraph.e -1 {
			break
		}
		path[tmp] = -1
		tmp++
	}
	ShortPath(alGraph,0,5)
}

