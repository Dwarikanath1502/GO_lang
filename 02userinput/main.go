package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input..."
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza: ")

	//comma okay syntax // error okay syntax
	//here input and error come if u don't care about them just put (,) like here i put for error
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T", input)

}
