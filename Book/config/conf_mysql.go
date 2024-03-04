package config

import (
	"fmt"
	"strconv"
)

type Mysql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Db       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log_level"`
}

func (m Mysql) Dsn() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Db + "?charset=utf8&parseTime=True&loc=Local&timeout=10s"
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"Password"`
	DB       string `yaml:"DB"`
}

func (r Redis) Addr() string {
	return r.Host + ":" + r.Port
}
func (r Redis) Num() int {
	id, _ := strconv.Atoi(r.DB)
	return id
}

type MongoDB struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func (d MongoDB) MyUri() (myuri string) {
	myuri = fmt.Sprintf("mongodb://%s:%s@%s:%s", d.User, d.Password, d.Host, d.Port)
	return myuri
}

type ES struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func (e ES) MyESUrl() (res string) {
	res = fmt.Sprintf("http://%s:%s", e.Host, e.Port)
	return res
}
func (e ES) MyESUserName() (res string) {
	res = fmt.Sprintf("%s", e.User)
	return res
}
func (e ES) MyESUserPassword() (res string) {
	res = fmt.Sprintf("%s", e.Password)
	return res
}
