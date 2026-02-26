package main

import "testing"

//english testing
func TestGreet_English(t *testing.T) {
	//preperation
	lang := language("en")
	want := "Hello, world"

	//execution
	got := greet(lang)

	//decison
	if got != want {
		//mark it as failure 
		t.Errorf("expected: %q, got: %q", want, got)
	}
}

//french testing
func TestGreet_French (t *testing.T) {
	//preperation
	lang := language("fr")
	want := "Bonjour le monde"

	//execution
	got := greet(lang)

	//decision
	if got != want {
		t.Errorf("expected: %q, got: %q", want, got)
	}
}

//non-existing langugae testing
func TestGreet_kerth(t *testing.T) {
	//preperation
	lang := language("kt")
	want := ""

	//execution
	got := greet(lang)

	//decision
	if got != want {
		t.Errorf("expected: %q, got: %q", want, got)
	}
}