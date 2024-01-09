package db

type Arcade struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	MachineCount int `gorm:"column:machineCount"`
	Address      string
	ProvinceCode int   `gorm:"foreignKey;column:provinceCode"`
	CityCode     int   `gorm:"foreignKey;column:cityCode"`
	UpdatedAt    int64 `gorm:"column:updatedAt"`
}

func GetArcadeList() ([]Arcade, error) {
	var arcades []Arcade
	result := SqliteDB.Table("arcade").Find(&arcades)
	return arcades, result.Error
}

func GetArcadeByArcadeID[T any](arcadeID T) (Arcade, error) {
	var arcade Arcade
	result := SqliteDB.Table("arcade").First(&arcade, arcadeID)
	return arcade, result.Error
}

func GetArcadeListByCityCode[T any](cityCode T) ([]Arcade, error) {
	var arcades []Arcade
	result := SqliteDB.Table("arcade").Where("cityCode = ?", cityCode).Find(&arcades)
	return arcades, result.Error
}

func GetArcadeListByProvinceCode[T any](provinceCode T) ([]Arcade, error) {
	var arcades []Arcade
	result := SqliteDB.Table("arcade").Where("provinceCode = ?", provinceCode).Find(&arcades)
	return arcades, result.Error
}

func CreateArcade(arcade Arcade) error {
	result := SqliteDB.Table("arcade").Create(&arcade)
	return result.Error
}

func UpdateArcade(arcade Arcade) error {
	result := SqliteDB.Table("arcade").Save(&arcade)
	return result.Error
}

func DeleteArcadeByArcadeID[T any](arcadeID T) error {
	result := SqliteDB.Table("arcade").Delete(&Arcade{}, arcadeID)
	return result.Error
}
