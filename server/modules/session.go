package modules

import "monitor_system/internal/dao"

func IsValid(sessionID string, s dao.ISession) (bool, error) {
	return s.IsValid(sessionID)
}

func Set(userID string, s dao.ISession) (string, error) {
	return s.Set(userID)
}

func Update(sessionID string, s dao.ISession) error {
	return s.Update(sessionID)
}
