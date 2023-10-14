package main

import "fmt"

func main() {
	// Masukkan nilai tagihan
	var tagihan float64
	fmt.Print("Nilai tagihan : ")
	fmt.Scan(&tagihan)

	// Menghitung tip
	var tip float64
	if tagihan >= 50 && tagihan <= 300 {
		tip = tagihan * 0.15 // Tip 15% jika tagihan antara 50 dan 300
	} else {
		tip = tagihan * 0.20 // Tip 20% untuk kasus lain
	}

	// Menghitung total nilai
	total := tagihan + tip

	// Menampilkan hasil
	fmt.Printf("Tagihannya %.2f, tipnya %.2f, dan total nilainya %.2f\n", tagihan, tip, total)
}
