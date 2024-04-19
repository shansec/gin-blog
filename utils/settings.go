package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	Model    string
	HttpPort string
	JwtKey   string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	AccessKey string
	SecretKey string
	Bucket    string
	QiNServe  string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Print("配置读取出错", err)
	}
	loadServe(file)
	loadData(file)
	loadImg(file)
}

func loadServe(file *ini.File) {
	Model = file.Section("serve").Key("Model").MustString("debug")
	HttpPort = file.Section("serve").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("serve").Key("JwtKey").MustString("a1s458d4s5")
}

func loadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("debug")
	DbHost = file.Section("database").Key("DbHost").MustString("127.0.0.1")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("root")
	DbName = file.Section("database").Key("DbName").MustString("myginblog")
}

func loadImg(file *ini.File) {
	AccessKey = file.Section("qiNiu").Key("AccessKey").MustString("-zgPgP_kLxQVl6DsUrf1_diSu8hJcI3Nsd_1Xy14")
	SecretKey = file.Section("qiNiu").Key("SecretKey").MustString("SubcC2c6NWVxCsT5fvv5NDdhFG8wul2njX0ssQ-e")
	Bucket = file.Section("qiNiu").Key("Bucket").MustString("myginblogimg")
	QiNServe = file.Section("qiNiu").Key("QiNServe").MustString("http://r8tivcxrp.hd-bkt.clouddn.com")
}
