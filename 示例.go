package main

import "fmt"

func main() {
   fmt.Println("Hello, World!")
}

//打印结构体
fmt.Printf("%#v", pageInfo)

//get参数获取
page := c.QueryMap("page")
//post参数获取（Content-Type: application/x-www-form-urlencoded）
pageSize := c.PostFormMap("pageSize")
//post参数获取（Content-Type: application/json）
// var pageInfo systemReq.SearchApiParams
// err := c.ShouldBindJSON(&pageInfo)