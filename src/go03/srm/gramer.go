package srm

//语法
import (
	"fmt"
	"testing"
)

//if语法
func TestIf(t *testing.T) {
	if k := 9; k < 10 {
		fmt.Println("k", k)
	}
}


//Switch
func TestSwitch(t *testing.T) {
	//case后面的值 不能重复
	u := 0
	switch u {
	case 0:
		fmt.Println(0)
		fallthrough   //继续执行
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3, 4, 5:
		fmt.Println("")
	}

	//风格
	switch {
	case u > 0:
	case u < 5:
	}
}


//For语法
func TestFor(t *testing.T) {
	i := 9

	//for{}死循环的写法
	for {
		fmt.Println(i)
		if i == 5 {
			break
		}
		i--
	}
	for i > 0 {
		fmt.Println(i)
		i--
	}
}


//数组 和切片
func TestArr(t *testing.T) {
	//数组
	var arr1 [5]int = [5]int{0, 1, 2, 3, 2} // go数组定义是多长就是多长
	arr := arr1[3:]                         //  - low   了解即可，数组 的键都是索引
	fmt.Println("arr1", arr)
	fmt.Println("lep", len(arr)) //
	fmt.Println("cap", cap(arr))

	ar1 := arr1[:3] // 2
	fmt.Println("arr1", ar1)
	fmt.Println("lep", len(ar1)) //
	fmt.Println("cap", cap(ar1))

	ar2 := arr1[1:3] // 2
	fmt.Println("arr1", ar2)
	fmt.Println("lep", len(ar2)) //
	fmt.Println("cap", cap(ar2))
	// fmt.Println()
	// 0,1,2,3,4
	fmt.Println(arr1)
	var arr2 = [...]int{0, 1, 2, 3, 2}
	fmt.Println(arr2)
	var arr3 [2][2]int = [2][2]int{{0, 2}, {9, 0}}
	fmt.Println(arr3)

	// arr1[5] = 9
	// fmt.Println(arr1)
	// 以切片为主，常见的写法
	var sar1 []int = []int{0, 1, 2, 34} // 这种写法是常见
	fmt.Println("len:", len(sar1)) //4
	fmt.Println("cap:", cap(sar1))  //4
	sar1 = append(sar1, 10)
	fmt.Println("len1:", len(sar1))  //5
	fmt.Println("cap1:", cap(sar1))  //8  容量进行了扩充
	fmt.Println(sar1)
	sar1 = append(sar1, 20, 3, 8, 92)
	fmt.Println("len3:", len(sar1))  //6
	fmt.Println("cap3:", cap(sar1))  //  8 容量够没有增加
	fmt.Println(sar1)

	// make (t type, len , cap)
	sar2 := make([]int, 0, 5) // 常见定义方式
	// sar2[0] = 0
	fmt.Println(sar2)
}

func TestMap(t *testing.T) {
	// 类似于[
	// "kk" => fun
	// ]


	//局部写法
	map1 := make(map[string]string, 0)
	map1["name1"] = " shineyork"
	fmt.Println(map1)

	// 值类型是切片
	/*
	  [
	    "1123" => [
	        0: "p",
	        9: "k"
	    ]
	  ]
	*/

	//键：map[string]  ，值：[]string
	smap := make(map[string][]string, 0)
	s := make([]string, 0)
	s = append(s, "1", "0")
	smap["123"] = s
	fmt.Println(smap)
	// {map, map}
	// 切片为原型，value是map
	/*
			  [
			    0: [
			        "k":"v",
			        "o":"p"
			    ],
		      1 : [
		        "o":"ls"
		      ]
			  ]
	*/
	maps := make([]map[string]string, 2)
	// map2 := make(map[string]string, 0)
	// map2["name1"] = " shineyork"
	maps[0] = make(map[string]string, 0)
	maps[0]["l"] = "oos"
	maps[1] = map1
	fmt.Println(maps)

	for key, val := range maps {
		fmt.Println(key, val)
	}
}
