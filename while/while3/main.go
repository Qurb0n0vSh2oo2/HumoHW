package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	l := a
	// cnt := 0
	for l >= 0 {
		l -= b
	}
	
	fmt.Println("ostatok ot deleniya", b + l)
	for l >= 0 {
		l -= a
		if ( l >)
	}
	fmt.Println("deleniye", b + l)
}