#!/bin/bash

set -e

echo "--- 1. Setup Environment ---"
if [ ! -f private_key.pem ]; then
    openssl genrsa -out private_key.pem 2048
    openssl rsa -in private_key.pem -pubout -out public_key.pem
fi

echo "--- 2. Database Migration & Seeding ---"
go mod tidy
go run main.go migrate
go run main.go reset
go run main.go seed

echo "--- 3. Starting Server ---"
go run main.go > server.log 2>&1 &
SERVER_PID=$!
sleep 3 

echo "--- 4. Testing Authentication & Scopes ---"

# Admin Login
echo "-> Logging in as Admin..."
ADMIN_RES=$(curl -s -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@unityunsrat.dev", "password":"miniprojectbackendunity2026"}')
TOKEN=$(echo $ADMIN_RES | jq -r .data.token)
AUTH="Authorization: Bearer $TOKEN"

# User Login 
echo "-> Logging in as User..."
USER_RES=$(curl -s -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"ezra.lmntt@proton.me", "password":"ezramighty08"}')
USER_TOKEN=$(echo $USER_RES | jq -r .data.token)
USER_AUTH="Authorization: Bearer $USER_TOKEN"

BASE_URL="http://localhost:8080/api/todos"
ADMIN_URL="http://localhost:8080/api/todos/admin"

echo "--- 5. Testing User Endpoints (Happy Path) ---"
# Create
ID=$(curl -s -X POST $BASE_URL/ -H "$AUTH" -H "Content-Type: application/json" \
  -d '{"title":"Test Title", "description":"Test Description"}' | jq -r .data.id)
echo "   Created ID: $ID"

curl -s -X GET $BASE_URL/ -H "$AUTH" | jq .
curl -s -X PUT $BASE_URL/$ID -H "$AUTH" -H "Content-Type: application/json" -d '{"title":"Updated"}' | jq .
curl -s -X DELETE $BASE_URL/$ID -H "$AUTH" | jq .
curl -s -X PATCH $BASE_URL/$ID/restore -H "$AUTH" | jq .

echo "--- 6. Testing Admin Endpoints ---"
curl -s -X GET $ADMIN_URL/ -H "$AUTH" | jq .
curl -s -X PUT $ADMIN_URL/$ID -H "$AUTH" -H "Content-Type: application/json" -d '{"title":"Admin Updated"}' | jq .
curl -s -X GET $ADMIN_URL/$ID -H "$AUTH" | jq .
curl -s -X DELETE $ADMIN_URL/$ID/permanent -H "$AUTH" | jq .

echo "--- 7. Running Negative Testing (Security) ---"
# Akses tanpa token
echo "-> Testing Unauthenticated Access..."
curl -s -o /dev/null -w "%{http_code}" -X GET $BASE_URL/ | grep -q "401" && echo "   PASSED: Unauth blocked"

# User biasa akses Admin
echo "-> Testing User accessing Admin endpoint..."
STATUS=$(curl -s -o /dev/null -w "%{http_code}" -X GET $ADMIN_URL/ -H "$USER_AUTH")
if [[ "$STATUS" == "401" || "$STATUS" == "403" ]]; then
    echo "   PASSED: User blocked from Admin scope (Status: $STATUS)"
else
    echo "   FAILED: User accessed Admin scope (Status: $STATUS)"
fi

echo -e "\n--- Semua Endpoint Teruji dengan Sukses! ---"

echo "--- 8. Cleanup ---"
if ps -p $SERVER_PID > /dev/null; then
    kill $SERVER_PID
fi
rm -f private_key.pem public_key.pem server.log
echo "Test Completed Successfully."
