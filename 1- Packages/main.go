package main

import (
	"fmt"
	"module/assistant"

	"github.com/badoux/checkmail"
)


func main(){
	fmt.Println("Writing to main file")
	assistant.Writing("Hello World")
	fmt.Print("Exported func : ")
	assistant.ExportedFunc()

	erro := checkmail.ValidateFormat("pedropeixotovz@gmail.com")
	fmt.Println(erro);
}