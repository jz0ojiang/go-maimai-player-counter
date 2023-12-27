package db

type Province struct {
	Code int `gorm:"primaryKey"`
	Name string
}

func GetProvinceList() ([]Province, error) {
	var provinces []Province
	result := SqliteDB.Table("province").Find(&provinces)
	return provinces, result.Error
}

func GetProvinceByCode[T any](code T) (Province, error) {
	var province Province
	result := SqliteDB.Table("province").First(&province, code)
	return province, result.Error
}

func GetProvinceByName(name string) (Province, error) {
	var province Province
	result := SqliteDB.Table("province").Where("name LIKE ?", "%"+name+"%").First(&province)
	return province, result.Error
}
