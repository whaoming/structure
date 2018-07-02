package graph

import (
	"fmt"
	"strconv"
)

var thisData = []string{"a", "b", "c", "d", "e", "f"}

func CreateMGraph() *MGraph {
	var mGrgph *MGraph
	mGrgph = &MGraph{}
	i := 0
	for {
		if i > 5 {
			break
		}
		var vex VertexType
		vex.no = i
		vex.data = thisData[i]
		mGrgph.vexs[i] = vex
		i++
	}
	mGrgph.edges[0][1] = 1
	mGrgph.edges[1][2] = 1
	mGrgph.edges[2][0] = 1
	mGrgph.edges[2][3] = 1
	mGrgph.edges[3][4] = 1
	mGrgph.edges[3][5] = 1
	mGrgph.edges[0][4] = 1
	mGrgph.edges[4][3] = 1
	mGrgph.edges[0][2] = 1
	mGrgph.n = 6
	mGrgph.e = 9
	return mGrgph
}

/**
 * 邻接矩阵转化为临接表
 */
func MatToList(mGraph *MGraph, alGraph *ALGraph) {
	i := 0
	j := 0
	for {
		if i > mGraph.n-1 {
			break
		}
		j = 0
		vex := mGraph.vexs[i]
		var vex2 VertexType //节点信息
		var vnode VNode     //一个节点
		vex2.no = i
		vex2.data = vex.data
		vnode.data = vex

		for {
			if j > mGraph.n-1 {
				break
			}

			if mGraph.edges[i][j] != 0 {
				var arcnode = &ArcNode{}
				arcnode.adjvex = j
				arcnode.info = 1
				arcnode.nextarc = vnode.firstarc
				vnode.firstarc = arcnode
			}
			j++
		}
		alGraph.n = mGraph.n
		alGraph.e = mGraph.e
		alGraph.adjlist[i] = vnode
		i++
	}
}

func CreateALGraph() *ALGraph {
	var alGraph *ALGraph
	i := 0
	for {
		if i > 5 {
			break
		}
		var vex VertexType //节点信息
		var vnode VNode    //一个节点
		vex.no = i
		vex.data = "aaa"
		vnode.data = vex

		switch i {
		case 0:
			//vnode.firstarc
		}
		alGraph.adjlist[i] = vnode
		i++
	}
	return alGraph
}

var Viste [MAXV]int

/**
 * 深度优先遍历
 */
func DFS(alGraph *ALGraph, v int) {
	vNode := alGraph.adjlist[v]
	print(vNode.data.data)
	Viste[v] = 1
	arcNode := vNode.firstarc
	for {
		if arcNode == nil {
			break
		}
		if Viste[arcNode.adjvex] == 0 {
			DFS(alGraph, arcNode.adjvex)
		}
		arcNode = arcNode.nextarc
	}
}

//广度优先遍历
func BFS(alGraph *ALGraph, v int) {
	queue := make([]int,MAXV)
	front := -1
	rear := -1
	rear++
	queue[rear] = v
	Viste[v] = 1
	fmt.Print(alGraph.adjlist[v].data.data)
	for {
		if rear == front {
			break
		}
		front++
		tmp := queue[front]
		vNode := alGraph.adjlist[tmp]
		nextNode := vNode.firstarc
		for {
			if nextNode == nil {
				break
			}
			if Viste[nextNode.adjvex] == 0 {
				rear++
				queue[rear] = nextNode.adjvex
				fmt.Print(alGraph.adjlist[nextNode.adjvex].data.data)
				Viste[nextNode.adjvex] = 1
			}
			nextNode = nextNode.nextarc
		}
	}
	fmt.Println("")
}

//假设图G采用邻接表存储，设计一个算法输出图G中从顶点u到v的一条简单路径
//d的作用很大 想清楚
func FindPath1(alGraph *ALGraph, u, v, d int, path []int) {
	//fmt.Println("目前检查的点："+strconv.Itoa(u))
	vNode := alGraph.adjlist[u]
	d++
	path[d] = u
	Viste[u] = 1
	if u == v {
		tmp := 0
		for {
			if tmp > len(path)-1 {
				break
			}
			if path[tmp] != -1 {
				fmt.Println(alGraph.adjlist[path[tmp]].data.data)
				if path[tmp] == v {
					return
				}
			}
			tmp++
		}

	}
	arcNode := vNode.firstarc
	for {
		if arcNode == nil {
			break
		}
		if Viste[arcNode.adjvex] == 0 {
			FindPath1(alGraph, arcNode.adjvex, v, d, path)
		}
		arcNode = arcNode.nextarc
	}
}

//假设图G采用邻接表存储，设计一个算法输出图G中从顶点u到v的所有简单路径
func FindPath2(alGraph *ALGraph, u, v, d int, path []int) {
	fmt.Println("目前检查的点：" + strconv.Itoa(u))
	vNode := alGraph.adjlist[u]
	d++
	path[d] = u
	Viste[u] = 1
	if u == v {
		tmp := 0
		for {
			if tmp > len(path)-1 {
				break
			}
			if path[tmp] != -1 {
				//
				fmt.Print(alGraph.adjlist[path[tmp]].data.data)
				if path[tmp] == v {
					break
				}
			}
			tmp++
		}
		fmt.Println("")
	}
	arcNode := vNode.firstarc
	for {
		if arcNode == nil {
			break
		}
		if Viste[arcNode.adjvex] == 0 {
			FindPath2(alGraph, arcNode.adjvex, v, d, path)
		}
		Viste[arcNode.adjvex] = 0
		arcNode = arcNode.nextarc

	}
}

//假设图G采用邻接表存储，设计一个算法输出图G中从顶点u到v的所有长度为L的所有简单路径
func FindPath3(alGraph *ALGraph, u, v, d, l int, path []int) {
	vNode := alGraph.adjlist[u]
	d++
	path[d] = u
	Viste[u] = 1
	if u == v && d == l {
		tmp := 0
		for {
			if tmp > len(path)-1 {
				break
			}
			if path[tmp] != -1 {
				//
				fmt.Print(alGraph.adjlist[path[tmp]].data.data)
				if path[tmp] == v {
					break
				}
			}
			tmp++
		}
		fmt.Println("")
	}
	arcNode := vNode.firstarc
	for {
		if arcNode == nil {
			break
		}
		if Viste[arcNode.adjvex] == 0 {
			FindPath3(alGraph, arcNode.adjvex, v, d, l, path)
		}
		Viste[arcNode.adjvex] = 0
		arcNode = arcNode.nextarc

	}
}

//假设图G采用邻接表存储，设计一个算法求图中通过某顶点k的所有简单回路
func FindPath4(alGraph *ALGraph, u, v, d int, path []int) {
	//fmt.Println("目前检查的点：" + strconv.Itoa(u))
	d++
	path[d] = u
	Viste[u] = 1
	vNode := alGraph.adjlist[u]
	nextNode := vNode.firstarc
	for {
		if nextNode == nil {
			break
		}
		//打印
		if nextNode.adjvex == v && d > 0 {
			tmp := 0
			for {
				if tmp > len(path)-1 {
					break
				}
				if path[tmp] != -1 {
					fmt.Print(alGraph.adjlist[path[tmp]].data.data)
				}
				tmp++
			}
			fmt.Print(alGraph.adjlist[v].data.data)
			fmt.Println("")
		}

		if Viste[nextNode.adjvex] == 0 {
			FindPath4(alGraph, nextNode.adjvex, v, d, path)
		}
		nextNode = nextNode.nextarc
	}
	//fmt.Println("回滚：" + strconv.Itoa(u))
	path[d] = -1
	Viste[u] = 0
}

//假设图G采用邻接表存储，设计一个算法求不带权无向连通图G中从定点u到定点v的一条最短路径
func ShortPath(alGraph *ALGraph, u, v int){
	type QUERE struct {
		data int
		parent int
	}
	rear := -1
	front := -1
	qu := make([]QUERE,alGraph.e)
	Viste[u] = 1
	rear ++
	qu[rear] = QUERE{
		data : u,
		parent : -1,
	}

	for {
		if rear == front {
			break
		}
		front ++
		quItem := qu[front]
		if quItem.data == v {
			i := front
			for{
				if qu[i].parent == -1 {
					break
				}
				fmt.Print(alGraph.adjlist[qu[i].data].data.data)
				i = qu[i].parent
			}
			fmt.Print(alGraph.adjlist[qu[i].data].data.data)
			break
		}
		arcNode := alGraph.adjlist[quItem.data].firstarc

		for{
			if arcNode == nil {
				break
			}
			if Viste[arcNode.adjvex] == 0{
				rear ++
				Viste[arcNode.adjvex] = 1
				qu[rear] = QUERE{
					data : arcNode.adjvex,
					parent : front,
				}
			}
			arcNode = arcNode.nextarc
		}

	}

}