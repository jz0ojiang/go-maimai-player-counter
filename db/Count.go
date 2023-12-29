package db

import (
	"sort"
	"strconv"

	"github.com/bytedance/sonic"
)

type CountLog struct {
	ArcadeId        int
	Count           int
	UpdateTimestamp int64
	Type            int
}

// key: timestamp
// value：Count.(json)

// CountLog 以时间戳降序排列（最新的在前面）
func GetLatestCountByArcadeID(arcadeID int) (int, error) {
	counts, err := GetCountLogsByArcadeID(arcadeID)
	if err != nil {
		return 0, err
	}
	if len(counts) == 0 {
		return 0, nil
	}
	sort.Slice(counts, func(i, j int) bool {
		return counts[i].UpdateTimestamp > counts[j].UpdateTimestamp
	})
	return counts[0].Count, nil
}

// GetCountLogByTimeStamp 根据时间戳获取 CountLog
func GetCountLogByTimeStamp(timestamp int64) (CountLog, error) {
	var count CountLog
	data, err := LevelDB.Get([]byte(strconv.FormatInt(timestamp, 10)), nil)
	if err != nil {
		return count, err
	}
	err = sonic.Unmarshal(data, &count)
	return count, err
}

// GetCountLogsByArcadeID 根据 ArcadeID 获取 CountLog
// 返回的 CountLog 以时间戳降序排列（最新的在前面）
func GetCountLogsByArcadeID(arcadeID int) ([]CountLog, error) {
	var counts []CountLog
	iter := LevelDB.NewIterator(nil, nil)
	defer iter.Release()
	for iter.Next() {
		var log CountLog
		err := sonic.Unmarshal(iter.Value(), &log)
		if err != nil {
			continue
		}
		if log.ArcadeId == arcadeID {
			counts = append(counts, log)
		}
	}
	// 排序
	sort.Slice(counts, func(i, j int) bool {
		return counts[i].UpdateTimestamp > counts[j].UpdateTimestamp
	})
	return counts, nil
}

func AddCountLog(count CountLog) error {
	data, err := sonic.Marshal(count)
	if err != nil {
		return err
	}
	return LevelDB.Put([]byte(strconv.FormatInt(count.UpdateTimestamp, 10)), data, nil)
}

func DeleteCountLog(timestamp int64) error {
	return LevelDB.Delete([]byte(strconv.FormatInt(timestamp, 10)), nil)
}

func GetAllCountLogs() ([]CountLog, error) {
	var counts []CountLog
	iter := LevelDB.NewIterator(nil, nil)
	defer iter.Release()
	for iter.Next() {
		var log CountLog
		err := sonic.Unmarshal(iter.Value(), &log)
		if err != nil {
			return nil, err
		}
		counts = append(counts, log)
	}
	// 排序
	sort.Slice(counts, func(i, j int) bool {
		return counts[i].UpdateTimestamp > counts[j].UpdateTimestamp
	})
	return counts, nil
}

func DeleteAllCountLogs() error {
	iter := LevelDB.NewIterator(nil, nil)
	defer iter.Release()
	for iter.Next() {
		err := LevelDB.Delete(iter.Key(), nil)
		if err != nil {
			return err
		}
	}
	return nil
}
