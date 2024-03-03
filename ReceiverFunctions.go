package main

import (
	"fmt"
	"os"
)

type bill1 struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills

func newBill1(name string) bill1 {
	b := bill1{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

func (b *bill1) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	// list items
	for k,v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":",v)
		total += v
	}

	//total 
	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "total:", total)
	fs += fmt.Sprintf("%-25v ...$%0.2f", "tip:", b.tip)

	return fs
}

// update tip
func (b *bill1) updateTip(tip float64){
	b.tip = tip
}

// add an item to the bill
func (b *bill1) addItem(name string, price float64){
	b.items[name] = price
}

// save bill

func (b *bill1) save(){
	data := []byte(b.format())

	err := os.WriteFile(b.name+".txt",data,0644)

	if err !=nil {
		panic(err)
	}
	fmt.Println("bill was saved ti file")
} 