package main

func main() {
	nums := []int{110, 20, 3333, 40, 50, 60}
	println(findNumbers(nums))

}

//func findNumbers(nums []int) int {
//	count :=0
//	 len := len(nums)
//	numss := make([]string,len)  ///动态数组写法
//	for i := range nums {
//		numss[i] =strconv.Itoa(nums[i])
//		v:=strings.Count(numss[i],"")-1
//		if v%2==0  {
//			count++
//		}
//
//	}
//
//
//	return count
//}

func findNumbers(nums []int) int {
	n := 0
	for _, v := range nums {
		cur := 0
		for 0 < v {
			v /= 10
			cur++
		}
		if cur%2 == 0 {
			n++
		}
	}
	return n
}
