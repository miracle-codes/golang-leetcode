package main

////返回字符串中出现频率最多的次数
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
