# MyFoodList - Aplikasi Pencatatan Persediaan Makanan

Selamat datang di MyFoodList! Aplikasi sederhana ini membantu Anda mencatat dan mengelola persediaan bahan makanan dan minuman Anda. Dengan MyFoodList, Anda dapat dengan mudah menambahkan, mengubah, melihat, mencari, dan mengurutkan item makanan, serta melacak item yang akan segera kadaluwarsa atau yang telah habis digunakan.

## Fitur Utama

* **Pencatatan Awal:** Memasukkan daftar bahan makanan yang ada saat pertama kali menggunakan aplikasi.
* **Tambah Item:** Menambahkan item makanan baru atau menambah jumlah stok item yang sudah ada.
* **Ubah Item:** Memperbarui detail item makanan (ID, nama, kategori, jumlah, tanggal kadaluwarsa).
* **Tampilkan Persediaan:** Melihat daftar lengkap bahan makanan yang tersedia dan daftar bahan yang sudah habis dipakai.
* **Cari Item:** Mencari item makanan tertentu berdasarkan nama atau ID.
* **Peringatan Kadaluwarsa:** Menampilkan item makanan yang akan kadaluwarsa dalam 30 hari ke depan.
* **Urutkan Item:** Mengurutkan daftar makanan berdasarkan tanggal kadaluwarsa atau jumlah stok.
* **Catat Pemakaian:** Memindahkan item yang telah digunakan ke daftar "Habis Pakai" dan mengurangi stoknya atau menghapusnya dari persediaan utama.

## Struktur Data

Program ini menggunakan beberapa struktur data utama:

* `Makanan`: Menyimpan detail setiap item makanan (ID, nama, jumlah, kategori, dan tanggal kadaluwarsa).
* `tanggal`: Menyimpan detail tanggal (nomor hari, bulan, tahun).
* `bhnHabisPakai`: Menyimpan detail item makanan yang telah habis dipakai.
* Array (`arrMakanan`, `arrbhnHabisPakai`): Digunakan untuk menyimpan kumpulan data makanan dan bahan habis pakai.

### 1. Interaksi dengan Program

1.  **Mulai Pencatatan:**
    * Saat program dimulai, Anda akan ditanya: `Apakah anda ingin mencatat persediaan anda? (1 = Ya, 0 = Tidak)`
    * Masukkan `1` untuk melanjutkan atau `0` untuk keluar.

2.  **Input Data Awal (jika memilih "Ya"):**
    * Program akan meminta: `Masukkan jumlah data yang ingin anda masukkan: `
    * Untuk setiap item, Anda akan diminta memasukkan:
        * `ID makanan / minuman:` (harus angka antara 100-999)
        * `kategori bahan:` (misalnya, Sayuran, Buah, Minuman, Daging, Bumbu)
        * `nama bahan:` (misalnya, Apel Fuji, Susu UHT Coklat)
        * `jumlah bahan:` (harus lebih dari 0)
        * `tanggal kadaluarsa baru (dd mm yyyy):` (misalnya, 25 12 2025)
            * Hari (dd) harus valid (1-31).
            * Bulan (mm) harus valid (1-12).

3.  **Menu Utama:**
    Setelah input data awal (atau jika Anda sudah pernah input sebelumnya dan data ada), menu utama akan tampil:
    ```
    ==========================================================
    |                               MENU                                 |
    ==========================================================
    1. Menambah bahan makanan / minuman
    2. Mengubah bahan makanan / minuman
    3. Menampilkan persediaan bahan makanan / minuman
    4. Mencari bahan makanan tertentu
    5. Menampilkan kadaluarsa bahan dalam 30 hari ke depan
    6. Mengurutkan bahan makanan / minuman
    7. Masukkan bahan yang habis dipakai
    0. Keluar
    Apa yang ingin anda lakukan:
    ```
    Masukkan nomor pilihan Anda dan tekan Enter.

    * **Pilihan 1: Menambah bahan makanan / minuman**
        * Anda akan ditanya:
            ```
            1.Menambahkan jumlah bahan
            2.Menambahkan bahan
            Anda ingin menambahkan berdasarkan :
            ```
        * Pilih `1` untuk menambah jumlah stok item yang sudah ada (Anda akan diminta ID item dan jumlah tambahan).
        * Pilih `2` untuk menambah item makanan baru (Anda akan diminta berapa banyak item baru dan detail untuk masing-masing item seperti pada input data awal).

    * **Pilihan 2: Mengubah bahan makanan / minuman**
        * Masukkan `ID bahan yang ingin diubah:`.
        * Jika ID ditemukan, Anda akan diminta untuk memasukkan semua data baru untuk item tersebut (ID baru, kategori baru, nama baru, jumlah baru, tanggal kadaluwarsa baru).

    * **Pilihan 3: Menampilkan persediaan bahan makanan / minuman**
        * Program akan menampilkan dua daftar:
            1.  `Daftar bahan makanan yang tersedia:` (item yang masih ada di stok).
            2.  `Daftar bahan makanan yang sudah digunakan:` (item yang telah dicatat sebagai habis pakai).

    * **Pilihan 4: Mencari bahan makanan tertentu**
        * Anda akan ditanya:
            ```
            1. Berdasarkan nama
            2. Berdasarkan ID
            Pilihan pencarian:
            ```
        * Pilih `1` untuk mencari berdasarkan nama (masukkan nama bahan). *Catatan: Untuk hasil optimal pencarian nama (binary search), idealnya daftar sudah terurut berdasarkan nama, meskipun program saat ini tidak mengurutkan otomatis sebelum pencarian nama.*
        * Pilih `2` untuk mencari berdasarkan ID (masukkan ID bahan).
        * Jika ditemukan, program akan mengonfirmasi dan kemudian menampilkan seluruh daftar makanan saat ini.

    * **Pilihan 5: Menampilkan kadaluarsa bahan dalam 30 hari ke depan**
        * Masukkan `tanggal hari ini (dd mm yyyy):`.
        * Program akan menampilkan daftar item yang akan kadaluwarsa dalam 30 hari dari tanggal yang Anda masukkan.

    * **Pilihan 6: Mengurutkan bahan makanan / minuman**
        * Anda akan ditanya:
            ```
            1. Tanggal kadaluarsa
            2. Jumlah stok bahan
            Pilihan pengurutan:
            ```
        * Pilih `1` untuk mengurutkan berdasarkan tanggal kadaluwarsa (dari yang paling dekat).
        * Pilih `2` untuk mengurutkan berdasarkan jumlah stok (dari yang paling sedikit).
        * Setelah diurutkan, daftar makanan akan ditampilkan.

    * **Pilihan 7: Masukkan bahan yang habis dipakai**
        * Masukkan `ID bahan:` yang telah digunakan.
        * Masukkan `jumlah bahan yang sudah digunakan:`.
        * Jika jumlah yang digunakan sama dengan stok, item akan dipindahkan sepenuhnya ke daftar "Habis Pakai".
        * Jika jumlah yang digunakan kurang dari stok, stok item akan dikurangi, dan sejumlah yang digunakan akan dicatat di daftar "Habis Pakai".
        * Jika jumlah yang digunakan melebihi stok, akan ada pesan error.

    * **Pilihan 0: Keluar**
        * Program akan berhenti dan menampilkan pesan `Program Selesai : TERIMA KASIH`.

### Contoh Sesi Penggunaan

1.  Jalankan program.
2.  Pilih `1` untuk mulai mencatat.
3.  Masukkan jumlah data, misal `1`.
4.  Masukkan detail makanan: ID `101`, kategori `Buah`, nama `Apel`, jumlah `10`, kadaluwarsa `30 06 2025`.
5.  Dari menu utama, pilih `3` untuk menampilkan. Anda akan melihat item Apel.
6.  Pilih `1` (Menambah), lalu `2` (Menambahkan bahan). Tambah `1` item baru.
7.  Masukkan detail makanan baru: ID `102`, kategori `Minuman`, nama `Susu`, jumlah `5`, kadaluwarsa `15 06 2025`.
8.  Pilih `5` (Kadaluwarsa). Masukkan tanggal hari ini (misal `01 06 2025`). Jika ada, item yang mendekati kadaluwarsa akan tampil.
9.  Pilih `7` (Habis dipakai). Masukkan ID `101` (Apel), jumlah `3`.
10. Pilih `3` lagi. Anda akan lihat stok Apel berkurang dan item Apel (sebanyak 3) muncul di daftar habis pakai.
11. Pilih `0` untuk keluar.

Semoga panduan ini membantu Anda menggunakan aplikasi MyFoodList!
