package main

import(
	"go03/srm"
	"fmt"
	"strconv"
	"strings"
)
//map(无序的键值对的集合)
func Map(){
	var countryMay map[string]string
	countryMay = make(map[string]string)



	countryMay["ZG"] = "北京"
	countryMay["US"] = "美国"

	for country := range countryMay{
		fmt.Println(country,"首都是",countryMay[country])
	}


}

///go数组
func main()  {
	var i int
	var n [5]int
	n[0] = 1
	n[1] = 2
	n[2] = 3
	n[3] = 4
	n[4] = 5
	var name = [...]string{"name","age"}
	fmt.Println(name[0])

	for i = 0 ; i < 5 ; i++ {
		//fmt.Println("name:",name[i])
		fmt.Println("n:",n[i])
	}
	Map()
	srm.Point();

}
//基本数据类型
func Count(){
	var a int8 =127 //int8: -128 - 127
	//如果超出会怎样
	a = a +1
	fmt.Println(a)  //超出不会报错，结果 ：-128

	//如果说a = 129 的话会报错
	/*
	那a = -128 呢？
	a = a-1
	fmt.Println(a)  // a = 127
	*/

	var a1 int8 = 1
	var a2 int = 1
	fmt.Println(a1 == a2)  //会报错，数据类型不一样

	/*
	解决：转换成同类型就可以
	a1 == int8(a2)
	*/


	// uint8 == byte
	var u8 uint8 = 1
	var b1 byte = 1
	fmt.Println(u8 == b1)  //true


	//字符 串清理
	name := "   kk   "
	fmt.Println(strings.Trim(name, " "))

	// string与int转换
	// _ 代表省略
	// 对import使用只会引入包下面的 init初始化可以不引用调用
	sum1 := "1" // mysql , 参数中获取
	num1, _:= strconv.ParseInt(sum1, 10, 64)  // _线不做任何处理 ， 原代码 num1,err = ...
	// if err != nil {
	// 	fmt.Println(err)
	// }
	fmt.Println(num1 + 1)

	fmt.Println(string(100))
	fmt.Println()

}
