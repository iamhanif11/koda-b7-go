package minitask3

import (
	"fmt"
	"slices"
)

// func InsertNumber() {
// 	num := []int{50, 75, 66, 20, 32, 90}
// 	insertNum := 88

// 	backSlice := append([]int{insertNum}, num[3:]...)

// 	num = append(num[:3], backSlice...)

// 	for i, v := range num {
// 		fmt.Printf("Elemen ke-%d: %d\n", i, v)
// 	}

// }

func InsertNumber(angka int, index int) {
	data := []int{50, 75, 66, 20, 32, 90}

	fmt.Println("Data: ", data)

	data = slices.Insert(data, index, angka)

	fmt.Println("New Data: ", data)
	fmt.Println()

	for i, v := range data {
		fmt.Printf("Index %d: %d\n", i, v)
	}
}
