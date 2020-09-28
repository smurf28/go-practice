package base

import (
	"fmt"
	"testing"
)

/*
数组的几种定义方式
	var a [3]int                    // 定义长度为3的int型数组, 元素全部为0
	var b = [...]int{1, 2, 3}       // 定义长度为3的int型数组, 元素为 1, 2, 3
	var c = [...]int{2: 3, 1: 2}    // 定义长度为3的int型数组, 元素为 0, 2, 3
	var d = [...]int{1, 2, 4: 5, 6} // 定义长度为6的int型数组, 元素为 1, 2, 0, 0, 5, 6
第一种方式是定义一个数组变量的最基本的方式，数组的长度明确指定，数组中的每个元素都以零值初始化。

第二种方式定义数组，可以在定义的时候顺序指定全部元素的初始化值，数组的长度根据初始化元素的数目自动计算。

第三种方式是以索引的方式来初始化数组的元素，因此元素的初始化值出现顺序比较随意。这种初始化方式和map[int]Type类型的初始化语法类似。数组的长度以出现的最大的索引为准，没有明确初始化的元素依然用0值初始化。

第四种方式是混合了第二种和第三种的初始化方式，前面两个元素采用顺序初始化，第三第四个元素零值初始化，第五个元素通过索引初始化，最后一个元素跟在前面的第五个元素之后采用顺序初始化。
*/

func TestArray(t *testing.T) {

	var a = [...]int{1, 2, 3}    // a 是一个数组
	var b = &a                   // b 是指向数组的指针
	var c = [...]int{2: 3, 1: 2} // 定义长度为3的int型数组, 元素为 0, 2, 3

	fmt.Println(&a[0], a[1]) // 打印数组的前2个元素
	fmt.Println(&b[0], b[1]) // 通过数组指针访问数组元素的方式和数组类似

	for i, v := range b { // 通过数组指针迭代数组的元素
		fmt.Println(i, v)
	}

	for i := range a {
		fmt.Printf("a[%d]: %d\n", i, a[i])
	}
	/*
		用for range方式迭代的性能可能会更好一些，因为这种迭代可以保证不会出现数组越界的情形，每轮迭代对数组元素的访问时可以省去对下标越界的判断。
	*/
	for i, v := range b {
		fmt.Printf("b[%d]: %d\n", i, v)
	}
	/*
		内置函数len可以用于计算数组的长度，cap函数可以用于计算数组的容量。不过对于数组类型来说，len和cap函数返回的结果始终是一样的，都是对应数组类型的长度。
	*/
	for i := 0; i < len(c); i++ {
		fmt.Printf("c[%d]: %d\n", i, c[i])
	}
}

/*
通过times数组实现迭代
其中times对应一个[5][0]int类型的数组，虽然第一维数组有长度，但是数组的元素[0]int大小是0，因此整个数组占用的内存大小依然是0。没有付出额外的内存代价，我们就通过for range方式实现了times次快速迭代。
*/
func TestArrayTimes(T *testing.T) {
	var times [5][0]int
	for range times {
		fmt.Println("hello")
	}
}

/*
 // 字符串数组
	var s1 = [2]string{"hello", "world"}
	var s2 = [...]string{"你好", "世界"}
	var s3 = [...]string{1: "世界", 0: "你好"}

	// 结构体数组
	var line1 [2]image.Point
	var line2 = [...]image.Point{image.Point{X: 0, Y: 0}, image.Point{X: 1, Y: 1}}
	var line3 = [...]image.Point{{0, 0}, {1, 1}}

	// 图像解码器数组
	var decoder1 [2]func(io.Reader) (image.Image, error)
	var decoder2 = [...]func(io.Reader) (image.Image, error){
		png.Decode,
		jpeg.Decode,
	}

	// 接口数组
	var unknown1 [2]interface{}
	var unknown2 = [...]interface{}{123, "你好"}

	// 管道数组
	var chanList = [2]chan int{}

	定义空数组
	var d [0]int       // 定义一个长度为0的数组
  var e = [0]int{}   // 定义一个长度为0的数组
	var f = [...]int{} // 定义一个长度为0的数组

	c1 := make(chan [0]int)
	go func() {
			fmt.Println("c1")
			c1 <- [0]int{}
	}()
	<-c1
*/
func TestArrayType(T *testing.T) {
	// 我们可以用fmt.Printf函数提供的%T或%#v谓词语法来打印数组的类型和详细信息：
	var b = [...]int{1, 2, 3}
	fmt.Printf("b: %T\n", b)  // b: [3]int
	fmt.Printf("b: %#v\n", b) // b: [3]int{1, 2, 3}
}
