package main

import "fmt"

func main(){
	var a int = 10
	b :=20

	fmt.Println("a:", a)
	fmt.Println("b:", b)

	var name string ="Kondo"
	var score float64 = 98.5
	var isOk bool = true

	fmt.Println("name", name)
	fmt.Println("score", score)
	fmt.Println("isOk", isOk)

	result := add(5,7)
	fmt.Println("5 + 7 =", result)
	
	result2 := add(6,7)
	fmt.Println("6 + 7 =", result2)

}

func add(x int, y int) int {
	return x+ y
}
