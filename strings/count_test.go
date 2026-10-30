package strings

import "fmt"

func ExampleCountA() {
	fmt.Println(CountA("ABC"))
	fmt.Println(CountA("ABCABC"))
	fmt.Println(CountA("AAAABBBA"))
	fmt.Println(CountA("DEFGH"))

	// Output:
	// 1
	// 2
	// 5
	// 0
}
