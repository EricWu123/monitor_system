package model

import (
	"errors"
	"monitor_system/config"
	"monitor_system/internal/db"
	"monitor_system/logging"
	"strings"
	"time"
)

type SystemInfo struct {
	HostName string   `json:"HostName"`
	OS       string   `json:"OS"`
	IPs      []string `json:"IPs"`
}

func (s *SystemInfo) Save() error {
	conf := config.GetConfig()
	mysql := db.NewMysql(conf.DB)
	tx, err := mysql.Begin() // 开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		logging.LogInfo("insert error, begin trasaction failed.")
		return err
	}
	sqlStr := "insert into system_info (hostName,os,IPs,created) values(?,?,?,?);"
	ret, err := tx.Exec(sqlStr, s.HostName, s.OS, strings.Join(s.IPs, ","), time.Now().Unix())
	if err != nil {
		tx.Rollback() // 回滚
		logging.LogInfo("exec sqlStr failed, err:%v, sql:%v", err, sqlStr)
		return err
	}
	affRow, err := ret.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		logging.LogInfo("exec RowsAffected failed, err:%v", err)
		return err
	}

	if affRow == 1 {
		tx.Commit() // 提交事务
	} else {
		tx.Rollback() // 回滚
		logging.LogInfo("insert failed, affRow:%v", affRow)
		return errors.New("insert failed")
	}

	return nil
}

func GetSystemInfo(condSql string) ([]SystemInfo, error) {
	conf := config.GetConfig()
	mysql := db.NewMysql(conf.DB)
	var sysInfos []SystemInfo
	rows, err := mysql.Query("select hostName, os, IPs from system_info" + condSql)
	if err != nil {
		logging.LogInfo("Query system info failed. err:%v, sql:%v", err, condSql)
		return nil, errors.New("query system info failed")
	}
	for rows.Next() {
		sysInfo := SystemInfo{}
		var IPs string

		if err := rows.Scan(&sysInfo.HostName, &sysInfo.OS, &IPs); err != nil {
			logging.LogInfo("Query product failed. err:", err)
			return nil, errors.New("query product failed")
		}
		sysInfo.IPs = strings.Split(IPs, ",")
		sysInfos = append(sysInfos, sysInfo)
	}
	return sysInfos, nil
}
