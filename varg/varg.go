package main

import("fmt")

func MyPrintf(args ...interface{}){
	for _,arg := range args{
		switch arg.(type) {
		case int:
			fmt.Println(arg,"is an int value.")
		case string:
			fmt.Println(arg,"is a string float.")
		case int64:
			fmt.Println(arg,"is an int64 value.")
		default:
			fmt.Println(arg,"is an unknown type.")
		}

	}
}

func main(){
	var v1 int = 1
	var v2 string = "abcd"
	var v3 int64 = 1
	var v4 float32 = 1.234
	MyPrintf(v1,v2,v3,v4)
}