package minitask6

import (
	"fmt"
	"io"
	"os"
)

func ReadFile(path string) error {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("PANIC: ", path, r)
			fmt.Println("continue...")
		}
	}()

	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Gagal membuka: %w", err)
	}

	defer func() {
		fmt.Println("Close: ", path)
		file.Close()
	}()

	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Isi File:\n%s\n", string(content))
	return nil
}
