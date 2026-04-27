package main

import "fmt"

type user struct{
	name string
	age int
}
func main(){
	newUser := user{"Pedro",19}
	users := map[string]string{
		"testing" : "Testing",
	}
	ages := map[int]int{
		1 : 1,
	}

	users2 := map[string]map[string]string{
		"nicknames":{
			"nickOne" : "Peixotim",
		},
		"Rewards":{
			"rewardOne" : "One Beer",
		},
	}
	fmt.Println(users)
	fmt.Println(ages)
	fmt.Println(newUser)
	fmt.Println(users2)
	delete(users2,"Rewards")
	fmt.Println(users2)
}