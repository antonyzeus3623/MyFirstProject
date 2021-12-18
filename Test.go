package Test

import "fmt"

func testSlice() {
	var slice []int
	var arr [5]int = [...]int{1, 2, 3, 4, 5}

	slice = arr[:4]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

}

func main() {
	testSlice()
}
