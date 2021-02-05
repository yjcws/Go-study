package srm

import "fmt"


//指针
func Point(){
	var n int = 20
	var ip *int

	//赋值
	ip = &n

	fmt.Println("指针访问",ip) // ip = &n = 0xc00000a0b8
	fmt.Println("指针访问",*ip)//*ip = 20



}

func sum(a *int) {
	*a = *a + 10
}

