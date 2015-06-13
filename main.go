package main

import "fmt"

func main() {
	b, v, r := extraiReal("")
	fmt.Printf("%v, %v, %v\n", b, v, r)

	b, v, r = extraiReal("123")
	fmt.Printf("%v, %v, %v\n", b, v, r)

	b, v, r = extraiReal("123abc")
	fmt.Printf("%v, %v, %v\n", b, v, r)

	b, v, r = extraiReal("123.45")
	fmt.Printf("%v, %v, %v\n", b, v, r)

	b, v, r = extraiReal("123.45abc")
	fmt.Printf("%v, %v, %v\n", b, v, r)

	b, v, r = extraiReal("123.45.67")
	fmt.Printf("%v, %v, %v\n", b, v, r)

	b, v, r = extraiReal("123.45.abc")
	fmt.Printf("%v, %v, %v\n", b, v, r)

	b, v, r = extraiReal("abc123")
	fmt.Printf("%v, %v, %v\n", b, v, r)
}
