package main

import "fmt"

func main(){
	func(){
		fmt.Println("Hello World !")
	}()

	func (number int){
		fmt.Println(number)
	}(10)

	func (nums ...int){
		totalSum := 0
		for _ , number := range nums{
			totalSum += number;
	}
		fmt.Println(totalSum)
	}(10,10,10,10,10,10,10,10,10,10)
	
}