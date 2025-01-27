# api_sanbox

This is a API sanbox for testing simple API calls without using toools like postman
To run in linux
```
go mod tiny
go build -o bin/api_sandbox
./bin/api_sanbox
```
for windows
```
go mod tidy
go run .
```

It supports these requests
```
curl -X POST http://localhost:1651/persons -H "Content-Type: application/json" -d '{
    "name": "John Doe",
    "language": "English",
    "id": "JD123456",
    "bio": "A brief bio about John Doe.",
    "version": 1.0
}'

curl http://localhost:1651/persons
curl http://localhost:1651/fetchname?name=Adeel_Solangi
curl http://localhost:1651/fetchid/VPK9MQRKX2L847HQ
```

It uses gin & net/http for basic backend utils.
