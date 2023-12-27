package service

import (
	"errors"

	"gorm.io/gorm"

	"github.com/jz0ojiang/go-maimai-player-counter/db"
)

type Province struct {
	Code     int    `json:"code"`
	Name     string `json:"name"`
	FullCode string `json:"full_code"`
}

func GetProvinceList() ([]Province, error) {
	province, err := db.GetProvinceList()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []Province{}, nil
		}
		return nil, err
	}
	var provinces []Province
	for _, v := range province {
		provinces = append(provinces, Province{
			Code:     v.Code,
			Name:     v.Name,
			FullCode: ToFullCode(v.Code),
		})
	}
	return provinces, nil
}

func GetProvinceByCode[T int | string](provinceCode T) (Province, error) {
	var code any = provinceCode
	if strCode, ok := code.(string); ok && len(strCode) == 12 {
		code = strCode[0:2]
	}
	if intCode, ok := code.(int); ok && intCode > 1e10 {
		code = intCode / 1e10
	}
	province, err := db.GetProvinceByCode(code)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Province{}, nil
		}
		return Province{}, err
	}
	return Province{
		Code:     province.Code,
		Name:     province.Name,
		FullCode: ToFullCode(province.Code),
	}, nil
}

func GetProvinceByName(name string) (Province, error) {
	province, err := db.GetProvinceByName(name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Province{}, nil
		}
		return Province{}, err
	}
	return Province{
		Code:     province.Code,
		Name:     province.Name,
		FullCode: ToFullCode(province.Code),
	}, nil
}
