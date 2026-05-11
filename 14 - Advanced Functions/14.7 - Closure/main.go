package main

import "fmt"

func closure() func(){
	text := "Function Closure";
	function := func(){
			fmt.Println(text);
	}
	return function;
}
func main(){
	text := "Function Main";
	fmt.Println(text);

	newFunction := closure();

	newFunction();
}