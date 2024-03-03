package main
 
import "fmt"

func updateName(x *string) {
	*x = "wedge"
}

func main() {
	name := "tifa"

	updateName(&name) // Pass the address of the variable 'name'

	fmt.Println("memory address is ", &name)

	m := &name

	// fmt.Println(name)
	fmt.Println(m)
	fmt.Println(*m)

	updateName(m)
	fmt.Println(name)
}
