package service

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jz0ojiang/go-maimai-player-counter/db"
)

var provinceListCache []Province

func init() {
	provinceListCache, _ = GetProvinceList()
}

func ToFullCode(code int) string {
	return fmt.Sprintf("%s%s", strconv.Itoa(code), strings.Repeat("0", 12-len(strconv.Itoa(code))))
}

func ToShortCode(fullCode string) int {
	// 110100000000 => 1101
	// 110000000000 => 11
	if len(fullCode) != 12 {
		return 0
	}
	code, err := strconv.Atoi(fullCode[0:4])
	if err != nil {
		return 0
	}
	// 1100 => 11
	if code%100 == 0 {
		code = code / 100
	}
	return code
}

func RFC3339ToTimestamp(timeStr string) int64 {
	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return 0
	}
	return t.UnixNano() / 1e6
}

func GetCityByAddress(address string, province string) City {
	provinceInfo, err := db.GetProvinceByName(province)
	if err != nil {
		return City{}
	}

	cities, err := GetCityListByProvinceCode(provinceInfo.Code)
	if err != nil {
		return City{}
	}

	for _, city := range cities {
		if strings.Contains(city.Name, "市") {
			cityNameWithoutSuffix := removeSuffix(city.Name, "市")
			if strings.Contains(address, cityNameWithoutSuffix) {
				return city
			}
		} else {
			if strings.Contains(address, city.Name[0:2]) {
				return city
			}
		}
	}

	provinces, _ := GetProvinceList()
	for _, province := range provinces {
		provinceNameWithoutSuffix := removeSuffix(province.Name, "省")
		if strings.Contains(address, provinceNameWithoutSuffix) {
			cityList, _ := GetCityListByProvinceCode[int](province.Code)
			if len(cityList) > 0 && cityList[0].Name == "市辖区" {
				return City{
					Code:         cityList[0].Code,
					Name:         province.Name,
					ProvinceCode: province.Code,
					FullCode:     ToFullCode(cityList[0].Code),
				}
			}
			break
		}
	}

	return City{}
}

func removeSuffix(name string, suffix string) string {
	return strings.Replace(name, suffix, "", -1)
}

// 获取本日 0:00 的时间戳
// 兼容 javascript 的时间戳
func GetTodayZeroTimestamp() int64 {
	now := time.Now()
	zeroTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return zeroTime.UnixNano() / 1e6
}
