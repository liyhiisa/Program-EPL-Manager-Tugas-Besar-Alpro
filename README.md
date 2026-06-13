# Program-EPL-Manager-Tugas-Besar-Alpro

EPL Manager adalah sebuah aplikasi berbasis Command Line Interface (CLI) yang ditulis menggunakan bahasa pemrograman Go (Golang). Aplikasi ini dirancang untuk menyimulasikan sistem manajemen liga sepak bola profesional layaknya English Premier League (EPL).
Melalui aplikasi ini, admin dapat mengelola klub peserta, men-generate jadwal kompetisi penuh (double round-robin) secara otomatis menggunakan algoritma Tabel Berger, mencatat skor pertandingan per pekan, serta menampilkan klasemen yang diperbarui secara real-time.

 Fitur Utama
Manajemen Klub: Tambah, ubah, dan hapus klub peserta liga (batas nama 3 karakter, misal: ARS, LIV, MCI).
Generator Jadwal: Pembuatan jadwal satu musim penuh secara otomatis dengan format Home-Away menggunakan sistem rotasi Tabel Berger.
Pencatatan Skor: Input hasil pertandingan setiap pekan yang akan otomatis mengalkulasi statistik (Menang, Seri, Kalah, Gol, Selisih Gol, dan Poin).
Klasemen Otomatis: Menampilkan tabel peringkat klub yang diurutkan berdasarkan Poin tertinggi dan tie-breaker Selisih Gol.
Statistik Gol: Menampilkan daftar produktivitas gol klub yang bisa diurutkan dari yang terbanyak (Descending) maupun paling sedikit (Ascending).
Sistem Koreksi: Fasilitas untuk mengedit atau menghapus riwayat hasil pertandingan (statistik akan dikembalikan seperti semula secara otomatis).
