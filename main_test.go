package main

import (
	"testing"
)

func TestPembayaranBarang(t *testing.T) {
	type args struct {
		hargaTotal float64
		metode     string
		dicicil    bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		errMsg  string
	}{

		{
			name:    "Total Negatif",
			args:    args{hargaTotal: 0, metode: "credit", dicicil: true},
			wantErr: true,
			errMsg:  "harga tidak bisa nol",
		},
		{
			name:    "Metode Tidak Dikenal",
			args:    args{hargaTotal: 100000, metode: "bitcoin", dicicil: false},
			wantErr: true,
			errMsg:  "metode tidak dikenali",
		},
		{
			name:    "Kredit Tanpa Cicilan",
			args:    args{hargaTotal: 200000, metode: "credit", dicicil: false},
			wantErr: true,
			errMsg:  "credit harus dicicil",
		},
		{
			name:    "Cicilan Tidak Memenuhi Syarat",
			args:    args{hargaTotal: 400000, metode: "credit", dicicil: true},
			wantErr: true,
			errMsg:  "cicilan tidak memenuhi syarat",
		},
		{
			name:    "Transaksi Berhasil dengan Debit",
			args:    args{hargaTotal: 200000, metode: "debit", dicicil: false},
			wantErr: false,
		},
		{
			name:    "Transaksi Kredit Berhasil untuk Nilai Tinggi",
			args:    args{hargaTotal: 1200000, metode: "credit", dicicil: true},
			wantErr: false,
		},

		{
			name:    "Cicilan dengan Metode Non-Kredit",
			args:    args{hargaTotal: 500000, metode: "debit", dicicil: true},
			wantErr: true,
			errMsg:  "cicilan hanya bisa dengan metode credit",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := PembayaranBarang(tt.args.hargaTotal, tt.args.metode, tt.args.dicicil)
			if (err != nil) != tt.wantErr {
				t.Errorf("%s error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			if err != nil && err.Error() != tt.errMsg {
				t.Errorf("%s pesan kesalahan yang diharapkan = %s, tapi mendapat = %s", tt.name, tt.errMsg, err.Error())
			}
		})
	}
}
