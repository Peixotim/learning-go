package main

import "fmt"

func recorvery(){
		if r:= recover(); r != nil{
			fmt.Println("Execution successfully recovered !")
		}
}

func studentIsApproved(n1 , n2 float64) bool{
		defer recorvery()
		average := (n1 + n2) / 2;
		
		if average > 6{
			return true
		}else if average < 6{
			return false
		}

		panic("The average is exactly 6.")
}

func main(){
	studentNotes := studentIsApproved(6,6);
	fmt.Println(studentNotes)
	fmt.Println("Post execution")
}