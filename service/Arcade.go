package service

import (
	"errors"
	"math/rand"
	"time"

	"github.com/jz0ojiang/go-maimai-player-counter/db"
	"gorm.io/gorm"
)

type Arcade struct {
	ID           int      `json:"arcade_id"`
	Name         string   `json:"arcade_name"`
	MachineCount int      `json:"machine_count"`
	Address      string   `json:"address"`
	Province     Province `json:"province"`
	City         City     `json:"city"`
	UpdatedAt    int64    `json:"updated_at"`
}

// 获取机厅列表（从数据库中获取）
func GetArcadeList() ([]Arcade, error) {
	arcade, err := db.GetArcadeList()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []Arcade{}, nil
		}
		return nil, err
	}
	var arcades []Arcade
	for _, v := range arcade {
		province, err := GetProvinceByCode(v.ProvinceCode)
		if err != nil {
			province = Province{}
		}
		city, err := GetCityByCode(v.CityCode)
		if err != nil {
			city = City{}
		}
		arcades = append(arcades, Arcade{
			ID:           v.ID,
			Name:         v.Name,
			MachineCount: v.MachineCount,
			Address:      v.Address,
			Province:     province,
			City:         city,
		})
	}
	return arcades, nil
}

// 获取机厅信息，如果数据库中没有，则从 Wahlap 获取并存入数据库
// 如果 Wahlap 中也没有，则返回错误
func GetArcadeByArcadeID(arcadeID int) (Arcade, error) {
	if !CheckArcadeExistByArcadeID(arcadeID) {
		if err := AddArcadeByArcadeIdWithArcadeMap(arcadeID); err != nil {
			return Arcade{}, errors.New("arcade not found")
		}
	}
	_arcade, _ := db.GetArcadeByArcadeID(arcadeID)
	province, err := GetProvinceByCode(_arcade.ProvinceCode)
	if err != nil {
		province = Province{}
	}
	city, err := GetCityByCode(_arcade.CityCode)
	if err != nil {
		city = City{}
	}
	return Arcade{
		ID:           _arcade.ID,
		Name:         _arcade.Name,
		MachineCount: _arcade.MachineCount,
		Address:      _arcade.Address,
		Province:     province,
		City:         city,
	}, nil
}

// 根据城市代码获取机厅列表（从数据库中获取）
func GetArcadeListByCityCode[T any](cityCode T) ([]Arcade, error) {
	arcade, err := db.GetArcadeListByCityCode(cityCode)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []Arcade{}, nil
		}
		return nil, err
	}
	var arcades []Arcade
	for _, v := range arcade {
		province, err := GetProvinceByCode(v.ProvinceCode)
		if err != nil {
			province = Province{}
		}
		city, err := GetCityByCode(v.CityCode)
		if err != nil {
			city = City{}
		}
		arcades = append(arcades, Arcade{
			ID:           v.ID,
			Name:         v.Name,
			MachineCount: v.MachineCount,
			Address:      v.Address,
			Province:     province,
			City:         city,
		})
	}
	return arcades, nil
}

// 根据省份代码获取机厅列表（从数据库中获取）
func GetArcadeListByProvinceCode[T any](provinceCode T) ([]Arcade, error) {
	arcade, err := db.GetArcadeListByProvinceCode(provinceCode)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []Arcade{}, nil
		}
		return nil, err
	}
	var arcades []Arcade
	for _, v := range arcade {
		province, err := GetProvinceByCode(v.ProvinceCode)
		if err != nil {
			province = Province{}
		}
		city, err := GetCityByCode(v.CityCode)
		if err != nil {
			city = City{}
		}
		arcades = append(arcades, Arcade{
			ID:           v.ID,
			Name:         v.Name,
			MachineCount: v.MachineCount,
			Address:      v.Address,
			Province:     province,
			City:         city,
		})
	}
	return arcades, nil
}

func CreateArcade(arcade Arcade) error {
	err := db.CreateArcade(db.Arcade{
		ID:           arcade.ID,
		Name:         arcade.Name,
		MachineCount: arcade.MachineCount,
		Address:      arcade.Address,
		ProvinceCode: arcade.Province.Code,
		CityCode:     arcade.City.Code,
	})
	return err
}

// 向数据库中添加机厅
// name 机厅名称
// machineCount 机台数量
// address 机厅地址
// provinceCode 省份代码
// cityCode 城市代码
func CreateCustomArcade(name string,
	machineCount int,
	address string,
	provinceCode int,
	cityCode int) (int, error) {
	timestamp := time.Now().Unix()
	uniqueID := timestamp*100 + int64(rand.Intn(100))
	err := db.CreateArcade(db.Arcade{
		ID:           int(uniqueID),
		Name:         name,
		MachineCount: machineCount,
		Address:      address,
		ProvinceCode: provinceCode,
		CityCode:     cityCode,
	})
	return int(uniqueID), err
}

// 更新数据库中的机厅信息
func UpdateArcade(arcade Arcade) error {
	err := db.UpdateArcade(db.Arcade{
		ID:           arcade.ID,
		Name:         arcade.Name,
		MachineCount: arcade.MachineCount,
		Address:      arcade.Address,
		ProvinceCode: arcade.Province.Code,
		CityCode:     arcade.City.Code,
	})
	return err
}

// 根据机厅 ID 删除机厅（从数据库中删除）
func DeleteArcadeByArcadeID[T any](arcadeID T) error {
	err := db.DeleteArcadeByArcadeID(arcadeID)
	return err
}

// 检查数据库中是否存在某个机厅
func CheckArcadeExistByArcadeID[T any](arcadeID T) bool {
	arcade, err := db.GetArcadeByArcadeID(arcadeID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false
		}
		return false
	}
	if arcade.ID > 0 {
		return true
	}
	return false
}
