package cmd

import (
	daoAuth "api/internal/dao/auth"
	"api/internal/service"
	"api/internal/utils"
	"context"
	"fmt"
	"os/exec"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

/*
后台常用生成示例：./main myGen -sceneCode=platform -dbGroup=default -dbTable=auth_test -removePrefix=auth_ -moduleDir=auth -commonName=权限管理/测试 -isList=1 -isCount=1 -isInfo=1 -isCreate=1 -isUpdate=1 -isDelete=1 -isApi=1 -isAuthAction=1 -isView=1 -isCover=0
APP常用生成示例：./main myGen -sceneCode=app -dbGroup=xxxx -dbTable=user -removePrefix= -moduleDir=xxxx/user -commonName=用户 -isList=1 -isCount=0 -isInfo=1 -isCreate=0 -isUpdate=0 -isDelete=0 -isApi=1 -isAuthAction=0 -isView=0 -isCover=0

强烈建议搭配Git使用
表字段命名需要遵守以下规则，否则只会根据字段类型做默认处理
主键必须在第一个字段。否则需要在dao层重写PrimaryKey方法返回主键字段
表内尽量根据表名设置xxxxId主键和xxxxName名称两个字段（作用1：常用于前端部分组件，如MySelect.vue组件；作用2：当其它表存在与该表主键同名的关联字段时，会自动生成联表查询代码）
每个字段都必须有注释。以下符号[\n\r.。:：(（]之前的部分或整个注释，将作为字段名称使用

	部分常用字段：
		密码 		命名：password|passwd；类型：char(32)；注意：password|passwd,salt同时存在时，有特殊处理
		加密盐 		命名：salt；类型：char(8)；注意：password|passwd,salt同时存在时，(才)有特殊处理
		父级		命名：pid；类型：int等类型；注意：pid,level,idPath|id_path同时存在时，有特殊处理
		层级		命名：level；类型：int等类型；注意：pid,level,idPath|id_path同时存在时，(才)有特殊处理
		层级路径	命名：idPath|id_path；类型：varchar；注意：pid,level,idPath|id_path同时存在时，(才)有特殊处理
		排序		命名：sort；类型：int等类型；
		权重		命名：weight；类型：int等类型；
		性别		命名：gender；类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（性别：0未设置 1男 2女）
		头像		命名：avatar；类型：varchar；
	其他类型字段：
		名称		命名：name后缀；类型：varchar；
		标识		命名：code后缀；类型：varchar；
		手机		命名：mobile,phone后缀；类型：varchar；
		链接		命名：url,link后缀；类型：varchar；
		关联ID		命名：id后缀；类型：int等类型；
		图片		命名：icon,cover,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；类型：单图片varchar，多图片json或text
		视频		命名：video,video_list,videoList,video_arr,videoArr等后缀；类型：单视频varchar，多视频json或text
		数组		命名：list,arr等后缀；类型：json或text；
		IP			命名：IP后缀；类型：varchar；
		(富)文本	命名：remark,desc,msg,message,intro,content后缀；类型：varchar或text；前端生成表单组件：textarea文本输入框或tinymce富文本编辑器
		状态		命名：status后缀；类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回）
		类型		命名：type后缀；类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（类型：0安卓 1苹果）
		是否		命名：is_前缀；类型：int等类型；注释：示例（停用：0否 1是）
		开始时间	命名：start_前缀；类型：timestamp或datetime或date；
		结束时间	命名：end_前缀；类型：timestamp或datetime或date；
*/
type MyGenOption struct {
	SceneCode    string `json:"sceneCode"`    //场景标识，必须在数据库表auth_scene已存在。示例：platform
	DbGroup      string `json:"dbGroup"`      //db分组。示例：default
	DbTable      string `json:"dbTable"`      //db表。示例：auth_test
	RemovePrefix string `json:"removePrefix"` //要删除的db表前缀。必须和hack/config.yaml内removePrefix保持一致，示例：auth_
	ModuleDir    string `json:"moduleDir"`    //模块目录，支持多目录。必须和hack/config.yaml内daoPath的后面部分保持一致，示例：auth，xxxx/user
	CommonName   string `json:"commonName"`   //公共名称，将同时在swagger文档Tag标签，权限菜单和权限操作中使用。示例：用户，权限管理/测试
	IsList       bool   `json:"isList" `      //是否生成列表接口(0和no为false，1和yes为true)
	IsCount      bool   `json:"isCount" `     //列表接口是否返回总数
	IsInfo       bool   `json:"isInfo" `      //是否生成详情接口
	IsCreate     bool   `json:"isCreate"`     //是否生成创建接口
	IsUpdate     bool   `json:"isUpdate"`     //是否生成更新接口
	IsDelete     bool   `json:"isDelete"`     //是否生成删除接口
	IsApi        bool   `json:"isApi"`        //是否生成后端接口文件
	IsAuthAction bool   `json:"isAuthAction"` //是否判断操作权限，如是，则同时会生成操作权限
	IsView       bool   `json:"isView"`       //是否生成前端视图文件
	IsCover      bool   `json:"isCover"`      //是否覆盖原文件(设置为true时，建议与git一起使用，防止代码覆盖风险)
}

type MyGenTpl struct {
	TableColumnList                gdb.Result //表字段详情
	SceneName                      string     //场景名称
	SceneId                        int        //场景ID
	TableNameCaseCamel             string     //去除前缀表名（大驼峰）
	TableNameCaseCamelLower        string     //去除前缀表名（小驼峰）
	TableNameCaseSnake             string     //去除前缀表名（蛇形）
	ModuleDirCaseCamel             string     //模块目录（大驼峰，/会被去除）
	ModuleDirCaseCamelLower        string     //模块目录（小驼峰，/会被保留）
	ModuleDirCaseCamelLowerReplace string     //模块目录（小驼峰，/会被替换成.）
	LogicStructName                string     //logic层结构体名称，也是权限操作前缀（大驼峰，由ModuleDirCaseCamel+TableNameCaseCamel组成。命名原因：gf gen service只支持logic单层目录，可能导致service层重名）
	PrimaryKey                     string     //表主键
	DeletedField                   string     //表删除时间字段
	UpdatedField                   string     //表更新时间字段
	CreatedField                   string     //表创建时间字段
	LabelField                     string     //dao层label对应的字段(常用于前端组件)
	// 以下字段用于对某些表字段做特殊处理
	RelTableMap    map[string]RelTableItem //当id后缀字段，且关联表能确定时，会自动生成联表查询代码
	PasswordHandle struct {                //password|passwd,salt同时存在时，需特殊处理
		IsCoexist     bool   //是否同时存在
		PasswordField string //密码字段
		SaltField     string //加密盐字段
	}
	PidHandle struct { //pid,level,idPath|id_path同时存在时，需特殊处理
		IsCoexist   bool   //是否同时存在
		PidField    string //父级字段
		LevelField  string //层级字段gstr.Replace(tpl.ModuleDirCaseCamelLower, `/`, `.`)
		IdPathField string //层级路径字段
		SortField   string //排序字段
	}
}

type RelTableItem struct {
	IsExistRelTableDao         bool   //是否存在关联表dao层
	RelDaoDir                  string //关联表dao层目录
	RelDaoDirCaseCamel         string //关联表dao层目录（大驼峰，/会被去除）
	RelDaoDirCaseCamelLower    string //关联表dao层目录（小驼峰，/会被保留）
	IsSameDir                  bool   //关联表dao层是否与当前生成dao层在相同目录下
	RelTableNameCaseCamel      string //关联表名（大驼峰）
	RelTableNameCaseCamelLower string //关联表名（小驼峰）
	RelTableNameCaseSnake      string //关联表名（蛇形）
	RelNameFieldCaseCamelLower string //关联表名称字段（小驼峰）
	RelNameFieldName           string //关联表名称字段名称
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
	// 公共名称，将同时在swagger文档Tag标签，权限菜单和权限操作中使用。示例：场景
	_, ok = optionMap[`commonName`]
	if !ok {
		option.CommonName = gcmd.Scan("> 请输入公共名称，将同时在swagger文档Tag标签，权限菜单和权限操作中使用:\n")
	}
	for {
		if option.CommonName != `` {
			break
		}
		option.CommonName = gcmd.Scan("> 请输入公共名称，将同时在swagger文档Tag标签，权限菜单和权限操作中使用:\n")
	}
noAllRestart:
	// 是否生成列表接口
	isList, ok := optionMap[`isList`]
	if !ok {
		isList = gcmd.Scan("> 是否生成列表接口，默认(yes):\n")
	}
isListEnd:
	for {
		switch isList {
		case ``, `1`, `yes`:
			option.IsList = true
			break isListEnd
		case `0`, `no`:
			option.IsList = false
			break isListEnd
		default:
			isList = gcmd.Scan("> 输入错误，请重新输入，是否生成列表接口，默认(yes):\n")
		}
	}
	// 列表接口是否返回总数
	isCount, ok := optionMap[`isCount`]
	if !ok {
		isCount = gcmd.Scan("> 列表接口是否返回总数，默认(yes):\n")
	}
isCountEnd:
	for {
		switch isCount {
		case ``, `1`, `yes`:
			option.IsCount = true
			break isCountEnd
		case `0`, `no`:
			option.IsCount = false
			break isCountEnd
		default:
			isCount = gcmd.Scan("> 输入错误，请重新输入，列表接口是否返回总数，默认(yes):\n")
		}
	}
	// 是否生成详情接口
	isInfo, ok := optionMap[`isInfo`]
	if !ok {
		isInfo = gcmd.Scan("> 是否生成详情接口，默认(yes):\n")
	}
isInfoEnd:
	for {
		switch isInfo {
		case ``, `1`, `yes`:
			option.IsInfo = true
			break isInfoEnd
		case `0`, `no`:
			option.IsInfo = false
			break isInfoEnd
		default:
			isInfo = gcmd.Scan("> 输入错误，请重新输入，是否生成详情接口，默认(yes):\n")
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
		case ``, `1`, `yes`:
			option.IsCreate = true
			break isCreateEnd
		case `0`, `no`:
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
		case ``, `1`, `yes`:
			option.IsUpdate = true
			break isUpdateEnd
		case `0`, `no`:
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
		case ``, `1`, `yes`:
			option.IsDelete = true
			break isDeleteEnd
		case `0`, `no`:
			option.IsDelete = false
			break isDeleteEnd
		default:
			isDelete = gcmd.Scan("> 输入错误，请重新输入，是否生成删除接口，默认(yes):\n")
		}
	}
	if !(option.IsList || option.IsInfo || option.IsCreate || option.IsUpdate || option.IsDelete) {
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
		case ``, `1`, `yes`:
			option.IsApi = true
			break isApiEnd
		case `0`, `no`:
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
			case ``, `1`, `yes`:
				option.IsAuthAction = true
				break isAuthActionEnd
			case `0`, `no`:
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
		case ``, `1`, `yes`:
			option.IsView = true
			break isViewEnd
		case `0`, `no`:
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
		case `1`, `yes`:
			option.IsCover = true
			break isCoverEnd
		case ``, `0`, `no`:
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
		TableColumnList:         tableColumnList,
		SceneName:               sceneInfo[daoAuth.Scene.Columns().SceneName].String(),
		SceneId:                 sceneInfo[daoAuth.Scene.Columns().SceneId].Int(),
		TableNameCaseCamel:      gstr.CaseCamel(tableName),
		TableNameCaseCamelLower: gstr.CaseCamelLower(tableName),
		TableNameCaseSnake:      gstr.CaseSnakeFirstUpper(tableName),
		RelTableMap:             map[string]RelTableItem{},
	}
	moduleDirArr := gstr.Split(option.ModuleDir, `/`)
	moduleDirCaseCamelArr := []string{}
	moduleDirCaseCamelLowerArr := []string{}
	for _, v := range moduleDirArr {
		moduleDirCaseCamelArr = append(moduleDirCaseCamelArr, gstr.CaseCamel(v))
		moduleDirCaseCamelLowerArr = append(moduleDirCaseCamelLowerArr, gstr.CaseCamelLower(v))
	}
	tpl.ModuleDirCaseCamel = gstr.Join(moduleDirCaseCamelArr, ``)
	tpl.ModuleDirCaseCamelLower = gstr.Join(moduleDirCaseCamelLowerArr, `/`)
	tpl.ModuleDirCaseCamelLowerReplace = gstr.Replace(tpl.ModuleDirCaseCamelLower, `/`, `.`)
	if gstr.CaseSnakeFirstUpper(moduleDirCaseCamelArr[len(moduleDirCaseCamelArr)-1]) == tpl.TableNameCaseSnake {
		tpl.LogicStructName = tpl.ModuleDirCaseCamel
	} else {
		tpl.LogicStructName = tpl.ModuleDirCaseCamel + tpl.TableNameCaseCamel
	}

	fieldArr := make([]string, len(tpl.TableColumnList))
	fieldCaseCamelArr := make([]string, len(tpl.TableColumnList))
	for index, column := range tpl.TableColumnList {
		field := column[`Field`].String()
		fieldCaseCamel := gstr.CaseCamel(field)
		tmp, _ := gregex.MatchString(`[^\n\r\.。:：\(（]*`, column[`Comment`].String())
		fieldName := gstr.Trim(tmp[0])
		fieldArr[index] = field
		fieldCaseCamelArr[index] = fieldCaseCamel

		if garray.NewStrArrayFrom([]string{`DeletedAt`, `DeleteAt`, `DeletedTime`, `DeleteTime`}).Contains(fieldCaseCamel) {
			tpl.DeletedField = field
			continue
		}
		if garray.NewStrArrayFrom([]string{`UpdatedAt`, `UpdateAt`, `UpdatedTime`, `UpdateTime`}).Contains(fieldCaseCamel) {
			tpl.UpdatedField = field
			continue
		}
		if garray.NewStrArrayFrom([]string{`CreatedAt`, `CreateAt`, `CreatedTime`, `CreateTime`}).Contains(fieldCaseCamel) {
			tpl.CreatedField = field
			continue
		}
		//主键
		if column[`Key`].String() == `PRI` && column[`Extra`].String() == `auto_increment` {
			tpl.PrimaryKey = field
			continue
		}
		//password|passwd
		if garray.NewStrArrayFrom([]string{`password`, `passwd`}).Contains(field) && column[`Type`].String() == `char(32)` {
			tpl.PasswordHandle.PasswordField = field
			continue
		}
		//salt
		if garray.NewStrArrayFrom([]string{`salt`}).Contains(field) && column[`Type`].String() == `char(8)` {
			tpl.PasswordHandle.SaltField = field
			continue
		}
		//pid
		if garray.NewStrArrayFrom([]string{`pid`}).Contains(field) && gstr.Pos(column[`Type`].String(), `int`) != -1 {
			tpl.PidHandle.PidField = field
			continue
		}
		//level
		if garray.NewStrArrayFrom([]string{`level`}).Contains(field) && gstr.Pos(column[`Type`].String(), `int`) != -1 {
			tpl.PidHandle.LevelField = field
			continue
		}
		//idPath|id_path
		if fieldCaseCamel == `IdPath` && gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
			tpl.PidHandle.IdPathField = field
			continue
		}
		//sort
		if garray.NewStrArrayFrom([]string{`sort`}).Contains(field) && gstr.Pos(column[`Type`].String(), `int`) != -1 {
			tpl.PidHandle.SortField = field
			continue
		}
		//id后缀
		if gstr.SubStr(fieldCaseCamel, -2) == `Id` && gstr.Pos(column[`Type`].String(), `int`) != -1 {
			relTableNameRaw := gstr.SubStr(field, 0, -2)
			relTableItem := RelTableItem{
				IsExistRelTableDao:         false,
				RelDaoDir:                  ``,
				IsSameDir:                  false,
				RelTableNameCaseCamel:      gstr.CaseCamel(relTableNameRaw),
				RelTableNameCaseCamelLower: gstr.CaseCamelLower(relTableNameRaw),
				RelTableNameCaseSnake:      gstr.CaseSnakeFirstUpper(relTableNameRaw),
				RelNameFieldCaseCamelLower: gstr.CaseCamelLower(relTableNameRaw) + `Name`,
				RelNameFieldName:           fieldName,
			}
			if gstr.ToUpper(gstr.SubStr(relTableItem.RelNameFieldName, -2)) == `ID` {
				relTableItem.RelNameFieldName = gstr.SubStr(relTableItem.RelNameFieldName, 0, -2)
			}

			fileArr, _ := gfile.ScanDirFile(gfile.SelfDir()+`/internal/dao/`, relTableItem.RelTableNameCaseSnake+`\.go`, true)
			relDaoDirList := []string{}
			for _, v := range fileArr {
				if gstr.Count(v, `/internal/`) == 1 {
					relDaoDirTmp := gstr.Replace(v, gfile.SelfDir()+`/internal/dao/`, ``)
					relDaoDirTmp = gstr.Replace(relDaoDirTmp, `/`+relTableItem.RelTableNameCaseSnake+`.go`, ``)
					if relDaoDirTmp == tpl.ModuleDirCaseCamelLower { //关联dao层在相同目录下时，直接返回
						relTableItem.IsExistRelTableDao = true
						relTableItem.RelDaoDir = relDaoDirTmp
						relTableItem.IsSameDir = true
						break
					}
					relDaoDirList = append(relDaoDirList, relDaoDirTmp)
				}
			}
			if !relTableItem.IsExistRelTableDao {
				if len(relDaoDirList) == 1 {
					relTableItem.IsExistRelTableDao = true
					relTableItem.RelDaoDir = relDaoDirList[0]
				} else {
					// gstr.SubStr(tpl.ModuleDirCaseCamelLower, 0, gstr.PosR(tpl.ModuleDirCaseCamelLower, `/`))
					relDaoCount := 0 //计算dao层同层目录的数量
					for _, v := range relDaoDirList {
						if gstr.Count(v, `/`) == gstr.Count(tpl.ModuleDirCaseCamelLower, `/`) {
							relDaoCount++
							relTableItem.RelDaoDir = v
						}
					}
					if relDaoCount == 1 { //当同层目录只存在一个关联dao时，以这个为准
						relTableItem.IsExistRelTableDao = true
					}
				}
			}
			relDaoDirArr := gstr.Split(relTableItem.RelDaoDir, `/`)
			relDaoDirCaseCamelArr := []string{}
			relDaoDirCaseCamelLowerArr := []string{}
			for _, v := range relDaoDirArr {
				relDaoDirCaseCamelArr = append(relDaoDirCaseCamelArr, gstr.CaseCamel(v))
				relDaoDirCaseCamelLowerArr = append(relDaoDirCaseCamelLowerArr, gstr.CaseCamelLower(v))
			}
			relTableItem.RelDaoDirCaseCamel = gstr.Join(relDaoDirCaseCamelArr, ``)
			relTableItem.RelDaoDirCaseCamelLower = gstr.Join(relDaoDirCaseCamelLowerArr, `/`)
			tpl.RelTableMap[field] = relTableItem
			continue
		}
	}

	fieldCaseCamelArrG := garray.NewStrArrayFrom(fieldCaseCamelArr)
	// 根据name字段优先级排序
	nameFieldList := []string{
		tpl.TableNameCaseCamel + `Name`,
		gstr.SubStr(gstr.CaseCamel(tpl.PrimaryKey), 0, -2) + `Name`,
		`Name`,
		`Phone`,
		`Account`,
		`Nickname`,
		`Title`,
	}
	for _, v := range nameFieldList {
		index := fieldCaseCamelArrG.Search(v)
		if index != -1 {
			tpl.LabelField = fieldArr[index]
			break
		}
	}

	if tpl.PasswordHandle.PasswordField != `` && tpl.PasswordHandle.SaltField != `` {
		tpl.PasswordHandle.IsCoexist = true
	}

	if tpl.PidHandle.PidField != `` && tpl.PidHandle.LevelField != `` && tpl.PidHandle.IdPathField != `` {
		tpl.PidHandle.IsCoexist = true
	}
	return
}

// 自动生成操作
func MyGenAction(ctx context.Context, sceneId int, actionCode string, actionName string) {
	daoThis := daoAuth.Action
	columnsThis := daoAuth.Action.Columns()
	actionName = gstr.Replace(actionName, `/`, `-`)

	idVar, _ := daoThis.ParseDbCtx(ctx).Where(columnsThis.ActionCode, actionCode).Value(daoThis.PrimaryKey())
	id := idVar.Int64()
	if id == 0 {
		id, _ = daoThis.ParseDbCtx(ctx).Data(map[string]interface{}{
			columnsThis.ActionCode: actionCode,
			columnsThis.ActionName: actionName,
		}).InsertAndGetId()
	} else {
		daoThis.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), id).Data(columnsThis.ActionName, actionName).Update()
	}
	daoAuth.ActionRelToScene.ParseDbCtx(ctx).Data(map[string]interface{}{
		daoAuth.ActionRelToScene.Columns().ActionId: id,
		daoAuth.ActionRelToScene.Columns().SceneId:  sceneId,
	}).Save()
}

// 自动生成菜单
func MyGenMenu(ctx context.Context, sceneId int, menuUrl string, menuName string, menuNameOfEn string) {
	daoThis := daoAuth.Menu
	columnsThis := daoAuth.Menu.Columns()
	menuNameArr := gstr.Split(menuName, `/`)

	var pid int64 = 0
	for _, v := range menuNameArr[:len(menuNameArr)-1] {
		pidVar, _ := daoThis.ParseDbCtx(ctx).Where(columnsThis.SceneId, sceneId).Where(columnsThis.MenuName, v).Value(daoThis.PrimaryKey())
		if pidVar.Int() == 0 {
			pid, _ = service.AuthMenu().Create(ctx, g.Map{
				columnsThis.SceneId:   sceneId,
				columnsThis.Pid:       pid,
				columnsThis.MenuName:  v,
				columnsThis.MenuIcon:  `AutoiconEpLink`,
				columnsThis.MenuUrl:   ``,
				columnsThis.ExtraData: `{"i18n": {"title": {"en": "", "zh-cn": "` + v + `"}}}`,
			})
		} else {
			pid = pidVar.Int64()
		}
	}

	menuName = menuNameArr[len(menuNameArr)-1]
	idVar, _ := daoThis.ParseDbCtx(ctx).Where(columnsThis.SceneId, sceneId).Where(columnsThis.MenuUrl, menuUrl).Value(daoThis.PrimaryKey())
	id := idVar.Int()
	if id == 0 {
		service.AuthMenu().Create(ctx, g.Map{
			columnsThis.SceneId:   sceneId,
			columnsThis.Pid:       pid,
			columnsThis.MenuName:  menuName,
			columnsThis.MenuIcon:  `AutoiconEpLink`,
			columnsThis.MenuUrl:   menuUrl,
			columnsThis.ExtraData: `{"i18n": {"title": {"en": "` + menuNameOfEn + `", "zh-cn": "` + menuName + `"}}}`,
		})
	} else {
		service.AuthMenu().Update(ctx, g.Map{daoThis.PrimaryKey(): id}, g.Map{
			columnsThis.MenuName:  menuName,
			columnsThis.Pid:       pid,
			columnsThis.ExtraData: `{"i18n": {"title": {"en": "` + menuNameOfEn + `", "zh-cn": "` + menuName + `"}}}`,
		})
	}
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

	daoParseInsert := ``
	daoHookInsert := ``
	daoParseUpdate := ``
	daoHookUpdateBefore := ``
	daoHookUpdateAfter := ``
	daoParseField := ``
	daoHookSelect := ``
	daoParseFilter := ``
	daoParseOrder := ``
	daoParseJoin := ``
	daoFunc := ``
	daoImportOtherDao := ``

	if tpl.LabelField != `` {
		daoParseFieldTmp := `
			case ` + "`label`" + `:
				m = m.Fields(daoThis.Table() + ` + "`.`" + ` + daoThis.Columns().` + gstr.CaseCamel(tpl.LabelField) + ` + ` + "` AS `" + ` + v)`
		if gstr.Pos(tplDao, daoParseFieldTmp) == -1 {
			daoParseField += daoParseFieldTmp
		}
		daoParseFilterTmp := `
			case ` + "`label`" + `:
				m = m.WhereLike(daoThis.Table()+` + "`.`" + `+daoThis.Columns().` + gstr.CaseCamel(tpl.LabelField) + `, ` + "`%`" + `+gconv.String(v)+` + "`%`" + `)`
		if gstr.Pos(tplDao, daoParseFilterTmp) == -1 {
			daoParseFilter += daoParseFilterTmp
		}
	}

	for _, column := range tpl.TableColumnList {
		field := column[`Field`].String()
		fieldCaseCamel := gstr.CaseCamel(field)
		fieldCaseSnake := gstr.CaseSnakeFirstUpper(field)

		if garray.NewStrArrayFrom([]string{`DeletedAt`, `DeleteAt`, `DeletedTime`, `DeleteTime`}).Contains(fieldCaseCamel) {
			continue
		}
		if garray.NewStrArrayFrom([]string{`UpdatedAt`, `UpdateAt`, `UpdatedTime`, `UpdateTime`}).Contains(fieldCaseCamel) {
			continue
		}
		if garray.NewStrArrayFrom([]string{`CreatedAt`, `CreateAt`, `CreatedTime`, `CreateTime`}).Contains(fieldCaseCamel) {
			daoParseFilterTmp := `
			case ` + "`timeRangeStart`" + `:
				m = m.WhereGTE(daoThis.Table()+` + "`.`" + `+daoThis.Columns().` + fieldCaseCamel + `, v)
			case ` + "`timeRangeEnd`" + `:
				m = m.WhereLTE(daoThis.Table()+` + "`.`" + `+daoThis.Columns().` + fieldCaseCamel + `, v)`
			if gstr.Pos(tplDao, daoParseFilterTmp) == -1 {
				daoParseFilter += daoParseFilterTmp
			}
			continue
		}
		//主键
		if column[`Key`].String() == `PRI` && column[`Extra`].String() == `auto_increment` {
			continue
		}
		//password|passwd
		if garray.NewStrArrayFrom([]string{`password`, `passwd`}).Contains(field) && column[`Type`].String() == `char(32)` {
			continue
		}
		//salt
		if garray.NewStrArrayFrom([]string{`salt`}).Contains(field) && column[`Type`].String() == `char(8)` && tpl.PasswordHandle.IsCoexist {
			continue
		}
		//pid
		if garray.NewStrArrayFrom([]string{`pid`}).Contains(field) && gstr.Pos(column[`Type`].String(), `int`) != -1 {
			if tpl.LabelField != `` {
				daoParseFieldTmp := `
			case ` + "`p" + gstr.CaseCamel(tpl.LabelField) + "`" + `:
				m = m.Fields(` + "`p_`" + ` + daoThis.Table() + ` + "`.`" + ` + daoThis.Columns().` + gstr.CaseCamel(tpl.LabelField) + ` + ` + "` AS `" + ` + v)
				m = daoThis.ParseJoin(` + "`p_`" + `+daoThis.Table(), joinTableArr)(m)`
				if gstr.Pos(tplDao, daoParseFieldTmp) == -1 {
					daoParseField += daoParseFieldTmp
				}
			}
			daoParseFieldTmp := `
			case ` + "`tree`" + `:
				m = m.Fields(daoThis.Table() + ` + "`.`" + ` + daoThis.PrimaryKey())
				m = m.Fields(daoThis.Table() + ` + "`.`" + ` + daoThis.Columns().` + fieldCaseCamel + `)
				m = daoThis.ParseOrder([]string{` + "`tree`" + `}, joinTableArr)(m)`
			if gstr.Pos(tplDao, daoParseFieldTmp) == -1 {
				daoParseField += daoParseFieldTmp
			}
			daoParseOrderTmp := `
			case ` + "`tree`" + `:
				m = m.OrderAsc(daoThis.Table() + ` + "`.`" + ` + daoThis.Columns().` + fieldCaseCamel + `)`
			if tpl.PidHandle.SortField != `` {
				daoParseOrderTmp += `
				m = m.OrderAsc(daoThis.Table() + ` + "`.`" + ` + daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.SortField) + `)`
			}
			daoParseOrderTmp += `
				m = m.OrderAsc(daoThis.Table() + ` + "`.`" + ` + daoThis.PrimaryKey())`
			if gstr.Pos(tplDao, daoParseOrderTmp) == -1 {
				daoParseOrder += daoParseOrderTmp
			}
			daoParseJoinTmp := `
		case ` + "`p_`" + ` + daoThis.Table():
			relTable := ` + "`p_`" + ` + daoThis.Table()
			if !garray.NewStrArrayFrom(*joinTableArr).Contains(relTable) {
				*joinTableArr = append(*joinTableArr, relTable)
				m = m.LeftJoin(daoThis.Table()+` + "` AS `" + `+relTable, relTable+` + "`.`" + `+daoThis.PrimaryKey()+` + "` = `" + `+daoThis.Table()+` + "`.`" + `+daoThis.Columns().` + fieldCaseCamel + `)
			}`
			if gstr.Pos(tplDao, daoParseJoinTmp) == -1 {
				daoParseJoin += daoParseJoinTmp
			}
			continue
		}
		//level
		if garray.NewStrArrayFrom([]string{`level`}).Contains(field) && gstr.Pos(column[`Type`].String(), `int`) != -1 && tpl.PidHandle.IsCoexist {
			daoParseOrderTmp := `
			case daoThis.Columns().` + fieldCaseCamel + `:
				m = m.Order(daoThis.Table()+` + "`.`" + `+kArr[0], kArr[1])
				m = m.OrderDesc(daoThis.Table() + ` + "`.`" + ` + daoThis.PrimaryKey())` //追加主键倒序。mysql排序字段有重复值时，分页会导致同一条数据可能在不同页都出现
			if gstr.Pos(tplDao, daoParseOrderTmp) == -1 {
				daoParseOrder += daoParseOrderTmp
			}
			continue
		}
		//idPath|id_path
		if fieldCaseCamel == `IdPath` && gstr.Pos(column[`Type`].String(), `varchar`) != -1 && tpl.PidHandle.IsCoexist {
			continue
		}
		//sort
		//weight
		if garray.NewStrArrayFrom([]string{`sort`, `weight`}).Contains(field) && gstr.Pos(column[`Type`].String(), `int`) != -1 {
			daoParseOrderTmp := `
			case daoThis.Columns().` + fieldCaseCamel + `:
				m = m.Order(daoThis.Table()+` + "`.`" + `+kArr[0], kArr[1])
				m = m.OrderDesc(daoThis.Table() + ` + "`.`" + ` + daoThis.PrimaryKey())` //追加主键倒序。mysql排序字段有重复值时，分页会导致同一条数据可能在不同页都出现
			if gstr.Pos(tplDao, daoParseOrderTmp) == -1 {
				daoParseOrder += daoParseOrderTmp
			}
			continue
		}
		//id后缀
		if gstr.SubStr(fieldCaseCamel, -2) == `Id` && gstr.Pos(column[`Type`].String(), `int`) != -1 {
			if tpl.RelTableMap[field].IsExistRelTableDao {
				relTable := tpl.RelTableMap[field]
				daoPath := relTable.RelTableNameCaseCamel
				if !relTable.IsSameDir {
					daoPath = `dao` + relTable.RelDaoDirCaseCamel + `.` + relTable.RelTableNameCaseCamel
					daoImportOtherDaoTmp := `
	dao` + relTable.RelDaoDirCaseCamel + ` "api/internal/dao/` + relTable.RelDaoDir + `"`
					if gstr.Pos(tplDao, daoImportOtherDaoTmp) == -1 {
						daoImportOtherDao += daoImportOtherDaoTmp
					}
				}
				daoParseFieldTmp := `
			case ` + "`" + relTable.RelNameFieldCaseCamelLower + "`" + `:
				m = m.Fields(` + daoPath + `.Table() + ` + "`.`" + ` + v)
				m = daoThis.ParseJoin(` + daoPath + `.Table(), joinTableArr)(m)`
				if gstr.Pos(tplDao, daoParseFieldTmp) == -1 {
					daoParseField += daoParseFieldTmp
				}
				daoParseJoinTmp := `
		case ` + daoPath + `.Table():
			relTable := ` + daoPath + `.Table()
			if !garray.NewStrArrayFrom(*joinTableArr).Contains(relTable) {
				*joinTableArr = append(*joinTableArr, relTable)
				m = m.LeftJoin(relTable, relTable+` + "`.`" + `+` + daoPath + `.PrimaryKey()+` + "` = `" + `+daoThis.Table()+` + "`.`" + `+daoThis.Columns().` + fieldCaseCamel + `)
			}`
				if gstr.Pos(tplDao, daoParseJoinTmp) == -1 {
					daoParseJoin += daoParseJoinTmp
				}
			}
			continue
		}
		//start_前缀
		if gstr.SubStr(fieldCaseSnake, 0, 6) == `start_` && (gstr.Pos(column[`Type`].String(), `timestamp`) != -1 || gstr.Pos(column[`Type`].String(), `datetime`) != -1 || gstr.Pos(column[`Type`].String(), `date`) != -1) {
			daoParseFilterTmp := `
			case daoThis.Columns().` + fieldCaseCamel + `:
				m = m.WhereLTE(daoThis.Table()+` + "`.`" + `+k, v)`
			if column[`Null`].String() == `NO` && column[`Default`].String() == `` {
				daoParseFilterTmp = `
			case daoThis.Columns().` + fieldCaseCamel + `:
				m = m.Where(m.Builder().WhereLTE(daoThis.Table()+` + "`.`" + `+k, v).WhereOrNull(daoThis.Table() + ` + "`.`" + ` + k))`
			}
			if gstr.Pos(tplDao, daoParseFilterTmp) == -1 {
				daoParseFilter += daoParseFilterTmp
			}
			continue
		}
		//end_前缀
		if gstr.SubStr(fieldCaseSnake, 0, 4) == `end_` && (gstr.Pos(column[`Type`].String(), `timestamp`) != -1 || gstr.Pos(column[`Type`].String(), `datetime`) != -1 || gstr.Pos(column[`Type`].String(), `date`) != -1) {
			daoParseFilterTmp := `
			case daoThis.Columns().` + fieldCaseCamel + `:
				m = m.WhereGTE(daoThis.Table()+` + "`.`" + `+k, v)`
			if column[`Null`].String() == `NO` && column[`Default`].String() == `` {
				daoParseFilterTmp = `
			case daoThis.Columns().` + fieldCaseCamel + `:
				m = m.Where(m.Builder().WhereGTE(daoThis.Table()+` + "`.`" + `+k, v).WhereOrNull(daoThis.Table() + ` + "`.`" + ` + k))`
			}
			if gstr.Pos(tplDao, daoParseFilterTmp) == -1 {
				daoParseFilter += daoParseFilterTmp
			}
			continue
		}
		//timestamp或datetime类型
		if gstr.Pos(column[`Type`].String(), `timestamp`) != -1 || gstr.Pos(column[`Type`].String(), `datetime`) != -1 {
			continue
		}
		//date类型
		if gstr.Pos(column[`Type`].String(), `date`) != -1 {
			daoParseOrderTmp := `
			case daoThis.Columns().` + fieldCaseCamel + `:
				m = m.Order(daoThis.Table()+` + "`.`" + `+kArr[0], kArr[1])
				m = m.OrderDesc(daoThis.Table() + ` + "`.`" + ` + daoThis.PrimaryKey())` //追加主键倒序。mysql排序字段有重复值时，分页会导致同一条数据可能在不同页都出现
			if gstr.Pos(tplDao, daoParseOrderTmp) == -1 {
				daoParseOrder += daoParseOrderTmp
			}
			continue
		}
	}

	if tpl.PasswordHandle.IsCoexist {
		daoParseInsertTmp := `
			case daoThis.Columns().` + gstr.CaseCamel(tpl.PasswordHandle.PasswordField) + `:
				salt := grand.S(8)
				insertData[daoThis.Columns().` + gstr.CaseCamel(tpl.PasswordHandle.SaltField) + `] = salt
				insertData[daoThis.Columns().` + gstr.CaseCamel(tpl.PasswordHandle.PasswordField) + `] = gmd5.MustEncrypt(gconv.String(v) + salt)`
		if gstr.Pos(tplDao, daoParseInsertTmp) == -1 {
			daoParseInsert += daoParseInsertTmp
		}
		daoParseUpdateTmp := `
			case daoThis.Columns().` + gstr.CaseCamel(tpl.PasswordHandle.PasswordField) + `:
				salt := grand.S(8)
				updateData[daoThis.Table()+` + "`.`" + `+daoThis.Columns().` + gstr.CaseCamel(tpl.PasswordHandle.SaltField) + `] = salt
				updateData[daoThis.Table()+` + "`.`" + `+daoThis.Columns().` + gstr.CaseCamel(tpl.PasswordHandle.PasswordField) + `] = gmd5.MustEncrypt(gconv.String(v) + salt)`
		if gstr.Pos(tplDao, daoParseUpdateTmp) == -1 {
			daoParseUpdate += daoParseUpdateTmp
		}
	}

	if tpl.PidHandle.IsCoexist {
		daoParseInsertTmp := `
			case daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.PidField) + `:
				insertData[k] = v
				if gconv.Int(v) > 0 {
					pInfo, _ := daoThis.ParseDbCtx(m.GetCtx()).Where(daoThis.PrimaryKey(), v).Fields(daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.IdPathField) + `, daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.LevelField) + `).One()
					hookData[` + "`pIdPath`" + `] = pInfo[daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.IdPathField) + `].String()
					hookData[` + "`pLevel`" + `] = pInfo[daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.LevelField) + `].Int()
				} else {
					hookData[` + "`pIdPath`" + `] = ` + "`0`" + `
					hookData[` + "`pLevel`" + `] = 0
				}`
		if gstr.Pos(tplDao, daoParseInsertTmp) == -1 {
			daoParseInsert += daoParseInsertTmp
		}
		daoHookInsertTmp := `

			updateSelfData := map[string]interface{}{}
			for k, v := range data {
				switch k {
				case ` + "`pIdPath`" + `:
					updateSelfData[daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.IdPathField) + `] = gconv.String(v) + ` + "`-`" + ` + gconv.String(id)
				case ` + "`pLevel`" + `:
					updateSelfData[daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.LevelField) + `] = gconv.Int(v) + 1
				}
			}
			if len(updateSelfData) > 0 {
				daoThis.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), id).Data(updateSelfData).Update()
			}`
		if gstr.Pos(tplDao, daoHookInsertTmp) == -1 {
			daoHookInsert += daoHookInsertTmp
		}
		daoParseUpdateTmp := `
			case daoThis.Columns().Pid:
				updateData[daoThis.Table()+` + "`.`" + `+k] = v
				pIdPath := ` + "`0`" + `
				pLevel := 0
				if gconv.Int(v) > 0 {
					pInfo, _ := daoThis.ParseDbCtx(m.GetCtx()).Where(daoThis.PrimaryKey(), v).Fields(daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.IdPathField) + `, daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.LevelField) + `).One()
					pIdPath = pInfo[daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.IdPathField) + `].String()
					pLevel = pInfo[daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.LevelField) + `].Int()
				}
				updateData[daoThis.Table()+` + "`.`" + `+daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.IdPathField) + `] = gdb.Raw(` + "`CONCAT('`" + ` + pIdPath + ` + "`-', `" + ` + daoThis.PrimaryKey() + ` + "`)`" + `)
				updateData[daoThis.Table()+` + "`.`" + `+daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.LevelField) + `] = pLevel + 1`
		if gstr.Pos(tplDao, daoParseUpdateTmp) == -1 {
			daoParseUpdate += daoParseUpdateTmp
		}
		daoHookUpdateAfterTmp := `

			for k, v := range data {
				switch k {
				case ` + "`updateChildIdPathAndLevelList`" + `: //修改pid时，更新所有子孙级的idPath和level。参数：[]map[string]interface{}{newIdPath: 父级新idPath, oldIdPath: 父级旧idPath, newLevel: 父级新level, oldLevel: 父级旧level}
					val := v.([]map[string]interface{})
					for _, v1 := range val {
						daoThis.UpdateChildIdPathAndLevel(ctx, gconv.String(v1[` + "`newIdPath`" + `]), gconv.String(v1[` + "`oldIdPath`" + `]), gconv.Int(v1[` + "`newLevel`" + `]), gconv.Int(v1[` + "`oldLevel`" + `]))
					}
				}
			}`
		if gstr.Pos(tplDao, daoHookUpdateAfterTmp) == -1 {
			daoHookUpdateAfter += daoHookUpdateAfterTmp
		}
		daoFuncTmp := `
// 修改pid时，更新所有子孙级的idPath和level
func (daoThis *` + tpl.TableNameCaseCamelLower + `Dao) UpdateChildIdPathAndLevel(ctx context.Context, newIdPath string, oldIdPath string, newLevel int, oldLevel int) {
	daoThis.ParseDbCtx(ctx).WhereLike(daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.IdPathField) + `, oldIdPath+` + "`-%`" + `).Data(g.Map{
		daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.IdPathField) + `: gdb.Raw(` + "`REPLACE(`" + ` + daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.IdPathField) + ` + ` + "`, '`" + ` + oldIdPath + ` + "`', '`" + ` + newIdPath + ` + "`')`" + `),
		daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.LevelField) + `:  gdb.Raw(daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.LevelField) + ` + ` + "` + `" + ` + gconv.String(newLevel-oldLevel)),
	}).Update()
}`
		if gstr.Pos(tplDao, daoFuncTmp) == -1 {
			daoFunc += daoFuncTmp
		}
	}

	if daoParseInsert != `` {
		daoParseInsertPoint := `case ` + "`id`" + `:
				insertData[daoThis.PrimaryKey()] = v`
		tplDao = gstr.Replace(tplDao, daoParseInsertPoint, daoParseInsertPoint+daoParseInsert)
	}
	if daoHookInsert != `` {
		daoHookInsertPoint := `// id, _ := result.LastInsertId()`
		tplDao = gstr.Replace(tplDao, daoHookInsertPoint, `id, _ := result.LastInsertId()`+daoHookInsert)
	}
	if daoParseUpdate != `` {
		daoParseUpdatePoint := `case ` + "`id`" + `:
				updateData[daoThis.Table()+` + "`.`" + `+daoThis.PrimaryKey()] = v`
		tplDao = gstr.Replace(tplDao, daoParseUpdatePoint, daoParseUpdatePoint+daoParseUpdate)
	}
	if daoHookUpdateBefore != `` || daoHookUpdateAfter != `` {
		daoHookUpdatePoint := `

			/* row, _ := result.RowsAffected()
			if row == 0 {
				return
			} */`
		if daoHookUpdateBefore != `` {
			tplDao = gstr.Replace(tplDao, daoHookUpdatePoint, daoHookUpdateBefore+`

			/* row, _ := result.RowsAffected()
			if row == 0 {
				return
			} */`)
		}
		if daoHookUpdateAfter != `` {
			tplDao = gstr.Replace(tplDao, daoHookUpdatePoint, `

			row, _ := result.RowsAffected()
			if row == 0 {
				return
			}`+daoHookUpdateAfter)
		}
	}
	if daoParseField != `` {
		daoParseFieldPoint := `case ` + "`id`" + `:
				m = m.Fields(daoThis.Table() + ` + "`.`" + ` + daoThis.PrimaryKey() + ` + "` AS `" + ` + v)`
		tplDao = gstr.Replace(tplDao, daoParseFieldPoint, daoParseFieldPoint+daoParseField)
	}
	if daoHookSelect != `` {
		daoHookSelectPoint := `
						/* case ` + "`xxxx`" + `:
						record[v] = gvar.New(` + "``" + `) */`
		tplDao = gstr.Replace(tplDao, daoHookSelectPoint, daoHookSelect)
	}
	if daoParseFilter != `` {
		daoParseFilterPoint := `case ` + "`id`, `idArr`" + `:
				m = m.Where(daoThis.Table()+` + "`.`" + `+daoThis.PrimaryKey(), v)`
		tplDao = gstr.Replace(tplDao, daoParseFilterPoint, daoParseFilterPoint+daoParseFilter)
	}
	if daoParseOrder != `` {
		daoParseOrderPoint := `case ` + "`id`" + `:
				m = m.Order(daoThis.Table()+` + "`.`" + `+daoThis.PrimaryKey(), kArr[1])`
		tplDao = gstr.Replace(tplDao, daoParseOrderPoint, daoParseOrderPoint+daoParseOrder)
	}
	if daoParseJoin != `` {
		daoParseJoinPoint := `
		/* case Xxxx.Table():
		relTable := Xxxx.Table()
		if !garray.NewStrArrayFrom(*joinTableArr).Contains(relTable) {
			*joinTableArr = append(*joinTableArr, relTable)
			m = m.LeftJoin(relTable, relTable+` + "`.`" + `+daoThis.PrimaryKey()+` + "` = `" + `+daoThis.Table()+` + "`.`" + `+daoThis.PrimaryKey())
		} */`
		tplDao = gstr.Replace(tplDao, daoParseJoinPoint, daoParseJoin)
	}
	if daoFunc != `` {
		tplDao = tplDao + daoFunc
	}
	if daoImportOtherDao != `` {
		daoImportOtherDaoPoint := `"api/internal/dao/auth/internal"`
		tplDao = gstr.Replace(tplDao, daoImportOtherDaoPoint, daoImportOtherDaoPoint+daoImportOtherDao)
	}

	tplDao = gstr.Replace(tplDao, `"github.com/gogf/gf/v2/util/gconv"`, `"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/gogf/gf/v2/container/garray"
	`)

	gfile.PutContents(saveFile, tplDao)
	utils.GoFileFmt(saveFile)
}

// logic模板生成（文件不存在时增删改查全部生成，已存在不处理不覆盖）
func MyGenTplLogic(ctx context.Context, option *MyGenOption, tpl *MyGenTpl) {
	saveFile := gfile.SelfDir() + `/internal/logic/` + gstr.LcFirst(tpl.ModuleDirCaseCamel) + `/` + tpl.TableNameCaseSnake + `.go`
	if gfile.IsFile(saveFile) {
		return
	}

	tplLogic := `package logic

import (
	"api/internal/dao"
	dao` + tpl.ModuleDirCaseCamel + ` "api/internal/dao/` + tpl.ModuleDirCaseCamelLower + `"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type s` + tpl.LogicStructName + ` struct{}

func New` + tpl.LogicStructName + `() *s` + tpl.LogicStructName + ` {
	return &s` + tpl.LogicStructName + `{}
}

func init() {
	service.Register` + tpl.LogicStructName + `(New` + tpl.LogicStructName + `())
}

// 新增
func (logicThis *s` + tpl.LogicStructName + `) Create(ctx context.Context, data map[string]interface{}) (id int64, err error) {
	daoThis := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `
	id, err = dao.NewDaoHandler(ctx, &daoThis).Insert(data).GetModel().InsertAndGetId()
	return
}

// 修改
func (logicThis *s` + tpl.LogicStructName + `) Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error) {
	daoThis := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `
	daoHandlerThis := dao.NewDaoHandler(ctx, &daoThis).Filter(filter)
	idArr, _ := daoHandlerThis.GetModel(true).Array(daoThis.PrimaryKey())
	if len(idArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ` + "``" + `)
		return
	}
	hookData := map[string]interface{}{}
`
	if tpl.PidHandle.IsCoexist {
		tplLogic += `
	_, okPid := data[daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.PidField) + `]
	if okPid {
		pInfo := gdb.Record{}
		pid := gconv.Int(data[daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.PidField) + `])
		if pid > 0 {
			pInfo, _ = daoThis.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), pid).One()
			if len(pInfo) == 0 {
				err = utils.NewErrorCode(ctx, 29999997, ` + "``" + `)
				return
			}
		}
		updateChildIdPathAndLevelList := []map[string]interface{}{}
		for _, id := range idArr {
			oldInfo, _ := daoThis.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), id).One()
			if pid == oldInfo[daoThis.PrimaryKey()].Int() { //父级不能是自身
				err = utils.NewErrorCode(ctx, 29999996, ` + "``" + `)
				return
			}
			if pid != oldInfo[daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.PidField) + `].Int() {
				pIdPath := ` + "`0`" + `
				pLevel := 0
				if pid > 0 {
					if garray.NewStrArrayFrom(gstr.Split(pInfo[daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.IdPathField) + `].String(), ` + "`-`" + `)).Contains(oldInfo[daoThis.PrimaryKey()].String()) { //父级不能是自身的子孙级
						err = utils.NewErrorCode(ctx, 29999995, ` + "``" + `)
						return
					}
					pIdPath = pInfo[daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.IdPathField) + `].String()
					pLevel = pInfo[daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.LevelField) + `].Int()
				}
				updateChildIdPathAndLevelList = append(updateChildIdPathAndLevelList, map[string]interface{}{
					` + "`" + `newIdPath` + "`" + `: pIdPath + ` + "`-`" + ` + id.String(),
					` + "`" + `oldIdPath` + "`" + `: oldInfo[daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.IdPathField) + `],
					` + "`" + `newLevel` + "`" + `:  pLevel + 1,
					` + "`" + `oldLevel` + "`" + `:  oldInfo[daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.LevelField) + `],
				})
			}
		}

		if len(updateChildIdPathAndLevelList) > 0 {
			hookData[` + "`" + `updateChildIdPathAndLevelList` + "`" + `] = updateChildIdPathAndLevelList
		}
	}
`
	}
	tplLogic += `
	row, err = daoHandlerThis.Update(data).HookUpdate(hookData, gconv.SliceInt(idArr)...).GetModel().UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *s` + tpl.LogicStructName + `) Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error) {
	daoThis := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `
	daoHandlerThis := dao.NewDaoHandler(ctx, &daoThis).Filter(filter)
	idArr, _ := daoHandlerThis.GetModel(true).Array(daoThis.PrimaryKey())
	if len(idArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ` + "``" + `)
		return
	}
`
	if tpl.PidHandle.PidField != `` {
		tplLogic += `
	count, _ := daoThis.ParseDbCtx(ctx).Where(daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.PidField) + `, idArr).Count()
	if count > 0 {
		err = utils.NewErrorCode(ctx, 29999994, ` + "``" + `)
		return
	}
`
	}
	tplLogic += `
	result, err := daoHandlerThis.HookDelete(gconv.SliceInt(idArr)...).GetModel().Delete()
	row, _ = result.RowsAffected()
	return
}
`

	gfile.PutContents(saveFile, tplLogic)
	utils.GoFileFmt(saveFile)
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
	apiResColumnAlloweFieldList := ``
	if tpl.LabelField != `` {
		apiReqFilterColumn += `Label          string      ` + "`" + `json:"label,omitempty" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"标签。常用于前端组件"` + "`\n"
		apiResColumn += `Label       *string     ` + "`" + `json:"label,omitempty" dc:"标签。常用于前端组件"` + "`\n"
	}
	for _, column := range tpl.TableColumnList {
		field := column[`Field`].String()
		fieldCaseCamel := gstr.CaseCamel(field)
		fieldCaseSnake := gstr.CaseSnakeFirstUpper(field)
		comment := gstr.Trim(gstr.ReplaceByArray(column[`Comment`].String(), g.SliceStr{
			"\n", ` `,
			"\r", ` `,
			`"`, `\"`,
		}))
		resultStr, _ := gregex.MatchString(`.*\((\d*)\)`, column[`Type`].String())

		if garray.NewStrArrayFrom([]string{`DeletedAt`, `DeleteAt`, `DeletedTime`, `DeleteTime`}).Contains(fieldCaseCamel) {
			continue
		}
		if garray.NewStrArrayFrom([]string{`UpdatedAt`, `UpdateAt`, `UpdatedTime`, `UpdateTime`}).Contains(fieldCaseCamel) {
			apiResColumn += fieldCaseCamel + ` *gtime.Time ` + "`" + `json:"` + field + `,omitempty" dc:"` + comment + `"` + "`\n"
			continue
		}
		if garray.NewStrArrayFrom([]string{`CreatedAt`, `CreateAt`, `CreatedTime`, `CreateTime`}).Contains(fieldCaseCamel) {
			apiReqFilterColumn += `TimeRangeStart *gtime.Time ` + "`" + `json:"timeRangeStart,omitempty" v:"date-format:Y-m-d H:i:s" dc:"开始时间：YYYY-mm-dd HH:ii:ss"` + "`\n"
			apiReqFilterColumn += `TimeRangeEnd   *gtime.Time ` + "`" + `json:"timeRangeEnd,omitempty" v:"date-format:Y-m-d H:i:s|after-equal:TimeRangeStart" dc:"结束时间：YYYY-mm-dd HH:ii:ss"` + "`\n"
			apiResColumn += fieldCaseCamel + ` *gtime.Time ` + "`" + `json:"` + field + `,omitempty" dc:"` + comment + `"` + "`\n"
			continue
		}
		//主键
		if column[`Key`].String() == `PRI` && column[`Extra`].String() == `auto_increment` {
			if field != `id` {
				apiReqFilterColumn += fieldCaseCamel + ` *uint ` + "`" + `json:"` + field + `,omitempty" v:"integer|min:1" dc:"` + comment + `"` + "`\n"
				apiResColumn += fieldCaseCamel + ` *uint ` + "`" + `json:"` + field + `,omitempty" dc:"` + comment + `"` + "`\n"
			}
			continue
		}
		//password|passwd
		if garray.NewStrArrayFrom([]string{`password`, `passwd`}).Contains(field) && column[`Type`].String() == `char(32)` {
			apiReqCreateColumn += fieldCaseCamel + ` *string ` + "`" + `json:"` + field + `,omitempty" v:"size:` + resultStr[1] + `" dc:"` + comment + `"` + "`\n"
			apiReqUpdateColumn += fieldCaseCamel + ` *string ` + "`" + `json:"` + field + `,omitempty" v:"size:` + resultStr[1] + `" dc:"` + comment + `"` + "`\n"
			continue
		}
		//salt
		if garray.NewStrArrayFrom([]string{`salt`}).Contains(field) && column[`Type`].String() == `char(8)` && tpl.PasswordHandle.IsCoexist {
			continue
		}
		//pid
		if garray.NewStrArrayFrom([]string{`pid`}).Contains(field) && gstr.Pos(column[`Type`].String(), `int`) != -1 {
			apiReqFilterColumn += fieldCaseCamel + ` *uint ` + "`" + `json:"` + field + `,omitempty" v:"integer|min:0" dc:"` + comment + `"` + "`\n"
			apiReqCreateColumn += fieldCaseCamel + ` *uint ` + "`" + `json:"` + field + `,omitempty" v:"integer|min:0" dc:"` + comment + `"` + "`\n"
			apiReqUpdateColumn += fieldCaseCamel + ` *uint ` + "`" + `json:"` + field + `,omitempty" v:"integer|min:0" dc:"` + comment + `"` + "`\n"
			apiResColumn += fieldCaseCamel + ` *uint ` + "`" + `json:"` + field + `,omitempty" dc:"` + comment + `"` + "`\n"
			if tpl.LabelField != `` {
				apiResColumnAlloweFieldList += `P` + gstr.CaseCamel(tpl.LabelField) + ` *string ` + "`" + `json:"p` + gstr.CaseCamel(tpl.LabelField) + `,omitempty" dc:"父级"` + "`\n"
			}
			continue
		}
		//level
		if garray.NewStrArrayFrom([]string{`level`}).Contains(field) && gstr.Pos(column[`Type`].String(), `int`) != -1 && tpl.PidHandle.IsCoexist {
			apiReqFilterColumn += fieldCaseCamel + ` *uint ` + "`" + `json:"` + field + `,omitempty" v:"integer|min:0" dc:"` + comment + `"` + "`\n"
			apiResColumn += fieldCaseCamel + ` *uint ` + "`" + `json:"` + field + `,omitempty" dc:"` + comment + `"` + "`\n"
			continue
		}
		//idPath|id_path
		if fieldCaseCamel == `IdPath` && gstr.Pos(column[`Type`].String(), `varchar`) != -1 && tpl.PidHandle.IsCoexist {
			apiResColumn += fieldCaseCamel + ` *string ` + "`" + `json:"` + field + `,omitempty" dc:"` + comment + `"` + "`\n"
			continue
		}
		//sort
		//weight
		if garray.NewStrArrayFrom([]string{`sort`, `weight`}).Contains(field) && gstr.Pos(column[`Type`].String(), `int`) != -1 {
			apiReqCreateColumn += fieldCaseCamel + ` *uint ` + "`" + `json:"` + field + `,omitempty" v:"integer|between:0,100" dc:"` + comment + `"` + "`\n"
			apiReqUpdateColumn += fieldCaseCamel + ` *uint ` + "`" + `json:"` + field + `,omitempty" v:"integer|between:0,100" dc:"` + comment + `"` + "`\n"
			apiResColumn += fieldCaseCamel + ` *uint ` + "`" + `json:"` + field + `,omitempty" dc:"` + comment + `"` + "`\n"
			continue
		}
		//id后缀
		if gstr.SubStr(fieldCaseCamel, -2) == `Id` && gstr.Pos(column[`Type`].String(), `int`) != -1 {
			apiReqFilterColumn += fieldCaseCamel + ` *uint ` + "`" + `json:"` + field + `,omitempty" v:"integer|min:1" dc:"` + comment + `"` + "`\n"
			apiReqCreateColumn += fieldCaseCamel + ` *uint ` + "`" + `json:"` + field + `,omitempty" v:"integer|min:1" dc:"` + comment + `"` + "`\n"
			apiReqUpdateColumn += fieldCaseCamel + ` *uint ` + "`" + `json:"` + field + `,omitempty" v:"integer|min:1" dc:"` + comment + `"` + "`\n"
			apiResColumn += fieldCaseCamel + ` *uint ` + "`" + `json:"` + field + `,omitempty" dc:"` + comment + `"` + "`\n"
			if tpl.RelTableMap[field].IsExistRelTableDao {
				relTable := tpl.RelTableMap[field]
				apiResColumnAlloweFieldList += gstr.CaseCamel(relTable.RelNameFieldCaseCamelLower) + ` *string ` + "`" + `json:"` + relTable.RelNameFieldCaseCamelLower + `,omitempty" dc:"` + relTable.RelNameFieldName + `"` + "`\n"
			}
			continue
		}
		//name或code后缀
		if (gstr.SubStr(fieldCaseCamel, -4) == `Name` || gstr.SubStr(fieldCaseCamel, -4) == `Code`) && gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
			apiReqFilterColumn += fieldCaseCamel + ` string ` + "`" + `json:"` + field + `,omitempty" v:"length:1,` + resultStr[1] + `|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"` + comment + `"` + "`\n"
			apiReqCreateColumn += fieldCaseCamel + ` *string ` + "`" + `json:"` + field + `,omitempty" v:"length:1,` + resultStr[1] + `|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"` + comment + `"` + "`\n"
			apiReqUpdateColumn += fieldCaseCamel + ` *string ` + "`" + `json:"` + field + `,omitempty" v:"length:1,` + resultStr[1] + `|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"` + comment + `"` + "`\n"
			apiResColumn += fieldCaseCamel + ` *string ` + "`" + `json:"` + field + `,omitempty" dc:"` + comment + `"` + "`\n"
			continue
		}
		//mobile,phone后缀
		if (gstr.SubStr(fieldCaseCamel, -5) == `Phone` || gstr.SubStr(fieldCaseCamel, -6) == `Mobile`) && gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
			apiReqFilterColumn += fieldCaseCamel + ` string ` + "`" + `json:"` + field + `,omitempty" v:"phone|length:1,` + resultStr[1] + `" dc:"` + comment + `"` + "`\n"
			apiReqCreateColumn += fieldCaseCamel + ` *string ` + "`" + `json:"` + field + `,omitempty" v:"phone|length:1,` + resultStr[1] + `" dc:"` + comment + `"` + "`\n"
			apiReqUpdateColumn += fieldCaseCamel + ` *string ` + "`" + `json:"` + field + `,omitempty" v:"phone|length:1,` + resultStr[1] + `" dc:"` + comment + `"` + "`\n"
			apiResColumn += fieldCaseCamel + ` *string ` + "`" + `json:"` + field + `,omitempty" dc:"` + comment + `"` + "`\n"
			continue
		}
		//url,link后缀
		if (gstr.SubStr(fieldCaseCamel, -3) == `Url` || gstr.SubStr(fieldCaseCamel, -4) == `Link`) && gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
			apiReqFilterColumn += fieldCaseCamel + ` string ` + "`" + `json:"` + field + `,omitempty" v:"url|length:1,` + resultStr[1] + `" dc:"` + comment + `"` + "`\n"
			apiReqCreateColumn += fieldCaseCamel + ` *string ` + "`" + `json:"` + field + `,omitempty" v:"url|length:1,` + resultStr[1] + `" dc:"` + comment + `"` + "`\n"
			apiReqUpdateColumn += fieldCaseCamel + ` *string ` + "`" + `json:"` + field + `,omitempty" v:"url|length:1,` + resultStr[1] + `" dc:"` + comment + `"` + "`\n"
			apiResColumn += fieldCaseCamel + ` *string ` + "`" + `json:"` + field + `,omitempty" dc:"` + comment + `"` + "`\n"
			continue
		}
		//avatar
		//icon,cover,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀
		//video,video_list,videoList,video_arr,videoArr等后缀
		if field == `avatar` || gstr.SubStr(fieldCaseCamel, -4) == `Icon` || gstr.SubStr(fieldCaseCamel, -5) == `Cover` || gstr.SubStr(fieldCaseCamel, -3) == `Img` || gstr.SubStr(fieldCaseCamel, -7) == `ImgList` || gstr.SubStr(fieldCaseCamel, -6) == `ImgArr` || gstr.SubStr(fieldCaseCamel, -5) == `Image` || gstr.SubStr(fieldCaseCamel, -9) == `ImageList` || gstr.SubStr(fieldCaseCamel, -8) == `ImageArr` || gstr.SubStr(fieldCaseCamel, -5) == `Video` || gstr.SubStr(fieldCaseCamel, -9) == `VideoList` || gstr.SubStr(fieldCaseCamel, -8) == `VideoArr` {
			if column[`Type`].String() == `json` || column[`Type`].String() == `text` {
				apiReqCreateColumn += fieldCaseCamel + ` *[]string ` + "`" + `json:"` + field + `,omitempty" v:"distinct|foreach|url|foreach|min-length:1" dc:"` + comment + `"` + "`\n"
				apiReqUpdateColumn += fieldCaseCamel + ` *[]string ` + "`" + `json:"` + field + `,omitempty" v:"distinct|foreach|url|foreach|min-length:1" dc:"` + comment + `"` + "`\n"
				apiResColumn += fieldCaseCamel + ` []string ` + "`" + `json:"` + field + `,omitempty" dc:"` + comment + `"` + "`\n"
			} else {
				apiReqCreateColumn += fieldCaseCamel + ` *string ` + "`" + `json:"` + field + `,omitempty" v:"url|length:1,` + resultStr[1] + `" dc:"` + comment + `"` + "`\n"
				apiReqUpdateColumn += fieldCaseCamel + ` *string ` + "`" + `json:"` + field + `,omitempty" v:"url|length:1,` + resultStr[1] + `" dc:"` + comment + `"` + "`\n"
				apiResColumn += fieldCaseCamel + ` *string ` + "`" + `json:"` + field + `,omitempty" dc:"` + comment + `"` + "`\n"
			}
			continue
		}
		//list,arr等后缀
		if (gstr.SubStr(fieldCaseCamel, -4) == `List` || gstr.SubStr(fieldCaseCamel, -3) == `Arr`) && (column[`Type`].String() == `json` || column[`Type`].String() == `text`) {
			apiReqCreateColumn += fieldCaseCamel + ` *[]interface{} ` + "`" + `json:"` + field + `,omitempty" v:"distinct" dc:"` + comment + `"` + "`\n"
			apiReqUpdateColumn += fieldCaseCamel + ` *[]interface{} ` + "`" + `json:"` + field + `,omitempty" v:"distinct" dc:"` + comment + `"` + "`\n"
			apiResColumn += fieldCaseCamel + ` []interface{} ` + "`" + `json:"` + field + `,omitempty" dc:"` + comment + `"` + "`\n"
			continue
		}
		//IP后缀
		if gstr.SubStr(fieldCaseCamel, -2) == `Ip` && gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
			apiReqFilterColumn += fieldCaseCamel + ` string ` + "`" + `json:"` + field + `,omitempty" v:"ip|length:1,` + resultStr[1] + `" dc:"` + comment + `"` + "`\n"
			apiReqCreateColumn += fieldCaseCamel + ` *string ` + "`" + `json:"` + field + `,omitempty" v:"ip|length:1,` + resultStr[1] + `" dc:"` + comment + `"` + "`\n"
			apiReqUpdateColumn += fieldCaseCamel + ` *string ` + "`" + `json:"` + field + `,omitempty" v:"ip|length:1,` + resultStr[1] + `" dc:"` + comment + `"` + "`\n"
			apiResColumn += fieldCaseCamel + ` *string ` + "`" + `json:"` + field + `,omitempty" dc:"` + comment + `"` + "`\n"
			continue
		}
		//remark,desc,msg,message,intro,content后缀
		if (gstr.SubStr(fieldCaseCamel, -6) == `Remark` || gstr.SubStr(fieldCaseCamel, -4) == `Desc` || gstr.SubStr(fieldCaseCamel, -3) == `Msg` || gstr.SubStr(fieldCaseCamel, -7) == `Message` || gstr.SubStr(fieldCaseCamel, -5) == `Intro` || gstr.SubStr(fieldCaseCamel, -7) == `Content`) && (gstr.Pos(column[`Type`].String(), `varchar`) != -1 || column[`Type`].String() == `text`) {
			apiReqCreateColumn += fieldCaseCamel + ` *string ` + "`" + `json:"` + field + `,omitempty" v:"" dc:"` + comment + `"` + "`\n"
			apiReqUpdateColumn += fieldCaseCamel + ` *string ` + "`" + `json:"` + field + `,omitempty" v:"" dc:"` + comment + `"` + "`\n"
			apiResColumn += fieldCaseCamel + ` *string ` + "`" + `json:"` + field + `,omitempty" dc:"` + comment + `"` + "`\n"
			continue
		}
		//gender
		//status或type后缀
		if (field == `gender` || gstr.SubStr(fieldCaseCamel, -6) == `Status` || gstr.SubStr(fieldCaseCamel, -4) == `Type`) && gstr.Pos(column[`Type`].String(), `int`) != -1 {
			statusList := MyGenStatusList(comment)
			statusArr := make([]string, len(statusList))
			for index, status := range statusList {
				statusArr[index] = status[0]
			}
			statusStr := gstr.Join(statusArr, `,`)
			apiReqFilterColumn += fieldCaseCamel + ` *uint ` + "`" + `json:"` + field + `,omitempty" v:"integer|in:` + statusStr + `" dc:"` + comment + `"` + "`\n"
			apiReqCreateColumn += fieldCaseCamel + ` *uint ` + "`" + `json:"` + field + `,omitempty" v:"integer|in:` + statusStr + `" dc:"` + comment + `"` + "`\n"
			apiReqUpdateColumn += fieldCaseCamel + ` *uint ` + "`" + `json:"` + field + `,omitempty" v:"integer|in:` + statusStr + `" dc:"` + comment + `"` + "`\n"
			apiResColumn += fieldCaseCamel + ` *uint ` + "`" + `json:"` + field + `,omitempty" dc:"` + comment + `"` + "`\n"
			continue
		}
		//is_前缀
		if gstr.SubStr(fieldCaseSnake, 0, 3) == `is_` && gstr.Pos(column[`Type`].String(), `int`) != -1 {
			apiReqFilterColumn += fieldCaseCamel + ` *uint ` + "`" + `json:"` + field + `,omitempty" v:"integer|in:0,1" dc:"` + comment + `"` + "`\n"
			apiReqCreateColumn += fieldCaseCamel + ` *uint ` + "`" + `json:"` + field + `,omitempty" v:"integer|in:0,1" dc:"` + comment + `"` + "`\n"
			apiReqUpdateColumn += fieldCaseCamel + ` *uint ` + "`" + `json:"` + field + `,omitempty" v:"integer|in:0,1" dc:"` + comment + `"` + "`\n"
			apiResColumn += fieldCaseCamel + ` *uint ` + "`" + `json:"` + field + `,omitempty" dc:"` + comment + `"` + "`\n"
			continue
		}
		//start_前缀
		//end_前缀
		if gstr.SubStr(fieldCaseSnake, 0, 6) == `start_` || gstr.SubStr(fieldCaseSnake, 0, 4) == `end_` {
			if gstr.Pos(column[`Type`].String(), `timestamp`) != -1 || gstr.Pos(column[`Type`].String(), `datetime`) != -1 {
				apiReqFilterColumn += fieldCaseCamel + ` *gtime.Time ` + "`" + `json:"` + field + `,omitempty" v:"date-format:Y-m-d H:i:s" dc:"` + comment + `"` + "`\n"
				if column[`Null`].String() == `NO` && column[`Default`].String() == `` {
					apiReqCreateColumn += fieldCaseCamel + ` *gtime.Time ` + "`" + `json:"` + field + `,omitempty" v:"required|date-format:Y-m-d H:i:s" dc:"` + comment + `"` + "`\n"
				} else {
					apiReqCreateColumn += fieldCaseCamel + ` *gtime.Time ` + "`" + `json:"` + field + `,omitempty" v:"date-format:Y-m-d H:i:s" dc:"` + comment + `"` + "`\n"
				}
				apiReqUpdateColumn += fieldCaseCamel + ` *gtime.Time ` + "`" + `json:"` + field + `,omitempty" v:"date-format:Y-m-d H:i:s" dc:"` + comment + `"` + "`\n"
				apiResColumn += fieldCaseCamel + ` *gtime.Time ` + "`" + `json:"` + field + `,omitempty" dc:"` + comment + `"` + "`\n"
				continue
			}
			/* if gstr.Pos(column[`Type`].String(), `date`) != -1 {
				continue
			} */
		}
		//int类型
		if gstr.Pos(column[`Type`].String(), `int`) != -1 {
			if gstr.Pos(column[`Type`].String(), `unsigned`) != -1 {
				apiReqFilterColumn += fieldCaseCamel + ` *uint ` + "`" + `json:"` + field + `,omitempty" v:"integer|min:0" dc:"` + comment + `"` + "`\n"
				apiReqCreateColumn += fieldCaseCamel + ` *uint ` + "`" + `json:"` + field + `,omitempty" v:"integer|min:0" dc:"` + comment + `"` + "`\n"
				apiReqUpdateColumn += fieldCaseCamel + ` *uint ` + "`" + `json:"` + field + `,omitempty" v:"integer|min:0" dc:"` + comment + `"` + "`\n"
			} else {
				apiReqFilterColumn += fieldCaseCamel + ` *int ` + "`" + `json:"` + field + `,omitempty" v:"integer" dc:"` + comment + `"` + "`\n"
				apiReqCreateColumn += fieldCaseCamel + ` *int ` + "`" + `json:"` + field + `,omitempty" v:"integer" dc:"` + comment + `"` + "`\n"
				apiReqUpdateColumn += fieldCaseCamel + ` *int ` + "`" + `json:"` + field + `,omitempty" v:"integer" dc:"` + comment + `"` + "`\n"
			}
			apiResColumn += fieldCaseCamel + ` *uint ` + "`" + `json:"` + field + `,omitempty" dc:"` + comment + `"` + "`\n"
			continue
		}
		//float类型
		if gstr.Pos(column[`Type`].String(), `decimal`) != -1 || gstr.Pos(column[`Type`].String(), `double`) != -1 || gstr.Pos(column[`Type`].String(), `float`) != -1 {
			if gstr.Pos(column[`Type`].String(), `unsigned`) != -1 {
				apiReqFilterColumn += fieldCaseCamel + ` *float64 ` + "`" + `json:"` + field + `,omitempty" v:"float|min:0" dc:"` + comment + `"` + "`\n"
				apiReqCreateColumn += fieldCaseCamel + ` *float64 ` + "`" + `json:"` + field + `,omitempty" v:"float|min:0" dc:"` + comment + `"` + "`\n"
				apiReqUpdateColumn += fieldCaseCamel + ` *float64 ` + "`" + `json:"` + field + `,omitempty" v:"float|min:0" dc:"` + comment + `"` + "`\n"
			} else {
				apiReqFilterColumn += fieldCaseCamel + ` *float64 ` + "`" + `json:"` + field + `,omitempty" v:"float" dc:"` + comment + `"` + "`\n"
				apiReqCreateColumn += fieldCaseCamel + ` *float64 ` + "`" + `json:"` + field + `,omitempty" v:"float" dc:"` + comment + `"` + "`\n"
				apiReqUpdateColumn += fieldCaseCamel + ` *float64 ` + "`" + `json:"` + field + `,omitempty" v:"float" dc:"` + comment + `"` + "`\n"
			}
			apiResColumn += fieldCaseCamel + ` *float64 ` + "`" + `json:"` + field + `,omitempty" dc:"` + comment + `"` + "`\n"
			continue
		}
		//varchar类型
		if gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
			apiReqFilterColumn += fieldCaseCamel + ` string ` + "`" + `json:"` + field + `,omitempty" v:"length:1,` + resultStr[1] + `" dc:"` + comment + `"` + "`\n"
			apiReqCreateColumn += fieldCaseCamel + ` *string ` + "`" + `json:"` + field + `,omitempty" v:"length:1,` + resultStr[1] + `" dc:"` + comment + `"` + "`\n"
			apiReqUpdateColumn += fieldCaseCamel + ` *string ` + "`" + `json:"` + field + `,omitempty" v:"length:1,` + resultStr[1] + `" dc:"` + comment + `"` + "`\n"
			apiResColumn += fieldCaseCamel + ` *string ` + "`" + `json:"` + field + `,omitempty" dc:"` + comment + `"` + "`\n"
			continue
		}
		//char类型
		if gstr.Pos(column[`Type`].String(), `char`) != -1 {
			apiReqFilterColumn += fieldCaseCamel + ` string ` + "`" + `json:"` + field + `,omitempty" v:"length:1,` + resultStr[1] + `" dc:"` + comment + `"` + "`\n"
			apiReqCreateColumn += fieldCaseCamel + ` *string ` + "`" + `json:"` + field + `,omitempty" v:"size:` + resultStr[1] + `" dc:"` + comment + `"` + "`\n"
			apiReqUpdateColumn += fieldCaseCamel + ` *string ` + "`" + `json:"` + field + `,omitempty" v:"size:` + resultStr[1] + `" dc:"` + comment + `"` + "`\n"
			apiResColumn += fieldCaseCamel + ` *string ` + "`" + `json:"` + field + `,omitempty" dc:"` + comment + `"` + "`\n"
			continue
		}
		//timestamp或datetime类型
		if gstr.Pos(column[`Type`].String(), `timestamp`) != -1 || gstr.Pos(column[`Type`].String(), `datetime`) != -1 {
			if column[`Null`].String() == `NO` && column[`Default`].String() == `` {
				apiReqCreateColumn += fieldCaseCamel + ` *gtime.Time ` + "`" + `json:"` + field + `,omitempty" v:"required|date-format:Y-m-d H:i:s" dc:"` + comment + `"` + "`\n"
			} else {
				apiReqCreateColumn += fieldCaseCamel + ` *gtime.Time ` + "`" + `json:"` + field + `,omitempty" v:"date-format:Y-m-d H:i:s" dc:"` + comment + `"` + "`\n"
			}
			apiReqUpdateColumn += fieldCaseCamel + ` *gtime.Time ` + "`" + `json:"` + field + `,omitempty" v:"date-format:Y-m-d H:i:s" dc:"` + comment + `"` + "`\n"
			apiResColumn += fieldCaseCamel + ` *gtime.Time ` + "`" + `json:"` + field + `,omitempty" dc:"` + comment + `"` + "`\n"
			continue
		}
		//date类型
		if gstr.Pos(column[`Type`].String(), `date`) != -1 {
			apiReqFilterColumn += fieldCaseCamel + ` *gtime.Time ` + "`" + `json:"` + field + `,omitempty" v:"date-format:Y-m-d" dc:"` + comment + `"` + "`\n"
			if column[`Null`].String() == `NO` && column[`Default`].String() == `` {
				apiReqCreateColumn += fieldCaseCamel + ` *gtime.Time ` + "`" + `json:"` + field + `,omitempty" v:"required|date-format:Y-m-d" dc:"` + comment + `"` + "`\n"
			} else {
				apiReqCreateColumn += fieldCaseCamel + ` *gtime.Time ` + "`" + `json:"` + field + `,omitempty" v:"date-format:Y-m-d" dc:"` + comment + `"` + "`\n"
			}
			apiReqUpdateColumn += fieldCaseCamel + ` *gtime.Time ` + "`" + `json:"` + field + `,omitempty" v:"date-format:Y-m-d" dc:"` + comment + `"` + "`\n"
			apiResColumn += fieldCaseCamel + ` *string ` + "`" + `json:"` + field + `,omitempty" dc:"` + comment + `"` + "`\n"
			continue
		}
		//json类型
		if gstr.Pos(column[`Type`].String(), `json`) != -1 {
			if column[`Null`].String() == `NO` {
				apiReqCreateColumn += fieldCaseCamel + ` *string ` + "`" + `json:"` + field + `,omitempty" v:"required|json" dc:"` + comment + `"` + "`\n"
			} else {
				apiReqCreateColumn += fieldCaseCamel + ` *string ` + "`" + `json:"` + field + `,omitempty" v:"json" dc:"` + comment + `"` + "`\n"
			}
			apiReqUpdateColumn += fieldCaseCamel + ` *string ` + "`" + `json:"` + field + `,omitempty" v:"json" dc:"` + comment + `"` + "`\n"
			apiResColumn += fieldCaseCamel + ` *string ` + "`" + `json:"` + field + `,omitempty" dc:"` + comment + `"` + "`\n"
			continue
		}
		//text类型
		if gstr.Pos(column[`Type`].String(), `text`) != -1 {
			if column[`Null`].String() == `NO` {
				apiReqCreateColumn += fieldCaseCamel + ` *string ` + "`" + `json:"` + field + `,omitempty" v:"required" dc:"` + comment + `"` + "`\n"
			} else {
				apiReqCreateColumn += fieldCaseCamel + ` *string ` + "`" + `json:"` + field + `,omitempty" v:"" dc:"` + comment + `"` + "`\n"
			}
			apiReqUpdateColumn += fieldCaseCamel + ` *string ` + "`" + `json:"` + field + `,omitempty" v:"" dc:"` + comment + `"` + "`\n"
			apiResColumn += fieldCaseCamel + ` *string ` + "`" + `json:"` + field + `,omitempty" dc:"` + comment + `"` + "`\n"
			continue
		}
		//默认处理
		apiReqFilterColumn += fieldCaseCamel + ` string ` + "`" + `json:"` + field + `,omitempty" v:"" dc:"` + comment + `"` + "`\n"
		apiReqCreateColumn += fieldCaseCamel + ` *string ` + "`" + `json:"` + field + `,omitempty" v:"" dc:"` + comment + `"` + "`\n"
		apiReqUpdateColumn += fieldCaseCamel + ` *string ` + "`" + `json:"` + field + `,omitempty" v:"" dc:"` + comment + `"` + "`\n"
		apiResColumn += fieldCaseCamel + ` *string ` + "`" + `json:"` + field + `,omitempty" dc:"` + comment + `"` + "`\n"
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
	Field  []string        ` + "`" + `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"` + "`" + `
	Sort   string          ` + "`" + `json:"sort" default:"id DESC" dc:"排序"` + "`" + `
	Page   int             ` + "`" + `json:"page" v:"integer|min:1" default:"1" dc:"页码"` + "`" + `
	Limit  int             ` + "`" + `json:"limit" v:"integer|min:0" default:"10" dc:"每页数量。可传0取全部"` + "`" + `
}

type ` + tpl.TableNameCaseCamel + `ListFilter struct {
	Id             *uint       ` + "`" + `json:"id,omitempty" v:"integer|min:1" dc:"ID"` + "`" + `
	IdArr          []uint      ` + "`" + `json:"idArr,omitempty" v:"distinct|foreach|integer|foreach|min:1" dc:"ID数组"` + "`" + `
	ExcId          *uint       ` + "`" + `json:"excId,omitempty" v:"integer|min:1" dc:"排除ID"` + "`" + `
	ExcIdArr       []uint      ` + "`" + `json:"excIdArr,omitempty" v:"distinct|foreach|integer|foreach|min:1" dc:"排除ID数组"` + "`" + `
	` + apiReqFilterColumn + `
}

type ` + tpl.TableNameCaseCamel + `ListRes struct {`
		if option.IsCount {
			tplApi += `
	Count int         ` + "`" + `json:"count" dc:"总数"` + "`"
		}
		tplApi += `
	List  []` + tpl.TableNameCaseCamel + `Item ` + "`" + `json:"list" dc:"列表"` + "`" + `
}

type ` + tpl.TableNameCaseCamel + `Item struct {
	Id          *uint       ` + "`" + `json:"id,omitempty" dc:"ID"` + "`" + `
	` + apiResColumn + `
	` + apiResColumnAlloweFieldList + `
}

/*--------列表 结束--------*/

`
	}
	if option.IsInfo {
		tplApi += `/*--------详情 开始--------*/
type ` + tpl.TableNameCaseCamel + `InfoReq struct {
	g.Meta ` + "`" + `path:"/` + tpl.TableNameCaseCamelLower + `/info" method:"post" tags:"` + tpl.SceneName + `/` + option.CommonName + `" sm:"详情"` + "`" + `
	Id     uint     ` + "`" + `json:"id" v:"required|integer|min:1" dc:"ID"` + "`" + `
	Field  []string ` + "`" + `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"` + "`" + `
}

type ` + tpl.TableNameCaseCamel + `InfoRes struct {
	Info ` + tpl.TableNameCaseCamel + `Info ` + "`" + `json:"info" dc:"详情"` + "`" + `
}

type ` + tpl.TableNameCaseCamel + `Info struct {
	Id          *uint       ` + "`" + `json:"id,omitempty" dc:"ID"` + "`" + `
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
	IdArr       []uint  ` + "`" + `json:"idArr,omitempty" v:"required|distinct|foreach|integer|foreach|min:1" dc:"ID数组"` + "`" + `
	` + apiReqUpdateColumn + `
}

/*--------修改 结束--------*/

`
	}

	if option.IsDelete {
		tplApi += `/*--------删除 开始--------*/
type ` + tpl.TableNameCaseCamel + `DeleteReq struct {
	g.Meta ` + "`" + `path:"/` + tpl.TableNameCaseCamelLower + `/del" method:"post" tags:"` + tpl.SceneName + `/` + option.CommonName + `" sm:"删除"` + "`" + `
	IdArr  []uint ` + "`" + `json:"idArr,omitempty" v:"required|distinct|foreach|integer|foreach|min:1" dc:"ID数组"` + "`" + `
}

/*--------删除 结束--------*/
`
	}

	if option.IsList && tpl.PidHandle.PidField != `` {
		tplApi += `
/*--------列表（树状） 开始--------*/
type ` + tpl.TableNameCaseCamel + `TreeReq struct {
	g.Meta ` + "`" + `path:"/` + tpl.TableNameCaseCamelLower + `/tree" method:"post" tags:"` + tpl.SceneName + `/` + option.CommonName + `" sm:"列表（树状）"` + "`" + `
	Field  []string       ` + "`" + `json:"field" v:"foreach|min-length:1"` + "`" + `
	Filter ` + tpl.TableNameCaseCamel + `ListFilter ` + "`" + `json:"filter" dc:"过滤条件"` + "`" + `
}

type ` + tpl.TableNameCaseCamel + `TreeRes struct {
	Tree []` + tpl.TableNameCaseCamel + `Tree ` + "`" + `json:"tree" dc:"列表（树状）"` + "`" + `
}

type ` + tpl.TableNameCaseCamel + `Tree struct {
	Id       *uint       ` + "`" + `json:"id,omitempty" dc:"ID"` + "`" + `
	Children interface{} ` + "`" + `json:"children" dc:"子级列表"` + "`" + `
	` + apiResColumn + `
}

/*--------列表（树状） 结束--------*/
`
	}

	gfile.PutContents(saveFile, tplApi)
	utils.GoFileFmt(saveFile)
}

// controller模板生成
func MyGenTplController(ctx context.Context, option *MyGenOption, tpl *MyGenTpl) {
	saveFile := gfile.SelfDir() + `/internal/controller/` + option.SceneCode + `/` + tpl.ModuleDirCaseCamelLower + `/` + tpl.TableNameCaseSnake + `.go`
	if !option.IsCover && gfile.IsFile(saveFile) {
		return
	}

	controllerAlloweFieldList := "`id`, "
	controllerAlloweFieldInfo := "`id`, "
	controllerAlloweFieldTree := "`id`, "
	controllerAlloweFieldNoAuth := "`id`, "
	if tpl.LabelField != `` {
		controllerAlloweFieldList += "`label`, "
		controllerAlloweFieldInfo += "`label`, "
		controllerAlloweFieldTree += "`label`, "
		if tpl.PidHandle.PidField != `` {
			controllerAlloweFieldList += "`p" + gstr.CaseCamel(tpl.LabelField) + "`, "
			// controllerAlloweFieldInfo += "`p" + gstr.CaseCamel(tpl.LabelField) + "`, "
		}
		controllerAlloweFieldNoAuth += "`label`, "
		controllerAlloweFieldNoAuth += `columnsThis.` + gstr.CaseCamel(tpl.LabelField) + `, `
	}
	controllerAlloweFieldDiff := ``
	for _, column := range tpl.TableColumnList {
		field := column[`Field`].String()
		fieldCaseCamel := gstr.CaseCamel(field)
		//主键
		if column[`Key`].String() == `PRI` && column[`Extra`].String() == `auto_increment` {
			if field != `id` {
				controllerAlloweFieldNoAuth += `columnsThis.` + fieldCaseCamel + `, `
			}
			continue
		}
		//password|passwd
		if garray.NewStrArrayFrom([]string{`password`, `passwd`}).Contains(field) && column[`Type`].String() == `char(32)` {
			controllerAlloweFieldDiff += `columnsThis.` + fieldCaseCamel + `, `
			continue
		}
		//salt
		if garray.NewStrArrayFrom([]string{`salt`}).Contains(field) && column[`Type`].String() == `char(8)` && tpl.PasswordHandle.IsCoexist {
			controllerAlloweFieldDiff += `columnsThis.` + fieldCaseCamel + `, `
			continue
		}
		//id后缀
		if gstr.SubStr(fieldCaseCamel, -2) == `Id` && gstr.Pos(column[`Type`].String(), `int`) != -1 {
			if tpl.RelTableMap[field].IsExistRelTableDao {
				relTable := tpl.RelTableMap[field]
				controllerAlloweFieldList += "`" + relTable.RelNameFieldCaseCamelLower + "`, "
			}
			continue
		}
	}
	controllerAlloweFieldList = gstr.SubStr(controllerAlloweFieldList, 0, -len(`, `))
	controllerAlloweFieldInfo = gstr.SubStr(controllerAlloweFieldInfo, 0, -len(`, `))
	controllerAlloweFieldTree = gstr.SubStr(controllerAlloweFieldTree, 0, -len(`, `))
	controllerAlloweFieldNoAuth = gstr.SubStr(controllerAlloweFieldNoAuth, 0, -len(`, `))
	controllerAlloweFieldDiff = gstr.SubStr(controllerAlloweFieldDiff, 0, -len(`, `))

	tplController := `package controller

import (
	"api/api"
	api` + tpl.ModuleDirCaseCamel + ` "api/api/` + option.SceneCode + `/` + tpl.ModuleDirCaseCamelLower + `"
	"api/internal/dao"
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
		tplController += `
// 列表
func (controllerThis *` + tpl.TableNameCaseCamel + `) List(ctx context.Context, req *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `ListReq) (res *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `ListRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.MapDeep(req.Filter)
	if filter == nil {
		filter = map[string]interface{}{}
	}
	order := []string{req.Sort}
	page := req.Page
	limit := req.Limit
`
		if controllerAlloweFieldDiff != `` || (option.IsAuthAction && controllerAlloweFieldNoAuth != ``) {
			tplController += `
	columnsThis := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `.Columns()`
		}
		tplController += `
	allowField := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `.ColumnArr()
	allowField = append(allowField, ` + controllerAlloweFieldList + `)`
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
			actionCode := gstr.CaseCamelLower(tpl.LogicStructName) + `Look`
			actionName := option.CommonName + `-查看`
			MyGenAction(ctx, tpl.SceneId, actionCode, actionName) // 数据库权限操作处理
			tplController += `
	/**--------权限验证 开始--------**/
	isAuth, _ := service.AuthAction().CheckAuth(ctx, ` + "`" + actionCode + "`" + `)
	if !isAuth {
		field = []string{` + controllerAlloweFieldNoAuth + `}
	}
	/**--------权限验证 结束--------**/
`
		}
		tplController += `
	daoHandlerThis := dao.NewDaoHandler(ctx, &dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `)
	daoHandlerThis.Filter(filter)`
		if option.IsCount {
			tplController += `
	count, err := daoHandlerThis.Count()
	if err != nil {
		return
	}`
		}
		tplController += `
	list, err := daoHandlerThis.Field(field).Order(order).JoinGroupByPrimaryKey().GetModel().Page(page, limit).All()
	if err != nil {
		return
	}

	res = &api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `ListRes{`
		if option.IsCount {
			tplController += `
		Count: count,
	`
		}
		tplController += `}
	list.Structs(&res.List)
	return
}
`
	}
	if option.IsInfo {
		tplController += `
// 详情
func (controllerThis *` + tpl.TableNameCaseCamel + `) Info(ctx context.Context, req *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `InfoReq) (res *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `InfoRes, err error) {
	/**--------参数处理 开始--------**/
	allowField := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `.ColumnArr()
	allowField = append(allowField, ` + controllerAlloweFieldInfo + `)`
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
			actionCode := gstr.CaseCamelLower(tpl.LogicStructName) + `Look`
			actionName := option.CommonName + `-查看`
			MyGenAction(ctx, tpl.SceneId, actionCode, actionName) // 数据库权限操作处理
			tplController += `
	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, ` + "`" + actionCode + "`" + `)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/
`
		}
		tplController += `
	info, err := dao.NewDaoHandler(ctx, &dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `).Filter(filter).Field(field).JoinGroupByPrimaryKey().GetModel().One()
	if err != nil {
		return
	}
	if len(info) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ` + "``" + `)
		return
	}

	res = &api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `InfoRes{}
	info.Struct(&res.Info)
	return
}
`
	}
	if option.IsCreate {
		tplController += `
// 新增
func (controllerThis *` + tpl.TableNameCaseCamel + `) Create(ctx context.Context, req *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `CreateReq) (res *api.CommonCreateRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.MapDeep(req)
	/**--------参数处理 结束--------**/
`
		if option.IsAuthAction {
			actionCode := gstr.CaseCamelLower(tpl.LogicStructName) + `Create`
			actionName := option.CommonName + `-新增`
			MyGenAction(ctx, tpl.SceneId, actionCode, actionName) // 数据库权限操作处理
			tplController += `
	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, ` + "`" + actionCode + "`" + `)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/
`
		}
		tplController += `
	id, err := service.` + tpl.LogicStructName + `().Create(ctx, data)
	if err != nil {
		return
	}
	res = &api.CommonCreateRes{Id: id}
	return
}
`
	}

	if option.IsUpdate {
		tplController += `
// 修改
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
			actionCode := gstr.CaseCamelLower(tpl.LogicStructName) + `Update`
			actionName := option.CommonName + `-编辑`
			MyGenAction(ctx, tpl.SceneId, actionCode, actionName) // 数据库权限操作处理
			tplController += `
	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, ` + "`" + actionCode + "`" + `)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/
`
		}
		tplController += `
	_, err = service.` + tpl.LogicStructName + `().Update(ctx, filter, data)
	return
}
`
	}

	if option.IsDelete {
		tplController += `
// 删除
func (controllerThis *` + tpl.TableNameCaseCamel + `) Delete(ctx context.Context, req *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `DeleteReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	filter := map[string]interface{}{` + "`id`" + `: req.IdArr}
	/**--------参数处理 结束--------**/
`
		if option.IsAuthAction {
			actionCode := gstr.CaseCamelLower(tpl.LogicStructName) + `Delete`
			actionName := option.CommonName + `-删除`
			MyGenAction(ctx, tpl.SceneId, actionCode, actionName) // 数据库权限操作处理
			tplController += `
	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, ` + "`" + actionCode + "`" + `)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/
`
		}
		tplController += `
	_, err = service.` + tpl.LogicStructName + `().Delete(ctx, filter)
	return
}
`
	}

	if option.IsList && tpl.PidHandle.PidField != `` {
		tplController += `
// 列表（树状）
func (controllerThis *` + tpl.TableNameCaseCamel + `) Tree(ctx context.Context, req *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `TreeReq) (res *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `TreeRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.MapDeep(req.Filter)
	if filter == nil {
		filter = map[string]interface{}{}
	}

	columnsThis := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `.Columns()
	allowField := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `.ColumnArr()
	allowField = append(allowField, ` + controllerAlloweFieldTree + `)`
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
			actionCode := gstr.CaseCamelLower(tpl.LogicStructName) + `Look`
			actionName := option.CommonName + `-查看`
			MyGenAction(ctx, tpl.SceneId, actionCode, actionName) // 数据库权限操作处理
			tplController += `
	/**--------权限验证 开始--------**/
	isAuth, _ := service.AuthAction().CheckAuth(ctx, ` + "`" + actionCode + "`" + `)
	if !isAuth {
		field = []string{` + controllerAlloweFieldNoAuth + `}
	}
	/**--------权限验证 结束--------**/
`
		}
		tplController += `
	field = append(field, ` + "`tree`" + `)

	list, err := dao.NewDaoHandler(ctx, &dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `).Filter(filter).Field(field).JoinGroupByPrimaryKey().GetModel().All()
	if err != nil {
		return
	}
	tree := utils.Tree(list, 0, columnsThis.` + gstr.CaseCamel(tpl.PrimaryKey) + `, columnsThis.` + gstr.CaseCamel(tpl.PidHandle.PidField) + `)

	utils.HttpWriteJson(ctx, map[string]interface{}{
		` + "`tree`" + `: tree,
	}, 0, ` + "``" + `)
	return
}
`
	}

	gfile.PutContents(saveFile, tplController)
	utils.GoFileFmt(saveFile)
}

// 后端路由生成
func MyGenTplRouter(ctx context.Context, option *MyGenOption, tpl *MyGenTpl) {
	saveFile := gfile.SelfDir() + `/internal/router/` + option.SceneCode + `.go`

	tplRouter := gfile.GetContents(saveFile)

	//控制器不存在时导入
	importControllerStr := `controller` + tpl.ModuleDirCaseCamel + ` "api/internal/controller/` + option.SceneCode + `/` + tpl.ModuleDirCaseCamelLower + `"`
	if gstr.Pos(tplRouter, importControllerStr) == -1 {
		tplRouter = gstr.Replace(tplRouter, `"api/internal/middleware"`, importControllerStr+`
	"api/internal/middleware"`)
		//路由生成
		tplRouter = gstr.Replace(tplRouter, `/*--------后端路由自动代码生成锚点（不允许修改和删除，否则将不能自动生成路由）--------*/`, `group.Group(`+"`"+`/`+tpl.ModuleDirCaseCamelLower+"`"+`, func(group *ghttp.RouterGroup) {
				group.Bind(controller`+tpl.ModuleDirCaseCamel+`.New`+tpl.TableNameCaseCamel+`())
			})

			/*--------后端路由自动代码生成锚点（不允许修改和删除，否则将不能自动生成路由）--------*/`)
		gfile.PutContents(saveFile, tplRouter)
	} else {
		//路由不存在时需生成
		if gstr.Pos(tplRouter, `group.Bind(controller`+tpl.ModuleDirCaseCamel+`.New`+tpl.TableNameCaseCamel+`())`) == -1 {
			//路由生成
			tplRouter = gstr.Replace(tplRouter, `group.Group(`+"`"+`/`+tpl.ModuleDirCaseCamelLower+"`"+`, func(group *ghttp.RouterGroup) {`, `group.Group(`+"`"+`/`+tpl.ModuleDirCaseCamelLower+"`"+`, func(group *ghttp.RouterGroup) {
				group.Bind(controller`+tpl.ModuleDirCaseCamel+`.New`+tpl.TableNameCaseCamel+`())`)
			gfile.PutContents(saveFile, tplRouter)
		}
	}

	utils.GoFileFmt(saveFile)
}

// 视图模板Index生成
func MyGenTplViewIndex(ctx context.Context, option *MyGenOption, tpl *MyGenTpl) {
	saveFile := gfile.SelfDir() + `/../view/` + option.SceneCode + `/src/views/` + tpl.ModuleDirCaseCamelLower + `/` + tpl.TableNameCaseCamelLower + `/Index.vue`
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
	saveFile := gfile.SelfDir() + `/../view/` + option.SceneCode + `/src/views/` + tpl.ModuleDirCaseCamelLower + `/` + tpl.TableNameCaseCamelLower + `/List.vue`
	if !option.IsCover && gfile.IsFile(saveFile) {
		return
	}

	tableRowHeight := 50
	viewListColumn := ``
	for _, column := range tpl.TableColumnList {
		field := column[`Field`].String()
		fieldCaseCamel := gstr.CaseCamel(field)
		fieldCaseSnake := gstr.CaseSnakeFirstUpper(field)
		comment := gstr.Trim(gstr.ReplaceByArray(column[`Comment`].String(), g.SliceStr{
			"\n", ` `,
			"\r", ` `,
			`"`, `\"`,
		}))

		if garray.NewStrArrayFrom([]string{`DeletedAt`, `DeleteAt`, `DeletedTime`, `DeleteTime`}).Contains(fieldCaseCamel) {
			continue
		}
		if garray.NewStrArrayFrom([]string{`UpdatedAt`, `UpdateAt`, `UpdatedTime`, `UpdateTime`}).Contains(fieldCaseCamel) {
			viewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('common.name.updatedAt'),
		key: '` + field + `',
		align: 'center',
		width: 150,
		sortable: true,
	},`
			continue
		}
		if garray.NewStrArrayFrom([]string{`CreatedAt`, `CreateAt`, `CreatedTime`, `CreateTime`}).Contains(fieldCaseCamel) {
			viewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('common.name.createdAt'),
		key: '` + field + `',
		align: 'center',
		width: 150,
		sortable: true
	},`
			continue
		}
		//主键
		if column[`Key`].String() == `PRI` && column[`Extra`].String() == `auto_increment` {
			continue
		}
		//password|passwd
		if garray.NewStrArrayFrom([]string{`password`, `passwd`}).Contains(field) && column[`Type`].String() == `char(32)` {
			continue
		}
		//salt
		if garray.NewStrArrayFrom([]string{`salt`}).Contains(field) && column[`Type`].String() == `char(8)` && tpl.PasswordHandle.IsCoexist {
			continue
		}
		//pid
		if garray.NewStrArrayFrom([]string{`pid`}).Contains(field) && gstr.Pos(column[`Type`].String(), `int`) != -1 {
			viewListColumn += `
	{
		dataKey: 'p` + gstr.CaseCamel(tpl.LabelField) + `',
		title: t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
		key: '` + field + `',
		align: 'center',
		width: 150,
	},`
			continue
		}
		//level
		if garray.NewStrArrayFrom([]string{`level`}).Contains(field) && gstr.Pos(column[`Type`].String(), `int`) != -1 && tpl.PidHandle.IsCoexist {
			viewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
		key: '` + field + `',
		align: 'center',
        width: 100,
        sortable: true,
	},`
			continue
		}
		//idPath|id_path
		if fieldCaseCamel == `IdPath` && gstr.Pos(column[`Type`].String(), `varchar`) != -1 && tpl.PidHandle.IsCoexist {
			viewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
		key: '` + field + `',
		align: 'center',
        width: 150,
        hidden: true,
	},`
			continue

		}
		//sort
		//weight
		if garray.NewStrArrayFrom([]string{`sort`, `weight`}).Contains(field) && gstr.Pos(column[`Type`].String(), `int`) != -1 {
			viewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
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
						'placeholder': t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.tip.` + field + `'),
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
			continue
		}
		//id后缀
		if gstr.SubStr(fieldCaseCamel, -2) == `Id` && gstr.Pos(column[`Type`].String(), `int`) != -1 {
			relNameField := field
			if tpl.RelTableMap[field].IsExistRelTableDao {
				relTable := tpl.RelTableMap[field]
				relNameField = relTable.RelNameFieldCaseCamelLower
			}
			viewListColumn += `
	{
		dataKey: '` + relNameField + `',
		title: t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
		key: '` + field + `',
		align: 'center',
		width: 150,
	},`
			continue
		}
		//name或code后缀
		if (gstr.SubStr(fieldCaseCamel, -4) == `Name` || gstr.SubStr(fieldCaseCamel, -4) == `Code`) && gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
			viewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
		key: '` + field + `',
		align: 'center',
		width: 150,
	},`
			continue
		}
		//mobile,phone后缀
		if (gstr.SubStr(fieldCaseCamel, -5) == `Phone` || gstr.SubStr(fieldCaseCamel, -6) == `Mobile`) && gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
			viewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
		key: '` + field + `',
		align: 'center',
		width: 150,
	},`
			continue
		}
		//url,link后缀
		if (gstr.SubStr(fieldCaseCamel, -3) == `Url` || gstr.SubStr(fieldCaseCamel, -4) == `Link`) && gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
			viewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
		key: '` + field + `',
		align: 'center',
		width: 300,
	},`
			continue
		}
		//avatar
		//icon,cover,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀
		if field == `avatar` || gstr.SubStr(fieldCaseCamel, -4) == `Icon` || gstr.SubStr(fieldCaseCamel, -5) == `Cover` || gstr.SubStr(fieldCaseCamel, -3) == `Img` || gstr.SubStr(fieldCaseCamel, -7) == `ImgList` || gstr.SubStr(fieldCaseCamel, -6) == `ImgArr` || gstr.SubStr(fieldCaseCamel, -5) == `Image` || gstr.SubStr(fieldCaseCamel, -9) == `ImageList` || gstr.SubStr(fieldCaseCamel, -8) == `ImageArr` {
			viewListColumn += `
	{
        dataKey: '` + field + `',
        title: t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
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
        title: t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
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
		//list,arr等后缀
		if (gstr.SubStr(fieldCaseCamel, -4) == `List` || gstr.SubStr(fieldCaseCamel, -3) == `Arr`) && (column[`Type`].String() == `json` || column[`Type`].String() == `text`) {
			viewListColumn += `
	{
        dataKey: '` + field + `',
        title: t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
        key: '` + field + `',
        align: 'center',
        width: 100,
        cellRenderer: (props: any): any => {
            if (!props.rowData.` + field + `) {
                return
            }`
			if column[`Type`].String() == `json` || column[`Type`].String() == `text` {
				viewListColumn += `
			let arrList: any[]
			if (Array.isArray(props.rowData.` + field + `)) {
				arrList = props.rowData.` + field + `
			} else {
				arrList = JSON.parse(props.rowData.` + field + `)
			}`
			} else {
				viewListColumn += `
			const arrList = [props.rowData.` + field + `]`
			}
			viewListColumn += `
			let typeArr: string[] = ['', 'success', 'danger', 'info', 'warning']
            return [
                h(ElScrollbar, {
                    'wrap-style': 'display: flex; align-items: center;',
                    'view-style': 'margin: auto;',
                }, {
                    default: () => {
                        const content = arrList.map((item, index) => {
                            return h(ElTag as any, {
                                'style': 'margin: auto 5px 5px auto;',
                                'type': typeArr[index % 5]
                            }, {
								default: () => {
									return item
								}
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
		//IP后缀
		if gstr.SubStr(fieldCaseCamel, -2) == `Ip` && gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
			viewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
		key: '` + field + `',
		align: 'center',
		width: 150,
	},`
			continue
		}
		//remark,desc,msg,message,intro,content后缀
		if (gstr.SubStr(fieldCaseCamel, -6) == `Remark` || gstr.SubStr(fieldCaseCamel, -4) == `Desc` || gstr.SubStr(fieldCaseCamel, -3) == `Msg` || gstr.SubStr(fieldCaseCamel, -7) == `Message` || gstr.SubStr(fieldCaseCamel, -5) == `Intro` || gstr.SubStr(fieldCaseCamel, -7) == `Content`) && (gstr.Pos(column[`Type`].String(), `varchar`) != -1 || column[`Type`].String() == `text`) {
			viewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
		key: '` + field + `',
		align: 'center',
		width: 150,
		hidden: true,
	},`
			continue
		}
		//gender
		//status或type后缀
		if (field == `gender` || gstr.SubStr(fieldCaseCamel, -6) == `Status` || gstr.SubStr(fieldCaseCamel, -4) == `Type`) && gstr.Pos(column[`Type`].String(), `int`) != -1 {
			statusList := MyGenStatusList(comment)
			tagTypeStr := ``
			tagTypeArr := []string{``, `success`, `danger`, `info`, `warning`}
			tagTypeLen := len(tagTypeArr)
			for index, status := range statusList {
				tagTypeStr += status[0] + `: '` + tagTypeArr[index%tagTypeLen] + `', `
			}
			tagTypeStr = gstr.SubStr(tagTypeStr, 0, -len(`, `))
			viewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
		key: '` + field + `',
		align: 'center',
		width: 100,
		cellRenderer: (props: any): any => {
			let typeObj: any = { ` + tagTypeStr + ` }
			return [
				h(ElTag as any, {
					type: typeObj[props.rowData.` + field + `]
				}, {
					default: () => (tm('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.status.` + field + `') as any).find((item: any) => { return item.value == props.rowData.` + field + ` })?.label
				})
			]
		}
	},`
			continue
		}
		//is_前缀
		if gstr.SubStr(fieldCaseSnake, 0, 3) == `is_` && gstr.Pos(column[`Type`].String(), `int`) != -1 {
			viewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
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
		//start_前缀
		//end_前缀
		/* if gstr.SubStr(fieldCaseSnake, 0, 6) == `start_` || gstr.SubStr(fieldCaseSnake, 0, 4) == `end_` {
			if gstr.Pos(column[`Type`].String(), `timestamp`) != -1 || gstr.Pos(column[`Type`].String(), `datetime`) != -1 {
				continue
			}
			if gstr.Pos(column[`Type`].String(), `date`) != -1 {
				continue
			}
		} */
		//int类型
		if gstr.Pos(column[`Type`].String(), `int`) != -1 {
			viewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
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
		title: t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
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
		title: t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
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
		title: t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
		key: '` + field + `',
		align: 'center',
		width: 150,
	},`
			continue
		}
		//timestamp或datetime类型
		if gstr.Pos(column[`Type`].String(), `timestamp`) != -1 || gstr.Pos(column[`Type`].String(), `datetime`) != -1 {
			viewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
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
		title: t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
		key: '` + field + `',
		align: 'center',
		width: 100,
        sortable: true
	},`
			continue
		}
		//json类型
		//text类型
		if column[`Type`].String() == `json` || gstr.Pos(column[`Type`].String(), `text`) != -1 {
			viewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
		key: '` + field + `',
		align: 'center',
		width: 200,
        hidden: true
	},`
			continue
		}
		//默认处理
		viewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),
		key: '` + field + `',
		align: 'center',
		width: 150,
	},`
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
	},` + viewListColumn
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
                <MyExportButton :headerList="table.columns"
                    :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseCamelLower + `/` + tpl.TableNameCaseCamelLower + `/list', param: { filter: queryCommon.data, sort: table.sort.key + ' ' + table.sort.order } }" />
				<ElDropdown max-height="300" :hide-on-click="false">
					<ElButton type="info" :circle="true">
						<AutoiconEpHide />
					</ElButton>
					<template #dropdown>
						<ElDropdownMenu>
							<ElDropdownItem v-for="(item, index) in table.columns" :key="index">
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
	saveFile := gfile.SelfDir() + `/../view/` + option.SceneCode + `/src/views/` + tpl.ModuleDirCaseCamelLower + `/` + tpl.TableNameCaseCamelLower + `/Query.vue`
	if !option.IsCover && gfile.IsFile(saveFile) {
		return
	}

	viewQueryDataInit := ``
	viewQueryField := ``
	for _, column := range tpl.TableColumnList {
		field := column[`Field`].String()
		fieldCaseCamel := gstr.CaseCamel(field)
		fieldCaseSnake := gstr.CaseSnakeFirstUpper(field)
		resultStr, _ := gregex.MatchString(`.*\((\d*)\)`, column[`Type`].String())
		resultFloat, _ := gregex.MatchString(`.*\((\d*),(\d*)\)`, column[`Type`].String())
		if len(resultFloat) < 3 {
			resultFloat = []string{``, `10`, `2`}
		}

		if garray.NewStrArrayFrom([]string{`DeletedAt`, `DeleteAt`, `DeletedTime`, `DeleteTime`}).Contains(fieldCaseCamel) {
			continue
		}
		if garray.NewStrArrayFrom([]string{`UpdatedAt`, `UpdateAt`, `UpdatedTime`, `UpdateTime`}).Contains(fieldCaseCamel) {
			continue
		}
		if garray.NewStrArrayFrom([]string{`CreatedAt`, `CreateAt`, `CreatedTime`, `CreateTime`}).Contains(fieldCaseCamel) {
			viewQueryDataInit += `
	timeRange: (() => {
		// const date = new Date()
		return [
			// new Date(date.getFullYear(), date.getMonth(), date.getDate(), 0, 0, 0),
			// new Date(date.getFullYear(), date.getMonth(), date.getDate(), 23, 59, 59),
		]
	})(),
	timeRangeStart: computed(() => {
		if (queryCommon.data.timeRange?.length) {
			return dayjs(queryCommon.data.timeRange[0]).format('YYYY-MM-DD HH:mm:ss')
		}
		return ''
	}),
	timeRangeEnd: computed(() => {
		if (queryCommon.data.timeRange?.length) {
			return dayjs(queryCommon.data.timeRange[1]).format('YYYY-MM-DD HH:mm:ss')
		}
		return ''
	}),`
			viewQueryField += `
		<ElFormItem prop="timeRange">
			<ElDatePicker v-model="queryCommon.data.timeRange" type="datetimerange" range-separator="-" :default-time="[new Date(2000, 0, 1, 0, 0, 0), new Date(2000, 0, 1, 23, 59, 59)]" :start-placeholder="t('common.name.timeRangeStart')" :end-placeholder="t('common.name.timeRangeEnd')" />
		</ElFormItem>`
			continue
		}
		//主键
		if column[`Key`].String() == `PRI` && column[`Extra`].String() == `auto_increment` {
			continue
		}
		//password|passwd
		if garray.NewStrArrayFrom([]string{`password`, `passwd`}).Contains(field) && column[`Type`].String() == `char(32)` {
			continue
		}
		//salt
		if garray.NewStrArrayFrom([]string{`salt`}).Contains(field) && column[`Type`].String() == `char(8)` && tpl.PasswordHandle.IsCoexist {
			continue
		}
		//pid
		if garray.NewStrArrayFrom([]string{`pid`}).Contains(field) && gstr.Pos(column[`Type`].String(), `int`) != -1 {
			viewQueryField += `
		<ElFormItem prop="` + field + `">
			<MyCascader v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseCamelLower + `/` + tpl.TableNameCaseCamelLower + `/tree' }" :defaultOptions="[{ id: 0, label: t('common.name.allTopLevel') }]" :props="{ checkStrictly: true, emitPath: false }" />
		</ElFormItem>`
			continue
		}
		/* //level
		if garray.NewStrArrayFrom([]string{`level`}).Contains(field) && gstr.Pos(column[`Type`].String(), `int`) != -1 && tpl.PidHandle.IsCoexist {
			continue
		} */
		//idPath|id_path
		if fieldCaseCamel == `IdPath` && gstr.Pos(column[`Type`].String(), `varchar`) != -1 && tpl.PidHandle.IsCoexist {
			continue
		}
		//sort
		//weight
		if garray.NewStrArrayFrom([]string{`sort`, `weight`}).Contains(field) && gstr.Pos(column[`Type`].String(), `int`) != -1 {
			continue
		}
		//id后缀
		if gstr.SubStr(fieldCaseCamel, -2) == `Id` && gstr.Pos(column[`Type`].String(), `int`) != -1 {
			apiUrl := tpl.ModuleDirCaseCamelLower + `/` + gstr.CaseCamelLower(gstr.SubStr(field, 0, -2))
			if tpl.RelTableMap[field].IsExistRelTableDao {
				relTable := tpl.RelTableMap[field]
				apiUrl = relTable.RelDaoDirCaseCamelLower + `/` + relTable.RelTableNameCaseCamelLower
			}
			viewQueryField += `
		<ElFormItem prop="` + field + `">
			<MySelect v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + apiUrl + `/list' }" />
		</ElFormItem>`
			continue
		}
		//name或code后缀
		if (gstr.SubStr(fieldCaseCamel, -4) == `Name` || gstr.SubStr(fieldCaseCamel, -4) == `Code`) && gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
			viewQueryField += `
		<ElFormItem prop="` + field + `">
			<ElInput v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" minlength="1" maxlength="` + resultStr[1] + `" :clearable="true" />
		</ElFormItem>`
			continue
		}
		//mobile,phone后缀
		if (gstr.SubStr(fieldCaseCamel, -5) == `Phone` || gstr.SubStr(fieldCaseCamel, -6) == `Mobile`) && gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
			viewQueryField += `
		<ElFormItem prop="` + field + `">
			<ElInput v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" minlength="1" maxlength="` + resultStr[1] + `" :clearable="true" />
		</ElFormItem>`
			continue
		}
		//url,link后缀
		if (gstr.SubStr(fieldCaseCamel, -3) == `Url` || gstr.SubStr(fieldCaseCamel, -4) == `Link`) && gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
			viewQueryField += `
		<ElFormItem prop="` + field + `">
			<ElInput v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" minlength="1" maxlength="` + resultStr[1] + `" :clearable="true" />
		</ElFormItem>`
			continue
		}
		//avatar
		//icon,cover,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀
		//video,video_list,videoList,video_arr,videoArr等后缀
		if field == `avatar` || gstr.SubStr(fieldCaseCamel, -4) == `Icon` || gstr.SubStr(fieldCaseCamel, -5) == `Cover` || gstr.SubStr(fieldCaseCamel, -3) == `Img` || gstr.SubStr(fieldCaseCamel, -7) == `ImgList` || gstr.SubStr(fieldCaseCamel, -6) == `ImgArr` || gstr.SubStr(fieldCaseCamel, -5) == `Image` || gstr.SubStr(fieldCaseCamel, -9) == `ImageList` || gstr.SubStr(fieldCaseCamel, -8) == `ImageArr` || gstr.SubStr(fieldCaseCamel, -5) == `Video` || gstr.SubStr(fieldCaseCamel, -9) == `VideoList` || gstr.SubStr(fieldCaseCamel, -8) == `VideoArr` {
			continue
		}
		//list,arr等后缀
		if (gstr.SubStr(fieldCaseCamel, -4) == `List` || gstr.SubStr(fieldCaseCamel, -3) == `Arr`) && (column[`Type`].String() == `json` || column[`Type`].String() == `text`) {
			continue
		}
		//IP后缀
		if gstr.SubStr(fieldCaseCamel, -2) == `Ip` && gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
			viewQueryField += `
		<ElFormItem prop="` + field + `">
			<ElInput v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" minlength="1" maxlength="` + resultStr[1] + `" :clearable="true" />
		</ElFormItem>`
			continue
		}
		//gender
		//status或type后缀
		if (field == `gender` || gstr.SubStr(fieldCaseCamel, -6) == `Status` || gstr.SubStr(fieldCaseCamel, -4) == `Type`) && gstr.Pos(column[`Type`].String(), `int`) != -1 {
			viewQueryField += `
		<ElFormItem prop="` + field + `" style="width: 120px;">
			<ElSelectV2 v-model="queryCommon.data.` + field + `" :options="tm('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.status.` + field + `')" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :clearable="true" />
		</ElFormItem>`
			continue
		}
		//remark,desc,msg,message,intro,content后缀
		if (gstr.SubStr(fieldCaseCamel, -6) == `Remark` || gstr.SubStr(fieldCaseCamel, -4) == `Desc` || gstr.SubStr(fieldCaseCamel, -3) == `Msg` || gstr.SubStr(fieldCaseCamel, -7) == `Message` || gstr.SubStr(fieldCaseCamel, -5) == `Intro` || gstr.SubStr(fieldCaseCamel, -7) == `Content`) && (gstr.Pos(column[`Type`].String(), `varchar`) != -1 || column[`Type`].String() == `text`) {
			continue
		}
		//is_前缀
		if gstr.SubStr(fieldCaseSnake, 0, 3) == `is_` && gstr.Pos(column[`Type`].String(), `int`) != -1 {
			viewQueryField += `
		<ElFormItem prop="` + field + `" style="width: 120px;">
			<ElSelectV2 v-model="queryCommon.data.` + field + `" :options="tm('common.status.whether')" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :clearable="true" />
		</ElFormItem>`
			continue
		}
		//start_前缀
		if gstr.SubStr(fieldCaseSnake, 0, 6) == `start_` {
			if gstr.Pos(column[`Type`].String(), `timestamp`) != -1 || gstr.Pos(column[`Type`].String(), `datetime`) != -1 {
				viewQueryField += `
		<ElFormItem prop="` + field + `">
			<ElDatePicker v-model="queryCommon.data.` + field + `" type="datetime" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :default-time="new Date(2000, 0, 1, 0, 0, 0)" format="YYYY-MM-DD HH:mm:ss" value-format="YYYY-MM-DD HH:mm:ss" />
		</ElFormItem>`
				continue
			}
			/* if gstr.Pos(column[`Type`].String(), `date`) != -1 {
				continue
			} */
		}
		//end_前缀
		if gstr.SubStr(fieldCaseSnake, 0, 4) == `end_` {
			if gstr.Pos(column[`Type`].String(), `timestamp`) != -1 || gstr.Pos(column[`Type`].String(), `datetime`) != -1 {
				viewQueryField += `
		<ElFormItem prop="` + field + `">
			<ElDatePicker v-model="queryCommon.data.` + field + `" type="datetime" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :default-time="new Date(2000, 0, 1, 23, 59, 59)" format="YYYY-MM-DD HH:mm:ss" value-format="YYYY-MM-DD HH:mm:ss" />
		</ElFormItem>`
				continue
			}
			/* if gstr.Pos(column[`Type`].String(), `date`) != -1 {
				continue
			} */
		}
		//int类型
		if gstr.Pos(column[`Type`].String(), `int`) != -1 {
			if gstr.Pos(column[`Type`].String(), `unsigned`) != -1 {
				viewQueryField += `
		<ElFormItem prop="` + field + `">
			<ElInputNumber v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :min="0" :controls="false" />
		</ElFormItem>`
			} else {
				viewQueryField += `
		<ElFormItem prop="` + field + `">
			<ElInputNumber v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :controls="false" />
		</ElFormItem>`
			}
			continue
		}
		//float类型
		if gstr.Pos(column[`Type`].String(), `decimal`) != -1 || gstr.Pos(column[`Type`].String(), `double`) != -1 || gstr.Pos(column[`Type`].String(), `float`) != -1 {
			if gstr.Pos(column[`Type`].String(), `unsigned`) != -1 {
				viewQueryField += `
		<ElFormItem prop="` + field + `">
			<ElInputNumber v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :min="0" :precision="` + resultFloat[2] + `" :controls="false" />
		</ElFormItem>`
			} else {
				viewQueryField += `
		<ElFormItem prop="` + field + `">
			<ElInputNumber v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :precision="` + resultFloat[2] + `" :controls="false" />
		</ElFormItem>`
			}
			continue
		}
		//varchar类型
		if gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
			viewQueryField += `
		<ElFormItem prop="` + field + `">
			<ElInput v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" minlength="1" maxlength="` + resultStr[1] + `" :clearable="true" />
		</ElFormItem>`
			continue
		}
		//char类型
		if gstr.Pos(column[`Type`].String(), `char`) != -1 {
			viewQueryField += `
		<ElFormItem prop="` + field + `">
			<ElInput v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" minlength="` + resultStr[1] + `" maxlength="` + resultStr[1] + `" :clearable="true" />
		</ElFormItem>`
			continue
		}
		//timestamp或datetime类型
		if gstr.Pos(column[`Type`].String(), `timestamp`) != -1 || gstr.Pos(column[`Type`].String(), `datetime`) != -1 {
			continue
		}
		//date类型
		if gstr.Pos(column[`Type`].String(), `date`) != -1 {
			viewQueryField += `
		<ElFormItem prop="` + field + `">
			<ElDatePicker v-model="queryCommon.data.` + field + `" type="date" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" format="YYYY-MM-DD" value-format="YYYY-MM-DD" />
		</ElFormItem>`
			continue
		}
		//json类型
		//text类型
		if column[`Type`].String() == `json` || gstr.Pos(column[`Type`].String(), `text`) != -1 {
			continue
		}
		//默认处理
		viewQueryField += `
		<ElFormItem prop="` + field + `">
			<ElInput v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :clearable="true" />
		</ElFormItem>`
	}

	tplView := `<script setup lang="ts">
import dayjs from 'dayjs'

const { t, tm } = useI18n()

const queryCommon = inject('queryCommon') as { data: { [propName: string]: any } }
queryCommon.data = {
	...queryCommon.data,` + viewQueryDataInit + `
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
	saveFile := gfile.SelfDir() + `/../view/` + option.SceneCode + `/src/views/` + tpl.ModuleDirCaseCamelLower + `/` + tpl.TableNameCaseCamelLower + `/Save.vue`
	if !option.IsCover && gfile.IsFile(saveFile) {
		return
	}

	viewSaveDataInit := ``
	viewSaveRule := ``
	viewSaveField := ``
	viewFieldHandle := ``
	for _, column := range tpl.TableColumnList {
		field := column[`Field`].String()
		fieldCaseCamel := gstr.CaseCamel(field)
		fieldCaseSnake := gstr.CaseSnakeFirstUpper(field)
		comment := gstr.Trim(gstr.ReplaceByArray(column[`Comment`].String(), g.SliceStr{
			"\n", ` `,
			"\r", ` `,
			`"`, `\"`,
		}))
		resultStr, _ := gregex.MatchString(`.*\((\d*)\)`, column[`Type`].String())
		resultFloat, _ := gregex.MatchString(`.*\((\d*),(\d*)\)`, column[`Type`].String())
		if len(resultFloat) < 3 {
			resultFloat = []string{``, `10`, `2`}
		}

		if garray.NewStrArrayFrom([]string{`DeletedAt`, `DeleteAt`, `DeletedTime`, `DeleteTime`}).Contains(fieldCaseCamel) {
			continue
		}
		if garray.NewStrArrayFrom([]string{`UpdatedAt`, `UpdateAt`, `UpdatedTime`, `UpdateTime`}).Contains(fieldCaseCamel) {
			continue
		}
		if garray.NewStrArrayFrom([]string{`CreatedAt`, `CreateAt`, `CreatedTime`, `CreateTime`}).Contains(fieldCaseCamel) {
			continue
		}
		//主键
		if column[`Key`].String() == `PRI` && column[`Extra`].String() == `auto_increment` {
			continue
		}
		//password|passwd
		if garray.NewStrArrayFrom([]string{`password`, `passwd`}).Contains(field) && column[`Type`].String() == `char(32)` {
			viewSaveRule += `
		` + field + `: [
			{ type: 'string', required: computed((): boolean => { return saveForm.data.idArr?.length ? false : true; }), min: 1, max: ` + resultStr[1] + `, trigger: 'blur', message: t('validation.between.string', { min: 1, max: ` + resultStr[1] + ` }) }
		],`
			viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
                    <ElInput v-model="saveForm.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" minlength="1" maxlength="` + resultStr[1] + `" :show-word-limit="true" :clearable="true" :show-password="true" style="max-width: 250px;" />
                    <label v-if="saveForm.data.idArr?.length">
                        <ElAlert :title="t('common.tip.notRequired')" type="info" :show-icon="true" :closable="false" />
                    </label>
                </ElFormItem>`
			continue
		}
		//salt
		if garray.NewStrArrayFrom([]string{`salt`}).Contains(field) && column[`Type`].String() == `char(8)` && tpl.PasswordHandle.IsCoexist {
			continue
		}
		//pid
		if garray.NewStrArrayFrom([]string{`pid`}).Contains(field) && gstr.Pos(column[`Type`].String(), `int`) != -1 {
			viewSaveDataInit += `
		` + field + `: 0,`
			viewSaveRule += `
		` + field + `: [
			{ type: 'integer', min: 0, trigger: 'change', message: t('validation.select') }
		],`
			viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
                    <MyCascader v-model="saveForm.data.` + field + `" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseCamelLower + `/` + tpl.TableNameCaseCamelLower + `/tree', param: { filter: { excIdArr: saveForm.data.idArr } } }" :defaultOptions="[{ id: 0, label: t('common.name.without') }]" :clearable="false" :props="{ checkStrictly: true, emitPath: false }" />
                </ElFormItem>`
			continue
		}
		//level
		if garray.NewStrArrayFrom([]string{`level`}).Contains(field) && gstr.Pos(column[`Type`].String(), `int`) != -1 && tpl.PidHandle.IsCoexist {
			continue
		}
		//idPath|id_path
		if fieldCaseCamel == `IdPath` && gstr.Pos(column[`Type`].String(), `varchar`) != -1 && tpl.PidHandle.IsCoexist {
			continue
		}
		//sort
		//weight
		if garray.NewStrArrayFrom([]string{`sort`, `weight`}).Contains(field) && gstr.Pos(column[`Type`].String(), `int`) != -1 {
			viewSaveDataInit += `
		` + field + `: ` + column[`Default`].String() + `,`
			viewSaveRule += `
		` + field + `: [
			{ type: 'integer', min: 0, max: 100, trigger: 'change', message: t('validation.between.number', { min: 0, max: 100 }) }
		],`
			viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
                    <ElInputNumber v-model="saveForm.data.` + field + `" :precision="0" :min="0" :max="100" :step="1" :step-strictly="true" controls-position="right" :value-on-clear="` + column[`Default`].String() + `" />
                    <label>
                        <ElAlert :title="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.tip.` + field + `')" type="info" :show-icon="true" :closable="false" />
                    </label>
                </ElFormItem>`
			continue
		}
		//id后缀
		if gstr.SubStr(fieldCaseCamel, -2) == `Id` && gstr.Pos(column[`Type`].String(), `int`) != -1 {
			apiUrl := tpl.ModuleDirCaseCamelLower + `/` + gstr.CaseCamelLower(gstr.SubStr(field, 0, -2))
			if tpl.RelTableMap[field].IsExistRelTableDao {
				relTable := tpl.RelTableMap[field]
				apiUrl = relTable.RelDaoDirCaseCamelLower + `/` + relTable.RelTableNameCaseCamelLower
			}
			viewSaveRule += `
		` + field + `: [
			{ type: 'integer', min: 1, trigger: 'change', message: t('validation.select') }
		],`
			viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
                    <MySelect v-model="saveForm.data.` + field + `" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + apiUrl + `/list' }" />
                </ElFormItem>`
			continue
		}
		//name或code后缀
		if (gstr.SubStr(fieldCaseCamel, -4) == `Name` || gstr.SubStr(fieldCaseCamel, -4) == `Code`) && gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
			viewSaveRule += `
		` + field + `: [
			{ type: 'string', min: 1, max: ` + resultStr[1] + `, trigger: 'blur', message: t('validation.between.string', { min: 1, max: ` + resultStr[1] + ` }) },
			{ pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') }
		],`
			if column[`Key`].String() == `UNI` {
				viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
					<ElInput v-model="saveForm.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" minlength="1" maxlength="` + resultStr[1] + `" :show-word-limit="true" :clearable="true" style="max-width: 250px;" />
                    <label>
                        <ElAlert :title="t('common.tip.notDuplicate')" type="info" :show-icon="true" :closable="false" />
                    </label>
				</ElFormItem>`
			} else {
				viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
					<ElInput v-model="saveForm.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" minlength="1" maxlength="` + resultStr[1] + `" :show-word-limit="true" :clearable="true" />
				</ElFormItem>`
			}
			continue
		}
		//mobile,phone后缀
		if (gstr.SubStr(fieldCaseCamel, -5) == `Phone` || gstr.SubStr(fieldCaseCamel, -6) == `Mobile`) && gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
			viewSaveRule += `
		` + field + `: [
			{ type: 'string', min: 1, max: ` + resultStr[1] + `, trigger: 'blur', message: t('validation.between.string', { min: 1, max: ` + resultStr[1] + ` }) },
			{ pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') }
		],`
			viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
					<ElInput v-model="saveForm.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" minlength="1" maxlength="` + resultStr[1] + `" :show-word-limit="true" :clearable="true" />
				</ElFormItem>`
			continue
		}
		//url,link后缀
		if (gstr.SubStr(fieldCaseCamel, -3) == `Url` || gstr.SubStr(fieldCaseCamel, -4) == `Link`) && gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
			viewSaveRule += `
		` + field + `: [
			{ type: 'url', trigger: 'change', message: t('validation.url') },
			{ type: 'string', min: 1, max: ` + resultStr[1] + `, trigger: 'change', message: t('validation.between.string', { min: 1, max: ` + resultStr[1] + ` }) }
		],`
			viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
					<ElInput v-model="saveForm.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" minlength="1" maxlength="` + resultStr[1] + `" :show-word-limit="true" :clearable="true" />
				</ElFormItem>`
			continue
		}
		//avatar
		//icon,cover,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀
		if field == `avatar` || gstr.SubStr(fieldCaseCamel, -4) == `Icon` || gstr.SubStr(fieldCaseCamel, -5) == `Cover` || gstr.SubStr(fieldCaseCamel, -3) == `Img` || gstr.SubStr(fieldCaseCamel, -7) == `ImgList` || gstr.SubStr(fieldCaseCamel, -6) == `ImgArr` || gstr.SubStr(fieldCaseCamel, -5) == `Image` || gstr.SubStr(fieldCaseCamel, -9) == `ImageList` || gstr.SubStr(fieldCaseCamel, -8) == `ImageArr` {
			if column[`Type`].String() == `json` || column[`Type`].String() == `text` {
				viewSaveRule += `
		` + field + `: [
            { type: 'array', trigger: 'change', defaultField: { type: 'url', message: t('validation.url') }, message: t('validation.upload') },
            // { type: 'array', min: 1, trigger: 'change', message: t('validation.min.upload', { min: 1 }) },
            // { type: 'array', max: 10, trigger: 'change', message: t('validation.max.upload', { max: 10 }) }
        ],`
				viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
					<MyUpload v-model="saveForm.data.` + field + `" accept="image/*" :multiple="true" />
				</ElFormItem>`
			} else {
				viewSaveRule += `
		` + field + `: [
			{ type: 'url', trigger: 'change', message: t('validation.upload') },
			{ type: 'string', min: 1, max: ` + resultStr[1] + `, trigger: 'blur', message: t('validation.between.string', { min: 1, max: ` + resultStr[1] + ` }) }
        ],`
				viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
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
            // { type: 'array', min: 1, trigger: 'change', message: t('validation.min.upload', { min: 1 }) },
            // { type: 'array', max: 10, trigger: 'change', message: t('validation.max.upload', { max: 10 }) }
        ],`
				viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
					<MyUpload v-model="saveForm.data.` + field + `" accept="video/*" :isImage="false" :multiple="true" />
				</ElFormItem>`
			} else {
				viewSaveRule += `
		` + field + `: [
			{ type: 'url', trigger: 'change', message: t('validation.upload') },
			{ type: 'string', min: 1, max: ` + resultStr[1] + `, trigger: 'blur', message: t('validation.between.string', { min: 1, max: ` + resultStr[1] + ` }) }
        ],`
				viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
                    <MyUpload v-model="saveForm.data.` + field + `" accept="video/*" :isImage="false" />
                </ElFormItem>`
			}
			continue
		}
		//list,arr等后缀
		if (gstr.SubStr(fieldCaseCamel, -4) == `List` || gstr.SubStr(fieldCaseCamel, -3) == `Arr`) && (column[`Type`].String() == `json` || column[`Type`].String() == `text`) {
			viewSaveDataInit += `
		` + field + `: [],`
			viewSaveRule += `
		` + field + `: [
            // { type: 'array', trigger: 'change', defaultField: { type: 'string', message: '' }, message: '' },
            // { type: 'array', min: 1, trigger: 'change', message: '' },
            // { type: 'array', max: 10, trigger: 'change', message: '' }
        ],`
			viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
					<ElTag v-for="(item, index) in saveForm.data.` + field + `" :type="` + field + `Handle.typeArr[index % 5]" @close="` + field + `Handle.delValue(item)" :key="index" :closable="true" style="margin-right: 10px;">
						{{ item }}
					</ElTag>
					<!-- <ElInputNumber v-if="` + field + `Handle.visible" :ref="(el: any) => { ` + field + `Handle.ref = el }" v-model="` + field + `Handle.value" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" @keyup.enter="` + field + `Handle.addValue" @blur="` + field + `Handle.addValue"  :controls="false" size="small" style="width: 100px;" /> -->
					<ElInput v-if="` + field + `Handle.visible" :ref="(el: any) => { ` + field + `Handle.ref = el }" v-model="` + field + `Handle.value" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" @keyup.enter="` + field + `Handle.addValue" @blur="` + field + `Handle.addValue" size="small" style="width: 100px;" />
					<ElButton v-else type="primary" size="small" @click="` + field + `Handle.visibleChange">
						<AutoiconEpPlus />{{ t('common.add') }}
					</ElButton>
				</ElFormItem>`
			viewFieldHandle += `

const ` + field + `Handle = reactive({
	ref: null as any,
	visible: false,
	value: undefined,
	typeArr: ['', 'success', 'danger', 'info', 'warning'] as any,
	visibleChange: () => {
		` + field + `Handle.visible = true
		nextTick(() => {
			` + field + `Handle.ref?.focus()
		})
	},
	addValue: () => {
		if (` + field + `Handle.value) {
			saveForm.data.` + field + `.push(` + field + `Handle.value)
		}
		` + field + `Handle.visible = false
		` + field + `Handle.value = undefined
	},
	delValue: (item: any) => {
		saveForm.data.` + field + `.splice(saveForm.data.` + field + `.indexOf(item), 1)
	},
})`
			continue
		}
		//IP后缀
		if gstr.SubStr(fieldCaseCamel, -2) == `Ip` && gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
			viewSaveRule += `
		` + field + `: [
			{ type: 'string', min: 1, max: ` + resultStr[1] + `, trigger: 'blur', message: t('validation.between.string', { min: 1, max: ` + resultStr[1] + ` }) }
		],`
			viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
					<ElInput v-model="saveForm.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" minlength="1" maxlength="` + resultStr[1] + `" :show-word-limit="true" :clearable="true" />
				</ElFormItem>`
			continue
		}
		//gender
		//status或type后缀
		if (field == `gender` || gstr.SubStr(fieldCaseCamel, -6) == `Status` || gstr.SubStr(fieldCaseCamel, -4) == `Type`) && gstr.Pos(column[`Type`].String(), `int`) != -1 {
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
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">`
			//超过5个状态用select组件，小于5个用radio组件
			if len(statusArr) > 5 {
				viewSaveField += `
					<ElSelectV2 v-model="saveForm.data.` + field + `" :options="tm('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.status.` + field + `')" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :clearable="true" />`
			} else {
				viewSaveDataInit += `
		` + field + `: ` + statusList[0][0] + `,`
				viewSaveField += `
					<ElRadioGroup v-model="saveForm.data.` + field + `">
                        <ElRadio v-for="(item, index) in (tm('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.status.` + field + `') as any)" :key="index" :label="item.value">
                            {{ item.label }}
                        </ElRadio>
                    </ElRadioGroup>`
			}
			viewSaveField += `
				</ElFormItem>`
			continue
		}
		//remark,desc,msg,message,intro,content后缀
		if gstr.SubStr(fieldCaseCamel, -6) == `Remark` || gstr.SubStr(fieldCaseCamel, -4) == `Desc` || gstr.SubStr(fieldCaseCamel, -3) == `Msg` || gstr.SubStr(fieldCaseCamel, -7) == `Message` || gstr.SubStr(fieldCaseCamel, -5) == `Intro` || gstr.SubStr(fieldCaseCamel, -7) == `Content` {
			if gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
				viewSaveRule += `
		` + field + `: [
			{ type: 'string', max: ` + resultStr[1] + `, trigger: 'blur', message: t('validation.max.string', { max: ` + resultStr[1] + ` }) },
		],`
				viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
					<ElInput v-model="saveForm.data.` + field + `" type="textarea" :autosize="{ minRows: 3 }" />
				</ElFormItem>`
				continue
			} else if column[`Type`].String() == `text` {
				viewSaveRule += `
		` + field + `: [],`
				viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
					<MyEditor v-model="saveForm.data.` + field + `" />
				</ElFormItem>`
				continue
			}
		}
		//is_前缀
		if gstr.SubStr(fieldCaseSnake, 0, 3) == `is_` && gstr.Pos(column[`Type`].String(), `int`) != -1 {
			viewSaveRule += `
		` + field + `: [
            { type: 'enum', enum: [0, 1], trigger: 'change', message: t('validation.select') }
        ],`
			viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
                    <ElSwitch v-model="saveForm.data.` + field + `" :active-value="1" :inactive-value="0" :inline-prompt="true" :active-text="t('common.yes')" :inactive-text="t('common.no')" style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success);" />
                </ElFormItem>`
			continue
		}
		//start_前缀
		if gstr.SubStr(fieldCaseSnake, 0, 6) == `start_` {
			if gstr.Pos(column[`Type`].String(), `timestamp`) != -1 || gstr.Pos(column[`Type`].String(), `datetime`) != -1 {
				if column[`Null`].String() == `NO` && column[`Default`].String() == `` {
					viewSaveRule += `
		` + field + `: [
			{ type: 'string', required: true, trigger: 'change', message: t('validation.select') }
		],`
				} else {
					viewSaveRule += `
		` + field + `: [
			{ type: 'string', trigger: 'change', message: '' }
		],`
				}
				viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
					<ElDatePicker v-model="saveForm.data.` + field + `" type="datetime" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :default-time="new Date(2000, 0, 1, 0, 0, 0)" format="YYYY-MM-DD HH:mm:ss" value-format="YYYY-MM-DD HH:mm:ss" />
				</ElFormItem>`
				continue
			}
			/* if gstr.Pos(column[`Type`].String(), `date`) != -1 {
				continue
			} */
		}
		//end_前缀
		if gstr.SubStr(fieldCaseSnake, 0, 4) == `end_` {
			if gstr.Pos(column[`Type`].String(), `timestamp`) != -1 || gstr.Pos(column[`Type`].String(), `datetime`) != -1 {
				if column[`Null`].String() == `NO` && column[`Default`].String() == `` {
					viewSaveRule += `
		` + field + `: [
			{ type: 'string', required: true, trigger: 'change', message: t('validation.select') }
		],`
				} else {
					viewSaveRule += `
		` + field + `: [
			{ type: 'string', trigger: 'change', message: '' }
		],`
				}
				viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
					<ElDatePicker v-model="saveForm.data.` + field + `" type="datetime" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :default-time="new Date(2000, 0, 1, 23, 59, 59)" format="YYYY-MM-DD HH:mm:ss" value-format="YYYY-MM-DD HH:mm:ss" />
				</ElFormItem>`
				continue
			}
			/* if gstr.Pos(column[`Type`].String(), `date`) != -1 {
				continue
			} */
		}
		//int类型
		if gstr.Pos(column[`Type`].String(), `int`) != -1 {
			if gstr.Pos(column[`Type`].String(), `unsigned`) != -1 {
				viewSaveRule += `
		` + field + `: [
			{ type: 'integer', min: 0, trigger: 'change', message: t('validation.min.number', { min: 0 }) }
		],`
				viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
					<ElInputNumber v-model="saveForm.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :min="0" :controls="false"/>
				</ElFormItem>`
			} else {
				viewSaveRule += `
		` + field + `: [
			{ type: 'integer', trigger: 'change', message: '' }
		],`
				viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
					<ElInputNumber v-model="saveForm.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :controls="false"/>
				</ElFormItem>`
			}
			continue
		}
		//float类型
		if gstr.Pos(column[`Type`].String(), `decimal`) != -1 || gstr.Pos(column[`Type`].String(), `double`) != -1 || gstr.Pos(column[`Type`].String(), `float`) != -1 {
			if gstr.Pos(column[`Type`].String(), `unsigned`) != -1 {
				viewSaveRule += `
		` + field + `: [
			{ type: 'number'/* 'float' */, min: 0, trigger: 'change', message: t('validation.min.number', { min: 0 }) }	// 类型float值为0时验证不能通过
		],`
				viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
					<ElInputNumber v-model="saveForm.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :min="0" :precision="` + resultFloat[2] + `" :controls="false"/>
				</ElFormItem>`
			} else {
				viewSaveRule += `
		` + field + `: [
			{ type: 'number'/* 'float' */, trigger: 'change', message: '' }	// 类型float值为0时验证不能通过
		],`
				viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
					<ElInputNumber v-model="saveForm.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :precision="` + resultFloat[2] + `" :controls="false"/>
				</ElFormItem>`
			}
			continue
		}
		//varchar类型
		if gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
			viewSaveRule += `
		` + field + `: [
			{ type: 'string', min: 1, max: ` + resultStr[1] + `, trigger: 'blur', message: t('validation.between.string', { min: 1, max: ` + resultStr[1] + ` }) },
		],`
			if column[`Key`].String() == `UNI` {
				viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
					<ElInput v-model="saveForm.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" minlength="1" maxlength="` + resultStr[1] + `" :show-word-limit="true" :clearable="true" style="max-width: 250px;" />
                    <label>
                        <ElAlert :title="t('common.tip.notDuplicate')" type="info" :show-icon="true" :closable="false" />
                    </label>
				</ElFormItem>`
			} else {
				viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
					<ElInput v-model="saveForm.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" minlength="1" maxlength="` + resultStr[1] + `" :show-word-limit="true" :clearable="true" />
				</ElFormItem>`
			}
			continue
		}
		//char类型
		if gstr.Pos(column[`Type`].String(), `char`) != -1 {
			viewSaveRule += `
		` + field + `: [
			{ type: 'string', min: ` + resultStr[1] + `, max: ` + resultStr[1] + `, trigger: 'blur', message: t('validation.between.string', { min: ` + resultStr[1] + `, max: ` + resultStr[1] + ` }) },
		],`
			if column[`Key`].String() == `UNI` {
				viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
					<ElInput v-model="saveForm.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" minlength="` + resultStr[1] + `" maxlength="` + resultStr[1] + `" :show-word-limit="true" :clearable="true" style="max-width: 250px;" />
                    <label>
                        <ElAlert :title="t('common.tip.notDuplicate')" type="info" :show-icon="true" :closable="false" />
                    </label>
				</ElFormItem>`
			} else {
				viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
					<ElInput v-model="saveForm.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" minlength="` + resultStr[1] + `" maxlength="` + resultStr[1] + `" :show-word-limit="true" :clearable="true" />
				</ElFormItem>`
			}
			continue
		}
		//timestamp或datetime类型
		if gstr.Pos(column[`Type`].String(), `timestamp`) != -1 || gstr.Pos(column[`Type`].String(), `datetime`) != -1 {
			if column[`Null`].String() == `NO` && column[`Default`].String() == `` {
				viewSaveRule += `
		` + field + `: [
			{ type: 'string', required: true, trigger: 'change', message: t('validation.select') }
		],`
			} else {
				viewSaveRule += `
		` + field + `: [
			{ type: 'string', trigger: 'change', message: '' }
		],`
			}
			viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
					<ElDatePicker v-model="saveForm.data.` + field + `" type="datetime" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" format="YYYY-MM-DD HH:mm:ss" value-format="YYYY-MM-DD HH:mm:ss" />
				</ElFormItem>`
			continue
		}
		//date类型
		if gstr.Pos(column[`Type`].String(), `date`) != -1 {
			if column[`Null`].String() == `NO` && column[`Default`].String() == `` {
				viewSaveRule += `
		` + field + `: [
			{ type: 'string', required: true, trigger: 'change', message: t('validation.select') }
		],`
			} else {
				viewSaveRule += `
		` + field + `: [
			{ type: 'string', trigger: 'change', message: '' }
		],`
			}
			viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
					<ElDatePicker v-model="saveForm.data.` + field + `" type="date" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" format="YYYY-MM-DD" value-format="YYYY-MM-DD" />
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
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
                    <ElAlert :title="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.tip.` + field + `')" type="info" :show-icon="true" :closable="false" />
					<ElInput v-model="saveForm.data.` + field + `" type="textarea" :autosize="{ minRows: 3 }" />
				</ElFormItem>`
			continue
		}
		//text类型
		if gstr.Pos(column[`Type`].String(), `text`) != -1 {
			viewSaveRule += `
		` + field + `: [],`
			viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
					<MyEditor v-model="saveForm.data.` + field + `" />
				</ElFormItem>`
			continue
		}
		//默认处理
		viewSaveRule += `
		` + field + `: [],`
		viewSaveField += `
				<ElFormItem :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
					<ElInput v-model="saveForm.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :show-word-limit="true" :clearable="true" />
				</ElFormItem>`
	}

	tplView := `<script setup lang="ts">`
	if tpl.PasswordHandle.PasswordField != `` {
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
	data: {` + viewSaveDataInit + `
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
	if tpl.PasswordHandle.PasswordField != `` {
		tplView += `
            param.` + tpl.PasswordHandle.PasswordField + ` ? param.` + tpl.PasswordHandle.PasswordField + ` = md5(param.` + tpl.PasswordHandle.PasswordField + `) : delete param.` + tpl.PasswordHandle.PasswordField
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
})` + viewFieldHandle + `
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
	saveFile := gfile.SelfDir() + `/../view/` + option.SceneCode + `/src/i18n/language/zh-cn/` + tpl.ModuleDirCaseCamelLower + `/` + tpl.TableNameCaseCamelLower + `.ts`
	if !option.IsCover && gfile.IsFile(saveFile) {
		return
	}

	viewI18nName := ``
	viewI18nStatus := ``
	viewI18nTip := ``
	for _, column := range tpl.TableColumnList {
		field := column[`Field`].String()
		fieldCaseCamel := gstr.CaseCamel(field)
		tmp, _ := gregex.MatchString(`[^\n\r\.。:：\(（]*`, column[`Comment`].String())
		fieldName := gstr.Trim(tmp[0])
		comment := gstr.Trim(gstr.ReplaceByArray(column[`Comment`].String(), g.SliceStr{
			"\n", ` `,
			"\r", ` `,
			`"`, `\"`,
		}))
		tip := gstr.Replace(comment, fieldName, ``, 1)
		tmp, _ = gregex.MatchString(`\n\r\.。:：\(（`, column[`Comment`].String())
		if len(tmp) > 0 {
			gstr.TrimLeft(tip, tmp[0])
		}
		for _, v := range []string{"\n", "\r", `.`, `。`, `:`, `：`, `(`, `（`, `)`, `）`, ` `, `,`, `，`, `;`, `；`} {
			tip = gstr.Trim(tip, v)
		}
		tip = gstr.ReplaceByArray(tip, g.SliceStr{
			`\"`, `"`,
			`}`, `' + "{'}'}" + '`,
			`{"`, `' + "{'{'}" + '"`,
		})

		if garray.NewStrArrayFrom([]string{`DeletedAt`, `DeleteAt`, `DeletedTime`, `DeleteTime`}).Contains(fieldCaseCamel) {
			continue
		}
		if garray.NewStrArrayFrom([]string{`UpdatedAt`, `UpdateAt`, `UpdatedTime`, `UpdateTime`}).Contains(fieldCaseCamel) {
			continue
		}
		if garray.NewStrArrayFrom([]string{`CreatedAt`, `CreateAt`, `CreatedTime`, `CreateTime`}).Contains(fieldCaseCamel) {
			continue
		}
		//主键
		if column[`Key`].String() == `PRI` && column[`Extra`].String() == `auto_increment` {
			continue
		}
		//password|passwd
		/* if garray.NewStrArrayFrom([]string{`password`, `passwd`}).Contains(field) && column[`Type`].String() == `char(32)` {
			continue
		} */
		//salt
		if garray.NewStrArrayFrom([]string{`salt`}).Contains(field) && column[`Type`].String() == `char(8)` && tpl.PasswordHandle.IsCoexist {
			continue
		}
		//pid
		if garray.NewStrArrayFrom([]string{`pid`}).Contains(field) && gstr.Pos(column[`Type`].String(), `int`) != -1 {
			viewI18nName += `
		` + field + `: '父级',`
			continue
		}
		/* //level
		if garray.NewStrArrayFrom([]string{`level`}).Contains(field) && gstr.Pos(column[`Type`].String(), `int`) != -1 && tpl.PidHandle.IsCoexist {
			continue
		}
		//idPath|id_path
		if fieldCaseCamel == `IdPath` && gstr.Pos(column[`Type`].String(), `varchar`) != -1 && tpl.PidHandle.IsCoexist {
			continue
		} */
		//sort
		//weight
		if garray.NewStrArrayFrom([]string{`sort`, `weight`}).Contains(field) && gstr.Pos(column[`Type`].String(), `int`) != -1 {
			viewI18nName += `
		` + field + `: '` + fieldName + `',`
			viewI18nTip += `
		` + field + `: '` + tip + `',`
			continue
		}
		//id后缀
		if gstr.SubStr(fieldCaseCamel, -2) == `Id` && gstr.Pos(column[`Type`].String(), `int`) != -1 {
			if tpl.RelTableMap[field].IsExistRelTableDao {
				relTable := tpl.RelTableMap[field]
				fieldName = relTable.RelNameFieldName
			}
			viewI18nName += `
		` + field + `: '` + fieldName + `',`
			continue
		}
		//icon,cover,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀
		//video,video_list,videoList,video_arr,videoArr等后缀
		if field == `avatar` || gstr.SubStr(fieldCaseCamel, -4) == `Icon` || gstr.SubStr(fieldCaseCamel, -5) == `Cover` || gstr.SubStr(fieldCaseCamel, -3) == `Img` || gstr.SubStr(fieldCaseCamel, -7) == `ImgList` || gstr.SubStr(fieldCaseCamel, -6) == `ImgArr` || gstr.SubStr(fieldCaseCamel, -5) == `Image` || gstr.SubStr(fieldCaseCamel, -9) == `ImageList` || gstr.SubStr(fieldCaseCamel, -8) == `ImageArr` || gstr.SubStr(fieldCaseCamel, -5) == `Video` || gstr.SubStr(fieldCaseCamel, -9) == `VideoList` || gstr.SubStr(fieldCaseCamel, -8) == `VideoArr` {
			viewI18nName += `
		` + field + `: '` + fieldName + `',`
			continue
		}
		//list,arr等后缀
		if (gstr.SubStr(fieldCaseCamel, -4) == `List` || gstr.SubStr(fieldCaseCamel, -3) == `Arr`) && (column[`Type`].String() == `json` || column[`Type`].String() == `text`) {
			viewI18nName += `
		` + field + `: '` + fieldName + `',`
			continue
		}
		//gender
		//status或type后缀
		if (field == `gender` || gstr.SubStr(fieldCaseCamel, -6) == `Status` || gstr.SubStr(fieldCaseCamel, -4) == `Type`) && gstr.Pos(column[`Type`].String(), `int`) != -1 {
			viewI18nName += `
		` + field + `: '` + fieldName + `',`
			statusList := MyGenStatusList(comment)
			viewI18nStatus += `
		` + field + `: [`
			for _, status := range statusList {
				viewI18nStatus += `
			{ value: ` + status[0] + `, label: '` + status[1] + `' },`
			}
			viewI18nStatus += `
		],`
			continue
		}
		//json类型
		if column[`Type`].String() == `json` {
			viewI18nName += `
		` + field + `: '` + fieldName + `',`
			viewI18nTip += `
		` + field + `: '` + tip + `',`
			continue
		}
		viewI18nName += `
		` + field + `: '` + fieldName + `',`
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
	saveFile := gfile.SelfDir() + `/../view/` + option.SceneCode + `/src/router/index.ts`

	tplViewRouter := gfile.GetContents(saveFile)

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

	if gstr.Pos(tplViewRouter, `'`+path+`'`) == -1 { //路由不存在时新增
		tplViewRouter = gstr.Replace(tplViewRouter, `/*--------前端路由自动代码生成锚点（不允许修改和删除，否则将不能自动生成路由）--------*/`, replaceStr+`
            /*--------前端路由自动代码生成锚点（不允许修改和删除，否则将不能自动生成路由）--------*/`)
	} else { //路由已存在则替换
		tplViewRouter, _ = gregex.ReplaceString(`\{
                path: '`+path+`',[\s\S]*'`+path+`' \}
            \},`, replaceStr, tplViewRouter)
	}
	gfile.PutContents(saveFile, tplViewRouter)

	MyGenMenu(ctx, tpl.SceneId, path, option.CommonName, tpl.TableNameCaseCamel) // 数据库权限菜单处理
}
