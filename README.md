# CariinKos API Docs

Dokumentasi REST API backend CariinKos. API menggunakan JSON kecuali endpoint unggah foto dan seluruh path API diawali `/api/v1`.

## Base URL

```text
http://localhost:3000/api/v1
```

Sesuaikan host dan port dengan konfigurasi deployment. Endpoint `GET /` pada base path mengembalikan `{"halo":"api"}`.

## Autentikasi

Endpoint login dan dokumentasi bersifat publik. Semua endpoint kos dan fasilitas memerlukan JWT dengan header:

```http
Authorization: Bearer <token>
```

Ganti `<token>` dengan JWT yang diterima dari `POST /login`. Header `Authorization` wajib berisi tepat dua bagian: skema `Bearer` dan token, dipisahkan oleh spasi. Skema tidak peka huruf besar/kecil, sehingga `Bearer` maupun `bearer` diterima. Token yang kosong, header yang tidak menggunakan format tersebut, token tidak valid, atau token kedaluwarsa menghasilkan `401 Unauthorized`.

Contoh respons jika format header salah atau tidak ada:

```json
{
  "message": "Unauthorized",
  "error": "Authorization header must use Bearer token format"
}
```

Contoh respons jika token tidak valid atau kedaluwarsa:

```json
{
  "message": "Unauthorized",
  "error": "Invalid or expired token"
}
```

## Format Respons

Respons resource umumnya dibungkus dengan `data`:

```json
{
  "data": {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "user_id": "123e4567-e89b-12d3-a456-426614174001",
    "name": "Kos Melati",
    "address": "Jl. Merdeka No. 10",
    "province": "Jawa Barat",
    "city": "Bandung",
    "district": "Coblong",
    "room_size": 20,
    "price_monthly": 1500000,
    "total_rooms": 10,
    "available_rooms": 2,
    "minimum_rent_months": 1,
    "type": "woman",
    "status": "published",
    "created_at": "2026-10-30T08:00:00Z",
    "updated_at": "2026-10-30T08:00:00Z"
  }
}
```

Operasi yang tidak mengembalikan resource memberikan `message`. Validasi gagal menggunakan status `422` dan bentuk:

```json
{
  "message": "Validation failed",
  "errors": {}
}
```

Error umum lainnya memakai `message`. Resource yang tidak ditemukan mengembalikan `{"message":"Data not found"}`; error server yang tidak terduga mengembalikan `{"message":"An unexpected error occurred","unrealmsg":"..."}`. Body JSON yang gagal di-bind biasanya menghasilkan `400 Invalid request body`, kecuali `PUT /kos/{id}` yang saat ini menghasilkan `404` untuk kondisi tersebut. `id` resource berupa UUID dan waktu ditulis sebagai timestamp. Respons daftar kos juga menyertakan metadata pagination (`page`, `per_page`, `total`, `last_page`).

## Contoh Response

Contoh berikut memakai nilai ilustratif. Field ber-pointer atau relasi bertag `omitempty` tidak muncul jika nilainya kosong/tidak dimuat. `GET /kos` mengisi `cover_photo` dengan URL foto paling awal dibuat atau URL fallback. Respons create/update kos dan detail kos tidak mengisi `cover_photo`.

### Login berhasil (`200`)

```json
{
  "message": "Login successful",
  "token": "<jwt>",
  "user": {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "name": "Pemilik Kos",
    "email": "pemilik@example.com",
    "created_at": "2026-10-30T08:00:00Z",
    "updated_at": "2026-10-30T08:00:00Z"
  }
}
```

### Daftar kos kosong (`GET /kos`, `200`)

```json
{"data": [], "meta": {"page": 1, "per_page": 10, "total": 0, "last_page": 0}}
```

### Resource kos (`POST /kos` mengembalikan `201`, `PUT /kos/{id}` mengembalikan `200`)

```json
{
  "data": {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "user_id": "123e4567-e89b-12d3-a456-426614174001",
    "name": "Kos Melati",
    "address": "Jl. Merdeka No. 10",
    "province": "Jawa Barat",
    "city": "Bandung",
    "district": "Coblong",
    "description": "Kos dekat kampus",
    "subdistrict": "Dago",
    "gmaps_link": "https://maps.example/kos-melati",
    "price_monthly": 1500000,
    "room_size": 20,
    "total_rooms": 10,
    "available_rooms": 2,
    "minimum_rent_months": 1,
    "type": "woman",
    "status": "published",
    "rules": "Tidak boleh merokok",
    "created_at": "2026-10-30T08:00:00Z",
    "updated_at": "2026-10-30T08:00:00Z"
  }
}
```

### Detail kos (`GET /kos/{id}`, `200`)

Detail kos menyertakan relasi pemilik, fasilitas, kontak, dan foto yang dimuat oleh endpoint.

```json
{
  "data": {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "user_id": "123e4567-e89b-12d3-a456-426614174001",
    "name": "Kos Melati",
    "address": "Jl. Merdeka No. 10",
    "province": "Jawa Barat",
    "city": "Bandung",
    "district": "Coblong",
    "description": "Kos dekat kampus",
    "subdistrict": "Dago",
    "gmaps_link": "https://maps.example/kos-melati",
    "price_monthly": 1500000,
    "room_size": 20,
    "total_rooms": 10,
    "available_rooms": 2,
    "minimum_rent_months": 1,
    "type": "woman",
    "status": "published",
    "rules": "Tidak boleh merokok",
    "created_at": "2026-10-30T08:00:00Z",
    "updated_at": "2026-10-30T08:00:00Z",
    "user": {
      "id": "123e4567-e89b-12d3-a456-426614174001",
      "name": "Pemilik Kos",
      "email": "pemilik@example.com",
      "role": "owner",
      "created_at": "2026-10-30T08:00:00Z",
      "updated_at": "2026-10-30T08:00:00Z"
    },
    "facilities": [
      {
        "id": "123e4567-e89b-12d3-a456-426614174002",
        "name": "Wi-Fi",
        "icon": "fa-solid fa-wifi",
        "created_at": "2026-10-30T08:00:00Z",
        "updated_at": "2026-10-30T08:00:00Z"
      }
    ],
    "contact": {
      "kos_id": "123e4567-e89b-12d3-a456-426614174000",
      "contact_name": "Penjaga Kos",
      "phone": "+62221234567",
      "whatsapp": "+6281234567890",
      "created_at": "2026-10-30T08:00:00Z",
      "updated_at": "2026-10-30T08:00:00Z"
    },
    "photos": [
      {
        "id": "123e4567-e89b-12d3-a456-426614174003",
        "kos_id": "123e4567-e89b-12d3-a456-426614174000",
        "image_url": "https://example.com/kos-melati/kamar.jpg",
        "public_id": "cariinkos/photos/kos-melati-kamar",
        "created_at": "2026-10-30T08:00:00Z",
        "updated_at": "2026-10-30T08:00:00Z"
      }
    ]
  }
}
```

### Fasilitas dibuat (`201`)

```json
{
  "data": {
    "id": "123e4567-e89b-12d3-a456-426614174002",
    "name": "Wi-Fi",
    "icon": "fa-solid fa-wifi",
    "created_at": "2026-10-30T08:00:00Z",
    "updated_at": "2026-10-30T08:00:00Z"
  }
}
```

### Kontak kos dibuat (`201`)

```json
{
  "data": {
    "kos_id": "123e4567-e89b-12d3-a456-426614174000",
    "contact_name": "Penjaga Kos",
    "phone": "+62221234567",
    "whatsapp": "+6281234567890",
    "created_at": "2026-10-30T08:00:00Z",
    "updated_at": "2026-10-30T08:00:00Z"
  }
}
```

### Foto kos diunggah (`201`)

```json
{
  "data": {
    "id": "123e4567-e89b-12d3-a456-426614174003",
    "kos_id": "123e4567-e89b-12d3-a456-426614174000",
    "image_url": "https://example.com/kos-melati/kamar.jpg",
    "created_at": "2026-10-30T08:00:00Z",
    "updated_at": "2026-10-30T08:00:00Z"
  }
}
```

### Fasilitas dikaitkan (`201`)

```json
{"message":"Facility attached to kos successfully"}
```

### Validasi gagal (`422`)

```json
{
  "message": "Validation failed",
  "errors": {
    "name": ["The name field is required."]
  }
}
```

### Kredensial login salah (`401`)

```json
{"message":"Invalid email or password"}
```

### Resource tidak ditemukan (`404`)

```json
{"message":"Data not found"}
```

## Ringkasan Endpoint

Semua method dan path di bawah relatif terhadap base URL.

| Modul | Method | Path | Keterangan | Auth |
|---|---|---|---|---|
| Health | GET | `/` | Cek API | Tidak |
| Auth | POST | `/login` | Login dan memperoleh token | Tidak |
| Auth | POST | `/logout` | Logout | Ya |
| Dokumentasi | GET | `/docs` | Halaman dokumentasi API | Tidak |
| Kos publik | GET | `/public/kos` | Daftar kos tanpa autentikasi | Tidak |
| Kos publik | GET | `/public/kos/{id}` | Detail kos tanpa autentikasi | Tidak |
| Kos | GET | `/kos` | Daftar kos | Ya |
| Kos | GET | `/kos/{id}` | Detail kos | Ya |
| Kos | POST | `/kos` | Buat kos | Ya |
| Kos | PUT | `/kos/{id}` | Perbarui kos | Ya |
| Kos | DELETE | `/kos/{id}` | Hapus kos | Ya |
| Fasilitas | GET | `/facilities` | Daftar fasilitas | Ya |
| Fasilitas | GET | `/facilities/{id}` | Detail fasilitas | Ya |
| Fasilitas | POST | `/facilities` | Buat fasilitas | Ya |
| Fasilitas | PUT | `/facilities/{id}` | Perbarui fasilitas | Ya |
| Fasilitas | DELETE | `/facilities/{id}` | Hapus fasilitas | Ya |
| Fasilitas kos | GET | `/kos/{id}/facilities` | Daftar fasilitas pada kos | Ya |
| Fasilitas kos | POST | `/kos/{id}/facilities` | Kaitkan fasilitas ke kos (`201`) | Ya |
| Fasilitas kos | DELETE | `/kos/{id}/facilities/{facility_id}` | Hapus fasilitas dari kos | Ya |
| Foto kos | GET | `/kos/{id}/photos` | Daftar foto kos | Ya |
| Foto kos | POST | `/kos/{id}/photos` | Unggah foto | Ya |
| Foto kos | PUT | `/kos/{id}/photos/{photo_id}` | Ganti foto | Ya |
| Foto kos | DELETE | `/kos/{id}/photos/{photo_id}` | Hapus foto | Ya |
| Kontak kos | GET | `/kos/{id}/contacts` | Lihat kontak kos | Ya |
| Kontak kos | POST | `/kos/{id}/contacts` | Simpan kontak kos | Ya |
| Kontak kos | PUT | `/kos/{id}/contacts` | Perbarui kontak kos | Ya |
| Kontak kos | DELETE | `/kos/{id}/contacts` | Hapus kontak kos | Ya |

## Auth

### `POST /login`

Request `application/json`:

```json
{
  "email": "pemilik@example.com",
  "password": "password"
}
```

`email` wajib berupa email valid dan `password` wajib berupa string. Kegagalan validasi mengembalikan `422`; kredensial salah mengembalikan `401` dengan `{"message":"Invalid email or password"}`.

Respons sukses:

```json
{
  "message": "Login successful",
  "token": "<jwt>",
  "user": {
    "id": "123e4567-e89b-12d3-a456-426614174001",
    "name": "Pemilik Kos",
    "email": "pemilik@example.com",
    "created_at": "2026-10-30T08:00:00Z",
    "updated_at": "2026-10-30T08:00:00Z"
  }
}
```

Kredensial salah mengembalikan `401`.

### `POST /logout`

Memerlukan Bearer token. Respons sukses: `{"message":"Logout successful"}`.

## Kos

### Endpoint publik

Dua route publik berikut tidak memerlukan Bearer token. Route ini menggunakan handler daftar dan detail kos yang sama dengan endpoint terautentikasi, sehingga saat ini tidak membatasi hasil berdasarkan status publikasi:

- `GET /public/kos` menerima query `page` dan `per_page` dengan default, batas, dan metadata pagination yang sama seperti `GET /kos`.
- `GET /public/kos/{id}` mengembalikan detail kos beserta relasi yang dimuat oleh endpoint, sama seperti `GET /kos/{id}`. Resource yang tidak ditemukan mengembalikan `404`.

### Endpoint terautentikasi

### Field kos

Request create dan update menerima JSON berikut. Field yang memiliki tanda `*` wajib saat create dan update; field lainnya opsional atau bernilai default nol sesuai tipe request.

| Field | Tipe | Aturan |
|---|---|---|
| `name`* | string | Wajib, maksimum 255 karakter |
| `description` | string/null | Opsional |
| `address`* | string | Wajib |
| `province`* | string | Wajib, maksimum 255 karakter |
| `city`* | string | Wajib, maksimum 255 karakter |
| `district`* | string | Wajib, maksimum 255 karakter |
| `subdistrict` | string/null | Opsional |
| `gmaps_link` | string/null | Opsional |
| `price_monthly`* | integer | Wajib, minimal 1 |
| `type` | string | Opsional saat create; jika dikirim, hanya `mixed`, `man`, atau `woman` yang diterima. Wajib saat update dan nilainya harus salah satu dari ketiga value tersebut. |
| `status` | string | Opsional saat create; jika dikirim, hanya `draft`, `published`, atau `archived` yang diterima. Saat update wajib dan hanya `draft` atau `published` yang diterima (`archived` ditolak). |
| `room_size` | integer/null | Opsional |
| `total_rooms` | integer | Opsional |
| `available_rooms` | integer | Opsional |
| `minimum_rent_months` | integer | Opsional |
| `rules` | string/null | Opsional |

Contoh body create/update:

```json
{
  "name": "Kos Melati",
  "description": "Kos dekat kampus",
  "address": "Jl. Merdeka No. 10",
  "province": "Jawa Barat",
  "city": "Bandung",
  "district": "Coblong",
  "subdistrict": "Dago",
  "gmaps_link": "https://maps.example/kos-melati",
  "price_monthly": 1500000,
  "type": "woman",
  "status": "published",
  "room_size": 20,
  "total_rooms": 10,
  "available_rooms": 2,
  "minimum_rent_months": 1,
  "rules": "Tidak boleh merokok"
}
```

`POST /kos` mengembalikan `201` dengan resource di `data`; daftar mengembalikan `data` dan `meta`, sedangkan detail mengembalikan `data`. `GET /kos` menerima `page` (default `1`) dan `per_page` (default `10`, maksimum `100`); nilai tidak valid atau kurang dari `1` kembali ke default. Update memakai `PUT` dan mengembalikan resource di `data`. Delete mengembalikan pesan sukses. Pada create, hanya field nama, alamat, provinsi, kota, distrik, dan harga yang diwajibkan; `type` dan `status` opsional, tetapi bila diisi harus memakai nilai enum yang didukung dan value lain ditolak. Pada update, `type` dan `status` wajib; `status` update hanya menerima `draft` atau `published`, meskipun enum database juga memiliki `archived`. `cover_photo` hanya diisi pada daftar kos: URL foto paling awal dibuat, atau `{http.url}/public/images_not_found.jpg` bila belum ada foto.

## Fasilitas

### `POST /facilities` dan `PUT /facilities/{id}`

```json
{
  "name": "Wi-Fi",
  "icon": "fa-solid fa-wifi"
}
```

`name` wajib berupa string maksimum 255 karakter. `icon` boleh string/null maksimum 255 karakter; untuk tampilan Font Awesome kirim nama kelas, misalnya `fa-solid fa-wifi`. API menyimpan dan mengembalikan string tersebut tanpa memvalidasi kelas icon. Create mengembalikan `201`; operasi lainnya mengembalikan resource di `data`, sedangkan delete mengembalikan `message`.

## Fasilitas Kos

### `POST /kos/{id}/facilities`

Mengaitkan fasilitas yang telah dibuat dengan kos. Body JSON:

```json
{
  "facility_id": "<uuid-fasilitas>"
}
```

Sukses mengembalikan status `201` dan `{"message":"Facility attached to kos successfully"}`. `facility_id` diwajibkan sebagai string (bukan validasi format UUID); ID yang tidak ditemukan dapat mengembalikan `404`. `DELETE /kos/{id}/facilities/{facility_id}` melepas relasi dan mengembalikan `{"message":"Facility detached from kos successfully"}`.

## Kontak Kos

### `POST /kos/{id}/contacts` dan `PUT /kos/{id}/contacts`

```json
{
  "contact_name": "Penjaga Kos",
  "phone": "+62221234567",
  "whatsapp": "+6281234567890"
}
```

Satu kos hanya memiliki satu kontak, yang ditentukan oleh `kos_id` dan tidak memiliki ID tersendiri. `contact_name` dan `phone` wajib berupa string; `whatsapp` opsional/null maksimum 50 karakter. Create mengembalikan `201`; get, create, dan update mengembalikan satu resource dalam `data`; delete mengembalikan `message`.

## Foto Kos

### `POST /kos/{id}/photos` dan `PUT /kos/{id}/photos/{photo_id}`

Kirim sebagai `multipart/form-data` dengan field file bernama `photo`. File harus berupa JPEG, PNG, atau WebP, berukuran lebih dari 0 byte dan maksimal 10 MB.

Create mengembalikan `201` dan resource foto di `data`. Update mengganti foto lalu mengembalikan resource. Resource foto menyediakan `image_url` dan `public_id` bila tersedia; delete mengembalikan `message`.

## Contoh Penggunaan

Login:

```bash
curl -X POST http://localhost:3000/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{"email":"pemilik@example.com","password":"password"}'
```

Buat kos (ganti `<token>` dengan JWT hasil login):

```bash
curl -X POST http://localhost:3000/api/v1/kos \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"name":"Kos Melati","address":"Jl. Merdeka No. 10","province":"Jawa Barat","city":"Bandung","district":"Coblong","price_monthly":1500000,"type":"woman","status":"published"}'
```

Unggah foto:

```bash
curl -X POST http://localhost:3000/api/v1/kos/<kos-id>/photos \
  -H "Authorization: Bearer <token>" \
  -F "photo=@./kamar.jpg"
```

## Catatan Cakupan

Dokumen ini mencakup route HTTP yang terdaftar di `routes/api.go` dan `routes/api/`. Modul internal Goravel seperti queue, cache, mail, dan database bukan endpoint HTTP aplikasi, sehingga tidak dicantumkan sebagai API.
