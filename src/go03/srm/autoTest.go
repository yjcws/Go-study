package srm
//自动测试
import (
	"fmt"
	"testing"
)


func TestGetUserByName(t *testing.T) {
	fmt.Println(getUserByName("shineyork"))
}

func TestGetUserById(t *testing.T) {
	fmt.Println(getUserById(9))
}


func getUserByName(name string) string {
	return "用户信息" + name
}

func getUserById(id int) int {
	// strconv.Itoa(id)
	return id
}

