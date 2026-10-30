# CariinKos API Docs

Dokumentasi REST API backend CariinKos. API menggunakan JSON kecuali endpoint unggah foto dan seluruh path API diawali `/api/v1`.

## Base URL

```text
http://localhost:3000/api/v1
```

Sesuaikan host dan port dengan konfigurasi deployment. Endpoint `GET /` pada base path mengembalikan `{"halo":"api"}`.

## Autentikasi

Endpoint login bersifat publik. Endpoint lain di modul kos dan fasilitas memerlukan JWT dengan header:

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
{"data": {}}
```

Operasi yang tidak mengembalikan resource memberikan `message`. Validasi gagal menggunakan status `422` dan bentuk:

```json
{
  "message": "Validation failed",
  "errors": {}
}
```

Error umum lainnya menggunakan `message`; format detail dapat berbeda menurut endpoint. `id` resource berupa UUID dan waktu ditulis sebagai timestamp.

## Ringkasan Endpoint

Semua method dan path di bawah relatif terhadap base URL.

| Modul | Method | Path | Keterangan | Auth |
|---|---|---|---|---|
| Health | GET | `/` | Cek API | Tidak |
| Auth | POST | `/login` | Login dan memperoleh token | Tidak |
| Auth | POST | `/logout` | Logout | Ya |
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
| Fasilitas kos | POST | `/kos/{id}/facilities` | Kaitkan fasilitas ke kos | Ya |
| Fasilitas kos | DELETE | `/kos/{id}/facilities/{facility_id}` | Hapus fasilitas dari kos | Ya |
| Foto kos | GET | `/kos/{id}/photos` | Daftar foto kos | Ya |
| Foto kos | POST | `/kos/{id}/photos` | Unggah foto | Ya |
| Foto kos | PUT | `/kos/{id}/photos/{photo_id}` | Ganti foto | Ya |
| Foto kos | DELETE | `/kos/{id}/photos/{photo_id}` | Hapus foto | Ya |
| Kontak kos | GET | `/kos/{id}/contacts` | Daftar kontak kos | Ya |
| Kontak kos | POST | `/kos/{id}/contacts` | Tambah kontak | Ya |
| Kontak kos | PUT | `/kos/{id}/contacts/{contact_id}` | Perbarui kontak | Ya |
| Kontak kos | DELETE | `/kos/{id}/contacts/{contact_id}` | Hapus kontak | Ya |

## Auth

### `POST /login`

Request `application/json`:

```json
{
  "email": "pemilik@example.com",
  "password": "password"
}
```

`email` wajib berupa email valid dan `password` wajib berupa string.

Respons sukses:

```json
{
  "message": "Login successful",
  "token": "<jwt>",
  "user": {}
}
```

Kredensial salah mengembalikan `401`.

### `POST /logout`

Memerlukan Bearer token. Respons sukses: `{"message":"Logout successful"}`.

## Kos

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
| `type`* | string | `mixed`, `man`, atau `woman` |
| `status`* | string | `draft`, `published`, atau `archived` |
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

`POST /kos` mengembalikan `201` dengan resource di `data`; daftar dan detail mengembalikan `data`. Update memakai `PUT` dan mengembalikan resource di `data`. Delete mengembalikan pesan sukses.

## Fasilitas

### `POST /facilities` dan `PUT /facilities/{id}`

```json
{
  "name": "Wi-Fi",
  "icon": "wifi"
}
```

`name` wajib berupa string maksimum 255 karakter. `icon` boleh string/null maksimum 255 karakter. Create mengembalikan `201`; operasi lainnya mengembalikan resource di `data`, sedangkan delete mengembalikan `message`.

## Fasilitas Kos

### `POST /kos/{id}/facilities`

Mengaitkan fasilitas yang telah dibuat dengan kos. Body JSON:

```json
{
  "facility_id": "<uuid-fasilitas>"
}
```

Sukses mengembalikan `{"message":"Facility attached to kos successfully"}`. `DELETE /kos/{id}/facilities/{facility_id}` melepas relasi dan mengembalikan pesan sukses.

## Kontak Kos

### `POST /kos/{id}/contacts` dan `PUT /kos/{id}/contacts/{contact_id}`

```json
{
  "phone": "+62221234567",
  "whatsapp": "+6281234567890"
}
```

`phone` wajib berupa string maksimum 50 karakter. `whatsapp` opsional/null maksimum 50 karakter. Create mengembalikan `201`; list, create, dan update mengembalikan resource dalam `data`; delete mengembalikan `message`.

## Foto Kos

### `POST /kos/{id}/photos` dan `PUT /kos/{id}/photos/{photo_id}`

Kirim sebagai `multipart/form-data` dengan field file bernama `photo`. File harus berupa JPEG, PNG, atau WebP, berukuran lebih dari 0 byte dan maksimal 10 MB.

Create mengembalikan `201` dan resource foto di `data`. Update mengganti foto lalu mengembalikan resource. Resource foto menyediakan `image_url`; delete mengembalikan `message`.

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
