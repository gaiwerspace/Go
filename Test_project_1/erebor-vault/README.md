# Erebor Vault System

🏔️ Complete treasure vault management system

## Quick Start

### 1. Database
```bash
createdb erebor
```

### 2. Backend (Terminal 1)
```bash
cd backend
go mod tidy
go run github.com/99designs/gqlgen generate
export DATABASE_URL="postgres://postgres@localhost:5432/erebor?sslmode=disable"
go run cmd/server/main.go
```

### 3. Frontend (Terminal 2)
```bash
cd frontend
npm install
npm run codegen
npm run dev
```

Visit: http://localhost:5173

## Features
✅ Create vaults | ✅ Transfer gold | ✅ View history | ✅ ACID transactions

## Pre-seeded Keepers
1. Thorin Oakenshield | 2. Balin Son of Fundin | 3. Dáin Ironfoot | 4. Gimli Son of Glóin
