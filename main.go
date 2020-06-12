package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type User struct {
	id                         int
	firstName, lastName, email string
}

type Group struct {
	role           string
	users          []User
	newestUser     User
	spaceAvailable bool
}

type Point struct {
	x, y int
}

// Describer interface describes the struct being passed in
type Describer interface {
	describe() string
}

func (g *Group) describe() string {
	if len(g.users) > 2 {
		g.spaceAvailable = false
	}
	return fmt.Sprintf("This user group has %d. The newest user is %s %s. Accepting New Users: %t", len(g.users),
		g.newestUser.firstName, g.newestUser.lastName, g.spaceAvailable)
}

func (u *User) describe() string {
	return fmt.Sprintf("id: %d\nfirstName: %s\nlastName: %s\nemail: %s", u.id, u.firstName, u.lastName, u.email)
}

// DoTheDescribing describes whatever struct is passed in
func DoTheDescribing(d Describer) string {
	return d.describe()
}

func describeUser(u User) string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s", u.firstName, u.lastName, u.email)
	return desc
}

func describeGroup(g Group) string {
	if len(g.users) > 2 {
		g.spaceAvailable = false
	}
	return fmt.Sprintf("This user group has %d. The newest user is %s %s. Accepting New Users: %t", len(g.users),
		g.newestUser.firstName, g.newestUser.lastName, g.spaceAvailable)
}

func changeName(name *string) {
	*name = strings.ToUpper(*name)
}

func updateEmail(user *User) {
	user.email = "Turn up"
	user.email = strings.ToUpper(*&user.email)

}

func isGreaterThanTen(n int) error {
	if n < 10 {
		return errors.New("Somethign went fucking bad")
	}
	return nil
}

func check(input int) error {
	if input < 0 {
		return errors.New("Invalid input")
	}
	return nil
}

func openTxtFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}

func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("Help! Ok here I come")
		fmt.Println(r)
	}
}
func doTings() {
	defer recoverFromPanic()
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i == 2 {
			panic("AhhhhHHH!!")
		}
	}
}

func main() {

	doTings()
	u0 := User{id: 0, firstName: "Charles", lastName: "Onyewuenyi", email: "charles@gmail.com"}
	u1 := User{id: 1, firstName: "Maya", lastName: "Onyewuenyi", email: "maya@gmail.com"}
	g0 := Group{
		role:           "admin",
		users:          []User{u0, u1},
		newestUser:     u1,
		spaceAvailable: true,
	}

	fmt.Println(DoTheDescribing(&u0))
	fmt.Println(DoTheDescribing(&g0))
	// path := "/Users/charlesonyewuendyi/go/src/sbox/file.txt"
	// if err := openTxtFile(path); err != nil {
	// 	fmt.Println(fmt.Errorf("%v", err))
	// }

	// p0 := Point{x: 55, y: 88}
	// fmt.Println(p0)

	// u0 := User{id: 0, firstName: "Charles", lastName: "Onyewuenyi", email: "charles@gmail.com"}
	// u1 := User{id: 1, firstName: "Maya", lastName: "Onyewuenyi", email: "maya@gmail.com"}

	// g0 := Group{
	// 	role:           "admin",
	// 	users:          []User{u0, u1},
	// 	newestUser:     u1,
	// 	spaceAvailable: true,
	// }

	// fmt.Println(describeUser(u0))
	// fmt.Println(describeGroup(g0))

	// var name string = "Charles"
	// var namePtr *string = &name
	// var nameVal string = *namePtr

	// fmt.Println(name)
	// fmt.Println(namePtr)
	// fmt.Println(nameVal)

	// name2 := "Ray"
	// changeName(&name2)
	// fmt.Println(name2)

	// u3 := User{id: 2, firstName: "James", lastName: "Hammond", email: "james@gmail.com"}

	// fmt.Println("Email = ", u3.email)
	// updateEmail(&u3)
	// fmt.Println("Updated email = ", u3.email)

}
