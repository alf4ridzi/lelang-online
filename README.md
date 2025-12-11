# 🏷️ Lelang Online

**Lelang Online** adalah aplikasi lelang real-time yang dibuat
menggunakan **Java (NetBeans)** sebagai frontend dan **Golang** sebagai
backend. Proyek ini dikembangkan untuk memenuhi tugas akhir semester
ganjil kelas 12.

## 🚀 Teknologi yang Digunakan

### Backend

-   **WebSocket** --- Koneksi real-time
-   **GIN** --- Web Framework
-   **GORM** --- ORM untuk manajemen database
-   **Air** --- Live reload untuk mempercepat proses pengembangan\
    Repo: https://github.com/air-verse/air

### Frontend

-   **Java Swing (NetBeans GUI Builder)**

## 📦 Instalasi

Clone repository:

``` bash
git clone https://github.com/alf4ridzi/lelang-online
cd lelang-online
```

## ⚙️ Setup Backend

``` bash
cd backend
go mod tidy
go install github.com/air-verse/air@latest
```

Seed database:

``` bash
go run cmd/main.go --seed
```

Jalankan server:

``` bash
air
```

## 🖥️ Setup Frontend

1.  Download NetBeans:\
    https://netbeans.apache.org/front/main/download

2.  Buka project:\
    `NetBeans → Open Project → pilih folder frontend`


## ✨ Fitur Utama

-   Lelang real-time menggunakan WebSocket
-   Manajemen item lelang
-   Penawaran (bid) langsung
-   Sinkronisasi backend ↔ frontend

## 📜 Lisensi

MIT
