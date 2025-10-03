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

go get github.com/rakyll/hey/requester

go run main.go

📦 fruits-api
├── 📂 cmd/
│   └── serve.go
│
├── 📂 config/
│   └── config.go
│
├── 📂 domain/
│   ├── user.go
│   └── product.go
│
├── 📂 product/
│   ├── port.go
│   └── service.go
│
├── 📂 infra/
│   └── 📂 db/
│       ├── connection.go
│       └── migrate.go
│
├── 📂 migrations/
│   ├── 000001-create-users.up.sql
│   ├── 000001-create-users.down.sql
│   ├── 000002-create-fruits.up.sql
│   └── 000002-create-fruits.down.sql
│
├── 📂 repo/
│   ├── product.go
│   └── user.go
│
├── 📂 rest/
│   └── 📂 handlers/
│       ├── 📂 product/
│       │   ├── create.go
│       │   ├── get.go
│       │   ├── update.go
│       │   ├── delete.go
│       │   ├── routes.go
│       │   ├── handler.go
│       │   └── port.go
│       │
│       ├── 📂 user/
│       │   ├── create.go
│       │   ├── get.go
│       │   ├── routes.go
│       │   ├── handler.go
│       │   └── port.go
│       │
│       └── 📂 middleware/
│           ├── jwt.go
│           ├── cors.go
│           ├── logger.go
│           └── manager.go
│
├── server.go
├── 📂 utils/
│   ├── send.go
│   └── jwt.go
│
├── .env
├── .gitignore
├── go.mod
├── go.sum
├── main.go
└── README.md

```
