package main

import (
	"html/template"
	"log"
	"net/http"
)

// Struct untuk nangkep data resi lu
type TrackingData struct {
	Resi     string
	Status   string
	Penerima string
	Layanan  string
	Isi      string
}

func main() {
	http.HandleFunc("/track", func(w http.ResponseWriter, r *http.Request) {
		resi := r.URL.Query().Get("nomor_resi")
		var data *TrackingData

		// Hardcode dummy data biar sama persis kayak screenshot lu sebelumnya
		if resi == "RB-IT-2026" {
			data = &TrackingData{
				Resi:     "RB-IT-2026",
				Status:   "DELIVERED",
				Penerima: "Tim Seleksi - Ruang Belajar IT",
				Layanan:  "Next Day (Anteraja)",
				Isi:      "Berkas Dokumen Pendaftaran Internship",
			}
		}

		// Render ke HTML
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "HTML-nya belom lu bikin, Nan! "+err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, data)
	})

	// Kalau buka localhost:8080 doang, langsung diarahin ke halaman track
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/track", http.StatusSeeOther)
	})

	log.Println("Server Golang jalan mulus di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}