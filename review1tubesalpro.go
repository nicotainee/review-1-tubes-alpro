package main

import "fmt"

const NMAX = 100

type Makanan struct {
	nama        string
	jumlah      int
	kategori    string
	kadarluarsa string
}

type arrMakanan [NMAX]Makanan

func main() {
	var pilihan int
	var n int
	var A arrMakanan
	for {
		fmt.Println("1. Input bahan makanan / minuman")
		fmt.Println("2. Menambah bahan makanan / minuman")
		fmt.Println("3. Menampilkan persediaan bahan makanan / minuman")
		fmt.Println("4. Mencari bahan makanan tertentu")
		fmt.Println("Apa yang ingin anda lakukan: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			bacaData(&A, &n)
		} else if pilihan == 2 {
			menambahMakanan(&A, &n)
		} else if pilihan == 3 {
			tampilkanMakanan(&A, n)
		} else if pilihan == 4 {
			fmt.Print("")
		}
	}
}

func bacaData(A *arrMakanan, n *int) {
	fmt.Print("Masukkan jumlah data yang ingin anda masukkan: ")
	fmt.Scan(n)

	for i := 0; i < *n; i++ {
		fmt.Print("Masukkan kategori makanan / minuman: ")
		fmt.Scan(&(*A)[i].kategori)
		fmt.Print("Nama makanan / minuman: ")
		fmt.Scan(&(*A)[i].nama)
		fmt.Print("Jumlah bahan / minuman: ")
		fmt.Scan(&(*A)[i].jumlah)
		fmt.Print("Masukkan kadarluarsa bahan makanan / minuman: ")
		fmt.Scan(&(*A)[i].kadarluarsa)
	}
}

func menambahMakanan(A *arrMakanan, n *int) {
	var tambah, i int
	fmt.Print("Mau menambahkan berapa bahan? ")
	fmt.Scan(&tambah)

	for i = 0; i < tambah; i++ {
		if *n >= len(*A) {
			fmt.Println("Data sudah penuh, tidak bisa menambah lagi.")
			break
		}
		fmt.Printf("Data ke-%d\n", i+1)
		fmt.Print("Masukkan kategori makanan / minuman: ")
		fmt.Scan(&(*A)[*n].kategori)

		fmt.Print("Nama makanan / minuman: ")
		fmt.Scan(&(*A)[*n].nama)

		fmt.Print("Jumlah bahan / minuman: ")
		fmt.Scan(&(*A)[*n].jumlah)

		fmt.Print("Masukkan kadaluarsa bahan makanan / minuman: ")
		fmt.Scan(&(*A)[*n].kadarluarsa)

		*n++
	}
}

func tampilkanMakanan(A *arrMakanan, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Printf("Data ke-%d:\n", i+1)
		fmt.Println("Kategori    :", A[i].kategori)
		fmt.Println("Nama        :", A[i].nama)
		fmt.Println("Jumlah      :", A[i].jumlah)
		fmt.Println("Kadaluarsa  :", A[i].kadarluarsa)

	}
}

func cariBahan(A *arrMakanan, n int) {
	var input string
	var ketemu, k int
	fmt.Print("Masukkan Nama Bahan: ")
	fmt.Scan(&input)
	ketemu = -1
	k = 0
	for ketemu == -1 && k < n {
		if A[k].nama == input {
			ketemu = k
		}
		k++
	}
	if ketemu != -1 {
		fmt.Println("Kategori    :", A[i].kategori)
		fmt.Println("Nama        :", A[i].nama)
		fmt.Println("Jumlah      :", A[i].jumlah)
		fmt.Println("Kadaluarsa  :", A[i].kadarluarsa)
	} else {
		fmt.Println("Data Tidak Ditemukan")
	}
}

func urutkanData(A *arrMakanan, n int) {

}
