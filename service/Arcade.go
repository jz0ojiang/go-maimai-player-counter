package service

import (
	"errors"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/bytedance/sonic"
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
	var arcade Arcade
	if !CheckArcadeExistByArcadeID(arcadeID) {
		if CheckArcadeExistInWahlapByArcadeID(arcadeID) {
			arcade, _ = GetArcadeWithWahlapById(arcadeID)
			CreateArcade(arcade)
		} else {
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

// Generated by https://quicktype.io

type WahlapArcade struct {
	PlaceID      string `json:"placeId"`
	MachineCount int64  `json:"machineCount"`
	ID           string `json:"id"`
	Province     string `json:"province"`
	ArcadeName   string `json:"arcadeName"`
	Mall         string `json:"mall"`
	Address      string `json:"address"`
}

// {
// 	"placeId": "1002",
// 	"machineCount": 3,
// 	"id": "1002",
// 	"province": "广东",
// 	"arcadeName": "环游嘉年华番禺易发店",
// 	"mall": "环游嘉年华番禺易发店",
// 	"address": "广东省广州市番禺区市桥街易发商业街新大新百货5楼"
// }

// 从 Wahlap 获取机厅列表
func GetArcadeListWithWahlap() ([]Arcade, error) {
	// 从 https://wc.wahlap.net/maidx/rest/location 获取数据并解析
	response, err := http.Get("https://wc.wahlap.net/maidx/rest/location")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var wahlapArcadeList []WahlapArcade
	err = sonic.Unmarshal(body, &wahlapArcadeList)
	if err != nil {
		return nil, err
	}
	var arcades []Arcade
	for _, v := range wahlapArcadeList {
		provinceObj, _ := GetProvinceByName(v.Province)
		id, err := strconv.ParseInt(v.ID, 10, 64)
		if err != nil {
			id, _ = strconv.ParseInt(v.PlaceID, 10, 64)
		}
		arcades = append(arcades, Arcade{
			ID:           int(id),
			Name:         v.ArcadeName,
			MachineCount: int(v.MachineCount),
			Address:      v.Address,
			Province:     provinceObj,
			City:         GetCityByAddress(v.Address, v.Province),
		})
	}
	return arcades, nil
}

// 从 Wahlap 获取某个城市的机厅列表
func GetArcadeListWithWahlapByCityCode[T int | string](cityCode T) ([]Arcade, error) {
	arcadeList, err := GetArcadeListWithWahlap()
	if err != nil {
		return nil, err
	}
	var arcades []Arcade
	city, err := GetCityByCode(cityCode)
	if err != nil {
		return nil, err
	}
	for _, v := range arcadeList {
		if v.City.Code == city.Code {
			arcades = append(arcades, v)
		}
	}
	return arcades, nil
}

// 根据 Wahlap 机厅 ID 获取机厅信息
func GetArcadeListWithWahlapByProvinceCode[T int | string](provinceCode T) ([]Arcade, error) {
	arcadeList, err := GetArcadeListWithWahlap()
	if err != nil {
		return nil, err
	}
	var arcades []Arcade
	province, err := GetProvinceByCode(provinceCode)
	if err != nil {
		return nil, err
	}
	for _, v := range arcadeList {
		if v.Province.Code == province.Code {
			arcades = append(arcades, v)
		}
	}
	return arcades, nil
}

// 从 Wahlap 获取某个机厅的信息
func GetArcadeWithWahlapById(id int) (Arcade, error) {
	arcadeList, err := GetArcadeListWithWahlap()
	if err != nil {
		return Arcade{}, err
	}
	for _, v := range arcadeList {
		if v.ID == id {
			return v, nil
		}
	}
	return Arcade{}, errors.New("arcade not found")
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

// 检查 Wahlap 中是否存在某个机厅
func CheckArcadeExistInWahlapByArcadeID(arcadeID int) bool {
	arcadeList, err := GetArcadeListWithWahlap()
	if err != nil {
		return false
	}
	for _, v := range arcadeList {
		if v.ID == arcadeID {
			return true
		}
	}
	return false
}

// 遍历从 Wahlap 获取的机厅列表
// 判断本地数据库中是否存在
// 如果不存在则添加
func CheckArcadeExistInWahlapAndAdd() error {
	arcadeList, err := GetArcadeListWithWahlap()
	if err != nil {
		return err
	}
	for _, v := range arcadeList {
		if !CheckArcadeExistByArcadeID(v.ID) {
			CreateArcade(v)
		}
	}
	return nil
}

// 遍历从 Wahlap 获取的机厅列表和本地数据库中的机厅列表
// 检查 Wahlap 中的信息是否有更新
// 如果有更新则更新本地数据库
func CheckArcadeExistInWahlapAndUpdate() error {
	wahlapArcadeList, err := GetArcadeListWithWahlap()
	if err != nil {
		return err
	}
	databaseArcadeList, err := GetArcadeList()
	if err != nil {
		return err
	}
	for _, v := range wahlapArcadeList {
		for _, _v := range databaseArcadeList {
			if v.ID == _v.ID {
				if v.Name != _v.Name || v.MachineCount != _v.MachineCount || v.Address != _v.Address {
					UpdateArcade(v)
				}
			}
		}
	}
	return nil
}
