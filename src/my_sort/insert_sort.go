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

func DisplayArray(data []int) {
	for _, item := range data {
		fmt.Println(fmt.Sprintf("%d", item))
	}
}
