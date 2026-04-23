package main

import (
	"bufio"
	"calculatorProject/calculator"
	"fmt"
	"os"
	"strconv"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("========================")
	fmt.Println("      CALCULATOR")
	fmt.Println("========================")
	fmt.Println("1) Add")
	fmt.Println("2) Subtract")
	fmt.Println("3) Multiply")
	fmt.Println("4) Divide")
	fmt.Println("0) Exit")
	fmt.Println("========================")
	fmt.Print("Enter your choice: ")
	scanner.Scan()
	option := scanner.Text(); //Captured Text

	switch option {
	case "1":
		fmt.Print("Enter a number to add: ")
		scanner.Scan()
		input := scanner.Text()
		
		numberOne, err := strconv.ParseFloat(input,64)
		if err != nil{
			fmt.Println("Invalid Number !")
			return
		}

		fmt.Print("Enter another number to add: ")
		scanner.Scan()
		inputN2 := scanner.Text()
		
		numberTwo , err2 := strconv.ParseFloat(inputN2, 64)

		if err2 != nil{
			fmt.Println("Invalid Number !")
			return
		}

		sum := calculator.Sum(numberOne,numberTwo)
		fmt.Printf("This result your sum is : %.2f\n" , sum)
		break;
	case "2" : 
		fmt.Print("Enter a number to subtraction: ")
		scanner.Scan()
		input := scanner.Text()

		numberOne , err := strconv.ParseFloat(input,64)

		if err != nil{
			fmt.Println("Invalid Number !")
			return
		}
		
		fmt.Print("Enter another number to subtraction: ")
		scanner.Scan()
		inputN2 := scanner.Text()

		numberTwo , err := strconv.ParseFloat(inputN2,64)

		if err != nil {
			fmt.Println("Invalid Number !")
			return
		}

		subtraction := calculator.Subtraction(numberOne,numberTwo);
		
		fmt.Printf("The result of this subtraction is : %2.f\n", subtraction)
		break;
		case "3" : 
		fmt.Print("Enter a number to multiply: ")
		scanner.Scan()
		input := scanner.Text()

		numberOne , err := strconv.ParseFloat(input,64)

		if err != nil{
			fmt.Println("Invalid Number !")
			return
		}
		
		fmt.Print("Enter another number to multiply: ")
		scanner.Scan()
		inputN2 := scanner.Text()

		numberTwo , err := strconv.ParseFloat(inputN2,64)

		if err != nil {
			fmt.Println("Invalid Number !")
			return
		}

		multiply := calculator.Multiplication(numberOne,numberTwo);
		
		fmt.Printf("The result of this subtraction is : %f\n", multiply)
		break;
	case "4" : 
		fmt.Print("Enter a number to division: ")
		scanner.Scan()
		input := scanner.Text()

		numberOne , err := strconv.ParseFloat(input,64)

		if err != nil{
			fmt.Println("Invalid Number !")
			return
		}
		
		fmt.Print("Enter another number to division: ")
		scanner.Scan()
		inputN2 := scanner.Text()

		numberTwo , err := strconv.ParseFloat(inputN2,64)

		if err != nil {
			fmt.Println("Invalid Number !")
			return
		}

		division := calculator.Division(numberOne,numberTwo);
		
		fmt.Printf("The result of this subtraction is : %2.f\n", division)
		break;
	}
}