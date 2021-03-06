package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func main() {

	fmt.Printf("launch Gin")

	r := gin.Default()
	r.GET("/get",HandleGet)
	r.POST("/getall",HandleGetAllData)

	//如果使用浏览器调试，那么响应Get方法
	//r.GET("/getall",HandleGetAllData)
	r.Run(":8080")

}

func HandleGet(c *gin.Context)  {
	c.JSON(200,gin.H{
		"receive":"65536",
	})

}

func HandleGetAllData(c *gin.Context)  {
	//log.Print("handle log")
	body,_ := ioutil.ReadAll(c.Request.Body)
	fmt.Println("---body/--- \r\n "+string(body))

	fmt.Println("---header/--- \r\n")
	for k,v :=range c.Request.Header {
		fmt.Println(k,v)
	}
	//fmt.Println("header \r\n",c.Request.Header)

	c.JSON(200,gin.H{
		"receive":"1024",
	})

}
