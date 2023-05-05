package main

func main() {
}

/*--------打印 开始--------*/
// %v	按值的本来值输出
// %+v	在 %v 基础上，对结构体字段名和值进行展开
// %#v	输出 Go 语言语法格式的值
// %T	输出 Go 语言语法格式的类型和值
// %%	输出 % 本体
// %b	整型以二进制方式显示
// %o	整型以八进制方式显示
// %d	整型以十进制方式显示
// %x	整型以十六进制方式显示
// %X	整型以十六进制、字母大写方式显示
// %U	Unicode 字符
// %f	浮点数
// %p	指针，十六进制方式显示
/* type SearchApiParams struct {
	OrderKey string `json:"orderKey"` // 排序
	Desc     bool   `json:"desc"`     // 排序方式:升序false(默认)|降序true
}
var pageInfo SearchApiParams
fmt.Printf("%#v\n", pageInfo) */
/*--------打印 结束--------*/

/*--------gin框架 开始--------*/
//"github.com/gin-gonic/gin"
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
/*--------gin框架 结束--------*/

/*--------验证器 开始--------*/
// //"github.com/go-playground/validator/v10"
// validate := validator.New()
// var err error
// err = validate.Struct(xxStruct)
// err = validate.Var(map[string]string{"aaaa": "aaaa", "bbbb": "", "": "cccc"}, "required,dive,keys,required,endkeys,required")
// var errs map[string]error
// errs = validate.ValidateMap(map[string]interface{}{"aaaa": "aaaa", "bbbb": "", "cccc": ""}, map[string]interface{}{"aaaa": "required", "bbbb": "required,gt=10", "cccc": "required"})
/*--------验证器 结束--------*/

/*--------时间相关 开始--------*/
// cstZone, _ := time.LoadLocation("Asia/Shanghai")	//设置时区
// time.Local = cstZone
// //2006-01-02 15:04:05相当于php的y-m-d H:i:s
// st, err := time.Parse("2006-01-02 15:04:05", "2023-01-01 00:00:00")
// st, err := time.ParseInLocation("2006-01-02 15:04:05", "2023-01-01 00:00:00", time.Local)
// 时间戳 := st.Unix()
/*--------时间相关 开始--------*/

/*--------json 开始--------*/
// var rawData map[string]interface{}{
// 	"a":"a"
// }
// rawDataJson, _ := json.Marshal(rawData)
// var orgData map[string]interface{}
// json.Unmarshal(rawDataJson, &orgData)
/*--------json 开始--------*/

/*--------gorm 开始--------*/
/*
var info map[string]interface{}
db.Table("table").Where("id", id).Take(&info)

var list []map[string]interface{}
db.Table("table").Joins("left join table1 on table1.user_id=table.id").Find(&list)

var list []interface{}
db.Table("table").Where("id", v).Pluck("id", &list)

var sum int
db.Model(&users).Where("id", 1).Pluck("SUM(price) as sum", &sum)

db.Table("table").Create(map[string]interface{}{"id":id})

db.Table("table").Updates(map[string]interface{}{"price": gorm.Expr("price + ?", 1)})

db.Where("id", id).Delete(&game.TabPromoteSettlement{})

db.RowsAffected

errors.Is(err, gorm.ErrRecordNotFound)	//Find方法不会报这个错
*/
/*--------gorm 开始--------*/

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

	//会使关联表字符串类型的字段取出来的值类型为[]btye，根本不能给前端使用。用Table("tab_game_server")替换Model(&game.TabGameServer{})可解决
	global.MustGetGlobalDBByDBName(dbName).Model(&game.TabGameServer{}).Joins("left join tab_game on tab_game.id = tab_game_server.game_id")

*/
