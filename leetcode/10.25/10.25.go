package main

import "fmt"

func main() {

	fmt.Println(longestCommonSubsequence("oxcpqrsvwf", "shmtulqrypy"))
}
func longestCommonSubsequence(text1 string, text2 string) int {
	short := ""
	long := ""

	if len(text1) > len(text2) {
		short = text2
		long = text1
	} else {
		short = text1
		long = text2
	}
	num := 0
	i := 0
	j := 0
	s := 0
	for ; j < len(short); j++ {
		i = s
		for ; i < len(long); i++ {

			if long[i] == short[j] {

				s = i + 1
				if j < len(short) {
					j++
				}
				num++
			}
		}
	}

	return num
}
