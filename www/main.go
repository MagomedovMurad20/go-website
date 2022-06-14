package main

import (
	"fmt"
	"net/http"
)

type User struct {
	name                  string
	age                   uint16
	money                 int16
	avg_grades, happyness float64
}

func (u *User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and money %d", u.name, u.age, u.money)
}

func (u *User) setNewName(newName string) {
	u.name = newName
}

func home_page(w http.ResponseWriter, request *http.Request) {
	// bob := User{name: "Bob", age: 25, money: -25, avg_grades: 4.2, happyness: 0.8}
	bob := User{"Bob", 25, -18000, 4.2, 0.8}
	bob.setNewName("Alex")
	fmt.Fprintf(w, bob.getAllInfo())
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "94564556288394")
} //w это страница, а r  это request отслеживаем передаются ли данные хттп методом

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts", contacts_page)
	http.ListenAndServe(":7979", nil)
}

func main() {

	handleRequest()
}
