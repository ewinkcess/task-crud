# Task CRUD API (Go + Supabase)

🚀 **Task CRUD** adalah RESTful API sederhana menggunakan **Golang** dengan arsitektur **Clean Architecture (Handler → Service → Repository)** dan database **PostgreSQL (Supabase)**.

Project ini cocok untuk:

* Belajar backend Golang
* Memahami relasi database (Foreign Key)
* Implementasi CRUD dengan struktur profesional

---

## ✨ Fitur

* CRUD **Products**
* CRUD **Categories**
* Relasi **Product ↔ Category** (Foreign Key)
* Response JSON dengan nested object (category di dalam product)
* Struktur kode rapi & scalable

---

## 🏗️ Arsitektur Project

```
.
├── handlers        # HTTP Handler (Controller)
├── services        # Business Logic
├── repositories    # Database Query
├── models          # Struct Model
├── config          # Config & DB
├── main.go
└── README.md
```

---

## 🛠️ Tech Stack

* **Golang**
* **PostgreSQL**
* **Supabase**
* **net/http**
* **database/sql**

---

## 📦 Database Schema

### categories

```sql
id SERIAL PRIMARY KEY
name TEXT
description TEXT
```

### products

```sql
id SERIAL PRIMARY KEY
name TEXT
price INT
stock INT
category_id INT REFERENCES categories(id)
```

---

## 🔗 Endpoint List

### Category

| Method | Endpoint             | Description        |
| ------ | -------------------- | ------------------ |
| GET    | /api/categories      | Get all categories |
| GET    | /api/categories/{id} | Get category by ID |
| POST   | /api/categories      | Create category    |
| PUT    | /api/categories/{id} | Update category    |
| DELETE | /api/categories/{id} | Delete category    |

### Product

| Method | Endpoint           | Description                       |
| ------ | ------------------ | --------------------------------- |
| GET    | /api/products      | Get all products (with category)  |
| GET    | /api/products/{id} | Get product by ID (with category) |
| POST   | /api/products      | Create product                    |
| PUT    | /api/products/{id} | Update product                    |
| DELETE | /api/products/{id} | Delete product                    |

---

## 📥 Contoh Response Product

```json
{
  "id": 29,
  "name": "Mesin Cuci",
  "price": 4500000,
  "stock": 10,
  "category_id": 2,
  "category": {
    "id": 2,
    "name": "Home Appliances",
    "description": "Peralatan rumah tangga"
  }
}
```

---

## ⚙️ Setup & Installation

### 1️⃣ Clone Repository

```bash
git clone https://github.com/ewinkcess/task-crud.git
cd task-crud
```

### 2️⃣ Setup Environment

Buat file `.env`

```env
DB_HOST=...
DB_PORT=5432
DB_USER=...
DB_PASSWORD=...
DB_NAME=...
```

### 3️⃣ Run Application

```bash
go run main.go
```

Server berjalan di:

```
http://localhost:8080
```

---

## 🧠 Catatan Penting

* Pastikan `category_id` **sudah ada** sebelum create product
* Relasi foreign key akan error jika category tidak valid
* Gunakan `JOIN` untuk menampilkan data category

---

## 📌 Roadmap

* [ ] Pagination
* [ ] Authentication (JWT)
* [ ] Validation
* [ ] Swagger / OpenAPI

---

## 🤝 Kontribusi

Pull request sangat terbuka 🙌

---

## 👨‍💻 Author/Mentor

**Muhammad Zuhrul Umam**
🔗 GitHub: [https://github.com/zuhrulumam](https://github.com/zuhrulumam)

---

⭐ Jika repo ini membantu, jangan lupa beri **star** ya!
