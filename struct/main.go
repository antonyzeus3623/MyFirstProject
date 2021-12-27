package main

import "fmt"

type Author struct {
	id   int
	name string
	addr string
	tel  string
	desc string
}

func main() {
	// var kk Author
	// fmt.Printf("%T:%#v\n", kk, kk)
	var kk Author = Author{
		01,
		"kk01",
		"西安市",
		"15500000000",
		"长风破浪会有时，直挂云帆济沧海",
	}
	fmt.Println(kk)
	fmt.Printf("%T:%#v\n", kk, kk)

}
