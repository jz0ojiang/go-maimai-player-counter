package db

type City struct {
	Code         int `gorm:"primaryKey"`
	Name         string
	ProvinceCode int `gorm:"foreignKey;column:provinceCode"`
}

func GetCityList() ([]City, error) {
	var cities []City
	result := SqliteDB.Table("city").Find(&cities)
	return cities, result.Error
}

func GetCityByCode[T any](code T) (City, error) {
	var city City
	result := SqliteDB.Table("city").First(&city, code)
	return city, result.Error
}

func GetCityListByProvinceCode[T any](provinceCode T) ([]City, error) {
	var cities []City
	result := SqliteDB.Table("city").Where("provinceCode = ?", provinceCode).Find(&cities)
	return cities, result.Error
}

func GetCityByName(name string) (City, error) {
	var city City
	result := SqliteDB.Table("city").Where("name LIKE ?", "%"+name+"%").First(&city)
	return city, result.Error
}
