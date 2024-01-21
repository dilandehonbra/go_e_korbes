// You can edit this code!
// Click here and start typing.
package main

import "fmt"

var oi string

func main() {

	somaInfo, somaValor, paramInfo, paramLen, oi := soma("", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50)

	fmt.Println(oi)

	fmt.Println(somaInfo, somaValor, paramInfo, paramLen)

}

func soma(s string, x ...int) (string, int, string, int, string) {

	if s == "" {
		oi = "Não sei o período em que estamos, mas em todo caso olá!"

	} else if s == "manhã" {
		oi = "Bom dia!"
	} else if s == "tarde" {
		oi = "Boa tarde!"
	} else {
		oi = "Boa noite!"
	}

	soma := 0
	for _, v := range x {
		soma += v
	}
	return "A soma dos números é - >", soma, "\nO total de parâmetros calculados foi ->", len(x), oi

}
