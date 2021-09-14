package main

import "fmt"

func main() {
	mapa := make(map[string]int)
	fmt.Println(mapa)

	mapa["Juan"] = 32
	mapa["Yessica"] = 41
	mapa["Darío"] = 32
	fmt.Println(mapa)

	for i, v := range mapa {
		fmt.Println(i, v)
	}

	valueJ, ok := mapa["Juan"]
	fmt.Println(valueJ, ok)

	valueM, ok := mapa["Maria"]
	fmt.Println(valueM, ok)

	// ciudades := make(map[string]string, 5)
	ciudades := make(map[string]string)
	ciudades["Mexico"] = "CDMX"
	ciudades["Brasil"] = "Rio de Janeiro"
	ciudades["Italia"] = "Roma"
	fmt.Println(ciudades)
	fmt.Println()

	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 30,
	}

	campeonato["River Plate"] = 25
	campeonato["Chivas"] = 25
	delete(campeonato, "Real Madrid")

	for key, value := range campeonato {
		fmt.Printf("El equipo %s, tiene un puntaje de: %d \n", key, value)
	}
	fmt.Println()

	mineiroValue, mineiroExists := campeonato["Mineiro"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe %t \n", mineiroValue, mineiroExists)

	chivasValue, chivasExists := campeonato["Chivas"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe %t \n", chivasValue, chivasExists)
}
