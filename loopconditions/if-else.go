package loopconditions

import "fmt"

func IfElse() {
	/*
		if (condition) {
			Codes...
		}
	*/

	x := 25
	if x%5 == 0 {
		fmt.Printf("%d is a multiple of 5\n", x)
	}

	/*
		Logical Operators
		----------------------------
		Operator	|	Description
		----------------------------
		&&				Logical AND
		||				Logical OR
		!				Logical NOT
	*/

	/*
		If-Else Statement
		if condition {
			code
		} else {
			code
		}
	*/
	if x%5 == 2 {
		fmt.Printf("%d is even\n", x)
	} else {
		fmt.Printf("%d is odd\n", x)
	}

	var BMI = 21.0
	if BMI < 18.5 {
		fmt.Println("You are underweight")
	} else if BMI >= 18.5 && BMI < 25.0 {
		fmt.Println("Your weight is normal")
	} else if BMI >= 25.0 && BMI < 30.0 {
		fmt.Println("You're overweight")
	} else {
		fmt.Println("You're obese")
	}

	/*
		If with a short statement
	*/
	if n := 10; n%2 == 0 {
		fmt.Printf("%d is even\n", n)
	}
}
