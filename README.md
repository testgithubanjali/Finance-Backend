# Finance Backend

## Overview
This is a role-based finance management backend built using Go, Gin, and PostgreSQL.

## Features
- JWT Authentication
- Role-based access control (RBAC)
- Financial records CRUD
- Dashboard analytics (income, expense, net)

## Tech Stack
- Go (Gin)
- PostgreSQL
- Docker

## Run Project
docker compose up --build

## APIs

### Auth
POST /signup  
POST /login  

### Records
GET /records  
POST /records  
PUT /records/:id  
DELETE /records/:id  

### Dashboard
GET /dashboard  

### User
PATCH /users/:id/role  