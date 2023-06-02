# Aplikasi Enigma Laundry

Aplikasi Enigma Laundry adalah sebuah aplikasi untuk mengelola layanan dan pesanan laundry.

### Cara menggunakan project ini adalah sebagai berikut :
- Pastikan Go sudah terinstal pada komputer beserta extensions yang ada di text editor VSCode
- Clone repositori ini ke direktori lokal Kamu.
- Buka terminal dan arahkan ke direktori proyek.
- Jalankan perintah berikut untuk menginstal dependensi:
```
go mod download
```

### Pengaturan Database :
- Pastikan PostgreSql di pgAdmin sudah berjalan dengan baik.
- Buatlah database baru dengan nama "enigmalaundry".
- Konfigurasikan koneksi database di file `config/database.go` dengan mengubah nilai variabel `dsn` sesuai dengan pengaturan PostgreSql Anda

### Penggunaan
- Jalankan aplikasi ini dengan melakukan
```
go run enigma-laundry.go
```
- Aplikasi akan menampilkan menu utama
- Jika Anda telah memasuki bagian Submenu dan ingin kembali ke menu utama, maka anda harus keluar terlebih dahulu
- Ikuti intruksi yang telah disediakan untuk melakukan inputan

## Lisensi

Proyek ini dilisensikan di bawah [MIT License](LICENSE).