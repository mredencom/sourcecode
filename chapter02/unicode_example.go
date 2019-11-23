package main

import (
	"fmt"
	"unicode"
	"unicode/utf16"
	"unicode/utf8"
)

// golang 中的编码 unicode  utf-8  utf-16三种编码，golang是utf-8最好的实现

func main() {
	// unicode rune
	//RuneExample()
	//unicode utf8 rune
	//UTF8RuneExample()
	//unicode utf16 rune
	UTF16RuneExample()
}

// 常用的方法使用
func RuneExample() {
	//IsControl 是否是控制字符
	single := '\u0015'
	fmt.Println(unicode.IsControl(single))
	single = '\ufe45'
	fmt.Println(unicode.IsControl(single))

	digit := '1'
	// 判断是否是阿拉伯数字
	fmt.Println(unicode.IsDigit(digit))
	// 是否数字字符，比如罗马数字Ⅷ也是数字字符
	fmt.Println(unicode.IsNumber(digit))

	letter := 'Ⅷ'
	// 判断是否是阿拉伯数字
	fmt.Println(unicode.IsDigit(letter))
	// 是否数字字符，比如罗马数字Ⅷ也是数字字符
	fmt.Println(unicode.IsNumber(letter))

	han := '你'
	fmt.Println(unicode.IsDigit(han))
	// 判断han是否在unicode编码中的
	fmt.Println(unicode.Is(unicode.Han, han))
	// 判断han是否在Gujaratih和White_Space其中的一个
	fmt.Println(unicode.In(han, unicode.Gujarati, unicode.White_Space))
}

// utf8中的rune
func UTF8RuneExample() {

	word := []byte("界")
	// 校验是否符合utf8编码的函数
	fmt.Println(utf8.Valid(word[:2]))
	fmt.Println(utf8.ValidRune('界'))
	fmt.Println(utf8.ValidString("世界"))

	fmt.Println(utf8.RuneLen('界'))
	// 统计长度
	fmt.Println(utf8.RuneCount(word))
	fmt.Println(utf8.RuneCountInString("世界你好1a"))
	// rune 编码函数
	p := make([]byte, 3)
	utf8.EncodeRune(p, '好')
	fmt.Println(p)
	r, size := utf8.DecodeRune(p)
	fmt.Println(string(r), size)
	fmt.Printf("%c %v\n", r, size)
	fmt.Println(utf8.DecodeRuneInString("你好"))
	fmt.Println(utf8.DecodeLastRune([]byte("你好")))
	fmt.Println(utf8.DecodeLastRuneInString("你好"))
	// 是否为完整的rune
	fmt.Println(utf8.FullRune(word[:2]))
	fmt.Println(utf8.FullRuneInString("你好"))

	fmt.Println(utf8.RuneStart(word[1]))
	fmt.Println(utf8.RuneStart(word[0]))
}

// utf16中的rune
func UTF16RuneExample() {
	words :=[]rune{'𝓐','𝓑'}

	u16:=utf16.Encode(words)
	fmt.Println(u16)
	fmt.Println(utf16.Decode(u16))

	r1,r2:=utf16.EncodeRune('𝓐')
	fmt.Println(r1,r2)
	fmt.Println(utf16.DecodeRune(r1,r2))
	fmt.Println(utf16.IsSurrogate(r1))
	fmt.Println(utf16.IsSurrogate(r2))
	fmt.Println(utf16.IsSurrogate(1234))

}
