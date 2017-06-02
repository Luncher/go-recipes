package main

import "fmt"

const u, v float32 = 0, 1.0
const (
	c0 = iota
	c1 = iota
)

const (
	c2 = 1 << iota
	c3 = 1 << iota
)

func getName() (firstName, lastName string) {
	return "chen", "linchen"
}

func example(x int) {
	switch x {
		case 0:
			fmt.Printf("0\n")
		case 1:
			fmt.Printf("1\n")
		case 2:
			fallthrough
		case 3:
			fmt.Printf("3\n")
		case 4, 5, 6:
			fmt.Printf("4, 5, 6\n")
		default:
			fmt.Printf("Default\n")
	}
}

func myFunc(arg1 string, args ...int) {
	fmt.Println(arg1)	
	for _,v := range args {
		fmt.Println(v)
	}
}

func myPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
			case int:
				fmt.Println(arg, "is an int value.")
			case string:
				fmt.Println(arg, "is a string value.")
			case int64:
				fmt.Println(arg, "is an int64 value.")
			default:
				fmt.Println(arg, "is an unknow type.")
		}
	}
}

func main() {
	// example(1)
	// example(2)
	// example(6)
	// myFunc("hehe", 1, 2, 3, 4, 5)
	var j = 5
	a := func()(func()) {
		var i = 10
		return func() {
			fmt.Printf("i, j: %d, %d\n", i, j)
		}
	}()

	a()

	j *= 2

	a()	
}