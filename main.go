package main

func main() {

	// no es necesario escribir las propi	edades pero si tiene que tener el mismo orden
	Go := Course{
		Name:    "Go desde Cero",
		Price:   12.34,
		IsFree:  false,
		UserIDs: []uint{12, 56, 89},
		Classes: map[uint]string{
			1: "Introduccion",
			2: "Estructuras",
			3: "Maps",
		},
	}

	Go.PrintClasses()
}
