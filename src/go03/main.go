package main

import "fmt"

func Map(){
	var countryMay map[string]string
	countryMay = make(map[string]string)



	countryMay["ZG"] = "北京"
	countryMay["US"] = "美国"

	for country := range countryMay{
		fmt.Println(country,"首都是",countryMay[country])
	}


}

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

}
