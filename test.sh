#!/bin/bash

# 1. Pastikan semua dependensi terunduh
echo "--- Running go mod tidy ---"
go mod tidy

# 2. Jalankan migrasi database
echo "--- Running database migration ---"
go run main.go migrate

# 3. Jalankan seeding (opsional, jika Anda punya fungsi seed)
echo "--- Running database seed ---"
go run main.go seed

# 4. Jalankan Server
echo "--- Starting the Server ---"
go run main.go

#!/bin/bash

# Konfigurasi
URL="http://localhost:8080/api/todos"
ADMIN_EMAIL="admin@unityunsrat.dev"
ADMIN_PASS="miniprojectbackendunity2026"

echo "--- 1. Login sebagai Admin ---"
TOKEN=$(curl -s -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d "{\"email\":\"$ADMIN_EMAIL\", \"password\":\"$ADMIN_PASS\"}" | jq -r .token)

echo "--- 2. Testing User Features (Scope) ---"
# Create
ID=$(curl -s -X POST $URL -H "Authorization: Bearer $TOKEN" -d '{"title":"Tugas Ezra", "description":"Coba fitur"}' | jq -r .id)
echo "Created: $ID"

# Get All Active
curl -s -X GET "$URL?limit=5" -H "Authorization: Bearer $TOKEN" | jq .

# Update
curl -s -X PUT "$URL/$ID" -H "Authorization: Bearer $TOKEN" -d '{"title":"Tugas Update", "is_done":true}' | jq .

# Soft Delete
curl -s -X DELETE "$URL/$ID" -H "Authorization: Bearer $TOKEN" | jq .

# Get Trash
curl -s -X GET "$URL/trash" -H "Authorization: Bearer $TOKEN" | jq .

# Restore
curl -s -X PATCH "$URL/$ID/restore" -H "Authorization: Bearer $TOKEN" | jq .

# Permanent Delete
curl -s -X DELETE "$URL/$ID/permanent" -H "Authorization: Bearer $TOKEN" | jq .

echo "--- 3. Testing Admin Features (Global) ---"
# Admin Get Active
curl -s -X GET "$URL/admin" -H "Authorization: Bearer $TOKEN" | jq .

# Admin Get Trash
curl -s -X GET "$URL/admin/trash" -H "Authorization: Bearer $TOKEN" | jq .

# Admin Get Specific ID
curl -s -X GET "$URL/admin/$ID" -H "Authorization: Bearer $TOKEN" | jq .

echo "Testing Selesai!"
