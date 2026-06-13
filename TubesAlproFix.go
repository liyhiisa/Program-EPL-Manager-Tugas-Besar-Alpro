package main

import (
    "fmt"
    "strings"
)

// --- KONSTANTA ---
const NMAX = 20
const MAX_JADWAL = 380

// --- STRUKTUR DATA ---
type Klub struct {
	Nama string
	Main, Menang, Seri, Kalah int
	GolMemasukkan, GolKemasukan, SelisihGol, Poin int
}

type Pertandingan struct {
	Pekan int
	Home, Away string
	GolHome, GolAway int
	SudahDimainkan bool
}

type TabKlub [NMAX]Klub
type TabJadwal [MAX_JADWAL]Pertandingan

// --- VARIABEL GLOBAL ---
var liga TabKlub
var nKlub int
var jadwal TabJadwal
var nJadwal int
var jadwalDibuat bool

// main adalah titik masuk (entry point) dari program.
// Fungsi ini menginisialisasi data awal dan menjalankan loop menu interaktif 
// agar pengguna dapat memilih fitur-fitur aplikasi (klasemen, jadwal, dll)
// sampai pengguna memilih untuk keluar (opsi 0).
func main() {
	initDataDummy()
	jadwalDibuat = false

	var pilihan int
	selesai := false // Flag pengganti break

	for !selesai {
		fmt.Println("\n=== APLIKASI EPL MANAGER (BERGER SYSTEM) ===")
		fmt.Println("1. Tampilkan Klasemen")
		fmt.Println("2. Kelola Klub (Tambah/Ubah/Hapus)")
		fmt.Println("3. Buat Jadwal Pertandingan (Tabel Berger)")
		fmt.Println("4. Catat Hasil Pertandingan per Pekan")
		fmt.Println("5. Edit Hasil Pertandingan")
		fmt.Println("6. Hapus Hasil Pertandingan")
		fmt.Println("7. Lihat Seluruh Jadwal (Generator Tabel Berger)")
		fmt.Println("8. Statistik Gol Klub")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			tampilKlasemen()
		} else if pilihan == 2 {
			menuKelolaKlub()
		} else if pilihan == 3 {
			buatJadwalBerger()
		} else if pilihan == 4 {
			catatPertandinganPerPekan()
		} else if pilihan == 5 {
			editHasilPertandingan()
		} else if pilihan == 6 {
			hapusHasilPertandingan()
		} else if pilihan == 7 {
			tampilSeluruhJadwal()
		} else if pilihan == 8 {
			menuStatistikGol()
		} else if pilihan == 0 {
			fmt.Println("Terima kasih telah menggunakan EPL Manager!")
			selesai = true
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// ==========================================
// SUBPROGRAM INISIALISASI
// ==========================================

// initDataDummy berfungsi untuk mengisi array liga dengan 10 tim awal (dummy).
// Fungsi ini berguna agar pengguna tidak perlu memasukkan tim satu per satu 
// saat aplikasi baru pertama kali dijalankan.
func initDataDummy() {
	liga[0] = Klub{"ARS", 0, 0, 0, 0, 0, 0, 0, 0}
	liga[1] = Klub{"MCI", 0, 0, 0, 0, 0, 0, 0, 0}
	liga[2] = Klub{"LIV", 0, 0, 0, 0, 0, 0, 0, 0}
	liga[3] = Klub{"CHE", 0, 0, 0, 0, 0, 0, 0, 0}
	liga[4] = Klub{"EVE", 0, 0, 0, 0, 0, 0, 0, 0}
	liga[5] = Klub{"FUL", 0, 0, 0, 0, 0, 0, 0, 0}
	liga[6] = Klub{"IPS", 0, 0, 0, 0, 0, 0, 0, 0}
	liga[7] = Klub{"LEI", 0, 0, 0, 0, 0, 0, 0, 0}
	liga[8] = Klub{"MUN", 0, 0, 0, 0, 0, 0, 0, 0}
	liga[9] = Klub{"TOT", 0, 0, 0, 0, 0, 0, 0, 0}
	nKlub = 10
}

// ==========================================
// SUBPROGRAM JADWAL & PERTANDINGAN
// ==========================================

// resetStatistikKlub mengembalikan semua nilai statistik setiap klub
// (Main, Menang, Seri, Kalah, Poin, Selisih Gol, dll) kembali menjadi 0.
// Biasanya dipanggil saat jadwal di-reset atau dibuat ulang.
func resetStatistikKlub() {
	i := 0
	for i < nKlub {
		liga[i].Main = 0
		liga[i].Menang = 0
		liga[i].Seri = 0
		liga[i].Kalah = 0

		liga[i].GolMemasukkan = 0
		liga[i].GolKemasukan = 0

		liga[i].SelisihGol = 0
		liga[i].Poin = 0
		i++
	}
}

// adaPertandinganDimainkan mengecek apakah di dalam array jadwal
// terdapat setidaknya satu pertandingan yang sudah memiliki hasil/dimainkan.
// Mengembalikan nilai true jika ada, dan false jika belum ada sama sekali.
func adaPertandinganDimainkan() bool {
	i := 0
	for i < nJadwal {
		if jadwal[i].SudahDimainkan {
			return true
		}
		i++
	}
	return false
}

// buatJadwalBerger adalah fungsi krusial yang meng-generate jadwal pertandingan
// format Double Round-Robin (Tandang-Kandang) menggunakan Algoritma Tabel Berger.
// Syarat mutlaknya adalah jumlah klub harus genap.
func buatJadwalBerger() {
	if jadwalDibuat {
		var konfirmasi string
		if adaPertandinganDimainkan() {
			fmt.Println("====================================================")
			fmt.Println("PERINGATAN!")
			fmt.Println("Membuat ulang jadwal akan menghapus:")
			fmt.Println("- Semua hasil pertandingan")
			fmt.Println("- Semua statistik klub")
			fmt.Println("- Semua klasemen")
			fmt.Println("====================================================")

			fmt.Print("Lanjutkan? (Y/T): ")
			fmt.Scan(&konfirmasi)

			if konfirmasi != "Y" && konfirmasi != "y" {
				fmt.Println("Pembuatan jadwal dibatalkan.")
				return
			}
			resetStatistikKlub()
			nJadwal = 0
			fmt.Println("Data lama berhasil dihapus.")
			fmt.Println("Membuat jadwal baru...")

		} else {
			fmt.Println("Jadwal sudah pernah dibuat.")
			fmt.Print("Buat ulang jadwal? (Y/T): ")
			fmt.Scan(&konfirmasi)

			if konfirmasi != "Y" && konfirmasi != "y" {
				fmt.Println("Pembuatan jadwal dibatalkan.")
				return
			}
			nJadwal = 0
			fmt.Println("Membuat jadwal baru...")
		}
	}

	if nKlub%2 != 0 {
		fmt.Println("Syarat Tabel Berger: Jumlah klub harus genap!")
		return
	}
	nJadwal = 0

	var rotasi [NMAX]int

	i := 0
	for i < nKlub {
		rotasi[i] = i
		i++
	}

	// ==========================
	// PUTARAN 1 (Setengah Musim Pertama)
	// ==========================
	pekan := 0
	for pekan < nKlub-1 {
		j := 0
		for j < nKlub/2 {
			idxHome := rotasi[j]
			idxAway := rotasi[nKlub-1-j]

			jadwal[nJadwal].Pekan = pekan + 1
			jadwal[nJadwal].Home = liga[idxHome].Nama
			jadwal[nJadwal].Away = liga[idxAway].Nama
			jadwal[nJadwal].GolHome = 0
			jadwal[nJadwal].GolAway = 0
			jadwal[nJadwal].SudahDimainkan = false
			nJadwal++
			j++
		}
		// Rotasi tim (elemen pertama/index 0 tetap di tempat)
		temp := rotasi[nKlub-1]
		k := nKlub - 1
		for k > 1 {
			rotasi[k] = rotasi[k-1]
			k--
		}
		rotasi[1] = temp
		pekan++
	}

	// ==========================
	// PUTARAN 2 (Setengah Musim Kedua)
	// ==========================
	// Membalikkan posisi Home dan Away dari putaran 1
	totalPutaran1 := nJadwal
	l := 0
	for l < totalPutaran1 {
		jadwal[nJadwal].Pekan = jadwal[l].Pekan + (nKlub - 1)
		jadwal[nJadwal].Home = jadwal[l].Away
		jadwal[nJadwal].Away = jadwal[l].Home
		jadwal[nJadwal].GolHome = 0
		jadwal[nJadwal].GolAway = 0
		jadwal[nJadwal].SudahDimainkan = false
		nJadwal++
		l++
	}

	jadwalDibuat = true
	fmt.Println("Jadwal Double Round-Robin berhasil dibuat!")
	fmt.Println("Silakan buka menu 7 untuk melihat jadwal.")
}

// tampilSeluruhJadwal mencetak seluruh daftar pertandingan dari pekan pertama
// hingga pekan terakhir yang telah dibuat oleh algoritma Berger.
func tampilSeluruhJadwal() {
	if !jadwalDibuat {
		fmt.Println("Jadwal belum dibuat! Silakan pilih menu 3 terlebih dahulu.")
		return
	}

	fmt.Println("\n=======================================================")
	fmt.Println("          TABEL BERGER (JADWAL LENGKAP SEMUSIM)        ")
	fmt.Println("=======================================================")

	totalPekan := (nKlub - 1) * 2
	lagaPerPekan := nKlub / 2

	pekan := 1
	for pekan <= totalPekan {
		fmt.Printf("\n--- PEKAN %d ---\n", pekan)
		
		idxMulai := (pekan - 1) * lagaPerPekan
		idxSelesai := idxMulai + lagaPerPekan
		
		i := idxMulai
		for i < idxSelesai {
			fmt.Printf("Tuan Rumah: %-4s vs  Tamu: %-4s", jadwal[i].Home, jadwal[i].Away)
			
			if jadwal[i].SudahDimainkan {
				fmt.Printf("  [Skor: %d - %d]\n", jadwal[i].GolHome, jadwal[i].GolAway)
			} else {
				fmt.Printf("  [Belum Main]\n")
			}
			i++
		}
		pekan++
	}
	fmt.Println("\n=======================================================")
}

// catatPertandinganPerPekan meminta input hasil skor dari pengguna 
// untuk semua pertandingan yang ada pada suatu pekan tertentu.
// Fungsi ini juga otomatis meng-update klasemen dan mengecek jika liga sudah berakhir.
func catatPertandinganPerPekan() {
	if !jadwalDibuat {
		fmt.Println("Jadwal belum ada! Silakan pilih menu 3 dulu.")
		return
	}

	var pekanInput int
	totalPekan := (nKlub - 1) * 2
	fmt.Printf("Masukkan Nomor Pekan (1 - %d) [0 untuk kembali]: ", totalPekan)
	fmt.Scan(&pekanInput)
	
	if pekanInput == 0 {
		return
	}

	lagaPerPekan := nKlub / 2
	idxMulai := (pekanInput - 1) * lagaPerPekan
	idxSelesai := idxMulai + lagaPerPekan

	if idxMulai < 0 || idxMulai >= nJadwal {
		fmt.Println("Nomor pekan tidak valid!")
		return
	}

	fmt.Printf("\n--- MENCATAT JADWAL PEKAN %d ---\n", pekanInput)
	
	i := idxMulai
	for i < idxSelesai {
		if jadwal[i].SudahDimainkan {
			fmt.Printf("[SELESAI] Tuan Rumah (%s) %d - %d Tamu (%s)\n", 
				jadwal[i].Home, jadwal[i].GolHome, jadwal[i].GolAway, jadwal[i].Away)
		} else {
			fmt.Printf("[BELUM MAIN] Tuan Rumah: %s vs Tamu: %s\n", jadwal[i].Home, jadwal[i].Away)
			
			var skorHome, skorAway int
			fmt.Printf("Masukkan Skor Tuan Rumah (%s) [Ketik -1 untuk lewat]: ", jadwal[i].Home)
			fmt.Scan(&skorHome)
			
			if skorHome < -1 {
			fmt.Println("Skor tidak boleh negatif.")
				return
			}
			if skorHome != -1 {
				fmt.Printf("Masukkan Skor Tamu (%s): ", jadwal[i].Away)
				fmt.Scan(&skorAway)
				
				if skorAway < 0 {
					fmt.Println("Skor tidak boleh negatif.")
					return
				}

				jadwal[i].GolHome = skorHome
				jadwal[i].GolAway = skorAway
				jadwal[i].SudahDimainkan = true

				updateStatistik(jadwal[i].Home, jadwal[i].Away, skorHome, skorAway)
				fmt.Println("-> Hasil pertandingan tersimpan!")
			}
		}
		i++
	}

	// Mengecek apakah dengan hasil barusan, seluruh kompetisi telah selesai
	if cekSemuaSelesai() {
		insertionSortKlasemen()

		fmt.Println("\n======================================")
		fmt.Println("             JUARA LIGA")
		fmt.Println("======================================")

		fmt.Printf("Klub : %s\n", liga[0].Nama)
		fmt.Printf("Poin : %d\n", liga[0].Poin)
		fmt.Printf("Main : %d\n", liga[0].Main)

		fmt.Printf("M-S-K: %d-%d-%d\n",
			liga[0].Menang,
			liga[0].Seri,
			liga[0].Kalah)

		fmt.Printf("SG   : %+d\n", liga[0].SelisihGol)

		fmt.Println("======================================")
		fmt.Printf("Selamat kepada %s sebagai Juara Liga!\n", liga[0].Nama)
		fmt.Println("======================================")
	}
}

// editHasilPertandingan digunakan untuk merevisi skor pertandingan 
// yang sudah terlanjur dicatat. Fungsi ini akan mengurangkan data lama, 
// kemudian menambahkan data skor yang baru.
func editHasilPertandingan() {
	var pekan int
	totalPekan := (nKlub - 1) * 2

	fmt.Printf("Masukkan Nomor Pekan (1 - %d) [0 untuk kembali]: ", totalPekan)
	fmt.Scan(&pekan)

	if pekan == 0 {
		return
	}

	if pekan < 1 || pekan > totalPekan {
		fmt.Println("Nomor pekan tidak valid.")
		return
	}

	fmt.Printf("\n=== PERTANDINGAN PEKAN %d ===\n", pekan)
	
	var daftar [20]int
	var jumlah int
	jumlah = 0
	i := 0
	
	for i < nJadwal {
		if jadwal[i].Pekan == pekan && jadwal[i].SudahDimainkan {
			fmt.Printf("%d. %s %d - %d %s\n",
				jumlah+1, jadwal[i].Home, jadwal[i].GolHome, jadwal[i].GolAway, jadwal[i].Away)
			daftar[jumlah] = i
			jumlah++
		}
		i++
	}

	if jumlah == 0 {
		fmt.Println("Belum ada hasil pertandingan pada pekan tersebut.")
		return
	}

	var pilih int
	fmt.Printf("Pilih nomor pertandingan (1-%d) [0 untuk kembali] : ", jumlah)
	fmt.Scan(&pilih)
	
	if pilih == 0 {
		return
	}

	if pilih >= 1 && pilih <= jumlah {
		idx := daftar[pilih-1]

		// Batalkan statistik lama terlebih dahulu
		batalkanStatistik(jadwal[idx].Home, jadwal[idx].Away, jadwal[idx].GolHome, jadwal[idx].GolAway)

		var golBaruHome, golBaruAway int
		fmt.Printf("Skor Baru %s : ", jadwal[idx].Home)
		fmt.Scan(&golBaruHome)
		fmt.Printf("Skor Baru %s : ", jadwal[idx].Away)
		fmt.Scan(&golBaruAway)
		
		if golBaruHome < 0 || golBaruAway < 0 {
			fmt.Println("Skor tidak boleh negatif.")
			return
		}

		// Perbarui skor
		jadwal[idx].GolHome = golBaruHome
		jadwal[idx].GolAway = golBaruAway

		// Masukkan statistik baru
		updateStatistik(jadwal[idx].Home, jadwal[idx].Away, golBaruHome, golBaruAway)
		fmt.Println("Hasil pertandingan berhasil diubah.")
	} else {
		fmt.Println("Pilihan pertandingan tidak valid.")
	}
}

// hapusHasilPertandingan berguna untuk mengosongkan kembali skor 
// dari sebuah pertandingan yang telah dicatat, dan mengembalikan statusnya 
// menjadi 'Belum Dimainkan'.
func hapusHasilPertandingan() {
	var pekan int
	fmt.Println("\n=== HAPUS HASIL PERTANDINGAN ===")
	totalPekan := (nKlub - 1) * 2

	fmt.Printf("Masukkan Nomor Pekan (1 - %d) [0 untuk kembali]: ", totalPekan)
	fmt.Scan(&pekan)
	if pekan == 0 {
		return
	}

	if pekan < 1 || pekan > totalPekan {
		fmt.Printf("Pekan tidak valid. Masukkan angka antara 1 sampai %d.\n", totalPekan)
		return
	}

	fmt.Println("\n=== DAFTAR HASIL PEKAN", pekan, "===")
	var daftar [20]int
	var jumlah int
	jumlah = 0
	i := 0
	
	for i < nJadwal {
		if jadwal[i].Pekan == pekan && jadwal[i].SudahDimainkan {
			fmt.Printf("%d. %s %d - %d %s\n",
				jumlah+1, jadwal[i].Home, jadwal[i].GolHome, jadwal[i].GolAway, jadwal[i].Away)
			daftar[jumlah] = i
			jumlah++
		}
		i++
	}

	if jumlah == 0 {
		fmt.Println("Belum ada hasil pertandingan yang dapat dihapus.")
		return
	}

	var pilih int
	fmt.Print("Pilih pertandingan [0 untuk kembali] : ")
	fmt.Scan(&pilih)
	if pilih == 0 {
		return
	}

	if pilih < 1 || pilih > jumlah {
		fmt.Println("Pilihan pertandingan tidak valid.")
		return
	}

	idx := daftar[pilih-1]
	var konfirmasi string
	fmt.Print("Yakin ingin menghapus hasil ini? (Y/T): ")
	fmt.Scan(&konfirmasi)

	if konfirmasi != "Y" && konfirmasi != "y" {
		fmt.Println("Penghapusan dibatalkan.")
		return
	}

	// Kembalikan statistik klub karena laga dibatalkan
	batalkanStatistik(jadwal[idx].Home, jadwal[idx].Away, jadwal[idx].GolHome, jadwal[idx].GolAway)

	// Reset hasil pertandingan
	jadwal[idx].GolHome = 0
	jadwal[idx].GolAway = 0
	jadwal[idx].SudahDimainkan = false

	fmt.Println("Hasil pertandingan berhasil dihapus.")
}

// cekSemuaSelesai melakukan pengecekan apakah seluruh pertandingan 
// di array jadwal sudah dimainkan semua. Jika ya (true), liga dianggap selesai.
func cekSemuaSelesai() bool {
	if !jadwalDibuat || nJadwal == 0 {
		return false
	}
	
	semuaSelesai := true
	i := 0
	for i < nJadwal && semuaSelesai {
		if !jadwal[i].SudahDimainkan {
			semuaSelesai = false 
		}
		i++
	}
	return semuaSelesai
}

// updateStatistik memperbarui statistik kedua tim (main, menang/seri/kalah, gol, poin) 
// berdasarkan hasil jumlah gol yang dimasukkan saat pencatatan pertandingan.
func updateStatistik(namaHome, namaAway string, golHome, golAway int) {
	idxHome := sequentialSearch(namaHome)
	idxAway := sequentialSearch(namaAway)

	if idxHome != -1 && idxAway != -1 {
		liga[idxHome].Main++
		liga[idxAway].Main++

		liga[idxHome].GolMemasukkan += golHome
		liga[idxHome].GolKemasukan += golAway
		liga[idxAway].GolMemasukkan += golAway
		liga[idxAway].GolKemasukan += golHome

		if golHome > golAway {
			liga[idxHome].Menang++
			liga[idxHome].Poin += 3
			liga[idxAway].Kalah++
		} else if golHome < golAway {
			liga[idxAway].Menang++
			liga[idxAway].Poin += 3
			liga[idxHome].Kalah++
		} else {
			liga[idxHome].Seri++
			liga[idxAway].Seri++
			liga[idxHome].Poin += 1
			liga[idxAway].Poin += 1
		}

		liga[idxHome].SelisihGol = liga[idxHome].GolMemasukkan - liga[idxHome].GolKemasukan
		liga[idxAway].SelisihGol = liga[idxAway].GolMemasukkan - liga[idxAway].GolKemasukan
	}
}

// batalkanStatistik mengurangi atau me-revert statistik kedua tim.
// Fungsi ini dipanggil ketika kita menghapus atau mengedit hasil pertandingan
// untuk memastikan klasemen tidak terakumulasi secara salah.
func batalkanStatistik(namaHome, namaAway string, golHome, golAway int) {
	idxHome := sequentialSearch(namaHome)
	idxAway := sequentialSearch(namaAway)

	if idxHome != -1 && idxAway != -1 {
		liga[idxHome].Main--
		liga[idxAway].Main--

		liga[idxHome].GolMemasukkan -= golHome
		liga[idxHome].GolKemasukan -= golAway

		liga[idxAway].GolMemasukkan -= golAway
		liga[idxAway].GolKemasukan -= golHome

		if golHome > golAway {
			liga[idxHome].Menang--
			liga[idxHome].Poin -= 3
			liga[idxAway].Kalah--
		} else if golHome < golAway {
			liga[idxAway].Menang--
			liga[idxAway].Poin -= 3
			liga[idxHome].Kalah--
		} else {
			liga[idxHome].Seri--
			liga[idxAway].Seri--
			liga[idxHome].Poin--
			liga[idxAway].Poin--
		}

		liga[idxHome].SelisihGol = liga[idxHome].GolMemasukkan - liga[idxHome].GolKemasukan
		liga[idxAway].SelisihGol = liga[idxAway].GolMemasukkan - liga[idxAway].GolKemasukan
	}
}

// ==========================================
// SUBPROGRAM KELOLA KLUB
// ==========================================

// menuKelolaKlub menangani logika untuk Menambah (Create),
// Mengubah (Update), atau Menghapus (Delete) klub dalam liga. 
// Jika struktur liga berubah, jadwal harus diulang pembuatannya.
func menuKelolaKlub() {
	var pilihSub int
	fmt.Println("\n-- Kelola Klub --")
	fmt.Println("1. Tambah Klub")
	fmt.Println("2. Ubah Klub")
	fmt.Println("3. Hapus Klub")
	fmt.Println("0. Kembali ke Menu Utama")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilihSub)

	if pilihSub == 0 {
		return
	} else if pilihSub == 1 {
		if nKlub < NMAX {
			var nama string
			fmt.Print("Masukkan nama klub (3 karakter) [0 untuk kembali]: ")
			fmt.Scan(&nama)
			nama = strings.ToUpper(nama)

			if nama == "0" {
				return
			}
			if len(nama) != 3 {
			fmt.Println("Nama klub harus tepat 3 karakter.")
			return
			}
			if sequentialSearch(nama) == -1 {
				liga[nKlub] = Klub{Nama: nama}
				nKlub++
				jadwalDibuat = false // Reset jadwal karena ada tim baru
				fmt.Println("Klub ditambahkan.")
			} else {
				fmt.Println("Klub sudah ada.")
			}
		} else {
			fmt.Println("Liga penuh.")
		}
	} else if pilihSub == 2 {
        var namaLama, namaBaru string
        fmt.Print("Nama klub lama [0 untuk kembali]: ")
		fmt.Scan(&namaLama)
		namaLama = strings.ToUpper(namaLama)

		if namaLama == "0" {
			return
		}

        idx := sequentialSearch(namaLama)
        if idx != -1 {
            fmt.Print("Nama klub baru [0 untuk kembali]: ")
            fmt.Scan(&namaBaru)
			namaBaru = strings.ToUpper(namaBaru)
			
			if namaBaru == "0" {
				return
			}	
			if len(namaBaru) != 3 {
				fmt.Println("Kode klub harus 3 karakter.")
				return
			}	
			if namaBaru == namaLama {
				fmt.Println("Nama baru sama dengan nama lama.")
				return
			}
			if sequentialSearch(namaBaru) != -1 {
				fmt.Println("Nama klub sudah digunakan.")
				return
			}
            liga[idx].Nama = namaBaru

            i := 0
            for i < nJadwal {
                if jadwal[i].Home == namaLama {
                   jadwal[i].Home = namaBaru
                }
                if jadwal[i].Away == namaLama {
                   jadwal[i].Away = namaBaru
                }
                i++
            }
            fmt.Println("Klub diubah.")
        } else {
            fmt.Println("Tidak ditemukan.")
		}
	} else if pilihSub == 3 {
		var nama string
		fmt.Print("Nama klub dihapus [0 untuk kembali]: ")
		fmt.Scan(&nama)
		nama = strings.ToUpper(nama)

		if nama == "0" {
			return
		}
		
		selectionSortNama() 
		idx := binarySearch(nama) // Membutuhkan data terurut
		
		if idx != -1 {
			i := idx
			for i < nKlub-1 {
				liga[i] = liga[i+1]
				i++
			}
			nKlub--
			jadwalDibuat = false // Reset jadwal karena tim berkurang
			fmt.Println("Klub dihapus.")
			insertionSortKlasemen()
		} else {
			fmt.Println("Tidak ditemukan.")
		}
	}
}

// tampilKlasemen mengurutkan data liga berdasarkan poin tertinggi 
// (dan selisih gol) lalu mencetak tabel klasemen sementara/akhir di layar.
func tampilKlasemen() {
	insertionSortKlasemen()
	fmt.Println("\n=========================================================================")
	fmt.Println("Pos | Klub | Main | Menang | Seri | Kalah | GM | GK | SG | Poin")
	fmt.Println("=========================================================================")
	i := 0
	for i < nKlub {
		k := liga[i]
		fmt.Printf("%3d | %4s | %4d | %6d | %4d | %5d | %2d | %2d | %2d | %4d\n",
			i+1, k.Nama, k.Main, k.Menang, k.Seri, k.Kalah, k.GolMemasukkan, k.GolKemasukan, k.SelisihGol, k.Poin)
		i++
	}
	fmt.Println("=========================================================================")
}

// menuStatistikGol menampilkan tabel klub yang diurutkan berdasarkan
// produktivitas gol mereka. Bisa diurutkan secara Ascending (terendah ke tertinggi)
// atau Descending (tertinggi ke terendah).
func menuStatistikGol() {
	var pilih int
	fmt.Println("\n=== STATISTIK GOL KLUB ===")
	fmt.Println("1. Gol Terendah - Tertinggi")
	fmt.Println("2. Gol Tertinggi - Terendah")
	fmt.Println("0. Kembali ke Menu Utama")
	fmt.Print("Pilih : ")
	fmt.Scan(&pilih)

	if pilih == 1 {
		selectionSortGol(true)
		fmt.Println("\n=== STATISTIK GOL KLUB (ASCENDING) ===")
	} else if pilih == 2 {
		selectionSortGol(false)
		fmt.Println("\n=== STATISTIK GOL KLUB (DESCENDING) ===")
	} else {
		return
	}

	fmt.Println("====================================")
	fmt.Println("Pos | Klub | Gol")
	fmt.Println("====================================")

	i := 0
	for i < nKlub {
		fmt.Printf("%3d | %4s | %3d\n",
			i+1,
			liga[i].Nama,
			liga[i].GolMemasukkan)
		i++
	}
	fmt.Println("====================================")
	
	insertionSortKlasemen()
}

// ==========================================
// ALGORITMA SEARCHING & SORTING
// ==========================================

// sequentialSearch mencari indeks dari sebuah klub di dalam array
// dengan cara mengecek elemen satu per satu dari awal hingga akhir.
func sequentialSearch(nama string) int {
	idx := -1
	ketemu := false
	i := 0
	for i < nKlub && !ketemu {
		if liga[i].Nama == nama {
			idx = i
			ketemu = true
		}
		i++
	}
	return idx
}

// binarySearch mencari indeks dari sebuah klub dengan membelah area pencarian
// menjadi dua secara berulang. Syaratnya: array liga harus terurut secara abjad.
func binarySearch(nama string) int {
	kiri := 0
	kanan := nKlub - 1
	idx := -1
	ketemu := false

	for kiri <= kanan && !ketemu {
		tengah := (kiri + kanan) / 2
		if liga[tengah].Nama == nama {
			idx = tengah
			ketemu = true
		} else if liga[tengah].Nama < nama {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return idx
}

// selectionSortNama mengurutkan array liga secara alfabetis berdasarkan nama klub.
// Algoritma ini mencari nilai minimum lalu menukarnya ke posisi paling depan.
func selectionSortNama() {
	i := 0
	for i < nKlub-1 {
		idxMin := i
		j := i + 1
		for j < nKlub {
			if liga[j].Nama < liga[idxMin].Nama {
				idxMin = j
			}
			j++
		}
		temp := liga[i]
		liga[i] = liga[idxMin]
		liga[idxMin] = temp
		i++
	}
}

// selectionSortGol mengurutkan array liga berdasarkan jumlah Gol Memasukkan.
// Parameter `ascending` mengatur apakah dari kecil ke besar (true) atau sebaliknya (false).
func selectionSortGol(ascending bool) {
	i := 0
	for i < nKlub-1 {
		idx := i
		j := i + 1
		for j < nKlub {
			if ascending {
				if liga[j].GolMemasukkan < liga[idx].GolMemasukkan {
					idx = j
				}
			} else {
				if liga[j].GolMemasukkan > liga[idx].GolMemasukkan {
					idx = j
				}
			}
			j++
		}
		temp := liga[i]
		liga[i] = liga[idx]
		liga[idx] = temp
		i++
	}
}

// insertionSortKlasemen mengurutkan array liga berdasarkan Poin tertinggi ke terendah.
// Jika Poin sama, maka akan diurutkan berdasarkan Selisih Gol tertinggi.
// Menggunakan algoritma penyisipan (Insertion Sort).
func insertionSortKlasemen() {
	i := 1
	for i < nKlub {
		temp := liga[i]
		j := i
		for j > 0 && (liga[j-1].Poin < temp.Poin || (liga[j-1].Poin == temp.Poin && liga[j-1].SelisihGol < temp.SelisihGol)) {
			liga[j] = liga[j-1]
			j--
		}
		liga[j] = temp
		i++
	}
}