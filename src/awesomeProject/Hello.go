package main

import (
	"awesomeProject/internal"
	fm "fmt"
)

func main() {
	message := greetMe("world")
	fm.Println(message)
	myStrunct1 := internal.MyStrunct{}
	myStrunct2 := internal.MyStrunct{FirstName: "sam", LastName: "cardonis", Age: 23}
	myStrunct3 := internal.MyStrunct{FirstName: "sam"}
	myStrunct1.FirstName="1"
	myStrunct2.FirstName="1"
	myStrunct3.FirstName="1"
	println(8,3)
	println(8,0)
	res, sheerit := devide(8,3)
	fm.Printf("%d, %d", res, sheerit)

	res2, sheerit2 := devide(8,0)
	fm.Printf("%d, %d", res2, sheerit2)
	slc2 := make([]int,2,4)
}

func sum(a,b int, c bool) (int,int) {
	return a+b, 2
}


func devide(a,b int) (int,int) {
	var res int = a / b
	sheerit := a%b  //use := to avoid type declaration
	return res,sheerit

}

func greetMe(name string) string {
	return "Hello, " + name + "!"
}