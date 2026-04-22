package main

import "fmt"

func main(){
	var variableOne string = "Variable One"
	variableTwo := "Variable Two"

	fmt.Printf("%s",variableOne)
	fmt.Println("")
	fmt.Printf("%s",variableTwo)
	fmt.Println("")

	var(
		variableTree string = "Variable Tree"
		variableFour string = "Variable Four"
	)

	fmt.Println(variableTree)
	fmt.Println(variableFour)


	variableFive , variableSix := "Variable Five" , "Variable Six"

	fmt.Println(variableFive)
	fmt.Println(variableSix)

}