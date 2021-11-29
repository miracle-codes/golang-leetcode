package main

//给你一个字符串 s ，字符串的「能量」定义为：只包含一种字符的最长非空子字符串的长度。
//
//请你返回字符串的能量。
func main() {
	println(maxPower("asdasdasdaaaa"))
}
func maxPower(s string) int {

	b := []byte(s)

	num := len(b)
	if num == 0 {
		return 0
	}
	if num == 1 {
		return 1
	}
	max := 0
	for len(b) != 0 {

		power := 0
		value := b[0]
		for j := 1; j < len(b); j++ {
			if value == b[j] {
				b = append(b[:(j-1)], b[j:]...)
				j--
				power++
			}
		}
		b = append(b[1:])
		power++
		if power > max {
			max = power
		}
	}

	return max
}
