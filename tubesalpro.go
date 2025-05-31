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
	var pilihan, n ,m,tarSeq int
	var selesai,cek bool
	var tarbin string
	var A arrMakanan
	var B arrTanggal
	var C arrbhnHabisPakai
	selesai = true
	cek = false
	
		fmt.Println("==========================================================")
		fmt.Println("Selamat datang di MyFoodList!")
		fmt.Println("Di sini anda dapat mencatat persediaan bahan makanan anda.")
		fmt.Println("Apakah anda ingin mencatat persediaan anda? (1 = Ya, 0 = Tidak)")
		fmt.Print("Masukkan pilihan anda: ")
		fmt.Scan(&pilihan)
		fmt.Println("==========================================================")
		if pilihan == 1 {
			// Memanngil fungsi untuk input data
			fmt.Println("1. Silahkan input bahan makanan / minuman")
			bacaData(&A, &n)
			cek = true
		}else{
			fmt.Print("Terimakasih telah berkunjung ^^")
		}
		// Jika sudah inputdata maka menu akan ditampilkan
		for selesai {
			if cek == true {
		fmt.Println("==========================================================")
		fmt.Println("|                          MENU                          |")
		fmt.Println("==========================================================")
		fmt.Println("1. Menambah bahan makanan / minuman")
		fmt.Println("2. Mengubah bahan makanan / minuman")
		fmt.Println("3. Menampilkan persediaan bahan makanan / minuman")
		fmt.Println("4. Mencari bahan makanan tertentu")
		fmt.Println("5. Menampilkan kadaluarsa bahan dalam 30 hari ke depan")
		fmt.Println("6. Mengurutkan bahan makanan / minuman")
		fmt.Println("7. Masukkan bahan yang habis dipakai")
		fmt.Println("0. Keluar")
		fmt.Print("Apa yang ingin anda lakukan: ")
		fmt.Scan(&pilihan)
		fmt.Println("==========================================================")
		} 
			if pilihan == 1 {
				// Menambah stok bahan makanan
				fmt.Println("1.Menambhkan jumlah bahan  ")
				fmt.Println("2.Menambahkan bahan")
				fmt.Println("Anda ingin menambahkan berdasarkan :")
				fmt.Scan(&pilihan)
			if pilihan == 1{
				tambahJumlah(&A,&n)
			}else if pilihan == 2{
				menambahMakanan(&A, &n)
			}else {
				fmt.Print("Input tidak sesuai")
			}
		} else if pilihan == 2 {
			// Mengubah data bahan makanan
				ubahData(&A, n)
		} else if pilihan == 3 {
			// Menampilkan stok bahan makanan dan bahan makanan habis pakai
			fmt.Println("Daftar bahan makanan yang tersedia: ")
			tampilkanMakanan(A, n)
			fmt.Println("Daftar bahan makanan yang sudah digunakan: ")
			tampilkanMakananHbsPakai(C,m)
		} else if pilihan == 4 {
			// mencari bahan makanan tertentu berdasarkan nama atau Id
			fmt.Println("1. Berdasarkan nama")
			fmt.Println("2. Berdasarkan ID")
			fmt.Print("Pilihan pencarian: ")
			fmt.Scan(&pilihan)
			if pilihan == 1 {
				fmt.Print("Masukkan nama bahan yang mau dicari :")
				fmt.Scan(&tarbin)

				if binarySearchNama(&A,n,tarbin) != -1 {
					tampilkanMakanan(A, n)
				} else {
					fmt.Println("Data tidak ditemukan")
		}
			} else if pilihan == 2 {
				fmt.Print("Masukkan ID bahan: ")
				fmt.Scan(&tarSeq)
				
				if cariId(&A, n,tarSeq) != -1 {
					tampilkanMakanan(A, n)
				} else {
					fmt.Println("Data tidak ditemukan")
				
			} 
		} else if pilihan == 5 {
			// Menampilkan kadarluarsa bahan selama 30 hari kedepan
			kadaluarsa(&B, &A, n)
		} else if pilihan == 6 {
			// Mengurutkan bahan makanan berdasarkan tanggal kadarluarsa dan jumlah stok
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
		}else if pilihan == 7{
			//Mencatat bahan makanan yang sudah digunakan
			hapusMakananDanSimpan(&A,&C,&n,&m)
		}else if pilihan == 0 {
			// Mengakhiri program
			selesai = false
			fmt.Println("Program Selesai  : TERIMA KASIH  ")
		} else {
			fmt.Println("Input tidak sesuai")
		}
		fmt.Println()
	}}}


// Fungsi untuk input data awal bahan makanan
//I.S. (initial atate): masukkan dr procedurenya apa
//F.S. (final state): keluaran dari procedurenya apa
func bacaData(A *arrMakanan, n *int) {
	var id,i,jml,no,bulan int
	var validId,validJml,validTgl bool 
	fmt.Print("Masukkan jumlah data yang ingin anda masukkan: ")
	fmt.Scan(n)
	
	for i = 0; i < *n; i++ {
	validId = false
	validJml = false
	validTgl= false
		fmt.Println("==========================================================")
		fmt.Print("Masukkan ID makanan / minuman: ")
		fmt.Scan(&id)
		// Memvalidasi Id harus diantara 100 dan 999
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
		// Memvalidasi jumlah harus diatas 0
		for  !validJml {
			if jml >= 1  {
       		 A[i].jumlah = jml
			 validJml = true
    		} else {
       		 fmt.Print("Data Tidak Sesuai (jumlah harus lebih dari 0)")
			 fmt.Print("Masukkan jumlah bahan: ")
			fmt.Scan(&jml)
		}}

		fmt.Print("Masukkan tanggal kadaluarsa baru (dd mm yyyy): ")
		fmt.Scan(&no, &bulan, &A[i].kadarluarsa.tahun)
		// Memvalidasi tanggal tidak boleh lebih dari 31 dan bulan tidak lebih dari 12
			for !validTgl {
				if no <= 31 && bulan <= 12 {
					A[i].kadarluarsa.no = no
					A[i].kadarluarsa.bulan =bulan
					validTgl=true
				}else {
					fmt.Println("Input tidak sesuai")
					fmt.Print("Masukkan tanggal kadaluarsa baru (dd mm yyyy): ")
					fmt.Scan(&no, &bulan, &A[i].kadarluarsa.tahun)
				}
			}
		fmt.Println("==========================================================")
	}
}
// fungsi menambah bahan makanan baru ke data yang sudah ada
// I.S.: Array bahan makanan (A) sudah berisi n data, n <= NMAX
// F.S.: Data bahan baru ditambahkan ke array A, n bertambah sesuai jumlah bahan baru yang valid

func menambahMakanan(A *arrMakanan, n *int) {
    var tambah,i,id,jml,no,bulan int
	var validId, validJml,validTgl bool
    fmt.Print("Mau menambahkan berapa bahan? ")
    fmt.Scan(&tambah)
	
	for i = 0; i < tambah && *n < NMAX; i++ {
	validId = false
	validJml = false
	validTgl = false
		fmt.Println("==========================================================")
		fmt.Print("Masukkan ID makanan / minuman: ")
		fmt.Scan(&id)
		// Memvalidasi Id harus diantara 100 dan 999
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
		// Memvalidasi jumlah harus lebih dari 0
		for  !validJml {
			if jml >= 1 {
       		 (*A)[*n].jumlah = jml
			 validJml = true
    		} else {
       		 fmt.Print("Data Tidak Sesuai (jumlah harus lebih dari 0)")
			 fmt.Print("Masukkan jumlah bahan: ")
			 fmt.Scan(&jml)
		}}

		fmt.Print("Masukkan tanggal kadaluarsa baru (dd mm yyyy): ")
		fmt.Scan(&no, &bulan, &A[*n].kadarluarsa.tahun)
		// Memvalidasi tanggal tidak boleh lebih dari 31 dan bulan tidak lebih dari 12
			for !validTgl {
				if no <= 31 && bulan <= 12 {
					A[*n].kadarluarsa.no = no
					A[*n].kadarluarsa.bulan =bulan
					validTgl=true
				}else {
					fmt.Println("Input tidak sesuai")
					fmt.Print("Masukkan tanggal kadaluarsa baru (dd mm yyyy): ")
					fmt.Scan(&no, &bulan, &A[*n].kadarluarsa.tahun)
				}
			}
		fmt.Println("==========================================================")
        *n = *n +1
    }
}
// Fungsi untuk menambahkan jumlah stok bahan makanan berdasarkan ID
// I.S.: Array bahan makanan (A) sudah berisi n data, user memasukkan ID bahan yang ada
// F.S.: Jumlah bahan pada ID yang sesuai bertambah sesuai input user, atau pesan error jika ID tidak ditemukan
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
			fmt.Print("Data tidak ditemukan")
		}
	}
}
// Fungsi mengubah data bahan makanan berdasarkan ID
// I.S.: Array bahan makanan (A) berisi n data, user memasukkan ID bahan yang ingin diubah
// F.S.: Data bahan dengan ID yang sesuai diubah sesuai input user, atau pesan error jika ID tidak ditemukan
func ubahData(A *arrMakanan, n int) {
	var idCari,id,i,jml,no,bulan int
	var found,validId,validJml,validTgl bool
	
	fmt.Print("Masukkan ID bahan yang ingin diubah: ")
	fmt.Scan(&idCari)

	found = false
	for i = 0; i < n; i++ {
	validId=false
	validJml=false
	validTgl=false
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
				if jml >= 1 {
       				 (*A)[i].jumlah = jml
					 validJml = true
    			} else {
       		 		fmt.Print("Data Tidak Sesuai (jumlah harus lebih dari 0)")
			 		fmt.Print("Masukkan jumlah bahan                  : ")
					 fmt.Scan(&jml)
		}}
			fmt.Print("Masukkan tanggal kadaluarsa (dd mm yyyy): ")
			
			fmt.Scan(&no, &bulan, &A[i].kadarluarsa.tahun)
			fmt.Print("Masukkan tanggal kadaluarsa baru (dd mm yyyy): ")
		fmt.Scan(&no, &bulan, &A[i].kadarluarsa.tahun)
			for !validTgl {
				if  no <= 31  && bulan <= 12 {
					A[i].kadarluarsa.no = no
					A[i].kadarluarsa.bulan =bulan
					validTgl=true
				}else {
					fmt.Println("Input tidak sesuai")
					fmt.Print("Masukkan tanggal kadaluarsa baru (dd mm yyyy): ")
					fmt.Scan(&no, &bulan, &A[i].kadarluarsa.tahun)
				}
			}
		}
	}

	if !found {
		fmt.Println("ID tidak ditemukan.")
	}
}

// fungsi untuk menampilkan daftar bahan makanan yang tersedia
// I.S.: Array bahan makanan (A) berisi n data
// F.S.: Menampilkan semua data bahan makanan dari array A ke layar
func tampilkanMakanan(A arrMakanan, n int) {
    var i int
    for i = 0; i < n; i++ {
        fmt.Println("==========================================================")
		fmt.Println("ID :", A[i].id)
        fmt.Println("Kategori    :", A[i].kategori)
        fmt.Println("Nama        :", A[i].nama)
        fmt.Println("Jumlah      :", A[i].jumlah)
        fmt.Printf("Kadaluarsa  : %02d %02d %04d\n", A[i].kadarluarsa.no,  A[i].kadarluarsa.bulan, A[i].kadarluarsa.tahun)
		fmt.Println("==========================================================")
		fmt.Println()
    }
}
//  Fungsi binary search untuk mencari bahan berdasarkan nama
// I.S.: Array bahan makanan (A) berisi n data yang sudah terurut berdasarkan nama, user memasukkan nama bahan yang dicari
// F.S.: Menampilkan data bahan sesuai nama jika ditemukan, atau pesan error jika tidak ditemukan
func binarySearchNama(A *arrMakanan, n int, target string) int {
	var left, right, mid,found int
	found = -1 // -1 artinya tidak ditemukan

	left = 0
	right = n - 1

	for left <= right && found == -1 {
		mid = (left + right) / 2
		if A[mid].nama == target {
			found = mid
		} else if A[mid].nama < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return found
}
	

// Fungsi Sequential Search untuk mencari bahan berdasarkan ID 
//// I.S.: Array bahan makanan (A) berisi n data, user memasukkan ID bahan yang dicari
// F.S.: Menampilkan data bahan sesuai ID jika ditemukan, atau pesan error jika tidak ditemukan
func cariId(A *arrMakanan, n int, X int) int {
	var i int
	for i = 0; i < n; i++ {
		if A[i].id == X {
			return i 
		}
	}
	return -1
}

// Fungsi menampilkan bahan yang akan kadaluarsa dalam 30 hari ke depan
// I.S.: Array bahan makanan (A) berisi n data dengan tanggal kadaluarsa, dan array tanggal B (tidak digunakan optimal)
// F.S.: Menampilkan bahan yang akan kadaluarsa dalam 30 hari ke depan 
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


// Fungsi Insertion Sort untuk mengurutkan bahan berdasarkan tanggal kadaluarsa 
// I.S.: Array bahan makanan (A) berisi n data dengan tanggal kadaluarsa tidak berurutan
// F.S.: Array A terurutkan berdasarkan tanggal kadaluarsa dari yang paling dekat ke yang paling jauh
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
// Fungsi pembanding tanggal, mengembalikan -1 jika t1 < t2, 0 jika sama, 1 jika t1 > t2
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

// Fungsi selection sort untuk mengurutkan bahan berdasarkan jumlah stok 
// I.S.: Array bahan makanan (A) berisi n data dengan jumlah stok tidak berurutan
// F.S.: Array A terurutkan berdasarkan jumlah stok dari yang terkecil ke terbesar
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
// Fungsi menghapus bahan yang sudah habis pakai dan menyimpannya ke list bahan habis pakai
// I.S.: Array bahan makanan (A) berisi n data, array bahan habis pakai (C) berisi m data, user memasukkan ID bahan yang habis pakai
// F.S.: Data bahan dengan ID yang sesuai dipindahkan ke array C dan dihapus dari array A, n berkurang 1, m bertambah 1, atau pesan error jika ID tidak ditemukan
func hapusMakananDanSimpan(A *arrMakanan, B *arrbhnHabisPakai, n *int, m *int) {
	var i,j,id,jumlah int
	var found bool 
	found = false

	fmt.Print("Masukkan ID bahan: ")
	fmt.Scan(&id)

	for i = 0; i < *n; i++ {
		if (*A)[i].id == id {
			found = true
			fmt.Print("Masukkn jumlah bahan yang sudah digunakan")
			fmt.Scan(&jumlah)
				
			// Salin data bahan yang habis ke array C
			(*B)[*m].idPakai = (*A)[i].id
			(*B)[*m].namaPakai = (*A)[i].nama
			(*B)[*m].jumlahPakai = jumlah
			(*B)[*m].kategoriPakai = (*A)[i].kategori
			(*B)[*m].kadarluarsaPakai = (*A)[i].kadarluarsa
			*m++
			if jumlah == (*A)[i].jumlah {
				for j = i; j < *n-1; j++ {
				(*A)[j] = (*A)[j+1]
				}
			*n--
			
		}else if jumlah > (*A)[i].jumlah{
			fmt.Print("Jumlah melebiihi stok yang ada")
			fmt.Print("Masukkn jumlah bahan yang sudah digunakan")
			fmt.Scan(&jumlah)
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
// Fungsi menampilkan daftar bahan yang sudah habis pakai
// I.S.: Array bahan habis pakai (C) berisi m data
// F.S.: Menampilkan semua data bahan habis pakai dari array C ke layar
func tampilkanMakananHbsPakai(A arrbhnHabisPakai, n int) {
    var i int
    for i = 0; i < n; i++ {
		fmt.Println("==========================================================")
		fmt.Println("ID 		 :", A[i].idPakai)
        fmt.Println("Kategori    :", A[i].kategoriPakai)
        fmt.Println("Nama        :", A[i].namaPakai)
        fmt.Println("Jumlah      :", A[i].jumlahPakai)
        fmt.Printf("Kadaluarsa  : %02d %02d %04d\n", A[i].kadarluarsaPakai.no,  A[i].kadarluarsaPakai.bulan, A[i].kadarluarsaPakai.tahun)
		fmt.Println("==========================================================")
		fmt.Println()
    }
}
