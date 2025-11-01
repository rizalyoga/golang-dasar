# Langkah awal pembuatan program :

1. Instal gin dengan:

```
# go get -u github.com/gin-gonic/gin
```

2. Intall Air, fitur live reload yg berguna untuk reload aplikasi ketika ada perubahan:

```
# go install github.com/air-verse/air@latest

// pada folder project jalankan command:
# air init

// Seteleh install air, maka untuk menjalankan program cukup dengan command:
# air
```

3. Buat file env, pada go
   untuk menambahkan file env bisa menggunakan library dari luar yaitu godotenv:

```
# go get github.com/joho/godotenv
```

4. Install gorm, gorm adalah sebuah library ORM yang bisa membuat interaksi antara app dan database menjadi lebih sederhana.
   dengan ORM ini kita tidak perlu lagi menulis sql code.

```
// Install gorm
# go get -u gorm.io/gorm

// Install driver postgres
# go get -u gorm.io/driver/postgres
// jika menggunakan DB lain kamu bisa menggantinya menjadi '/driver/mysql' atau '/driver/sqlite'
```

5. Install JWT

```
# go get -u github.com/golang-jwt/jwt/v5
```

5. Install validator

- validator digunakan untuk mengani error dari sebuah input.

```
# go get github.com/go-playground/validator/v10
```

6. Install cors

```
# go get github.com/gin-contrib/cors
```
