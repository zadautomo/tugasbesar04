package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Konten struct {
	Judul       string
	TanggalPost string
	Kategori    string
	Interaksi   int
}

var reader = bufio.NewReader(os.Stdin)

func inputString(prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func parseTanggal(tanggal string) time.Time {
	parsed, err := time.Parse("02-01-2006", tanggal)
	if err != nil {
		fmt.Println("Format tanggal salah. Gunakan DD-MM-YYYY.")
		return time.Time{}
	}
	return parsed
}

func formatTanggal(t time.Time) string {
	return t.Format("02-01-2006")
}

func tampilkanMenu() {
	fmt.Println("\nMenu:")
	fmt.Println("1. Tambah Konten")
	fmt.Println("2. Lihat Daftar Konten")
	fmt.Println("3. Edit Konten")
	fmt.Println("4. Hapus Konten")
	fmt.Println("5. Cari Konten (Binary Search)")
	fmt.Println("6. Urutkan Konten (Selection Sort - Tanggal)")
	fmt.Println("7. Urutkan Konten (Insertion Sort - Interaksi)")
	fmt.Println("8. Tampilkan Konten Engagement Tertinggi di Periode")
	fmt.Println("9. Keluar")
	fmt.Print("Pilih menu: ")
}

func tambahKonten(daftar *[]Konten) {
	judul := inputString("Judul Konten: ")
	tanggal := inputString("Tanggal Posting (DD-MM-YYYY): ")
	kategori := inputString("Kategori Konten: ")
	var interaksi int
	fmt.Print("Jumlah Interaksi: ")
	fmt.Scanln(&interaksi)

	*daftar = append(*daftar, Konten{judul, tanggal, kategori, interaksi})
	fmt.Println("Konten berhasil ditambahkan.")
}

func lihatKonten(daftar []Konten) {
	fmt.Println("\nDaftar Konten:")
	fmt.Println(strings.Repeat("=", 90))
	fmt.Printf("| %-3s | %-10s | %-13s | %-35s | %-10s |\n", "No", "Tanggal", "Kategori", "Judul", "Interaksi")
	fmt.Println(strings.Repeat("-", 90))
	for i, k := range daftar {
		fmt.Printf("| %-3d | %-10s | %-13s | %-35s | %-10d |\n", i+1, k.TanggalPost, k.Kategori, k.Judul, k.Interaksi)
	}
	fmt.Println(strings.Repeat("=", 90))
}

func editKonten(daftar *[]Konten) {
	judul := inputString("Masukkan judul konten yang ingin diedit: ")
	for i, k := range *daftar {
		if strings.ToLower(k.Judul) == strings.ToLower(judul) {
			(*daftar)[i].Judul = inputString("Judul Baru: ")
			(*daftar)[i].TanggalPost = inputString("Tanggal Baru (DD-MM-YYYY): ")
			(*daftar)[i].Kategori = inputString("Kategori Baru: ")
			fmt.Print("Interaksi Baru: ")
			fmt.Scanln(&(*daftar)[i].Interaksi)
			fmt.Println("Konten berhasil diperbarui.")
			return
		}
	}
	fmt.Println("Konten tidak ditemukan.")
}

func hapusKonten(daftar *[]Konten) {
	judul := inputString("Masukkan judul konten yang ingin dihapus: ")
	for i, k := range *daftar {
		if strings.ToLower(k.Judul) == strings.ToLower(judul) {
			*daftar = append((*daftar)[:i], (*daftar)[i+1:]...)
			fmt.Println("Konten berhasil dihapus.")
			return
		}
	}
	fmt.Println("Konten tidak ditemukan.")
}

func selectionSortKategori(daftar []Konten) {
	n := len(daftar)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if strings.ToLower(daftar[j].Kategori) < strings.ToLower(daftar[minIdx].Kategori) {
				minIdx = j
			}
		}
		daftar[i], daftar[minIdx] = daftar[minIdx], daftar[i]
	}
}

func binarySearchKategori(daftar []Konten, target string) {
	selectionSortKategori(daftar)
	target = strings.ToLower(target)
	i, j := 0, len(daftar)-1
	for i <= j {
		mid := (i + j) / 2
		kat := strings.ToLower(daftar[mid].Kategori)
		if kat == target {
			fmt.Printf("\nHasil ditemukan: %s | %s | %s | Interaksi: %d\n",
				daftar[mid].TanggalPost, daftar[mid].Kategori, daftar[mid].Judul, daftar[mid].Interaksi)
			return
		} else if kat < target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	fmt.Println("Kategori tidak ditemukan.")
}

func selectionSortTanggal(daftar []Konten) {
	n := len(daftar)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if parseTanggal(daftar[j].TanggalPost).Before(parseTanggal(daftar[minIdx].TanggalPost)) {
				minIdx = j
			}
		}
		daftar[i], daftar[minIdx] = daftar[minIdx], daftar[i]
	}
	lihatKonten(daftar)
}

func insertionSortInteraksi(daftar []Konten) {
	for i := 1; i < len(daftar); i++ {
		key := daftar[i]
		j := i - 1
		for j >= 0 && daftar[j].Interaksi > key.Interaksi {
			daftar[j+1] = daftar[j]
			j--
		}
		daftar[j+1] = key
	}
	lihatKonten(daftar)
}

func tampilkanEngagementTertinggi(daftar []Konten) {
	awalStr := inputString("Masukkan tanggal awal (DD-MM-YYYY): ")
	akhirStr := inputString("Masukkan tanggal akhir (DD-MM-YYYY): ")

	awal := parseTanggal(awalStr)
	akhir := parseTanggal(akhirStr)

	max := -1
	var topKonten Konten
	for _, k := range daftar {
		tanggal := parseTanggal(k.TanggalPost)
		if !tanggal.Before(awal) && !tanggal.After(akhir) {
			if k.Interaksi > max {
				max = k.Interaksi
				topKonten = k
			}
		}
	}

	if max == -1 {
		fmt.Println("Tidak ada konten dalam periode tersebut.")
	} else {
		fmt.Printf("Konten engagement tertinggi: %s | %s | Interaksi: %d\n", topKonten.Judul, topKonten.TanggalPost, topKonten.Interaksi)
	}
}

func main() {
	daftarKonten := []Konten{
		{"Tutorial Go Dasar", "01-05-2025", "Edukasi", 1200},
		{"Vlog Harian", "04-03-2025", "Hiburan", 980},
		{"Review Gadget", "03-08-2025", "Teknologi", 1450},
		{"Resep Masakan Simple", "04-06-2025", "Kuliner", 1100},
		{"Workout di Rumah", "04-07-2025", "Kesehatan", 1340},
	}

	for {
		tampilkanMenu()
		var pilihan string
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			tambahKonten(&daftarKonten)
		case "2":
			lihatKonten(daftarKonten)
		case "3":
			editKonten(&daftarKonten)
		case "4":
			hapusKonten(&daftarKonten)
		case "5":
			kat := inputString("Masukkan kategori: ")
			binarySearchKategori(daftarKonten, kat)
		case "6":
			selectionSortTanggal(daftarKonten)
		case "7":
			insertionSortInteraksi(daftarKonten)
		case "8":
			tampilkanEngagementTertinggi(daftarKonten)
		case "9":
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}