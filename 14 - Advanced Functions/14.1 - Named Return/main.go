package main

import (
	"fmt"
)
func calculatorSumAndSubtraction(numOne int , numTwo int) (sum int , sub int){
		sum = numOne + numTwo;
		sub = numOne - numTwo;

		return sum , sub
}

func main(){

	fmt.Println(calculatorSumAndSubtraction(20,20))
}