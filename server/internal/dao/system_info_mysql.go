package dao

import (
	"database/sql"
	"errors"
	"monitor_system/config"
	"monitor_system/internal/db"
	"monitor_system/logging"
	"monitor_system/model"
	"strings"
	"time"
)

type SystemInfoMysqlRepo struct {
	mysql *sql.DB
}

func NewSystemInfoMysql(conf *config.DBConfig) *SystemInfoMysqlRepo {
	sys := &SystemInfoMysqlRepo{}
	sys.mysql = db.NewMysql(conf)
	return sys
}

func (s *SystemInfoMysqlRepo) Get(condSql string) ([]model.SystemInfo, error) {
	var sysInfos []model.SystemInfo
	rows, err := s.mysql.Query("select hostName, os, IPs from system_info" + condSql)
	if err != nil {
		logging.LogInfo("Query system info failed. err:%v, sql:%v", err, condSql)
		return nil, errors.New("query system info failed")
	}
	for rows.Next() {
		sysInfo := model.SystemInfo{}
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

func (s *SystemInfoMysqlRepo) Save(info *model.SystemInfo) error {
	tx, err := s.mysql.Begin() // 开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		logging.LogInfo("insert error, begin trasaction failed.")
		return err
	}
	sqlStr := "insert into system_info (hostName,os,IPs,created) values(?,?,?,?);"
	ret, err := tx.Exec(sqlStr, info.HostName, info.OS, strings.Join(info.IPs, ","), time.Now().Unix())
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
