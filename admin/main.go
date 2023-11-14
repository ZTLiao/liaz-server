package main

import (
	_ "admin/router"
	"core/cmd"
)

// @title 管理后台接口文档
// @version 1.0.0
// @contact.name liaozetao
// @contact.email 1107136310@qq.com
// @securityDefinitions.apikey  ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cmd.Execute("liaz-admin")
}
