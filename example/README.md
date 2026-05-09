# Example

Folder ini berisi aplikasi sederhana menggunakan:

- Node.js
- Express
- TypeScript
- PostgreSQL

Fitur aplikasi:

- GET notes
- CREATE notes

## Dockerfile

Di folder ini terdapat 2 Dockerfile:

- `Dockerfile` → Single-stage build
- `multi.Dockerfile` → Multi-stage build

---

# Step 1 — Clone Repository

```bash
git clone https://github.com/mqqff/workshop-container.git
````

---

# Step 2 — Masuk ke Folder Example

```bash
cd workshop-container/example
```

---

# Step 3 — Copy File Environment

```bash
cp .env.example .env
```

Lalu sesuaikan isi `.env`.

Contoh:

```env
PORT=3000

DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=app
```

Pastikan value database sesuai.

---

# Step 4 — Jalankan Docker Compose

```bash
docker compose up --build
```

Jika berhasil, aplikasi akan berjalan di:

```txt
http://localhost:3000
```

---

# Step 5 — Testing Endpoint

Testing endpoint bisa menggunakan:

* Postman
* Bruno
* Insomnia
* VSCode REST Client
* dan tools lainnya

Contoh request sudah disediakan di file:

```txt
createNote.http
getNotes.http
```

Endpoint yang tersedia:

```txt
GET  /notes
POST /notes
```

---

# Step 6 — Menggunakan Multi-Stage Dockerfile

Sekarang kita akan mencoba menggunakan multi-stage build.

Buka file:

```txt
docker-compose.yaml
```

Lalu ubah:

```yaml
build:
  context: .
  dockerfile: Dockerfile
```

Menjadi:

```yaml
build:
  context: .
  dockerfile: multi.Dockerfile
```

---

# Step 7 — Build Ulang

Jalankan kembali:

```bash
docker compose up --build
```

Sekarang aplikasi menggunakan multi-stage Docker build.
