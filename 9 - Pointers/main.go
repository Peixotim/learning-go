package main

//Ponteiros apontam para o endereco de memoria ou seja
//Eh como se fosse uma mira que mostra onde esta algo, e retorna o seu valor de referencia
//Referencia = localizacao
import "fmt"
func main(){
	
	var variableOne int = 10;
	var variableTwo int = variableOne 

	fmt.Println(variableOne,variableTwo)
	variableOne++

	fmt.Println(variableOne,variableTwo)

	var variableThree int = 100
	var pointer *int = &variableThree


	fmt.Println(variableThree,pointer) // Assim recebe o valor da referencia de memoria
	fmt.Println(variableThree,*pointer) //Assim recebe o valor de onde esta a referencia de memoria
	variableThree+= 100

	fmt.Println(variableThree,*pointer)
}