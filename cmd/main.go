package main

import (
	"BookStore/common/db"
	"BookStore/pkg/books"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./common/envs/.env")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()
	h := db.Init(dbUrl)

	books.RegisterRoutes(r, h)
	// register more routes here

	r.Run(port)
}

//viper.SetConfigFile("./common/envs/.env")

////viper.ReadInConfig()
//
//port := viper.Get("PORT").(string)
//dbUrl := viper.Get("DB_URL").(string)
//
//r := gin.Default()
//db.Init(dbUrl)
//
//r.GET("/", func(c *gin.Context) {
//	c.JSON(200, gin.H{
//		"port":  port,
//		"dbUrl": dbUrl,
//	})
//})
//
//r.Run(port)
//}
