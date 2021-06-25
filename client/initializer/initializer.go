package initializer

import "github.com/jjonline/golang-backend/client"

// region 全局句柄初始化相关

// Init 初始化
func Init() {
	client.Logger = iniLogger() // 初始化logger，需要优先执行
	client.Redis = initRedis()  // 初始化redis
	client.DB = initDB()        // 初始化db，内部操作不允许重新赋值，除非清楚您知道现在找干什么！！！
}

// endregion

// region 全局句柄热重载后初始化相关

// Reload 热重载再次初始化调用
// 当配置变更时又无需重新启动进程触发，监听调用
func Reload() {

}

// endregion