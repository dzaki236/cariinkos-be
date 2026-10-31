# CariinKos API

Dokumentasi REST API backend CariinKos. Semua path API menggunakan prefix `/api/v1`. Request dan response menggunakan JSON, kecuali endpoint unggah dan ganti foto yang menggunakan `multipart/form-data`.

## Base URL

```text
http://localhost:3000/api/v1
```

Sesuaikan host dan port dengan deployment. Health check berada di `GET /`.

## Autentikasi dan Role

Endpoint publik, login, register, dan dokumentasi tidak memerlukan token. Endpoint lain yang dilindungi memakai header berikut:

```http
Authorization: Bearer <token>
```

Skema `Bearer` tidak peka huruf besar/kecil. Header harus terdiri dari tepat dua bagian dan token tidak boleh kosong.

| Role | Akses |
|---|---|
| `owner` | Mengelola kos miliknya, fasilitas, foto, kontak, daftar user, dan status user biasa |
| `user` | Melihat serta mengelola review dan favorite miliknya |

Kesalahan autentikasi menghasilkan `401 Unauthorized`. Token valid dengan role yang tidak sesuai menghasilkan `403 Forbidden`.

Contoh error autentikasi:

```json
{"message":"Unauthorized","error":"Authorization header must use Bearer token format"}
```

```json
{"message":"Unauthorized","error":"Invalid or expired token"}
```

## Format Respons

Resource dikembalikan dalam properti `data`. Daftar resource menggunakan array di `data`; endpoint daftar kos juga mengembalikan `meta` pagination. Operasi tanpa resource menggunakan `message`.

Validasi gagal menggunakan `422` dan bentuk berikut:

```json
{
  "message": "Validation failed",
  "errors": {"name": ["The name field is required."]}
}
```

Error umum dari service dapat menyertakan `unrealmsg`. Resource yang tidak ditemukan mengembalikan `404` dengan `message: "Data not found"`. Body JSON yang gagal di-bind mengembalikan `400 Invalid request body`, kecuali `PUT /kos/{id}` yang saat ini mengembalikan `404` untuk kondisi tersebut.

## Ringkasan Endpoint

Semua path di bawah relatif terhadap base URL.

| Method | Path | Auth | Role | Keterangan |
|---|---|---|---|---|
| GET | `/` | Tidak | - | Health check, status `201` |
| GET | `/docs` | Tidak | - | Halaman dokumentasi HTML |
| POST | `/users/login` | Tidak | - | Login dan memperoleh JWT |
| POST | `/users/register` | Tidak | - | Registrasi user biasa |
| POST | `/users/logout` | Ya | Semua | Logout |
| PUT | `/users/profile` | Ya | Semua | Perbarui profil dan password |
| GET | `/users` | Ya | `owner` | Daftar seluruh user |
| PUT | `/users/{id}/status` | Ya | `owner` | Aktifkan/nonaktifkan user biasa |
| GET | `/public/kos` | Tidak | - | Daftar kos publik |
| GET | `/public/kos/{id}` | Tidak | - | Detail kos publik |
| GET | `/public/kos/{id}/reviews` | Tidak | - | Daftar review kos |
| GET | `/kos` | Ya | `owner` | Daftar kos |
| GET | `/kos/{id}` | Ya | `owner` | Detail kos |
| POST | `/kos` | Ya | `owner` | Buat kos |
| PUT | `/kos/{id}` | Ya | `owner` | Perbarui kos |
| DELETE | `/kos/{id}` | Ya | `owner` | Hapus kos |
| GET | `/facilities` | Ya | Semua | Daftar fasilitas |
| GET | `/facilities/{id}` | Ya | Semua | Detail fasilitas |
| POST | `/facilities` | Ya | Semua | Buat fasilitas |
| PUT | `/facilities/{id}` | Ya | Semua | Perbarui fasilitas |
| DELETE | `/facilities/{id}` | Ya | Semua | Hapus fasilitas |
| GET | `/kos/{id}/facilities` | Ya | `owner` | Daftar fasilitas kos |
| POST | `/kos/{id}/facilities` | Ya | `owner` | Kaitkan fasilitas ke kos |
| DELETE | `/kos/{id}/facilities/{facility_id}` | Ya | `owner` | Lepas fasilitas dari kos |
| GET | `/kos/{id}/photos` | Ya | `owner` | Daftar foto kos |
| POST | `/kos/{id}/photos` | Ya | `owner` | Unggah foto |
| PUT | `/kos/{id}/photos/{photo_id}` | Ya | `owner` | Ganti foto |
| DELETE | `/kos/{id}/photos/{photo_id}` | Ya | `owner` | Hapus foto |
| GET | `/kos/{id}/contacts` | Ya | `owner` | Lihat kontak kos |
| POST | `/kos/{id}/contacts` | Ya | `owner` | Buat kontak kos |
| PUT | `/kos/{id}/contacts` | Ya | `owner` | Perbarui kontak kos |
| DELETE | `/kos/{id}/contacts` | Ya | `owner` | Hapus kontak kos |
| POST | `/kos/{id}/reviews` | Ya | `user` | Buat review |
| PUT | `/kos/{id}/reviews` | Ya | `user` | Perbarui review milik user |
| DELETE | `/kos/{id}/reviews` | Ya | `user` | Hapus review milik user |
| GET | `/favorites` | Ya | `user` | Daftar favorite milik user |
| POST | `/kos/{id}/favorite` | Ya | `user` | Tambahkan favorite |
| PUT | `/kos/{id}/favorite` | Ya | `user` | Perbarui catatan favorite |
| DELETE | `/kos/{id}/favorite` | Ya | `user` | Hapus favorite |

## Auth dan User

### `POST /users/login`

Request JSON:

```json
{"email":"user@example.com","password":"password"}
```

`email` wajib berupa email valid dan `password` wajib berupa string. Sukses mengembalikan `200` dengan `message`, JWT `token`, dan `user`. Kredensial salah mengembalikan `401 Invalid email or password`.

### `POST /users/register`

Request JSON:

```json
{"name":"Pengguna","email":"user@example.com","password":"password-minimal-8"}
```

`name` dan `email` maksimal 255 karakter. Password minimal 8 karakter. Registrasi selalu membuat role `user`, mengembalikan `201`, dan tidak mengembalikan password. Email yang sudah terdaftar mengembalikan `409`.

### `POST /users/logout`

Memerlukan token. Sukses: `200 {"message":"Logout successful"}`.

### `PUT /users/profile`

Request JSON:

```json
{
  "name":"Nama Baru",
  "email":"baru@example.com",
  "old_password":"password-lama",
  "new_password":"password-baru",
  "new_password_confirmation":"password-baru"
}
```

`name` dan `email` wajib. Field password opsional; jika `new_password` dikirim, `old_password` wajib cocok dan konfirmasi harus sama. Password baru minimal 8 karakter. Sukses mengembalikan `200` dengan `message` dan `user`.

### Manajemen user owner

`GET /users` mengembalikan `{ "data": [...] }`. `PUT /users/{id}/status` menerima `{ "active": true }` atau `{ "active": false }`. Hanya user dengan role `user` yang dapat diubah statusnya; user lain menghasilkan `422`. Status aktif direpresentasikan dengan `disabled_at: null`, sedangkan user nonaktif memiliki timestamp pada `disabled_at`.

## Kos

`GET /public/kos` dan `GET /public/kos/{id}` tidak memerlukan autentikasi. Daftar publik hanya mengambil kos dengan status `published`. `GET /public/kos/{id}/reviews` juga publik dan mengembalikan daftar review beserta relasi user yang dimuat.

`GET /public/kos` menerima `page` default `1` dan `per_page` default `10`, maksimal `100`. Nilai query kurang dari `1` atau tidak valid kembali ke nilai default. Metadata berisi `page`, `per_page`, `total`, dan `last_page`.

### Field kos

| Field | Tipe | Aturan |
|---|---|---|
| `name` | string | Wajib, maksimal 255 |
| `description` | string/null | Opsional |
| `address` | string | Wajib |
| `province` | string | Wajib, maksimal 255 |
| `city` | string | Wajib, maksimal 255 |
| `district` | string | Wajib, maksimal 255 |
| `subdistrict` | string/null | Opsional |
| `gmaps_link` | string/null | Opsional |
| `price_monthly` | integer | Wajib, minimal 1 |
| `type` | string | Create opsional; update wajib; `mixed`, `man`, atau `woman` |
| `status` | string | Create opsional: `draft`, `published`, `archived`; update wajib: `draft` atau `published` |
| `room_size` | integer/null | Opsional |
| `total_rooms` | integer | Opsional, default nol jika tidak diisi |
| `available_rooms` | integer | Opsional, default nol jika tidak diisi |
| `minimum_rent_months` | integer | Opsional, default nol jika tidak diisi |
| `rules` | string/null | Opsional |

Create mengembalikan `201`, update `200`, dan delete `200` dengan message. List kos menyertakan `rattings`, `favorite_count`, serta `cover_photo`. `cover_photo` adalah URL foto pertama berdasarkan waktu pembuatan; jika belum ada foto, nilainya URL `public/images_not_found.jpg`. Detail dan create/update tidak mengisi `cover_photo`.

## Fasilitas

`POST /facilities` dan `PUT /facilities/{id}` menerima:

```json
{"name":"Wi-Fi","icon":"fa-solid fa-wifi"}
```

`name` wajib dan maksimal 255 karakter. `icon` opsional/null dan maksimal 255 karakter. Nilai icon disimpan sebagai string tanpa validasi kelas Font Awesome. Create mengembalikan `201`; list, detail, dan update `200`; delete mengembalikan message.

## Fasilitas Kos

`POST /kos/{id}/facilities` menerima `{ "facility_id": "<id>" }` dan mengembalikan `201 {"message":"Facility attached to kos successfully"}`. `facility_id` wajib berupa string. Delete mengembalikan `Facility detached from kos successfully`.

## Foto Kos

Kirim request sebagai `multipart/form-data` dengan field file `photo`. File harus JPEG, PNG, atau WebP, berukuran lebih dari 0 byte dan maksimal 10 MB. Create mengembalikan `201`, update `200`, dan delete `200`. Resource foto memiliki `id`, `kos_id`, `image_url`, `public_id`, `created_at`, dan `updated_at`.

## Kontak Kos

`POST` dan `PUT /kos/{id}/contacts` menerima:

```json
{"contact_name":"Penjaga Kos","phone":"+62221234567","whatsapp":"+6281234567890"}
```

`contact_name` dan `phone` wajib; `whatsapp` opsional/null. Batas masing-masing adalah 255 karakter untuk `contact_name` dan 50 karakter untuk nomor. Satu kos hanya memiliki satu kontak dan resource memakai `kos_id` sebagai identitas.

## Review

`POST` dan `PUT /kos/{id}/reviews` menerima:

```json
{"rating":5,"comment":"Lokasi bersih dan dekat kampus"}
```

Keduanya hanya untuk role `user`. Saat create, `rating` wajib bernilai 1 sampai 5 dan `comment` wajib string maksimal 2000 karakter. Satu user hanya dapat memiliki satu review untuk satu kos. `POST` mengembalikan `201`, `PUT` mengembalikan `200`, dan `DELETE` mengembalikan message.

## Favorite

`GET /favorites` mengembalikan favorite milik user beserta relasi kos yang dimuat. `POST` dan `PUT /kos/{id}/favorite` menerima `{ "note": "Dekat kantor" }`. Field `note` opsional. Satu user hanya dapat memiliki satu favorite untuk satu kos. Create mengembalikan `201`, update `200`, dan delete `200` dengan message.

## Contoh cURL

Login:

```bash
curl -X POST http://localhost:3000/api/v1/users/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"password"}'
```

Daftar kos publik:

```bash
curl "http://localhost:3000/api/v1/public/kos?page=1&per_page=10"
```

Buat review:

```bash
curl -X POST http://localhost:3000/api/v1/kos/<kos-id>/reviews \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"rating":5,"comment":"Dekat kampus"}'
```

Unggah foto:

```bash
curl -X POST http://localhost:3000/api/v1/kos/<kos-id>/photos \
  -H "Authorization: Bearer <token>" \
  -F "photo=@./kamar.jpg"
```

## Sumber Dokumentasi

OpenAPI tersedia pada [openapi.json](openapi.json). Halaman HTML API tersedia di `GET /api/v1/docs`. Dokumentasi ini mengikuti route yang terdaftar di `routes/api.go` dan `routes/api/`.
