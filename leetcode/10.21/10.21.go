package main

func main() {
	num := []int{1, 1, 1, 3, 5}
	println(countQuadruplets(num))
}

func countQuadruplets(nums []int) int {

	num := 0
	for y := 3; y < len(nums); y++ {
		for h := 2; h < y; h++ {
			for j := 1; j < h; j++ {
				for i := 0; i < j; i++ {
					if nums[i]+nums[j]+nums[h] == nums[y] {
						num++

					}
				}
			}
		}
	}

	return num
}
