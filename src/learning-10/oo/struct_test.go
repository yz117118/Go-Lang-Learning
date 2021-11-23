package oo

import (
	"fmt"
	"testing"
	"unsafe"
)

type User struct {
	Id   string
	Name string
	Age  int
}

func TestCreateUser(t *testing.T) {
	u := User{"1", "Aylee", 18}
	u1 := User{Name: "Aylee"}
	u2 := new(User) //返回指针
	u2.Id = "3"
	u2.Name = "Terry"
	t.Log(u)
	t.Log(u1)
	t.Log(u1.Name)
	t.Log(u2)
	t.Logf("e is %T", u)  //user
	t.Logf("e is %T", u2) // user ptr
}

func GetUserName(user User) string {
	//有对象复制的开销
	fmt.Println("Address is %\n", unsafe.Pointer(&user.Name))
	return user.Name
}

func GetUserId(user *User) string {
	//没有对象复制产生， 推荐
	fmt.Println("Pointer Address is %\n", unsafe.Pointer(&user.Name))
	return user.Id
}

func TestStructOperation(t *testing.T) {
	user := User{"1", "Aylee", 18}
	userPtr := &User{"1", "Aylee", 18}
	fmt.Println("Address is %\n", unsafe.Pointer(&user.Name))
	t.Log(GetUserName(user))
	t.Log(GetUserId(userPtr))

}
