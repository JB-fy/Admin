package main

func main() {
}

// type SearchApiParams struct {
// 	OrderKey string `json:"orderKey"` // 排序
// 	Desc     bool   `json:"desc"`     // 排序方式:升序false(默认)|降序true
// }
// /* 打印结构体
// %v占位符是不会打印结构体字段名称的，字段之间以空格隔开；
// %+v占位符会打印字段名称，字段之间也是以空格隔开；
// %#v占位符则会打印结构体类型和字段名称，字段之间以逗号分隔 */
// var pageInfo SearchApiParams
// fmt.Printf("%#v\n", pageInfo)
// /*--------gin框架 开始--------*/
// //"github.com/gin-gonic/gin"
// c := *gin.Context
// //path参数获取（/user/:page/*action"）
// page := c.Param("page")
// action := c.Param("action")
// //get参数获取
// page := c.Query("page")
// //post参数获取，对应application/x-www-form-urlencoded和from-data格式参数
// page := c.PostForm("page")
// //page := c.PostFormMap("page")
// //post参数获取（Content-Type: application/json）
// // var pageInfo systemReq.SearchApiParams
// // err := c.ShouldBindJSON(&pageInfo)
// /*--------gin框架 结束--------*/

// /*--------验证器 开始--------*/
// //"github.com/go-playground/validator/v10"
// validate := validator.New()
// var err error
// err = validate.Struct(xxStruct)
// err = validate.Var(map[string]string{"aaaa": "aaaa", "bbbb": "", "": "cccc"}, "required,dive,keys,required,endkeys,required")
// var errs map[string]error
// errs = validate.ValidateMap(map[string]interface{}{"aaaa": "aaaa", "bbbb": "", "cccc": ""}, map[string]interface{}{"aaaa": "required", "bbbb": "required,gt=10", "cccc": "required"})
// /*--------验证器 结束--------*/

// /*--------时间相关 开始--------*/
// //2006-01-02 15:04:05相当于php的y-m-d H:i:s
// st, err := time.Parse("2006-01-02 15:04:05", "2023-01-01 00:00:00")
// 时间戳 := st.Unix()
// /*--------时间相关 开始--------*/

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


注意事项
	使用结构体做创建和更新操作时，必须把设置全部数据库字段的值，或使用Select()或Omit()方法指定或排除某些字段。
		如果字段没传时，创建时会插入null，导致数据库报错，除非在结构体中设置默认值gorm:"default:0;"，但是这也会导致更新操作时，未设置值的字段会被更新成默认值。
		所以强烈不建议用数据库模型结构体接收前端参数，做数据库创建和更新
	使用map方式做创建和更新操作时，字段名必须与数据库字段名一致

	//不建议用结构体直接插入，会有默认值问题（Column 'xxxx' cannot be null）
		err = global.MustGetGlobalDBByDBName(dbName).Create(&tabGameServer).Error
	//多行数据同时插入。每行数据字段都一样时使用，否则差异字段会被插入null，导致数据库报错
		err = global.MustGetGlobalDBByDBName(dbName).Model(&game.TabGameServer{}).Create(data).Error
	//多行数据循环插入。每行数据字段不一样时使用
		for _, one := range data {
			err = global.MustGetGlobalDBByDBName(dbName).Model(&game.TabGameServer{}).Create(one).Error
		}


*/
