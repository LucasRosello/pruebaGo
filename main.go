package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type calc struct{
	 
}

func (calc) operate(entrada string, operador string) int {
	
	numeros := strings.Split(entrada, operador)
	numero0 := Parsear(numeros[0])
	numero1 := Parsear(numeros[1])

	switch operador{
	case "+":
		fmt.Println(numero0 + numero1)
	case "-":
		fmt.Println(numero0 - numero1)
	case "*":
		fmt.Println(numero0 * numero1)
	case "/":
		fmt.Println(numero0 / numero1)
	default:
		fmt.Println("che, ¿Que carajo pusiste?")
	}

	return numero0
}

func Parsear(entrada string) int {
	numero, _ := strconv.Atoi(entrada)
	return numero
}

func leerEntrada() string{
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main()  {
	entrada := leerEntrada()
	operador := leerEntrada()
	c := calc{}
	fmt.Println(c.operate(entrada, operador))






}