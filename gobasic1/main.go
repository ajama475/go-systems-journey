package main

import "fmt"

/*
An interface for all users 
*/
type User interface {
	getName() string
	getRole() string
	getAccessLevel() int
}

/*
Admin struct that implements User inteface
*/
type Admin struct {
	Name string
	Department string
}

/*
This methods returns admin's name
*/
func (a Admin) getName() string {
	return a.Name
}

/*
This methods returns admin's role
*/
func (a Admin) getRole() string {
	return "Admin"
}

/*
This methods returns admin's access level
for now it's hardcorded to 10 for simplicity
*/
func (a Admin) getAccessLevel() int {
	return 10
}

/*
RegularUser struct that implements User inteface
*/
type RegularUser struct {
	Name string
	MembershipLevel int
}

/*
This methods returns regular user's name
*/
func (r RegularUser) getName() string {
	return r.Name
}

/*
This methods returns regular user's role
*/
func (r RegularUser) getRole() string {
	return "User"
}

/*
This methods returns regular user's access level
for now it's hardcorded to 1 for simplicity
*/
func (r RegularUser) getAccessLevel() int {
	return 1
}

/*
The function prints the user's name, role, access leve and 
further explaining sentence
*/
func displayInfo(user User) {
	name := user.getName()
	role := user.getRole()
	accessLevel := user.getAccessLevel()
	
	//print info
	fmt.Println("User:", name)
	fmt.Println("Role:", role)
	fmt.Println("Access Level:", accessLevel)

	if role == "Admin" {
		fmt.Println("Administrative privileges granted")

	}else if role == "User" {
		fmt.Println("Standard user permissions")
	}

	switch accessLevel {
	case 1:
		fmt.Println("Limited access")
	case 10:
		fmt.Println("Full system access")
	}
}

func main(){
	fmt.Printf("Getting information...\n")
  	displayInfo(Admin{
		Name: "Alice",
		Department: "Science",
	})
	
	fmt.Printf("\nGetting information...\n")
	displayInfo(RegularUser{
		Name: "Ken",
		MembershipLevel: 4,
	})
}