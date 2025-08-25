package cdlm_my_go_module

import "fmt"

type Figura interface {
	area() float64
	perimetro() float64
}

func Medida(f Figura) {
	fmt.Println(f)
	fmt.Printf("Area: %.2f, Perimetro: %.2f\n", f.area(), f.perimetro())
}
