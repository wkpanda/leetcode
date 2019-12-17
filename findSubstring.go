// 30.串联所有单词的子串

// 给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。

// 注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。

//  

// 示例 1：

// 输入：
//   s = "barfoothefoobarman",
//   words = ["foo","bar"]
// 输出：[0,9]
// 解释：
// 从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
// 输出的顺序不重要, [9,0] 也是有效答案。
// 示例 2：

// 输入：
//   s = "wordgoodgoodgoodbestword",
//   words = ["word","good","best","word"]
// 输出：[]

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


func findSubstring(s string, words []string)(result []int) {    
	if len(words)==0 || len(s) == 0{
		return
	}
	wordLen := len(words[0])
	wordsNum := len(words)

	//将ｗords放入Ｍａｐ中
    wordsMap := map[string]int{}
	for _,word := range words {
		wordsMap[word]++
	}

	length := wordLen*wordsNum
	
	//使用ｓｔａｒｔ开头长度为ｌｅｎ的窗口滑动
	for start:=0;start < len(s)-length+1; start++{
		tmpMap := map[string]int {}
		j:=0
		for ;j<length;j+=wordLen{
			tmpS := s[start+j:start+j+wordLen]
			if wordsMap[tmpS]!=0 {
				tmpMap[tmpS]++
				if tmpMap[tmpS] > wordsMap[tmpS]{
					break
				}
			} else {
				break
			}

		}
		if j==length{
			result = append(result,start)
		}
	}
	return
}
