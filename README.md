# 📘 Fruits API (Go + PostgreSQL)

এই প্রোজেক্টে আমরা **Go (Golang)** ব্যবহার করেছি **PostgreSQL database** এর সাথে।  
Database connection, environment variables, migration সবকিছু setup করা আছে।

---

## 🚀 Requirements

- Go 1.20+
- PostgreSQL
- Git

---

## 📦 Dependencies

- **[github.com/jmoiron/sqlx](https://github.com/jmoiron/sqlx)**  
  → সহজ database access এর জন্য

- **[github.com/joho/godotenv](https://github.com/joho/godotenv)**  
  → `.env` ফাইল থেকে environment variables লোড করার জন্য

- **[github.com/lib/pq](https://github.com/lib/pq)**  
  → PostgreSQL driver

- **[github.com/golang-migrate/migrate](https://github.com/golang-migrate/migrate)**  
  → database migrations এর জন্য

- **[github.com/rubenv/sql-migrate](https://github.com/rubenv/sql-migrate)** _(optional)_  
  → alternative migration tool

---

## ⚙️ Installation

প্রথমে module initialize করো:

```bash
go mod init fruits-api

go get github.com/jmoiron/sqlx
go get github.com/joho/godotenv
go get github.com/lib/pq
go get github.com/golang-migrate/migrate/v4
# Alternative (optional)
go get github.com/rubenv/sql-migrate/...

go run main.go

domain/
   fruits.go
   user.go
   order.go
   favorite.go
order/
   service.go
   port.go
favorite/
   service.go
   port.go
repo/
   fruits.go
   user.go
   order.go
   favorite.go
rest/
   handlers/
      product/
      user/
      order/
      favorite/

```
