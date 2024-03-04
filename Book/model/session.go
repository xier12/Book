package model

import (
	"github.com/gin-gonic/gin"
	"github.com/rbcervilla/redisstore/v9"
)

var store *redisstore.RedisStore

// var store = sessions.NewCookieStore([]byte("lexington"))
var sessionName = "session-name"

func GetSession(c *gin.Context) map[interface{}]interface{} {
	session, _ := store.Get(c.Request, sessionName)
	//fmt.Printf("session:%+v\n", session.Values)
	return session.Values
}

func SetSession(c *gin.Context, name string, id int64) error {
	session, _ := store.Get(c.Request, sessionName)
	session.Values["name"] = name
	session.Values["id"] = id
	return session.Save(c.Request, c.Writer)
}

func FlushSession(c *gin.Context) error {
	session, _ := store.Get(c.Request, sessionName)
	//fmt.Printf("session : %+v\n", session.Values)
	session.Values["name"] = ""
	session.Values["id"] = int64(0)
	return session.Save(c.Request, c.Writer)
}
