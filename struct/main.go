package main

import "fmt"

func main() {
	type Peserta struct {
		NIM         string
		namaLengkap string
		NoHp        int
		IPK         float64
		Hadir       bool
	}

	Maira := Peserta{
		NIM:         "210170077",
		namaLengkap: "Humaira",
		NoHp:        82267844887,
		IPK:         3.23,
		Hadir:       true,
	}

	fmt.Printf(`
	Saya adalah seorang mahasiswa yang bernama %s
	memiliki Nim %s berikut adalah no hp yang bisa dihubungi %d
	saya memiliki ipk %.2f dan kehadiran saya adalah %t 
	  `, Maira.namaLengkap, Maira.NIM, Maira.NoHp, Maira.IPK, Maira.Hadir)
}
