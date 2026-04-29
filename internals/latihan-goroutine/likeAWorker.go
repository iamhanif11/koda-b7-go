package latihangoroutine

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	DailyRoutine()
}

func shower(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Sedang mandi...")
	time.Sleep(2 * time.Second)
	fmt.Println("Mandi Selesai")
}

func makeCoffee(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Siapkan gelas dan panaskan air")
	time.Sleep(3 * time.Second)
	fmt.Println("Masukkan Kopi, Gula dan Air Panas ke dalam gelas")
	fmt.Println("Aduk....")
	time.Sleep(1 * time.Second)
	fmt.Println("Nikmat Kopi")
}

func prepBreakfast(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Siapkan sarapan")
	time.Sleep(2 * time.Second)
	fmt.Println("Menyantap sarapan")
}

func cleanBedroom(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Ambil alat kebersihan")
	time.Sleep(2 * time.Second)
	fmt.Println("Bersihkan kasur")
}

func DailyRoutine() {
	var wg sync.WaitGroup
	fmt.Println("Bangun Tidur")

	wg.Add(4)

	go cleanBedroom(&wg)
	go shower(&wg)
	go makeCoffee(&wg)
	go prepBreakfast(&wg)

	wg.Wait()

	fmt.Printf("Ready! Berangkat kerja")

}
