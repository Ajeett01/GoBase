package main

import "fmt"

func main(){
	mybill := newBill1("mario's bill")

	mybill.updateTip(10)

	fmt.Println(mybill.format())
}