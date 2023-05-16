package main

import "fmt"

const ebulicaoK = 373.0

func main() {

	K := ebulicaoK
	C := ebulicaoK - 273.0

	fmt.Printf("A temperatura de ebulição da água em C = %g e a temperatura de ebulição da água em K = %g", C, K)

}
