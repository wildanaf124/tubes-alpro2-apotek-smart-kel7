package main
import "fmt"
const MAX = 1000
type Kategori struct {
	IDKategori string
	NamaKategori string
}

type Obat struct {
	IDObat string
	NamaObat string
	IDKategori string
	Indikasi string
	StokTotal int
}

type TransaksiStok struct {
	IDTransaksi string
	IDObat string
	JumlahStok int
	TanggalKadaluarsa string
}

func main() {
	var dataKategori [MAX]Kategori
	var dataObat [MAX]Obat
	var dataTransaksi [MAX]TransaksiStok
	var nKategori, nObat, nTransaksi int
	var endApoteker,endKat,endObat bool
	var endGudang bool
	var end bool = false
	var act int

	for !end {
		fmt.Print("\033[H\033[2J")
		fmt.Printf("========================================\n||   SELAMAT DATANG DI APOTEK-SMART   ||\n========================================\n")
		fmt.Print("Silakan Login:\n1. Login sebagai Apoteker\n2. Login sebagai Staff Gudang\n0. Keluar Aplikasi\nPilih Menu: ")
		fmt.Scan(&act)

		if act == 1 {
			endApoteker = false
			for !endApoteker {
				fmt.Print("\033[H\033[2J")
				fmt.Printf("========================================\n||         DASHBOARD APOTEKER         ||\n========================================\n")
				fmt.Printf("[KELOLA MASTER DATA]\n1. Kelola Data Kategori Gejala\n2. Kelola Data Obat\n")
				fmt.Printf("\n[PENCARIAN & INFORMASI]\n3. Cari Obat\n4. Lihat Laporan Obat & Peringatan\n\n0. Logout\nPilih Menu: ")
				fmt.Scan(&act)
				if act == 1 {
					endKat = false
					for !endKat {
						fmt.Print("\033[H\033[2J")
						fmt.Printf("========================================\n||       KELOLA DATA KATEGORI         ||\n========================================\n")
						fmt.Print("1. Tambah Kategori\n2. Lihat Data Kategori\n3. Ubah Kategori\n4. Hapus Kategori\n0. Kembali\nPilih Menu: ")
						fmt.Scan(&act)
						if act == 1 {
							inputKategori(&dataKategori, &nKategori)
						} else if act == 2 {
							outDataKategori(dataKategori, nKategori)
						} else if act == 3 {
							UpdateKategori(&dataKategori, nKategori)
						} else if act == 4 {
							HapusKategori(&dataKategori, &nKategori)
						} else {
							endKat = true
						}
					}
				} else if act == 2 {
					endObat = false
					for !endObat {
						fmt.Print("\033[H\033[2J")
						fmt.Printf("========================================\n||          KELOLA DATA OBAT          ||\n========================================\n")
						fmt.Print("1. Tambah Obat\n2. Lihat Data Obat\n3. Ubah Obat\n4. Hapus Obat\n0. Kembali\nPilih Menu: ")
						fmt.Scan(&act)
						if act == 1 {
							inputObat(&dataObat, &nObat, dataKategori, nKategori)
						} else if act == 2 {
							outDataObat(dataObat, nObat, dataKategori, nKategori)
						} else if act == 3 {
							UpdateObat(&dataObat, nObat)
						} else if act == 4 {
							HapusObat(&dataObat, &nObat)
						} else {
							endObat = true
						}
					}
				} else if act == 3 {
					cariObat(dataObat, nObat, dataKategori, nKategori)
				} else if act == 4 {
					fmt.Print("\033[H\033[2J")
					fmt.Print("\n======================================================\n||             LAPORAN PERINGATAN APOTEK            ||\n======================================================\n")
					fmt.Print("\n[!] PERINGATAN: OBAT HAMPIR HABIS (STOK < 10)\n")
					outHabis(dataObat, nObat)
					fmt.Print("\n[!] PERINGATAN: OBAT SEGERA KADALUARSA (DIBAWAH TAHUN 2027)\n")
					outKadaluarsa(dataTransaksi, nTransaksi, dataObat, nObat)
					fmt.Print("\n======================================================\n")
					backMenu()
				} else {
					endApoteker = true
				}
			}
		} else if act == 2 {
			endGudang = false
			for !endGudang {
				fmt.Print("\033[H\033[2J")
				fmt.Printf("========================================\n||       DASHBOARD STAFF GUDANG       ||\n========================================\n")
				fmt.Printf("[KELOLA INVENTORI]\n1. Catat Stok Masuk\n")
				fmt.Printf("\n[SORTING & PENGECEKAN]\n2. Urutkan Stok Berdasarkan Kadaluarsa Terdekat\n3. Lihat Daftar Obat Hampir Habis\n\n0. Logout\nPilih Menu: ")
				fmt.Scan(&act)
				if act == 1 {
					inputTransaksi(&dataTransaksi, &nTransaksi, &dataObat, nObat)
				} else if act == 2 {
					selectionSortTransaksiByKadaluarsa(&dataTransaksi, nTransaksi)
					outDaftarStokTransaksi(dataTransaksi, nTransaksi, dataObat, nObat)
				} else if act == 3 {
					fmt.Print("\033[H\033[2J")
					fmt.Print("======================================================\n||               APOTEK-SMART SYSTEM                ||\n======================================================\n")
					fmt.Print("[MENU STAFF GUDANG > DAFTAR OBAT HAMPIR HABIS]\n")
					fmt.Print("\n[!] PERHATIAN: Obat berikut memiliki stok di bawah 10!\nSilakan lakukan pengadaan barang ke supplier.\n")
					fmt.Print("\n------------------------------------------------------\n")
					outHabis(dataObat, nObat)
					fmt.Print("------------------------------------------------------\n")
					backMenu()
				} else {
					endGudang = true
				}
			}
		} else {
			end = true
		}
	}
}

func backMenu() {
	var dummy string
	fmt.Print("\nKetik apa saja dan tekan Enter untuk kembali... ")
	fmt.Scan(&dummy)
}
func seqSearchThereExistObat(K [MAX]Obat, n int, target string) int { //digunakan di: inputTransaksi, outKadaluarsa, outDaftarStokTransaksi, outDataObat, HapusObat, UpdateObat, cariObat
	var i int = 0
	var found int = -1
	for i < n && found == -1 {
		if K[i].IDObat == target {
			found = i
		}
		i++
	}
	return found
}
func seqSearchThereExistKategori(K [MAX]Kategori, n int, target string) int { //digunakan di: inputObat, outDataObat, HapusKategori, UpdateKategori, cariObat
	var i int = 0
	var found int = -1
	for i < n && found == -1 {
		if K[i].IDKategori == target {
			found = i
		}
		i++
	}
	return found
}
func binarySearchObatByNama(K [MAX]Obat, n int, target string) int { //digunakan di: cariObat
	var left, right, mid int = 0, n - 1, 0
	var found int = -1
	for left <= right && found == -1 {
		mid = (left + right) / 2
		if K[mid].NamaObat == target {
			found = mid
		} else if K[mid].NamaObat < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return found
}
func insertionSortObatByNama(K [MAX]Obat, n int) [MAX]Obat { //digunakan di: cariObat
	var i, j int
	var key Obat
	i = 1
	for i < n {
		key = K[i]
		j = i - 1
		for j >= 0 && K[j].NamaObat > key.NamaObat {
			K[j+1] = K[j]
			j--
		}
		K[j+1] = key
		i++
	}
	return K
}
func selectionSortTransaksiByKadaluarsa(K *[MAX]TransaksiStok, n int) { //digunakan di: Gudang(2)
	var i, j, minIdx int
	var temp TransaksiStok
	i = 0
	for i < n-1 {
		minIdx = i
		j = i + 1
		for j < n {
			if K[j].TanggalKadaluarsa < K[minIdx].TanggalKadaluarsa {
				minIdx = j
			}
			j++
		}
		temp = K[i]
		K[i] = K[minIdx]
		K[minIdx] = temp
		i++
	}
}
func inputKategori(K *[MAX]Kategori, n *int) { //Apoteker(1-1)
	var end bool = false
	var data Kategori
	fmt.Print("\033[H\033[2J")
	fmt.Print("======================================================\n||               APOTEK-SMART SYSTEM                ||\n||     Tulis 'Back' di ID Kategori untuk kembali    ||\n======================================================\n")
	fmt.Print("[MENU APOTEKER > KELOLA KATEGORI > TAMBAH KATEGORI]\n")
	for !end {
		if *n >= MAX {
			fmt.Print("Data kategori sudah penuh (maks 1000)!\n")
			end = true
		} else {
			fmt.Print("\nMasukkan ID Kategori : ")
			fmt.Scan(&data.IDKategori)
			if data.IDKategori != "Back" {
				fmt.Print("Masukkan Nama Kategori: ")
				fmt.Scan(&data.NamaKategori)
				K[*n] = data
				*n++
				fmt.Print("Data kategori berhasil ditambahkan!\n")
			} else {
				end = true
			}
		}
	}
}
func inputObat(K *[MAX]Obat, n *int, L [MAX]Kategori, nL int) { //Apoteker(2-1), menggunakan: seqSearchThereExistKategori
	var end bool = false
	var data Obat
	var idx int
	fmt.Print("\033[H\033[2J")
	fmt.Print("======================================================\n||               APOTEK-SMART SYSTEM                ||\n||       Tulis 'Back' di ID Obat untuk kembali      ||\n======================================================\n")
	fmt.Print("[MENU APOTEKER > KELOLA OBAT > TAMBAH OBAT]\n")
	for !end {
		if *n >= MAX {
			fmt.Print("Data obat sudah penuh (maks 1000)!\n")
			end = true
		} else {
			fmt.Print("\nMasukkan ID Obat : ")
			fmt.Scan(&data.IDObat)
			if data.IDObat != "Back" {
				fmt.Print("Masukkan Nama Obat : ")
				fmt.Scan(&data.NamaObat)
				fmt.Print("Masukkan ID Kategori : ")
				fmt.Scan(&data.IDKategori)
				idx = seqSearchThereExistKategori(L, nL, data.IDKategori)
				if idx == -1 {
					fmt.Print("Data Kategori Tidak Ditemukan! Coba Lagi\n")
				} else {
					fmt.Print("Masukkan Indikasi Gejala : ")
					fmt.Scan(&data.Indikasi)
					data.StokTotal = 0
					K[*n] = data
					*n++
					fmt.Print("Data obat berhasil ditambahkan!\n")
				}
			} else {
				end = true
			}
		}
	}
}
func inputTransaksi(K *[MAX]TransaksiStok, n *int, L *[MAX]Obat, nL int) { //Gudang(1), menggunakan: seqSearchThereExistObat
	var end bool = false
	var data TransaksiStok
	var idx int
	fmt.Print("\033[H\033[2J")
	fmt.Print("======================================================\n||               APOTEK-SMART SYSTEM                ||\n||    Tulis 'Back' di ID Transaksi untuk kembali    ||\n======================================================\n")
	fmt.Print("[MENU STAFF GUDANG > CATAT STOK MASUK]\n")
	for !end {
		if *n >= MAX {
			fmt.Print("Data transaksi sudah penuh (maks 1000)!\n")
			end = true
		} else {
			fmt.Print("\nMasukkan ID Transaksi : ")
			fmt.Scan(&data.IDTransaksi)
			if data.IDTransaksi != "Back" {
				fmt.Print("Masukkan ID Obat : ")
				fmt.Scan(&data.IDObat)
				idx = seqSearchThereExistObat(*L, nL, data.IDObat)
				if idx == -1 {
					fmt.Print("Data Obat Tidak Ditemukan! Coba Lagi\n")
				} else {
					fmt.Print("Masukkan Jumlah Stok Masuk : ")
					fmt.Scan(&data.JumlahStok)
					fmt.Print("Masukkan Tgl Kadaluarsa (YYYY-MM-DD) : ")
					fmt.Scan(&data.TanggalKadaluarsa)
					(*L)[idx].StokTotal += data.JumlahStok
					fmt.Printf("Stok berhasil dicatat! (Stok %s bertambah %d)\n", (*L)[idx].NamaObat, data.JumlahStok)
					K[*n] = data
					*n++
				}
			} else {
				end = true
			}
		}
	}
}
func outHabis(K [MAX]Obat, n int) { //Apoteker(4), Gudang(3)
	var i, count int
	count = 1
	i = 0
	for i < n {
		if K[i].StokTotal < 10 {
			fmt.Printf("%d. %s (ID: %s) - Sisa Stok: %d\n", count, K[i].NamaObat, K[i].IDObat, K[i].StokTotal)
			count++
		}
		i++
	}
	if count == 1 {
		fmt.Print("   (Tidak ada obat dengan stok hampir habis)\n")
	}
}
func outKadaluarsa(K [MAX]TransaksiStok, n int, L [MAX]Obat, nL int) { //Apoteker(4), menggunakan: seqSearchThereExistObat
	var i, count, idx int
	count = 1
	i = 0
	for i < n {
		if K[i].TanggalKadaluarsa < "2027-01-01" {
			idx = seqSearchThereExistObat(L, nL, K[i].IDObat)
			if idx != -1 {
				fmt.Printf("%d. %s (Batch %s) - Expired: %s\n", count, L[idx].NamaObat, K[i].IDTransaksi, K[i].TanggalKadaluarsa)
				count++
			}
		}
		i++
	}
	if count == 1 {
		fmt.Print("   (Tidak ada obat yang segera kadaluarsa)\n")
	}
}
func outDaftarStokTransaksi(K [MAX]TransaksiStok, n int, L [MAX]Obat, nL int) { //Gudang(2), menggunakan: seqSearchThereExistObat, backMenu
	var i, idx int
	fmt.Print("\033[H\033[2J")
	fmt.Print("===============================================================================\n||                             DAFTAR STOK OBAT                              ||\n===============================================================================\n")
	fmt.Printf("| %-8s | %-9s | %-18s | %-14s | %-14s |\n", "ID TRX", "ID OBAT", "NAMA OBAT", "TGL KADALUARSA", "SISA STOK TRX")
	fmt.Print("-------------------------------------------------------------------------------\n")
	i = 0
	for i < n {
		idx = seqSearchThereExistObat(L, nL, K[i].IDObat)
		if idx != -1 {
			fmt.Printf("| %-8s | %-9s | %-18s | %-14s | %-14d |\n", K[i].IDTransaksi, K[i].IDObat, L[idx].NamaObat, K[i].TanggalKadaluarsa, K[i].JumlahStok)
		}
		i++
	}
	fmt.Print("===============================================================================\n")
	backMenu()
}
func outDataObat(K [MAX]Obat, n int, L [MAX]Kategori, nL int) { //Apoteker(2-2), menggunakan: seqSearchThereExistKategori, backMenu
	var i, idx int
	fmt.Print("\033[H\033[2J")
	fmt.Print("======================================================================================\n||                                DAFTAR DATA OBAT                                  ||\n======================================================================================\n")
	fmt.Printf("| %-9s | %-19s | %-15s | %-21s | %-6s |\n", "ID OBAT", "NAMA OBAT", "KATEGORI", "INDIKASI", "STOK")
	fmt.Print("--------------------------------------------------------------------------------------\n")
	i = 0
	for i < n {
		idx = seqSearchThereExistKategori(L, nL, K[i].IDKategori)
		if idx != -1 {
			fmt.Printf("| %-9s | %-19s | %-15s | %-21s | %-6d |\n", K[i].IDObat, K[i].NamaObat, L[idx].NamaKategori, K[i].Indikasi, K[i].StokTotal)
		}
		i++
	}
	fmt.Print("======================================================================================\n")
	backMenu()
}
func outDataKategori(K [MAX]Kategori, n int) { //Apoteker(1-2), menggunakan: backMenu
	var i int
	fmt.Print("\033[H\033[2J")
	fmt.Print("============================================================\n||                  DAFTAR DATA KATEGORI                  ||\n============================================================\n")
	fmt.Printf("| %-15s | %-38s |\n", "ID KATEGORI", "NAMA KATEGORI")
	fmt.Print("------------------------------------------------------------\n")
	i = 0
	for i < n {
		fmt.Printf("| %-15s | %-38s |\n", K[i].IDKategori, K[i].NamaKategori)
		i++
	}
	fmt.Print("============================================================\n")
	backMenu()
}
func HapusObat(K *[MAX]Obat, n *int) { //Apoteker(2-4), menggunakan: seqSearchThereExistObat
	var end bool = false
	var target, confirm string
	var idx, i int
	fmt.Print("\033[H\033[2J")
	fmt.Print("======================================================\n||               APOTEK-SMART SYSTEM                ||\n======================================================\n")
	fmt.Print("[MENU APOTEKER > KELOLA OBAT > HAPUS DATA]\n")
	for !end {
		fmt.Print("\nMasukkan ID Obat yang ingin dihapus (atau 'Back' untuk kembali): ")
		fmt.Scan(&target)
		if target != "Back" {
			idx = seqSearchThereExistObat(*K, *n, target)
			if idx == -1 {
				fmt.Print("Data Tidak Ditemukan! Coba Lagi\n")
			} else {
				fmt.Printf("Data Ditemukan: %s (Stok saat ini: %d)\n", (*K)[idx].NamaObat, (*K)[idx].StokTotal)
				fmt.Print("Peringatan: Menghapus data ini akan menghilangkannya dari sistem!\n")
				fmt.Print("Ketik 'HAPUS' untuk konfirmasi (atau ketik 'Back' untuk batal): ")
				fmt.Scan(&confirm)
				if confirm == "HAPUS" {
					i = idx
					for i < *n-1 {
						(*K)[i] = (*K)[i+1]
						i++
					}
					*n--
					fmt.Printf("\n[SUKSES] Data Obat %s berhasil dihapus dari sistem!\n", target)
				} else if confirm == "Back" {
					fmt.Printf("\n[BATAL] Penghapusan Obat %s dibatalkan.\n", target)
				}
			}
		} else {
			end = true
		}
	}
}
func HapusKategori(K *[MAX]Kategori, n *int) { //Apoteker(1-4), menggunakan: seqSearchThereExistKategori
	var end bool = false
	var target, confirm string
	var idx, i int
	fmt.Print("\033[H\033[2J")
	fmt.Print("======================================================\n||               APOTEK-SMART SYSTEM                ||\n======================================================\n")
	fmt.Print("[MENU APOTEKER > KELOLA KATEGORI > HAPUS DATA]\n")
	for !end {
		fmt.Print("\nMasukkan ID Kategori yang ingin dihapus (atau 'Back' untuk kembali): ")
		fmt.Scan(&target)
		if target != "Back" {
			idx = seqSearchThereExistKategori(*K, *n, target)
			if idx == -1 {
				fmt.Print("Data Tidak Ditemukan! Coba Lagi\n")
			} else {
				fmt.Printf("Data Ditemukan: %s\n", (*K)[idx].NamaKategori)
				fmt.Print("Peringatan: Menghapus data ini akan menghilangkannya dari sistem!\n")
				fmt.Print("Ketik 'HAPUS' untuk konfirmasi (atau ketik 'Back' untuk batal): ")
				fmt.Scan(&confirm)
				if confirm == "HAPUS" {
					i = idx
					for i < *n-1 {
						(*K)[i] = (*K)[i+1]
						i++
					}
					*n--
					fmt.Printf("\n[SUKSES] Data Kategori %s berhasil dihapus dari sistem!\n", target)
				} else if confirm == "Back" {
					fmt.Printf("\n[BATAL] Penghapusan Kategori %s dibatalkan.\n", target)
				}
			}
		} else {
			end = true
		}
	}
}
func UpdateObat(K *[MAX]Obat, n int) { //Apoteker(2-3), menggunakan: seqSearchThereExistObat
	var end bool = false
	var target, confirm string
	var idx int
	fmt.Print("\033[H\033[2J")
	fmt.Print("======================================================\n||               APOTEK-SMART SYSTEM                ||\n======================================================\n")
	fmt.Print("[MENU APOTEKER > KELOLA OBAT > UBAH DATA]\n")
	for !end {
		fmt.Print("\nMasukkan ID Obat yang ingin diubah (atau 'Back' untuk kembali): ")
		fmt.Scan(&target)
		if target != "Back" {
			idx = seqSearchThereExistObat(*K, n, target)
			if idx == -1 {
				fmt.Print("Data Tidak Ditemukan! Coba Lagi\n")
			} else {
				fmt.Printf("Data Ditemukan!\n[Data Lama]\nNama Obat : %s\nKategori  : %s\nIndikasi  : %s\n", (*K)[idx].NamaObat, (*K)[idx].IDKategori, (*K)[idx].Indikasi)
				fmt.Printf("\n[Masukkan Data Baru]\n")
				fmt.Printf("Nama Obat Baru (ketik '-' jika tidak ingin diubah)       : ")
				fmt.Scan(&confirm)
				if confirm != "-" {
					(*K)[idx].NamaObat = confirm
				}
				fmt.Printf("ID Kategori Baru (ketik '-' jika tidak ingin diubah)     : ")
				fmt.Scan(&confirm)
				if confirm != "-" {
					(*K)[idx].IDKategori = confirm
				}
				fmt.Printf("Indikasi Gejala Baru (ketik '-' jika tidak ingin diubah) : ")
				fmt.Scan(&confirm)
				if confirm != "-" {
					(*K)[idx].Indikasi = confirm
				}
				fmt.Printf("[SUKSES] Data Obat %s berhasil diperbarui!\n", target)
			}
		} else {
			end = true
		}
	}
}
func UpdateKategori(K *[MAX]Kategori, n int) { //Apoteker(1-3), menggunakan: seqSearchThereExistKategori
	var end bool = false
	var target, replacement string
	var idx int
	fmt.Print("\033[H\033[2J")
	fmt.Print("======================================================\n||               APOTEK-SMART SYSTEM                ||\n======================================================\n")
	fmt.Print("[MENU APOTEKER > KELOLA KATEGORI > UBAH DATA]\n")
	for !end {
		fmt.Print("\nMasukkan ID Kategori yang ingin diubah (atau 'Back' untuk kembali): ")
		fmt.Scan(&target)
		if target != "Back" {
			idx = seqSearchThereExistKategori(*K, n, target)
			if idx == -1 {
				fmt.Print("Data Tidak Ditemukan! Coba Lagi\n")
			} else {
				fmt.Printf("Data Ditemukan: %s\n", (*K)[idx].NamaKategori)
				fmt.Print("\nMasukkan Nama Kategori Baru: ")
				fmt.Scan(&replacement)
				(*K)[idx].NamaKategori = replacement
				fmt.Printf("[SUKSES] Data kategori %s berhasil diperbarui!\n", target)
			}
		} else {
			end = true
		}
	}
}
func cariObat(K [MAX]Obat, n int, L [MAX]Kategori, nL int) { //Apoteker(3), menggunakan: insertionSortObatByNama, binarySearchObatByNama, seqSearchThereExistKategori
	var end bool = false
	var keyword string
	var sortedObat [MAX]Obat
	var idx, idxKat int

	fmt.Print("\033[H\033[2J")
	fmt.Print("======================================================\n||               APOTEK-SMART SYSTEM                ||\n======================================================\n")
	fmt.Print("[MENU APOTEKER > PENCARIAN OBAT]\n")
	fmt.Print("Catatan: Pencarian menggunakan nama obat yang tepat (case-sensitive).\n")

	for !end {
		fmt.Print("\nMasukkan Nama Obat yang dicari (atau 'Back' untuk kembali): ")
		fmt.Scan(&keyword)
		if keyword != "Back" {
			sortedObat = insertionSortObatByNama(K, n)
			idx = binarySearchObatByNama(sortedObat, n, keyword)
			if idx == -1 {
				fmt.Printf("Obat dengan nama '%s' tidak ditemukan.\n", keyword)
			} else {
				idxKat = seqSearchThereExistKategori(L, nL, sortedObat[idx].IDKategori)
				fmt.Print("\n--- HASIL PENCARIAN ---\n")
				fmt.Printf("ID Obat    : %s\n", sortedObat[idx].IDObat)
				fmt.Printf("Nama Obat  : %s\n", sortedObat[idx].NamaObat)
				if idxKat != -1 {
					fmt.Printf("Kategori   : %s\n", L[idxKat].NamaKategori)
				} else {
					fmt.Printf("Kategori   : %s (ID tidak ditemukan)\n", sortedObat[idx].IDKategori)
				}
				fmt.Printf("Indikasi   : %s\n", sortedObat[idx].Indikasi)
				fmt.Printf("Stok Total : %d\n", sortedObat[idx].StokTotal)
				fmt.Print("-----------------------\n")
			}
		} else {
			end = true
		}
	}
}
