package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader)(string,error){
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill1{
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill1(name)
	fmt.Println("Created the bill -", b.name)

	return b
}

func promptOptions(b bill1){
	reader := bufio.NewReader(os.Stdin)

	opt,_ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)
	fmt.Println(opt)

	switch opt{
	case "a":
		name,_ := getInput("Item name:", reader)
		price,_ := getInput("Item price:", reader)

		p,err := strconv.ParseFloat(price,64)
		if err !=nil{
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.addItem(name,p)

		fmt.Println(name,price)
		promptOptions(b)
 
	case "t":
		tip,_ := getInput("Tip AMount :", reader)
		t,err := strconv.ParseFloat(tip,64)
		if err !=nil{
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)

		
		fmt.Println(tip)
		fmt.Println("Tip added")
		promptOptions(b)
 
	case "s":
		b.save()
		fmt.Println("you chose save file",b.name)

	default:
		fmt.Println("that was not a valid option...")
		promptOptions(b)
	}

}

func main(){
	mybill := createBill()
	promptOptions(mybill)
	fmt.Println(mybill)
}