package auth

import (
	"context"
	"database/sql"
	"log"
	"strconv"
	"time"

	"github.com/adafatya/noobeeid-product_catalog/entity"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type AuthRepository struct {
	db  *sqlx.DB
	rdb *redis.Client
}

func NewAuthRepository(db *sqlx.DB, rdb *redis.Client) AuthRepository {
	return AuthRepository{
		db:  db,
		rdb: rdb,
	}
}

func (r AuthRepository) Store(auth entity.Auth) error {
	// insert new data to db
	_, err := r.db.Exec(`INSERT INTO auth (email, password, role) VALUES($1, $2, $3)`, auth.Email, auth.Password, auth.Role)
	if err != nil {
		log.Println("terdapat error saat menyimpan auth: ", err)
		return err
	}

	return nil
}

func (r AuthRepository) GetByEmail(email string) (entity.Auth, error) {
	var auth entity.Auth

	// get data from db
	err := r.db.Get(&auth, `SELECT id, email, password, role FROM auth WHERE email = $1`, email)
	if err == sql.ErrNoRows {
		return entity.Auth{}, nil
	}
	if err != nil {
		log.Println("terdapat error saat mengambil auth: ", err)
		return entity.Auth{}, err
	}

	return auth, nil
}

func (r AuthRepository) SetAuthJWTToken(id int, token string) error {
	idStr := strconv.Itoa(id)

	err := r.rdb.Set(context.Background(), idStr, token, 24*time.Hour).Err()
	if err != nil {
		log.Println("terdapat error saat set token ke redis: ", err)
	}
	return err
}

func (r AuthRepository) GetAuthJWTToken(id int) (string, error) {
	idStr := strconv.Itoa(id)

	token, err := r.rdb.Get(context.Background(), idStr).Result()
	if err != nil {
		log.Println("terdapat error saat mengambil token dari redis : ", err)
		return "", err
	}

	return token, nil
}
