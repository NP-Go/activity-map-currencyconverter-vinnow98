package main

import "fmt"

type currencyData struct {
	currency     string
	currencyName string
	exchangeRate float64
}

func main() {
	m := make(map[string]currencyData)

	m["USD"] = currencyData{"USD", "US dollar", 1.1318}
	m["JPY"] = currencyData{"JPY", "Japanese yen", 121.05}
	m["GBP"] = currencyData{"GBP", "Pound sterling", 0.90630}
	m["CNY"] = currencyData{"CNY", "Chinese yuan renminbi", 7.9944}
	m["SGD"] = currencyData{"SGD", "Singapore dollar", 1.5743}
	m["MYR"] = currencyData{"MYR", "Malaysian ringgit", 4.8390}

	fmt.Println(m["USD"])

	var currencyFrom string
	var howMuch float64
	var currencyTo string
	fmt.Println("From what currency?")
	fmt.Scanln(&currencyFrom)
	fmt.Println("How much?")
	fmt.Scanln(&howMuch)
	fmt.Println("To what currency?")
	fmt.Scanln(&currencyTo)

	result := (howMuch / m[currencyFrom].exchangeRate) * m[currencyTo].exchangeRate
	fmt.Println(result)

}
