# Belajar-Golang-Rastful-API

Link Pembelajaran <a href="https://www.youtube.com/watch?v=bJ2ZFt9D0uI&list=PL-CtdCApEFH-0i9dzMzLw6FKVrFWv3QvQ&index=14&ab_channel=ProgrammerZamanNow">Programmer Zaman Now</a>

Link Documentasi

- <a href="https://github.com/go-sql-driver/mysql">Driver MySQL</a>
- <a href="https://github.com/julienschmidt/httprouter">HTTP Router</a>
- <a href="https://github.com/go-playground/validator">Validation</a>
- <a href="https://github.com/stretchr/testify">Testify</a>

> Setup

Driver MySQL

```golang
go get -u github.com/go-sql-driver/mysql
```

HTTP Router

```golang
go get github.com/julienschmidt/httprouter
```

Validator

```golang
go get github.com/go-playground/validator
```

Testing

```golang
go get github.com/stretchr/testify
```

> Create Open Api Spec

- `apispec.json`
  - like `sweger` atau dokumentasi api yang akan di buat

> Create Database

Create Database `belajar_golang_restful_api`

```sql
create database belajar_golang_restful_api;
&
create database belajar_golang_restful_api_test;
```

Create Table `category`

```sql
CREATE TABLE category (
	id INTEGER PRIMARY KEY auto_increment,
	name VARCHAR(200) NOT NULL
)
```

> Step By Step

1. Create `Repository` Folder

   - `category_repository.go`
   - `category_repository_impl.go`
     - Untuk `query` ke `Database`

2. Create `Service` Folder

   - `category_service.go`
   - `category_service_impl.go`
     - untuk `logic` dan pengolahan data yang di dapat dari `Databasae`

3. Create `Controller` Folder

   - `category_controller.go`
   - `category_controller_impl.go`
     - untuk mengambil data yang telah di olah di `Sevice` dan dan di return sebagai `response` API

4. `app` Folder -> config

   - `database.go`
     - setup config database
   - `router.go`
     - setup router dan kumpulan dari semua `endpoint` API

5. `exception` Folder -> handler `error`

   - `error_handler.go`
     - handler err speerti:
       - `notFoundErr`
       - `validator`
       - `internal server err`
   - `error_not_found.go`
     - karana error notfound tidak err yang di dapat dari database dan tidak ada triker spesifik maka butuh di buat `struct` agar bisa di persifikasi dan balikan dari kondisi `notfound`
     - sedangakan `validate` sudah di handle oleh `validator`
     - dan `internal server err` pun sama karna server yang tidak dapat response atau mati

6. `Helper` Folder -> function yang `reuseble`

   - `error.go`
     - handle err balikan di varible
   - `tx.go`
     - handle `commit` dan `rollback` dari `tx`
   - `json.go`
     - handle `read` data `body` dan di `decode` agar bisa di olah
     - handle `write` nge `encode` data yang telah di olah untuk menjadi `response`
   - `model.go`
     - `ToCategoryRes`
       - template response di properti `data` untuk singgle data
     - `ToCategoryResult`
       - template response di properti `data` untuk data yang banyak, seperti untuk `findAll`

7. `model` Folder -> jika di TS seperti `entity`

   - `domain`
     - model dari `category` dan diperuntukan supaya kita bisa nge `filter` data yang mau di `show` atau di `hide` data dari `database`
   - `web`
     - `category_create_req.go` ->
       - `entity` atau `model` dari data `body` yang di kirim dari Api `create category`
     - `category_res.go`
       - model untuk response api
     - `category_update_req.go`
       - model untuk response api update
     - `web_res.go`
       - struktur date response api

8. `middleware` Folder

   - `auth_middleware.go`
     - handle `AuthMiddleware` api

9. `test` Folder
   - `category_controller_test.go`
     - test by endpoint
