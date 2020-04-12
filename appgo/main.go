package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"samtake/justdoit/common"
	"samtake/justdoit/router"

	_ "github.com/go-sql-driver/mysql"

	"github.com/spf13/viper"
)

func main() {
	InitConfig()
	db := common.InitDB()
	defer db.Close()

	r := gin.New()
	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r = router.InitRouter(r)

	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}else{
		panic(r.Run(":8099"))
	}	
}

//InitConfig 初始化配置文件.
func InitConfig() {
	workDir, _ := os.Getwd()                 
	viper.SetConfigName("application")       
	viper.SetConfigType("yml")               
	viper.AddConfigPath(workDir + "/config") 
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
