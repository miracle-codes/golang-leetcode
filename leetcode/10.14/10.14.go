package main

//给定一个整数，写一个函数来判断它是否是 3 的幂次方。如果是，返回 true ；否则，返回 false 。
//
//整数 n 是 3 的幂次方需满足：存在整数 x 使得 n == 3x
//

func main() {
	println(division(27))
	//println(5%3)

}

func division(n int32) (m bool) {
	if n == 0 {
		return false
	}

	if n == 1 {
		return true
	}

	if n%3 == int32(0) {
		n = n / 3
		return division(n)
	} else {
		return false
	}

}

//官方解答
//func isPowerOfThree(n int) bool {
//    for n > 0 && n%3 == 0 {
//        n /= 3
//    }
//    return n == 1
//}
//
//
