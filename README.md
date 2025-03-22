# Golang Skill Test PT. Bima Sakti Multi Sinergi

README ini menyediakan informasi tentang versi Go yang digunakan dalam proyek ini beserta instruksi untuk instalasi dan menjalankan aplikasi.

## Go Version

Proyek ini menggunakan Go versi 1.23.1. Disarankan untuk menggunakan versi ini atau yang lebih baru untuk memastikan kompatibilitas.

## Prerequisites

Sebelum memulai, pastikan Anda telah menginstal:
- [Go](https://golang.org/dl/) (versi 1.23.1 atau lebih baru)
- Git (untuk mengkloning repositori)

## Installation

Ikuti langkah-langkah berikut untuk menginstal proyek:

1. Kloning repositori:
   ```
   git clone https://github.com/Kalveir/Skill-Test.git
   cd Skill-Test
   ```

2. Instal dependensi:
   ```
   go mod tidy
   ```

## Running the Code

Untuk menjalankan script, gunakan perintah berikut:

```
go run <script-name>.go
```

Atau, Anda dapat membuat dan kemudian menjalankan executable:

```
go build -o app <script-name>.go
./app
```