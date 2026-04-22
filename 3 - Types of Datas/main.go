package main

import (
	"errors"
	"fmt"
)

func main(){
	var numberOne int8 = 1;
	var numberTwo int16 = 12233;
	var numberTree int32 = 1223131231;
	var numberFour int64 = 1223131231321223131;
	//Based on the user's computer architecture
	var number int = 1223131231321223131;
	//Rune = int32
	var numberRune rune = 123131;
	//Uint Collection not allowed negative numbers
	var numberUint64 uint64 = 12313123
	
	var numberFloat32 float32 = 67.67
	var numberFloat64 float64 = 6767.6767

	var typeString string = "Hello World"
	var char = 'P'

	var boolean bool = true
	typeStringToper := "Hello World"

	var erro error = errors.New("Error")

	fmt.Println(numberOne)
	fmt.Println(numberTwo)
	fmt.Println(numberTree)
	fmt.Println(numberFour)
	fmt.Println(number)
	fmt.Println(numberRune)
	fmt.Println(numberUint64)
	fmt.Println(numberFloat32)
	fmt.Println(numberFloat64)
	fmt.Println(typeString)
	fmt.Println(typeStringToper)
	fmt.Println(char)
	fmt.Println(boolean)
	fmt.Println(erro)
}