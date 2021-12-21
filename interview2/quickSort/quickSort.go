package main

import (
	"fmt"
	"sort"
)

func quickSortV1(arr []int, low, hight int) {
	// 当 low = hight  时跳出
	if low >= hight {
		return
	}

	left, right := low, hight
	pivot := arr[left] // 为了简单起见，直接取左边的第一个数

	for left < right {
		// 先从右边开始迭代

		// 右边的数如果比pivot大，那么就应该将他放在右边，继续向左滑动，遇到一个比他小的为止
		for left < right && arr[right] >= pivot {
			right--
		}

		// 小数移动到左边
		if left < right {
			arr[left] = arr[right]
		}

		// 左边的数如果比pivot小，那么就应该将他放在左边，继续向右滑动，遇到一个比他大的为止
		for left < right && arr[left] < pivot {
			left++
		}

		// 大数移动到右边
		if left < right {
			arr[right] = arr[left]
		}

		// left == right ,pivot 即是最终位置
		if left <= right {
			arr[left] = pivot
		}
	}

	//因为 pivot 的最终位置已锁定

	// 继续排序左边部分
	quickSortV1(arr, low, right-1)
	// 继续排序右边部分
	quickSortV1(arr, right+1, hight)
}

func quickSort(arr []int, left, right int) {

	if left < right {
		pivot := arr[left]
		j := left
		for i := left; i < right; i++ {
			if arr[i] < pivot {
				j++
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
		arr[left], arr[j] = arr[j], arr[left]
		fmt.Println(arr, "  ", arr[left], "  ", arr[j])
		quickSort(arr, left, j)
		quickSort(arr, j+1, right)
	}
}

func quicksort2(arr []int, left, right int) {
	if left < right {
		j := left
		pivot := arr[left]
		for i := left; i < right; i++ {
			if arr[i] < pivot {
				j++
				arr[i], arr[j] = arr[j], arr[i]

			}

		}
		arr[left], arr[j] = arr[j], arr[left]
		quicksort2(arr, left, j)
		quicksort2(arr, j+1, right)
	}
}

func main() {
	arr := []int{3, 7, 9, 8, 38, 93, 12, 222, 45, 93, 23, 84, 65, 2}
	//arr:=[]int{10,9,18,7,16,5,4,3,2,1,1}
	//arr := []int{3,2,1,10,11,1}
	//quickSortV1(arr, 0, len(arr)-1)
	//quicksort2(arr,0,len(arr))
	//arr = sort.Ints(arr)
	//fmt.Println(arr)

	ls := sort.IntSlice(arr)
	fmt.Println(ls) //[1 4 5 3 2]
	sort.Ints(ls)
	fmt.Println(ls) //[1 2 3 4 5]

}

type One struct {
	Num  int
	Name string
}
type OneList []*One

func (this OneList) Len() int {
	return len(this)
}
func (this OneList) Less(i, j int) bool {
	return this[i].Num < this[j].Num
}
func (this OneList) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
