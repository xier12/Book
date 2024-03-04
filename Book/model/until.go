package model

import (
	"Book/config"
	"Book/tools"
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/rbcervilla/redisstore/v9"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
)

var ConnectMysql *gorm.DB
var ConnectRedis *redis.Client
var ConnectMongoDB *mongo.Client
var ConnectES *elastic.Client
var AliPayClient *tools.AliPayClient

const (
	kAppID               = "9021000131684453"
	kPrivateKey          = "MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCCN7HrrPsyAwoDo+vMTNt7NxCwwD3hYBNrM3/hURnvi88IXHc95z3AeJ05021MmZ1x3ZbS3vLXfrYPd4rsy+VUnFs3flZarrAz4SmotvVLcLr+Ng0oWWpJp0Xeeck5MHQXcEENbLF1Tr6/7Hcg+oFTjdg12SUx1n3iqU3CrPJDDHKkbZ66umFBUMZ8xF4jXdqEQZddRSt6t+uKYwZD/86OUIDm1UzTsDlWa3XKysbBrLemowYLEIy+O36JyqgO94LgzHw0RXbwhIxbsjBkKrBRmdBDkBYunQWJNg2ZVZ5uh0dBM2pLCL0zzWoqwto2cMB8GzZ9YW1y8FyCbdTJDhexAgMBAAECggEAE6ecb1o4wB+9TzdazAd9yWdSWZtqu63owHGRh8zaAVI0+aobRXi11WtfL+89vCYjoaA0t7d3KBe0GzmL+pn8D31aN1IWbrGXXv1JcHHNLInjb6Nw8PouFNfur9nqCXvKyg7jgsc5Md+k4wsqIAwvoRDz5YzVFGSvs5qiZhP8rSnopxUuG+q+66dk/kl6gpgKO+sxW3IwAJWsPuv+2HZFPTf9mVTd7a4XhH8FrQaWQBQJd4AdEgxYmsGCCliGQvA3y4yI8wbf31QBPpjocdlIRMkY1zMpBWiNNiEQitfoyfgyLFM46SbFoyOiuRamCGnF6cOmybz8f8F6MsMldNMfEQKBgQC/ItYZCawl8yDyWIziJTFWM0EcqekX3uuaKCnhY4a074FuIRdZjS9zXJRHaqSeiCE11N1qURao5yZOURnWNlx2yRrEEl2XeQiK0pFO4CKQ1NB3KEaIHetHXPcwRftADpMSDLlKRNQZatmOL3+bH1/gTLYyN8yHlZoB8zjinyLtlwKBgQCuaH47pdqAtDIW3feyVoMj1rXPT1+ByJP7jUfMQOxFhnNIkUxOgyn+iIaxA5jfsXaOdekZpuRko8lkgHc/fE6UiQmyWMZckB/2s9Ne6IDFKo7RfYaf9YgqEp4ir15qVZOfMaVdFxZXIddyeljnicTVU+VdXWWIQ99caqFS2a1d9wKBgFYOVjzxqOtxOv1CNzQ+sKbx7rf8HDGeMY+305tiFy1xxGGUTpIISjvzi+NXtYSXH/S/wWzz03L6l7mdNDfJQ/pLc1yiNDdVzC3MvpW3vnhdCPsNTCxO5Da+OaEncQcXSMpQpkA1GxXSkN8+rYJnLuCrSUD09IZ7KdsE6jDGm5BDAoGAf9AcEmJtSrzDqHZqu8sit/T1sEOe3mG78TGMWGeLvzfU7G1u2mrNL7el3buhIO3Q0H4goafo4MgXIwvyWBglDj2PWaAtXlBQ3F7UnE0PcW7K00OGkVtCunKCaTDTpGGqmZOseBgfSyOF3GNFHAmesieVg2Kasc3rtrJ+H1Ve8+kCgYBhwC3FGR3ZKUoOFCuxU+GieE3Sdjp9FF2GLDuAGeOHjt1Yb57eZLAPolWSE01TI3iL1tkSbHyL6G+54t8NBgE33rRu21ZYoj/zaweOOksDZOETKrlI+nZvenR50oQokS2TqrP1rmSBl/iBW4iuEDhLntMnkijbk/lpV3QPW65hWg=="
	kServerDomain        = "http://192.168.20.23:8080"
	AppPublicCertPath    = "config/crt/appPublicCert.crt"    // app公钥证书路径
	AliPayRootCertPath   = "config/crt/alipayRootCert.crt"   // alipay根证书路径
	AliPayPublicCertPath = "config/crt/alipayPublicCert.crt" // alipay公钥证书路径
	NotifyURL            = kServerDomain + "/notify"
	ReturnURL            = kServerDomain + "/callback"
	IsProduction         = false
)

func InitConf() {
	//CoreConf()
	InitConfig()
	UntilMysql()
	tools.InitRoles()
	UntilRedis()
	UntilMongoDB()
	AliPayClient = tools.Init(tools.Config{
		KAppID:               kAppID,
		KPrivateKey:          kPrivateKey,
		IsProduction:         IsProduction,
		AppPublicCertPath:    AppPublicCertPath,
		AliPayRootCertPath:   AliPayRootCertPath,
		AliPayPublicCertPath: AliPayPublicCertPath,
		NotifyURL:            NotifyURL,
		ReturnURL:            ReturnURL,
	})
	//UntilES()
}
func CloseUntil() {
	//ESClose(ConnectES)
	MongoDBClose(ConnectMongoDB)
	RedisClose(ConnectRedis)
	MysqlClose(ConnectMysql)
}
func CoreConf() {
	fmt.Println("配置初始化开始")
	const ConfigFile = "setting.yaml"
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		fmt.Println(1)
		panic(fmt.Errorf("error:%s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		fmt.Println(2)
		log.Fatalf("error:%v", err)
	}
	config.NewConfig = c
	fmt.Println("配置初始化完成")
}
func InitConfig() {
	v := viper.New()
	v.AddConfigPath(".") // 添加配置文件搜索路径，点号为当前目录
	//v.AddConfigPath("./configs") // 添加多个搜索目录
	v.SetConfigType("yaml")    // 如果配置文件没有后缀，可以不用配置
	v.SetConfigName("setting") // 文件名，没有后缀
	// v.SetConfigFile("configs/app.yml")
	// 读取配置文件
	if err := v.ReadInConfig(); err == nil {
		log.Printf("use config file -> %s\n", v.ConfigFileUsed())
	} else {
		return
	}
	a := config.Config{
		Mysql: config.Mysql{
			Host:     v.GetString("mysql.host"),
			Port:     v.GetString("mysql.port"),
			Db:       v.GetString("mysql.db"),
			User:     v.GetString("mysql.user"),
			Password: v.GetString("mysql.password"),
			LogLevel: v.GetString("mysql.log_level"),
		},
		Redis: config.Redis{
			Host:     v.GetString("Redis.host"),
			Port:     v.GetString("Redis.port"),
			Password: v.GetString("Redis.Password"),
			DB:       v.GetString("Redis.DB"),
		},
		MongoDB: config.MongoDB{
			Host:     v.GetString("MongoDB.host"),
			Port:     v.GetString("MongoDB.port"),
			User:     v.GetString("MongoDB.user"),
			Password: v.GetString("MongoDB.password"),
		},
		ES: config.ES{
			Host:     v.GetString("ES.host"),
			Port:     v.GetString("ES.port"),
			User:     v.GetString("ES.user"),
			Password: v.GetString("ES.password"),
		},
	}
	config.NewConfig = &a
	fmt.Println(a)
}
func UntilMysql() {
	db, err := gorm.Open(mysql.Open(config.NewConfig.Mysql.Dsn()), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	ConnectMysql = db
	fmt.Println("连接成功")
}
func UntilRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.NewConfig.Redis.Addr(),   // Redis服务器地址和端口
		Password: config.NewConfig.Redis.Password, // Redis密码，如果没有设置密码则为空
		DB:       config.NewConfig.Redis.Num(),    // Redis数据库索引，默认为0
	})
	ConnectRedis = rdb

	store, _ = redisstore.NewRedisStore(context.TODO(), ConnectRedis)
}
func UntilMongoDB() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.NewConfig.MongoDB.MyUri()))
	if err != nil {
		panic(err)
	}
	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
	ConnectMongoDB = client
}
func UntilES() {
	client, err := elastic.NewClient(
		elastic.SetURL(config.NewConfig.ES.MyESUrl()),
		elastic.SetSniff(false),
		elastic.SetBasicAuth(config.NewConfig.ES.MyESUserName(), config.NewConfig.ES.MyESUserPassword()),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	ConnectES = client
}
func ESClose(edb *elastic.Client) {
}
func MongoDBClose(mdb *mongo.Client) {
	err := mdb.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
}
func RedisClose(rdb *redis.Client) {
	err := rdb.Close()
	if err != nil {
		panic(err)
	}
}
func MysqlClose(db *gorm.DB) {
	DB, err := db.DB()
	if err != nil {
		panic(err)
	}
	err1 := DB.Close()
	if err1 != nil {
		panic(err1)
	}
}
func UntilRedis1() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.20.21:6379", // Redis服务器地址和端口
		Password: "root",               // Redis密码，如果没有设置密码则为空
		DB:       0,                    // Redis数据库索引，默认为0
	})
	return rdb
}
func UntilMysql1() *gorm.DB {
	db, _ := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",
		"root", "root", "192.168.20.21", 3306, "vote", "10s")), &gorm.Config{})
	return db
}
