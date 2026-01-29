# Aplikasi POS

Buatlah sebuah aplikasi Point of Sales (POS) App berbasis RESTful API menggunakan bahasa pemrograman Golang. Aplikasi ini dirancang untuk membantu toko atau restoran dalam mengelola data staf, menu makanan/minuman, pemesanan, serta pelacakan penjualan dasar.

## Berikut fitur-fitur yang harus tersedia:

### Authentication

- API untuk Login.
- API untuk pengecekan email.
- API untuk validasi OTP.
- API untuk reset password.

### Dashboard

- Ringkasan penjualan harian, bulanan, jumlah meja.
- Daftar produk populer.
- Daftar produk baru (produk < 30 hari).

### Staff Management

- List, tambah, edit, detail, dan hapus data staf atau admin.
- Menggunakan pagination dan sorting berdasarkan nama atau email.
- Admin yang dihapus tidak dapat login lagi.

### Menu

- API kategori menu (CRUD).
- API produk makanan/minuman (CRUD).
- Filter berdasarkan kategori.
- Menggunakan pagination.

### Orders

- List order, tambah order, edit, hapus.
- List table dan payment method.
- Kursi hanya menampilkan meja yang tersedia.
- Data meja disiapkan via seeder.
- Pajak dibuat statis.
- Edit order memungkinkan pengubahan nama pemesan, produk & jumlah, serta metode pembayaran.

### Inventory

- List produk inventaris.
- Tambah, edit, dan hapus data inventaris.
- Tambahkan parameter filter untuk pencarian data produk.

### Notification

- List notifikasi.
- Update status notifikasi (new dan readed).
- Hapus notifikasi.

### Profile

- Edit profil user.
- List data admin.
- Edit akses admin (khusus superadmin).
- Logout.
- Password dikirim via email saat akun admin dibuat.

### Revenue Report

- Total revenue dan breakdown berdasarkan status.
- Total revenue per bulan.
- List produk beserta detail revenue.

### Reservation

- List, tambah, update, dan detail data reservasi.
- Update hanya bisa mengubah meja dan status pembatalan.

## Ketentuan Umum

- Project dikerjakan secara berkelompok sesuai pembagian tim.
- Bahasa pemrograman: Golang.
- Framework API: Gin.
- Tools bebas, akses API via Postman diperbolehkan.
- Kode harus ditulis dalam gaya idiomatik Go dan disimpan dalam struktur folder yang rapi.
- Gunakan komentar seperlunya untuk membantu pembacaan kode.

## Ketentuan Utama

Project harus mencakup dan mengimplementasikan hal-hal berikut:

- Gunakan PostgreSQL untuk penyimpanan data permanen.
- Gunakan ORM GORM untuk interaksi database.
- Implementasi automigrate dan seeder data awal.
- Validasi input menggunakan tag binding, dan tampilkan error spesifik pada input yang salah.
- Gunakan status code HTTP yang sesuai untuk setiap kondisi response.
- Gunakan Zap Logger untuk mencatat log penting dan error.
- Implementasi pagination untuk endpoint list.
- Implementasi graceful shutdown saat aplikasi dimatikan.
- Implementasi fitur pengiriman email (contoh: OTP, notifikasi password admin).

## Ketentuan Tambahan

- Terapkan prinsip DRY (Don’t Repeat Yourself) untuk menghindari pengulangan kode.
- Buat unit test untuk setiap function, dengan target code coverage rata-rata 50% ke atas.

## Ketentuan Pengumpulan

- Upload ke akun GitHub masing-masing dengan nama repository project-POS-APP-golang-nama-team
- Sertakan:
  - Collection Postman untuk pengujian endpoint.
  - File database dalam format .sql atau database plan.
  - Board Trello dan backlog task yang dibagikan ke: email admin (sudah ada)
  - Link Figma (referensi desain): (Link nya sudah saya dapatkan)
