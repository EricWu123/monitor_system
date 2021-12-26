package session

import (
	"database/sql"
	"errors"
	"fmt"
	"monitor_system/config"
	mylog "monitor_system/logging"
	"sync"
	"time"

	"github.com/rs/xid"
)

var (
	mysession    *Session
	sessionMutex sync.Mutex
)

type Session struct {
	UserID       string
	SessionID    string
	TimeAccessed time.Time
	MaxAge       int64
	mysqlDB      *sql.DB
}

func NewSession(db *sql.DB) *Session {
	if mysession != nil {
		return mysession
	}

	sessionMutex.Lock()
	defer sessionMutex.Unlock()

	// double check
	if mysession != nil {
		return mysession
	}

	mysession = &Session{}
	mysession.mysqlDB = db
	mysession.MaxAge = int64(config.GetConfig().Session.MaxAge)
	return mysession
}

func (s *Session) genSessionId() string {
	return xid.New().String()
}

func (s *Session) Get() bool {
	row := s.mysqlDB.QueryRow("select * from session where sessionID = ?;", s.SessionID)
	e := row.Scan(&s.UserID, &s.SessionID, &s.TimeAccessed)
	if e != nil {
		mylog.LogInfo("query failed. session id:%v, err:%v", s.SessionID, e)
		return false
	}
	curTime := time.Now().Unix()
	if curTime-s.TimeAccessed.Unix() > s.MaxAge {
		mylog.LogInfo("sid time out. session id:%v", s.SessionID)
		return false
	}
	return true
}

func (s *Session) Set(userID string) error {
	tx, err := s.mysqlDB.Begin() // 开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		mylog.LogInfo("insert error, begin trasaction failed.")
		return err
	}

	sqlStr := fmt.Sprintf("select count(*) from session where userID = '%s' limit 1;", userID)
	result := tx.QueryRow(sqlStr)
	var count int
	err = result.Scan(&count)
	if err != nil {
		mylog.LogInfo("exec sqlStr failed, err:%v, sql:%v", err, sqlStr)
		tx.Rollback() // 回滚
		return err
	}

	curTime := time.Now().Format("2006-01-02 15:04:05")
	s.SessionID = s.genSessionId()
	if count == 0 {
		sqlStr = fmt.Sprintf(
			"insert into session values('%s', '%s', '%s')",
			userID,
			s.SessionID,
			curTime,
		)
	} else {
		sqlStr = fmt.Sprintf(
			"update session set expirationTime='%s',sessionID = '%s' where userID = '%s'",
			curTime,
			s.SessionID,
			userID,
		)
	}

	ret, err := tx.Exec(sqlStr)
	if err != nil {
		tx.Rollback() // 回滚
		mylog.LogInfo("exec sqlStr failed, err:%v, sql:%v", err, sqlStr)
		return err
	}
	affRow, err := ret.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		mylog.LogInfo("exec ret.RowsAffected() failed, err:%v", err)
		return err
	}

	if affRow != 1 {
		mylog.LogInfo("Set affRow:", affRow)
		tx.Rollback()
		return errors.New("insert session failed")
	}
	tx.Commit() // 提交事务
	mylog.LogInfo("insert session success. sessionid:", s.SessionID)
	return nil
}

func (s *Session) Update() error {
	tx, err := s.mysqlDB.Begin() // 开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		mylog.LogInfo("update error, begin trasaction failed.")
		return err
	}
	curTime := time.Now().Format("2006-01-02 15:04:05")
	sqlStr := fmt.Sprintf(
		"update session set expirationTime='%s' where sessionID='%s'",
		curTime,
		s.SessionID,
	)
	ret, err := tx.Exec(sqlStr)
	if err != nil {
		tx.Rollback() // 回滚
		mylog.LogInfo("exec sqlStr failed, err:%v, sql:%v", err, sqlStr)
		return err
	}
	affRow, err := ret.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		mylog.LogInfo("exec RowsAffected failed, err:%v", err)
		return err
	}

	if affRow != 1 {
		mylog.LogInfo("Update affRow:", affRow)
		tx.Rollback()
		return errors.New("update session failed")
	}
	tx.Commit() // 提交事务
	mylog.LogInfo("update session success.")
	return nil
}

func (s *Session) Delete() error {
	tx, err := s.mysqlDB.Begin() // 开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		mylog.LogInfo("delete error, begin trasaction failed.")
		return err
	}
	sqlStr := fmt.Sprintf("delete from session where sessionID='%s'", s.SessionID)
	ret, err := tx.Exec(sqlStr)
	if err != nil {
		tx.Rollback() // 回滚
		mylog.LogInfo("exec sqlStr failed, err:%v, sql:%v", err, sqlStr)
		return err
	}
	affRow, err := ret.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		mylog.LogInfo("exec RowsAffected failed, err:%v", err)
		return err
	}
	if affRow != 1 {
		mylog.LogInfo("Delete affRow:", affRow)
		tx.Rollback()
		return errors.New("delete session failed")
	}

	tx.Commit()
	mylog.LogInfo("delete session success.")
	return nil
}
