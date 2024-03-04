package main

import (
	"Book/model"
	"Book/router"
	"Book/tools"
	"fmt"
)

// @contact.name   Book API
// @contact.email  香香编程喵喵喵
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	model.InitConf()
	defer model.RedisClose(model.ConnectRedis)
	defer model.MysqlClose(model.ConnectMysql)
	fmt.Println(model.ConnectRedis)
	tools.Snow, _ = tools.NewSnowflakeIDGenerator(1)
	router.InitRouter()
}

// var s = gin.Default()
// s.GET("/alipay", payUrl)
// s.GET("/callback", callback)
// s.POST("/notify", notify)
// s.Run(":8080")
