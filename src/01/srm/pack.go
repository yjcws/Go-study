package srm

import (
	"fmt"
)

func Hello(){

	fmt.Print("hello")
}

type User struct {
	Name string
	Age int

}

func (u User) ToString(){
	fmt.Println("tostring:",u.Name)

}

func (u User) GetAge(){
	fmt.Println("getage():",u.Age)
}
