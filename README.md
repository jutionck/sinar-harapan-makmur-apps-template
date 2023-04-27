# Sinar Harapan Makmur Apps

Adalah aplikasi Jual Beli Kendaraan

## Note

- Ini dipersiapkan sebagai template untuk melakuan _Unit Testing_

## Run Apps

Aplikasi juga dapat berjalan dan digunakan, beberapa hal yang perlu dipersiapkan sebelum running apps:

1. Silahkan copy file `.env.example` dan buat file baru dengan nama `.env` kemudian sesuaikan dengan konfigurasi di device teman-teman.
2. Untuk menjalankan pertama kali pastikan di file `.env` pada bagian ini `ENV=MIGRATION` dan silahkan lakukan `go run .` agar tabel-tabel ter-migrasi.
3. Jika sudah silahkan ubah `ENV=MIGRATION` menjadi `ENV=DEV` dan bisa menjalankan `go run .`

## Setup

Beberapa yang perlu disiapkan sebelum melakukan unit testing:

1. Install dependency `testify`:

```bash
go get github.com/stretchr/testify
```

2. Install dependency `sql-mock`:

```bash
go get github.com/DATA-DOG/go-sqlmock
```

## Run Test

Untuk menjalankan unit testing bisa gunakan command-command berikut:

1. Menjalankan semua unit testing:

```bash
go test -v ./...
```

2. Menjalankan fungsi coverage:

```bash
go test ./... -coverprofile=cover.out  && go tool cover -html=cover.out
```
