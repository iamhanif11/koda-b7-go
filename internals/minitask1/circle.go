package minitask1

import "math"

func GetKeliling(radius float32) (keliling float32) {
	keliling = 2 * float32(math.Pi) * radius
	return keliling
}

func GetLuas(radius float32) (luas float32) {
	luas = float32(math.Pi) * radius * radius

	return luas
}

func GetKelilingAndLuas(radius float32) (keliling float32, luas float32) {
	keliling = 2 * float32(math.Pi) * radius
	luas = float32(math.Pi) * radius * radius

	return keliling, luas
}
