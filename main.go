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

// func main() {
// 	for i := 5; i >= 1; i-- {
// 		for c := 1; c <= i; c++ {
// 			fmt.Print(" * ")
// 		}
// 		fmt.Print("\n")
// 	}
// }

func main() {
	calculation()

}

func calculation() {

	var operator, x, y, z int

	fmt.Print("Inter Operator \n")
	fmt.Print("1 for Addtion \n")
	fmt.Print("2 for Substraction \n")
	fmt.Print("3 for Multiplication \n")
	fmt.Print("4 for Division \n")
	fmt.Scanln(&operator)

	if operator > 4 {
		calculation()

	}
	if operator == 1 {
		for z == 0 {
			fmt.Println("Enter two number to add")
			fmt.Scanln(&x, &y)
			fmt.Println("the answer is", x+y)
			fmt.Println("Want to add some other number\n 0 = Yes \n 1 = No")
			fmt.Scan(&z)
		}
	}
	if operator == 2 {
		for z == 0 {
			fmt.Println("Enter two number to sub")
			fmt.Scanln(&x, &y)
			fmt.Println("the answer is", x-y)
			fmt.Println("Want to add some other number\n 0 = Yes \n 1 = No")
			fmt.Scan(&z)
		}
	}
	if operator == 3 {
		for z == 0 {
			fmt.Println("Enter two number to multiply")
			fmt.Scanln(&x, &y)
			fmt.Println("the answer is", x*y)
			fmt.Println("Want to add some other number\n 0 = Yes \n 1 = No")
			fmt.Scan(&z)
		}
	}
	if operator == 4 {
		for z == 0 {
			fmt.Println("Enter two number to devide")
			fmt.Scanln(&x, &y)

			for y < 1 {
				fmt.Println("number can't be divided by 0")
				fmt.Println("Enter two number to devide")
				fmt.Scanln(&x, &y)

			}
			fmt.Println("the answer is", x/y)
			fmt.Println("Want to add some other number\n 0 = Yes \n 1 = No")
			fmt.Scan(&z)

		}
	}

}

// }
// if operator == 1 {
// 	for z == 0 {
// 		fmt.Println("Enter two number to add")
// 		fmt.Scanln(&x, &y)
// 		fmt.Println("the answer is", x+y)
// 		fmt.Println("Want to add some other number\n 0 = Yes \n 1 = No")
// 		fmt.Scan(&z)
// 	}
// }
// if operator == 2 {
// 	for z == 0 {
// 		fmt.Println("Enter two number to add")
// 		fmt.Scanln(&x, &y)
// 		fmt.Println("the answer is", x-y)
// 		fmt.Println("Want to add some other number\n 0 = Yes \n 1 = No")
// 		fmt.Scan(&z)
// 	}
// }
// if operator == 3 {
// 	for z == 0 {
// 		fmt.Println("Enter two number to add")
// 		fmt.Scanln(&x, &y)
// 		fmt.Println("the answer is", x*y)
// 		fmt.Println("Want to add some other number\n 0 = Yes \n 1 = No")
// 		fmt.Scan(&z)
// 	}
// }
// if operator == 4 {
// 	for z == 0 {
// 		fmt.Println("Enter two number to add")
// 		fmt.Scanln(&x, &y)
// 		fmt.Println("the answer is", x/y)
// 		fmt.Println("Want to add some other number\n 0 = Yes \n 1 = No")
// 		fmt.Scan(&z)
// 	}
// }
