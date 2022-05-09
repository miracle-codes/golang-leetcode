package main

func main() {

}

func canConstruct(ransomNote string, magazine string) bool {
	var m [26]int
	for _, v := range magazine {
		m[v-'a']++
	}
	for _, v := range ransomNote {
		m[v-'a']--
		if m[v-'a'] < 0 {
			return false
		}
	}
	return true

}
