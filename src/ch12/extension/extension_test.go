package extension

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Println("...")
}

func (p *Pet) SpeakTo(name string) {
	p.Speak()
	fmt.Println(" ", name)
}

type Dog struct {
	//p *Pet
	Pet
}

func (d *Dog) Speak() {
	fmt.Println("Wang!")
}

func (d *Dog) SpeakTo(name string) {
	//d.p.Speak()
	d.Speak()
	fmt.Println(" ", name)
}

func TestDog(t *testing.T) {
	//dog := new(Dog)
	// ↓↓ 无法将 'new(Dog)' (类型 *Dog) 用作类型 Pet ↓↓
	//var dog Pet = new(Dog)
	var dog = new(Dog)
	// ↓↓ 无法支持 LSP ↓↓
	//var p = *Pet(dog)
	dog.SpeakTo("SSS")
}
