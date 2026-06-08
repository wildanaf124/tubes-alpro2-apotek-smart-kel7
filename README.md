# Apotek-Smart 
**Tugas Besar Mata Kuliah Algoritma Pemrograman 2**
**Fakultas Informatika - Telkom University**

Apotek-Smart adalah aplikasi manajemen stok dan inventaris apotek berbasis *Command Line Interface* (CLI) yang dibangun menggunakan bahasa Go (Golang). Aplikasi ini dirancang untuk mengelola persediaan obat, memantau masa berlaku (kadaluarsa) stok, serta memudahkan pencarian dan pengurutan data persediaan.

Sistem ini ditujukan untuk dua peran pengguna utama: **Apoteker** dan **Staf Gudang Farmasi**.

##  Kelompok 7 (Kelas IF-49-09)
* Muhammad Wildan Al-Fattah (103012500157)
* Kesya Tesalonika - (103012500317)


## Fitur dan Spesifikasi Sistem

Sesuai dengan spesifikasi tugas, program ini mengimplementasikan konsep *array*, *struct*, serta algoritma dasar untuk pencarian dan pengurutan:

1. **Kelola Master Data (CRUD):** Pengguna dapat menambahkan, melihat, mengubah, dan menghapus (menggeser elemen *array*) data Obat maupun data Kategori Gejala Penyakit.
2. **Pencatatan Transaksi Stok:** Staf Gudang dapat mencatat stok masuk, yang secara otomatis akan menambahkan jumlah total stok pada master data obat terkait beserta tanggal kadaluarsanya.
3. **Algoritma Pencarian (*Searching*):**
   * **Sequential Search:** Digunakan pada internal sistem untuk memvalidasi keberadaan *ID Kategori* dan *ID Obat* saat melakukan *input*, *update*, atau penghapusan data.
   * **Binary Search:** Digunakan pada fitur pencarian utama untuk mencari detail obat berdasarkan *Nama Obat* yang dimasukkan oleh Apoteker.
4. **Algoritma Pengurutan (*Sorting*):**
   * **Insertion Sort:** Digunakan untuk mengurutkan daftar nama obat secara alfabetis (menaik) sebagai syarat mutlak sebelum *Binary Search* dijalankan.
   * **Selection Sort:** Digunakan oleh Staf Gudang untuk mengurutkan dan menampilkan daftar riwayat stok (transaksi masuk) berdasarkan tanggal kadaluarsa paling dekat.
5. **Sistem Peringatan Otomatis:** Sistem dapat menampilkan statistik berupa:
   * Daftar obat yang stok totalnya hampir habis (di bawah 10 unit).
   * Daftar obat yang akan segera kadaluarsa (sebelum tahun 2027).
