package main

import "fmt"

const NMAX = 100

type data struct {
	namaBarang, kategori string
	jumlah               int
}

type tabInt [NMAX]data

func main() {
	var A tabInt
	inputData(&A)
}

func inputData(A *tabInt) {
	i := 0
	var angkaInput int
	for {
		fmt.Println("---------Selamat Datang di toko kelontong---------")
		fmt.Println("Silahkan pilih kategori barang:")
		fmt.Println("1.Makanan")
		fmt.Println("2.Minuman")
		fmt.Println("3.Pakaian")
		fmt.Println("4.List barang pesanan")
		fmt.Println("5.Selesai")
		fmt.Println()
		fmt.Print("Masukkan angka untuk memilih kategori: ")
		fmt.Scan(&angkaInput)
		if angkaInput == 1 {
			A[i].kategori = "Makanan"
			fmt.Print("Masukkan nama makanan: ")
			fmt.Scan(&A[i].namaBarang)
			fmt.Print("Masukkan jumlah barang: ")
			fmt.Scan(&A[i].jumlah)
			i++
			sortingBarang(A, i)
		} else if angkaInput == 2 {
			A[i].kategori = "Minuman"
			fmt.Print("Masukkan nama makanan: ")
			fmt.Scan(&A[i].namaBarang)
			fmt.Print("Masukkan jumlah barang: ")
			fmt.Scan(&A[i].jumlah)
			i++
			sortingBarang(A, i)
		} else if angkaInput == 3 {
			A[i].kategori = "Pakaian"
			fmt.Print("Masukkan nama makanan: ")
			fmt.Scan(&A[i].namaBarang)
			fmt.Print("Masukkan jumlah barang: ")
			fmt.Scan(&A[i].jumlah)
			i++
			sortingBarang(A, i)
		} else if angkaInput == 4 {
			outputPesanan(A, i)
		} else if angkaInput == 5 {
			break
		}
	}
}

func sortingBarang(A *tabInt, n int) {
	for i := 0; i < n; i++ {
		max := i
		for j := i + 1; j < n; j++ {
			if A[j].jumlah > A[max].jumlah {
				max = j
			}
		}
		A[i], A[max] = A[max], A[i]
	}
}

func outputPesanan(A *tabInt, n int) {
	var namaKategori string
	fmt.Print("Masukkan kategori: ")
	fmt.Scan(&namaKategori)
	fmt.Println("---Nama Barang---Kategori---Jumlah---")
	for i := 0; i < n; i++ {
		if A[i].kategori == namaKategori {
			fmt.Println("---", A[i].namaBarang, "---", A[i].kategori, "---", A[i].jumlah, "---")
		}
	}
}
