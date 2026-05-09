# Hands-On

Folder ini akan dikerjakan sendiri oleh peserta.

Project ini menggunakan:

- Golang
- MySQL

Fitur aplikasi:

- GET workshop
- CREATE workshop

Endpoint tersedia:

```txt
GET  /workshops
POST /workshops
````

---

# Struktur Project

Di folder ini terdapat:

* `Dockerfile`
* `multi.Dockerfile`
* `docker-compose.yaml`

Namun file-file tersebut belum lengkap.

Isinya hanya berupa komentar step-by-step.

Tugas peserta adalah mengubah komentar tersebut menjadi konfigurasi Docker yang benar.

---

# Step 1 — Masuk ke Folder Hands-On

```bash
cd workshop-container/hands-on
```

---

# Step 2 — Copy Environment File

```bash
cp .env.example .env
```

Lalu sesuaikan value database.

Contoh:

```env
DB_URL=<user>:<password>@tcp(<host>:<port)/<name>
```

---

# Step 3 — Lengkapi Dockerfile

Lengkapi:

* `Dockerfile`
* `multi.Dockerfile`

Sesuai instruksi komentar yang tersedia.

---

# Step 4 — Lengkapi Docker Compose

Lengkapi file:

```txt
docker-compose.yaml
```

Buat:

* Service app
* Service database MySQL

Pastikan:

* App dapat terhubung ke database
* Port dapat diakses
* Environment variable terbaca
* Database memiliki volume

---

# Step 5 — Build dan Jalankan Container

```bash
docker compose up --build
```

---

# Step 6 — Masuk ke Container Database

Jika container berhasil berjalan, masuk ke container database.

Contoh:

```bash
docker exec -it <nama_container> bash
```

Masuk ke MySQL:

```bash
mysql -u <user> -p
```

Pilih database:

```sql
USE <nama_db>
```

---

# Step 7 — Jalankan SQL

Copy isi file:

```txt
db.sql
```

Lalu jalankan di MySQL.

Tujuannya untuk membuat tabel yang dibutuhkan aplikasi.

---

# Step 8 — Testing Endpoint

Contoh request sudah disediakan di file:

```txt
createWorkshop.http
getWorkshops.http
```

Endpoint dapat dites menggunakan:

* Postman
* Bruno
* Insomnia
* VSCode REST Client

Jika berhasil:

* Data workshop dapat dibuat
* Data workshop dapat diambil
