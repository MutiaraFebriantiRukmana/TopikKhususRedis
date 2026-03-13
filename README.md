# Go Redis Clean Architecture
Tugas Praktikum Redis menggunakan Golang dengan standar Clean Architecture.
# Prom AI [menggunakan extention blackbox AI di vscode]
Please generate a complete Golang backend project structure for a Redis-based application using Clean Architecture.
I need you to provide code for these specific files:
domain/user.go: (User struct and UserRepository interface)
repository/redis/user_repository.go: (Implementation using go-redis/v9)
usecase/user_usecase.go: (Business logic)
usecase/user_usecase_test.go: (Unit test with manual mock using 'testify/assert')
main.go: (The entry point to connect everything)
Use professional, clean code and provide each file in a separate code block so I can easily apply them to my workspace.

# download go sdk dan aplikasi redis

## Cara Menjalankan
1. Pastikan Redis Server jalan di port 6379.
2. `go mod tidy`
3. `go run main.go`
