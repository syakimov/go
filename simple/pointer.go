package main

import "fmt"

type User struct {
	Name string
}

func Rename(u *User, name string) {
	fmt.Println("Old name", u.Name)
	u.Name = name
}

func Copy(u User) User {
	return u
}

func main() {
	p := &User{"Stoyan"}
	Rename(p, "Stefan")

	c := Copy(*p)
	Rename(&c, "New Name")
	fmt.Println("Passed by ref", p.Name == c.Name)

	u := *p
	fmt.Println(u)
}
