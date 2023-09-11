package conf

import (
	"github.com/go-ini/ini"
)

var (
	AppModel string
	HttpPort string

	DB         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string

	RedisDb     string
	RedisAddr   string
	RedisPw     string
	RedisDbName string

	Host        string
	ProductPath string
	AvatarPath  string
)

func init() {
	//读取config
	file, err := ini.Load("./conf")
	if err != nil {
		panic(err)
	}
	LoadServer(file)
	LoadMysql(file)
	LoadRedis(file)
	LoadEmail(file)
	LoadPhotoPath(file)
}
func LoadServer(file *ini.File) {
	AppModel = file.Section("service").Key("AppMode").String()
	AppModel = file.Section("service").Key("HttpPort").String()
}
func LoadMysql(file *ini.File) {
	AppModel = file.Section("mysql").Key("DB").String()
	AppModel = file.Section("mysql").Key("DbHost").String()
	AppModel = file.Section("mysql").Key("DbPort").String()
	AppModel = file.Section("mysql").Key("DbUser").String()
	AppModel = file.Section("mysql").Key("DbPassword").String()
	AppModel = file.Section("mysql").Key("DbName").String()
}
func LoadRedis(file *ini.File) {
	AppModel = file.Section("redis").Key("AppMode").String()

}
func LoadEmail(file *ini.File) {
	AppModel = file.Section("email").Key("AppMode").String()

}
func LoadPhotoPath(file *ini.File) {
	AppModel = file.Section("path").Key("AppMode").String()

}
