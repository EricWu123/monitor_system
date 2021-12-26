package dao

import (
	"database/sql"
	"monitor_system/config"
	"monitor_system/internal/cipher"
	"monitor_system/internal/db"
	"monitor_system/model"
)

type UserRepo struct {
	mysql     *sql.DB
	cipherKey string
}

func NewUserMysql(conf *config.Conf) *UserRepo {
	user := &UserRepo{}
	user.mysql = db.NewMysql(conf.DB)
	user.cipherKey = conf.Server.CipherKey
	return user
}

func (u *UserRepo) LoginAuth(loginPass string, userName string) (bool, error) {
	user := &model.User{}
	row := u.mysql.QueryRow("select * from user where userName = ?;", userName)
	e := row.Scan(&user.UserName, &user.Password)
	if e != nil {
		return false, e
	}
	// 解密数据库中的密码
	password, e := cipher.AesDecrypt(user.Password, []byte(u.cipherKey))
	if e != nil {
		return false, e
	}

	// 解密用户输入的密码
	loginPassord, e := cipher.AesDecrypt(loginPass, []byte(u.cipherKey))
	if e != nil {
		return false, e
	}

	if password == loginPassord {
		return true, nil
	}
	return false, nil
}
