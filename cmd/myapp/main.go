package main

import (
	"fmt"
	"net/http"

	"golang-contact-form/internal/config" // 1. Pastikan mengimpor package config
	"golang-contact-form/internal/handlers"
)

func main() {
	// 2. Panggil fungsi inisialisasi database sebelum server jalan
	config.InitDB()

	// 3. Pastikan koneksi database ditutup dengan aman saat aplikasi dimatikan
	defer config.DB.Close()

	http.HandleFunc("/kontak", handlers.KontakHandler)
	http.HandleFunc("/admin/pesan", handlers.PesanAdminHandler)

	fmt.Println("Aplikasi berjalan di: http://localhost:8080/kontak")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("gagal menjalankan server", err)
	}
}
