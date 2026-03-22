/*
This program builds a small user directory manager for a school club.

Each user in the system has:
- A name
- A role
- An age
- A list of interests
*/
package main

import (
	"fmt"
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
	interest []string
}

/*
user type Member
*/
type Member struct {
	name string
	age int
	role string
	interest []string
}

/*
Admin implements User inteface with getName()
return string
*/
func (a Admin) getName() string {
	return a.name
}

/*
Admin implements User inteface with getRole()
return string
*/
func (a Admin) getRole() string {
	return a.role
}

/*
Admin implements User inteface with getAge()
return int
*/
func (a Admin) getAge() int {
	return a.age
}


/*
Member implements User inteface with getName()
return string
*/
func (m Member) getName() string {
	return m.name
}

/*
Member implements User inteface with getRole()
return string
*/
func (m Member) getRole() string {
	return m.role
}

/*
Member implements User inteface with getAge()
return string
*/
func (m Member) getAge() int {
	return m.age
}

/*
three supported membership types in the club
return an array length of 3
*/
func membershipTypes() [3]string {
	types := [3]string{
		"general",
		"prospective",
		"pro",
	}
	return types

}

func test(user User,membership [3]string) {
	name := user.getName()
	age := user.getAge()
	role := user.getRole()
	
	fmt.Println("Getting User Information...")
	fmt.Printf("Name: %v\nAge: %d\nRole: %v\n", 
		name,
		age,
		role,
	)
	
	if role == "admin" {
	  fmt.Printf("Congratulations %v, you have been assigned to %v membership\n\n",
		    name,
		    membership[2],
	  )
	}
	
	if role == "member" {
	  fmt.Printf("Congratulations %v, you have been assigned to %v membership\n\n",
		    name,
		    membership[0],
	  )
	}
	 
	
	
}

func main() {
  membership := membershipTypes()
  fmt.Println(membership[0])
	
	user1 := Member{
	  name: "Cole",
		age: 20,
		role: "member",
		interest: []string{
			"baminton",
			"coding",
			"reading",
		},
	}
	test(user1, membership)
	
	user2 := Admin{
	  name: "Kali",
		age: 19,
		role: "admin",
		interest: []string{
			"Design",
			"Neurology",
			"Running",
		},
	}
	test(user2, membership)
}