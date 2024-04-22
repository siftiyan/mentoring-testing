package main

import (
	"errors"
	"fmt"
)

func PembayaranBarang(hargaTotal float64, metodePembayaran string, cicilan bool) error {
	if hargaTotal <= 0 {
		return errors.New("harga tidak bisa nol")
	}

	metodeValid := map[string]bool{"cod": true, "transfer": true, "debit": true, "credit": true, "gerai": true}
	if _, ok := metodeValid[metodePembayaran]; !ok {
		return errors.New("metode tidak dikenali")
	}

	diskon, biaya := hitungPenyesuaian(hargaTotal, metodePembayaran)
	totalSetelahPenyesuaian := hargaTotal - diskon + biaya

	if cicilan {
		if metodePembayaran != "credit" {
			return errors.New("cicilan hanya bisa dengan metode credit")
		}
		if totalSetelahPenyesuaian < 500000 {
			return errors.New("cicilan tidak memenuhi syarat")
		}
	} else {
		if metodePembayaran == "credit" {
			return errors.New("credit harus dicicil")
		}
	}

	return nil
}

// ini uji coba penambahan kasus dari tugas yang diberikan
func hitungPenyesuaian(total float64, metode string) (diskon, biaya float64) {
	// kondisi ini digunakan ketika pengguna memilih credit maka secara lumrahnya/umunnya terdapat fee didalamnya
	if metode == "credit" {
		biaya = 5000
	}
	// kondisi ini digunakan ketika pengguna mengluarkan transaksi lebih dari 1 juta (estimasi diskon toko oren biasanya)
	if total >= 1000000 {
		diskon = total * 0.05
	}

	return
}

func prosesTransaksi(transaksi []struct {
	hargaTotal float64
	metode     string
	cicilan    bool
}) {
	for _, trx := range transaksi {
		err := PembayaranBarang(trx.hargaTotal, trx.metode, trx.cicilan)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		} else {
			fmt.Println("Transaksi berhasil!")
		}
	}
}

func main() {
	transaksi := []struct {
		hargaTotal float64
		metode     string
		cicilan    bool
	}{
		{500000, "credit", true},
	}
	prosesTransaksi(transaksi)
}
