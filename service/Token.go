package service

import (
	"crypto/md5"
	"fmt"

	"github.com/google/uuid"
	"github.com/jz0ojiang/go-maimai-player-counter/db"
)

func VerifyToken(token string) bool {
	return db.CheckTokenExistByToken(token)
}

func GenerateToken(remark ...string) (string, error) {
	id := uuid.New()
	has := md5.Sum([]byte(id.String()))
	token := fmt.Sprintf("%x", has)
	remark = append(remark, "")
	err := db.CreateToken(db.Token{
		Token:  token,
		Remark: remark[0],
	})
	if err != nil {
		return "", err
	}
	return token, nil
}

func DeleteToken(token string) error {
	return db.DeleteTokenByToken(token)
}
