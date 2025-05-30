# Aplikasi Manajemen Stok Bahan Makanan

## Deskripsi Aplikasi

Aplikasi ini membantu Anda mengelola persediaan makanan di rumah dengan lebih mudah. Catat detail bahan, pantau stok, dapatkan pengingat kedaluwarsa, dan gunakan fitur cari serta urut untuk mengurangi pemborosan dan mengoptimalkan belanja.

## Cara Mengoperasikan Aplikasi

Aplikasi ini dijalankan melalui menu teks:

* **Input Data Awal:** Masukkan detail semua bahan makanan Anda (ID, nama, kategori, stok, tgl. kedaluwarsa).
* **Menu 1: Tambah Bahan**
    * **Submenu 1 (`tambahjumlah`):** Update jumlah stok bahan yang sudah ada (berdasarkan ID).
    * **Submenu 2 (`menambahMakanan`):** Tambah bahan makanan baru ke sistem.
* **Menu 3: Ubah Data (`ubahData`):** Modifikasi data bahan yang sudah ada (ID, nama, kategori, jumlah, tgl. kedaluwarsa) berdasarkan ID.
* **Menu 4: Tampilkan Stok & Laporan**
    * **`tampilkanMakanan`:** Lihat semua stok aktif.
    * **`tampilkanMakananHbsPakai`:** Lihat riwayat bahan yang sudah habis atau digunakan.
* **Menu 5: Cari Bahan**
    * **Submenu 1 (`cariBahan`):** Cari bahan berdasarkan nama.
    * **Submenu 2 (`cariId`):** Cari bahan berdasarkan ID.
* **Menu 6: Peringatan Kedaluwarsa (`kadaluarsa`):** Cek bahan yang akan kedaluwarsa dalam 30 hari (dihitung dari tanggal yang Anda masukkan).
* **Menu 7: Urutkan Data**
    * **Submenu 1 (`insertionSortTanggal`):** Urutkan bahan berdasarkan tanggal kedaluwarsa (paling dekat).
    * **Submenu 2 (`selectionSortJumlah`):** Urutkan bahan berdasarkan jumlah stok (paling sedikit).
* **Menu 8: Hapus & Simpan Habis Pakai (`hapusMakananDanSimpan`):** Kurangi stok bahan yang digunakan. Jika habis, data dipindah ke laporan bahan habis pakai.
* **Menu 0: Keluar:** Mengakhiri aplikasi.

## Tujuan Aplikasi
* Mengimplementasikan konsep algoritma pemrograman dalam aplikasi praktis sehari-hari.
* Menyediakan fitur manajemen data stok bahan makanan (tambah, edit, hapus, lihat).
* Menerapkan algoritma pengurutan (`Selection Sort`, `Insertion Sort`) untuk data bahan.
* Mengimplementasikan algoritma pencarian (`Sequential Search`, `Binary Search`) untuk data bahan.
* Menyajikan laporan stok yang ringkas dan informatif.

## Tautan GitHub Proyek
[Masukkan tautan ke repositori GitHub Anda di sini, contoh: https://github.com/nicotainee/review-1-tubes-alpro]
