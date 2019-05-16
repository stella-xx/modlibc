package lib

import "fmt"

type User struct {
	name string
}

func NewUser(name string) *User {
	return &User{name: name}
}

func (u *User) Read(book string) {
	fmt.Printf("User:%s is reading Book:%s", u.name, book)
}

type Reader interface {
	Read()
}

