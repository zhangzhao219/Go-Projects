package chapter02

import "fmt"

var i int = 5
var I int = 6

func main() {

	fmt.Println("df\tfd")
	fmt.Println("df\nfd")
	fmt.Println("df\rfd")
	fmt.Println("df\\fd")
	fmt.Println("df\"fd")
	fmt.Println("abdvgdsf\rfd") // 从当前行的最前面开始输出，覆盖掉前面的内容

	fmt.Println("姓名\t年龄\t籍贯\t住址\njohn\t12\t河北\t北京")
	fmt.Println(i)
	fmt.Println(I)
}
