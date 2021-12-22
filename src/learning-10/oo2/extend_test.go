package oo2

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) speak() {
	fmt.Println("...")
}

func (p *Pet) speakTo(host string) {
	p.speak()
	fmt.Println(" ", host)
}

type Dog struct {
	p *Pet
}

type Cat struct {
	p *Pet
}

//extend
func (c *Cat) speak() {
	c.p.speak()
}

func (c *Cat) speakTo(host string) {
	c.p.speakTo(host)
}

//complex, implement by method
func (d *Dog) speak() {
	fmt.Println("Wang")
}

func (d *Dog) speakTo(host string) {
	d.speak()
	fmt.Println(host)
}

//匿名嵌套类型
type Horse struct {
	Pet
}

//尝试重载
type Duck struct {
	Pet
}

func (d *Duck) speak() {
	fmt.Println("Ga")
}

func TestDogSpeak(t *testing.T) {
	cat := new(Cat)
	cat.speakTo("Yi  ")
	dog := new(Dog)
	dog.speakTo("Yi")
	horse := new(Horse)
	horse.speakTo("Yi")

	//不支持显示转换和强转
	//var duck Pet := new(Duck)
	//var duck Duck := new(Duck)
	//var d = *Pet(duck)
	//d.speakTo("Yi")
	//duck.speakTo("Yi")

}
