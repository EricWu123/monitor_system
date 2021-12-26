package dao

import (
	"database/sql"
	"errors"
	"fmt"
	"monitor_system/config"
	"monitor_system/internal/db"
	"monitor_system/logging"
	"monitor_system/model"
	"time"

	"github.com/rs/xid"
)

type SessionMysqlRepo struct {
	mysqlDB *sql.DB
	conf    *config.Conf
}

func NewMysqlSession(conf *config.Conf) *SessionMysqlRepo {
	s := &SessionMysqlRepo{}
	s.mysqlDB = db.NewMysql(conf.DB)
	s.conf = conf
	return s
}

func (s *SessionMysqlRepo) genSessionId() string {
	return xid.New().String()
}

func (s *SessionMysqlRepo) IsValid(sessionID string) (bool, error) {
	sess := &model.Session{}
	row := s.mysqlDB.QueryRow("select * from session where sessionID = ?;", sessionID)
	e := row.Scan(&sess.UserID, &sess.SessionID, &sess.TimeAccessed)
	if e != nil {
		logging.LogInfo("query failed. session id:%v, err:%v", sessionID, e)
		return false, e
	}
	curTime := time.Now().Unix()
	if curTime-sess.TimeAccessed.Unix() > int64(s.conf.Session.MaxAge) {
		logging.LogInfo("sid time out. session id:%v", sessionID)
		return false, errors.New("session is out of date")
	}
	return true, nil
}

func (s *SessionMysqlRepo) Set(userID string) (string, error) {
	tx, err := s.mysqlDB.Begin() // 开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		logging.LogInfo("insert error, begin trasaction failed.")
		return "", err
	}

	sqlStr := fmt.Sprintf("select count(*) from session where userID = '%s' limit 1;", userID)
	result := tx.QueryRow(sqlStr)
	var count int
	err = result.Scan(&count)
	if err != nil {
		logging.LogInfo("exec sqlStr failed, err:%v, sql:%v", err, sqlStr)
		tx.Rollback() // 回滚
		return "", err
	}

	curTime := time.Now().Format("2006-01-02 15:04:05")
	sessionID := s.genSessionId()
	if count == 0 {
		sqlStr = fmt.Sprintf(
			"insert into session values('%s', '%s', '%s')",
			userID,
			sessionID,
			curTime,
		)
	} else {
		sqlStr = fmt.Sprintf(
			"update session set expirationTime='%s',sessionID = '%s' where userID = '%s'",
			curTime,
			sessionID,
			userID,
		)
	}

	ret, err := tx.Exec(sqlStr)
	if err != nil {
		tx.Rollback() // 回滚
		logging.LogInfo("exec sqlStr failed, err:%v, sql:%v", err, sqlStr)
		return "", err
	}
	affRow, err := ret.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		logging.LogInfo("exec ret.RowsAffected() failed, err:%v", err)
		return "", err
	}

	if affRow != 1 {
		logging.LogInfo("Set affRow:", affRow)
		tx.Rollback()
		return "", errors.New("insert session failed")
	}
	tx.Commit() // 提交事务
	logging.LogInfo("insert session success. sessionid:", sessionID)
	return sessionID, nil
}

func (s *SessionMysqlRepo) Update(sessionID string) error {
	tx, err := s.mysqlDB.Begin() // 开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		logging.LogInfo("update error, begin trasaction failed.")
		return err
	}
	curTime := time.Now().Format("2006-01-02 15:04:05")
	sqlStr := fmt.Sprintf(
		"update session set expirationTime='%s' where sessionID='%s'",
		curTime,
		sessionID,
	)
	ret, err := tx.Exec(sqlStr)
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

	if affRow != 1 {
		logging.LogInfo("Update affRow:", affRow)
		tx.Rollback()
		return errors.New("update session failed")
	}
	tx.Commit() // 提交事务
	logging.LogInfo("update session success.")
	return nil
}

func (s *SessionMysqlRepo) Delete(sessionID string) error {
	tx, err := s.mysqlDB.Begin() // 开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		logging.LogInfo("delete error, begin trasaction failed.")
		return err
	}
	sqlStr := fmt.Sprintf("delete from session where sessionID='%s'", sessionID)
	ret, err := tx.Exec(sqlStr)
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
	if affRow != 1 {
		logging.LogInfo("Delete affRow:", affRow)
		tx.Rollback()
		return errors.New("delete session failed")
	}

	tx.Commit()
	logging.LogInfo("delete session success.")
	return nil
}
