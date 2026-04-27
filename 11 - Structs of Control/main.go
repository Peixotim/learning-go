package main

import "fmt"


var numberOne int = 10;

func main(){
	numberTwo := 20;

	if numberOne > numberTwo{
		fmt.Println("Number One > Number Two")
	}else{
		fmt.Println("Number Two > Number One")
	}

	if otherNumber := numberTwo; otherNumber > 10{
		fmt.Println("Other Number it is larger that 10")
	}


}