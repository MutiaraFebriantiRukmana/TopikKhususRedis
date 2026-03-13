package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9" 
	
	"redis/domain"
	
	repoRedis "redis/repository/redis" 
	
	"redis/usecase"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", 
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("Gagal konek Redis: %v. Pastikan aplikasi Redis Server sudah kamu jalankan!", err)
	}
	fmt.Println("1. Koneksi ke Redis: BERHASIL!")

	userRepo := repoRedis.NewUserRepository(rdb)
	userUsecase := usecase.NewUserUsecase(userRepo)

	newUser := &domain.User{
		ID:    "101",
		Name:  "Daffa Praktikum",
		Email: "daffa@kampus.ac.id",
	}

	err := userUsecase.CreateUser(ctx, newUser)
	if err != nil {
		log.Fatalf("Gagal simpan user: %v", err)
	}
	fmt.Println("2. Simpan User: BERHASIL!")

	retrievedUser, err := userUsecase.GetUser(ctx, "101")
	if err != nil {
		log.Fatalf("Gagal ambil user: %v", err)
	}
	fmt.Printf("3. Ambil User: BERHASIL! (Data: %+v)\n", retrievedUser)
}