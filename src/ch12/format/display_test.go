package format

func ExampleDisplay() {
	e := struct{a,c int }{
		10,
		0,
	}
	Display("e", e)
	// Output:
	// e.a = 10
	// e.c = 0
}
