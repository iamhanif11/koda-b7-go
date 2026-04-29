package latihangoroutine

import (
	"fmt"
	"sync"
	"time"
)

//  func main (){
// 	var wg sync.WaitGroup
// 	wg.Go()
// 		coffe := []string{"Espresso", "Latte", "Cappucino"}
// 	wg.Go(func (){
// 		barista(coffe)
// 	})
// 	wg.Wait()
//  }

// func Coffeshop(coffe string) {

// }

func CoffeShop() {
	var wg sync.WaitGroup
	coffees := []string{"Kopi A", "Kopi B", "Kopi C"}

	for _, coffee := range coffees {
		wg.Add(1)
		go Barista(coffee, &wg)
	}

	wg.Wait()
}

func Barista(coffee string, wg *sync.WaitGroup) {
	fmt.Printf("[barista] sedang membuat pesanan: %s nama", coffee)

	time.Sleep(2 * time.Second)
	fmt.Printf("[barista] Pesanan Selesai: %s", coffee)

}
