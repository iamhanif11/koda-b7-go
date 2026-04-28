package minitask8

import (
	"errors"
	"fmt"
)

type Payment interface {
	Pay(amount []int) error
}

type BankPayment struct {
	BankName string
}

func (b BankPayment) Pay(amount []int) error {
	total := 0
	for _, amt := range amount {
		if amt <= 0 {
			return errors.New("Pembayaran tidak boleh 0")
		}
		total += amt
	}
	fmt.Printf("[Bank%s] berhasil bayar Rp. %d\n", b.BankName, total)
	return nil
}

type OnlinePayment struct {
	Platform string
}

func (o OnlinePayment) Pay(amount []int) error {
	total := 0
	for _, amt := range amount {
		if amt <= 0 {
			return errors.New("Pembayaran tidak boleh 0")
		}
		total += amt
	}
	fmt.Printf("[Platform %s] berhasil bayar Rp. %d\n", o.Platform, total)
	return nil
}

type FictionPayment struct {
	TotalPayment []int
}

func (f *FictionPayment) Pay(amount []int) error {
	total := 0
	for _, amt := range amount {
		if amt <= 0 {
			return errors.New("Pemabayaran tidak boleh 0")
		}
		total += amt
	}
	f.TotalPayment = append(f.TotalPayment, total)
	return nil
}
