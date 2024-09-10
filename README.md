# **Library Management System**

## **Overview**

The **Library Management System** is a project designed to manage books, categories, and users with two user roles:
- **Admin**: Full access to manage books, categories, and recommendations.
- **User**: Can view books and manage personal book recommendations.

## **MVP Features**

1. **Book Management**
   - **Admin**: Can perform GET, POST, UPDATE, DELETE on books.
   - **User**: Can perform GET (read-only access).

2. **Category Management**
   - **Admin**: Can perform GET, POST, UPDATE, DELETE on categories.
   - **User**: Can perform GET (read-only access).

3. **User Management**
   - Admin and Users can log in with separate login endpoints.

4. **Recommendation Feature**
   - **Admin**: Full management access to recommendations (GET, POST, UPDATE, DELETE).
   - **User**: Manage personal recommendations (GET, POST, UPDATE, DELETE).

## **API Endpoints**

### **Authentication**
- **Register**:  
  *Admin : `POST /register/admin`
  *User : `POST /register`
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
