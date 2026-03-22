 package main

 import ("testing")

 //admin GetName()
func TestGetName_admin(t *testing.T) {
	//preperation
	admin := Admin{
		Name: "Alice",
	}
	want := "Alice"

	//execution
	got := admin.getName()

	//decision
	if got != want {
		t.Errorf("Expected: %q got: %q", want, got)
	}
	
}

 //admin GetRole()
func TestGetRole_admin(t *testing.T) {
	//preperation
	admin := Admin{
		Name: "Alice",
	}
	want := "Admin"

	//execution
	got := admin.getRole()

	//decision
	if got != want {
		t.Errorf("Expected: %q got: %q", want, got)
	}
	
}

 //admin GetAccessLevel
func TestGetAccessLevel_admin(t *testing.T) {
	//preperation
	admin := Admin{
		Name: "Alice",
	}
	want := 10

	//execution
	got := admin.getAccessLevel()

	//decision
	if got != want {
		t.Errorf("Expected: %q got: %q", want, got)
	}
	
}

 //Regular user getName
func TestGetName_regular(t *testing.T) {
	//preperation
	regularUser := RegularUser{
		Name: "Ken",
	}
	want := "Ken"

	//execution
	got := regularUser.getName()

	//decision
	if got != want {
		t.Errorf("Expected: %q got: %q", want, got)
	}
	
}

//regular user GetRole()
func TestGetRole_regular(t *testing.T) {
	//preperation
	regularUser := RegularUser{
		Name: "Ken",
	}
	want := "User"

	//execution
	got := regularUser.getRole()

	//decision
	if got != want {
		t.Errorf("Expected: %q got: %q", want, got)
	}
	
}

//regular user getAccessLevel()
func TestGetAcessLevel_regular(t *testing.T) {
	//preperation
	regularUser := RegularUser{
		Name: "Ken",
	}
	want := 1

	//execution
	got := regularUser.getAccessLevel()

	//decision
	if got != want {
		t.Errorf("Expected: %q got: %q", want, got)
	}
	
}
