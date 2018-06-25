package graph

import "testing"

func TestDFS(t *testing.T) {
	var alGraph = &ALGraph{}
	m := CreateMGraph()
	MatToList(m,alGraph)
	DFS(alGraph,1)
}