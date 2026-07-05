package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"

	"golang-contact-form/internal/config" // Import package config untuk akses DB
	"golang-contact-form/internal/models"
)

func KontakHandler(w http.ResponseWriter, r *http.Request) {
	// Arahkan path ke folder web/templates/form.html
	targetFile := filepath.Join("web", "templates", "form.html")

	tmpl, err := template.ParseFiles(targetFile)
	if err != nil {
		http.Error(w, "Gagal memuat template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodGet {
		tmpl.Execute(w, nil)
		return
	}

	if r.Method == http.MethodPost {
		nama := r.FormValue("nama")
		email := r.FormValue("email")
		pesan := r.FormValue("pesan")

		// Gunakan tanda tanya (?) sebagai placeholder kueri untuk MariaDB/MySQL
		query := `INSERT INTO pesan_kontak (nama, email, pesan) VALUES (?, ?, ?)`

		// Eksekusi kueri menggunakan objek DB global dari package config
		_, err := config.DB.Exec(query, nama, email, pesan)
		if err != nil {
			http.Error(w, "Gagal menyimpan pesan ke MariaDB: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Kirim status sukses kembali ke HTML
		dataBalik := models.HalamanData{
			Success: true,
			Nama:    nama,
		}
		tmpl.Execute(w, dataBalik)
	}
}
