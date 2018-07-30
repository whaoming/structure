package main

import (
	"my_sort"
)

func main()  {
	//myTree := tree.GetChildBroTree()
	//result := tree.GetTreeHeight(myTree)
	//
	//println("获得树的高度是：%d",result)
	//
	//tree.TestOrder()
	a := []int{20,9,5,6,3,2,8,7,1,10,56,88}
	my_sort.InsertSort(a)
}
