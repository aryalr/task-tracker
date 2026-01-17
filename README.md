# Task Tracker CLI (WIP)

Aplikasi Command Line Interface (CLI) sederhana untuk mengelola tugas, ditulis menggunakan Go. Proyek ini sedang dalam tahap pengembangan aktif.

## Deskripsi

Task Tracker CLI bertujuan untuk memberikan cara cepat dan mudah untuk mencatat, melihat, memperbarui, dan menghapus tugas langsung dari terminal. Data tugas akan disimpan secara lokal dalam file JSON (`mytask.json`).

## Fitur (Direncanakan & Berjalan)

Berikut adalah status fitur berdasarkan spesifikasi pengembangan:

- [x] **Routing Perintah Dasar**: Struktur dasar untuk menangani input perintah.
- [ ] **Add Task**: Menambahkan tugas baru dengan prioritas (p1, p2, p3). (Implementasi parsial)
- [ ] **List Tasks**: Menampilkan daftar tugas yang tersimpan. (Implementasi parsial)
- [ ] **Update Task**: Memperbarui nama atau prioritas tugas berdasarkan ID.
- [ ] **Delete Task**: Menghapus tugas berdasarkan ID.
- [ ] **Mark Done**: Menandai tugas sebagai selesai.
- [ ] **Penyimpanan Data**: Persistensi data ke file `mytask.json`.

## Cara Menjalankan (Development)

Saat ini, Anda dapat menjalankan aplikasi langsung menggunakan perintah `go`:

### Prasyarat
- Go 1.25 atau lebih baru.

### Perintah Dasar

```bash
# Menambahkan Tugas (Contoh)
go run . add "Nama Tugas" p1

# Melihat Daftar Tugas
go run . list
```

## Spesifikasi Perintah

Aplikasi ini dirancang dengan struktur perintah sebagai berikut (lihat `spec.md` untuk detail lengkap):

- `add <task-name> <priority>`: Tambah tugas.
  - Prioritas: `p1` (High), `p2` (Medium), `p3` (Low).
- `list`: Lihat semua tugas.
- `update <id> <task-name> <priority>`: Update tugas.
- `delete <id>`: Hapus tugas.
- `done <id>`: Tandai tugas selesai.

## Struktur Proyek

- `main.go`: Entry point aplikasi dan routing perintah.
- `handleInput.go`: Logika penanganan input untuk setiap perintah.
- `spec.md`: Spesifikasi desain dan fitur aplikasi.
