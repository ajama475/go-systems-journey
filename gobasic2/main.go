package main

import (
	"fmt"
	"errors"
)
/*
Represents a user in the system
*/

type User interface {
	getName() string
	getRole() string
	getAge() int
}

/*
user type Admin
*/
type Admin struct {
	name string
	age int 
	role string
	interest []string{}
}

/*
user type Member
*/
type Member struct {
	name string
	age int
	role string
	interest []string{}
}

/*
Admin implements User inteface wiht getName()
return string
*/
func (a Admin) getName() string {
	return a.name
}

/*
Admin implements User inteface wiht getRole()
return string
*/
func (a Admin) getRole() string {
	return a.role
}

/*
Admin implements User inteface wiht getAge()
return int
*/
func (a Admin) getAge() int {
	return a.age
}

/*
Member implements User inteface wiht getName()
return string
*/
func (m Member) getName() string {
	return m.name
}

/*
Member implements User inteface wiht getRole()
return string
*/
func (m Member) getRole() string {
	return m.role
}

/*
Member implements User inteface wiht getAge()
return string
*/
func (m Member) getAge() int {
	return m.age
}


func main() {
	//TODO
}