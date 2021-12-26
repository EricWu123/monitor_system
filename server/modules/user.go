package modules

import "monitor_system/internal/dao"

func LoginAuth(loginPass string, userName string, user dao.IUser) (bool, error) {
	return user.LoginAuth(loginPass, userName)
}
