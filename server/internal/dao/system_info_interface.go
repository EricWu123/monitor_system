package dao

import "monitor_system/model"

type ISystemInfo interface {
	Get(condtion string) ([]model.SystemInfo, error)
	Save(*model.SystemInfo) error
}
