package dao

type IUser interface {
	LoginAuth(loginpass string, userName string) (bool, error)
}
