package main

func main() {

}

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 && len(nums2) == 0 {
		return nil
	}
	long := 0
	if len(nums1) > len(nums2) {
		long = len(nums1)
	} else {
		long = len(nums2)
	}

	res := make([]int, long)
	return res

}
