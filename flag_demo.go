//package main
//import (
//	"fmt"
//	"os"
//)
//func main() {
//	if len(os.Args) > 0 {
//		for index, arg := range os.Args {
//			fmt.Printf("args[%d]=%v\n", index, arg)
//		}
//	}
//}

package main

import (
	"flag"
	"fmt"
	"time"
)

type User struct {
	Name string
}

func main() {
	//m := make(map[int]User)
	//
	//m[1] = User{Name:"1111"}
	//fmt.Println(m)
	//delete(m, 1)
	//m[1] = User{Name:"222"}
	//fmt.Println(m)

	//定义命令行参数方式1
	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "延迟的时间间隔")

	//解析命令行参数
	flag.Parse()
	fmt.Println(name, age, married, delay)
	//返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	//返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	//返回使用的命令行参数个数
	fmt.Println(flag.NFlag())

}
