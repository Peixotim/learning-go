package main

import "fmt"


func main(){

	var arrayOne [4] string
	arrayOne[0] = "One Position"
	fmt.Println(arrayOne);

	arrayTwo := [4]string{"Zero","One","Two","Three"}
	fmt.Println(arrayTwo);
	fmt.Println(arrayTwo[0]);

	sliceOne := []string{"Hello World"}
	fmt.Println("SliceOne : ",sliceOne)
	sliceOne = append(sliceOne, "Good Morning !")
	fmt.Println("SliceOne with append : " , sliceOne)
}