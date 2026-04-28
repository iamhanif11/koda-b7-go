package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/iamhanif11/koda-b7-go/internals/minitask1"
	"github.com/iamhanif11/koda-b7-go/internals/minitask2"
	"github.com/iamhanif11/koda-b7-go/internals/minitask3"
	"github.com/iamhanif11/koda-b7-go/internals/minitask6"
	"github.com/iamhanif11/koda-b7-go/internals/minitask8"

	"github.com/iamhanif11/koda-b7-go/internals/minitask7"
	"github.com/iamhanif11/koda-b7-go/internals/models"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("===MENU===")
		fmt.Println("1. Keliling Lingkaran")
		fmt.Println("2. Luas Lingkaran")
		fmt.Println("3. Keliling & Luas Lingkaran")
		fmt.Println("4. Segitiga Siku-siku")
		fmt.Println("5. Sisip angka")
		fmt.Println("6. User")
		fmt.Println("7. ReadFile")
		fmt.Println("8. Person")
		fmt.Println("9. Payment")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih Menu: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "0":
			fmt.Println("Keluar...")
			return
		case "1":
			fmt.Print("Masukkan Jari-jari Lingkaran: ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			radius64, _ := strconv.ParseFloat(input, 32)
			radius := float32(radius64)
			result := minitask1.GetKeliling(radius)
			fmt.Printf("\n Keliling Lingkaran: %.2f\n\n\n", result)
		case "2":
			fmt.Print("Masukkan Jari-jari Lingkaran: ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			radius64, _ := strconv.ParseFloat(input, 32)
			radius := float32(radius64)
			result := minitask1.GetLuas(radius)
			fmt.Printf("\n Luas Lingkaran: %.2f\n\n\n", result)
		case "3":
			fmt.Print("Masukkan Jari-jari Lingkaran: ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			radius64, _ := strconv.ParseFloat(input, 32)
			radius := float32(radius64)
			k, l := minitask1.GetKelilingAndLuas(radius)
			fmt.Printf("\n Keliling Lingkaran: %.2f\n Luas Lingkaran: %2f\n\n", k, l)
		case "4":
			fmt.Print("Masukkan tinggi Segitiga: ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			tinggi, _ := strconv.Atoi(input)
			minitask2.GetTriangle(tinggi)
		case "5":
			fmt.Print("Masukkan angka: ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			angka, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Error")
				continue
			}
			fmt.Println("Masukkan index (0-6): ")
			inputIndex, _ := reader.ReadString('\n')
			inputIndex = strings.TrimSpace(inputIndex)

			index, err := strconv.Atoi(inputIndex)
			if err != nil {
				fmt.Println("Error")
				continue
			}

			minitask3.InsertNumber(angka, index)
			fmt.Println()

		case "6":
			data := models.TampilinData()
			fmt.Println(data)

		case "7":
			fmt.Println("Masukkan path file")
			filepath, _ := reader.ReadString('\n')
			filepath = strings.TrimSpace(filepath)

			if filepath == "" {
				fmt.Println("Path tidak boleh kosong")
				continue
			}

			err := minitask6.ReadFile(filepath)
			if err != nil {
				fmt.Printf("Error: %n\n", err)
			}
			fmt.Println()

		case "8":
			p := minitask7.NewPerson("Hanif", "Padang", "12343221")

			fmt.Println(p.Print())
			p.Greet()

			p.SetName("Hanif Irfan")
			fmt.Println("Setelah set nama: ")
			p.Greet()
			fmt.Println()

		case "9":

			prices := []int{5000, 10000, 25000, 50000}
			invalidPrices := []int{10000, -5000}

			bank := minitask8.BankPayment{BankName: "BNI"}
			online := minitask8.OnlinePayment{Platform: "Gopay"}
			fiction := &minitask8.FictionPayment{}
			fmt.Println("---Transaksi Bank---")
			if err := bank.Pay(prices); err != nil {
				fmt.Println(err)
			}

			fmt.Println("\n---Transaksi Online---")
			if err := online.Pay(prices); err != nil {
				fmt.Println(err)
			}

			fmt.Println("\n---Transaksi Invalid---")
			if err := bank.Pay(invalidPrices); err != nil {
				fmt.Println(err)
			}

			fmt.Println("\n---Fiction Transaksi----")
			fiction.Pay(prices)
			fiction.Pay([]int{20000})

			fmt.Printf("Data di Slice: %v\n", fiction.TotalPayment)
		}

	}

	// hasilKeliling := minitask1.GetKeliling(7)

	// hasilLuas := minitask1.GetLuas(7)
	// fmt.Printf("Luas Lingkaran: %.2f\n", hasilLuas)

	// keliling, luas := minitask1.GetKelilingAndLuas(7)
	// fmt.Printf("Keliling = %.2f\nLuas= %.2f", keliling, luas)

	// minitask2.GetTriangle(2)

	// minitask3.InsertNumber()

	// data := models.TampilinData()
	// fmt.Println(data)

	// fmt.Printf("Nama:%s\n", profil.nama)
	// fmt.Printf("Foto:%s\n", profil.foto)
	// fmt.Printf("Email:%s\n", profil.email)
	// fmt.Printf("Umur:%d\n", profil.umur)
	// fmt.Printf("NomorTelepon:	%s\n", profil.nomorTelepon)
	// fmt.Printf("Status Menikah:%t\n", profil.statusMenikah)
	// fmt.Printf("Umur:%d\n", profil.umur)
	// for i, edu := range profil.riwayatPendidikan {
	// 	fmt.Printf("%d. %s - Jurasan %s\n", i+1, edu.namaSekolah, edu.jurusan)
	// }

}
