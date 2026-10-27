package lists

import "fmt"

func ExampleMinList() {

	l1 := MinList([]int{1, 2, 3, 4, 5})
	fmt.Println(l1)

	fmt.Println(MinList([]int{5, 4, 3, 2, 1}))
	fmt.Println(MinList([]int{1, 7, 2, -1, 5, -3}))
	fmt.Println(MinList([]int{}))

	// Output:
	// 1
	// 1
	// -3
	// 0
}

func ExampleMinList2() {

	l1 := MinList2([]int{1, 2, 3, 4, 5})
	fmt.Println(l1)

	fmt.Println(MinList2([]int{5, 4, 3, 2, 1}))
	fmt.Println(MinList2([]int{1, 7, 2, -1, 5, -3}))
	fmt.Println(MinList2([]int{}))

	// Output:
	// 1
	// 1
	// -3
	// 0
}
