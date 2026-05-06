package main

import "fmt"

func functionOne(){
	fmt.Println("Function One")
}

func functionTwo(){
	fmt.Println("Function Two")
}

func functionThree(){
	fmt.Println("Function Three")
}

func main(){

	defer functionOne() //Atrasa o retorno do codigo, primeiro executa tudo e depois volta nele
	defer functionThree()
	functionTwo()
}