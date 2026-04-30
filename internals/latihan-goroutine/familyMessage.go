package latihangoroutine

import (
	"fmt"
	"sync"
	"time"
)

// func main() {
// 	// BoardMessage()
// 	var wg sync.WaitGroup
// 	chn := make(chan string)

// 	wg.Add(1)
// 	go BoardMessage(chn, &wg)

// 	Sender(chn)

// 	wg.Wait()

// }

func BoardMessage(chn chan string, wg *sync.WaitGroup) {

	// chn := make(chan string)
	messages := []string{"A: Nanti kita mabar", "P: -1 satu Roam", "B: Nasi goreng dikulkas udah basi!", "N: Saya nanti pulang jam 2"}
	defer wg.Done()

	// wg.Add(1)
	// go Sender(chn, &wg)

	for _, message := range messages {
		chn <- message
	}
	close(chn)
}

func Sender(messageChn chan string) {
	defer func() {
		fmt.Println("Pesan telah dibaca")

	}()
	for message := range messageChn {
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("Ini pesan dari %s\n", message)
	}
}
