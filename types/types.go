package main

import "fmt"

//Here i have created a custom datatype User using struct (struct is a datatype containing multiple fields)
type User struct {
	name    stringbyte
	age     int
	adharId float32
}

//Here i have created a custom datatype stringbyte.
type stringbyte string

//interfaces
type activity interface {
	dance()
}

func (u *User) greet() string {
	u.name = "Shri krishna"
	return "hello, " + string(u.name)
}

func (u1 User) dance() {
	fmt.Println("u1 ", u1.name, "dances")
}

type Player struct {
	playername string
	playerId   int
}

func (p1 Player) dance() {
	fmt.Println("p1 ", p1.playername, "dances")
}

func whoDances(A activity) {
	A.dance()
}

func main() {
	// var string1 stringbyte = "Shiv"
	// fmt.Println("string1 : ", string1)

	// user := User{
	// 	name:    "Shri Ram",
	// 	age:     43,
	// 	adharId: 203.211,
	// }

	// fmt.Println(user.greet())
	// fmt.Println("user.name : ", user.name)
	// fmt.Println("user.age : ", user.age)

	u1 := User{name: "Shri Ram", age: 214, adharId: 10112031}
	p1 := Player{playername: "Virat Kohli", playerId: 18}

	whoDances(u1)
	whoDances(p1)

}
