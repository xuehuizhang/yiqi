package initialize

import "com.yiqi/dao/db"

func InitDb() {
	//db.SetUp("mysql", "root", "Zhangxh106!", "127.0.0.1", 61166, "yiqi")
	db.SetUp("mysql", "root", "123456", "127.0.0.1", 3306, "yiqi")

}
