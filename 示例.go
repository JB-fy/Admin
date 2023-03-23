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

/*
go开发流程
	自动化package->新增
	代码生成器->新增
	断开服务
	如果数据库中表已存在，删除server/initialize/gorm.go该表的自动创建，否则将导致原表被修改
	如果数据库中表已存在，且没有created_by updated_by deleted_by字段，则将server/model/user/tab_user.go中
		global.GVA_MODEL
		替换成
		ID        uint           `gorm:"primarykey"` // 主键ID
	重启服务
*/
