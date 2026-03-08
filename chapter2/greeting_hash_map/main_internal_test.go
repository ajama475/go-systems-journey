package main

import ("testing")

//english
func TestGreet_english(t *tesing.T) {
	//preperation
	lang := phrasebook[language("en")]
	want := "Hello, world"

	//execution
	got := greet(lang)

	//decision
	if got != want {
		t.Errorf("expected: %q, got: %q", want, got)
	}

}

//french
func TestGreet_french(t *tesing.T) {
	//preperation
	lang := phrasebook[language("fr")]
	want := "Bonjour le monde"

	//execution
	got := greet(lang)

	//decision
	if got != want {
		t.Errorf("expected: %q, got: %q", want, got)
	}
}

//arabic
func TestGreet_arabic(t *tesing.T) {
	//preperation
	lang := phrasebook[language("ar")]
	want := "مرحبا بالعالم"

	//execution
	got := greet(lang)

	//decision
	if got != want {
		t.Errorf("expected: %q, got: %q", want, got)
	}
	
}

//chinese
func TestGreet_chinese(t *tesing.T) {
	//preperation
	lang := phrasebook[language("ch")]
	want := "你好世界"

	//execution
	got := greet(lang)

	//decision
	if got != want {
		t.Errorf("expected: %q, got: %q", want, got)
	}
	
}

//non existed language
func TestGreet_hindi(t *tesing.T) {
	//preperation
	lang := phrasebook[language("hi")]
	want := "हैलो वर्ल्ड"

	//execution
	got := greet(lang)

	//decision
	if got != want {
		t.Errorf("expected: %q, got: %q", want, got)
	}
	
}