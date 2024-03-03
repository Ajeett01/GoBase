package main

import "fmt"

func main(){

	menu := map[string]float64{
		"spi": 4.99,
		"pie": 7.99,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looking maps
	for k, v := range menu{
		fmt.Println(k, "-",v)
	}

	// ints as key type
	phonebook := map[int]string{
		268455:"mario",
		625451:"ludo",
	}

	phonebook[268455] = "bowser"
	fmt.Println(phonebook)
}
