package latihangoroutine

// import "fmt"

// func main() {
// 	n := 10

// 	pipeline := FindSquare(FilterEvenNum(GenerateNum(n)))

// 	for result := range pipeline {
// 		fmt.Println("Result: ", result)
// 	}

// 	fmt.Println("---Selesai---")
// }

func GenerateNum(n int) <-chan int {
	num := make(chan int)
	go func() {
		for i := 1; i <= n; i++ {
			num <- i
		}
		close(num)
	}()
	return num
}

func FilterEvenNum(input <-chan int) <-chan int {
	num := make(chan int)
	go func() {
		for n := range input {
			if n%2 == 0 {
				num <- n
			}
		}
		close(num)
	}()
	return num
}

func FindSquare(input <-chan int) <-chan int {
	num := make(chan int)
	go func() {
		for n := range input {
			num <- n * n
		}
		close(num)
	}()
	return num
}
