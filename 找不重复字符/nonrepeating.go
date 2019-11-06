package main

import "fmt"

//此处bytes只能作用于英文字符，并无法判断中文句子中的非重复字符
/*
可以使用rune，相当于其他语言的char
s="moon，我爱你"
for i ,ch := range []rune(s) {
	fmt.Printf("(%d,%c)",i,ch)
}  //使用rune转换字符串
 */

//寻找最长不含有重复字符的子串
func lengthofNonRepeatingSubstr(s string) int {
	//lastOccured := make(map[byte] int)  //字符上一次出现的位置
	lastOccured := make(map[rune] int)  //使用rune可以识别中文字符
	start := 0
	maxLength := 0

	//for i,ch := range []byte(s) {
	for i,ch := range []rune(s) {
		if lastI,ok := lastOccured[ch]; ok && lastI >=start{ //因为lastoccurred不一定存在，所以加上ok而不是直接>=
			start =lastI + 1
		} //若lastOccured[ch] >= start，更新start
		if i-start+1 > maxLength{
			maxLength = i-start+1
		} //更新maxlength
		lastOccured[ch] = i //更新lastoccurred
	}
	return maxLength
}

/* 对于每一个字母X
1、lastOccurred[x]不存在，或者位置小于start则不操作
2、lastOccurred[x]>=start ,则更新start
3、更新lastOccurred[x],更新maxLength
 */
func main()  {
	fmt.Println(lengthofNonRepeatingSubstr("sajhvdsshjddd"))
	fmt.Println(lengthofNonRepeatingSubstr("sassskjsw"))
	fmt.Println(lengthofNonRepeatingSubstr("我是谁是我"))
}