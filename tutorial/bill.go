package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	return bill{
		name: name,
		items: map[string]float64{
			"pie":  5.99,
			"cake": 3.99,
		},
		tip: 0,
	}
}

func (b bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%v ... $%v \n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%v ... %v\n", "tip:", b.tip)

	fs += fmt.Sprintf("%v ... $%0.2f", "total:", total)

	return fs
}

func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func (b *bill) save() {
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}
	fmt.Println("bill was saved to file")
}
