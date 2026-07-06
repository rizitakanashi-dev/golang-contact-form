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

func PesanAdminHandler(w http.ResponseWriter, r *http.Request) {
	query := `SELECT id, nama, email, pesan, waktu_kirim FROM pesan_kontak ORDER BY waktu_kirim DESC`
	rows, err := config.DB.Query(query)
	if err != nil {
		http.Error(w, "gagal mengambil string dari database: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var listPesan []models.KontakPesan
	for rows.Next() {
		var p models.KontakPesan

		err := rows.Scan(&p.ID, &p.Nama, &p.Email, &p.Pesan, &p.WaktuKirim)
		if err != nil {
			http.Error(w, "gagal memindai data: "+err.Error(), http.StatusInternalServerError)
			return
		}
		listPesan = append(listPesan, p)
	}

	targetFile := filepath.Join("web", "templates", "admin.html")
	tmpl, err := template.ParseFiles(targetFile)
	if err != nil {
		http.Error(w, "gagal memuat template admin: "+err.Error(), http.StatusInternalServerError)
		return
	}

	dataBalik := models.HalamanAdminData{
		SemuaPesan: listPesan,
	}
	tmpl.Execute(w, dataBalik)
}
