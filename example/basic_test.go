package example

import "fmt"

func ExampleBasicEma() {
	fmt.Println(BasicEma().Calculate(3).FormattedString(2))
	// Output: 102.50
}
