package main

import "fmt"


func main(){
	//Arithmetic
	sum := 100 + 204
	subtract := 202 - 394
	multply := 548 * 21389
	divison := 22 / 2
	rest := 44 % 2

	fmt.Println(sum)
	fmt.Println(subtract)
	fmt.Println(multply)
	fmt.Println(divison)
	fmt.Println(rest)

	//Assignment

	var variable string = "String One"
	variable2 := "String Two"

	fmt.Println(variable)
	fmt.Println(variable2)

	//Relational Operators

	fmt.Println(1 > 2)
	fmt.Println(2 < 1)
	fmt.Println(2 == 1)
	fmt.Println(2 != 4)
	fmt.Println(2 >= 4)
	fmt.Println(2 <= 5)

	//Logical Operators

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!false)


	//Unary Operators
	number := 2
	number++
	number += 1

	number--
	number -= 2
}