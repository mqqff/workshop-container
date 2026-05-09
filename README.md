# README.md

# Workshop Container

Repository ini berisi 2 folder utama:

- `example` → Akan didemokan bersama-sama saat workshop.
- `hands-on` → Akan dikerjakan sendiri oleh peserta.

## Struktur Folder

```txt
├── example/
└── hands-on/
```

## Penjelasan

### Example

Folder `example` berisi aplikasi sederhana menggunakan:

* Node.js
* Express
* TypeScript
* PostgreSQL

App ini hanya memiliki 2 endpoint:

* GET notes
* CREATE notes

Pada sesi ini peserta akan melihat:

* Cara menggunakan Dockerfile single-stage
* Cara menggunakan Dockerfile multi-stage
* Cara menjalankan PostgreSQL menggunakan Docker Compose
* Cara melakukan build dan run container

➡️ Lanjut ke dokumentasi:

* [README Example](./example/README.md)

---

### Hands-On

Folder `hands-on` berisi project sederhana menggunakan:

* Golang
* PostgreSQL

Peserta akan mengerjakan sendiri:

* Dockerfile single-stage
* Dockerfile multi-stage
* Docker Compose
* Build dan menjalankan container

➡️ Lanjut ke dokumentasi:

* [README Hands-On](./hands-on/README.md)
