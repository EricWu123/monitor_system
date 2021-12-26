package modules

import (
	"monitor_system/internal/dao"
	"monitor_system/model"
)

func GetSystemInfo(cond string, i dao.ISystemInfo) ([]model.SystemInfo, error) {
	return i.Get(cond)
}

func SaveSystemInfo(s *model.SystemInfo, i dao.ISystemInfo) error {
	return i.Save(s)
}
