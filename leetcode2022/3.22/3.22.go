package main

import "fmt"

func qsort(arr []int, start int, end int) {
	if start >= end {
		return
	}

	key := start
	value := arr[start] //记录当前基准值位置
	fmt.Println(arr)
	for n := start + 1; n <= end; n++ {
		// a[n] < arr[start]
		if arr[n] < value {
			arr[key] = arr[n]   //就将a[n]挪至arr[key]所在位置
			arr[n] = arr[key+1] //a[n]空缺了，将arr[key]向后移动一位
			// 理论上现在值为这个 {84, 84, 543, 5, 100, 23}
			key++ //key的位置改变了1位，key++
			//最后在将进位后的arr[key] = 之前保存的value即为
			// value=321 {84, 312, 543, 5, 100, 23}
			//这样完成了替换
		}
	}

	arr[key] = value
	fmt.Println("---------------------------")
	//一轮循环后该数组为[84 5 100 23 312 543]
	//将基准数两边的数进行进行排序 此时 key=3 start=0 左边为 start-key-1
	qsort(arr, start, key-1)
	qsort(arr, key+1, end) //右边为key+1-end
}

func main() {
	//var array = [...]int{312, 84, 543, 5, 100, 23,25,64,1236,435}
	var array = [...]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	qsort(array[:], 0, len(array)-1)
	fmt.Println(array)
}
