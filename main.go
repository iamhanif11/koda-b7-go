package main

import (
	"fmt"
	"math"
)

func main() {
	hasilKeliling := getKeliling(7)
	fmt.Printf("Keliling Lingkaran: %.2f\n", hasilKeliling)

	hasilLuas := getLuas(7)
	fmt.Printf("Luas Lingkaran: %.2f\n", hasilLuas)

	keliling, luas := getKelilingAndLuas(7)
	fmt.Printf("Keliling = %.2f\nLuas= %.2f", keliling, luas)
}

func getKeliling(radius float32) (keliling float32) {
	keliling = 2 * float32(math.Pi) * radius
	return keliling
}

func getLuas(radius float32) (luas float32) {
	luas = float32(math.Pi) * radius * radius

	return luas
}

func getKelilingAndLuas(radius float32) (keliling float32, luas float32) {
	keliling = 2 * float32(math.Pi) * radius
	luas = float32(math.Pi) * radius * radius

	return keliling, luas
}
