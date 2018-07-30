package my_sort

import (
	"testing"
)

func TestInsertSort(t *testing.T) {
	a := []int{20,9,5,6,3,2,8,7,1,10,56,88}
	InsertSort(a)
	DisplayArray(a)
}

func TestInsertSort1(t *testing.T) {
	a := []int{20,9,5,6,3,2,8,7,1,10,56,88}
	InsertSort1(a)
	DisplayArray(a)
}

func TestShellSort(t *testing.T) {
	a := []int{20,9,5,6,3,2,8,7,1,10,56,88}
	ShellSort(a)
	DisplayArray(a)
}

func TestBubbleSort(t *testing.T) {
	a := []int{20,9,5,6,3,2,8,7,1,10,56,88}
	BubbleSort(a)
	DisplayArray(a)
}
