package service

import (
	"errors"

	"github.com/jz0ojiang/go-maimai-player-counter/db"
	"gorm.io/gorm"
)

type City struct {
	Code         int    `json:"code"`
	Name         string `json:"name"`
	ProvinceCode int    `json:"province_code"`
	FullCode     string `json:"full_code"`
}

func GetCityList() ([]City, error) {
	city, err := db.GetCityList()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []City{}, nil
		}
		return nil, err
	}
	var cities []City
	for _, v := range city {
		cities = append(cities, City{
			Code:         v.Code,
			Name:         v.Name,
			ProvinceCode: v.ProvinceCode,
			FullCode:     ToFullCode(v.Code),
		})
	}
	return cities, nil
}

func GetCityByCode[T int | string](cityCode T) (City, error) {
	var code any = cityCode
	if strCode, ok := code.(string); ok && len(strCode) == 12 {
		code = strCode[0:4]
	}
	if intCode, ok := code.(int); ok && intCode > 1e8 {
		code = intCode / 1e8
	}
	city, err := db.GetCityByCode(code)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return City{}, nil
		}
		return City{}, err
	}
	return City{
		Code:         city.Code,
		Name:         city.Name,
		ProvinceCode: city.ProvinceCode,
		FullCode:     ToFullCode(city.Code),
	}, nil
}

func GetCityListByProvinceCode[T int | string](provinceCode T) ([]City, error) {
	var code any = provinceCode
	if strCode, ok := code.(string); ok && len(strCode) == 12 {
		code = strCode[0:2]
	}
	if intCode, ok := code.(int); ok && intCode > 1e10 {
		code = intCode / 1e10
	}
	city, err := db.GetCityListByProvinceCode(code)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []City{}, nil
		}
		return nil, err
	}
	var cities []City
	for _, v := range city {
		cities = append(cities, City{
			Code:         v.Code,
			Name:         v.Name,
			ProvinceCode: v.ProvinceCode,
			FullCode:     ToFullCode(v.Code),
		})
	}
	return cities, nil
}

func GetCityByName(name string) (City, error) {
	city, err := db.GetCityByName(name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return City{}, nil
		}
		return City{}, err
	}
	return City{
		Code:         city.Code,
		Name:         city.Name,
		ProvinceCode: city.ProvinceCode,
		FullCode:     ToFullCode(city.Code),
	}, nil
}
