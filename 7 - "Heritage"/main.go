package main

import "fmt"

type person struct{
	name string
	age uint8
}

type student struct{
	person
	registration string
	college string
}

func main(){
 personOne := person{"Pedro",18}
 student := student{personOne,"PE18112006UD","UDEMY"}
 fmt.Printf("Student data : %v\n" , student)
}