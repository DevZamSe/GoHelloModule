package helloWord

import "fmt"

func imprimirHola() {
	defer fmt.Println("Mundo")
	fmt.Println("Hola")
}
