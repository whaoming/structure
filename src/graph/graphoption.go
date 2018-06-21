package graph

func CreateMGraph() (*MGraph) {
	var mGrgph *MGraph
	i := 0
	for {
		if i > 5 {
			break
		}
		var vex VertexType
		vex.no = i
		vex.data = 1
		mGrgph.vexs[i] = vex
		i++
	}
	mGrgph.edges[0][1] = 1
	mGrgph.edges[1][2] = 1
	mGrgph.edges[2][0] = 1
	mGrgph.edges[2][3] = 1
	mGrgph.edges[3][4] = 1
	mGrgph.edges[3][5] = 1
	mGrgph.n = 6
	mGrgph.e = 6
	return mGrgph
}



/**
 * 邻接矩阵转化为临接表
 */
func MatToList(mGraph MGraph, alGraph *ALGraph){
	i := 0
	j := 0
	for {
		if i>mGraph.n-1 {
			break
		}
		j = 0
		vex := mGraph.vexs[i]
		var vex2 VertexType	//节点信息
		var vnode VNode		//一个节点
		vex2.no = i
		vex2.data = vex.data
		vnode.data = vex

		for{
			if j>mGraph.n-1 {
				break
			}

			if mGraph.edges[i][j] !=0 {
				var arcnode ArcNode
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

func CreateALGraph() (*ALGraph){
	var alGraph *ALGraph
	i := 0
	for {
		if i > 5 {
			break
		}
		var vex VertexType	//节点信息
		var vnode VNode		//一个节点
		vex.no = i
		vex.data = 1
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
func DFS(alGraph ALGraph, v int) {
	vNode :=alGraph.adjlist[v]
	print(vNode.data)
	Viste[v] = 1
	arcNode := vNode.firstarc
	for{
		if arcNode.info == 0 {
			break
		}
		if Viste[arcNode.adjvex] == 0 {
			DFS(alGraph, arcNode.adjvex)
		}
		arcNode = arcNode.nextarc
	}
}
