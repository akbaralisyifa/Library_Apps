# **Sistem Manajemen Perpustakaan**

## **Deskripsi**

**Sistem Manajemen Perpustakaan** adalah proyek yang dirancang untuk mengelola buku, kategori buku, dan pengguna dengan dua peran utama:
- **Admin**: Memiliki akses penuh untuk mengelola buku, kategori, dan rekomendasi.
- **User**: Dapat melihat buku, mencari buku berdasarkan 'title' dan mengelola rekomendasi buku.

## **Fitur Utama (MVP)**

1. **Manajemen Buku**
   - **Admin**: Bisa melakukan GET, POST, UPDATE, DELETE pada buku.
   - **User**: Bisa melakukan GET (hanya akses baca).

2. **Manajemen Kategori Buku**
   - **Admin**: Bisa melakukan GET, POST, UPDATE, DELETE pada kategori buku.
   - **User**: Bisa melakukan GET (hanya akses baca).

3. **Manajemen Pengguna**
   - Admin dan User bisa melakukan login dengan endpoint login yang berbeda.

4. **Fitur Rekomendasi**
   - **Admin**: Memiliki akses penuh untuk mengelola rekomendasi (GET, POST, UPDATE, DELETE).
   - **User**: Dapat mengelola rekomendasi buku pribadi (GET, POST, UPDATE, DELETE).

     
## **API Endpoints**

### **Authentication**
- **Register**:  
  *Admin : `POST /register/admin`
  *User : `POST /register`
  **Request body**:
  ```json
  {
    "username": "admin",
    "email": "admin@example.com",
    "password": "password123"
  }   

- **Login**:  
  `POST /login`  
  **Request body**:
  ```json
  {
    "email": "admin@example.com",
    "password": "password123"
  }

- **Books**
  *Admin : `POST /books`, `GET /books`, `PUT /books/:book_id`, `DELETE /books/:book_id`
  *User : `GET /books`
    **Request body**:
  ```json
  {
    "category_id": 1,
    "title": "days",
    "author": "bambang",
    "published_year": "2023"
  }

- **Category**
  *Admin : `POST /category`, `GET /category`, `PUT /category/:category_id`, `DELETE /category/:category_id`
    **Request body**:
  ```json
  {
    "name":"horor"
  }

- **Recomendation**
  *Admin & User : `POST /recomendation/book_id`, `GET /recomendation`, `PUT /recomendation/:recomendation_id`, `DELETE /recomendation/:recomendation_id`
    **Request body**:
  ```json
  {
    "reason":"this book is perfect"
  }

### **ERD**
