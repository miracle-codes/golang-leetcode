package main

import "fmt"

func main() {
	arr := "asdasddd"
	arr2 := "asdasddda"
	fmt.Println(isAnagram(arr, arr2))

}

func isAnagram(s string, t string) bool {
	record := [26]int{}
	record2 := [26]int{}
	for _, value := range s {
		record[value-'a']++
	}
	for _, value := range t {
		record2[value-'a']++
	}

	return record == record2
}
