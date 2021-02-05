package main

import (
	"01/srm"
	"fmt"
)

func main(){
	srm.Hello()
	tu := srm.User{"yjc",45}
	fmt.Println("结构 体",tu)
	tu.ToString()
	tu.GetAge()
//	fmt.Println("")
}
