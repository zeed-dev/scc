# 🏥 Smart Command Center

Sistem backend berbasis Go (Golang) untuk mendukung Smart Command Center. Backend ini bertanggung jawab untuk menangani data pasien darurat, panic button, pelacakan ambulans, serta integrasi real-time dan peta.

## 🚀 Fitur Utama Backend

- Autentikasi & Manajemen Pengguna (Admin, Petugas, Pasien)
- Penerimaan Panic Button (lokasi, identitas, audio)
- Penghitungan rute ambulans otomatis (Google Maps API)
- Update posisi ambulans real-time
- Push Notification (Firebase Cloud Messaging) / MQTT
- Logging insiden lengkap & statistik laporan
- Integrasi dengan layanan eksternal (BPBD, Ambulans Swasta)
- Dashboard API untuk menampilkan data peta & status insiden

## 🛠️ Teknologi yang Digunakan

| Komponen         | Teknologi                    |
|------------------|------------------------------|
| Bahasa Backend   | Go (Golang)                  |
| Framework Web    | [Gin](https://github.com/gin-gonic/gin)          |
| Real-time        | [MQTT](https://mqtt.org/) / [Firebase Cloud Messaging](https://firebase.google.com/docs/cloud-messaging) |
| Database         | MySql  |
| Storage Rekaman  | Firebase Storage / Cloudinary |
| Maps & Routing   | Google Maps API              |
| Auth             | JWT (JSON Web Tokens)        |

## 📂 Struktur Proyek

```
smart-command-center-backend/
├── cmd/                    # Entry point (main.go)
├── config/                 # Konfigurasi dan environment
├── controllers/            # HTTP handlers / controller logic
├── models/                 # Model database & schema
├── routes/                 # HTTP routes
├── services/               # Business logic (panic button, maps, notifikasi, dll)
├── utils/                  # Helper dan utilities (JWT, validasi, logger)
├── storage/                # Upload rekaman suara
├── mqtt/                   # Handler dan listener MQTT
├── firebase/               # Integrasi Firebase FCM
└── README.md
```

## ⚙️ Instalasi & Penggunaan

### 1. Clone Repository
```bash
git clone https://github.com/your-username/smart-command-center-backend.git
cd smart-command-center-backend
```

### 2. Atur Environment
Buat file `.env`:
```env
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=smart_center
JWT_SECRET=your_jwt_secret
GOOGLE_MAPS_API_KEY=your_maps_api_key
FIREBASE_CREDENTIALS=path/to/your/firebase.json
```

### 3. Jalankan Aplikasi
```bash
go run cmd/main.go
```

## 🧪 Contoh Endpoint API

| Method | Endpoint                | Deskripsi                                |
|--------|--------------------------|-------------------------------------------|
| POST   | `/api/v1/auth/login`     | Login pengguna                            |
| POST   | `/api/v1/panic`          | Kirim panic button                        |
| GET    | `/api/v1/incidents`      | Ambil daftar insiden                      |
| GET    | `/api/v1/maps/route`     | Ambil rute ambulans                       |
| POST   | `/api/v1/ambulance/track`| Update lokasi ambulans secara real-time   |

## 🧩 Potensi Integrasi Eksternal

- Dinas Kesehatan: Webhook laporan insiden.
- BPBD: Sinkronisasi status bencana.
- Layanan Ambulans Swasta: Notifikasi otomatis via MQTT.

## ✅ To-do (Roadmap)

- [x] Struktur dasar backend
- [x] Endpoint Login
- [x] Endpoint Role
- [x] Endpoint CRUD User
- [x] Endpoint kirim panic button (lihat `routes/panic_routes.go`)
- [ ] Integrasi Google Maps API untuk routing
- [ ] Realtime tracking ambulans
- [ ] Firebase Notification
- [ ] Dashboard Analytics Endpoint

## 👨‍💻 Kontributor

- Muhammad Ziad Alfian – Fullstack Developer

## 📄 Lisensi

MIT License – Silakan gunakan dan modifikasi sesuai kebutuhan.

> Smart Command Center: Membantu menyelamatkan nyawa dengan kecepatan, akurasi, dan koordinasi teknologi.