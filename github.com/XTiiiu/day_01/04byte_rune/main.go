package main

import "fmt"

func main() {
	s := "hello 你好"
	n := len(s)
	fmt.Println(n)

	for _, c := range s {
		fmt.Printf("%c\n", c)
	}

	s2 := "白萝卜"
	s3 := []rune(s2)        //把字符串强行转化为一个rune切片
	s3[0] = '红'             //rune(int32)
	fmt.Println(string(s3)) //把rune切片强制转化为一个rune切片

	c1 := "红" //string
	c2 := '红' //int32
	fmt.Printf("%T,%T", c1, c2)
	c3 := "H"
	c4 := byte('H')
	fmt.Printf("%T,%T", c3, c4)
}
