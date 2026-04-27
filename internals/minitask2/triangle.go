package minitask2

import "fmt"

func GetTriangle(tinggi int) {
	for i := 0; i <= tinggi; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}
