package main

import "fmt"

// Ser vivo
type serVivo interface {
	estaVivo() bool
}

func estoyVivo(v serVivo) bool {
	return v.estaVivo()
}

// Humano
type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
}
type hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	esHombre   bool
	vivo       bool
}
type mujer struct {
	hombre
}

func (h *hombre) respirar() {
	h.respirando = true
}
func (h *hombre) comer() {
	h.comiendo = true
}
func (h *hombre) pensar() {
	h.pensando = true
}
func (h *hombre) sexo() string {
	if h.esHombre {
		return "Hombre"
	} else {
		return "Mujer"
	}
}
func (h *hombre) estaVivo() bool {
	return h.vivo
}
func HumanosRespirando(hu humano) {
	hu.respirar()
	fmt.Printf("Soy un/a %s, y ya estoy respirando \n", hu.sexo())
}

// Animal
type animal interface {
	respirar()
	comer()
	EsCarnivoro() bool
}
type perro struct {
	respirando bool
	comiendo   bool
	carnivoro  bool
	vivo       bool
}

func (p *perro) respirar() {
	p.respirando = true
}
func (p *perro) comer() {
	p.comiendo = true
}
func (p *perro) EsCarnivoro() bool {
	return p.carnivoro
}
func (p *perro) estaVivo() bool {
	return p.vivo
}
func AnimalesRespirar(an animal) {
	an.respirar()
	fmt.Println("Soy un animal y estoy respirando")
}
func AnimalesCarnivoros(an animal) int {
	if an.EsCarnivoro() == true {
		return 1
	}
	return 0
}

type vegetal interface {
	ClasificacionVegetal()
}

type figuras2D interface {
	area() float64
}

type cuadrado struct {
	base float64
}

type rectangulo struct {
	altura float64
	base   float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func calcular(f figuras2D) {
	fmt.Println("Area", f.area())
}

func main() {
	miCuadrado := cuadrado{
		base: 10,
	}

	miRectangulo := rectangulo{
		altura: 10,
		base:   20,
	}

	calcular(miCuadrado)
	calcular(miRectangulo)

	// Lista de interfaces
	myInterface := []interface{}{"Hola", 12, 4.14}
	fmt.Println(myInterface...)

	Pedro := new(hombre)
	Pedro.esHombre = true
	HumanosRespirando(Pedro)
	Pedro.vivo = true
	fmt.Printf("Pedro vive %t \n", estoyVivo(Pedro))

	Maria := new(mujer)
	HumanosRespirando(Maria)
	// Maria.vivo = true
	fmt.Printf("Maria vive %t \n", estoyVivo(Maria))

	totalCarnivoros := 0
	Dogo := new(perro)
	Dogo.carnivoro = true
	Dogo.vivo = true
	AnimalesRespirar(Dogo)
	totalCarnivoros += AnimalesCarnivoros(Dogo)
	fmt.Printf("Dogo vive %t \n", estoyVivo(Dogo))

	fmt.Printf("Total carnivoros %d \n", totalCarnivoros)
}
