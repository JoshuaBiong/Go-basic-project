// Output Right Triangle:
// *
// * *
// * * *
// * * * *
// * * * * *
//

package main

import (
	"fmt"
)

// func main() {
// 	for i := 1; i <= 5; i++ {
// 		for c := 1; c <= i; c++ {
// 			fmt.Print("* ")
// 		}
// 		fmt.Print("\n")
// 	}
// }

// Output Reverse Right Triangle

// * * * * *
// * * * *
// * * *
// * *
// *

func main() {
	for i := 5; i >= 1; i-- {
		for c := 1; c <= i; c++ {
			fmt.Print(" * ")
		}

		fmt.Print("\n")
	}
}
