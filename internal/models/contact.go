package models

import "time"

type HalamanData struct {
	Success bool
	Nama    string
}

type KontakPesan struct {
	ID         int
	Nama       string
	Email      string
	Pesan      string
	WaktuKirim time.Time
}

type HalamanAdminData struct {
	SemuaPesan []KontakPesan
}
