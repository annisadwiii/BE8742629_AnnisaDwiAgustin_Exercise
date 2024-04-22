package main

import (
	"testing"
)

func Test_hargaTotal(t *testing.T) {
	type args struct {
		hargaItem float64
		ongkir    float64
		qty       int
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "gagalin harga item",
			args: args{
				hargaItem: 0,
				ongkir:    1,
				qty:       1,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "gagalin kuantitas",
			args: args{
				hargaItem: 1,
				ongkir:    1,
				qty:       0,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "gagalin ongkir",
			args: args{
				hargaItem: 1,
				ongkir:    0,
				qty:       1,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "sukses",
			args: args{
				hargaItem: 15000,
				ongkir:    10000,
				qty:       2,
			},
			want:    45000,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := hargaTotal(tt.args.hargaItem, tt.args.ongkir, tt.args.qty)
			if (err != nil) != tt.wantErr {
				t.Errorf("hargaTotal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("hargaTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPembayaranBarang(t *testing.T) {
	type args struct {
		hargaTotal       float64
		metodePembayaran string
		dicicil          bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "gagalin harga total",
			args: args{
				hargaTotal:       0,
				metodePembayaran: "cod",
				dicicil:          true,
			},
			wantErr: true,
		},
		{
			name: "gagalin metode pembayaran",
			args: args{
				hargaTotal:       670000,
				metodePembayaran: "gak bayar",
				dicicil:          true,
			},
			wantErr: true,
		},
		{
			name: "gagalin di metode pembayaran bukan kredit tapi dicicil",
			args: args{
				hargaTotal:       670000,
				metodePembayaran: "cod",
				dicicil:          true,
			},
			wantErr: true,
		},
		{
			name: "gagalin metode pembayaran kredit, dicicil, tapi harga total kurang dari 500000",
			args: args{
				hargaTotal:       450000,
				metodePembayaran: "credit",
				dicicil:          true,
			},
			wantErr: true,
		},
		{
			name: "gagalin metode pembayaran kredit, tapi tidak dicicil",
			args: args{
				hargaTotal:       670000,
				metodePembayaran: "credit",
				dicicil:          false,
			},
			wantErr: true,
		},
		{
			name: "suksesin metode pembayaran kredit, dicicl, harga total lebih dari 500000",
			args: args{
				hargaTotal:       670000,
				metodePembayaran: "credit",
				dicicil:          true,
			},
			wantErr: false,
		},
		{
			name: "suksesin metode pembayaran cod, tidak dicicil",
			args: args{
				hargaTotal:       340000,
				metodePembayaran: "cod",
				dicicil:          false,
			},
			wantErr: false,
		},
		{
			name: "suksesin metode pembayaran transfer, tidak dicicil",
			args: args{
				hargaTotal:       340000,
				metodePembayaran: "transfer",
				dicicil:          false,
			},
			wantErr: false,
		},
		{
			name: "suksesin metode pembayaran debit, tidak dicicil",
			args: args{
				hargaTotal:       340000,
				metodePembayaran: "debit",
				dicicil:          false,
			},
			wantErr: false,
		},
		{
			name: "suksesin metode pembayaran gerai, tidak dicicil",
			args: args{
				hargaTotal:       340000,
				metodePembayaran: "gerai",
				dicicil:          false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PembayaranBarang(tt.args.hargaTotal, tt.args.metodePembayaran, tt.args.dicicil); (err != nil) != tt.wantErr {
				t.Errorf("PembayaranBarang() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
