package cdlm_my_go_module

type Cuadrado struct {
	Lado float64
}

func (c *Cuadrado) area() float64 {
	return c.Lado * c.Lado
}

func (c *Cuadrado) perimetro() float64 {
	return 2 * (c.Lado + c.Lado)
}
