package main

//给定一个赎金信 (ransom) 字符串和一个杂志(magazine)字符串，判断第一个字符串 ransom 能不能由第二个字符串 magazines 里面的字符构成。如果可以构成，返回 true ；否则返回 false。
//
//(题目说明：为了不暴露赎金信字迹，要从杂志上搜索各个需要的字母，组成单词来表达意思。杂志字符串中的每个字符只能在赎金信字符串中使用一次。)
//
//示例 1：
//
//输入：ransomNote = "a", magazine = "b"
//输出：false
//示例 2：
//
//输入：ransomNote = "aa", magazine = "ab"
//输出：false
//示例 3：
//
//输入：ransomNote = "aa", magazine = "aab"
//输出：true
//
//
//提示：
//
//你可以假设两个字符串均只含有小写字母。

func main() {
	println(canConstruct("asdaad", "asdffaa"))
}

func canConstruc2t(ransomNote string, magazine string) bool {
	num := 0
	for i := 0; i < len(ransomNote); i++ {
		for j := 0; j < len(magazine); j++ {
			if ransomNote[i] == magazine[j] {
				magazine = magazine[:j] + magazine[j+1:]
				num++
				break
			}
		}
	}
	if num == len(ransomNote) {
		return true
	}

	return false
}

///官方
func canConstruct(ransomNote string, magazine string) bool {
	record := make([]int, 26)
	for _, v := range magazine {
		record[v-'a']++
	}
	for _, v := range ransomNote {
		record[v-'a']--
		if record[v-'a'] < 0 {
			return false
		}
	}
	return true
}
