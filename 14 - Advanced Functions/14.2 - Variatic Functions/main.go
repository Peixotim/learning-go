package main

import "fmt"


func sum(numbers ...int)int{
	totalSum := 0;

	for _ , number := range numbers{
			totalSum += number;
	}

	return totalSum;
}
func main(){
	fmt.Println(sum(2,2,4,4))
}