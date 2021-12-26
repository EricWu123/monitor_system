package dao

type ISession interface {
	IsValid(sessionID string) (bool, error)
	Set(userID string) (string, error)
	Update(sessionID string) error
	Delete(sessionID string) error
}
