package main

import "fmt"

func sum (x int64 , y int64) int64{
	return x + y
}

func sub (x int64 , y int64) int64{
	return x -y
}

//The function has two return values
func mathematicalCalculations(x,y int8) (int8, int8){
	sum := x + y;
	subtract := x-y
	return sum,subtract
}


func main(){
	count := sum(10,20)

	fmt.Printf("Sum is result : %d\n" , count)

	var ageVerify = func (age int8){
		if age >= 18 {
			fmt.Println("You has age bigger 18 years")
		} else {
			fmt.Println("You are under 18 years of age.")
		}
	}
	ageVerify(18);

	resultSum , resultSub := mathematicalCalculations(10,20);
 	resultSumOne , _ := mathematicalCalculations(10,20);
	fmt.Println(resultSum)
	fmt.Println(resultSub)
	fmt.Println(resultSumOne)
}