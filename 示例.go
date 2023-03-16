package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type SearchApiParams struct {
	OrderKey string `json:"orderKey"` // 排序
	Desc     bool   `json:"desc"`     // 排序方式:升序false(默认)|降序true
}

func main() {
	fmt.Println("Hello, World!")
	//打印结构体
	var pageInfo SearchApiParams
	fmt.Printf("%#v", pageInfo)

	/* gin框架使用方法 */
	c := *gin.Context
	//path参数获取（/user/:page/*action"）
	page := c.Param("page")
	action := c.Param("action")
	//get参数获取
	page := c.Query("page")
	//post参数获取，对应application/x-www-form-urlencoded和from-data格式参数
	page := c.PostForm("page")
	//page := c.PostFormMap("page")
	//post参数获取（Content-Type: application/json）
	// var pageInfo systemReq.SearchApiParams
	// err := c.ShouldBindJSON(&pageInfo)
}
