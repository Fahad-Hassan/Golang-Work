package main

import "fmt"

type Person struct{
	fname string
	lname string
	age int
	personcontact Contact
}
type Contact struct {
	Email string
	Phone string
}
func main(){
	var p Person
	p.fname ="fahad"
	p.lname = "hass"
	p.age = 33
	fmt.Println(p)
	ff :=33
	fmt.Println("hhhh",ff)
	pers :=Person{
		fname: "sss",
		lname: "saa",
		age: 11,
	}
	fmt.Println(pers)

	var pp =new(Person)
	pp.fname="ert"
	pp.lname="dd"
	pp.age=55
	pp.personcontact.Email="rrrrr"
	fmt.Println(pp)
}