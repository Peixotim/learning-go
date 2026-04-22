package main

import "fmt"

type player struct{
	id uint8
	username string
	email string
	password string
	ranks ranks
}

type ranks struct{
	rank string
}

func main(){
	fmt.Println("Hello World")
	var playerOne player
	playerOne.id = 10
	playerOne.username = "peixotim"
	playerOne.email = "pedropeixotovz@gmail.com"
	playerOne.password = "1234"
	fmt.Println("Player one : " , playerOne)

	rankExample := ranks{"Example"}
	playerTwo := player{10,"pedro","pedropeixotovz@gmail.com","1234",rankExample}

	fmt.Println("Player two : ", playerTwo)
	
	playerThree := player{username:"peixotim"}
	fmt.Println("Player three" , playerThree)

	playerFour := player{
		id:1,
		username: "peixotim",
		email:"pedropeixotovz@gmail.com",
		password: "1234"}
	fmt.Println(playerFour)
}
