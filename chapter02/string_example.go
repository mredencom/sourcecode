package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

// 字符串的操作
func main() {
	//CompareExample("a", "c")
	//ContainsExample("hello world1", "1")
	//CountExample()
	//FieldExample()
	//SplitExample()
	//PrefixSuffixExample()
	//fmt.Println(JoinExample([]string{"name=xxx", "age=xx", "ages=11"}, "&&&&"))
	//RepeatExample()
	//MapExample()
	//ReplaceExample()
	//LowerUpperExample()
	//TitleExample()
	//TrimExample()
	//NewReplacerExample()
	BuilderExample()
}

// 字符串比较
// a<b 返回-1
// a>b 返回1
// a=b 返回1
func CompareExample(a, b string) {
	// Compare 函数，用于比较两个字符串的大小，如果两个字符串相等，返回为 0。如果 a 小于 b ，返回 -1 ，反之返回 1 。
	// 不推荐使用这个函数，直接使用 == != > < >= <= 等一系列运算符更加直观。
	n := strings.Compare(a, b)
	fmt.Println(n)
	fmt.Println(strings.Compare(a, b))
	fmt.Println(strings.Compare(a, a))
	fmt.Println(strings.Compare(b, a))

	fmt.Println(strings.EqualFold("GO", "go"))
	fmt.Println(strings.EqualFold("壹", "一"))
}

// 是否存在某个字符或子串
func ContainsExample(s, substr string) {
	contains := strings.Contains(s, substr)
	fmt.Println("contains：", contains)
	containsAny := strings.ContainsAny(s, substr)
	fmt.Println("containsAny：", containsAny)
	containsRune := strings.ContainsRune(s, 1)
	fmt.Println("containsRune：", containsRune)
	lastIndex := strings.LastIndex(s, substr)
	fmt.Println("lastIndex：", lastIndex)
	indexBytes := strings.IndexByte(s, byte(1))
	fmt.Println("indexBytes：", indexBytes)
}

// 检查字符串出现的次数
func CountExample() {
	s := "hello world"
	sep := "ll"
	n := strings.Count(s, sep)
	fmt.Println(n)
}

// 字符串切割
func FieldExample() {
	strs := strings.Fields("  foo bar  baz   ")
	fmt.Printf("%q", strs)
}

// 字符串切割
func SplitExample() {
	s := "foo,bar,baz"
	fmt.Printf("%q\n", strings.Split(s, ","))
	fmt.Printf("%q\n", strings.SplitAfter(s, ","))
	fmt.Printf("%q\n", strings.SplitN(s, ",", -1))
	// 来自官方的例子
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
}

// prefix and suffix use
func PrefixSuffixExample() {
	fmt.Println(strings.HasPrefix("Gopher", "Go"))
	fmt.Println(strings.HasPrefix("Gopher", "o"))
	fmt.Println(strings.HasPrefix("Gopher", ""))
	fmt.Println(strings.HasSuffix("Amigo", "go"))
	fmt.Println(strings.HasSuffix("Amigo", "Ami"))
	fmt.Println(strings.HasSuffix("Amigo", ""))
}

// 自己实现的Join的方法
func JoinExample(a []string, sep string) string {
	if len(a) == 0 {
		return ""
	}
	if len(a) == 1 {
		return a[0]
	}
	n := len(sep) * (len(a) - 1)
	for i := 0; i < len(a); i++ {
		n += len(a[i])
	}
	b := make([]byte, n)
	bp := copy(b, a[0])
	for _, s := range a[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], s)
	}
	return string(b)
}

// 字符串重复出现的次数
// 在已知的字符串的后面追加新的字符串
func RepeatExample() {
	s := strings.Repeat("ba", 2)
	fmt.Println("ba" + s)
}

// 字符串的替换
func MapExample() {
	mapping := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z': // 大写字母转小写
			return r + 32
		case r >= 'a' && r <= 'z': // 小写字母不处理
			return r
		case unicode.Is(unicode.Han, r): // 汉字换行
			return '\n'
		}
		return -1 // 过滤所有非字母、汉字的字符
	}
	fmt.Println(strings.Map(mapping, "Hello你#￥%……\n（'World\n,好Hello^(&(*界gopher..."))
}

// 字符串替换

func ReplaceExample() {
	// Replace
	// 用 new 替换 s 中的 old，一共替换 n 个。
	// 如果 n < 0，则不限制替换次数，即全部替换
	// ReplaceAll
	// 该函数内部直接调用了函数 Replace(s, old, new , -1)
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
	fmt.Println(strings.ReplaceAll("oink oink oink", "oink", "moo"))
	fmt.Println(strings.ReplaceAll("This is <b>HTML</b>", ">", "---"))
}

// 大小写转换
// 大小写转换包含了 4 个相关函数，ToLower,ToUpper 用于大小写转换。ToLowerSpecial,ToUpperSpecial 可以转换特殊字符的大小写
func LowerUpperExample() {
	fmt.Println(strings.ToLower("HELLO WORLD"))
	fmt.Println(strings.ToLower("Ā Á Ǎ À"))
	fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "壹"))
	fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "HELLO WORLD"))
	fmt.Println(strings.ToLower("Önnek İş"))
	fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "Önnek İş"))

	fmt.Println(strings.ToUpper("hello world"))
	fmt.Println(strings.ToUpper("ā á ǎ à"))
	fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "一"))
	fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "hello world"))
	fmt.Println(strings.ToUpper("örnek iş"))
	fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "örnek iş"))
}

// 标题处理包含 3 个相关函数，其中 Title 会将 s 每个单词的首字母大写，不处理该单词的后续字符。
// ToTitle 将 s 的每个字母大写。ToTitleSpecial 将 s 的每个字母大写，并且会将一些特殊字母转换为其对应的特殊大写字母。
func TitleExample() {
	fmt.Println(strings.Title("hElLo wOrLd"))
	fmt.Println(strings.ToTitle("hElLo wOrLd"))
	fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "hElLo wOrLd"))
	fmt.Println(strings.Title("āáǎà ōóǒò êēéěè"))
	fmt.Println(strings.ToTitle("āáǎà ōóǒò êēéěè"))
	fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "āáǎà ōóǒò êēéěè"))
	fmt.Println(strings.Title("dünyanın ilk borsa yapısı Aizonai kabul edilir"))
	fmt.Println(strings.ToTitle("dünyanın ilk borsa yapısı Aizonai kabul edilir"))
	fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "dünyanın ilk borsa yapısı Aizonai kabul edilir"))
}

// 取出空格字符的
func TrimExample() {
	x := "!!!@@@你好,!@#$ Gophers###$$$"
	fmt.Println(strings.Trim(x, "@#$!%^&*()_+=-"))
	fmt.Println(strings.TrimLeft(x, "@#$!%^&*()_+=-"))
	fmt.Println(strings.TrimRight(x, "@#$!%^&*()_+=-"))
	fmt.Println(strings.TrimSpace(" \t\n Hello, Gophers \n\t\r\n"))
	fmt.Println(strings.TrimPrefix(x, "!"))
	fmt.Println(strings.TrimSuffix(x, "$"))

	f := func(r rune) bool {
		return !unicode.Is(unicode.Han, r) // 非汉字返回 true
	}
	fmt.Println(strings.TrimFunc(x, f))
	fmt.Println(strings.TrimLeftFunc(x, f))
	fmt.Println(strings.TrimRightFunc(x, f))
}

// 这是一个结构，没有导出任何字段，实例化通过 func NewReplacer(oldnew ...string) *Replacer 函数进行，
// 其中不定参数 oldnew 是 old-new 对，即进行多个替换。如果 oldnew 长度与奇数，会导致 panic.
func NewReplacerExample() {
	r := strings.NewReplacer("<", "&lt------;", ">", "&gt----;")
	fmt.Println(r.Replace("This is <b>HTML</b>!"))

	r.WriteString(os.Stdout, "This is <b>HTML</b>!")
}

// Builder类型
func BuilderExample() {
	b := strings.Builder{}
	_ = b.WriteByte('7')
	n, _ := b.WriteRune('夕')
	fmt.Println(n)
	n, _ = b.Write([]byte("Hello, World"))
	fmt.Println(n)
	n, _ = b.WriteString("你好，世界")
	fmt.Println(n)
	fmt.Println(b.Len())
	fmt.Println(b.Cap())
	b.Grow(100)
	fmt.Println(b.Len())
	fmt.Println(b.Cap())
	fmt.Println(b.String())
	b.Reset()
	fmt.Println(b.String())
}
