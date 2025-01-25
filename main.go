package main

import (
	"github.com/hcd233/Aris-url-gen/cmd"
)

// @title Aris URL Generator API
// @version 1.0
// @description 短链接生成服务API文档
// @termsOfService http://swagger.io/terms/

// @contact.name centonhuang
// @contact.email lvlvko233@qq.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cmd.Execute()
}
