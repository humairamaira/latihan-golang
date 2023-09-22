package main

import (
	"fmt"
	"os"
)

// Data barang
var daftarBarang = map[string]map[string]int{
	"Minuman": {
		"Teh":  5000,
		"Kopi": 10000,
		"Jus":  5000,
	},
	"Makanan": {
		"Seblak": 10000,
		"Sate":   25000,
		"Bakso":  15000,
	},
}

// Variabel global untuk menyimpan hasil total/invoice
var invoice string

func main() {
	keranjang := make(map[string]map[string]int)

	for {
		fmt.Println("\nAplikasi Kasir")
		fmt.Println("1. Tampilkan Daftar Barang")
		fmt.Println("2. Tambahkan Barang ke Keranjang")
		fmt.Println("3. Lihat Keranjang Belanja")
		fmt.Println("4. Tampilkan Total Belanja dan Total Biaya")
		fmt.Println("5. Keluar")

		var pilihan string
		fmt.Print("Pilih menu (1/2/3/4/5): ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			tampilkanDaftarBarang()
		case "2":
			tambahkanBarangKeKeranjang(keranjang)
		case "3":
			lihatKeranjangBelanja(keranjang)
		case "4":
			tampilkanTotalBelanjaDanBiaya(keranjang)
		case "5":
			fmt.Println("Terima kasih telah menggunakan aplikasi kasir.")
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih menu yang benar.")
		}
	}
}

func tampilkanDaftarBarang() {
	fmt.Println("Daftar Barang:")
	for kategori, barang := range daftarBarang {
		fmt.Printf("\nKategori: %s\n", kategori)
		for item, harga := range barang {
			fmt.Printf("%s: $%d\n", item, harga)
		}
	}
}

func tambahkanBarangKeKeranjang(keranjang map[string]map[string]int) {
	var kategori, item string
	var jumlah int

	fmt.Print("Masukkan nama kategori: ")
	fmt.Scanln(&kategori)
	fmt.Print("Masukkan nama barang: ")
	fmt.Scanln(&item)
	fmt.Print("Masukkan jumlah barang yang dibeli: ")
	fmt.Scanln(&jumlah)

	if _, ok := daftarBarang[kategori]; ok {
		if _, ok := keranjang[kategori]; !ok {
			keranjang[kategori] = make(map[string]int)
		}
		keranjang[kategori][item] += jumlah
		fmt.Printf("%d %s telah ditambahkan ke keranjang.\n", jumlah, item)
	} else {
		fmt.Println("Kategori atau barang tidak valid.")
	}
}

func lihatKeranjangBelanja(keranjang map[string]map[string]int) {
	fmt.Println("\nKeranjang Belanja:")
	for kategori, barang := range keranjang {
		for item, jumlah := range barang {
			fmt.Printf("%s - %s - %d pcs\n", kategori, item, jumlah)
		}
	}
}

func tampilkanTotalBelanjaDanBiaya(keranjang map[string]map[string]int) {
	var totalBiaya int

	for kategori, barang := range keranjang {
		for item, jumlah := range barang {
			if harga, ok := daftarBarang[kategori][item]; ok {
				totalBiaya += harga * jumlah
			}
		}
	}

	invoice = fmt.Sprintf("\nTotal Belanja: $%d\n", totalBiaya)

	// Simpan hasil ke dalam file "invoice.txt"
	filename := "invoice.txt"
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Gagal membuka file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(invoice)
	if err != nil {
		fmt.Println("Gagal menyimpan hasil ke file:", err)
		return
	}

	fmt.Printf("Hasil telah disimpan dalam file '%s'.\n", filename)
}
