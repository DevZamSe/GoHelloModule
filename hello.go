package helloWord

import "fmt"

func ImprimirHola() {
	defer fmt.Println("Mundo")
	fmt.Println("Hola")
}
