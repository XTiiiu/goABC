package main

import (
	"fmt"
	"strings"
)

func main() {

	//转义
	path := "\"E:\\gocode\\src\\github.com\\XTiiiu\\day_01\\02fmt>02fmt.exe\""
	fmt.Println(path)

	s1 := "I`m OK"
	fmt.Println(s1)

	//反引号  `
	s2 := `你好
	你在干什么
	我在打游戏
	`
	fmt.Println(s2)
	s3 := `E:\gocode\src\github.com\XTiiiu\day_01\03string>03string.exe`
	fmt.Println(s3)

	//字符串拼接
	name := "理想"
	world := "dsb"

	ss := name + world
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s%s", name, world)
	fmt.Printf("%s%s\n", name, world)
	fmt.Println(ss1)

	//字符串分割
	ret := strings.Split(s3, "\\")
	fmt.Println(ret)

	//包含
	fmt.Println(strings.Contains(ss, "理想"))
	fmt.Println(strings.Contains(ss, "理性"))

	//前缀
	fmt.Println(strings.HasPrefix(ss, "理想"))
	//后缀
	fmt.Println(strings.HasSuffix(ss, "理想"))

	//搜索下标
	s4 := "abcdeb"
	fmt.Println(strings.Index(s4, "c"))
	fmt.Println(strings.LastIndex(s4, "b"))

	//拼接
	fmt.Println(strings.Join(ret, "+"))

}
