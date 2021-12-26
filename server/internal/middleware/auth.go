package middleware

import (
	"monitor_system/config"
	"monitor_system/global"
	"monitor_system/logging"
	"monitor_system/session"
	"net/http"
	"net/url"
	"sync"

	"github.com/gin-gonic/gin"
)

type Auth struct {
	session *session.Session
	conf    *config.Conf
}

var (
	auth      *Auth
	authMutex sync.Mutex
)

func NewAuth(s *session.Session, c *config.Conf) *Auth {
	if auth != nil {
		return auth
	}

	authMutex.Lock()
	defer authMutex.Unlock()

	// double check
	if auth != nil {
		return auth
	}

	auth = &Auth{}
	auth.conf = c
	auth.session = s
	return auth
}

func (a *Auth) Autheticate(context *gin.Context) {
	sessionid, err := context.Request.Cookie("sessionid")
	if err != nil {
		logging.LogInfo("no rights access. err:", err)
		context.JSON(http.StatusOK, gin.H{"code": global.FAILED, "msg": "no rights access."})
		context.Abort()
		return
	}
	a.session.SessionID, _ = url.QueryUnescape(sessionid.Value)
	if !a.session.Get() {
		logging.LogInfo("session id is not valid.")
		context.JSON(http.StatusOK, gin.H{"code": global.FAILED, "msg": "session id is not valid."})
		context.Abort()
		return
	}
	a.session.Update()
	logging.LogInfo("auth success. sess:", a.session.SessionID)
}

func (a *Auth) ReportAuth(context *gin.Context) {
	authorization := context.Request.Header.Get("Authorization")
	if authorization != "APPCODE "+a.conf.Server.AppCode {
		logging.LogInfo("appcode is not valid.")
		context.JSON(http.StatusOK, gin.H{"code": global.FAILED, "msg": "appcode is not valid."})
		context.Abort()
	}
	logging.LogInfo("ReportAuth success.")
}
