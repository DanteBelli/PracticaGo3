package main

import "fmt"

func main() {
	ints := map[string]int64{
		"primero": 11,
		"segund":  45,
	}
	floats := map[string]float64{
		"primero": 11.55,
		"segundo": 48.75,
	}
	fmt.Printf("Las sumas de los numeros pasados a int y float son: %v y %v\n",
		sumInt(ints),
		sumFloat(floats))
}
func sumInt(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}
func sumFloat(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}
