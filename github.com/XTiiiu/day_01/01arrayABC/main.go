package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

//全局变量
var n int = 100

//遍历字符串
func traversalString() {
	s := "pprof.cn博客"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%c ", s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%c ", r)
	}
	fmt.Println()
}

//修改字符串
func changeString() {
	s1 := "hello"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'H'
	fmt.Println(string(byteS1))

	s2 := "博客"
	runeS2 := []rune(s2)
	runeS2[0] = '狗'
	fmt.Println(string(runeS2))
}

//强制类型转换
func sqrtDemo() {
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

//一维数组初始化
// var arr0 [5]int = [5]int{1, 2, 3}
// var arr1 = [5]int{1, 2, 3, 4, 5}
// var arr2 = [...]int{1, 2, 3, 4, 5, 6}
// var str = [5]string{2: "hello world", 4: "tom"}

// func maohao() {
// 	a := [3]int{1, 2}           // 未初始化元素值为 0。
// 	b := [...]int{1, 2, 3, 4}   // 通过初始化值确定数组长度。
// 	c := [5]int{2: 100, 4: 200} // 使用引号初始化元素。
// 	d := [...]struct {
// 		name string
// 		age  uint8
// 	}{
// 		{"user1", 10}, // 可省略元素类型。
// 		{"user2", 20}, // 别忘了最后一行的逗号。
// 	}
// 	fmt.Println(arr0, arr1, arr2, str)
// 	fmt.Println(a, b, c, d)
// }

//多维数组初始化
var arr0 [5][3]int
var arr1 [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}

func arraysInit() {
	a := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	b := [...][2]int{{1, 1}, {2, 2}, {3, 3}} // 第 2 纬度不能用 "..."。
	fmt.Println(arr0, arr1)
	fmt.Println(a, b)
}

//多维数组遍历
func arrayPrint() {

	var f [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}

	for k1, v1 := range f {
		for k2, v2 := range v1 {
			fmt.Printf("(%d,%d)=%d ", k1, k2, v2)
		}
		fmt.Println()
	}
}

func printArr(arr *[5]int) {
	arr[0] = 10
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func arrayTest() {
	var arr1 [5]int
	printArr(&arr1)
	fmt.Println(arr1)
	arr2 := [...]int{2, 4, 6, 8, 10}
	printArr(&arr2)
	fmt.Println(arr2)
}

func sumArr(x [10]int) int {
	var sum int = 0
	for i := 0; i < len(x); i++ {
		sum += x[i]
	}
	return sum
}

func main() {
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
		PB = 1 << (10 * iota)
	)
	fmt.Println(KB)
	fmt.Println(MB)

	s1 := `第一行
    第二行
    第三行
	`

	fmt.Println(s1)
	// traversalString()
	// changeString()
	// sqrtDemo()
	// maohao()
	// arraysInit()
	// arrayPrint()
	// arrayTest()
	rand.Seed(time.Now().Unix())
	var b [10]int
	for i := 0; i < len(b); i++ {
		b[i] = rand.Intn(1000)
	}
	sum := sumArr(b)
	fmt.Println("sum = %d", sum)
	fmt.Println("hello world...")
}
