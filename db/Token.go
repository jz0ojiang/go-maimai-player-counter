package db

type Token struct {
	ID     int `gorm:"primaryKey"`
	Token  string
	Remark string
}

func GetTokenList() ([]Token, error) {
	var tokens []Token
	result := SqliteDB.Table("token").Find(&tokens)
	return tokens, result.Error
}

func GetTokenByToken[T any](t T) (Token, error) {
	var token Token
	result := SqliteDB.Table("token").Where("token = ?", t).First(&token)
	return token, result.Error
}

func GetTokenById[T any](id T) (Token, error) {
	var token Token
	result := SqliteDB.Table("token").First(&token, id)
	return token, result.Error
}

func CreateToken(token Token) error {
	result := SqliteDB.Table("token").Create(&token)
	return result.Error
}

func UpdateToken(token Token) error {
	result := SqliteDB.Table("token").Save(&token)
	return result.Error
}

func DeleteTokenByToken[T any](token T) error {
	result := SqliteDB.Table("token").Where("token = ?", token).Delete(&Token{})
	return result.Error
}

func DeleteTokenById[T any](id T) error {
	result := SqliteDB.Table("token").Delete(&Token{}, id)
	return result.Error
}

func CheckTokenExistByToken[T any](t T) bool {
	token, err := GetTokenByToken(t)
	if err != nil {
		return false
	}
	if token.ID > 0 {
		return true
	}
	return false
}
