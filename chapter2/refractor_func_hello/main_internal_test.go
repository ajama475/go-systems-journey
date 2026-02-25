package main

import "testing"

func Example_main() {
	main()
	// Output:
	// Hello, world
}

func TestGreet(t *testing.T) {
	//preparation phase
	expectedGreeting := "Hello, world"

	//execution phase
	greeting := greet()

	//decision phase: check the returned value
	if greeting != expectedGreeting {
		//mark this as failed and see what you got 
		t.Errorf("expected: %q, got: %q", expectedGreeting, greeting)
	}


	
}