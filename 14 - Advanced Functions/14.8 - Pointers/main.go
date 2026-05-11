package main

import "fmt"

//& = Busca o endereco de memoria
//* = Faz um apontamento para o endereco de memoria

func signalInversion(number int) int{
	return number * -1;
}

func signalInversionWithPointers(number * int){
	*number= *number * -1;
}
func main(){
	number := 20;
	numberInversion := signalInversion(number);
	fmt.Println(numberInversion);

	newNumber := 40
	fmt.Printf("New number : %d\n" ,newNumber)
	signalInversionWithPointers(&newNumber);
	fmt.Println(newNumber);
}