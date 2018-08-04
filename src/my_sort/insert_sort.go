package my_sort

import "fmt"

//直接插入排序
//也就是先把要排的主角先拿出来，然后一层一层往后面排序
func InsertSort(data []int) {
	for index, item := range data[1:]{
		for {
			if index < 0 || index > len(data) - 1 || item < data[index] {
				break
			}
			data[index+1] = data[index]
			index--
		}
		data[index+1] = item
	}
}

//折半插入排序
func InsertSort1(data []int){
	for index,item := range data[1:] {
		low := 0
		high := index + 1
		mid := 0
		for {
			if low > high {
				break
			}
			mid = (low + high) / 2
			if data[mid] > item {
				high = mid - 1
			}else{
				low = mid + 1
			}
		}

		j := index
		for {
			if j < mid {
				break
			}
			data[j + 1] = data[j]
			j--
		}
		data[mid] = item
	}
}

//希尔排序
func ShellSort(data []int){
	length := len(data)
	gap := length / 2
	var i,j int
	for {
		if gap <= 0 {
			break
		}
		i = gap
		for {
			if i >= length {
				break
			}
			tmp := data[i]
			j = i - gap
			//进行直接插入排序
			for {
				if j < 0 || data[j] < tmp{
					break
				}
				data[j + gap] = data[j]
				j = j - gap
			}
			data[j + gap] = tmp
			i++
		}
		gap = gap / 2
	}
}

//冒泡排序
func BubbleSort(data []int){
	var i,j int
	i = 0
	for {
		if i >= len(data) {
			break
		}
		j = len(data) - 1
		for {
			if j <= i {
				break
			}
			if data[j] < data[j-1] {
				tmp := data[j]
				data[j] = data[j-1]
				data[j-1] = tmp
			}
			j--
		}
		i++
	}
}

//快速排序
func QuickSort(data []int, s,t int){
	i := s
	j := t

	if s < t {
		tmp := data[s]
		for {
			if i == j {
				break
			}
			for {
				if data[j] < tmp || j <= i {
					break
				}
				j--
			}
			data[i] = data[j]
			for {
				if data[i] > tmp || i >= j {
					break
				}
				i++
			}
			data[j] = data[i]

		}
		data[i] = tmp
		QuickSort(data, s, i-1)
		QuickSort(data, i+1, t)
	}
}

func SelectSort(data []int){
	length := len(data)
	for i:=0;i<length-1;i++{
		m := i
		for j:=i+1;j<length-1;j++{
			if data[j] < data[m] {
				m = j
			}
		}
		if m != i {
			tmp := data[i]
			data[i] = data[m]
			data[m] = tmp

		}
	}
}

func Sift(data []int,low, high int){
	i := low
	j := 2 * i
	tmp := data[i]
	for {
		if j > high {
			break
		}
		if data[j] > data[j + 1] {
			j++
		}
		if tmp < data[j] {
			data[i] = data[j]
			i = j
			j = 2*i
		}else{
			break
		}
	}
	data[i] = tmp
}

func HeapSort(data []int){
	//先构造一个大跟堆
	length := len(data) - 1
	for start:= len(data) / 2 -1 ; start>0;start--{
		Sift(data,start,length)
	}

	for i:=length;i>1;i--{
		tmp := data[i]
		data[i] = data[1]
		data[1] = tmp
		Sift(data,1,i-1)
	}

}

func DisplayArray(data []int) {
	for _, item := range data {
		fmt.Println(fmt.Sprintf("%d", item))
	}
}
