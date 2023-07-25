package cmd

import (
	daoAuth "api/internal/dao/auth"
	"context"
	"fmt"
	"os/exec"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

/*
使用示例：./main myGen -sceneCode=platform -dbGroup=default -dbTable=auth_test -removePrefix=auth_ -moduleDir=auth -commonName=测试 -isList=yes -isCreate=yes -isUpdate=yes -isDelete=yes -isApi=yes -isAuthAction=yes -isView=yes -isCover=no

强烈建议搭配Git使用

表字段命名需要遵守以下规则，否则会根据字段类型默认处理
主键必须是第一个字段。否则需要在dao层重写PrimaryKey方法返回主键字段
表内尽量根据表名设置xxxxId和xxxxName两个字段(这两字段，常用于前端组件)

	部分常用字段：
		password	密码
		passwd		密码
		pid			父级（指向本表）
		sort		排序
		weight 		权重
		gender 		性别
		avatar		头像
	其他类型字段：
		名称和标识字段，命名用name或code后缀
		手机号码字段，命名用mobile或phone后缀
		链接地址字段，命名用url或link后缀
		关联id字段和关联表主键保持一致，命名用id后缀
		图片字段，命名用icon,cover或img,img_list,imgList,img_arr,imgArr或image,image_list,imageList,image_arr,imageArr等后缀（多图片时字段类型用json或text，保存格式为JSON格式）
		视频字段，命名用video,video_list,videoList,video_arr,videoArr等后缀（多视频时字段类型用json或text，保存格式为JSON格式）
		ip字段，命名用Ip后缀
		备注字段，命名用remark后缀
		状态和类型字段，命名用status或type后缀
		是否字段，命名用is_前缀
*/
type MyGenOption struct {
	SceneCode    string `c:"sceneCode"`    //场景标识。示例：platform
	DbGroup      string `c:"dbGroup"`      //db分组。示例：default
	DbTable      string `c:"dbTable"`      //db表。示例：auth_scene
	RemovePrefix string `c:"removePrefix"` //要删除的db表前缀。示例：auth_
	ModuleDir    string `c:"moduleDir"`    //模块目录，只支持单目录。必须和hcak/config.yaml内daoPath的后面部分保持一致，示例：auth
	CommonName   string `c:"commonName"`   //公共名称，将同时在swagger文档Tag标签名称，菜单名称和操作名称中使用。示例：场景
	IsList       bool   `c:"isList" `      //是否生成列表接口(0,false,off,no,""为false，其他都为true)
	IsCreate     bool   `c:"isCreate"`     //是否生成创建接口
	IsUpdate     bool   `c:"isUpdate"`     //是否生成更新接口
	IsDelete     bool   `c:"isDelete"`     //是否生成删除接口
	IsApi        bool   `c:"isApi"`        //是否生成后端接口文件
	IsAuthAction bool   `c:"isAuthAction"` //是否判断操作权限，如是，则同时会生成操作权限
	IsView       bool   `c:"isView"`       //是否生成前端视图文件
	IsCover      bool   `c:"isCover"`      //是否覆盖原文件(设置为true时，建议与git一起使用，防止代码覆盖风险)
}

type MyGenTpl struct {
	TableColumnList            gdb.Result //表字段详情
	PrimaryKey                 string     //表主键
	LabelField                 string     //dao层label对应的字段(常用于前端组件)
	SceneName                  string     //场景名称
	SceneId                    int        //场景ID
	RawTableNameCaseCamelLower string     //原始表名（小驼峰）
	TableNameCaseCamelLower    string     //去除前缀表名（小驼峰）
	TableNameCaseCamel         string     //去除前缀表名（大驼峰）
	TableNameCaseSnake         string     //去除前缀表名（蛇形）
	ModuleDirCaseCamelLower    string     //路径后缀（小驼峰）
	ModuleDirCaseCamel         string     //路径后缀（大驼峰）
}

func MyGenFunc(ctx context.Context, parser *gcmd.Parser) (err error) {
	option := MyGenOptionHandle(ctx, parser)
	tpl := MyGenTplHandle(ctx, option)

	MyGenTplDao(ctx, option, tpl)                         // dao层存在时，增加或修改部分字段的解析代码
	MyGenTplLogic(ctx, option, tpl)                       // logic模板生成（文件不存在时增删改查全部生成，已存在不处理不覆盖）
	exec.Command(`gf`, `gen`, `service`).CombinedOutput() // service生成

	if option.IsApi {
		MyGenTplApi(ctx, option, tpl)        // api模板生成
		MyGenTplController(ctx, option, tpl) // controller模板生成
		MyGenTplRouter(ctx, option, tpl)     // 后端路由生成
	}

	if option.IsView {
		MyGenTplViewIndex(ctx, option, tpl)  // 视图模板Index生成
		MyGenTplViewList(ctx, option, tpl)   // 视图模板List生成
		MyGenTplViewQuery(ctx, option, tpl)  // 视图模板Query生成
		MyGenTplViewSave(ctx, option, tpl)   // 视图模板Save生成
		MyGenTplViewI18n(ctx, option, tpl)   // 视图模板I18n生成
		MyGenTplViewRouter(ctx, option, tpl) // 前端路由生成
	}
	return
}

// 参数处理
func MyGenOptionHandle(ctx context.Context, parser *gcmd.Parser) (option *MyGenOption) {
	optionMap := parser.GetOptAll()
	option = &MyGenOption{}
	gconv.Struct(optionMap, option)

	// 场景标识
	_, ok := optionMap[`sceneCode`]
	if !ok {
		option.SceneCode = gcmd.Scan("> 请输入场景标识:\n")
	}
	for {
		if option.SceneCode != `` {
			count, _ := daoAuth.Scene.ParseDbCtx(ctx).Where(daoAuth.Scene.Columns().SceneCode, option.SceneCode).Count()
			if count > 0 {
				break
			}
		}
		option.SceneCode = gcmd.Scan("> 场景标识不存在，请重新输入:\n")
	}
	// db分组
	var db gdb.DB
	if option.DbGroup == `` {
		option.DbGroup = gcmd.Scan("> 请输入db分组，默认(default):\n")
		if option.DbGroup == `` {
			option.DbGroup = `default`
		}
	}
	for {
		err := g.Try(ctx, func(ctx context.Context) {
			db = g.DB(option.DbGroup)
		})
		if err == nil {
			break
		}
		option.DbGroup = gcmd.Scan("> db分组不存在，请重新输入，默认(default):\n")
		if option.DbGroup == `` {
			option.DbGroup = `default`
		}
	}
	// db表
	tableArrTmp, _ := db.GetArray(ctx, `SHOW TABLES`)
	tableArr := gconv.SliceStr(tableArrTmp)
	if option.DbTable == `` {
		option.DbTable = gcmd.Scan("> 请输入db表:\n")
	}
	for {
		if option.DbTable != `` && garray.NewStrArrayFrom(tableArr).Contains(option.DbTable) {
			break
		}
		option.DbTable = gcmd.Scan("> db表不存在，请重新输入:\n")
	}
	// db表前缀
	_, ok = optionMap[`removePrefix`]
	if !ok {
		option.RemovePrefix = gcmd.Scan("> 请输入要删除的db表前缀，默认(空):\n")
	}
	for {
		if option.RemovePrefix == `` || gstr.Pos(option.DbTable, option.RemovePrefix) == 0 {
			break
		}
		option.RemovePrefix = gcmd.Scan("> 要删除的db表前缀不存在，请重新输入，默认(空):\n")
	}
	// 模块目录
	_, ok = optionMap[`moduleDir`]
	if !ok {
		option.ModuleDir = gcmd.Scan("> 请输入模块目录:\n")
	}
	for {
		if option.ModuleDir != `` {
			break
		}
		option.ModuleDir = gcmd.Scan("> 请输入模块目录:\n")
	}
	// 公共名称，将同时在swagger文档Tag标签名称，菜单名称和操作名称中使用。示例：场景
	_, ok = optionMap[`commonName`]
	if !ok {
		option.CommonName = gcmd.Scan("> 请输入公共名称，将同时在swagger文档Tag标签名称，菜单名称和操作名称中使用:\n")
	}
	for {
		if option.CommonName != `` {
			break
		}
		option.CommonName = gcmd.Scan("> 请输入公共名称，将同时在swagger文档Tag标签名称，菜单名称和操作名称中使用:\n")
	}
	// 是否生成列表接口
noAllRestart:
	isList, ok := optionMap[`isList`]
	if !ok {
		isList = gcmd.Scan("> 是否生成列表接口，默认(yes):\n")
	}
isListEnd:
	for {
		switch isList {
		case ``, `yes`:
			option.IsList = true
			break isListEnd
		case `no`:
			option.IsList = false
			break isListEnd
		default:
			isList = gcmd.Scan("> 输入错误，请重新输入，是否生成列表接口，默认(yes):\n")
		}
	}
	// 是否生成创建接口
	isCreate, ok := optionMap[`isCreate`]
	if !ok {
		isCreate = gcmd.Scan("> 是否生成创建接口，默认(yes):\n")
	}
isCreateEnd:
	for {
		switch isCreate {
		case ``, `yes`:
			option.IsCreate = true
			break isCreateEnd
		case `no`:
			option.IsCreate = false
			break isCreateEnd
		default:
			isCreate = gcmd.Scan("> 输入错误，请重新输入，是否生成创建接口，默认(yes):\n")
		}
	}
	// 是否生成更新接口
	isUpdate, ok := optionMap[`isUpdate`]
	if !ok {
		isUpdate = gcmd.Scan("> 是否生成更新接口，默认(yes):\n")
	}
isUpdateEnd:
	for {
		switch isUpdate {
		case ``, `yes`:
			option.IsUpdate = true
			break isUpdateEnd
		case `no`:
			option.IsUpdate = false
			break isUpdateEnd
		default:
			isUpdate = gcmd.Scan("> 输入错误，请重新输入，是否生成更新接口，默认(yes):\n")
		}
	}
	// 是否生成删除接口
	isDelete, ok := optionMap[`isDelete`]
	if !ok {
		isDelete = gcmd.Scan("> 是否生成删除接口，默认(yes):\n")
	}
isDeleteEnd:
	for {
		switch isDelete {
		case ``, `yes`:
			option.IsDelete = true
			break isDeleteEnd
		case `no`:
			option.IsDelete = false
			break isDeleteEnd
		default:
			isDelete = gcmd.Scan("> 输入错误，请重新输入，是否生成删除接口，默认(yes):\n")
		}
	}
	if !(option.IsList || option.IsCreate || option.IsUpdate || option.IsDelete) {
		fmt.Println("请重新选择生成哪些接口，不能全是no！")
		goto noAllRestart
	}
	// 是否生成后端接口文件
	isApi, ok := optionMap[`isApi`]
	if !ok {
		isApi = gcmd.Scan("> 是否生成后端接口文件，默认(yes):\n")
	}
isApiEnd:
	for {
		switch isApi {
		case ``, `yes`:
			option.IsApi = true
			break isApiEnd
		case `no`:
			option.IsApi = false
			break isApiEnd
		default:
			isApi = gcmd.Scan("> 输入错误，请重新输入，是否生成后端接口文件，默认(yes):\n")
		}
	}
	if option.IsApi {
		// 是否判断操作权限，如是，则同时会生成操作权限
		isAuthAction, ok := optionMap[`isAuthAction`]
		if !ok {
			isAuthAction = gcmd.Scan("> 是否判断操作权限，如是，则同时会生成操作权限，默认(yes):\n")
		}
	isAuthActionEnd:
		for {
			switch isAuthAction {
			case ``, `yes`:
				option.IsAuthAction = true
				break isAuthActionEnd
			case `no`:
				option.IsAuthAction = false
				break isAuthActionEnd
			default:
				isAuthAction = gcmd.Scan("> 输入错误，请重新输入，是否判断操作权限，如是，则同时会生成操作权限，默认(yes):\n")
			}
		}
	}
	// 是否生成前端视图文件
	isView, ok := optionMap[`isView`]
	if !ok {
		isView = gcmd.Scan("> 是否生成前端视图文件，默认(yes):\n")
	}
isViewEnd:
	for {
		switch isView {
		case ``, `yes`:
			option.IsView = true
			break isViewEnd
		case `no`:
			option.IsView = false
			break isViewEnd
		default:
			isView = gcmd.Scan("> 输入错误，请重新输入，是否生成前端视图文件，默认(yes):\n")
		}
	}
	// 是否覆盖原文件
	isCover, ok := optionMap[`isCover`]
	if !ok {
		isCover = gcmd.Scan("> 是否覆盖原文件(设置为yes时，建议与git一起使用，防止代码覆盖风险)，默认(no):\n")
	}
isCoverEnd:
	for {
		switch isCover {
		case `yes`:
			option.IsCover = true
			break isCoverEnd
		case ``, `no`:
			option.IsCover = false
			break isCoverEnd
		default:
			isCover = gcmd.Scan("> 输入错误，请重新输入，是否覆盖原文件(设置为yes时，建议与git一起使用，防止代码覆盖风险)，默认(no):\n")
		}
	}
	return
}

// 模板参数处理
func MyGenTplHandle(ctx context.Context, option *MyGenOption) (tpl *MyGenTpl) {
	tableColumnList, _ := g.DB(option.DbGroup).GetAll(ctx, `SHOW FULL COLUMNS FROM `+option.DbTable)
	sceneInfo, _ := daoAuth.Scene.ParseDbCtx(ctx).Where(daoAuth.Scene.Columns().SceneCode, option.SceneCode).One()
	tableName := gstr.Replace(option.DbTable, option.RemovePrefix, ``, 1)
	tpl = &MyGenTpl{
		TableColumnList:            tableColumnList,
		SceneName:                  sceneInfo[daoAuth.Scene.Columns().SceneName].String(),
		SceneId:                    sceneInfo[daoAuth.Scene.Columns().SceneId].Int(),
		RawTableNameCaseCamelLower: gstr.CaseCamelLower(option.DbTable),
		TableNameCaseCamelLower:    gstr.CaseCamelLower(tableName),
		TableNameCaseCamel:         gstr.CaseCamel(tableName),
		TableNameCaseSnake:         gstr.CaseSnakeFirstUpper(tableName),
		ModuleDirCaseCamelLower:    gstr.CaseCamelLower(option.ModuleDir),
		ModuleDirCaseCamel:         gstr.CaseCamel(option.ModuleDir),
	}
	fieldArr := make([]string, len(tpl.TableColumnList))
	fieldCaseCamelArr := make([]string, len(tpl.TableColumnList))
	for index, column := range tpl.TableColumnList {
		field := column[`Field`].String()
		fieldArr[index] = field
		fieldCaseCamel := gstr.CaseCamel(field)
		fieldCaseCamelArr[index] = fieldCaseCamel
		if column[`Key`].String() == `PRI` && column[`Extra`].String() == `auto_increment` {
			tpl.PrimaryKey = field
		}
	}

	fieldCaseCamelArrG := garray.NewStrArrayFrom(fieldCaseCamelArr)
	// 根据name字段优先级排序
	nameFieldList := []string{
		tpl.TableNameCaseCamel + `Name`,
		gstr.SubStr(gstr.CaseCamel(tpl.PrimaryKey), 0, -2) + `Name`,
		`Phone`,
		`Account`,
		`nickname`,
	}
	for _, v := range nameFieldList {
		index := fieldCaseCamelArrG.Search(v)
		if index != -1 {
			tpl.LabelField = fieldArr[index]
			return
		}
	}
	return
}

// status字段注释解析
func MyGenStatusList(comment string) (statusList [][2]string) {
	tmp, _ := gregex.MatchAllString(`(\d+)([^\d\s,，;；]+)`, comment)
	statusList = make([][2]string, len(tmp))
	for k, v := range tmp {
		statusList[k] = [2]string{v[1], gstr.TrimLeft(gstr.TrimLeft(v[2], `:`), `：`)}
	}
	return
}

// dao层存在时，增加或修改部分字段的解析代码
func MyGenTplDao(ctx context.Context, option *MyGenOption, tpl *MyGenTpl) {
	saveFile := gfile.SelfDir() + `/internal/dao/` + tpl.ModuleDirCaseCamelLower + `/` + tpl.TableNameCaseSnake + `.go`
	if !gfile.IsFile(saveFile) {
		return
	}
	tplDao := gfile.GetContents(saveFile)

	if tpl.LabelField != `` {
		if gstr.Pos(tplDao, `case `+"`label`"+`:
				m = m.Fields(`) == -1 {
			tplDao = gstr.Replace(tplDao, `/*--------ParseField自动代码生成锚点（不允许修改和删除，否则将不能自动生成代码）--------*/`, `case `+"`label`"+`:
				m = m.Fields(daoThis.Table() + `+"`.`"+` + daoThis.Columns().`+gstr.CaseCamel(tpl.LabelField)+` + `+"` AS `"+` + v)
			/*--------ParseField自动代码生成锚点（不允许修改和删除，否则将不能自动生成代码）--------*/`)
		}
		if gstr.Pos(tplDao, `case `+"`label`"+`:
				m = m.WhereLike(`) == -1 {
			tplDao = gstr.Replace(tplDao, `/*--------ParseFilter自动代码生成锚点（不允许修改和删除，否则将不能自动生成代码）--------*/`, `case `+"`label`"+`:
				m = m.WhereLike(daoThis.Table()+`+"`.`"+`+daoThis.Columns().`+gstr.CaseCamel(tpl.LabelField)+`, `+"`%`"+`+gconv.String(v)+`+"`%`"+`)
			/*--------ParseFilter自动代码生成锚点（不允许修改和删除，否则将不能自动生成代码）--------*/`)
		}
	}

	imageVideoJsonFieldArr := []string{}
	for _, column := range tpl.TableColumnList {
		field := column[`Field`].String()
		fieldCaseCamel := gstr.CaseCamel(field)
		switch field {
		case `deletedAt`, `deleted_at`, `createdAt`, `created_at`, `updatedAt`, `updated_at`:
		case `password`, `passwd`:
		default:
			//icon,cover或img,img_list,imgList,img_arr,imgArr或image,image_list,imageList,image_arr,imageArr等后缀
			//video,video_list,videoList,video_arr,videoArr等后缀
			if field == `avatar` || gstr.SubStr(fieldCaseCamel, -4) == `Icon` || gstr.SubStr(fieldCaseCamel, -5) == `Cover` || gstr.SubStr(fieldCaseCamel, -3) == `Img` || gstr.SubStr(fieldCaseCamel, -7) == `ImgList` || gstr.SubStr(fieldCaseCamel, -6) == `ImgArr` || gstr.SubStr(fieldCaseCamel, -5) == `Image` || gstr.SubStr(fieldCaseCamel, -9) == `ImageList` || gstr.SubStr(fieldCaseCamel, -8) == `ImageArr` || gstr.SubStr(fieldCaseCamel, -5) == `Video` || gstr.SubStr(fieldCaseCamel, -9) == `VideoList` || gstr.SubStr(fieldCaseCamel, -8) == `VideoArr` {
				if column[`Type`].String() == `json` || column[`Type`].String() == `text` {
					imageVideoJsonFieldArr = append(imageVideoJsonFieldArr, `daoThis.Columns().`+gstr.CaseCamel(field))
				}
			}
		}
	}
	if len(imageVideoJsonFieldArr) > 0 {
		imageVideoJsonFieldStr := gstr.Join(imageVideoJsonFieldArr, `, `)
		if gstr.Pos(tplDao, `m = m.Fields(daoThis.Table() + `+"`.`"+` + v)
				afterField = append(afterField, v)`) == -1 {
			tplDao = gstr.Replace(tplDao, `/*--------ParseField自动代码生成锚点（不允许修改和删除，否则将不能自动生成代码）--------*/`, `case `+imageVideoJsonFieldStr+`:
				m = m.Fields(daoThis.Table() + `+"`.`"+` + v)
				afterField = append(afterField, v)
			/*--------ParseField自动代码生成锚点（不允许修改和删除，否则将不能自动生成代码）--------*/`)
		} else {
			tmp, _ := gregex.MatchString(`case ([^:]*):
				m = m\.Fields\(daoThis\.Table\(\) \+ `+"`"+`\.`+"`"+` \+ v\)
				afterField = append\(afterField, v\)`, tplDao)
			tplDao, _ = gregex.ReplaceString(`case [^:]*:
				m = m\.Fields\(daoThis\.Table\(\) \+ `+"`"+`\.`+"`"+` \+ v\)
				afterField = append\(afterField, v\)`, `case `+gstr.Join(gset.NewStrSetFrom(imageVideoJsonFieldArr).Union(gset.NewStrSetFrom(gstr.Split(tmp[1], `, `))).Slice(), `, `)+`:
				m = m.Fields(daoThis.Table() + `+"`.`"+` + v)
				afterField = append(afterField, v)`, tplDao)
		}

		if gstr.Pos(tplDao, `record[v] = gvar.New(record[v].Slice())`) == -1 {
			tplDao = gstr.Replace(tplDao, `/*--------HookSelect自动代码生成锚点（不允许修改和删除，否则将不能自动生成代码）--------*/`, `case `+imageVideoJsonFieldStr+`:
						record[v] = gvar.New(record[v].Slice())
					/*--------HookSelect自动代码生成锚点（不允许修改和删除，否则将不能自动生成代码）--------*/`)
		} else {
			tmp, _ := gregex.MatchString(`case ([^:]*):
						record\[v\] = gvar\.New\(record\[v\]\.Slice\(\)\)`, tplDao)
			tplDao, _ = gregex.ReplaceString(`case [^:]*:
						record\[v\] = gvar\.New\(record\[v\]\.Slice\(\)\)`, `case `+gstr.Join(gset.NewStrSetFrom(imageVideoJsonFieldArr).Union(gset.NewStrSetFrom(gstr.Split(tmp[1], `, `))).Slice(), `, `)+`:
						record[v] = gvar.New(record[v].Slice())`, tplDao)
		}
	}
	gfile.PutContents(saveFile, tplDao)
}

// logic模板生成（文件不存在时增删改查全部生成，已存在不处理不覆盖）
func MyGenTplLogic(ctx context.Context, option *MyGenOption, tpl *MyGenTpl) {
	saveFile := gfile.SelfDir() + `/internal/logic/` + tpl.ModuleDirCaseCamelLower + `/` + tpl.TableNameCaseSnake + `.go`
	if gfile.IsFile(saveFile) {
		return
	}

	tplLogic := `package logic

import (
	dao` + tpl.ModuleDirCaseCamel + ` "api/internal/dao/` + tpl.ModuleDirCaseCamelLower + `"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gconv"
)

type s` + tpl.TableNameCaseCamel + ` struct{}

func New` + tpl.TableNameCaseCamel + `() *s` + tpl.TableNameCaseCamel + ` {
	return &s` + tpl.TableNameCaseCamel + `{}
}

func init() {
	service.Register` + tpl.TableNameCaseCamel + `(New` + tpl.TableNameCaseCamel + `())
}

// 总数
func (logicThis *s` + tpl.TableNameCaseCamel + `) Count(ctx context.Context, filter map[string]interface{}) (count int, err error) {
	daoThis := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `
	joinTableArr := []string{}
	model := daoThis.ParseDbCtx(ctx)
	if len(filter) > 0 {
		model = model.Handler(daoThis.ParseFilter(filter, &joinTableArr))
	}
	if len(joinTableArr) > 0 {
		model = model.Group(daoThis.Table() + ` + "`.`" + ` + daoThis.PrimaryKey()).Distinct().Fields(daoThis.Table() + ` + "`.`" + ` + daoThis.PrimaryKey())
	}
	count, err = model.Count()
	return
}

// 列表
func (logicThis *s` + tpl.TableNameCaseCamel + `) List(ctx context.Context, filter map[string]interface{}, field []string, order []string, page int, limit int) (list gdb.Result, err error) {
	daoThis := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `
	joinTableArr := []string{}
	model := daoThis.ParseDbCtx(ctx)
	if len(filter) > 0 {
		model = model.Handler(daoThis.ParseFilter(filter, &joinTableArr))
	}
	if len(field) > 0 {
		model = model.Handler(daoThis.ParseField(field, &joinTableArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoThis.ParseOrder(order, &joinTableArr))
	}
	if len(joinTableArr) > 0 {
		model = model.Group(daoThis.Table() + ` + "`.`" + ` + daoThis.PrimaryKey())
	}
	if limit > 0 {
		model = model.Offset((page - 1) * limit).Limit(limit)
	}
	list, err = model.All()
	return
}

// 详情
func (logicThis *s` + tpl.TableNameCaseCamel + `) Info(ctx context.Context, filter map[string]interface{}, field ...[]string) (info gdb.Record, err error) {
	daoThis := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `
	joinTableArr := []string{}
	model := daoThis.ParseDbCtx(ctx)
	model = model.Handler(daoThis.ParseFilter(filter, &joinTableArr))
	if len(field) > 0 && len(field[0]) > 0 {
		model = model.Handler(daoThis.ParseField(field[0], &joinTableArr))
	}
	if len(joinTableArr) > 0 {
		model = model.Group(daoThis.Table() + ` + "`.`" + ` + daoThis.PrimaryKey())
	}
	info, err = model.One()
	if err != nil {
		return
	}
	if len(info) == 0 {
		err = utils.NewErrorCode(ctx, 29999999, ` + "``" + `)
		return
	}
	return
}

// 新增
func (logicThis *s` + tpl.TableNameCaseCamel + `) Create(ctx context.Context, data map[string]interface{}) (id int64, err error) {
	daoThis := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `
	id, err = daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseInsert(data)).InsertAndGetId()
	return
}

// 修改
func (logicThis *s` + tpl.TableNameCaseCamel + `) Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (err error) {
	daoThis := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `
	idArr, _ := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Array(daoThis.PrimaryKey())
	if len(idArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999999, ` + "``" + `)
		return
	}
	hookData := map[string]interface{}{}

	model := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{}), daoThis.ParseUpdate(data))
	if len(hookData) > 0 {
		model = model.Hook(daoThis.HookUpdate(hookData, gconv.SliceInt(idArr)...))
	}
	_, err = model.UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *s` + tpl.TableNameCaseCamel + `) Delete(ctx context.Context, filter map[string]interface{}) (err error) {
	daoThis := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `
	idArr, _ := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Array(daoThis.PrimaryKey())
	if len(idArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999999, ` + "``" + `)
		return
	}

	_, err = daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Hook(daoThis.HookDelete(gconv.SliceInt(idArr)...)).Delete()
	return
}
`

	gfile.PutContents(saveFile, tplLogic)
}

// api模板生成
func MyGenTplApi(ctx context.Context, option *MyGenOption, tpl *MyGenTpl) {
	saveFile := gfile.SelfDir() + `/api/` + option.SceneCode + `/` + tpl.ModuleDirCaseCamelLower + `/` + tpl.TableNameCaseSnake + `.go`
	if !option.IsCover && gfile.IsFile(saveFile) {
		return
	}

	apiReqFilterColumn := ``
	apiReqCreateColumn := ``
	apiReqUpdateColumn := ``
	apiResColumn := ``
	for _, column := range tpl.TableColumnList {
		field := column[`Field`].String()
		fieldCaseCamel := gstr.CaseCamel(field)
		fieldCaseSnake := gstr.CaseSnakeFirstUpper(field)
		comment := gstr.Trim(gstr.ReplaceByArray(column[`Comment`].String(), g.SliceStr{
			"\n", " ",
			"\r", " ",
		}))
		result, _ := gregex.MatchString(`.*\((\d*)\)`, column[`Type`].String())

		switch field {
		case `deletedAt`, `deleted_at`:
		case `createdAt`, `created_at`, `updatedAt`, `updated_at`:
			apiResColumn += fieldCaseCamel + ` *gtime.Time ` + "`" + `json:"` + field + `" dc:"` + comment + `"` + "`\n"
		case `password`, `passwd`:
			apiReqCreateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"size:` + result[1] + `" dc:"` + comment + `"` + "`\n"
			apiReqUpdateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"size:` + result[1] + `" dc:"` + comment + `"` + "`\n"
		case `sort`, `weight`:
			apiReqFilterColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer|between:0,100" dc:"` + comment + `"` + "`\n"
			apiReqCreateColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer|between:0,100" dc:"` + comment + `"` + "`\n"
			apiReqUpdateColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer|between:0,100" dc:"` + comment + `"` + "`\n"
			apiResColumn += fieldCaseCamel + ` uint ` + "`" + `json:"` + field + `" dc:"` + comment + `"` + "`\n"
		case `pid`:
			apiReqFilterColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer|min:0" dc:"` + comment + `"` + "`\n"
			apiReqCreateColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer|min:0" dc:"` + comment + `"` + "`\n"
			apiReqUpdateColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer|min:0" dc:"` + comment + `"` + "`\n"
			apiResColumn += fieldCaseCamel + ` uint ` + "`" + `json:"` + field + `" dc:"` + comment + `"` + "`\n"
		default:
			//主键
			if column[`Key`].String() == `PRI` && column[`Extra`].String() == `auto_increment` && field != `id` {
				apiReqFilterColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer|min:1" dc:"` + comment + `"` + "`\n"
				apiResColumn += fieldCaseCamel + ` uint ` + "`" + `json:"` + field + `" dc:"` + comment + `"` + "`\n"
				continue
			}
			//id后缀
			if gstr.SubStr(fieldCaseCamel, -2) == `Id` {
				apiReqFilterColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer|min:1" dc:"` + comment + `"` + "`\n"
				apiReqCreateColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer|min:1" dc:"` + comment + `"` + "`\n"
				apiReqUpdateColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer|min:1" dc:"` + comment + `"` + "`\n"
				apiResColumn += fieldCaseCamel + ` uint ` + "`" + `json:"` + field + `" dc:"` + comment + `"` + "`\n"
				continue
			}
			//name或code后缀
			if gstr.SubStr(fieldCaseCamel, -4) == `Name` || gstr.SubStr(fieldCaseCamel, -4) == `Code` {
				apiReqFilterColumn += fieldCaseCamel + ` string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"length:1,` + result[1] + `|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"` + comment + `"` + "`\n"
				apiReqCreateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"length:1,` + result[1] + `|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"` + comment + `"` + "`\n"
				apiReqUpdateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"length:1,` + result[1] + `|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"` + comment + `"` + "`\n"
				apiResColumn += fieldCaseCamel + ` string ` + "`" + `json:"` + field + `" dc:"` + comment + `"` + "`\n"
				continue
			}
			//mobile或phone后缀
			if gstr.SubStr(fieldCaseCamel, -5) == `Phone` || gstr.SubStr(fieldCaseCamel, -6) == `Mobile` {
				apiReqFilterColumn += fieldCaseCamel + ` string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"phone|length:1,` + result[1] + `" dc:"` + comment + `"` + "`\n"
				apiReqCreateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"phone|length:1,` + result[1] + `" dc:"` + comment + `"` + "`\n"
				apiReqUpdateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"phone|length:1,` + result[1] + `" dc:"` + comment + `"` + "`\n"
				apiResColumn += fieldCaseCamel + ` string ` + "`" + `json:"` + field + `" dc:"` + comment + `"` + "`\n"
				continue
			}
			//url或link后缀
			if gstr.SubStr(fieldCaseCamel, -3) == `Url` || gstr.SubStr(fieldCaseCamel, -4) == `Link` {
				apiReqFilterColumn += fieldCaseCamel + ` string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"url|length:1,` + result[1] + `" dc:"` + comment + `"` + "`\n"
				apiReqCreateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"url|length:1,` + result[1] + `" dc:"` + comment + `"` + "`\n"
				apiReqUpdateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"url|length:1,` + result[1] + `" dc:"` + comment + `"` + "`\n"
				apiResColumn += fieldCaseCamel + ` string ` + "`" + `json:"` + field + `" dc:"` + comment + `"` + "`\n"
				continue
			}
			//icon,cover或img,img_list,imgList,img_arr,imgArr或image,image_list,imageList,image_arr,imageArr等后缀
			if field == `avatar` || gstr.SubStr(fieldCaseCamel, -4) == `Icon` || gstr.SubStr(fieldCaseCamel, -5) == `Cover` || gstr.SubStr(fieldCaseCamel, -3) == `Img` || gstr.SubStr(fieldCaseCamel, -7) == `ImgList` || gstr.SubStr(fieldCaseCamel, -6) == `ImgArr` || gstr.SubStr(fieldCaseCamel, -5) == `Image` || gstr.SubStr(fieldCaseCamel, -9) == `ImageList` || gstr.SubStr(fieldCaseCamel, -8) == `ImageArr` {
				if column[`Type`].String() == `json` || column[`Type`].String() == `text` {
					apiReqCreateColumn += fieldCaseCamel + ` *[]string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"distinct|foreach|url|foreach|min-length:1" dc:"` + comment + `"` + "`\n"
					apiReqUpdateColumn += fieldCaseCamel + ` *[]string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"distinct|foreach|url|foreach|min-length:1" dc:"` + comment + `"` + "`\n"
					apiResColumn += fieldCaseCamel + ` []string ` + "`" + `json:"` + field + `" dc:"` + comment + `"` + "`\n"
				} else {
					apiReqCreateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"url|length:1,` + result[1] + `" dc:"` + comment + `"` + "`\n"
					apiReqUpdateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"url|length:1,` + result[1] + `" dc:"` + comment + `"` + "`\n"
					apiResColumn += fieldCaseCamel + ` string ` + "`" + `json:"` + field + `" dc:"` + comment + `"` + "`\n"
				}
				continue
			}
			//video,video_list,videoList,video_arr,videoArr等后缀
			if gstr.SubStr(fieldCaseCamel, -5) == `Video` || gstr.SubStr(fieldCaseCamel, -9) == `VideoList` || gstr.SubStr(fieldCaseCamel, -8) == `VideoArr` {
				if column[`Type`].String() == `json` || column[`Type`].String() == `text` {
					apiReqCreateColumn += fieldCaseCamel + ` *[]string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"distinct|foreach|url|foreach|min-length:1" dc:"` + comment + `"` + "`\n"
					apiReqUpdateColumn += fieldCaseCamel + ` *[]string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"distinct|foreach|url|foreach|min-length:1" dc:"` + comment + `"` + "`\n"
					apiResColumn += fieldCaseCamel + ` []string ` + "`" + `json:"` + field + `" dc:"` + comment + `"` + "`\n"
				} else {
					apiReqCreateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"url|length:1,` + result[1] + `" dc:"` + comment + `"` + "`\n"
					apiReqUpdateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"url|length:1,` + result[1] + `" dc:"` + comment + `"` + "`\n"
					apiResColumn += fieldCaseCamel + ` string ` + "`" + `json:"` + field + `" dc:"` + comment + `"` + "`\n"
				}
				continue
			}
			//Ip后缀
			if gstr.SubStr(fieldCaseCamel, -2) == `Ip` {
				apiReqFilterColumn += fieldCaseCamel + ` string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"ip|length:1,` + result[1] + `" dc:"` + comment + `"` + "`\n"
				apiReqCreateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"ip|length:1,` + result[1] + `" dc:"` + comment + `"` + "`\n"
				apiReqUpdateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"ip|length:1,` + result[1] + `" dc:"` + comment + `"` + "`\n"
				apiResColumn += fieldCaseCamel + ` string ` + "`" + `json:"` + field + `" dc:"` + comment + `"` + "`\n"
				continue
			}
			//status或type后缀
			if field == `gender` || gstr.SubStr(fieldCaseCamel, -6) == `Status` || gstr.SubStr(fieldCaseCamel, -4) == `Type` {
				statusList := MyGenStatusList(comment)
				statusArr := make([]string, len(statusList))
				for index, status := range statusList {
					statusArr[index] = status[0]
				}
				statusStr := gstr.Join(statusArr, `,`)
				apiReqFilterColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer|in:` + statusStr + `" dc:"` + comment + `"` + "`\n"
				apiReqCreateColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer|in:` + statusStr + `" dc:"` + comment + `"` + "`\n"
				apiReqUpdateColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer|in:` + statusStr + `" dc:"` + comment + `"` + "`\n"
				apiResColumn += fieldCaseCamel + ` uint ` + "`" + `json:"` + field + `" dc:"` + comment + `"` + "`\n"
				continue
			}
			//is_前缀
			if gstr.SubStr(fieldCaseSnake, 0, 3) == `is_` {
				apiReqFilterColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer|in:0,1" dc:"` + comment + `"` + "`\n"
				apiReqCreateColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer|in:0,1" dc:"` + comment + `"` + "`\n"
				apiReqUpdateColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer|in:0,1" dc:"` + comment + `"` + "`\n"
				apiResColumn += fieldCaseCamel + ` uint ` + "`" + `json:"` + field + `" dc:"` + comment + `"` + "`\n"
				continue
			}
			//int类型
			if gstr.Pos(column[`Type`].String(), `int`) != -1 {
				if gstr.Pos(column[`Type`].String(), `unsigned`) != -1 {
					apiReqFilterColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer" dc:"` + comment + `"` + "`\n"
					apiReqCreateColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer" dc:"` + comment + `"` + "`\n"
					apiReqUpdateColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer" dc:"` + comment + `"` + "`\n"
					apiResColumn += fieldCaseCamel + ` uint ` + "`" + `json:"` + field + `" dc:"` + comment + `"` + "`\n"
				} else {
					apiReqFilterColumn += fieldCaseCamel + ` *int ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer" dc:"` + comment + `"` + "`\n"
					apiReqCreateColumn += fieldCaseCamel + ` *int ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer" dc:"` + comment + `"` + "`\n"
					apiReqUpdateColumn += fieldCaseCamel + ` *int ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer" dc:"` + comment + `"` + "`\n"
					apiResColumn += fieldCaseCamel + ` int ` + "`" + `json:"` + field + `" dc:"` + comment + `"` + "`\n"
				}
				continue
			}
			//float类型
			if gstr.Pos(column[`Type`].String(), `decimal`) != -1 || gstr.Pos(column[`Type`].String(), `double`) != -1 || gstr.Pos(column[`Type`].String(), `float`) != -1 {
				apiReqFilterColumn += fieldCaseCamel + ` *float64 ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"float" dc:"` + comment + `"` + "`\n"
				apiReqCreateColumn += fieldCaseCamel + ` *float64 ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"float" dc:"` + comment + `"` + "`\n"
				apiReqUpdateColumn += fieldCaseCamel + ` *float64 ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"float" dc:"` + comment + `"` + "`\n"
				apiResColumn += fieldCaseCamel + ` float64 ` + "`" + `json:"` + field + `" dc:"` + comment + `"` + "`\n"
				continue
			}
			//varchar类型
			if gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
				apiReqFilterColumn += fieldCaseCamel + ` string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"length:1,` + result[1] + `" dc:"` + comment + `"` + "`\n"
				apiReqCreateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"length:1,` + result[1] + `" dc:"` + comment + `"` + "`\n"
				apiReqUpdateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"length:1,` + result[1] + `" dc:"` + comment + `"` + "`\n"
				apiResColumn += fieldCaseCamel + ` string ` + "`" + `json:"` + field + `" dc:"` + comment + `"` + "`\n"
				continue
			}
			//char类型
			if gstr.Pos(column[`Type`].String(), `char`) != -1 {
				apiReqFilterColumn += fieldCaseCamel + ` string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"length:1,` + result[1] + `" dc:"` + comment + `"` + "`\n"
				apiReqCreateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"size:` + result[1] + `" dc:"` + comment + `"` + "`\n"
				apiReqUpdateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"size:` + result[1] + `" dc:"` + comment + `"` + "`\n"
				apiResColumn += fieldCaseCamel + ` string ` + "`" + `json:"` + field + `" dc:"` + comment + `"` + "`\n"
				continue
			}
			//datetime和timestamp类型
			if gstr.Pos(column[`Type`].String(), `datetime`) != -1 || gstr.Pos(column[`Type`].String(), `timestamp`) != -1 {
				apiReqFilterColumn += fieldCaseCamel + ` *gtime.Time ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"date-format:Y-m-d H:i:s" dc:"` + comment + `"` + "`\n"
				if column[`Null`].String() == `NO` && column[`Default`].String() == `` {
					apiReqCreateColumn += fieldCaseCamel + ` *gtime.Time ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"required|date-format:Y-m-d H:i:s" dc:"` + comment + `"` + "`\n"
				} else {
					apiReqCreateColumn += fieldCaseCamel + ` *gtime.Time ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"date-format:Y-m-d H:i:s" dc:"` + comment + `"` + "`\n"
				}
				apiReqUpdateColumn += fieldCaseCamel + ` *gtime.Time ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"date-format:Y-m-d H:i:s" dc:"` + comment + `"` + "`\n"
				apiResColumn += fieldCaseCamel + ` *gtime.Time ` + "`" + `json:"` + field + `" dc:"` + comment + `"` + "`\n"
				continue
			}
			//date类型
			if gstr.Pos(column[`Type`].String(), `date`) != -1 {
				apiReqFilterColumn += fieldCaseCamel + ` *gtime.Time ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"date-format:Y-m-d" dc:"` + comment + `"` + "`\n"
				if column[`Null`].String() == `NO` && column[`Default`].String() == `` {
					apiReqCreateColumn += fieldCaseCamel + ` *gtime.Time ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"required|date-format:Y-m-d" dc:"` + comment + `"` + "`\n"
				} else {
					apiReqCreateColumn += fieldCaseCamel + ` *gtime.Time ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"date-format:Y-m-d" dc:"` + comment + `"` + "`\n"
				}
				apiReqUpdateColumn += fieldCaseCamel + ` *gtime.Time ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"date-format:Y-m-d" dc:"` + comment + `"` + "`\n"
				apiResColumn += fieldCaseCamel + ` *gtime.Time ` + "`" + `json:"` + field + `" dc:"` + comment + `"` + "`\n"
				continue
			}
			//json类型
			if gstr.Pos(column[`Type`].String(), `json`) != -1 {
				apiReqFilterColumn += fieldCaseCamel + ` string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"json" dc:"` + comment + `"` + "`\n"
				if column[`Null`].String() == `NO` {
					apiReqCreateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"required|json" dc:"` + comment + `"` + "`\n"
				} else {
					apiReqCreateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"json" dc:"` + comment + `"` + "`\n"
				}
				apiReqUpdateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"json" dc:"` + comment + `"` + "`\n"
				apiResColumn += fieldCaseCamel + ` string ` + "`" + `json:"` + field + `" dc:"` + comment + `"` + "`\n"
				continue
			}
			//text类型
			if gstr.Pos(column[`Type`].String(), `text`) != -1 {
				apiReqFilterColumn += fieldCaseCamel + ` string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"" dc:"` + comment + `"` + "`\n"
				if column[`Null`].String() == `NO` {
					apiReqCreateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"required" dc:"` + comment + `"` + "`\n"
				} else {
					apiReqCreateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"" dc:"` + comment + `"` + "`\n"
				}
				apiReqUpdateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"" dc:"` + comment + `"` + "`\n"
				apiResColumn += fieldCaseCamel + ` string ` + "`" + `json:"` + field + `" dc:"` + comment + `"` + "`\n"
				continue
			}
			//默认处理
			apiReqFilterColumn += fieldCaseCamel + ` string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"" dc:"` + comment + `"` + "`\n"
			apiReqCreateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"" dc:"` + comment + `"` + "`\n"
			apiReqUpdateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"" dc:"` + comment + `"` + "`\n"
			apiResColumn += fieldCaseCamel + ` string ` + "`" + `json:"` + field + `" dc:"` + comment + `"` + "`\n"
		}
	}
	apiReqFilterColumn = gstr.SubStr(apiReqFilterColumn, 0, -len("\n"))
	apiReqCreateColumn = gstr.SubStr(apiReqCreateColumn, 0, -len("\n"))
	apiReqUpdateColumn = gstr.SubStr(apiReqUpdateColumn, 0, -len("\n"))
	apiResColumn = gstr.SubStr(apiResColumn, 0, -len("\n"))

	tplApi := `package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

`
	if option.IsList {
		tplApi += `
/*--------列表 开始--------*/
type ` + tpl.TableNameCaseCamel + `ListReq struct {
	g.Meta ` + "`" + `path:"/` + tpl.TableNameCaseCamelLower + `/list" method:"post" tags:"` + tpl.SceneName + `/` + option.CommonName + `" sm:"列表"` + "`" + `
	Filter ` + tpl.TableNameCaseCamel + `ListFilter ` + "`" + `json:"filter" dc:"查询条件"` + "`" + `
	Field  []string        ` + "`" + `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段。默认会返回全部查询字段。如果需要的字段较少，建议指定字段，传值参考默认返回的字段"` + "`" + `
	Sort   string          ` + "`" + `json:"sort" default:"id DESC" dc:"排序"` + "`" + `
	Page   int             ` + "`" + `json:"page" v:"integer|min:1" default:"1" dc:"页码"` + "`" + `
	Limit  int             ` + "`" + `json:"limit" v:"integer|min:0" default:"10" dc:"每页数量。可传0取全部"` + "`" + `
}

type ` + tpl.TableNameCaseCamel + `ListFilter struct {
	/*--------公共参数 开始--------*/
	Id        *uint       ` + "`" + `c:"id,omitempty" json:"id" v:"integer|min:1" dc:"ID"` + "`" + `
	IdArr     []uint      ` + "`" + `c:"idArr,omitempty" json:"idArr" v:"distinct|foreach|integer|foreach|min:1" dc:"ID数组"` + "`" + `
	ExcId     *uint       ` + "`" + `c:"excId,omitempty" json:"excId" v:"integer|min:1" dc:"排除ID"` + "`" + `
	ExcIdArr  []uint      ` + "`" + `c:"excIdArr,omitempty" json:"excIdArr" v:"distinct|foreach|integer|foreach|min:1" dc:"排除ID数组"` + "`" + `
	StartTime *gtime.Time ` + "`" + `c:"startTime,omitempty" json:"startTime" v:"date-format:Y-m-d H:i:s" dc:"开始时间。示例：2000-01-01 00:00:00"` + "`" + `
	EndTime   *gtime.Time ` + "`" + `c:"endTime,omitempty" json:"endTime" v:"date-format:Y-m-d H:i:s|after-equal:StartTime" dc:"结束时间。示例：2000-01-01 00:00:00"` + "`" + `
	Label     string      ` + "`" + `c:"label,omitempty" json:"label" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"标签。常用于前端组件"` + "`" + `
	/*--------公共参数 结束--------*/
	` + apiReqFilterColumn + `
}

type ` + tpl.TableNameCaseCamel + `ListRes struct {
	Count int         ` + "`" + `json:"count" dc:"总数"` + "`" + `
	List  []` + tpl.TableNameCaseCamel + `Item ` + "`" + `json:"list" dc:"列表"` + "`" + `
}

type ` + tpl.TableNameCaseCamel + `Item struct {
	Id          uint        ` + "`" + `json:"id" dc:"ID"` + "`" + `
	Label       string      ` + "`" + `json:"label" dc:"标签。常用于前端组件"` + "`" + `
	` + apiResColumn + `
}

/*--------列表 结束--------*/

`
	}
	if option.IsUpdate {
		tplApi += `/*--------详情 开始--------*/
type ` + tpl.TableNameCaseCamel + `InfoReq struct {
	g.Meta ` + "`" + `path:"/` + tpl.TableNameCaseCamelLower + `/info" method:"post" tags:"` + tpl.SceneName + `/` + option.CommonName + `" sm:"详情"` + "`" + `
	Id     uint     ` + "`" + `json:"id" v:"required|integer|min:1" dc:"ID"` + "`" + `
	Field  []string ` + "`" + `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段。默认会返回全部查询字段。如果需要的字段较少，建议指定字段，传值参考默认返回的字段"` + "`" + `
}

type ` + tpl.TableNameCaseCamel + `InfoRes struct {
	Info ` + tpl.TableNameCaseCamel + `Info ` + "`" + `json:"info" dc:"详情"` + "`" + `
}

type ` + tpl.TableNameCaseCamel + `Info struct {
	Id          uint        ` + "`" + `json:"id" dc:"ID"` + "`" + `
	Label       string      ` + "`" + `json:"label" dc:"标签。常用于前端组件"` + "`" + `
	` + apiResColumn + `
}

/*--------详情 结束--------*/

`
	}
	if option.IsCreate {
		tplApi += `/*--------新增 开始--------*/
type ` + tpl.TableNameCaseCamel + `CreateReq struct {
	g.Meta      ` + "`" + `path:"/` + tpl.TableNameCaseCamelLower + `/create" method:"post" tags:"` + tpl.SceneName + `/` + option.CommonName + `" sm:"创建"` + "`" + `
	` + apiReqCreateColumn + `
}

/*--------新增 结束--------*/

`
	}

	if option.IsUpdate {
		tplApi += `/*--------修改 开始--------*/
type ` + tpl.TableNameCaseCamel + `UpdateReq struct {
	g.Meta      ` + "`" + `path:"/` + tpl.TableNameCaseCamelLower + `/update" method:"post" tags:"` + tpl.SceneName + `/` + option.CommonName + `" sm:"更新"` + "`" + `
	IdArr       []uint  ` + "`" + `c:"idArr,omitempty" json:"idArr" v:"required|distinct|foreach|integer|foreach|min:1" dc:"ID数组"` + "`" + `
	` + apiReqUpdateColumn + `
}

/*--------修改 结束--------*/

`
	}

	if option.IsDelete {
		tplApi += `/*--------删除 开始--------*/
type ` + tpl.TableNameCaseCamel + `DeleteReq struct {
	g.Meta ` + "`" + `path:"/` + tpl.TableNameCaseCamelLower + `/del" method:"post" tags:"` + tpl.SceneName + `/` + option.CommonName + `" sm:"删除"` + "`" + `
	IdArr  []uint ` + "`" + `c:"idArr,omitempty" json:"idArr" v:"required|distinct|foreach|integer|foreach|min:1" dc:"ID数组"` + "`" + `
}

/*--------删除 结束--------*/
`
	}

	gfile.PutContents(saveFile, tplApi)
}

// controller模板生成
func MyGenTplController(ctx context.Context, option *MyGenOption, tpl *MyGenTpl) {
	saveFile := gfile.SelfDir() + `/internal/controller/` + option.SceneCode + `/` + tpl.ModuleDirCaseCamelLower + `/` + tpl.TableNameCaseSnake + `.go`
	if !option.IsCover && gfile.IsFile(saveFile) {
		return
	}

	controllerAlloweFieldAppend := ``
	if tpl.PrimaryKey != `id` {
		controllerAlloweFieldAppend += `columnsThis.` + gstr.CaseCamel(tpl.PrimaryKey) + `, `
	}
	if tpl.LabelField != `` {
		controllerAlloweFieldAppend += `columnsThis.` + gstr.CaseCamel(tpl.LabelField) + `, `
	}
	controllerAlloweFieldDiff := ``
	for _, column := range tpl.TableColumnList {
		field := column[`Field`].String()
		fieldCaseCamel := gstr.CaseCamel(field)
		switch field {
		case `password`, `passwd`:
			controllerAlloweFieldDiff += `columnsThis.` + fieldCaseCamel + `, `
		}
	}
	controllerAlloweFieldAppend = gstr.SubStr(controllerAlloweFieldAppend, 0, -len(`, `))
	controllerAlloweFieldDiff = gstr.SubStr(controllerAlloweFieldDiff, 0, -len(`, `))

	tplController := `package controller

import (
	"api/api"
	api` + tpl.ModuleDirCaseCamel + ` "api/api/` + option.SceneCode + `/` + tpl.ModuleDirCaseCamelLower + `"
	dao` + tpl.ModuleDirCaseCamel + ` "api/internal/dao/` + tpl.ModuleDirCaseCamelLower + `"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/util/gconv"
)

type ` + tpl.TableNameCaseCamel + ` struct{}

func New` + tpl.TableNameCaseCamel + `() *` + tpl.TableNameCaseCamel + ` {
	return &` + tpl.TableNameCaseCamel + `{}
}

`
	if option.IsList {
		tplController += `// 列表
func (controllerThis *` + tpl.TableNameCaseCamel + `) List(ctx context.Context, req *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `ListReq) (res *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `ListRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.MapDeep(req.Filter)
	if filter == nil {
		filter = map[string]interface{}{}
	}
	order := []string{req.Sort}
	page := req.Page
	limit := req.Limit

	columnsThis := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `.Columns()
	allowField := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `.ColumnArr()
	allowField = append(allowField, ` + "`id`, `label`" + `)`
		if controllerAlloweFieldDiff != `` {
			tplController += `
	allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{` + controllerAlloweFieldDiff + `})).Slice() //移除敏感字段`
		}
		tplController += `
	field := allowField
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	/**--------参数处理 结束--------**/
`
		if option.IsAuthAction {
			actionCode := tpl.RawTableNameCaseCamelLower + `Look`
			actionName := option.CommonName + `-查看`
			daoAuth.Action.MyGenAction(ctx, tpl.SceneId, actionCode, actionName) // 数据库权限操作处理
			tplController += `
	/**--------权限验证 开始--------**/
	isAuth, _ := service.Action().CheckAuth(ctx, ` + "`" + actionCode + "`" + `)
	if !isAuth {
		field = []string{` + "`id`, `label`"
			if controllerAlloweFieldAppend != `` {
				tplController += `, ` + controllerAlloweFieldAppend
			}
			tplController += `}
	}
	/**--------权限验证 结束--------**/
`
		}
		tplController += `
	count, err := service.` + tpl.TableNameCaseCamel + `().Count(ctx, filter)
	if err != nil {
		return
	}
	list, err := service.` + tpl.TableNameCaseCamel + `().List(ctx, filter, field, order, page, limit)
	if err != nil {
		return
	}
	/* //不建议用这个返回，指定字段获取时，返回时其他字段也会返回，但都是空
	res = &api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `ListRes{
		Count: count,
	}
	list.Structs(&res.List) */
	utils.HttpWriteJson(ctx, map[string]interface{}{
		` + "`count`" + `: count,
		` + "`list`" + `:  list,
	}, 0, ` + "``" + `)
	return
}

`
	}
	if option.IsUpdate {
		tplController += `// 详情
func (controllerThis *` + tpl.TableNameCaseCamel + `) Info(ctx context.Context, req *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `InfoReq) (res *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `InfoRes, err error) {
	/**--------参数处理 开始--------**/
	allowField := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `.ColumnArr()
	allowField = append(allowField, ` + "`id`, `label`" + `)`
		if controllerAlloweFieldDiff != `` {
			tplController += `
	columnsThis := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `.Columns()
	allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{` + controllerAlloweFieldDiff + `})).Slice() //移除敏感字段`
		}
		tplController += `
	field := allowField
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	filter := map[string]interface{}{` + "`id`" + `: req.Id}
	/**--------参数处理 结束--------**/
`
		if option.IsAuthAction {
			actionCode := tpl.RawTableNameCaseCamelLower + `Look`
			actionName := option.CommonName + `-查看`
			daoAuth.Action.MyGenAction(ctx, tpl.SceneId, actionCode, actionName) // 数据库权限操作处理
			tplController += `
	/**--------权限验证 开始--------**/
	_, err = service.Action().CheckAuth(ctx, ` + "`" + actionCode + "`" + `)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/
`
		}
		tplController += `
	info, err := service.` + tpl.TableNameCaseCamel + `().Info(ctx, filter, field)
	if err != nil {
		return
	}
	/* //不建议用这个返回，指定字段获取时，返回时其他字段也会返回，但都是空
	res = &api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `InfoRes{}
	info.Struct(&res.Info) */
	utils.HttpWriteJson(ctx, map[string]interface{}{
		` + "`info`" + `: info,
	}, 0, ` + "``" + `)
	return
}

`
	}
	if option.IsCreate {
		tplController += `// 新增
func (controllerThis *` + tpl.TableNameCaseCamel + `) Create(ctx context.Context, req *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `CreateReq) (res *api.CommonCreateRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.MapDeep(req)
	/**--------参数处理 结束--------**/
`
		if option.IsAuthAction {
			actionCode := tpl.RawTableNameCaseCamelLower + `Create`
			actionName := option.CommonName + `-新增`
			daoAuth.Action.MyGenAction(ctx, tpl.SceneId, actionCode, actionName) // 数据库权限操作处理
			tplController += `
	/**--------权限验证 开始--------**/
	_, err = service.Action().CheckAuth(ctx, ` + "`" + actionCode + "`" + `)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/
`
		}
		tplController += `
	id, err := service.` + tpl.TableNameCaseCamel + `().Create(ctx, data)
	if err != nil {
		return
	}
	res = &api.CommonCreateRes{Id: id}
	return
}

`
	}

	if option.IsUpdate {
		tplController += `// 修改
func (controllerThis *` + tpl.TableNameCaseCamel + `) Update(ctx context.Context, req *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `UpdateReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.MapDeep(req)
	delete(data, ` + "`idArr`" + `)
	if len(data) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ` + "``" + `)
		return
	}
	filter := map[string]interface{}{` + "`id`" + `: req.IdArr}
	/**--------参数处理 结束--------**/
`
		if option.IsAuthAction {
			actionCode := tpl.RawTableNameCaseCamelLower + `Update`
			actionName := option.CommonName + `-编辑`
			daoAuth.Action.MyGenAction(ctx, tpl.SceneId, actionCode, actionName) // 数据库权限操作处理
			tplController += `
	/**--------权限验证 开始--------**/
	_, err = service.Action().CheckAuth(ctx, ` + "`" + actionCode + "`" + `)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/
`
		}
		tplController += `
	err = service.` + tpl.TableNameCaseCamel + `().Update(ctx, filter, data)
	return
}

`
	}

	if option.IsDelete {
		tplController += `// 删除
func (controllerThis *` + tpl.TableNameCaseCamel + `) Delete(ctx context.Context, req *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `DeleteReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	filter := map[string]interface{}{` + "`id`" + `: req.IdArr}
	/**--------参数处理 结束--------**/
`
		if option.IsAuthAction {
			actionCode := tpl.RawTableNameCaseCamelLower + `Delete`
			actionName := option.CommonName + `-删除`
			daoAuth.Action.MyGenAction(ctx, tpl.SceneId, actionCode, actionName) // 数据库权限操作处理
			tplController += `
	/**--------权限验证 开始--------**/
	_, err = service.Action().CheckAuth(ctx, ` + "`" + actionCode + "`" + `)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/
`
		}
		tplController += `
	err = service.` + tpl.TableNameCaseCamel + `().Delete(ctx, filter)
	return
}
`
	}

	gfile.PutContents(saveFile, tplController)
}

// 后端路由生成
func MyGenTplRouter(ctx context.Context, option *MyGenOption, tpl *MyGenTpl) {
	saveFile := gfile.SelfDir() + `/internal/router/` + option.SceneCode + `.go`

	tplView := gfile.GetContents(saveFile)

	//控制器不存在时导入
	importControllerStr := `controller` + tpl.ModuleDirCaseCamel + ` "api/internal/controller/` + option.SceneCode + `/` + tpl.ModuleDirCaseCamelLower + `"`
	if gstr.Pos(tplView, importControllerStr) == -1 {
		tplView = gstr.Replace(tplView, `"api/internal/middleware"`, importControllerStr+`
	"api/internal/middleware"`)
		//路由生成
		tplView = gstr.Replace(tplView, `/*--------后端路由自动代码生成锚点（不允许修改和删除，否则将不能自动生成路由）--------*/`, `group.Group(`+"`"+`/`+tpl.ModuleDirCaseCamelLower+"`"+`, func(group *ghttp.RouterGroup) {
				group.Bind(controller`+tpl.ModuleDirCaseCamel+`.New`+tpl.TableNameCaseCamel+`())
			})

			/*--------后端路由自动代码生成锚点（不允许修改和删除，否则将不能自动生成路由）--------*/`)
		gfile.PutContents(saveFile, tplView)
	} else {
		//路由不存在时需生成
		if gstr.Pos(tplView, `group.Bind(controller`+tpl.ModuleDirCaseCamel+`.New`+tpl.TableNameCaseCamel+`())`) == -1 {
			//路由生成
			tplView = gstr.Replace(tplView, `group.Group(`+"`"+`/`+tpl.ModuleDirCaseCamelLower+"`"+`, func(group *ghttp.RouterGroup) {`, `group.Group(`+"`"+`/`+tpl.ModuleDirCaseCamelLower+"`"+`, func(group *ghttp.RouterGroup) {
				group.Bind(controller`+tpl.ModuleDirCaseCamel+`.New`+tpl.TableNameCaseCamel+`())`)
			gfile.PutContents(saveFile, tplView)
		}
	}
}

// 视图模板Index生成
func MyGenTplViewIndex(ctx context.Context, option *MyGenOption, tpl *MyGenTpl) {
	saveFile := gfile.SelfDir() + `/../dev/` + option.SceneCode + `/src/views/` + tpl.ModuleDirCaseCamelLower + `/` + tpl.TableNameCaseCamelLower + `/Index.vue`
	if !option.IsCover && gfile.IsFile(saveFile) {
		return
	}

	tplView := `<script setup lang="ts">
import List from './List.vue'
import Query from './Query.vue'`
	if option.IsCreate || option.IsUpdate {
		tplView += `
import Save from './Save.vue'`
	}
	tplView += `

//搜索
const queryCommon = reactive({
	data: {}
})
provide('queryCommon', queryCommon)

//列表
const listCommon = reactive({
	ref: null as any,
})
provide('listCommon', listCommon)`
	if option.IsCreate || option.IsUpdate {
		tplView += `

//保存
const saveCommon = reactive({
	visible: false,
	title: '',  //新增|编辑|复制
	data: {}
})
provide('saveCommon', saveCommon)`
	}
	tplView += `
</script>

<template>
	<ElContainer class="main-table-container">
		<ElHeader>
			<Query />
		</ElHeader>

		<List :ref="(el: any) => { listCommon.ref = el }" />`
	if option.IsCreate || option.IsUpdate {
		tplView += `

		<!-- 加上v-if每次都重新生成组件。可防止不同操作之间的影响；新增操作数据的默认值也能写在save组件内 -->
		<Save v-if="saveCommon.visible" />`
	}
	tplView += `
	</ElContainer>
</template>`

	gfile.PutContents(saveFile, tplView)
}

// 视图模板List生成
func MyGenTplViewList(ctx context.Context, option *MyGenOption, tpl *MyGenTpl) {
	saveFile := gfile.SelfDir() + `/../dev/` + option.SceneCode + `/src/views/` + tpl.ModuleDirCaseCamelLower + `/` + tpl.TableNameCaseCamelLower + `/List.vue`
	if !option.IsCover && gfile.IsFile(saveFile) {
		return
	}

	rawCreatedAtField := ``
	rawUpdatedAtField := ``
	// rawDeletedAtField := ``
	tableRowHeight := 50
	viewListColumn := ``
	for _, column := range tpl.TableColumnList {
		field := column[`Field`].String()
		fieldCaseCamel := gstr.CaseCamel(field)
		fieldCaseSnake := gstr.CaseSnakeFirstUpper(field)
		comment := gstr.Trim(gstr.ReplaceByArray(column[`Comment`].String(), g.SliceStr{
			"\n", " ",
			"\r", " ",
		}))
		switch field {
		case `deletedAt`, `deleted_at`:
			// rawDeletedAtField = field
		case `createdAt`, `created_at`:
			rawCreatedAtField = field
		case `updatedAt`, `updated_at`:
			rawUpdatedAtField = field
		case `password`, `passwd`:
		case `sort`, `weight`:
			viewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
		key: '` + field + `',
		align: 'center',
		width: 100,
		sortable: true,`
			if option.IsUpdate {
				viewListColumn += `
		cellRenderer: (props: any): any => {
			if (props.rowData.edit` + gstr.CaseCamel(field) + `) {
				let currentRef: any
				let currentVal = props.rowData.` + field + `
				return [
					h(ElInputNumber as any, {
						'ref': (el: any) => { currentRef = el; el?.focus() },
						'model-value': currentVal,
						'placeholder': t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.tip.` + field + `'),
						'precision': 0,
						'min': 0,
						'max': 100,
						'step': 1,
						'step-strictly': true,
						'controls': false,  //控制按钮会导致诸多问题。如：焦点丢失；` + field + `是0或100时，只一个按钮可点击
						'controls-position': 'right',
						onChange: (val: number) => {
							currentVal = val
						},
						onBlur: () => {
							props.rowData.edit` + gstr.CaseCamel(field) + ` = false
							if ((currentVal || currentVal === 0) && currentVal != props.rowData.` + field + `) {
								handleUpdate({
									idArr: [props.rowData.id],
									` + field + `: currentVal
								}).then((res) => {
									props.rowData.` + field + ` = currentVal
								}).catch((error) => {
								})
							}
						},
						onKeydown: (event: any) => {
							switch (event.keyCode) {
								//case 27:    //Esc键：Escape
								//case 32:    //空格键：" "
								case 13:    //Enter键：Enter
									//props.rowData.edit` + gstr.CaseCamel(field) + ` = false  //也会触发onBlur事件
									currentRef?.blur()
									break;
							}
						},
					})
				]
			}
			return [
				h('div', {
					class: 'inline-edit',
					onClick: () => {
						props.rowData.edit` + gstr.CaseCamel(field) + ` = true
					}
				}, {
					default: () => props.rowData.` + field + `
				})
			]
		}`
			}
			viewListColumn += `
	},`
		default:
			//主键
			if column[`Key`].String() == `PRI` && column[`Extra`].String() == `auto_increment` {
				continue
			}
			//id后缀
			if field == `pid` || gstr.SubStr(fieldCaseCamel, -2) == `Id` {
				viewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
		key: '` + field + `',
		align: 'center',
		width: 150,
	},`
				continue
			}
			//name或code后缀
			if gstr.SubStr(fieldCaseCamel, -4) == `Name` || gstr.SubStr(fieldCaseCamel, -4) == `Code` {
				viewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
		key: '` + field + `',
		align: 'center',
		width: 150,
	},`
				continue
			}
			//mobile或phone后缀
			if gstr.SubStr(fieldCaseCamel, -5) == `Phone` || gstr.SubStr(fieldCaseCamel, -6) == `Mobile` {
				viewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
		key: '` + field + `',
		align: 'center',
		width: 150,
	},`
				continue
			}
			//url或link后缀
			if gstr.SubStr(fieldCaseCamel, -3) == `Url` || gstr.SubStr(fieldCaseCamel, -4) == `Link` {
				viewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
		key: '` + field + `',
		align: 'center',
		width: 300,
	},`
				continue
			}
			//icon,cover或img,img_list,imgList,img_arr,imgArr或image,image_list,imageList,image_arr,imageArr等后缀
			if field == `avatar` || gstr.SubStr(fieldCaseCamel, -4) == `Icon` || gstr.SubStr(fieldCaseCamel, -5) == `Cover` || gstr.SubStr(fieldCaseCamel, -3) == `Img` || gstr.SubStr(fieldCaseCamel, -7) == `ImgList` || gstr.SubStr(fieldCaseCamel, -6) == `ImgArr` || gstr.SubStr(fieldCaseCamel, -5) == `Image` || gstr.SubStr(fieldCaseCamel, -9) == `ImageList` || gstr.SubStr(fieldCaseCamel, -8) == `ImageArr` {
				viewListColumn += `
	{
        dataKey: '` + field + `',
        title: t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
        key: '` + field + `',
        align: 'center',
        width: 100,
        cellRenderer: (props: any): any => {
            if (!props.rowData.` + field + `) {
                return
            }`
				if column[`Type`].String() == `json` || column[`Type`].String() == `text` {
					viewListColumn += `
			let imageList: string[]
			if (Array.isArray(props.rowData.` + field + `)) {
				imageList = props.rowData.` + field + `
			} else {
				imageList = JSON.parse(props.rowData.` + field + `)
			}`
				} else {
					viewListColumn += `
			const imageList = [props.rowData.` + field + `]`
				}
				viewListColumn += `
            return [
                h(ElScrollbar, {
                    'wrap-style': 'display: flex; align-items: center;',
                    'view-style': 'margin: auto;',
                }, {
                    default: () => {
                        const content = imageList.map((item) => {
                            return h(ElImage as any, {
                                'style': 'width: 45px;',    //不想显示滚动条，需设置table属性row-height增加行高
                                'src': item,
                                'lazy': true,
                                'hide-on-click-modal': true,
                                'preview-teleported': true,
                                'preview-src-list': imageList
                            })
                        })
                        return content
                    }
                })
            ]
        },
    },`
				continue
			}
			//video,video_list,videoList,video_arr,videoArr等后缀
			if gstr.SubStr(fieldCaseCamel, -5) == `Video` || gstr.SubStr(fieldCaseCamel, -9) == `VideoList` || gstr.SubStr(fieldCaseCamel, -8) == `VideoArr` {
				if tableRowHeight < 100 {
					tableRowHeight = 100
				}
				viewListColumn += `
	{
        dataKey: '` + field + `',
        title: t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
        key: '` + field + `',
        align: 'center',
        width: 150,
        cellRenderer: (props: any): any => {
            if (!props.rowData.` + field + `) {
                return
            }`
				if column[`Type`].String() == `json` || column[`Type`].String() == `text` {
					viewListColumn += `
			let videoList: string[]
			if (Array.isArray(props.rowData.` + field + `)) {
				videoList = props.rowData.` + field + `
			} else {
				videoList = JSON.parse(props.rowData.` + field + `)
			}`
				} else {
					viewListColumn += `
			const videoList = [props.rowData.` + field + `]`
				}
				viewListColumn += `
            return [
                h(ElScrollbar, {
                    'wrap-style': 'display: flex; align-items: center;',
                    'view-style': 'margin: auto;',
                }, {
                    default: () => {
                        const content = videoList.map((item) => {
                            return h('video', {
								'style': 'width: 120px; height: 80px;',    //不想显示滚动条，需设置table属性row-height增加行高
								'preload': 'none',
								'controls': true,
								'src': item
							})
                        })
                        return content
                    }
                })
            ]
        },
    },`
				continue
			}
			//Ip后缀
			if gstr.SubStr(fieldCaseCamel, -2) == `Ip` {
				viewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
		key: '` + field + `',
		align: 'center',
		width: 150,
	},`
				continue
			}
			//remark后缀
			if gstr.SubStr(fieldCaseCamel, -6) == `Remark` {
				viewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
		key: '` + field + `',
		align: 'center',
		width: 150,
		hidden: true,
	},`
				continue
			}
			//status或type后缀
			if field == `gender` || gstr.SubStr(fieldCaseCamel, -6) == `Status` || gstr.SubStr(fieldCaseCamel, -4) == `Type` {
				statusList := MyGenStatusList(comment)
				tagTypeStr := ``
				tagTypeArr := []string{``, `success`, `danger`, `info`, `warning`}
				for index, status := range statusList {
					if index < len(tagTypeArr) {
						tagTypeStr += status[0] + `: '` + tagTypeArr[index] + `', `
					} else {
						tagTypeStr += status[0] + `: '', `
					}
				}
				tagTypeStr = gstr.SubStr(tagTypeStr, 0, -len(`, `))
				viewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
		key: '` + field + `',
		align: 'center',
		width: 100,
		cellRenderer: (props: any): any => {
			let typeObj: any = { ` + tagTypeStr + ` }
			return [
				h(ElTag as any, {
					type: typeObj?.[props.rowData.` + field + `] ?? ''
				}, {
					default: () => (tm('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.status.` + field + `') as any).find((item: any) => { return item.value == props.rowData.` + field + ` })?.label
				})
			]
		}
	},`
				continue
			}
			//is_前缀
			if gstr.SubStr(fieldCaseSnake, 0, 3) == `is_` {
				viewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
		key: '` + field + `',
		align: 'center',
		width: 100,
		cellRenderer: (props: any): any => {
			return [
				h(ElSwitch as any, {
					'model-value': props.rowData.` + field + `,
					'active-value': 1,
					'inactive-value': 0,
					'inline-prompt': true,
					'active-text': t('common.yes'),
					'inactive-text': t('common.no'),
					style: '--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success)',`
				if option.IsUpdate {
					viewListColumn += `
					onChange: (val: number) => {
						handleUpdate({
							idArr: [props.rowData.id],
							` + field + `: val
						}).then((res) => {
							props.rowData.` + field + ` = val
						}).catch((error) => { })
					}`
				}
				viewListColumn += `
				})
			]
		}
	},`
				continue
			}
			//int类型
			if gstr.Pos(column[`Type`].String(), `int`) != -1 {
				viewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
		key: '` + field + `',
		align: 'center',
		width: 150,
	},`
				continue
			}
			//float类型
			if gstr.Pos(column[`Type`].String(), `decimal`) != -1 || gstr.Pos(column[`Type`].String(), `double`) != -1 || gstr.Pos(column[`Type`].String(), `float`) != -1 {
				viewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
		key: '` + field + `',
		align: 'center',
		width: 150,
	},`
				continue
			}
			//varchar类型
			if gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
				viewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
		key: '` + field + `',
		align: 'center',
		width: 150,
	},`
				continue
			}
			//char类型
			if gstr.Pos(column[`Type`].String(), `char`) != -1 {
				viewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
		key: '` + field + `',
		align: 'center',
		width: 150,
	},`
				continue
			}
			//json类型
			if column[`Type`].String() == `json` {
				viewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
		key: '` + field + `',
		align: 'center',
		width: 200,
        hidden: true
	},`
				continue
			}
			//datetime和timestamp类型
			if gstr.Pos(column[`Type`].String(), `datetime`) != -1 || gstr.Pos(column[`Type`].String(), `timestamp`) != -1 {
				viewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
		key: '` + field + `',
		align: 'center',
		width: 150,
        sortable: true
	},`
				continue
			}
			//date类型
			if gstr.Pos(column[`Type`].String(), `date`) != -1 {
				viewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
		key: '` + field + `',
		align: 'center',
		width: 150,
        sortable: true
	},`
				continue
			}
			//默认处理
			viewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
		key: '` + field + `',
		align: 'center',
		width: 150,
	},`
		}
	}

	tplView := `<script setup lang="ts">
const { t, tm } = useI18n()

const table = reactive({
	columns: [{
		dataKey: 'id',
		title: t('common.name.id'),
		key: 'id',
		align: 'center',
		width: 200,
		fixed: 'left',
		sortable: true,`
	if option.IsUpdate || option.IsDelete {
		tplView += `
		headerCellRenderer: () => {
			const allChecked = table.data.every((item: any) => item.checked)
			const someChecked = table.data.some((item: any) => item.checked)
			return [
				h('div', {
					class: 'id-checkbox',
					onClick: (event: any) => {
						event.stopPropagation();    //阻止冒泡
					},
				}, {
					default: () => [
						h(ElCheckbox as any, {
							'model-value': table.data.length ? allChecked : false,
							indeterminate: someChecked && !allChecked,
							onChange: (val: boolean) => {
								table.data.forEach((item: any) => {
									item.checked = val
								})
							}
						})
					]
				}),
				h('div', {}, {
					default: () => t('common.name.id')
				})
			]
		},
		cellRenderer: (props: any): any => {
			return [
				h(ElCheckbox as any, {
					class: 'id-checkbox',
					'model-value': props.rowData.checked,
					onChange: (val: boolean) => {
						props.rowData.checked = val
					}
				}),
				h('div', {}, {
					default: () => props.rowData.id
				})
			]
		},`
	}
	tplView += `
	},` + viewListColumn + `
	{
		dataKey: '` + rawUpdatedAtField + `',
		title: t('common.name.updatedAt'),
		key: '` + rawUpdatedAtField + `',
		align: 'center',
		width: 150,
		sortable: true,
	},
	{
		dataKey: '` + rawCreatedAtField + `',
		title: t('common.name.createdAt'),
		key: '` + rawCreatedAtField + `',
		align: 'center',
		width: 150,
		sortable: true
	},`
	if option.IsCreate || option.IsUpdate || option.IsDelete {
		tplView += `
	{
		title: t('common.name.action'),
		key: 'action',
		align: 'center',
		width: 250,
		fixed: 'right',
		cellRenderer: (props: any): any => {
			return [`
		if option.IsUpdate {
			tplView += `
				h(ElButton, {
					type: 'primary',
					size: 'small',
					onClick: () => handleEditCopy(props.rowData.id)
				}, {
					default: () => [h(AutoiconEpEdit), t('common.edit')]
				}),`
		}
		if option.IsDelete {
			tplView += `
				h(ElButton, {
					type: 'danger',
					size: 'small',
					onClick: () => handleDelete([props.rowData.id])
				}, {
					default: () => [h(AutoiconEpDelete), t('common.delete')]
				}),`
		}
		if option.IsCreate {
			tplView += `
				h(ElButton, {
					type: 'warning',
					size: 'small',
					onClick: () => handleEditCopy(props.rowData.id, 'copy')
				}, {
					default: () => [h(AutoiconEpDocumentCopy), t('common.copy')]
				}),`
		}
		tplView += `
			]
		},
	}`
	}
	tplView += `] as any,
	data: [],
	loading: false,
	sort: { key: 'id', order: 'desc' } as any,
	handleSort: (sort: any) => {
		table.sort.key = sort.key
		table.sort.order = sort.order
		getList()
	},
})`
	if option.IsCreate || option.IsUpdate {
		tplView += `

const saveCommon = inject('saveCommon') as { visible: boolean, title: string, data: { [propName: string]: any } }`
	}
	if option.IsCreate {
		tplView += `
//新增
const handleAdd = () => {
	saveCommon.data = {}
	saveCommon.title = t('common.add')
	saveCommon.visible = true
}`
	}
	if option.IsDelete {
		tplView += `
//批量删除
const handleBatchDelete = () => {
	const idArr: number[] = [];
	table.data.forEach((item: any) => {
		if (item.checked) {
			idArr.push(item.id)
		}
	})
	if (idArr.length) {
		handleDelete(idArr)
	} else {
		ElMessage.error(t('common.tip.selectDelete'))
	}
}`
	}
	if option.IsCreate || option.IsUpdate {
		tplView += `
//编辑|复制
const handleEditCopy = (id: number, type: string = 'edit') => {
	request(t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseCamelLower + `/` + tpl.TableNameCaseCamelLower + `/info', { id: id }).then((res) => {
		saveCommon.data = { ...res.data.info }
		switch (type) {
			case 'edit':
				saveCommon.data.idArr = [saveCommon.data.id]
				delete saveCommon.data.id
				saveCommon.title = t('common.edit')
				break;
			case 'copy':
				delete saveCommon.data.id
				saveCommon.title = t('common.copy')
				break;
		}
		saveCommon.visible = true
	}).catch(() => { })
}`
	}
	if option.IsDelete {
		tplView += `
//删除
const handleDelete = (idArr: number[]) => {
	ElMessageBox.confirm('', {
		type: 'warning',
		title: t('common.tip.configDelete'),
		center: true,
		showClose: false,
	}).then(() => {
		request(t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseCamelLower + `/` + tpl.TableNameCaseCamelLower + `/del', { idArr: idArr }, true).then((res) => {
			getList()
		}).catch(() => { })
	}).catch(() => { })
}`
	}
	if option.IsUpdate {
		tplView += `
//更新
const handleUpdate = async (param: { idArr: number[], [propName: string]: any }) => {
	await request(t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseCamelLower + `/` + tpl.TableNameCaseCamelLower + `/update', param, true)
}`
	}
	tplView += `

//分页
const settingStore = useSettingStore()
const pagination = reactive({
	total: 0,
	page: 1,
	size: settingStore.pagination.size,
	sizeList: settingStore.pagination.sizeList,
	layout: settingStore.pagination.layout,
	sizeChange: (val: number) => {
		getList()
	},
	pageChange: (val: number) => {
		getList()
	}
})

const queryCommon = inject('queryCommon') as { data: { [propName: string]: any } }
//列表
const getList = async (resetPage: boolean = false) => {
	if (resetPage) {
		pagination.page = 1
	}
	const param = {
		field: [],
		filter: removeEmptyOfObj(queryCommon.data),
		sort: table.sort.key + ' ' + table.sort.order,
		page: pagination.page,
		limit: pagination.size
	}
	table.loading = true
	try {
		const res = await request(t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseCamelLower + `/` + tpl.TableNameCaseCamelLower + `/list', param)
		table.data = res.data.list?.length ? res.data.list : []
		pagination.total = res.data.count
	} catch (error) { }
	table.loading = false
}
getList()

//暴露组件接口给父组件
defineExpose({
	getList
})
</script>

<template>
	<ElRow class="main-table-tool">
		<ElCol :span="16">
			<ElSpace :size="10" style="height: 100%; margin-left: 10px;">`
	if option.IsCreate {
		tplView += `
				<ElButton type="primary" @click="handleAdd">
					<AutoiconEpEditPen />{{ t('common.add') }}
				</ElButton>`
	}
	if option.IsDelete {
		tplView += `
				<ElButton type="danger" @click="handleBatchDelete">
					<AutoiconEpDeleteFilled />{{ t('common.batchDelete') }}
				</ElButton>`
	}
	tplView += `
			</ElSpace>
		</ElCol>
		<ElCol :span="8" style="text-align: right;">
			<ElSpace :size="10" style="height: 100%;">
				<ElDropdown max-height="300" :hide-on-click="false">
					<ElButton type="info" :circle="true">
						<AutoiconEpHide />
					</ElButton>
					<template #dropdown>
						<ElDropdownMenu>
							<ElDropdownItem v-for="(item, key) in table.columns" :key="key">
								<ElCheckbox v-model="item.hidden">
									{{ item.title }}
								</ElCheckbox>
							</ElDropdownItem>
						</ElDropdownMenu>
					</template>
				</ElDropdown>
			</ElSpace>
		</ElCol>
	</ElRow>

	<ElMain>
		<ElAutoResizer>
			<template #default="{ height, width }">
				<ElTableV2 class="main-table" :columns="table.columns" :data="table.data" :sort-by="table.sort"
					@column-sort="table.handleSort" :width="width" :height="height" :fixed="true" :row-height="` + gconv.String(tableRowHeight) + `">
					<template v-if="table.loading" #overlay>
						<ElIcon class="is-loading" color="var(--el-color-primary)" :size="25">
							<AutoiconEpLoading />
						</ElIcon>
					</template>
				</ElTableV2>
			</template>
		</ElAutoResizer>
	</ElMain>

	<ElRow class="main-table-pagination">
		<ElCol :span="24">
			<ElPagination :total="pagination.total" v-model:currentPage="pagination.page"
				v-model:page-size="pagination.size" @size-change="pagination.sizeChange"
				@current-change="pagination.pageChange" :page-sizes="pagination.sizeList" :layout="pagination.layout"
				:background="true" />
		</ElCol>
	</ElRow>
</template>`

	gfile.PutContents(saveFile, tplView)
}

// 视图模板Query生成
func MyGenTplViewQuery(ctx context.Context, option *MyGenOption, tpl *MyGenTpl) {
	saveFile := gfile.SelfDir() + `/../dev/` + option.SceneCode + `/src/views/` + tpl.ModuleDirCaseCamelLower + `/` + tpl.TableNameCaseCamelLower + `/Query.vue`
	if !option.IsCover && gfile.IsFile(saveFile) {
		return
	}

	viewQueryField := ``
	for _, column := range tpl.TableColumnList {
		field := column[`Field`].String()
		fieldCaseCamel := gstr.CaseCamel(field)
		fieldCaseSnake := gstr.CaseSnakeFirstUpper(field)

		switch field {
		case `deletedAt`, `deleted_at`, `createdAt`, `created_at`, `updatedAt`, `updated_at`:
		case `password`, `passwd`:
		case `pid`:
			viewQueryField += `
		<ElFormItem prop="` + field + `">
			<MyCascader v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseCamelLower + `/` + tpl.TableNameCaseCamelLower + `/tree' }" :defaultOptions="[{ id: 0, label: t('common.name.allTopLevel') }]" :props="{ checkStrictly: true, emitPath: false }" />
		</ElFormItem>`
		case `sort`, `weight`:
		default:
			//主键
			if column[`Key`].String() == `PRI` && column[`Extra`].String() == `auto_increment` {
				continue
			}
			//id后缀
			if gstr.SubStr(fieldCaseCamel, -2) == `Id` {
				viewQueryField += `
		<ElFormItem prop="` + field + `">
			<MySelect v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseCamelLower + `/` + gstr.CaseCamelLower(gstr.SubStr(field, 0, -2)) + `/list' }" />
		</ElFormItem>`
				continue
			}
			//name或code后缀
			if gstr.SubStr(fieldCaseCamel, -4) == `Name` || gstr.SubStr(fieldCaseCamel, -4) == `Code` {
				viewQueryField += `
		<ElFormItem prop="` + field + `">
			<ElInput v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :clearable="true" />
		</ElFormItem>`
				continue
			}
			//mobile或phone后缀
			if gstr.SubStr(fieldCaseCamel, -5) == `Phone` || gstr.SubStr(fieldCaseCamel, -6) == `Mobile` {
				viewQueryField += `
		<ElFormItem prop="` + field + `">
			<ElInput v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :clearable="true" />
		</ElFormItem>`
				continue
			}
			//url或link后缀
			if gstr.SubStr(fieldCaseCamel, -3) == `Url` || gstr.SubStr(fieldCaseCamel, -4) == `Link` {
				viewQueryField += `
		<ElFormItem prop="` + field + `">
			<ElInput v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :clearable="true" />
		</ElFormItem>`
				continue
			}
			//icon,cover或img,img_list,imgList,img_arr,imgArr或image,image_list,imageList,image_arr,imageArr等后缀
			if field == `avatar` || gstr.SubStr(fieldCaseCamel, -5) == `Cover` || gstr.SubStr(fieldCaseCamel, -3) == `Img` || gstr.SubStr(fieldCaseCamel, -7) == `ImgList` || gstr.SubStr(fieldCaseCamel, -6) == `ImgArr` || gstr.SubStr(fieldCaseCamel, -5) == `Image` || gstr.SubStr(fieldCaseCamel, -9) == `ImageList` || gstr.SubStr(fieldCaseCamel, -8) == `ImageArr` {
				continue
			}
			//video,video_list,videoList,video_arr,videoArr等后缀
			if gstr.SubStr(fieldCaseCamel, -5) == `Video` || gstr.SubStr(fieldCaseCamel, -9) == `VideoList` || gstr.SubStr(fieldCaseCamel, -8) == `VideoArr` {
				continue
			}
			//Ip后缀
			if gstr.SubStr(fieldCaseCamel, -2) == `Ip` {
				viewQueryField += `
		<ElFormItem prop="` + field + `">
			<ElInput v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :clearable="true" />
		</ElFormItem>`
				continue
			}
			//status或type后缀
			if field == `gender` || gstr.SubStr(fieldCaseCamel, -6) == `Status` || gstr.SubStr(fieldCaseCamel, -4) == `Type` {
				viewQueryField += `
		<ElFormItem prop="` + field + `" style="width: 100px;">
			<ElSelectV2 v-model="queryCommon.data.` + field + `" :options="tm('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.status.` + field + `')" :placeholder="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :clearable="true" />
		</ElFormItem>`
				continue
			}
			//remark后缀
			if gstr.SubStr(fieldCaseCamel, -6) == `Remark` {
				continue
			}
			//is_前缀
			if gstr.SubStr(fieldCaseSnake, 0, 3) == `is_` {
				viewQueryField += `
		<ElFormItem prop="` + field + `" style="width: 100px;">
			<ElSelectV2 v-model="queryCommon.data.` + field + `" :options="tm('common.status.whether')" :placeholder="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :clearable="true" />
		</ElFormItem>`
				continue
			}
			//int类型
			if gstr.Pos(column[`Type`].String(), `int`) != -1 {
				viewQueryField += `
		<ElFormItem prop="` + field + `">
			<ElInputNumber v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :controls="false" />
		</ElFormItem>`
				continue
			}
			//float类型
			if gstr.Pos(column[`Type`].String(), `decimal`) != -1 || gstr.Pos(column[`Type`].String(), `double`) != -1 || gstr.Pos(column[`Type`].String(), `float`) != -1 {
				viewQueryField += `
		<ElFormItem prop="` + field + `">
			<ElInputNumber v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :precision="2" :controls="false" />
		</ElFormItem>`
				continue
			}
			//varchar类型
			if gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
				viewQueryField += `
		<ElFormItem prop="` + field + `">
			<ElInput v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :clearable="true" />
		</ElFormItem>`
				continue
			}
			//char类型
			if gstr.Pos(column[`Type`].String(), `char`) != -1 {
				viewQueryField += `
		<ElFormItem prop="` + field + `">
			<ElInput v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :clearable="true" />
		</ElFormItem>`
				continue
			}
			//json类型
			if column[`Type`].String() == `json` {
				continue
			}
			//datetime和timestamp类型
			if gstr.Pos(column[`Type`].String(), `datetime`) != -1 || gstr.Pos(column[`Type`].String(), `timestamp`) != -1 {
				continue
			}
			//date类型
			if gstr.Pos(column[`Type`].String(), `date`) != -1 {
				continue
			}
			//默认处理
			viewQueryField += `
		<ElFormItem prop="` + field + `">
			<ElInput v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :clearable="true" />
		</ElFormItem>`
		}
	}

	tplView := `<script setup lang="ts">
import dayjs from 'dayjs'

const { t, tm } = useI18n()

const queryCommon = inject('queryCommon') as { data: { [propName: string]: any } }
queryCommon.data = {
	...queryCommon.data,
	timeRange: (() => {
		const date = new Date()
		return [
			// new Date(date.getFullYear(), date.getMonth(), date.getDate(), 0, 0, 0),
			// new Date(date.getFullYear(), date.getMonth(), date.getDate(), 23, 59, 59),
		]
	})(),
	startTime: computed(() => {
		if (queryCommon.data.timeRange?.length) {
			return dayjs(queryCommon.data.timeRange[0]).format('YYYY-MM-DD HH:mm:ss')
		}
		return ''
	}),
	endTime: computed(() => {
		if (queryCommon.data.timeRange?.length) {
			return dayjs(queryCommon.data.timeRange[1]).format('YYYY-MM-DD HH:mm:ss')
		}
		return ''
	})
}
const listCommon = inject('listCommon') as { ref: any }
const queryForm = reactive({
	ref: null as any,
	loading: false,
	submit: () => {
		queryForm.loading = true
		listCommon.ref.getList(true).finally(() => {
			queryForm.loading = false
		})
	},
	reset: () => {
		queryForm.ref.resetFields()
		//queryForm.submit()
	}
})
</script>

<template>
	<ElForm class="query-form" :ref="(el: any) => { queryForm.ref = el }" :model="queryCommon.data" :inline="true"
		@keyup.enter="queryForm.submit">
		<ElFormItem prop="id">
			<ElInputNumber v-model="queryCommon.data.id" :placeholder="t('common.name.id')" :min="1" :controls="false" />
		</ElFormItem>` + viewQueryField + `
		<ElFormItem prop="timeRange">
			<ElDatePicker v-model="queryCommon.data.timeRange" type="datetimerange" range-separator="-"
				:default-time="queryCommon.data.timeRange" :start-placeholder="t('common.name.startTime')"
				:end-placeholder="t('common.name.endTime')">
			</ElDatePicker>
		</ElFormItem>
		<ElFormItem>
			<ElButton type="primary" @click="queryForm.submit" :loading="queryForm.loading">
				<AutoiconEpSearch />{{ t('common.query') }}
			</ElButton>
			<ElButton type="info" @click="queryForm.reset">
				<AutoiconEpCircleClose />{{ t('common.reset') }}
			</ElButton>
		</ElFormItem>
	</ElForm>
</template>`

	gfile.PutContents(saveFile, tplView)
}

// 视图模板Save生成
func MyGenTplViewSave(ctx context.Context, option *MyGenOption, tpl *MyGenTpl) {
	if !(option.IsCreate || option.IsUpdate) {
		return
	}
	saveFile := gfile.SelfDir() + `/../dev/` + option.SceneCode + `/src/views/` + tpl.ModuleDirCaseCamelLower + `/` + tpl.TableNameCaseCamelLower + `/Save.vue`
	if !option.IsCover && gfile.IsFile(saveFile) {
		return
	}

	passwordField := ``
	viewSaveRule := ``
	viewSaveField := ``
	for _, column := range tpl.TableColumnList {
		field := column[`Field`].String()
		fieldCaseCamel := gstr.CaseCamel(field)
		fieldCaseSnake := gstr.CaseSnakeFirstUpper(field)
		comment := gstr.Trim(gstr.ReplaceByArray(column[`Comment`].String(), g.SliceStr{
			"\n", " ",
			"\r", " ",
		}))
		result, _ := gregex.MatchString(`.*\((\d*)\)`, column[`Type`].String())

		switch field {
		case `deletedAt`, `deleted_at`, `createdAt`, `created_at`, `updatedAt`, `updated_at`:
		case `password`, `passwd`:
			passwordField = field
			viewSaveRule += `
		` + field + `: [
			{ type: 'string', required: computed((): boolean => { return saveForm.data.idArr?.length ? false : true; }), min: 1, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 30 }) }
		],`
			viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
                    <ElInput v-model="saveForm.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" minlength="1" maxlength="30" :show-word-limit="true" :clearable="true" :show-password="true" style="max-width: 250px;" />
                    <label v-if="saveForm.data.id">
                        <ElAlert :title="t('common.tip.notRequired')" type="info" :show-icon="true" :closable="false" />
                    </label>
                </ElFormItem>`
		case `pid`:
			viewSaveRule += `
		` + field + `: [
			{ type: 'integer', min: 0, trigger: 'change', message: t('validation.select') }
		],`
			viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
                    <MyCascader v-model="saveForm.data.` + field + `" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseCamelLower + `/` + tpl.TableNameCaseCamelLower + `/tree', param: { filter: { excId: saveForm.data.id } } }" :defaultOptions="[{ id: 0, label: t('common.name.without') }]" :clearable="false" :props="{ checkStrictly: true, emitPath: false }" />
                </ElFormItem>`
		case `sort`, `weight`:
			viewSaveRule += `
		` + field + `: [
			{ type: 'integer', min: 0, max: 100, trigger: 'change', message: t('validation.between.number', { min: 0, max: 100 }) }
		],`
			viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
                    <ElInputNumber v-model="saveForm.data.` + field + `" :precision="0" :min="0" :max="100" :step="1" :step-strictly="true" controls-position="right" :value-on-clear="50" />
                    <label>
                        <ElAlert :title="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.tip.` + field + `')" type="info" :show-icon="true" :closable="false" />
                    </label>
                </ElFormItem>`
		default:
			//主键
			if column[`Key`].String() == `PRI` && column[`Extra`].String() == `auto_increment` {
				continue
			}
			//id后缀
			if gstr.SubStr(fieldCaseCamel, -2) == `Id` {
				viewSaveRule += `
		` + field + `: [
			{ type: 'integer', min: 1, trigger: 'change', message: t('validation.select') }
		],`
				viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
                    <MySelect v-model="saveForm.data.` + field + `" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseCamelLower + `/` + gstr.CaseCamelLower(gstr.SubStr(field, 0, -2)) + `/list' }" />
                </ElFormItem>`
				continue
			}
			//name或code后缀
			if gstr.SubStr(fieldCaseCamel, -4) == `Name` || gstr.SubStr(fieldCaseCamel, -4) == `Code` {
				viewSaveRule += `
		` + field + `: [
			{ type: 'string', required: true, min: 1, max: ` + result[1] + `, trigger: 'blur', message: t('validation.between.string', { min: 1, max: ` + result[1] + ` }) },
			{ pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') }
		],`
				viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
					<ElInput v-model="saveForm.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" minlength="1" maxlength="` + result[1] + `" :show-word-limit="true" :clearable="true" />
				</ElFormItem>`
				continue
			}
			//mobile或phone后缀
			if gstr.SubStr(fieldCaseCamel, -5) == `Phone` || gstr.SubStr(fieldCaseCamel, -6) == `Mobile` {
				viewSaveRule += `
		` + field + `: [
			{ type: 'string', min: 1, max: ` + result[1] + `, trigger: 'blur', message: t('validation.between.string', { min: 1, max: ` + result[1] + ` }) },
			{ pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') }
		],`
				viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
					<ElInput v-model="saveForm.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" minlength="1" maxlength="` + result[1] + `" :show-word-limit="true" :clearable="true" />
				</ElFormItem>`
				continue
			}
			//url或link后缀
			if gstr.SubStr(fieldCaseCamel, -3) == `Url` || gstr.SubStr(fieldCaseCamel, -4) == `Link` {
				viewSaveRule += `
		` + field + `: [
			{ type: 'url', trigger: 'change', message: t('validation.url') },
			{ type: 'string', min: 1, max: ` + result[1] + `, trigger: 'change', message: t('validation.between.string', { min: 1, max: ` + result[1] + ` }) }
		],`
				viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
					<ElInput v-model="saveForm.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" minlength="1" maxlength="` + result[1] + `" :show-word-limit="true" :clearable="true" />
				</ElFormItem>`
				continue
			}
			//icon,cover或img,img_list,imgList,img_arr,imgArr或image,image_list,imageList,image_arr,imageArr等后缀
			if field == `avatar` || gstr.SubStr(fieldCaseCamel, -4) == `Icon` || gstr.SubStr(fieldCaseCamel, -5) == `Cover` || gstr.SubStr(fieldCaseCamel, -3) == `Img` || gstr.SubStr(fieldCaseCamel, -7) == `ImgList` || gstr.SubStr(fieldCaseCamel, -6) == `ImgArr` || gstr.SubStr(fieldCaseCamel, -5) == `Image` || gstr.SubStr(fieldCaseCamel, -9) == `ImageList` || gstr.SubStr(fieldCaseCamel, -8) == `ImageArr` {
				if column[`Type`].String() == `json` || column[`Type`].String() == `text` {
					viewSaveRule += `
		` + field + `: [
            { type: 'array', trigger: 'change', defaultField: { type: 'url', message: t('validation.url') }, message: t('validation.upload') },
            { type: 'array', min: 1, trigger: 'change', message: t('validation.min.upload', { min: 1 }) },
            { type: 'array', max: 10, trigger: 'change', message: t('validation.max.upload', { max: 10 }) }
        ],`
					viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
					<MyUpload v-model="saveForm.data.` + field + `" accept="image/*" :multiple="true" />
				</ElFormItem>`
				} else {
					viewSaveRule += `
		` + field + `: [
			{ type: 'url', trigger: 'change', message: t('validation.upload') },
			{ type: 'string', min: 1, max: ` + result[1] + `, trigger: 'blur', message: t('validation.between.string', { min: 1, max: ` + result[1] + ` }) }
        ],`
					viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
                    <MyUpload v-model="saveForm.data.` + field + `" accept="image/*" />
                </ElFormItem>`
				}
				continue
			}
			//video,video_list,videoList,video_arr,videoArr等后缀
			if gstr.SubStr(fieldCaseCamel, -5) == `Video` || gstr.SubStr(fieldCaseCamel, -9) == `VideoList` || gstr.SubStr(fieldCaseCamel, -8) == `VideoArr` {
				if column[`Type`].String() == `json` || column[`Type`].String() == `text` {
					viewSaveRule += `
		` + field + `: [
            { type: 'array', trigger: 'change', defaultField: { type: 'url', message: t('validation.url') }, message: t('validation.upload') },
            { type: 'array', min: 1, trigger: 'change', message: t('validation.min.upload', { min: 1 }) },
            { type: 'array', max: 10, trigger: 'change', message: t('validation.max.upload', { max: 10 }) }
        ],`
					viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
					<MyUpload v-model="saveForm.data.` + field + `" accept="video/*" :isImage="false" :multiple="true" />
				</ElFormItem>`
				} else {
					viewSaveRule += `
		` + field + `: [
			{ type: 'url', trigger: 'change', message: t('validation.upload') },
			{ type: 'string', min: 1, max: ` + result[1] + `, trigger: 'blur', message: t('validation.between.string', { min: 1, max: ` + result[1] + ` }) }
        ],`
					viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
                    <MyUpload v-model="saveForm.data.` + field + `" accept="video/*" :isImage="false" />
                </ElFormItem>`
				}
				continue
			}
			//Ip后缀
			if gstr.SubStr(fieldCaseCamel, -2) == `Ip` {
				viewSaveRule += `
		` + field + `: [
			{ type: 'string', min: 1, max: ` + result[1] + `, trigger: 'blur', message: t('validation.between.string', { min: 1, max: ` + result[1] + ` }) }
		],`
				viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
					<ElInput v-model="saveForm.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" minlength="1" maxlength="` + result[1] + `" :show-word-limit="true" :clearable="true" />
				</ElFormItem>`
				continue
			}
			//status或type后缀
			if field == `gender` || gstr.SubStr(fieldCaseCamel, -6) == `Status` || gstr.SubStr(fieldCaseCamel, -4) == `Type` {
				statusList := MyGenStatusList(comment)
				statusArr := make([]string, len(statusList))
				for index, status := range statusList {
					statusArr[index] = status[0]
				}
				statusStr := gstr.Join(statusArr, `, `)
				viewSaveRule += `
		` + field + `: [
			{ type: 'enum', enum: [` + statusStr + `], trigger: 'change', message: t('validation.select') }
		],`
				viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">`
				//超过5个状态用select组件，小于5个用radio组件
				if len(statusArr) > 5 {
					viewSaveField += `
					<ElSelectV2 v-model="saveForm.data.` + field + `" :options="tm('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.status.` + field + `')" :placeholder="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :clearable="true" />`
				} else {
					viewSaveField += `
					<ElRadioGroup v-model="saveForm.data.` + field + `">
                        <ElRadio v-for="(item, key) in tm('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.status.` + field + `') as any" :key="key" :label="item.value">
                            {{ item.label }}
                        </ElRadio>
                    </ElRadioGroup>`
				}
				viewSaveField += `
				</ElFormItem>`
				continue
			}
			//remark后缀
			if gstr.SubStr(fieldCaseCamel, -6) == `Remark` {
				viewSaveRule += `
		` + field + `: [
			{ type: 'string', min: 1, max: ` + result[1] + `, trigger: 'blur', message: t('validation.between.string', { min: 1, max: ` + result[1] + ` }) },
		],`
				viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
					<ElInput v-model="saveForm.data.` + field + `" type="textarea" :autosize="{ minRows: 3 }" />
				</ElFormItem>`
				continue
			}
			//is_前缀
			if gstr.SubStr(fieldCaseSnake, 0, 3) == `is_` {
				viewSaveRule += `
		` + field + `: [
            { type: 'enum', enum: [0, 1], trigger: 'change', message: t('validation.select') }
        ],`
				viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
                    <ElSwitch v-model="saveForm.data.` + field + `" :active-value="1" :inactive-value="0" :inline-prompt="true" :active-text="t('common.yes')" :inactive-text="t('common.no')" style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success);" />
                </ElFormItem>`
				continue
			}
			//int类型
			if gstr.Pos(column[`Type`].String(), `int`) != -1 {
				viewSaveRule += `
		` + field + `: [
			{ type: 'integer', trigger: 'change', message: '' }
		],`
				viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
					<ElInputNumber v-model="saveForm.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :controls="false"/>
				</ElFormItem>`
				continue
			}
			//float类型
			if gstr.Pos(column[`Type`].String(), `decimal`) != -1 || gstr.Pos(column[`Type`].String(), `double`) != -1 || gstr.Pos(column[`Type`].String(), `float`) != -1 {
				viewSaveRule += `
		` + field + `: [
			{ type: 'float', trigger: 'change', message: '' }
		],`
				viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
					<ElInputNumber v-model="saveForm.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :precision="2" :controls="false"/>
				</ElFormItem>`
				continue
			}
			//varchar类型
			if gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
				viewSaveRule += `
		` + field + `: [
			{ type: 'string', min: 1, max: ` + result[1] + `, trigger: 'blur', message: t('validation.between.string', { min: 1, max: ` + result[1] + ` }) },
		],`
				viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
					<ElInput v-model="saveForm.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" minlength="1" maxlength="` + result[1] + `" :show-word-limit="true" :clearable="true" />
				</ElFormItem>`
				continue
			}
			//char类型
			if gstr.Pos(column[`Type`].String(), `char`) != -1 {
				viewSaveRule += `
		` + field + `: [
			{ type: 'string', min: ` + result[1] + `, max: ` + result[1] + `, trigger: 'blur', message: t('validation.between.string', { min: ` + result[1] + `, max: ` + result[1] + ` }) },
		],`
				viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
					<ElInput v-model="saveForm.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" minlength="` + result[1] + `" maxlength="` + result[1] + `" :show-word-limit="true" :clearable="true" />
				</ElFormItem>`
				continue
			}
			//json类型
			if column[`Type`].String() == `json` {
				viewSaveRule += `
		` + field + `: [
			{
				type: 'object',
				/* fields: {
					xxxx: { type: 'string', min: 1, message: 'xxxx' + t('validation.min.string', { min: 1 }) }
				}, */
				transform(value: any) {
					if (value === '' || value === null || value === undefined) {
						return undefined
					}
					try {
						return JSON.parse(value)
					} catch (e) {
						return value
					}
				},
				trigger: 'blur',
				message: t('validation.json')
			},
		],`
				viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
					<ElInput v-model="saveForm.data.` + field + `" type="textarea" :autosize="{ minRows: 3 }" />
				</ElFormItem>`
				continue
			}
			//datetime和timestamp类型
			if gstr.Pos(column[`Type`].String(), `datetime`) != -1 || gstr.Pos(column[`Type`].String(), `timestamp`) != -1 {
				viewSaveRule += `
		` + field + `: [],`
				viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
					<ElDatePicker v-model="saveForm.data.` + field + `" type="datetime" :placeholder="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" format="YYYY-MM-DD HH:mm:ss" value-format="YYYY-MM-DD HH:mm:ss" />
				</ElFormItem>`
				continue
			}
			//date类型
			if gstr.Pos(column[`Type`].String(), `date`) != -1 {
				viewSaveRule += `
		` + field + `: [],`
				viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
					<ElDatePicker v-model="saveForm.data.` + field + `" type="date" :placeholder="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" format="YYYY-MM-DD" value-format="YYYY-MM-DD" />
				</ElFormItem>`
				continue
			}
			//默认处理
			viewSaveRule += `
		` + field + `: [],`
			viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
					<ElInput v-model="saveForm.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLower + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :show-word-limit="true" :clearable="true" />
				</ElFormItem>`
		}
	}

	tplView := `<script setup lang="ts">`
	if passwordField != `` {
		tplView += `
import md5 from 'js-md5'
`
	}
	tplView += `
const { t, tm } = useI18n()

const saveCommon = inject('saveCommon') as { visible: boolean, title: string, data: { [propName: string]: any } }
const listCommon = inject('listCommon') as { ref: any }

const saveForm = reactive({
	ref: null as any,
	loading: false,
	data: {
		...saveCommon.data
	} as { [propName: string]: any },
	rules: {` + viewSaveRule + `
	} as any,
	submit: () => {
		saveForm.ref.validate(async (valid: boolean) => {
			if (!valid) {
				return false
			}
			saveForm.loading = true
			const param = removeEmptyOfObj(saveForm.data, false)`
	if passwordField != `` {
		tplView += `
            param.` + passwordField + ` ? param.` + passwordField + ` = md5(param.` + passwordField + `) : delete param.` + passwordField
	}
	tplView += `
			try {
				if (param?.idArr?.length > 0) {
					await request(t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseCamelLower + `/` + tpl.TableNameCaseCamelLower + `/update', param, true)
				} else {
					await request(t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseCamelLower + `/` + tpl.TableNameCaseCamelLower + `/create', param, true)
				}
				listCommon.ref.getList(true)
				saveCommon.visible = false
			} catch (error) { }
			saveForm.loading = false
		})
	}
})

const saveDrawer = reactive({
	ref: null as any,
	size: useSettingStore().saveDrawer.size,
	beforeClose: (done: Function) => {
		if (useSettingStore().saveDrawer.isTipClose) {
			ElMessageBox.confirm('', {
				type: 'info',
				title: t('common.tip.configExit'),
				center: true,
				showClose: false,
			}).then(() => {
				done()
			}).catch(() => { })
		} else {
			done()
		}
	},
	buttonClose: () => {
		//saveCommon.visible = false
		saveDrawer.ref.handleClose()    //会触发beforeClose
	}
})
</script>

<template>
	<ElDrawer class="save-drawer" :ref="(el: any) => { saveDrawer.ref = el }" v-model="saveCommon.visible"
		:title="saveCommon.title" :size="saveDrawer.size" :before-close="saveDrawer.beforeClose">
		<ElScrollbar>
			<ElForm :ref="(el: any) => { saveForm.ref = el }" :model="saveForm.data" :rules="saveForm.rules"
				label-width="auto" :status-icon="true" :scroll-to-error="true">` + viewSaveField + `
			</ElForm>
		</ElScrollbar>
		<template #footer>
			<ElButton @click="saveDrawer.buttonClose">{{ t('common.cancel') }}</ElButton>
			<ElButton type="primary" @click="saveForm.submit" :loading="saveForm.loading">
				{{ t('common.save') }}
			</ElButton>
		</template>
	</ElDrawer>
</template>`

	gfile.PutContents(saveFile, tplView)
}

// 视图模板I18n生成
func MyGenTplViewI18n(ctx context.Context, option *MyGenOption, tpl *MyGenTpl) {
	saveFile := gfile.SelfDir() + `/../dev/` + option.SceneCode + `/src/i18n/language/zh-cn/` + tpl.ModuleDirCaseCamelLower + `/` + tpl.TableNameCaseCamelLower + `.ts`
	if !option.IsCover && gfile.IsFile(saveFile) {
		return
	}

	viewI18nName := ``
	viewI18nStatus := ``
	viewI18nTip := ``
	for _, column := range tpl.TableColumnList {
		field := column[`Field`].String()
		fieldCaseCamel := gstr.CaseCamel(field)
		comment := gstr.Trim(gstr.ReplaceByArray(column[`Comment`].String(), g.SliceStr{
			"\n", " ",
			"\r", " ",
		}))

		switch field {
		case `deletedAt`, `deleted_at`, `createdAt`, `created_at`, `updatedAt`, `updated_at`:
		case `sort`, `weight`:
			viewI18nName += `
		` + field + `: '` + comment + `',`
			viewI18nTip += `
		` + field + `: '` + comment + `',`
		default:
			//主键
			if column[`Key`].String() == `PRI` && column[`Extra`].String() == `auto_increment` {
				continue
			}
			viewI18nName += `
		` + field + `: '` + comment + `',`

			//status或type后缀
			if field == `gender` || gstr.SubStr(fieldCaseCamel, -6) == `Status` || gstr.SubStr(fieldCaseCamel, -4) == `Type` {
				statusList := MyGenStatusList(comment)
				viewI18nStatus += `
		` + field + `: [`
				for _, status := range statusList {
					viewI18nStatus += `
			{ label: '` + status[1] + `', value: ` + status[0] + ` },`
				}
				viewI18nStatus += `
		],`
			}
		}
	}
	tplView := `export default {
    name:{` + viewI18nName + `
    },
    status: {` + viewI18nStatus + `
    },
    tip: {` + viewI18nTip + `
    }
}`

	gfile.PutContents(saveFile, tplView)
}

// 前端路由生成
func MyGenTplViewRouter(ctx context.Context, option *MyGenOption, tpl *MyGenTpl) {
	saveFile := gfile.SelfDir() + `/../dev/` + option.SceneCode + `/src/router/index.ts`

	tplView := gfile.GetContents(saveFile)

	path := `/` + tpl.ModuleDirCaseCamelLower + `/` + tpl.TableNameCaseCamelLower
	replaceStr := `{
                path: '` + path + `',
                component: async () => {
                    const component = await import('@/views` + path + `/Index.vue')
                    component.default.name = '` + path + `'
                    return component
                },
                meta: { isAuth: true, keepAlive: true, componentName: '` + path + `' }
            },`

	if gstr.Pos(tplView, `'`+path+`'`) == -1 { //路由不存在时新增
		tplView = gstr.Replace(tplView, `/*--------前端路由自动代码生成锚点（不允许修改和删除，否则将不能自动生成路由）--------*/`, replaceStr+`
            /*--------前端路由自动代码生成锚点（不允许修改和删除，否则将不能自动生成路由）--------*/`)
	} else { //路由已存在则替换
		tplView, _ = gregex.ReplaceString(`\{
                path: '`+path+`',[\s\S]*'`+path+`' \}
            \},`, replaceStr, tplView)
	}
	gfile.PutContents(saveFile, tplView)

	daoAuth.Menu.MyGenMenu(ctx, tpl.SceneId, path, option.CommonName, tpl.TableNameCaseCamel) // 数据库权限菜单处理
}
