package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis() *redis.Client {
	// get redis env
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	pass := os.Getenv("REDIS_PASS")
	dbEnv := os.Getenv("REDIS_DB")
	db, err := strconv.Atoi(dbEnv)
	if err != nil {
		log.Panic("error pada konversi db redis:", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: pass,
		DB:       db,
	})

	return rdb
}
