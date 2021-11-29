package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "zhangsan lisi wanger a b c d"
	fmt.Println(toGoatLatin(sentence))

}

func toGoatLatin(sentence string) string {
	context := strings.Fields(sentence)

	s := string('a')
	translation := ""
	for key, value := range context {
		if value[0] == 'a' || value[0] == 'e' || value[0] == 'i' || value[0] == 'o' || value[0] == 'u' ||
			value[0] == 'A' || value[0] == 'E' || value[0] == 'I' || value[0] == 'O' || value[0] == 'U' {
			if translation == "" {
				translation = context[key] + "ma" + s
			} else {
				translation = translation + " " + context[key] + "ma" + s
			}
		} else {
			frist := string([]byte(value)[:1])
			if translation == "" {
				translation = string([]byte(value)[1:]) + frist + "ma" + s
			} else {

				translation = translation + " " + string([]byte(value)[1:]) + frist + "ma" + s
			}

		}
		s = s + "a"

	}

	return translation
}
