package service

import (
	"errors"
	"strconv"

	"github.com/jz0ojiang/go-maimai-player-counter/db"
)

func GetCountsByCity(cityCode int) (map[string]int, error) {
	arcades, err := GetArcadeListByCityCode(cityCode)
	if err != nil {
		return nil, err
	}
	var counts = make(map[string]int)
	for _, arcade := range arcades {
		ClearExpiredCountLogsByArcadeID(arcade.ID)
		count, err := db.GetLatestCountByArcadeID(arcade.ID)
		if err != nil {
			counts[strconv.Itoa(arcade.ID)] = 0
			continue
		}
		counts[strconv.Itoa(arcade.ID)] = count
	}
	return counts, nil
}

type CountLog struct {
	ArcadeId        int   `json:"arcade_id"`
	Count           int   `json:"count"`
	UpdateTimestamp int64 `json:"update_timestamp"`
	Type            int   `json:"type"`
}

func GetCountLogByTimeStamp(timestamp int64) (CountLog, error) {
	data, err := db.GetCountLogByTimeStamp(timestamp)
	if err != nil {
		return CountLog{}, err
	}
	if data.UpdateTimestamp < GetTodayZeroTimestamp() {
		db.DeleteCountLog(data.UpdateTimestamp)
		return CountLog{}, errors.New("log expired")
	}
	return CountLog{
		ArcadeId:        data.ArcadeId,
		Count:           data.Count,
		UpdateTimestamp: data.UpdateTimestamp,
		Type:            data.Type,
	}, nil
}

func GetCountLogsByArcadeID(arcadeID int) ([]CountLog, error) {
	if !CheckArcadeExistByArcadeID(arcadeID) {
		if CheckArcadeExistInWahlapByArcadeID(arcadeID) {
			arcade, _ := GetArcadeWithWahlapById(arcadeID)
			CreateArcade(arcade)
		} else {
			return nil, errors.New("arcade not found")
		}
	}
	zerotimestamp := GetTodayZeroTimestamp()
	data, err := db.GetCountLogsByArcadeID(arcadeID)
	if err != nil {
		return nil, err
	}
	err = ClearExpiredCountLogsByArcadeID(arcadeID)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return []CountLog{{
			ArcadeId:        arcadeID,
			Count:           0,
			UpdateTimestamp: zerotimestamp,
			Type:            0,
		}}, nil
	}
	var logs []CountLog
	for _, v := range data {
		logs = append(logs, CountLog{
			ArcadeId:        v.ArcadeId,
			Count:           v.Count,
			UpdateTimestamp: v.UpdateTimestamp,
			Type:            v.Type,
		})
	}
	return logs, nil
}

func AddCountLog(count CountLog) error {
	if !CheckArcadeExistByArcadeID(count.ArcadeId) {
		if CheckArcadeExistInWahlapByArcadeID(count.ArcadeId) {
			arcade, _ := GetArcadeWithWahlapById(count.ArcadeId)
			CreateArcade(arcade)
		} else {
			return errors.New("arcade not found")
		}
	}
	return db.AddCountLog(db.CountLog{
		ArcadeId:        count.ArcadeId,
		Count:           count.Count,
		UpdateTimestamp: count.UpdateTimestamp,
		Type:            count.Type,
	})
}

func DeleteCountLog(timestamp int64) error {
	return db.DeleteCountLog(timestamp)
}

func GetAllCountLogs() ([]CountLog, error) {
	data, err := db.GetAllCountLogs()
	if err != nil {
		return nil, err
	}
	var logs []CountLog
	for _, v := range data {
		// if v.UpdateTimestamp < GetTodayZeroTimestamp() {
		// 	db.DeleteCountLog(v.UpdateTimestamp)
		// 	continue
		// }
		logs = append(logs, CountLog{
			ArcadeId:        v.ArcadeId,
			Count:           v.Count,
			UpdateTimestamp: v.UpdateTimestamp,
			Type:            v.Type,
		})
	}
	return logs, nil
}

func DeleteAllCountLogs() error {
	return db.DeleteAllCountLogs()
}

func ClearExpiredCountLogsByArcadeID(arcadeID int) error {
	zerotimestamp := GetTodayZeroTimestamp()
	data, err := db.GetCountLogsByArcadeID(arcadeID)
	if err != nil {
		return err
	}
	if len(data) == 0 {
		err = db.AddCountLog(db.CountLog{
			ArcadeId:        arcadeID,
			Count:           0,
			UpdateTimestamp: zerotimestamp,
			Type:            0,
		})
		return err
	}
	for _, v := range data {
		if v.UpdateTimestamp < zerotimestamp {
			db.DeleteCountLog(v.UpdateTimestamp)
		}
	}
	return nil
}
