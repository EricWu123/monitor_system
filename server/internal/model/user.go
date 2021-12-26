package model

import (
	"monitor_system/config"
	"monitor_system/internal/cipher"
	"monitor_system/internal/db"
	"monitor_system/session"

	"github.com/gin-gonic/gin"
)

// 用户
type User struct {
	UserName string // 用户名称
	Password string // 用户密码
}

func (user *User) LoginAuth(loginPass string) (bool, error) {
	conf := config.GetConfig()
	mysql := db.NewMysql(conf.DB)
	row := mysql.QueryRow("select * from user where userName = ?;", user.UserName)
	e := row.Scan(&user.UserName, &user.Password)
	if e != nil {
		return false, e
	}
	// 解密数据库中的密码
	password, e := cipher.AesDecrypt(user.Password, []byte(conf.Server.CipherKey))
	if e != nil {
		return false, e
	}

	// 解密用户输入的密码
	loginPassord, e := cipher.AesDecrypt(loginPass, []byte(conf.Server.CipherKey))
	if e != nil {
		return false, e
	}

	if password == loginPassord {
		return true, nil
	}
	return false, nil
}

func (user *User) SetSession(context *gin.Context) error {
	conf := config.GetConfig()
	mysql := db.NewMysql(conf.DB)
	session := session.NewSession(mysql)

	if e := session.Set(user.UserName); e != nil {
		return e
	}
	context.SetCookie(conf.Session.Name, session.SessionID, conf.Session.MaxAge, "/", "127.0.0.1", false, true)

	return nil
}
