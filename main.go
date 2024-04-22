package main

import (
	"errors"
	"fmt"
)

const (
	tax = 10
	adm = 2000
)

func main() {
	total, err := hargaTotal(15000, 10000, 2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Harga total yang harus dibayarkan adalah %v\n", total)
	fmt.Println(PembayaranBarang(45000, "gak bayar", false))
}
func hargaTotal(hargaItem float64, ongkir float64, qty int) (float64, error) {
	if hargaItem <= 0 {
		return 0, errors.New("Harga barang tidak boleh nol")
	}
	if qty <= 0 {
		return 0, errors.New("Jumlah barang tidak boleh nol")
	}
	if ongkir <= 0 {
		return 0, errors.New("harga ongkir tidak boleh nol")
	}
	hargatotalItem := hargaItem * float64(qty)
	hargaSetelahOngkir := hargatotalItem + ongkir
	pajak := hargatotalItem * tax / 100
	total := hargaSetelahOngkir + pajak + adm

	return total, nil
}

func PembayaranBarang(hargaTotal float64, metodePembayaran string, dicicil bool) error {
	if hargaTotal <= 0 {
		return errors.New("harga tidak bisa nol")
	}

	switch metodePembayaran {
	case "cod", "transfer", "debit", "credit", "gerai":
		if dicicil {
			if metodePembayaran != "credit" || hargaTotal <= 500000 {
				return errors.New("cicilan tidak memenuhi syarat")
			}
		} else {
			if metodePembayaran == "credit" {
				return errors.New("credit harus dicicil")
			}
		}
	default:
		return errors.New("metode tidak dikenali")
	}
	return nil
}
