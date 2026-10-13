package main

import "fmt"

func main() {
	fmt.Println(sumar(2, 3))

	resultado, err := dividir(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Resultado:", resultado)
	}
}

func sumar(a, b int) int {
	return a + b
}

func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("no se puede dividir entre cero")
	}
	return a / b, nil
}
