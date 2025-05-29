package main

import 	"fmt"

const NMAX = 100

type Makanan struct {
	id	int
	nama  string
	jumlah  int
	kategori string
	kadarluarsa tanggal
}
type tanggal struct{
	no,bulan,tahun int
}
type bhnHabisPakai struct {
	idPakai	int
	namaPakai   string
	jumlahPakai  int
	kategoriPakai  string
	kadarluarsaPakai tanggal
}
type arrMakanan [NMAX] Makanan
type arrTanggal [NMAX] tanggal
type arrbhnHabisPakai [NMAX] bhnHabisPakai
func main() {
	var pilihan, n ,m int
	var selesai bool
	var A arrMakanan
	var B arrTanggal
	var C arrbhnHabisPakai
	selesai = true

	for selesai {
		fmt.Println("1. Input bahan makanan / minuman")
		fmt.Println("2. Menambah bahan makanan / minuman")
		fmt.Println("3. Mengubah bahan makanan / minuman")
		fmt.Println("4. Menampilkan persediaan bahan makanan / minuman")
		fmt.Println("5. Mencari bahan makanan tertentu")
		fmt.Println("6. Kadaluarsa makanan dalam 30 hari ke depan")
		fmt.Println("7. Mengurutkan bahan makanan / minuman")
		fmt.Println("8. Masukkan bahan yang habis dipakai")
		fmt.Println("0. Keluar")
		fmt.Print("Apa yang ingin anda lakukan: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			bacaData(&A, &n)
		} else if pilihan == 2 {
			fmt.Println("1.Menambhkan jumlah bahan : ")
			fmt.Println("2.Menambahkan bahan")
			fmt.Scan(&pilihan)
			if pilihan == 1{
			tambahJumlah(&A,&n)
			}else if pilihan == 2{
			menambahMakanan(&A, &n)
			}else {
				fmt.Print("Input tidak sesuai")
			}
		} else if pilihan == 3 {
			ubahData(&A, n)
		} else if pilihan == 4 {
			fmt.Println("Daftar bahan makanan yang tersedia: ")
			tampilkanMakanan(A, n)
			fmt.Println("Daftar bahan makanan yang sudah digunakan: ")
			tampilkanMakananHbsPakai(C,m)
		} else if pilihan == 5 {
			fmt.Println("1. Berdasarkan nama")
			fmt.Println("2. Berdasarkan ID")
			fmt.Print("Pilihan pencarian: ")
			fmt.Scan(&pilihan)
			if pilihan == 1 {
				cariBahan(&A, n)
			} else if pilihan == 2 {
				cariId(&A, n)
			} else {
				fmt.Println("Input tidak sesuai")
			}
		} else if pilihan == 6 {
			kadaluarsa(&B, &A, n)
		} else if pilihan == 7 {
			fmt.Println("1. Tanggal kadaluarsa")
			fmt.Println("2. Jumlah stok bahan")
			fmt.Print("Pilihan pengurutan: ")
			fmt.Scan(&pilihan)
			if pilihan == 1 {
				insertSortTanggal(&A, n)
				fmt.Println("Data diurutkan berdasarkan tanggal kadaluarsa:")
				tampilkanMakanan(A, n)
			} else if pilihan == 2 {
				selectSortJumlah(&A, n)
				fmt.Println("Data diurutkan berdasarkan jumlah bahan:")
				tampilkanMakanan(A, n)
			} else {
				fmt.Println("Input tidak sesuai")
			}
		}else if pilihan == 8{
			hapusMakananDanSimpan(&A,&C,&n,&m)
		}else if pilihan == 0 {
			selesai = false
			fmt.Println("Program Selesai  : TERIMA KASIH  ")
		} else {
			fmt.Println("Input tidak sesuai")
		}
		fmt.Println()
	}}



func bacaData(A *arrMakanan, n *int) {
	var id,i,jml int
	var validId,validJml bool 
	fmt.Print("Masukkan jumlah data yang ingin anda masukkan: ")
	fmt.Scan(n)
	
	for i = 0; i < *n; i++ {
	validId = false
	validJml = false
		fmt.Println("============================================")
		fmt.Print("Masukkan ID makanan / minuman: ")
		fmt.Scan(&id)
		for  !validId {
			if id >= 100 && id <= 999 {
       		 A[i].id = id
			 validId = true
    		} else {
       		 fmt.Println("Data Tidak Sesuai (ID harus antara 100 - 999)")
			fmt.Print("Masukkan ID makanan / minuman: ")
			fmt.Scan(&id)

		}}
		fmt.Print("Masukkan kategori bahan: ")
		fmt.Scan(&A[i].kategori)
		fmt.Print("Masukkan nama bahan: ")
		fmt.Scan(&A[i].nama)

		fmt.Print("Masukkan jumlah bahan: ")
        fmt.Scan(&jml)
		for  !validJml {
			if jml >= 0 {
       		 A[i].jumlah = jml
			 validJml = true
    		} else {
       		 fmt.Print("Data Tidak Sesuai (jumlah harus lebih dari 0)")
			 fmt.Print("Masukkan jumlah bahan: ")
			fmt.Scan(&jml)
		}}

		fmt.Print("Masukkan tanggal kadaluarsa baru (dd mm yyyy): ")
			fmt.Scan(&A[i].kadarluarsa.no, &A[i].kadarluarsa.bulan, &A[i].kadarluarsa.tahun)
		fmt.Println("============================================")
	}
}

func menambahMakanan(A *arrMakanan, n *int) {
    var tambah,i,id,jml int
	var validId, validJml bool
    fmt.Print("Mau menambahkan berapa bahan? ")
    fmt.Scan(&tambah)
	validId = false
	validJml = false
	for i = 0; i < tambah && *n < NMAX; i++ {
		fmt.Println("============================================")
		fmt.Print("Masukkan ID makanan / minuman: ")
		fmt.Scan(&id)
		for  !validId {
			if id >= 100 && id <= 999 {
       		 (*A)[*n].id = id
			 validId = true
    		} else {
       		 fmt.Println("Data Tidak Sesuai (ID harus antara 100 - 999)")
				fmt.Print("Masukkan ID makanan / minuman: ")
				fmt.Scan(&id)
		}}
		fmt.Print("Masukkan kategori bahan: ")
		fmt.Scan(&A[*n].kategori)
		fmt.Print("Masukkan nama bahan: ")
		fmt.Scan(&A[*n].nama)
		fmt.Print("Masukkan jumlah bahan: ")
        fmt.Scan(&jml)
		for  !validJml {
			if jml >= 0 {
       		 (*A)[*n].jumlah = jml
			 validJml = true
    		} else {
       		 fmt.Print("Data Tidak Sesuai (jumlah harus lebih dari 0)")
			 fmt.Print("Masukkan jumlah bahan: ")
			 fmt.Scan(&jml)
		}}

		fmt.Print("Masukkan tanggal kadaluarsa baru (dd mm yyyy): ")
			fmt.Scan(&A[*n].kadarluarsa.no, &A[*n].kadarluarsa.bulan, &A[*n].kadarluarsa.tahun)
		fmt.Println("============================================")
        *n = *n +1
    }
}
func tambahJumlah(A *arrMakanan, n *int){
	var i,input,id int
	fmt.Print("Masukkan Id:")
	fmt.Scan(&id)
	for i = 0 ; i < *n ; i++{
		if A[i].id == id {
	fmt.Print("Masukkan jumlah bahan yang ingin ditambahkan :")
	fmt.Scan(&input)
	A[i].jumlah = A[i].jumlah + input
	fmt.Print("Data berhasil ditambahkan")
		}else {
			fmt.Print("Data tidak ditrmukan")
		}
	}
}
func ubahData(A *arrMakanan, n int) {
	var idCari,id,i,jml int
	var found,validId,validJml bool
	validId=false
	validJml=false
	fmt.Print("Masukkan ID bahan yang ingin diubah: ")
	fmt.Scan(&idCari)

	found = false
	for i = 0; i < n; i++ {
		if (*A)[i].id == idCari {
			found = true
			fmt.Println("Data ditemukan. Silakan masukkan data baru.")
			fmt.Print("Masukkan ID baru: ")
			fmt.Scan(&id)
			for  !validId {
			if id >= 100 && id <= 999 {
       		 (*A)[i].id = id
			 validId = true
    		} else {
       		 fmt.Println("Data Tidak Sesuai (ID harus antara 100 - 999)")
				fmt.Print("Masukkan ID makanan / minuman   : ")
				fmt.Scan(&id)
		}}
			fmt.Print("Masukkan kategori baru : ")
			fmt.Scan(&A[i].kategori)
			fmt.Print("Masukkan nama baru     :")
			fmt.Scan(&A[i].nama)
			fmt.Print("Masukkan jumlah bahan  : ")
			fmt.Scan(&jml)
			for  !validJml {
			if jml >= 0 {
       		 (*A)[i].jumlah = jml
			 validJml = true
    		} else {
       		 fmt.Print("Data Tidak Sesuai (jumlah harus lebih dari 0)")
			 fmt.Print("Masukkan jumlah bahan                  : ")
			 fmt.Scan(&jml)
		}}
			fmt.Print("Masukkan tanggal kadaluarsa (dd mm yyyy): ")
			fmt.Scan(&A[i].kadarluarsa.no, &A[i].kadarluarsa.bulan, &A[i].kadarluarsa.tahun)
		}
	}

	if !found {
		fmt.Println("ID tidak ditemukan.")
	}
}


func tampilkanMakanan(A arrMakanan, n int) {
    var i int
    for i = 0; i < n; i++ {
        fmt.Println("============================================")
		fmt.Println("ID :", A[i].id)
        fmt.Println("Kategori    :", A[i].kategori)
        fmt.Println("Nama        :", A[i].nama)
        fmt.Println("Jumlah      :", A[i].jumlah)
        fmt.Printf("Kadaluarsa  : %02d %02d %04d\n", A[i].kadarluarsa.no,  A[i].kadarluarsa.bulan, A[i].kadarluarsa.tahun)
		fmt.Println("============================================")
		fmt.Println()
    }
}
func binarySearchNama(A *arrMakanan, n int, target string) int {
	var left,right,mid int
	left = 0
	right = n - 1
	for left <= right {
		mid = (left + right) / 2
		if A[mid].nama == target {
			return mid
		} else if A[mid].nama < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}


func cariBahan(A *arrMakanan, n int){
	var input string 
	var ketemu, k int
	fmt.Println("Masukkan Nama Bahan: ")
	fmt.Scan(&input)
	ketemu = -1
	k = 0
	for ketemu == -1 && k < n {
		if A[k].nama == input {
			ketemu = k
		}
		k++
	}
	if ketemu != -1{
	fmt.Println("Nama :",A[ketemu].nama)
	fmt.Println("Jumlah :",A[ketemu].jumlah)
	fmt.Println("Kategori :",A[ketemu].kategori)
	fmt.Println("Kadarluarsa :",A[ketemu].kadarluarsa.no,A[ketemu].kadarluarsa.bulan,A[ketemu].kadarluarsa.tahun)
}else {
	fmt.Println("Data Tidak Ditemukan")
}
}
func cariId(A *arrMakanan, n int) {
	var x, ketemu int
	var found bool
	fmt.Print("Masukkan ID bahan: ")
	fmt.Scan(&x)
	found = false
	for i := 0; i < n; i++ {
		if A[i].id == x {
			ketemu = i
			found = true
			break
		}
	}
	if found {
		fmt.Println("Nama :", A[ketemu].nama)
		fmt.Println("Jumlah :", A[ketemu].jumlah)
		fmt.Println("Kategori :", A[ketemu].kategori)
		fmt.Println("Kadarluarsa :", A[ketemu].kadarluarsa.no, A[ketemu].kadarluarsa.bulan, A[ketemu].kadarluarsa.tahun)
	} else {
		fmt.Println("ID tidak ditemukan.")
	}
}


func kadaluarsa(B *arrTanggal, A *arrMakanan, n int) {
	var H, exp tanggal
	var totH, totK ,i int
	var ada bool = false

	fmt.Println("Masukkan tanggal hari ini (dd mm yyyy): ")
	fmt.Scan(&H.no, &H.bulan, &H.tahun)

	totH = H.tahun*365 + H.bulan*30 + H.no

	fmt.Println("Bahan makanan yang akan kadaluarsa dalam 30 hari ke depan:")
	for i = 0; i < n; i++ {
		exp = (*A)[i].kadarluarsa
		totK = exp.tahun*365 + exp.bulan*30 + exp.no

		if totK >= totH && totK <= totH+30 {
			ada = true
			fmt.Printf("- %s (kadaluarsa: %02d %02d %04d )\n ",(*A)[i].nama, exp.no, exp.bulan, exp.tahun )
		
		}
	}

	if !ada {
		fmt.Println("Tidak ada bahan makanan yang akan kadaluarsa dalam 30 hari ke depan.")
	}
}



func insertSortTanggal(A *arrMakanan, n int) {
	var i, j int
	var temp Makanan
	for i = 1; i < n; i++ {
		temp = (*A)[i]
		j = i - 1
		for j >= 0 && lebihLama((*A)[j].kadarluarsa, temp.kadarluarsa) {
			(*A)[j+1] = (*A)[j]
			j--
		}
		(*A)[j+1] = temp
	}
}

func lebihLama(t1, t2 tanggal) bool {
	if t1.tahun > t2.tahun {
		return true
	} else if t1.tahun == t2.tahun {
		if t1.bulan > t2.bulan {
			return true
		} else if t1.bulan == t2.bulan {
			return t1.no > t2.no
		}
	}
	return false
}


func selectSortJumlah(A *arrMakanan, n int) {
	var i, j, minIdx int
	var temp Makanan

	for i = 0; i < n-1; i++ {
		minIdx = i
		for j = i + 1; j < n; j++ {
			if (*A)[j].jumlah < (*A)[minIdx].jumlah {
				minIdx = j
			}
		}
		
		temp = (*A)[i]
		(*A)[i] = (*A)[minIdx]
		(*A)[minIdx] = temp
	}
}
func hapusMakananDanSimpan(A *arrMakanan, B *arrbhnHabisPakai, n *int, m *int) {
	var i,id,jumlah int
	var found bool 
	found = false

	fmt.Print("Masukkan ID bahan: ")
	fmt.Scan(&id)

	for i = 0; i < *n; i++ {
		if (*A)[i].id == id {
			found = true
			fmt.Print("Masukkn jumlah bahan yang sudah digunakan")
			fmt.Scan(&jumlah)
				

			(*B)[*m].idPakai = (*A)[i].id
			(*B)[*m].namaPakai = (*A)[i].nama
			(*B)[*m].jumlahPakai = jumlah
			(*B)[*m].kategoriPakai = (*A)[i].kategori
			(*B)[*m].kadarluarsaPakai = (*A)[i].kadarluarsa
			*m++
			if jumlah == (*A)[i].jumlah {
				for j := i; j < *n-1; j++ {
				(*A)[j] = (*A)[j+1]
				}
			*n--
			
		}else if jumlah > (*A)[i].jumlah{
			fmt.Print("Jumlah melebiihi stok yang ada")
		} else {
			(*A)[i].jumlah  = (*A)[i].jumlah  -  jumlah
		}
		}
	}

	if found {
		fmt.Println("Data berhasil dihapus dan disimpan ke laporan.")
	} else {
		fmt.Println("ID tidak ditemukan.")
	}
}
func tampilkanMakananHbsPakai(A arrbhnHabisPakai, n int) {
    var i int
    for i = 0; i < n; i++ {
		fmt.Println("============================================")
		fmt.Println("ID 		 :", A[i].idPakai)
        fmt.Println("Kategori    :", A[i].kategoriPakai)
        fmt.Println("Nama        :", A[i].namaPakai)
        fmt.Println("Jumlah      :", A[i].jumlahPakai)
        fmt.Printf("Kadaluarsa  : %02d %02d %04d\n", A[i].kadarluarsaPakai.no,  A[i].kadarluarsaPakai.bulan, A[i].kadarluarsaPakai.tahun)
		fmt.Println("============================================")
		fmt.Println()
    }
}
