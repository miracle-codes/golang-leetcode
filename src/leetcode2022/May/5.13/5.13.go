package main

import "fmt"

func main() {
	s := "We are happy."
	s = "We a r e a s d a s d"
	fmt.Println(replaceSpace(s))
}
func replaceSpace(s string) string {
	b := []byte(s)
	result := make([]byte, 0)
	for i := 0; i < len(b); i++ {
		if b[i] == ' ' {
			result = append(result, []byte("%20")...)
		} else {
			result = append(result, b[i])
		}
	}
	return string(result)

}

func replaceSpace2(s string) string {
	i := 0
	for _, v := range s {
		if string(v) == " " {
			i++
		}
	}
	length := len(s) + 2*i
	//length := 100000
	arr := make([]byte, length)
	for i := 0; i < length; i++ {
		if s[i] == ' ' {
			arr[i] = '%'
			arr[i+1] = '2'
			arr[i+2] = '0'

			fmt.Println("数字1", i, "值", string(s[i]))

		} else {
			arr[i] = byte(s[i])
			fmt.Println("数字2", i, "值", string(s[i]))

		}

	}

	return string(arr)
}
