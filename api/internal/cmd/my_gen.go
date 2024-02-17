package cmd

import (
	daoAuth "api/internal/dao/auth"
	"api/internal/service"
	"api/internal/utils"
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
后台常用生成示例：./main myGen -sceneCode=platform -dbGroup=default -dbTable=auth_test -removePrefix=auth_ -moduleDir=auth -commonName=权限管理/测试 -isList=1 -isCount=1 -isInfo=1 -isCreate=1 -isUpdate=1 -isDelete=1 -isApi=1 -isAuthAction=1 -isView=1 -isCover=0
APP常用生成示例：./main myGen -sceneCode=app -dbGroup=xxxx -dbTable=user -removePrefix= -moduleDir=xxxx/user -commonName=用户 -isList=1 -isCount=0 -isInfo=1 -isCreate=0 -isUpdate=0 -isDelete=0 -isApi=1 -isAuthAction=0 -isView=0 -isCover=0

强烈建议搭配Git使用
主键必须在第一个字段。否则需要在dao层重写PrimaryKey方法返回主键字段
表内尽量根据表名设置xxxxId主键和xxxxName名称两个字段（作用1：常用于前端部分组件，如MySelect.vue组件；作用2：当其它表存在与该表主键同名的关联字段时，会自动生成联表查询代码）
每个字段都必须有注释。以下符号[\n\r.。:：(（]之前的部分或整个注释，将作为字段名称使用
表字段按以下规则命名时，会做特殊处理，其它情况根据字段类型做默认处理

	固定命名：
		父级		命名：pid；      		类型：int等类型；		注意：pid,level,idPath|id_path同时存在时，有特殊处理
		层级		命名：level；          	类型：int等类型；		注意：pid,level,idPath|id_path同时存在时，(才)有特殊处理
		层级路径	命名：idPath|id_path；	类型：varchar或text；	注意：pid,level,idPath|id_path同时存在时，(才)有特殊处理
		排序		命名：sort；			类型：int等类型；		注意：pid,level,idPath|id_path|sort同时存在时，(才)有特殊处理

	常用命名(字段含[_of_]时，会忽略[_of_]及其之后的部分)：
		密码		命名：password,passwd后缀；		类型：char(32)；
		加密盐 		命名：salt后缀；     			类型：char；	注意：password,salt同时存在时，有特殊处理
		名称		命名：name后缀；				类型：varchar；
		标识		命名：code后缀；				类型：varchar；
		手机		命名：mobile,phone后缀；		类型：varchar；
		链接		命名：url,link后缀；			类型：varchar；
		IP			命名：IP后缀；					类型：varchar；
		关联ID		命名：id后缀；					类型：int等类型；
		排序|权重	命名：sort,weight等后缀；		类型：int等类型；
		是否		命名：is_前缀；					类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（停用：0否 1是）
		状态|类型	命名：status,type,method,pos,position,gender等后缀；类型：int等类型或varchar或char；注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
		开始时间	命名：start_前缀；				类型：timestamp或datetime或date；
		结束时间	命名：end_前缀；				类型：timestamp或datetime或date；
		(富)文本	命名：remark,desc,msg,message,intro,content后缀；类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器
		图片		命名：icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；类型：单图片varchar，多图片json或text
		视频		命名：video,video_list,videoList,video_arr,videoArr等后缀；类型：单视频varchar，多视频json或text
		数组		命名：list,arr等后缀；类型：json或text；
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
	SceneId                        uint       //场景ID
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
	// 以下字段用于对某些表字段做特殊处理
	LabelHandle struct { //dao层label对应的字段(常用于前端组件)
		LabelField string //是否同时存在
		IsCoexist  bool   //当LabelField=phone或account时，是否同时存在phone和account两个字段
	}
	RelTableMap       map[string]RelTableItem       //id后缀字段，能确定关联表时，会自动生成联表查询代码
	PasswordHandleMap map[string]PasswordHandleItem //password|passwd,salt同时存在时，需特殊处理
	PidHandle         struct {                      //pid,level,idPath|id_path同时存在时，需特殊处理
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
	RelTableField              string //关联表字段
	RelTableFieldName          string //关联表字段名称
	IsRedundRelNameField       bool   //当前表是否冗余关联表字段
	RelSuffix                  string //关联表字段后缀（原始，大驼峰或蛇形）。字段含[_of_]时，_of_及之后的部分。示例：userIdOfSend对应OfSend；user_id_of_send对应_of_send
	RelSuffixCaseCamel         string //关联表字段后缀（大驼峰）。字段含[_of_]时，_of_及其之后的部分。示例：userIdOfSend和user_id_of_send都对应OfSend
	RelSuffixCaseSnake         string //关联表字段后缀（蛇形）。字段含[_of_]时，_of_及其之后的部分。示例：userIdOfSend和user_id_of_send都对应_of_send
}

type PasswordHandleItem struct {
	IsCoexist      bool   //是否同时存在
	PasswordField  string //密码字段
	PasswordLength string //密码字段长度
	SaltField      string //加密盐字段
	SaltLength     string //加密盐字段长度
}

func MyGenFunc(ctx context.Context, parser *gcmd.Parser) (err error) {
	option := MyGenOptionHandle(ctx, parser)
	tpl := MyGenTplHandle(ctx, option)

	MyGenTplDao(ctx, option, tpl)   // dao层存在时，增加或修改部分字段的解析代码
	MyGenTplLogic(ctx, option, tpl) // logic模板生成（文件不存在时增删改查全部生成，已存在不处理不覆盖）
	// service生成
	fmt.Println(`service生成 开始`)
	command := exec.Command(`gf`, `gen`, `service`)
	stdout, _ := command.StdoutPipe()
	command.Start()
	buf := make([]byte, 1024)
	for {
		n, err := stdout.Read(buf)
		if err != nil {
			break
		}
		if n > 0 {
			fmt.Print(string(buf[:n]))
		}
	}
	command.Wait()
	fmt.Println(`service生成 结束`)

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
		// 前端代码格式化
		fmt.Println(`前端代码格式化 开始`)
		command := exec.Command(`npm`, `run`, `format`)
		command.Dir = gfile.SelfDir() + `/../view/` + option.SceneCode
		stdout, _ := command.StdoutPipe()
		command.Start()
		buf := make([]byte, 1024)
		for {
			n, err := stdout.Read(buf)
			if err != nil {
				break
			}
			if n > 0 {
				fmt.Print(string(buf[:n]))
			}
		}
		command.Wait()
		fmt.Println(`前端代码格式化 结束`)
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
		fmt.Println(`请重新选择生成哪些接口，不能全是no！`)
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
		SceneId:                 sceneInfo[daoAuth.Scene.Columns().SceneId].Uint(),
		TableNameCaseCamel:      gstr.CaseCamel(tableName),
		TableNameCaseCamelLower: gstr.CaseCamelLower(tableName),
		TableNameCaseSnake:      gstr.CaseSnakeFirstUpper(tableName),
		RelTableMap:             map[string]RelTableItem{},
		PasswordHandleMap:       map[string]PasswordHandleItem{},
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
		fieldCaseSnake := gstr.CaseSnakeFirstUpper(field)
		fieldCaseSnakeOfRemove := gstr.Split(fieldCaseSnake, `_of_`)[0]
		fieldCaseCamelOfRemove := gstr.CaseCamel(fieldCaseSnakeOfRemove)
		fieldSplitArr := gstr.Split(fieldCaseSnakeOfRemove, `_`)
		// fieldPrefix := fieldSplitArr[0]
		fieldSuffix := fieldSplitArr[len(fieldSplitArr)-1]
		resultStr, _ := gregex.MatchString(`.*\((\d*)\)`, column[`Type`].String())
		tmp, _ := gregex.MatchString(`[^\n\r\.。:：\(（]*`, column[`Comment`].String())
		fieldName := gstr.Trim(tmp[0])
		fieldArr[index] = field
		fieldCaseCamelArr[index] = fieldCaseCamel

		if garray.NewStrArrayFrom([]string{`DeletedAt`, `DeleteAt`, `DeletedTime`, `DeleteTime`}).Contains(fieldCaseCamel) {
			tpl.DeletedField = field
		} else if garray.NewStrArrayFrom([]string{`UpdatedAt`, `UpdateAt`, `UpdatedTime`, `UpdateTime`}).Contains(fieldCaseCamel) {
			tpl.UpdatedField = field
		} else if garray.NewStrArrayFrom([]string{`CreatedAt`, `CreateAt`, `CreatedTime`, `CreateTime`}).Contains(fieldCaseCamel) {
			tpl.CreatedField = field
		} else if column[`Key`].String() == `PRI` && column[`Extra`].String() == `auto_increment` { //主键
			tpl.PrimaryKey = field
		} else if fieldCaseCamel == `IdPath` && (gstr.Pos(column[`Type`].String(), `varchar`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1) { //idPath|id_path
			tpl.PidHandle.IdPathField = field
		} else if gstr.Pos(column[`Type`].String(), `varchar`) != -1 { //varchar类型
		} else if gstr.Pos(column[`Type`].String(), `char`) != -1 { //char类型
			if garray.NewStrArrayFrom([]string{`password`, `passwd`}).Contains(fieldSuffix) && column[`Type`].String() == `char(32)` { //password,passwd后缀
				passwordHandleMapKey := MyGenPasswordHandleMapKey(field)
				passwordHandleItem, ok := tpl.PasswordHandleMap[passwordHandleMapKey]
				if ok {
					passwordHandleItem.PasswordField = field
					passwordHandleItem.PasswordLength = resultStr[1]
				} else {
					passwordHandleItem = PasswordHandleItem{
						PasswordField:  field,
						PasswordLength: resultStr[1],
					}
				}
				tpl.PasswordHandleMap[passwordHandleMapKey] = passwordHandleItem
			} else if garray.NewStrArrayFrom([]string{`salt`}).Contains(fieldSuffix) { //salt后缀
				passwordHandleMapKey := MyGenPasswordHandleMapKey(field)
				passwordHandleItem, ok := tpl.PasswordHandleMap[passwordHandleMapKey]
				if ok {
					passwordHandleItem.SaltField = field
					passwordHandleItem.SaltLength = resultStr[1]
				} else {
					passwordHandleItem = PasswordHandleItem{
						SaltField:  field,
						SaltLength: resultStr[1],
					}
				}
				tpl.PasswordHandleMap[passwordHandleMapKey] = passwordHandleItem
			}
		} else if gstr.Pos(column[`Type`].String(), `int`) != -1 && gstr.Pos(column[`Type`].String(), `point`) == -1 { //int等类型
			if field == `pid` { //pid
				tpl.PidHandle.PidField = field
			} else if field == `level` { //level
				tpl.PidHandle.LevelField = field
			} else if garray.NewStrArrayFrom([]string{`sort`, `weight`}).Contains(fieldSuffix) { //sort,weight等后缀
				if field == `sort` { //sort
					tpl.PidHandle.SortField = field
				}
			} else if garray.NewStrArrayFrom([]string{`id`}).Contains(fieldSuffix) { //id后缀
				relTableNameCaseCamel := gstr.SubStr(fieldCaseCamelOfRemove, 0, -2)
				relTableItem := RelTableItem{
					IsExistRelTableDao:         false,
					RelDaoDir:                  ``,
					RelDaoDirCaseCamel:         ``,
					RelDaoDirCaseCamelLower:    ``,
					IsSameDir:                  false,
					RelTableNameCaseCamel:      relTableNameCaseCamel,
					RelTableNameCaseCamelLower: gstr.CaseCamelLower(relTableNameCaseCamel),
					RelTableNameCaseSnake:      gstr.CaseSnakeFirstUpper(relTableNameCaseCamel),
					RelTableField:              ``,
					RelTableFieldName:          fieldName,
					IsRedundRelNameField:       false,
					RelSuffix:                  ``,
					RelSuffixCaseCamel:         ``,
					RelSuffixCaseSnake:         ``,
				}
				fieldCaseSnakeArr := gstr.Split(fieldCaseSnake, `_of_`)
				if len(fieldCaseSnakeArr) > 1 {
					relTableItem.RelSuffixCaseSnake = `_of_` + gstr.Join(fieldCaseSnakeArr[1:], `_of_`)
					relTableItem.RelSuffixCaseCamel = gstr.CaseCamel(relTableItem.RelSuffixCaseSnake)
					relTableItem.RelSuffix = relTableItem.RelSuffixCaseCamel //默认：小驼峰
				}
				relTableItem.RelTableField = relTableItem.RelTableNameCaseCamelLower + `Name` //默认：小驼峰
				if gstr.CaseCamelLower(field) != field {                                      //判断字段是不是蛇形
					relTableItem.RelTableField = relTableItem.RelTableNameCaseSnake + `_name`
					if len(fieldCaseSnakeArr) > 1 {
						relTableItem.RelSuffix = relTableItem.RelSuffixCaseSnake
					}
				}
				if gstr.ToUpper(gstr.SubStr(relTableItem.RelTableFieldName, -2)) == `ID` {
					relTableItem.RelTableFieldName = gstr.SubStr(relTableItem.RelTableFieldName, 0, -2)
				}

				selfDir := gfile.SelfDir()
				fileArr, _ := gfile.ScanDirFile(selfDir+`/internal/dao/`, relTableItem.RelTableNameCaseSnake+`\.go`, true)
				relDaoDirList := []string{}
				for _, v := range fileArr {
					if gstr.Count(v, `/internal/`) == 1 {
						if v == selfDir+`/internal/dao/`+tpl.ModuleDirCaseCamelLower+`/`+tpl.TableNameCaseSnake+`.go` { //自身跳过
							continue
						}
						relDaoDirTmp := gstr.Replace(v, selfDir+`/internal/dao/`, ``, 1)
						relDaoDirTmp = gstr.Replace(relDaoDirTmp, `/`+relTableItem.RelTableNameCaseSnake+`.go`, ``, 1)
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
			}
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
			tpl.LabelHandle.LabelField = fieldArr[index]
			break
		}
	}
	if tpl.LabelHandle.LabelField == `phone` || tpl.LabelHandle.LabelField == `account` {
		if gset.NewStrSetFrom([]string{`Phone`, `Account`}).Intersect(gset.NewStrSetFrom(fieldCaseCamelArr)).Size() == 2 {
			tpl.LabelHandle.IsCoexist = true
		}
	}

	for k, v := range tpl.RelTableMap {
		if garray.NewStrArrayFrom(fieldArr).Contains(v.RelTableField + v.RelSuffix) {
			v.IsRedundRelNameField = true
			tpl.RelTableMap[k] = v
		}
	}

	for k, v := range tpl.PasswordHandleMap {
		if v.PasswordField != `` && v.SaltField != `` {
			v.IsCoexist = true
			tpl.PasswordHandleMap[k] = v
		}
	}

	if tpl.PidHandle.PidField != `` && tpl.PidHandle.LevelField != `` && tpl.PidHandle.IdPathField != `` {
		tpl.PidHandle.IsCoexist = true
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

	if tpl.LabelHandle.LabelField != `` {
		daoParseFieldTmp := `
			case ` + "`label`" + `:
				m = m.Fields(daoHandler.DbTable + ` + "`.`" + ` + daoThis.Columns().` + gstr.CaseCamel(tpl.LabelHandle.LabelField) + ` + ` + "` AS `" + ` + v)`
		daoParseFilterTmp := `
			case ` + "`label`" + `:
				m = m.WhereLike(daoHandler.DbTable+` + "`.`" + `+daoThis.Columns().` + gstr.CaseCamel(tpl.LabelHandle.LabelField) + `, ` + "`%`" + `+gconv.String(v)+` + "`%`" + `)`
		if tpl.LabelHandle.IsCoexist {
			daoParseFieldTmp = `
			case ` + "`label`" + `:
				m = m.Fields(` + "`IFNULL(` + daoHandler.DbTable + `.` + daoThis.Columns().Account + `, ` + daoHandler.DbTable + `.` + daoThis.Columns().Phone + `) AS ` + v)"
			daoParseFilterTmp = `
			case ` + "`label`" + `:
				m = m.Where(` + "m.Builder().WhereLike(daoHandler.DbTable+`.`+daoThis.Columns().Account, `%`+gconv.String(v)+`%`).WhereOrLike(daoHandler.DbTable+`.`+daoThis.Columns().Phone, `%`+gconv.String(v)+`%`))"
		}
		if gstr.Pos(tplDao, daoParseFieldTmp) == -1 {
			daoParseField += daoParseFieldTmp
		}
		if gstr.Pos(tplDao, daoParseFilterTmp) == -1 {
			daoParseFilter += daoParseFilterTmp
		}
	}

	for _, column := range tpl.TableColumnList {
		field := column[`Field`].String()
		fieldCaseCamel := gstr.CaseCamel(field)
		fieldCaseSnake := gstr.CaseSnakeFirstUpper(field)
		fieldCaseSnakeOfRemove := gstr.Split(fieldCaseSnake, `_of_`)[0]
		// fieldCaseCamelOfRemove := gstr.CaseCamel(fieldCaseSnakeOfRemove)
		fieldSplitArr := gstr.Split(fieldCaseSnakeOfRemove, `_`)
		fieldPrefix := fieldSplitArr[0]
		fieldSuffix := fieldSplitArr[len(fieldSplitArr)-1]

		if garray.NewStrArrayFrom([]string{`DeletedAt`, `DeleteAt`, `DeletedTime`, `DeleteTime`}).Contains(fieldCaseCamel) {
		} else if garray.NewStrArrayFrom([]string{`UpdatedAt`, `UpdateAt`, `UpdatedTime`, `UpdateTime`}).Contains(fieldCaseCamel) {
		} else if garray.NewStrArrayFrom([]string{`CreatedAt`, `CreateAt`, `CreatedTime`, `CreateTime`}).Contains(fieldCaseCamel) {
			daoParseFilterTmp := `
			case ` + "`timeRangeStart`" + `:
				m = m.WhereGTE(daoHandler.DbTable+` + "`.`" + `+daoThis.Columns().` + fieldCaseCamel + `, v)
			case ` + "`timeRangeEnd`" + `:
				m = m.WhereLTE(daoHandler.DbTable+` + "`.`" + `+daoThis.Columns().` + fieldCaseCamel + `, v)`
			if gstr.Pos(tplDao, daoParseFilterTmp) == -1 {
				daoParseFilter += daoParseFilterTmp
			}
		} else if column[`Key`].String() == `PRI` && column[`Extra`].String() == `auto_increment` { //主键
		} else if fieldCaseCamel == `IdPath` && (gstr.Pos(column[`Type`].String(), `varchar`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1) && tpl.PidHandle.IsCoexist { //idPath|id_path
		} else if gstr.Pos(column[`Type`].String(), `varchar`) != -1 { //varchar类型
			if garray.NewStrArrayFrom([]string{`name`}).Contains(fieldSuffix) { //name后缀
				daoParseFilterTmp := `
			case daoThis.Columns().` + fieldCaseCamel + `:
				m = m.WhereLike(daoHandler.DbTable+` + "`.`" + `+k, ` + "`%`" + `+gconv.String(v)+` + "`%`" + `)`
				if gstr.Pos(tplDao, daoParseFilterTmp) == -1 {
					daoParseFilter += daoParseFilterTmp
				}
			}

			if column[`Key`].String() == `UNI` && column[`Null`].String() == `YES` {
				daoParseInsertTmp := `
			case daoThis.Columns().` + fieldCaseCamel + `:
				insertData[k] = v
				if gconv.String(v) == ` + "``" + ` {
					insertData[k] = nil
				}`
				if gstr.Pos(tplDao, daoParseInsertTmp) == -1 {
					daoParseInsert += daoParseInsertTmp
				}
				daoParseUpdateTmp := `
			case daoThis.Columns().` + fieldCaseCamel + `:
				updateData[daoHandler.DbTable+` + "`.`" + `+k] = v
				if gconv.String(v) == ` + "``" + ` {
					updateData[daoHandler.DbTable+` + "`.`" + `+k] = nil
				}`
				if gstr.Pos(tplDao, daoParseUpdateTmp) == -1 {
					daoParseUpdate += daoParseUpdateTmp
				}
			}
		} else if gstr.Pos(column[`Type`].String(), `char`) != -1 { //char类型
			if garray.NewStrArrayFrom([]string{`password`, `passwd`}).Contains(fieldSuffix) && column[`Type`].String() == `char(32)` { //password,passwd后缀
				daoParseInsertTmp := `
			case daoThis.Columns().` + fieldCaseCamel + `:
				password := gconv.String(v)
				if len(password) != 32 {
					password = gmd5.MustEncrypt(password)
				}`
				daoParseUpdateTmp := `
			case daoThis.Columns().` + fieldCaseCamel + `:
				password := gconv.String(v)
				if len(password) != 32 {
					password = gmd5.MustEncrypt(password)
				}`
				passwordHandleMapKey := MyGenPasswordHandleMapKey(field)
				if tpl.PasswordHandleMap[passwordHandleMapKey].IsCoexist {
					daoParseInsertTmp += `
				salt := grand.S(` + tpl.PasswordHandleMap[passwordHandleMapKey].SaltLength + `)
				insertData[daoThis.Columns().` + gstr.CaseCamel(tpl.PasswordHandleMap[passwordHandleMapKey].SaltField) + `] = salt
				password = gmd5.MustEncrypt(password + salt)`
					daoParseUpdateTmp += `
				salt := grand.S(` + tpl.PasswordHandleMap[passwordHandleMapKey].SaltLength + `)
				updateData[daoHandler.DbTable+` + "`.`" + `+daoThis.Columns().` + gstr.CaseCamel(tpl.PasswordHandleMap[passwordHandleMapKey].SaltField) + `] = salt
				password = gmd5.MustEncrypt(password + salt)`
				}
				daoParseInsertTmp += `
				insertData[k] = password`
				daoParseUpdateTmp += `
				updateData[daoHandler.DbTable+` + "`.`" + `+k] = password`
				if gstr.Pos(tplDao, daoParseInsertTmp) == -1 {
					daoParseInsert += daoParseInsertTmp
				}
				if gstr.Pos(tplDao, daoParseUpdateTmp) == -1 {
					daoParseUpdate += daoParseUpdateTmp
				}
			} else if garray.NewStrArrayFrom([]string{`salt`}).Contains(fieldSuffix) && tpl.PasswordHandleMap[MyGenPasswordHandleMapKey(field)].IsCoexist { //salt后缀
			} else {
				if column[`Key`].String() == `UNI` && column[`Null`].String() == `YES` {
					daoParseInsertTmp := `
			case daoThis.Columns().` + fieldCaseCamel + `:
				insertData[k] = v
				if gconv.String(v) == ` + "``" + ` {
					insertData[k] = nil
				}`
					if gstr.Pos(tplDao, daoParseInsertTmp) == -1 {
						daoParseInsert += daoParseInsertTmp
					}
					daoParseUpdateTmp := `
			case daoThis.Columns().` + fieldCaseCamel + `:
				updateData[daoHandler.DbTable+` + "`.`" + `+k] = v
				if gconv.String(v) == ` + "``" + ` {
					updateData[daoHandler.DbTable+` + "`.`" + `+k] = nil
				}`
					if gstr.Pos(tplDao, daoParseUpdateTmp) == -1 {
						daoParseUpdate += daoParseUpdateTmp
					}
				}
			}
		} else if gstr.Pos(column[`Type`].String(), `int`) != -1 && gstr.Pos(column[`Type`].String(), `point`) == -1 { //int等类型
			if field == `pid` { //pid
				if tpl.LabelHandle.LabelField != `` {
					daoParseFieldTmp := `
			case ` + "`p" + gstr.CaseCamel(tpl.LabelHandle.LabelField) + "`" + `:
				tableP := ` + "`p_`" + ` + daoHandler.DbTable
				m = m.Fields(tableP + ` + "`.`" + ` + daoThis.Columns().` + gstr.CaseCamel(tpl.LabelHandle.LabelField) + ` + ` + "` AS `" + ` + v)
				m = m.Handler(daoThis.ParseJoin(tableP, daoHandler))`
					if gstr.Pos(tplDao, daoParseFieldTmp) == -1 {
						daoParseField += daoParseFieldTmp
					}
				}
				daoParseFieldTmp := `
			case ` + "`tree`" + `:
				m = m.Fields(daoHandler.DbTable + ` + "`.`" + ` + daoThis.PrimaryKey())
				m = m.Fields(daoHandler.DbTable + ` + "`.`" + ` + daoThis.Columns().` + fieldCaseCamel + `)
				m = m.Handler(daoThis.ParseOrder([]string{` + "`tree`" + `}, daoHandler))`
				if gstr.Pos(tplDao, daoParseFieldTmp) == -1 {
					daoParseField += daoParseFieldTmp
				}
				daoParseOrderTmp := `
			case ` + "`tree`" + `:
				m = m.OrderAsc(daoHandler.DbTable + ` + "`.`" + ` + daoThis.Columns().` + fieldCaseCamel + `)`
				if tpl.PidHandle.SortField != `` {
					daoParseOrderTmp += `
				m = m.OrderAsc(daoHandler.DbTable + ` + "`.`" + ` + daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.SortField) + `)`
				}
				daoParseOrderTmp += `
				m = m.OrderAsc(daoHandler.DbTable + ` + "`.`" + ` + daoThis.PrimaryKey())`
				if gstr.Pos(tplDao, daoParseOrderTmp) == -1 {
					daoParseOrder += daoParseOrderTmp
				}
				daoParseJoinTmp := `
		case ` + "`p_`" + ` + daoHandler.DbTable:
			m = m.LeftJoin(daoHandler.DbTable+` + "` AS `" + `+joinTable, joinTable+` + "`.`" + `+daoThis.PrimaryKey()+` + "` = `" + `+daoHandler.DbTable+` + "`.`" + `+daoThis.Columns().` + fieldCaseCamel + `)`
				if gstr.Pos(tplDao, daoParseJoinTmp) == -1 {
					daoParseJoin += daoParseJoinTmp
				}

				if tpl.PidHandle.IsCoexist {
					daoParseInsertTmp := `
			case daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.PidField) + `:
				insertData[k] = v
				if gconv.Uint(v) > 0 {
					pInfo, _ := daoThis.ParseDbCtx(m.GetCtx()).Where(daoThis.PrimaryKey(), v).Fields(daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.IdPathField) + `, daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.LevelField) + `).One()
					daoHandler.AfterInsert[` + "`pIdPath`" + `] = pInfo[daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.IdPathField) + `].String()
					daoHandler.AfterInsert[` + "`pLevel`" + `] = pInfo[daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.LevelField) + `].Uint()
				} else {
					daoHandler.AfterInsert[` + "`pIdPath`" + `] = ` + "`0`" + `
					daoHandler.AfterInsert[` + "`pLevel`" + `] = 0
				}`
					if gstr.Pos(tplDao, daoParseInsertTmp) == -1 {
						daoParseInsert += daoParseInsertTmp
					}
					daoHookInsertTmp := `

			updateSelfData := map[string]interface{}{}
			for k, v := range daoHandler.AfterInsert {
				switch k {
				case ` + "`pIdPath`" + `:
					updateSelfData[daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.IdPathField) + `] = gconv.String(v) + ` + "`-`" + ` + gconv.String(id)
				case ` + "`pLevel`" + `:
					updateSelfData[daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.LevelField) + `] = gconv.Uint(v) + 1
				}
			}
			if len(updateSelfData) > 0 {
				daoThis.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), id).Data(updateSelfData).Update()
			}`
					if gstr.Pos(tplDao, daoHookInsertTmp) == -1 {
						daoHookInsert += daoHookInsertTmp
					}
					daoParseUpdateTmp := `
			case daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.PidField) + `:
				updateData[daoHandler.DbTable+` + "`.`" + `+k] = v
				pIdPath := ` + "`0`" + `
				var pLevel uint = 0
				if gconv.Uint(v) > 0 {
					pInfo, _ := daoThis.ParseDbCtx(m.GetCtx()).Where(daoThis.PrimaryKey(), v).Fields(daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.IdPathField) + `, daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.LevelField) + `).One()
					pIdPath = pInfo[daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.IdPathField) + `].String()
					pLevel = pInfo[daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.LevelField) + `].Uint()
				}
				updateData[daoHandler.DbTable+` + "`.`" + `+daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.IdPathField) + `] = gdb.Raw(` + "`CONCAT('`" + ` + pIdPath + ` + "`-', `" + ` + daoThis.PrimaryKey() + ` + "`)`" + `)
				updateData[daoHandler.DbTable+` + "`.`" + `+daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.LevelField) + `] = pLevel + 1`
					if gstr.Pos(tplDao, daoParseUpdateTmp) == -1 {
						daoParseUpdate += daoParseUpdateTmp
					}
					daoHookUpdateAfterTmp := `

			for k, v := range daoHandler.AfterUpdate {
				switch k {
				case ` + "`updateChildIdPathAndLevelList`" + `: //修改pid时，更新所有子孙级的idPath和level。参数：[]map[string]interface{}{newIdPath: 父级新idPath, oldIdPath: 父级旧idPath, newLevel: 父级新level, oldLevel: 父级旧level}
					val := v.([]map[string]interface{})
					for _, v1 := range val {
						daoThis.UpdateChildIdPathAndLevel(ctx, gconv.String(v1[` + "`newIdPath`" + `]), gconv.String(v1[` + "`oldIdPath`" + `]), gconv.Uint(v1[` + "`newLevel`" + `]), gconv.Uint(v1[` + "`oldLevel`" + `]))
					}
				}
			}`
					if gstr.Pos(tplDao, daoHookUpdateAfterTmp) == -1 {
						daoHookUpdateAfter += daoHookUpdateAfterTmp
					}
					daoFuncTmp := `
// 修改pid时，更新所有子孙级的idPath和level
func (daoThis *` + tpl.TableNameCaseCamelLower + `Dao) UpdateChildIdPathAndLevel(ctx context.Context, newIdPath string, oldIdPath string, newLevel uint, oldLevel uint) {
	data := g.Map{
		daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.IdPathField) + `: gdb.Raw(` + "`REPLACE(`" + ` + daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.IdPathField) + ` + ` + "`, '`" + ` + oldIdPath + ` + "`', '`" + ` + newIdPath + ` + "`')`" + `),
		daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.LevelField) + `:  gdb.Raw(daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.LevelField) + ` + ` + "` + `" + ` + gconv.String(newLevel-oldLevel)),
	}
	if newLevel < oldLevel {
		data[daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.LevelField) + `] = gdb.Raw(daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.LevelField) + ` + ` + "` - `" + ` + gconv.String(oldLevel-newLevel))
	}
	daoThis.ParseDbCtx(ctx).WhereLike(daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.IdPathField) + `, oldIdPath+` + "`-%`" + `).Data(data).Update()
}`
					if gstr.Pos(tplDao, daoFuncTmp) == -1 {
						daoFunc += daoFuncTmp
					}
				}
			} else if field == `level` && tpl.PidHandle.IsCoexist { //level
				daoParseOrderTmp := `
			case daoThis.Columns().` + fieldCaseCamel + `:
				m = m.Order(daoHandler.DbTable + ` + "`.`" + ` + v)
				m = m.OrderDesc(daoHandler.DbTable + ` + "`.`" + ` + daoThis.PrimaryKey())` //追加主键倒序。mysql排序字段有重复值时，分页会导致同一条数据可能在不同页都出现
				if gstr.Pos(tplDao, daoParseOrderTmp) == -1 {
					daoParseOrder += daoParseOrderTmp
				}
			} else if garray.NewStrArrayFrom([]string{`sort`, `weight`}).Contains(fieldSuffix) { //sort,weight等后缀
				daoParseOrderTmp := `
			case daoThis.Columns().` + fieldCaseCamel + `:
				m = m.Order(daoHandler.DbTable + ` + "`.`" + ` + v)
				m = m.OrderDesc(daoHandler.DbTable + ` + "`.`" + ` + daoThis.PrimaryKey())` //追加主键倒序。mysql排序字段有重复值时，分页会导致同一条数据可能在不同页都出现
				if gstr.Pos(tplDao, daoParseOrderTmp) == -1 {
					daoParseOrder += daoParseOrderTmp
				}
			} else if garray.NewStrArrayFrom([]string{`id`}).Contains(fieldSuffix) { //id后缀
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
					if !tpl.RelTableMap[field].IsRedundRelNameField {
						daoParseFieldTmp := `//因前端页面已用该字段名显示，故不存在时改成` + "`" + relTable.RelTableField + relTable.RelSuffix + "`" + `（控制器也要改）。同时下面Fields方法改成m = m.Fields(table` + relTable.RelTableNameCaseCamel + relTable.RelSuffixCaseCamel + ` + ` + "`.`" + ` + ` + daoPath + `.Columns().Xxxx + ` + "` AS `" + ` + v)`
						if gstr.Pos(tplDao, daoParseFieldTmp) == -1 {
							if relTable.RelSuffix != `` {
								daoParseFieldTmp = `
			case ` + daoPath + `.Columns().` + gstr.CaseCamel(relTable.RelTableField) + " + `" + relTable.RelSuffix + "`: " + daoParseFieldTmp + `
				table` + relTable.RelTableNameCaseCamel + relTable.RelSuffixCaseCamel + ` := ` + daoPath + `.ParseDbTable(m.GetCtx()) + ` + "`" + relTable.RelSuffixCaseSnake + "`" + `
				m = m.Fields(table` + relTable.RelTableNameCaseCamel + relTable.RelSuffixCaseCamel + ` + ` + "`.`" + ` + ` + daoPath + `.Columns().` + gstr.CaseCamel(relTable.RelTableField) + ` + ` + "` AS `" + ` + v)
				m = m.Handler(daoThis.ParseJoin(table` + relTable.RelTableNameCaseCamel + relTable.RelSuffixCaseCamel + `, daoHandler))`
							} else {
								daoParseFieldTmp = `
			case ` + daoPath + `.Columns().` + gstr.CaseCamel(relTable.RelTableField) + `: ` + daoParseFieldTmp + `
				table` + relTable.RelTableNameCaseCamel + ` := ` + daoPath + `.ParseDbTable(m.GetCtx())
				m = m.Fields(table` + relTable.RelTableNameCaseCamel + ` + ` + "`.`" + ` + v)
				m = m.Handler(daoThis.ParseJoin(table` + relTable.RelTableNameCaseCamel + `, daoHandler))`
							}
							daoParseField += daoParseFieldTmp
						}
					}
					daoParseJoinTmp := `
		case ` + daoPath + `.ParseDbTable(m.GetCtx()):
			m = m.LeftJoin(joinTable, joinTable+` + "`.`" + `+` + daoPath + `.PrimaryKey()+` + "` = `" + `+daoHandler.DbTable+` + "`.`" + `+daoThis.Columns().` + fieldCaseCamel + `)`
					if relTable.RelSuffix != `` {
						daoParseJoinTmp = `
		case ` + daoPath + `.ParseDbTable(m.GetCtx()) + ` + "`" + relTable.RelSuffixCaseSnake + "`" + `:
			m = m.LeftJoin(` + daoPath + `.ParseDbTable(m.GetCtx())+` + "` AS `" + `+joinTable, joinTable+` + "`.`" + `+` + daoPath + `.PrimaryKey()+` + "` = `" + `+daoHandler.DbTable+` + "`.`" + `+daoThis.Columns().` + fieldCaseCamel + `)`
					}
					if gstr.Pos(tplDao, daoParseJoinTmp) == -1 {
						daoParseJoin += daoParseJoinTmp
					}
				}
			}
		} else if gstr.Pos(column[`Type`].String(), `timestamp`) != -1 || gstr.Pos(column[`Type`].String(), `date`) != -1 { //timestamp或datetime或date类型
			if gstr.Pos(column[`Type`].String(), `date`) != -1 && gstr.Pos(column[`Type`].String(), `datetime`) == -1 {
				daoParseOrderTmp := `
			case daoThis.Columns().` + fieldCaseCamel + `:
				m = m.Order(daoHandler.DbTable + ` + "`.`" + ` + v)
				m = m.OrderDesc(daoHandler.DbTable + ` + "`.`" + ` + daoThis.PrimaryKey())` //追加主键倒序。mysql排序字段有重复值时，分页会导致同一条数据可能在不同页都出现
				if gstr.Pos(tplDao, daoParseOrderTmp) == -1 {
					daoParseOrder += daoParseOrderTmp
				}
			}

			if garray.NewStrArrayFrom([]string{`start`}).Contains(fieldPrefix) { //start_前缀
				daoParseFilterTmp := `
			case daoThis.Columns().` + fieldCaseCamel + `:
				m = m.WhereLTE(daoHandler.DbTable+` + "`.`" + `+k, v)`
				if column[`Null`].String() == `NO` && column[`Default`].String() == `` {
					daoParseFilterTmp = `
			case daoThis.Columns().` + fieldCaseCamel + `:
				m = m.Where(m.Builder().WhereLTE(daoHandler.DbTable+` + "`.`" + `+k, v).WhereOrNull(daoHandler.DbTable + ` + "`.`" + ` + k))`
				}
				if gstr.Pos(tplDao, daoParseFilterTmp) == -1 {
					daoParseFilter += daoParseFilterTmp
				}
			} else if garray.NewStrArrayFrom([]string{`end`}).Contains(fieldPrefix) { //end_前缀
				daoParseFilterTmp := `
			case daoThis.Columns().` + fieldCaseCamel + `:
				m = m.WhereGTE(daoHandler.DbTable+` + "`.`" + `+k, v)`
				if column[`Null`].String() == `NO` && column[`Default`].String() == `` {
					daoParseFilterTmp = `
			case daoThis.Columns().` + fieldCaseCamel + `:
				m = m.Where(m.Builder().WhereGTE(daoHandler.DbTable+` + "`.`" + `+k, v).WhereOrNull(daoHandler.DbTable + ` + "`.`" + ` + k))`
				}
				if gstr.Pos(tplDao, daoParseFilterTmp) == -1 {
					daoParseFilter += daoParseFilterTmp
				}
			}
		} else if gstr.Pos(column[`Type`].String(), `json`) != -1 { //json类型
			if column[`Null`].String() == `YES` {
				daoParseInsertTmp := `
			case daoThis.Columns().` + fieldCaseCamel + `:
				insertData[k] = v
				if gconv.String(v) == ` + "``" + ` {
					insertData[k] = nil
				}`
				if gstr.Pos(tplDao, daoParseInsertTmp) == -1 {
					daoParseInsert += daoParseInsertTmp
				}
				daoParseUpdateTmp := `
			case daoThis.Columns().` + fieldCaseCamel + `:
				updateData[daoHandler.DbTable+` + "`.`" + `+k] = gvar.New(v)
				if gconv.String(v) == ` + "``" + ` {
					updateData[daoHandler.DbTable+` + "`.`" + `+k] = nil
				}`
				if gstr.Pos(tplDao, daoParseUpdateTmp) == -1 {
					daoParseUpdate += daoParseUpdateTmp
				}
			}
		}
	}

	if daoParseInsert != `` {
		daoParseInsertPoint := `case ` + "`id`" + `:
				insertData[daoThis.PrimaryKey()] = v`
		tplDao = gstr.Replace(tplDao, daoParseInsertPoint, daoParseInsertPoint+daoParseInsert, 1)
	}
	if daoHookInsert != `` {
		daoHookInsertPoint := `// id, _ := result.LastInsertId()`
		tplDao = gstr.Replace(tplDao, daoHookInsertPoint, `id, _ := result.LastInsertId()`+daoHookInsert, 1)
	}
	if daoParseUpdate != `` {
		daoParseUpdatePoint := `case ` + "`id`" + `:
				updateData[daoHandler.DbTable+` + "`.`" + `+daoThis.PrimaryKey()] = v`
		tplDao = gstr.Replace(tplDao, daoParseUpdatePoint, daoParseUpdatePoint+daoParseUpdate, 1)
	}
	if daoHookUpdateBefore != `` || daoHookUpdateAfter != `` {
		daoHookUpdatePoint := `

			/* row, _ := result.RowsAffected()
			if row == 0 {
				return
			} */`
		if daoHookUpdateBefore != `` {
			tplDao = gstr.Replace(tplDao, daoHookUpdatePoint, daoHookUpdateBefore+daoHookUpdatePoint, 1)
		}
		if daoHookUpdateAfter != `` {
			tplDao = gstr.Replace(tplDao, daoHookUpdatePoint, `

			row, _ := result.RowsAffected()
			if row == 0 {
				return
			}`+daoHookUpdateAfter, 1)
		}
	}
	if daoParseField != `` {
		daoParseFieldPoint := `case ` + "`id`" + `:
				m = m.Fields(daoHandler.DbTable + ` + "`.`" + ` + daoThis.PrimaryKey() + ` + "` AS `" + ` + v)`
		tplDao = gstr.Replace(tplDao, daoParseFieldPoint, daoParseFieldPoint+daoParseField, 1)
	}
	if daoHookSelect != `` {
		daoHookSelectPoint := `
					default:
						record[v] = gvar.New(nil)`
		tplDao = gstr.Replace(tplDao, daoHookSelectPoint, daoHookSelect+daoHookSelectPoint, 1)
	}
	if daoParseFilter != `` {
		daoParseFilterPoint := `case ` + "`id`, `idArr`" + `:
				m = m.Where(daoHandler.DbTable+` + "`.`" + `+daoThis.PrimaryKey(), v)`
		tplDao = gstr.Replace(tplDao, daoParseFilterPoint, daoParseFilterPoint+daoParseFilter, 1)
	}
	if daoParseOrder != `` {
		daoParseOrderPoint := `case ` + "`id`" + `:
				m = m.Order(daoHandler.DbTable + ` + "`.`" + ` + gstr.Replace(v, k, daoThis.PrimaryKey(), 1))`
		tplDao = gstr.Replace(tplDao, daoParseOrderPoint, daoParseOrderPoint+daoParseOrder, 1)
	}
	if daoParseJoin != `` {
		daoParseJoinPoint := `
		/* case Xxxx.ParseDbTable(m.GetCtx()):
		m = m.LeftJoin(joinTable, joinTable+` + "`.`" + `+Xxxx.Columns().XxxxId+` + "` = `" + `+daoHandler.DbTable+` + "`.`" + `+daoThis.PrimaryKey())
		// m = m.LeftJoin(Xxxx.ParseDbTable(m.GetCtx())+` + "` AS `" + `+joinTable, joinTable+` + "`.`" + `+Xxxx.Columns().XxxxId+` + "` = `" + `+daoHandler.DbTable+` + "`.`" + `+daoThis.PrimaryKey()) */`
		tplDao = gstr.Replace(tplDao, daoParseJoinPoint, daoParseJoinPoint+daoParseJoin, 1)
	}
	if daoFunc != `` {
		tplDao = tplDao + daoFunc
	}
	if daoImportOtherDao != `` {
		daoImportOtherDaoPoint := `"api/internal/dao/` + tpl.ModuleDirCaseCamelLower + `/internal"`
		tplDao = gstr.Replace(tplDao, daoImportOtherDaoPoint, daoImportOtherDaoPoint+daoImportOtherDao, 1)
	}

	tplDao = gstr.Replace(tplDao, `"github.com/gogf/gf/v2/util/gconv"`, `"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/gogf/gf/v2/container/garray"`, 1)

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
	daoThis := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel
	if tpl.PidHandle.PidField != `` {
		tplLogic += `

	_, okPid := data[daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.PidField) + `]
	if okPid {
		pid := gconv.Uint(data[daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.PidField) + `])
		if pid > 0 {
			pInfo, _ := daoThis.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), pid).One()
			if pInfo.IsEmpty() {
				err = utils.NewErrorCode(ctx, 29999997, ` + "``" + `)
				return
			}
		}
	}`
		if tpl.PidHandle.IsCoexist {
			tplLogic += ` else {
		data[daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.PidField) + `] = 0
	}`
		}
		tplLogic += `
`
	}
	tplLogic += `
	id, err = daoThis.HandlerCtx(ctx).Insert(data).GetModel().InsertAndGetId()
	return
}

// 修改
func (logicThis *s` + tpl.LogicStructName + `) Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error) {
	daoThis := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `
	daoHandlerThis := daoThis.HandlerCtx(ctx).Filter(filter, true)
	if len(daoHandlerThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ` + "``" + `)
		return
	}
`
	if tpl.PidHandle.PidField != `` {
		tplLogic += `
	_, okPid := data[daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.PidField) + `]
	if okPid {`
		if tpl.PidHandle.IsCoexist {
			tplLogic += `
		pInfo := gdb.Record{}
		pid := gconv.Uint(data[daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.PidField) + `])
		if pid > 0 {
			pInfo, _ = daoThis.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), pid).One()
			if pInfo.IsEmpty() {
				err = utils.NewErrorCode(ctx, 29999997, ` + "``" + `)
				return
			}
		}
		updateChildIdPathAndLevelList := []map[string]interface{}{}
		for _, id := range daoHandlerThis.IdArr {
			if pid == id { //父级不能是自身
				err = utils.NewErrorCode(ctx, 29999996, ` + "``" + `)
				return
			}
			oldInfo, _ := daoThis.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), id).One()
			if pid != oldInfo[daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.PidField) + `].Uint() {
				pIdPath := ` + "`0`" + `
				var pLevel uint = 0
				if pid > 0 {
					if garray.NewStrArrayFrom(gstr.Split(pInfo[daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.IdPathField) + `].String(), ` + "`-`" + `)).Contains(oldInfo[daoThis.PrimaryKey()].String()) { //父级不能是自身的子孙级
						err = utils.NewErrorCode(ctx, 29999995, ` + "``" + `)
						return
					}
					pIdPath = pInfo[daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.IdPathField) + `].String()
					pLevel = pInfo[daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.LevelField) + `].Uint()
				}
				updateChildIdPathAndLevelList = append(updateChildIdPathAndLevelList, map[string]interface{}{
					` + "`" + `newIdPath` + "`" + `: pIdPath + ` + "`-`" + ` + gconv.String(id),
					` + "`" + `oldIdPath` + "`" + `: oldInfo[daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.IdPathField) + `],
					` + "`" + `newLevel` + "`" + `:  pLevel + 1,
					` + "`" + `oldLevel` + "`" + `:  oldInfo[daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.LevelField) + `],
				})
			}
		}
		if len(updateChildIdPathAndLevelList) > 0 {
			daoHandlerThis.AfterUpdate[` + "`" + `updateChildIdPathAndLevelList` + "`" + `] = updateChildIdPathAndLevelList
		}`
		} else {
			tplLogic += `
		pid := gconv.Uint(data[daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.PidField) + `])
		if pid > 0 {
			pInfo, _ := daoThis.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), pid).One()
			if pInfo.IsEmpty() {
				err = utils.NewErrorCode(ctx, 29999997, ` + "``" + `)
				return
			}
		}
		for _, id := range daoHandlerThis.IdArr {
			if pid == id { //父级不能是自身
				err = utils.NewErrorCode(ctx, 29999996, ` + "``" + `)
				return
			}
		}`
		}
		tplLogic += `
	}
`
	}
	tplLogic += `
	row, err = daoHandlerThis.Update(data).GetModel().UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *s` + tpl.LogicStructName + `) Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error) {
	daoThis := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `
	daoHandlerThis := daoThis.HandlerCtx(ctx).Filter(filter, true)
	if len(daoHandlerThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ` + "``" + `)
		return
	}
`
	if tpl.PidHandle.PidField != `` {
		tplLogic += `
	count, _ := daoThis.ParseDbCtx(ctx).Where(daoThis.Columns().` + gstr.CaseCamel(tpl.PidHandle.PidField) + `, daoHandlerThis.IdArr).Count()
	if count > 0 {
		err = utils.NewErrorCode(ctx, 29999994, ` + "``" + `)
		return
	}
`
	}
	tplLogic += `
	result, err := daoHandlerThis.Delete().GetModel().Delete()
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
	if tpl.LabelHandle.LabelField != `` {
		apiReqFilterColumn += `Label          string      ` + "`" + `json:"label,omitempty" v:"max-length:30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"标签。常用于前端组件"` + "`\n"
		apiResColumn += `Label       *string     ` + "`" + `json:"label,omitempty" dc:"标签。常用于前端组件"` + "`\n"
	}
	for _, column := range tpl.TableColumnList {
		field := column[`Field`].String()
		fieldCaseCamel := gstr.CaseCamel(field)
		fieldCaseSnake := gstr.CaseSnakeFirstUpper(field)
		fieldCaseSnakeOfRemove := gstr.Split(fieldCaseSnake, `_of_`)[0]
		fieldCaseCamelOfRemove := gstr.CaseCamel(fieldCaseSnakeOfRemove)
		fieldSplitArr := gstr.Split(fieldCaseSnakeOfRemove, `_`)
		fieldPrefix := fieldSplitArr[0]
		fieldSuffix := fieldSplitArr[len(fieldSplitArr)-1]
		comment := gstr.Trim(gstr.ReplaceByArray(column[`Comment`].String(), g.SliceStr{
			"\n", ` `,
			"\r", ` `,
			`"`, `\"`,
		}))
		resultStr, _ := gregex.MatchString(`.*\((\d*)\)`, column[`Type`].String())

		typeReqFilter := ``
		typeReqCreate := ``
		typeReqUpdate := ``
		typeRes := ``
		ruleReqFilter := ``
		ruleReqCreate := ``
		ruleReqUpdate := ``
		isRequired := false

		if garray.NewStrArrayFrom([]string{`DeletedAt`, `DeleteAt`, `DeletedTime`, `DeleteTime`}).Contains(fieldCaseCamel) {
			continue
		} else if garray.NewStrArrayFrom([]string{`UpdatedAt`, `UpdateAt`, `UpdatedTime`, `UpdateTime`}).Contains(fieldCaseCamel) {
			typeRes = `*gtime.Time`
		} else if garray.NewStrArrayFrom([]string{`CreatedAt`, `CreateAt`, `CreatedTime`, `CreateTime`}).Contains(fieldCaseCamel) {
			apiReqFilterColumn += `TimeRangeStart *gtime.Time ` + "`" + `json:"timeRangeStart,omitempty" v:"date-format:Y-m-d H:i:s" dc:"开始时间：YYYY-mm-dd HH:ii:ss"` + "`\n"
			apiReqFilterColumn += `TimeRangeEnd   *gtime.Time ` + "`" + `json:"timeRangeEnd,omitempty" v:"date-format:Y-m-d H:i:s|after-equal:TimeRangeStart" dc:"结束时间：YYYY-mm-dd HH:ii:ss"` + "`\n"
			typeRes = `*gtime.Time`
		} else if column[`Key`].String() == `PRI` && column[`Extra`].String() == `auto_increment` { //主键
			if field == `id` {
				continue
			}
			typeReqFilter = `*uint`
			typeRes = `*uint`
			ruleReqFilter = `min:1`
		} else if fieldCaseCamel == `IdPath` && (gstr.Pos(column[`Type`].String(), `varchar`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1) && tpl.PidHandle.IsCoexist { //idPath|id_path
			typeRes = `*string`
		} else if garray.NewStrArrayFrom([]string{`status`, `type`, `method`, `pos`, `position`, `gender`}).Contains(fieldSuffix) && ((gstr.Pos(column[`Type`].String(), `int`) != -1 && gstr.Pos(column[`Type`].String(), `point`) == -1) || gstr.Pos(column[`Type`].String(), `char`) != -1) { //status,type,method,pos,position,gender等后缀
			typeReqFilter = `string`
			typeReqCreate = `*string`
			typeReqUpdate = `*string`
			typeRes = `*string`
			isStr := true
			if gstr.Pos(column[`Type`].String(), `int`) != -1 && gstr.Pos(column[`Type`].String(), `point`) == -1 {
				typeReqFilter = `*int`
				typeReqCreate = `*int`
				typeReqUpdate = `*int`
				typeRes = `*int`
				if gstr.Pos(column[`Type`].String(), `unsigned`) != -1 {
					typeReqFilter = `*uint`
					typeReqCreate = `*uint`
					typeReqUpdate = `*uint`
					typeRes = `*uint`
				}
				isStr = false
			}

			statusList := MyGenStatusList(comment, isStr)
			statusArr := make([]string, len(statusList))
			for index, status := range statusList {
				statusArr[index] = status[0]
			}
			statusStr := gstr.Join(statusArr, `,`)
			ruleReqFilter += `in:` + statusStr
			ruleReqCreate += `in:` + statusStr
			ruleReqUpdate += `in:` + statusStr
		} else if (garray.NewStrArrayFrom([]string{`icon`, `cover`, `avatar`, `img`, `image`}).Contains(fieldSuffix) || gstr.SubStr(fieldCaseCamelOfRemove, -7) == `ImgList` || gstr.SubStr(fieldCaseCamelOfRemove, -6) == `ImgArr` || gstr.SubStr(fieldCaseCamelOfRemove, -9) == `ImageList` || gstr.SubStr(fieldCaseCamelOfRemove, -8) == `ImageArr` || garray.NewStrArrayFrom([]string{`video`}).Contains(fieldSuffix) || gstr.SubStr(fieldCaseCamelOfRemove, -9) == `VideoList` || gstr.SubStr(fieldCaseCamelOfRemove, -8) == `VideoArr`) && (gstr.Pos(column[`Type`].String(), `varchar`) != -1 || gstr.Pos(column[`Type`].String(), `json`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1) { //icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀 //video,video_list,videoList,video_arr,videoArr等后缀
			if gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
				typeReqCreate = `*string`
				typeReqUpdate = `*string`
				typeRes = `*string`
				ruleReqCreate = `max-length:` + resultStr[1] + `|url`
				ruleReqUpdate = `max-length:` + resultStr[1] + `|url`
			} else {
				if column[`Null`].String() == `NO` {
					isRequired = true
				}
				typeReqCreate = `*[]string`
				typeReqUpdate = `*[]string`
				typeRes = `[]string`
				ruleReqCreate = `distinct|foreach|url|foreach|min-length:1`
				ruleReqUpdate = `distinct|foreach|url|foreach|min-length:1`
			}
		} else if garray.NewStrArrayFrom([]string{`list`, `arr`}).Contains(fieldSuffix) && (gstr.Pos(column[`Type`].String(), `json`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1) { //list,arr等后缀
			if column[`Null`].String() == `NO` {
				isRequired = true
			}
			typeReqCreate = `*[]interface{}`
			typeReqUpdate = `*[]interface{}`
			typeRes = `[]interface{}`
			ruleReqCreate = `distinct`
			ruleReqUpdate = `distinct`
		} else if garray.NewStrArrayFrom([]string{`remark`, `desc`, `msg`, `message`, `intro`, `content`}).Contains(fieldSuffix) && (gstr.Pos(column[`Type`].String(), `varchar`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1) { //remark,desc,msg,message,intro,content后缀
			typeReqCreate = `*string`
			typeReqUpdate = `*string`
			typeRes = `*string`
			if gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
				ruleReqCreate = `max-length:` + resultStr[1]
				ruleReqUpdate = `max-length:` + resultStr[1]
			}
		} else if gstr.Pos(column[`Type`].String(), `varchar`) != -1 { //varchar类型
			typeReqFilter = `string`
			typeReqCreate = `*string`
			typeReqUpdate = `*string`
			typeRes = `*string`
			ruleReqFilter = `max-length:` + resultStr[1]
			ruleReqCreate = `max-length:` + resultStr[1]
			ruleReqUpdate = `max-length:` + resultStr[1]
			if column[`Key`].String() == `UNI` && column[`Null`].String() == `NO` {
				isRequired = true
			}

			if garray.NewStrArrayFrom([]string{`name`}).Contains(fieldSuffix) { //name后缀
				if gstr.SubStr(gstr.CaseCamel(tpl.PrimaryKey), 0, -2)+`Name` == fieldCaseCamel {
					isRequired = true
				}
				ruleReqFilter += `|regex:^[\\p{L}\\p{M}\\p{N}_-]+$`
				ruleReqCreate += `|regex:^[\\p{L}\\p{M}\\p{N}_-]+$`
				ruleReqUpdate += `|regex:^[\\p{L}\\p{M}\\p{N}_-]+$`
			} else if garray.NewStrArrayFrom([]string{`code`}).Contains(fieldSuffix) { //code后缀
				ruleReqFilter += `|regex:^[\\p{L}\\p{M}\\p{N}_-]+$`
				ruleReqCreate += `|regex:^[\\p{L}\\p{M}\\p{N}_-]+$`
				ruleReqUpdate += `|regex:^[\\p{L}\\p{M}\\p{N}_-]+$`
			} else if garray.NewStrArrayFrom([]string{`phone`, `mobile`}).Contains(fieldSuffix) { //mobile,phone后缀
				ruleReqFilter += `|phone`
				ruleReqCreate += `|phone`
				ruleReqUpdate += `|phone`
			} else if garray.NewStrArrayFrom([]string{`url`, `link`}).Contains(fieldSuffix) { //url,link后缀
				ruleReqFilter += `|url`
				ruleReqCreate += `|url`
				ruleReqUpdate += `|url`
			} else if garray.NewStrArrayFrom([]string{`ip`}).Contains(fieldSuffix) { //IP后缀
				ruleReqFilter += `|ip`
				ruleReqCreate += `|ip`
				ruleReqUpdate += `|ip`
			}
		} else if gstr.Pos(column[`Type`].String(), `char`) != -1 { //char类型
			typeReqFilter = `string`
			typeReqCreate = `*string`
			typeReqUpdate = `*string`
			typeRes = `*string`
			ruleReqFilter = `max-length:` + resultStr[1]
			ruleReqCreate = `size:` + resultStr[1]
			ruleReqUpdate = `size:` + resultStr[1]
			if garray.NewStrArrayFrom([]string{`password`, `passwd`}).Contains(fieldSuffix) && column[`Type`].String() == `char(32)` { //password,passwd后缀
				typeReqFilter = ``
				typeRes = ``
				ruleReqFilter = ``
				isRequired = true
			} else if garray.NewStrArrayFrom([]string{`salt`}).Contains(fieldSuffix) && tpl.PasswordHandleMap[MyGenPasswordHandleMapKey(field)].IsCoexist { //salt后缀
				continue
			} else {
				if column[`Key`].String() == `UNI` && column[`Null`].String() == `NO` {
					isRequired = true
				}
			}
		} else if gstr.Pos(column[`Type`].String(), `int`) != -1 && gstr.Pos(column[`Type`].String(), `point`) == -1 { //int等类型
			typeReqFilter = `*int`
			typeReqCreate = `*int`
			typeReqUpdate = `*int`
			typeRes = `*int`
			ruleReqFilter = ``
			ruleReqCreate = ``
			ruleReqUpdate = ``
			if gstr.Pos(column[`Type`].String(), `unsigned`) != -1 {
				typeReqFilter = `*uint`
				typeReqCreate = `*uint`
				typeReqUpdate = `*uint`
				typeRes = `*uint`
			}

			if field == `pid` { //pid
				if tpl.LabelHandle.LabelField != `` {
					apiResColumnAlloweFieldList += `P` + gstr.CaseCamel(tpl.LabelHandle.LabelField) + ` *string ` + "`" + `json:"p` + gstr.CaseCamel(tpl.LabelHandle.LabelField) + `,omitempty" dc:"父级"` + "`\n"
				}
			} else if field == `level` && tpl.PidHandle.IsCoexist { //level
				typeReqCreate = ``
				typeReqUpdate = ``
				ruleReqFilter += `min:1`
			} else if garray.NewStrArrayFrom([]string{`sort`, `weight`}).Contains(fieldSuffix) { //sort,weight等后缀
				typeReqFilter = ``
				ruleReqCreate += `between:0,100`
				ruleReqUpdate += `between:0,100`
			} else if garray.NewStrArrayFrom([]string{`id`}).Contains(fieldSuffix) { //id后缀
				if tpl.RelTableMap[field].IsExistRelTableDao && !tpl.RelTableMap[field].IsRedundRelNameField {
					relTable := tpl.RelTableMap[field]
					apiResColumnAlloweFieldList += gstr.CaseCamel(relTable.RelTableField) + relTable.RelSuffixCaseCamel + ` *string ` + "`" + `json:"` + relTable.RelTableField + relTable.RelSuffix + `,omitempty" dc:"` + relTable.RelTableFieldName + `"` + "`\n"
				}
				ruleReqFilter += `min:1`
				ruleReqCreate += `min:1`
				ruleReqUpdate += `min:1`
			} else if garray.NewStrArrayFrom([]string{`is`}).Contains(fieldPrefix) { //is_前缀
				ruleReqFilter += `in:0,1`
				ruleReqCreate += `in:0,1`
				ruleReqUpdate += `in:0,1`
			} else { //默认处理（int等类型）
				typeReqFilter = ``
			}
		} else if gstr.Pos(column[`Type`].String(), `decimal`) != -1 || gstr.Pos(column[`Type`].String(), `double`) != -1 || gstr.Pos(column[`Type`].String(), `float`) != -1 { //float类型
			typeReqFilter = `*float64`
			typeReqCreate = `*float64`
			typeReqUpdate = `*float64`
			typeRes = `*float64`
			ruleReqFilter = ``
			ruleReqCreate = ``
			ruleReqUpdate = ``
			if gstr.Pos(column[`Type`].String(), `unsigned`) != -1 {
				ruleReqFilter += `min:0`
				ruleReqCreate += `min:0`
				ruleReqUpdate += `min:0`
			}

			//默认处理（float类型）
			typeReqFilter = ``
		} else if gstr.Pos(column[`Type`].String(), `timestamp`) != -1 || gstr.Pos(column[`Type`].String(), `date`) != -1 { //timestamp或datetime或date类型
			typeReqFilter = ``
			typeReqCreate = `*gtime.Time`
			typeReqUpdate = `*gtime.Time`
			typeRes = `*gtime.Time`
			ruleReqFilter = `date-format:Y-m-d H:i:s`
			ruleReqCreate = `date-format:Y-m-d H:i:s`
			ruleReqUpdate = `date-format:Y-m-d H:i:s`
			if gstr.Pos(column[`Type`].String(), `date`) != -1 && gstr.Pos(column[`Type`].String(), `datetime`) == -1 {
				typeReqFilter = `*gtime.Time`
				typeRes = `*string`
				ruleReqFilter = `date-format:Y-m-d`
				ruleReqCreate = `date-format:Y-m-d`
				ruleReqUpdate = `date-format:Y-m-d`
			}
			if column[`Null`].String() == `NO` && column[`Default`].String() == `` {
				isRequired = true
			}

			if garray.NewStrArrayFrom([]string{`start`}).Contains(fieldPrefix) || garray.NewStrArrayFrom([]string{`end`}).Contains(fieldPrefix) { //start_前缀 //end_前缀
				typeReqFilter = `*gtime.Time`
			}
		} else if gstr.Pos(column[`Type`].String(), `json`) != -1 { //json类型
			if column[`Null`].String() == `NO` {
				isRequired = true
			}
			typeReqCreate = `*string`
			typeReqUpdate = `*string`
			typeRes = `*string`
			ruleReqCreate = `json`
			ruleReqUpdate = `json`
		} else if gstr.Pos(column[`Type`].String(), `text`) != -1 { //text类型
			typeReqCreate = `*string`
			typeReqUpdate = `*string`
			typeRes = `*string`
		} else { //默认处理
			typeReqFilter = `string`
			typeReqCreate = `*string`
			typeReqUpdate = `*string`
			typeRes = `*string`
		}

		if typeReqFilter != `` {
			apiReqFilterColumn += fieldCaseCamel + ` ` + typeReqFilter + ` ` + "`" + `json:"` + field + `,omitempty" v:"` + ruleReqFilter + `" dc:"` + comment + `"` + "`\n"
		}
		if typeReqCreate != `` {
			if isRequired {
				if ruleReqCreate == `` {
					ruleReqCreate = `required`
				} else {
					ruleReqCreate = `required|` + ruleReqCreate
				}
			}
			apiReqCreateColumn += fieldCaseCamel + ` ` + typeReqCreate + ` ` + "`" + `json:"` + field + `,omitempty" v:"` + ruleReqCreate + `" dc:"` + comment + `"` + "`\n"
		}
		if typeReqUpdate != `` {
			apiReqUpdateColumn += fieldCaseCamel + ` ` + typeReqUpdate + ` ` + "`" + `json:"` + field + `,omitempty" v:"` + ruleReqUpdate + `" dc:"` + comment + `"` + "`\n"
		}
		if typeRes != `` {
			apiResColumn += fieldCaseCamel + ` ` + typeRes + ` ` + "`" + `json:"` + field + `,omitempty" dc:"` + comment + `"` + "`\n"
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
	Filter ` + tpl.TableNameCaseCamel + `ListFilter ` + "`" + `json:"filter" dc:"过滤条件"` + "`" + `
	Field  []string        ` + "`" + `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"` + "`" + `
	Sort   string          ` + "`" + `json:"sort" default:"id DESC" dc:"排序"` + "`" + `
	Page   int             ` + "`" + `json:"page" v:"min:1" default:"1" dc:"页码"` + "`" + `
	Limit  int             ` + "`" + `json:"limit" v:"min:0" default:"10" dc:"每页数量。可传0取全部"` + "`" + `
}

type ` + tpl.TableNameCaseCamel + `ListFilter struct {
	Id             *uint       ` + "`" + `json:"id,omitempty" v:"min:1" dc:"ID"` + "`" + `
	IdArr          []uint      ` + "`" + `json:"idArr,omitempty" v:"distinct|foreach|min:1" dc:"ID数组"` + "`" + `
	ExcId          *uint       ` + "`" + `json:"excId,omitempty" v:"min:1" dc:"排除ID"` + "`" + `
	ExcIdArr       []uint      ` + "`" + `json:"excIdArr,omitempty" v:"distinct|foreach|min:1" dc:"排除ID数组"` + "`" + `
	` + apiReqFilterColumn + `
}

type ` + tpl.TableNameCaseCamel + `ListRes struct {`
		if option.IsCount {
			tplApi += `
	Count int         ` + "`" + `json:"count" dc:"总数"` + "`"
		}
		tplApi += `
	List  []` + tpl.TableNameCaseCamel + `ListItem ` + "`" + `json:"list" dc:"列表"` + "`" + `
}

type ` + tpl.TableNameCaseCamel + `ListItem struct {
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
	Id     uint     ` + "`" + `json:"id" v:"required|min:1" dc:"ID"` + "`" + `
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
	g.Meta      ` + "`" + `path:"/` + tpl.TableNameCaseCamelLower + `/create" method:"post" tags:"` + tpl.SceneName + `/` + option.CommonName + `" sm:"新增"` + "`" + `
	` + apiReqCreateColumn + `
}

/*--------新增 结束--------*/

`
	}

	if option.IsUpdate {
		tplApi += `/*--------修改 开始--------*/
type ` + tpl.TableNameCaseCamel + `UpdateReq struct {
	g.Meta      ` + "`" + `path:"/` + tpl.TableNameCaseCamelLower + `/update" method:"post" tags:"` + tpl.SceneName + `/` + option.CommonName + `" sm:"修改"` + "`" + `
	IdArr       []uint  ` + "`" + `json:"idArr,omitempty" v:"required|distinct|foreach|min:1" dc:"ID数组"` + "`" + `
	` + apiReqUpdateColumn + `
}

/*--------修改 结束--------*/

`
	}

	if option.IsDelete {
		tplApi += `/*--------删除 开始--------*/
type ` + tpl.TableNameCaseCamel + `DeleteReq struct {
	g.Meta ` + "`" + `path:"/` + tpl.TableNameCaseCamelLower + `/del" method:"post" tags:"` + tpl.SceneName + `/` + option.CommonName + `" sm:"删除"` + "`" + `
	IdArr  []uint ` + "`" + `json:"idArr,omitempty" v:"required|distinct|foreach|min:1" dc:"ID数组"` + "`" + `
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
	Tree []` + tpl.TableNameCaseCamel + `TreeItem ` + "`" + `json:"tree" dc:"列表（树状）"` + "`" + `
}

type ` + tpl.TableNameCaseCamel + `TreeItem struct {
	Id       *uint       ` + "`" + `json:"id,omitempty" dc:"ID"` + "`" + `
	` + apiResColumn + `
	Children []` + tpl.TableNameCaseCamel + `TreeItem ` + "`" + `json:"children" dc:"子级列表"` + "`" + `
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
	// controllerAlloweFieldDiff := `` // 可以不要。数据返回时，会根据API文件中的结构体做过滤
	daoImportOtherDao := ``
	if tpl.LabelHandle.LabelField != `` {
		controllerAlloweFieldList += "`label`, "
		controllerAlloweFieldInfo += "`label`, "
		controllerAlloweFieldTree += "`label`, "
		if tpl.PidHandle.PidField != `` {
			controllerAlloweFieldList += "`p" + gstr.CaseCamel(tpl.LabelHandle.LabelField) + "`, "
			// controllerAlloweFieldInfo += "`p" + gstr.CaseCamel(tpl.LabelHandle.LabelField) + "`, "
		}
		controllerAlloweFieldNoAuth += "`label`, "
		if tpl.LabelHandle.IsCoexist {
			controllerAlloweFieldNoAuth += `dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `.Columns().Phone, ` + `dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `.Columns().Account, `
		} else {
			controllerAlloweFieldNoAuth += `dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `.Columns().` + gstr.CaseCamel(tpl.LabelHandle.LabelField) + `, `
		}
	}
	for _, column := range tpl.TableColumnList {
		field := column[`Field`].String()
		fieldCaseCamel := gstr.CaseCamel(field)
		fieldCaseSnake := gstr.CaseSnakeFirstUpper(field)
		fieldCaseSnakeOfRemove := gstr.Split(fieldCaseSnake, `_of_`)[0]
		// fieldCaseCamelOfRemove := gstr.CaseCamel(fieldCaseSnakeOfRemove)
		fieldSplitArr := gstr.Split(fieldCaseSnakeOfRemove, `_`)
		// fieldPrefix := fieldSplitArr[0]
		fieldSuffix := fieldSplitArr[len(fieldSplitArr)-1]

		if column[`Key`].String() == `PRI` && column[`Extra`].String() == `auto_increment` { //主键
			if field != `id` {
				controllerAlloweFieldNoAuth += `dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `.Columns().` + fieldCaseCamel + `, `
			}
		} else if gstr.Pos(column[`Type`].String(), `char`) != -1 { //char类型
			if garray.NewStrArrayFrom([]string{`password`, `passwd`}).Contains(fieldSuffix) && column[`Type`].String() == `char(32)` { //password,passwd后缀
				// controllerAlloweFieldDiff += `dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `.Columns().` + fieldCaseCamel + `, `
			} else if garray.NewStrArrayFrom([]string{`salt`}).Contains(fieldSuffix) && tpl.PasswordHandleMap[MyGenPasswordHandleMapKey(field)].IsCoexist { //salt后缀
				// controllerAlloweFieldDiff += `dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `.Columns().` + fieldCaseCamel + `, `
			}
		} else if gstr.Pos(column[`Type`].String(), `int`) != -1 && gstr.Pos(column[`Type`].String(), `point`) == -1 { //int等类型
			if garray.NewStrArrayFrom([]string{`id`}).Contains(fieldSuffix) { //id后缀
				if tpl.RelTableMap[field].IsExistRelTableDao && !tpl.RelTableMap[field].IsRedundRelNameField {
					relTable := tpl.RelTableMap[field]
					// controllerAlloweFieldList += "`" + relTable.RelNameField + "`, "
					daoPath := `dao` + relTable.RelDaoDirCaseCamel + `.` + relTable.RelTableNameCaseCamel
					daoImportOtherDao += `
	dao` + relTable.RelDaoDirCaseCamel + ` "api/internal/dao/` + relTable.RelDaoDir + `"`
					controllerAlloweFieldList += daoPath + `.Columns().` + gstr.CaseCamel(relTable.RelTableField)
					if relTable.RelSuffix != `` {
						controllerAlloweFieldList += "+`" + relTable.RelSuffix + "`"
					}
					controllerAlloweFieldList += ", "
				}
			}
		}
	}
	controllerAlloweFieldList = gstr.SubStr(controllerAlloweFieldList, 0, -len(`, `))
	controllerAlloweFieldInfo = gstr.SubStr(controllerAlloweFieldInfo, 0, -len(`, `))
	controllerAlloweFieldTree = gstr.SubStr(controllerAlloweFieldTree, 0, -len(`, `))
	controllerAlloweFieldNoAuth = gstr.SubStr(controllerAlloweFieldNoAuth, 0, -len(`, `))
	// controllerAlloweFieldDiff = gstr.SubStr(controllerAlloweFieldDiff, 0, -len(`, `))

	tplController := `package controller

import (
	"api/api"
	api` + tpl.ModuleDirCaseCamel + ` "api/api/` + option.SceneCode + `/` + tpl.ModuleDirCaseCamelLower + `"
	dao` + tpl.ModuleDirCaseCamel + ` "api/internal/dao/` + tpl.ModuleDirCaseCamelLower + `"` + daoImportOtherDao + `
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
	filter := gconv.Map(req.Filter, gconv.MapOption{Deep: true, OmitEmpty: true})
	if filter == nil {
		filter = map[string]interface{}{}
	}
`
		tplController += `
	allowField := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `.ColumnArr()
	allowField = append(allowField, ` + controllerAlloweFieldList + `)`
		/* if controllerAlloweFieldDiff != `` {
				tplController += `
		allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{` + controllerAlloweFieldDiff + `})).Slice() //移除敏感字段`
			} */
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
	daoHandlerThis := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `.HandlerCtx(ctx).Filter(filter)`
		if option.IsCount {
			tplController += `
	count, err := daoHandlerThis.Count()
	if err != nil {
		return
	}`
		}
		tplController += `
	list, err := daoHandlerThis.Field(field).Order([]string{req.Sort}).JoinGroupByPrimaryKey().GetModel().Page(req.Page, req.Limit).All()
	if err != nil {
		return
	}

	res = &api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `ListRes{`
		if option.IsCount {
			tplController += `Count: count, `
		}
		tplController += `List:  []api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `ListItem{}}
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
		/* if controllerAlloweFieldDiff != `` {
				tplController += `
		allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{` + controllerAlloweFieldDiff + `})).Slice() //移除敏感字段`
			} */
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
	info, err := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `.HandlerCtx(ctx).Filter(filter).Field(field).JoinGroupByPrimaryKey().GetModel().One()
	if err != nil {
		return
	}
	if info.IsEmpty() {
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
	data := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true})
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
	data := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true})
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
	filter := gconv.Map(req.Filter, gconv.MapOption{Deep: true, OmitEmpty: true})
	if filter == nil {
		filter = map[string]interface{}{}
	}

	allowField := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `.ColumnArr()
	allowField = append(allowField, ` + controllerAlloweFieldTree + `)`
		/* if controllerAlloweFieldDiff != `` {
				tplController += `
		allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{` + controllerAlloweFieldDiff + `})).Slice() //移除敏感字段`
			} */
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

	list, err :=dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `.HandlerCtx(ctx).Filter(filter).Field(field).JoinGroupByPrimaryKey().GetModel().All()
	if err != nil {
		return
	}
	tree := utils.Tree(list.List(), 0, dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `.Columns().` + gstr.CaseCamel(tpl.PrimaryKey) + `, dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `.Columns().` + gstr.CaseCamel(tpl.PidHandle.PidField) + `)

	res = &api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableNameCaseCamel + `TreeRes{}
	gconv.Structs(tree, &res.Tree)
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
	"api/internal/middleware"`, 1)
		//路由生成
		tplRouter = gstr.Replace(tplRouter, `/*--------后端路由自动代码生成锚点（不允许修改和删除，否则将不能自动生成路由）--------*/`, `group.Group(`+"`"+`/`+tpl.ModuleDirCaseCamelLower+"`"+`, func(group *ghttp.RouterGroup) {
				group.Bind(controller`+tpl.ModuleDirCaseCamel+`.New`+tpl.TableNameCaseCamel+`())
			})

			/*--------后端路由自动代码生成锚点（不允许修改和删除，否则将不能自动生成路由）--------*/`, 1)
		gfile.PutContents(saveFile, tplRouter)
	} else {
		//路由不存在时需生成
		if gstr.Pos(tplRouter, `group.Bind(controller`+tpl.ModuleDirCaseCamel+`.New`+tpl.TableNameCaseCamel+`())`) == -1 {
			//路由生成
			tplRouter = gstr.Replace(tplRouter, `group.Group(`+"`"+`/`+tpl.ModuleDirCaseCamelLower+"`"+`, func(group *ghttp.RouterGroup) {`, `group.Group(`+"`"+`/`+tpl.ModuleDirCaseCamelLower+"`"+`, func(group *ghttp.RouterGroup) {
				group.Bind(controller`+tpl.ModuleDirCaseCamel+`.New`+tpl.TableNameCaseCamel+`())`, 1)
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

	tplView := `<script setup lang="tsx">
import List from './List.vue'
import Query from './Query.vue'`
	if option.IsCreate || option.IsUpdate {
		tplView += `
import Save from './Save.vue'`
	}
	tplView += `

//搜索
const queryCommon = reactive({
    data: {},
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
    title: '', //新增|编辑|复制
    data: {},
})
provide('saveCommon', saveCommon)`
	}
	tplView += `
</script>

<template>
    <el-container class="main-table-container">
        <el-header>
            <query />
        </el-header>

        <list :ref="(el: any) => listCommon.ref = el" />`
	if option.IsCreate || option.IsUpdate {
		tplView += `

        <!-- 加上v-if每次都重新生成组件。可防止不同操作之间的影响；新增操作数据的默认值也能写在save组件内 -->
        <save v-if="saveCommon.visible" />`
	}
	tplView += `
    </el-container>
</template>
`

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
		fieldCaseSnakeOfRemove := gstr.Split(fieldCaseSnake, `_of_`)[0]
		fieldCaseCamelOfRemove := gstr.CaseCamel(fieldCaseSnakeOfRemove)
		fieldSplitArr := gstr.Split(fieldCaseSnakeOfRemove, `_`)
		fieldPrefix := fieldSplitArr[0]
		fieldSuffix := fieldSplitArr[len(fieldSplitArr)-1]

		dataKeyOfColumn := `dataKey: '` + field + `',`
		titleOfColumn := `title: t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `'),`
		keyOfColumn := `key: '` + field + `',`
		alignOfColumn := `align: 'center',`
		widthOfColumn := `width: 150,`
		sortableOfColumn := ``
		hiddenOfColumn := ``
		cellRendererOfColumn := ``

		if garray.NewStrArrayFrom([]string{`DeletedAt`, `DeleteAt`, `DeletedTime`, `DeleteTime`}).Contains(fieldCaseCamel) {
			continue
		} else if garray.NewStrArrayFrom([]string{`UpdatedAt`, `UpdateAt`, `UpdatedTime`, `UpdateTime`}).Contains(fieldCaseCamel) {
			titleOfColumn = `title: t('common.name.updatedAt'),`
			sortableOfColumn = `sortable: true,`
		} else if garray.NewStrArrayFrom([]string{`CreatedAt`, `CreateAt`, `CreatedTime`, `CreateTime`}).Contains(fieldCaseCamel) {
			titleOfColumn = `title: t('common.name.createdAt'),`
			sortableOfColumn = `sortable: true,`
		} else if column[`Key`].String() == `PRI` && column[`Extra`].String() == `auto_increment` { //主键
			continue
		} else if fieldCaseCamel == `IdPath` && (gstr.Pos(column[`Type`].String(), `varchar`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1) && tpl.PidHandle.IsCoexist { //idPath|id_path
			hiddenOfColumn = `hidden: true,`
		} else if garray.NewStrArrayFrom([]string{`status`, `type`, `method`, `pos`, `position`, `gender`}).Contains(fieldSuffix) && ((gstr.Pos(column[`Type`].String(), `int`) != -1 && gstr.Pos(column[`Type`].String(), `point`) == -1) || gstr.Pos(column[`Type`].String(), `char`) != -1) { //status,type,method,pos,position,gender等后缀
			widthOfColumn = `width: 100,`
			cellRendererOfColumn = `cellRenderer: (props: any): any => {
                let tagType = tm('config.const.tagType') as string[]
                let obj = tm('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.status.` + field + `') as { value: any, label: string }[]
                let index = obj.findIndex((item) => { return item.value == props.rowData.` + field + ` })
                return <el-tag type={tagType[index % tagType.length]}>{obj[index]?.label}</el-tag>
            },`
		} else if (garray.NewStrArrayFrom([]string{`icon`, `cover`, `avatar`, `img`, `image`}).Contains(fieldSuffix) || gstr.SubStr(fieldCaseCamelOfRemove, -7) == `ImgList` || gstr.SubStr(fieldCaseCamelOfRemove, -6) == `ImgArr` || gstr.SubStr(fieldCaseCamelOfRemove, -9) == `ImageList` || gstr.SubStr(fieldCaseCamelOfRemove, -8) == `ImageArr`) && (gstr.Pos(column[`Type`].String(), `varchar`) != -1 || gstr.Pos(column[`Type`].String(), `json`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1) { //icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀
			widthOfColumn = `width: 100,`
			cellRendererOfColumn = `cellRenderer: (props: any): any => {
                if (!props.rowData.` + field + `) {
                    return
                }`
			if gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
				cellRendererOfColumn += `
                const imageList = [props.rowData.` + field + `]`
			} else {
				cellRendererOfColumn += `
                let imageList: string[]
                if (Array.isArray(props.rowData.` + field + `)) {
                    imageList = props.rowData.` + field + `
                } else {
                    imageList = JSON.parse(props.rowData.` + field + `)
                }`
			}
			cellRendererOfColumn += `
                return [
                    <el-scrollbar wrap-style="display: flex; align-items: center;" view-style="margin: auto;">
                        {imageList.map((item) => {
                            //修改宽高时，可同时修改table属性row-height增加行高，则不会显示滚动条
                            return <el-image style="width: 45px;" src={item} lazy={true} hide-on-click-modal={true} preview-teleported={true} preview-src-list={imageList} />
                        })}
                    </el-scrollbar>
                ]
            },`
		} else if (garray.NewStrArrayFrom([]string{`video`}).Contains(fieldSuffix) || gstr.SubStr(fieldCaseCamelOfRemove, -9) == `VideoList` || gstr.SubStr(fieldCaseCamelOfRemove, -8) == `VideoArr`) && (gstr.Pos(column[`Type`].String(), `varchar`) != -1 || gstr.Pos(column[`Type`].String(), `json`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1) { //video,video_list,videoList,video_arr,videoArr等后缀
			if tableRowHeight < 100 {
				tableRowHeight = 100
			}
			cellRendererOfColumn = `cellRenderer: (props: any): any => {
                if (!props.rowData.` + field + `) {
                    return
                }`
			if gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
				cellRendererOfColumn += `
                const videoList = [props.rowData.` + field + `]`
			} else {
				cellRendererOfColumn += `
                let videoList: string[]
                if (Array.isArray(props.rowData.` + field + `)) {
                    videoList = props.rowData.` + field + `
                } else {
                    videoList = JSON.parse(props.rowData.` + field + `)
                }`
			}
			cellRendererOfColumn += `
                return [
                    <el-scrollbar wrap-style="display: flex; align-items: center;" view-style="margin: auto;">
                        {videoList.map((item) => {
                            //修改宽高时，可同时修改table属性row-height增加行高，则不会显示滚动条
                            return <video style="width: 120px; height: 80px;" preload="none" controls={true} src={item} />
                        })}
                    </el-scrollbar>,
                ]
            },`
		} else if garray.NewStrArrayFrom([]string{`list`, `arr`}).Contains(fieldSuffix) && (gstr.Pos(column[`Type`].String(), `json`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1) { //list,arr等后缀
			widthOfColumn = `width: 100,`
			cellRendererOfColumn = `cellRenderer: (props: any): any => {
                if (!props.rowData.` + field + `) {
                    return
                }`
			cellRendererOfColumn += `
                let arrList: any[]
                if (Array.isArray(props.rowData.` + field + `)) {
                    arrList = props.rowData.` + field + `
                } else {
                    arrList = JSON.parse(props.rowData.` + field + `)
                }
                let tagType = tm('config.const.tagType') as string[]
                return [
                    <el-scrollbar wrap-style="display: flex; align-items: center;" view-style="margin: auto;">
                        {arrList.map((item, index) => {
                            return [
                                <el-tag style="margin: auto 5px 5px auto;" type={tagType[index % tagType.length]}>
                                    {item}
                                </el-tag>,
                            ]
                        })}
                    </el-scrollbar>,
                ]
            },`
		} else if garray.NewStrArrayFrom([]string{`remark`, `desc`, `msg`, `message`, `intro`, `content`}).Contains(fieldSuffix) && (gstr.Pos(column[`Type`].String(), `varchar`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1) { //remark,desc,msg,message,intro,content后缀
			hiddenOfColumn = `hidden: true,`
		} else if gstr.Pos(column[`Type`].String(), `varchar`) != -1 { //varchar类型
			if garray.NewStrArrayFrom([]string{`url`, `link`}).Contains(fieldSuffix) { //url,link后缀
				widthOfColumn = `width: 200,`
			}
		} else if gstr.Pos(column[`Type`].String(), `char`) != -1 { //char类型
			if garray.NewStrArrayFrom([]string{`password`, `passwd`}).Contains(fieldSuffix) && column[`Type`].String() == `char(32)` { //password,passwd后缀
				continue
			} else if garray.NewStrArrayFrom([]string{`salt`}).Contains(fieldSuffix) && tpl.PasswordHandleMap[MyGenPasswordHandleMapKey(field)].IsCoexist { //salt后缀
				continue
			}
		} else if gstr.Pos(column[`Type`].String(), `int`) != -1 && gstr.Pos(column[`Type`].String(), `point`) == -1 { //int等类型
			if field == `pid` { //pid
				dataKeyOfColumn = `dataKey: 'p` + gstr.CaseCamel(tpl.LabelHandle.LabelField) + `',`
			} else if field == `level` && tpl.PidHandle.IsCoexist { //level
				widthOfColumn = `width: 100,`
				sortableOfColumn = `sortable: true,`
			} else if garray.NewStrArrayFrom([]string{`sort`, `weight`}).Contains(fieldSuffix) { //sort,weight等后缀
				widthOfColumn = `width: 100,`
				sortableOfColumn = `sortable: true,`
				if option.IsUpdate {
					cellRendererOfColumn = `cellRenderer: (props: any): any => {
                if (props.rowData.edit` + gstr.CaseCamel(field) + `) {
                    let currentRef: any
                    let currentVal = props.rowData.` + field + `
                    return [
                        <el-input-number
                            ref={(el: any) => {
                                currentRef = el
                                el?.focus()
                            }}
                            model-value={currentVal}
                            placeholder={t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.tip.` + field + `')}
                            precision={0}
                            min={0}
                            max={100}
                            step={1}
                            step-strictly={true}
                            controls={false} //控制按钮会导致诸多问题。如：焦点丢失；` + field + `是0或100时，只一个按钮可点击
                            controls-position="right"
                            onChange={(val: number) => (currentVal = val)}
                            onBlur={() => {
                                props.rowData.edit` + gstr.CaseCamel(field) + ` = false
                                if ((currentVal || currentVal === 0) && currentVal != props.rowData.` + field + `) {
                                    handleUpdate({
                                        idArr: [props.rowData.id],
                                        ` + field + `: currentVal,
                                    })
                                        .then((res) => {
                                            props.rowData.` + field + ` = currentVal
                                        })
                                        .catch((error) => {})
                                }
                            }}
                            onKeydown={(event: any) => {
                                switch (event.keyCode) {
                                    // case 27:    //Esc键：Escape
                                    // case 32:    //空格键：" "
                                    case 13: //Enter键：Enter
                                        // props.rowData.edit` + gstr.CaseCamel(field) + ` = false    //也会触发onBlur事件
                                        currentRef?.blur()
                                        break
                                }
                            }}
                        />,
                    ]
                }
                return [
                    <div class="inline-edit" onClick={() => (props.rowData.edit` + gstr.CaseCamel(field) + ` = true)}>
                        {props.rowData.` + field + `}
                    </div>,
                ]
            },`
				}
			} else if garray.NewStrArrayFrom([]string{`id`}).Contains(fieldSuffix) { //id后缀
				if tpl.RelTableMap[field].IsExistRelTableDao && !tpl.RelTableMap[field].IsRedundRelNameField {
					dataKeyOfColumn = `dataKey: '` + tpl.RelTableMap[field].RelTableField + tpl.RelTableMap[field].RelSuffix + `',`
				}
			} else if garray.NewStrArrayFrom([]string{`is`}).Contains(fieldPrefix) { //is_前缀
				widthOfColumn = `width: 100,`
				cellRendererOfColumn = `cellRenderer: (props: any): any => {
                return [
                    <el-switch
                        model-value={props.rowData.` + field + `}
                        // disabled={true}
                        active-value={1}
                        inactive-value={0}
                        inline-prompt={true}
                        active-text={t('common.yes')}
                        inactive-text={t('common.no')}
                        style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success);"`
				if option.IsUpdate {
					cellRendererOfColumn += `
                        onChange={(val: number) => {
                            handleUpdate({
                                idArr: [props.rowData.id],
                                ` + field + `: val,
                            })
                                .then((res) => {
                                    props.rowData.` + field + ` = val
                                })
                                .catch((error) => {})
                        }}`
				}
				cellRendererOfColumn += `
                    />,
                ]
            },`
			}
		} else if gstr.Pos(column[`Type`].String(), `timestamp`) != -1 || gstr.Pos(column[`Type`].String(), `date`) != -1 { //timestamp或datetime或date类型
			sortableOfColumn = `sortable: true,`
			if gstr.Pos(column[`Type`].String(), `date`) != -1 && gstr.Pos(column[`Type`].String(), `datetime`) == -1 {
				widthOfColumn = `width: 100,`
			}
		} else if gstr.Pos(column[`Type`].String(), `json`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1 { //json类型 //text类型
			widthOfColumn = `width: 200,`
			hiddenOfColumn = `hidden: true,`
		}

		viewListColumn += `
        {
            ` + dataKeyOfColumn + `
            ` + titleOfColumn + `
            ` + keyOfColumn + `
            ` + alignOfColumn + `
            ` + widthOfColumn
		if sortableOfColumn != `` {
			viewListColumn += `
            ` + sortableOfColumn
		}
		if hiddenOfColumn != `` {
			viewListColumn += `
            ` + hiddenOfColumn
		}
		if cellRendererOfColumn != `` {
			viewListColumn += `
            ` + cellRendererOfColumn
		}
		viewListColumn += `
        },`
	}

	tplView := `<script setup lang="tsx">
const { t, tm } = useI18n()

const table = reactive({
    columns: [
        {
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
                    //阻止冒泡
                    <div class="id-checkbox" onClick={(event: any) => event.stopPropagation()}>
                        <el-checkbox
                            model-value={table.data.length ? allChecked : false}
                            indeterminate={someChecked && !allChecked}
                            onChange={(val: boolean) => {
                                table.data.forEach((item: any) => {
                                    item.checked = val
                                })
                            }}
                        />
                    </div>,
                    <div>{t('common.name.id')}</div>,
                ]
            },
            cellRenderer: (props: any): any => {
                return [<el-checkbox class="id-checkbox" model-value={props.rowData.checked} onChange={(val: boolean) => (props.rowData.checked = val)} />, <div>{props.rowData.id}</div>]
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
                    <el-button type="primary" size="small" onClick={() => handleEditCopy(props.rowData.id)}>
                        <autoicon-ep-edit />
                        {t('common.edit')}
                    </el-button>,`
		}
		if option.IsDelete {
			tplView += `
                    <el-button type="danger" size="small" onClick={() => handleDelete(props.rowData.id)}>
                        <autoicon-ep-delete />
                        {t('common.delete')}
                    </el-button>,`
		}
		if option.IsCreate {
			tplView += `
                    <el-button type="warning" size="small" onClick={() => handleEditCopy(props.rowData.id, 'copy')}>
                        <autoicon-ep-document-copy />
                        {t('common.copy')}
                    </el-button>,`
		}
		tplView += `
                ]
            },
        },`
	}
	tplView += `
    ] as any,
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

const saveCommon = inject('saveCommon') as { visible: boolean; title: string; data: { [propName: string]: any } }`
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
    const idArr: number[] = []
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
    request(t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseCamelLower + `/` + tpl.TableNameCaseCamelLower + `/info', { id: id })
        .then((res) => {
            saveCommon.data = { ...res.data.info }
            switch (type) {
                case 'edit':
                    saveCommon.data.idArr = [saveCommon.data.id]
                    delete saveCommon.data.id
                    saveCommon.title = t('common.edit')
                    break
                case 'copy':
                    delete saveCommon.data.id
                    saveCommon.title = t('common.copy')
                    break
            }
            saveCommon.visible = true
        })
        .catch(() => {})
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
    })
        .then(() => {
            request(t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseCamelLower + `/` + tpl.TableNameCaseCamelLower + `/del', { idArr: idArr }, true)
                .then((res) => {
                    getList()
                })
                .catch(() => {})
        })
        .catch(() => {})
}`
	}
	if option.IsUpdate {
		tplView += `
//更新
const handleUpdate = async (param: { idArr: number[]; [propName: string]: any }) => {
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
    },
})

const queryCommon = inject('queryCommon') as { data: { [propName: string]: any } }
//列表
const getList = async (resetPage: boolean = false) => {
    if (resetPage) {
        pagination.page = 1
    }
    const param = {
        field: [],
        filter: removeEmptyOfObj(queryCommon.data, true, true),
        sort: table.sort.key + ' ' + table.sort.order,
        page: pagination.page,
        limit: pagination.size,
    }
    table.loading = true
    try {
        const res = await request(t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseCamelLower + `/` + tpl.TableNameCaseCamelLower + `/list', param)
        table.data = res.data.list?.length ? res.data.list : []
        pagination.total = res.data.count
    } catch (error) {}
    table.loading = false
}
getList()

//暴露组件接口给父组件
defineExpose({
    getList,
})
</script>

<template>
    <el-row class="main-table-tool">
        <el-col :span="16">
            <el-space :size="10" style="height: 100%; margin-left: 10px">`
	if option.IsCreate {
		tplView += `
                <el-button type="primary" @click="handleAdd"> <autoicon-ep-edit-pen />{{ t('common.add') }} </el-button>`
	}
	if option.IsDelete {
		tplView += `
                <el-button type="danger" @click="handleBatchDelete"> <autoicon-ep-delete-filled />{{ t('common.batchDelete') }} </el-button>`
	}
	tplView += `
            </el-space>
        </el-col>
        <el-col :span="8" style="text-align: right">
            <el-space :size="10" style="height: 100%">
                <my-export-button i18nPrefix="` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `" :headerList="table.columns" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseCamelLower + `/` + tpl.TableNameCaseCamelLower + `/list', param: { filter: queryCommon.data, sort: table.sort.key + ' ' + table.sort.order } }" />
                <el-dropdown max-height="300" :hide-on-click="false">
                    <el-button type="info" :circle="true">
                        <autoicon-ep-hide />
                    </el-button>
                    <template #dropdown>
                        <el-dropdown-menu>
                            <el-dropdown-item v-for="(item, index) in table.columns" :key="index">
                                <el-checkbox v-model="item.hidden">
                                    {{ item.title }}
                                </el-checkbox>
                            </el-dropdown-item>
                        </el-dropdown-menu>
                    </template>
                </el-dropdown>
            </el-space>
        </el-col>
    </el-row>

    <el-main>
        <el-auto-resizer>
            <template #default="{ height, width }">
                <el-table-v2 class="main-table" :columns="table.columns" :data="table.data" :sort-by="table.sort" @column-sort="table.handleSort" :width="width" :height="height" :fixed="true" :row-height="` + gconv.String(tableRowHeight) + `">
                    <template v-if="table.loading" #overlay>
                        <el-icon class="is-loading" color="var(--el-color-primary)" :size="25">
                            <autoicon-ep-loading />
                        </el-icon>
                    </template>
                </el-table-v2>
            </template>
        </el-auto-resizer>
    </el-main>

    <el-row class="main-table-pagination">
        <el-col :span="24">
            <el-pagination
                :total="pagination.total"
                v-model:currentPage="pagination.page"
                v-model:page-size="pagination.size"
                @size-change="pagination.sizeChange"
                @current-change="pagination.pageChange"
                :page-sizes="pagination.sizeList"
                :layout="pagination.layout"
                :background="true"
            />
        </el-col>
    </el-row>
</template>
`

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
		fieldCaseSnakeOfRemove := gstr.Split(fieldCaseSnake, `_of_`)[0]
		fieldCaseCamelOfRemove := gstr.CaseCamel(fieldCaseSnakeOfRemove)
		fieldSplitArr := gstr.Split(fieldCaseSnakeOfRemove, `_`)
		fieldPrefix := fieldSplitArr[0]
		fieldSuffix := fieldSplitArr[len(fieldSplitArr)-1]
		resultStr, _ := gregex.MatchString(`.*\((\d*)\)`, column[`Type`].String())
		/* resultFloat, _ := gregex.MatchString(`.*\((\d*),(\d*)\)`, column[`Type`].String())
		if len(resultFloat) < 3 {
			resultFloat = []string{``, `10`, `2`}
		} */

		if garray.NewStrArrayFrom([]string{`DeletedAt`, `DeleteAt`, `DeletedTime`, `DeleteTime`}).Contains(fieldCaseCamel) {
		} else if garray.NewStrArrayFrom([]string{`UpdatedAt`, `UpdateAt`, `UpdatedTime`, `UpdateTime`}).Contains(fieldCaseCamel) {
		} else if garray.NewStrArrayFrom([]string{`CreatedAt`, `CreateAt`, `CreatedTime`, `CreateTime`}).Contains(fieldCaseCamel) {
			viewQueryDataInit += `
    timeRange: (() => {
        return undefined
        /* const date = new Date()
        return [
            new Date(date.getFullYear(), date.getMonth(), date.getDate(), 0, 0, 0),
            new Date(date.getFullYear(), date.getMonth(), date.getDate(), 23, 59, 59),
        ] */
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
        <el-form-item prop="timeRange">
            <el-date-picker v-model="queryCommon.data.timeRange" type="datetimerange" range-separator="-" :default-time="[new Date(2000, 0, 1, 0, 0, 0), new Date(2000, 0, 1, 23, 59, 59)]" :start-placeholder="t('common.name.timeRangeStart')" :end-placeholder="t('common.name.timeRangeEnd')" />
        </el-form-item>`
		} else if column[`Key`].String() == `PRI` && column[`Extra`].String() == `auto_increment` { //主键
		} else if fieldCaseCamel == `IdPath` && (gstr.Pos(column[`Type`].String(), `varchar`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1) && tpl.PidHandle.IsCoexist { //idPath|id_path
		} else if garray.NewStrArrayFrom([]string{`status`, `type`, `method`, `pos`, `position`, `gender`}).Contains(fieldSuffix) && ((gstr.Pos(column[`Type`].String(), `int`) != -1 && gstr.Pos(column[`Type`].String(), `point`) == -1) || gstr.Pos(column[`Type`].String(), `char`) != -1) { //status,type,method,pos,position,gender等后缀
			viewQueryField += `
        <el-form-item prop="` + field + `" style="width: 120px">
            <el-select-v2 v-model="queryCommon.data.` + field + `" :options="tm('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.status.` + field + `')" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :clearable="true" />
        </el-form-item>`
		} else if (garray.NewStrArrayFrom([]string{`icon`, `cover`, `avatar`, `img`, `image`}).Contains(fieldSuffix) || gstr.SubStr(fieldCaseCamelOfRemove, -7) == `ImgList` || gstr.SubStr(fieldCaseCamelOfRemove, -6) == `ImgArr` || gstr.SubStr(fieldCaseCamelOfRemove, -9) == `ImageList` || gstr.SubStr(fieldCaseCamelOfRemove, -8) == `ImageArr` || garray.NewStrArrayFrom([]string{`video`}).Contains(fieldSuffix) || gstr.SubStr(fieldCaseCamelOfRemove, -9) == `VideoList` || gstr.SubStr(fieldCaseCamelOfRemove, -8) == `VideoArr`) && (gstr.Pos(column[`Type`].String(), `varchar`) != -1 || gstr.Pos(column[`Type`].String(), `json`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1) { //icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀 //video,video_list,videoList,video_arr,videoArr等后缀
		} else if garray.NewStrArrayFrom([]string{`list`, `arr`}).Contains(fieldSuffix) && (gstr.Pos(column[`Type`].String(), `json`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1) { //list,arr等后缀
		} else if garray.NewStrArrayFrom([]string{`remark`, `desc`, `msg`, `message`, `intro`, `content`}).Contains(fieldSuffix) && (gstr.Pos(column[`Type`].String(), `varchar`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1) { //remark,desc,msg,message,intro,content后缀
		} else if gstr.Pos(column[`Type`].String(), `varchar`) != -1 { //varchar类型
			viewQueryField += `
        <el-form-item prop="` + field + `">
            <el-input v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" maxlength="` + resultStr[1] + `" :clearable="true" />
        </el-form-item>`
		} else if gstr.Pos(column[`Type`].String(), `char`) != -1 { //char类型
			if garray.NewStrArrayFrom([]string{`password`, `passwd`}).Contains(fieldSuffix) && column[`Type`].String() == `char(32)` { //password,passwd后缀
			} else if garray.NewStrArrayFrom([]string{`salt`}).Contains(fieldSuffix) && tpl.PasswordHandleMap[MyGenPasswordHandleMapKey(field)].IsCoexist { //salt后缀
			} else {
				viewQueryField += `
        <el-form-item prop="` + field + `">
            <el-input v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" minlength="` + resultStr[1] + `" maxlength="` + resultStr[1] + `" :clearable="true" />
        </el-form-item>`
			}
		} else if gstr.Pos(column[`Type`].String(), `int`) != -1 && gstr.Pos(column[`Type`].String(), `point`) == -1 { //int等类型
			if field == `pid` { //pid
				viewQueryField += `
        <el-form-item prop="` + field + `">
            <my-cascader v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseCamelLower + `/` + tpl.TableNameCaseCamelLower + `/tree' }" :defaultOptions="[{ id: 0, label: t('common.name.allTopLevel') }]" :props="{ checkStrictly: true, emitPath: false }" />
        </el-form-item>`
			} else if field == `level` && tpl.PidHandle.IsCoexist { //level
				viewQueryField += `
        <el-form-item prop="` + field + `">
            <el-input-number v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :min="1" :controls="false" />
        </el-form-item>`
			} else if garray.NewStrArrayFrom([]string{`sort`, `weight`}).Contains(fieldSuffix) { //sort,weight等后缀
			} else if garray.NewStrArrayFrom([]string{`id`}).Contains(fieldSuffix) { //id后缀
				apiUrl := tpl.ModuleDirCaseCamelLower + `/` + gstr.CaseCamelLower(gstr.SubStr(field, 0, -2))
				if tpl.RelTableMap[field].IsExistRelTableDao {
					relTable := tpl.RelTableMap[field]
					apiUrl = relTable.RelDaoDirCaseCamelLower + `/` + relTable.RelTableNameCaseCamelLower
				}
				viewQueryField += `
        <el-form-item prop="` + field + `">
            <my-select v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + apiUrl + `/list' }" />
        </el-form-item>`
			} else if garray.NewStrArrayFrom([]string{`is`}).Contains(fieldPrefix) { //is_前缀
				viewQueryField += `
        <el-form-item prop="` + field + `" style="width: 120px">
            <el-select-v2 v-model="queryCommon.data.` + field + `" :options="tm('common.status.whether')" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :clearable="true" />
        </el-form-item>`
			} else { //默认处理（int等类型）
				/* if gstr.Pos(column[`Type`].String(), `unsigned`) != -1 {
				               viewQueryField += `
				   <el-form-item prop="` + field + `">
				       <el-input-number v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :min="0" :controls="false" />
				   </el-form-item>`
				           } else {
				               viewQueryField += `
				   <el-form-item prop="` + field + `">
				       <el-input-number v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :controls="false" />
				   </el-form-item>`
				           } */
			}
		} else if gstr.Pos(column[`Type`].String(), `decimal`) != -1 || gstr.Pos(column[`Type`].String(), `double`) != -1 || gstr.Pos(column[`Type`].String(), `float`) != -1 { //float类型
			/* if gstr.Pos(column[`Type`].String(), `unsigned`) != -1 {
			           viewQueryField += `
			   <el-form-item prop="` + field + `">
			       <el-input-number v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :min="0" :precision="` + resultFloat[2] + `" :controls="false" />
			   </el-form-item>`
			       } else {
			           viewQueryField += `
			   <el-form-item prop="` + field + `">
			       <el-input-number v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :precision="` + resultFloat[2] + `" :controls="false" />
			   </el-form-item>`
			       } */
		} else if gstr.Pos(column[`Type`].String(), `timestamp`) != -1 || gstr.Pos(column[`Type`].String(), `date`) != -1 { //timestamp或datetime或date类型
			typeDatePicker := ``
			formatDatePicker := `YYYY-MM-DD HH:mm:ss`
			defaultTimeDatePicker := ``
			if gstr.Pos(column[`Type`].String(), `date`) != -1 && gstr.Pos(column[`Type`].String(), `datetime`) == -1 {
				typeDatePicker = `date`
				formatDatePicker = `YYYY-MM-DD`
			}

			if garray.NewStrArrayFrom([]string{`start`}).Contains(fieldPrefix) { //start_前缀
				typeDatePicker = `datetime`
				if formatDatePicker == `YYYY-MM-DD HH:mm:ss` {
					defaultTimeDatePicker = ` :default-time="new Date(2000, 0, 1, 0, 0, 0)"`
				}
			} else if garray.NewStrArrayFrom([]string{`end`}).Contains(fieldPrefix) { //end_前缀
				typeDatePicker = `datetime`
				if formatDatePicker == `YYYY-MM-DD HH:mm:ss` {
					defaultTimeDatePicker = ` :default-time="new Date(2000, 0, 1, 23, 59, 59)"`
				}
			}

			if typeDatePicker != `` {
				viewQueryField += `
        <el-form-item prop="` + field + `">
            <el-date-picker v-model="queryCommon.data.` + field + `" type="` + typeDatePicker + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" format="` + formatDatePicker + `" value-format="` + formatDatePicker + `"` + defaultTimeDatePicker + ` />
        </el-form-item>`
			}
		} else if gstr.Pos(column[`Type`].String(), `json`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1 { //json类型 //text类型
		} else { //默认处理
			viewQueryField += `
        <el-form-item prop="` + field + `">
            <el-input v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :clearable="true" />
        </el-form-item>`
		}
	}

	tplView := `<script setup lang="tsx">
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
    },
})
</script>

<template>
    <el-form class="query-form" :ref="(el: any) => queryForm.ref = el" :model="queryCommon.data" :inline="true" @keyup.enter="queryForm.submit">
        <el-form-item prop="id">
            <el-input-number v-model="queryCommon.data.id" :placeholder="t('common.name.id')" :min="1" :controls="false" />
        </el-form-item>` + viewQueryField + `
        <el-form-item>
            <el-button type="primary" @click="queryForm.submit" :loading="queryForm.loading"> <autoicon-ep-search />{{ t('common.query') }} </el-button>
            <el-button type="info" @click="queryForm.reset"> <autoicon-ep-circle-close />{{ t('common.reset') }} </el-button>
        </el-form-item>
    </el-form>
</template>
`

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

	viewSaveImport := ``
	viewSaveParamHandle := ``
	viewSaveDataInit := ``
	viewSaveRule := ``
	viewSaveField := ``
	viewFieldHandle := ``
	for _, column := range tpl.TableColumnList {
		field := column[`Field`].String()
		fieldCaseCamel := gstr.CaseCamel(field)
		fieldCaseSnake := gstr.CaseSnakeFirstUpper(field)
		fieldCaseSnakeOfRemove := gstr.Split(fieldCaseSnake, `_of_`)[0]
		fieldCaseCamelOfRemove := gstr.CaseCamel(fieldCaseSnakeOfRemove)
		fieldSplitArr := gstr.Split(fieldCaseSnakeOfRemove, `_`)
		fieldPrefix := fieldSplitArr[0]
		fieldSuffix := fieldSplitArr[len(fieldSplitArr)-1]
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
		} else if garray.NewStrArrayFrom([]string{`UpdatedAt`, `UpdateAt`, `UpdatedTime`, `UpdateTime`}).Contains(fieldCaseCamel) {
		} else if garray.NewStrArrayFrom([]string{`CreatedAt`, `CreateAt`, `CreatedTime`, `CreateTime`}).Contains(fieldCaseCamel) {
		} else if column[`Key`].String() == `PRI` && column[`Extra`].String() == `auto_increment` { //主键
		} else if fieldCaseCamel == `IdPath` && (gstr.Pos(column[`Type`].String(), `varchar`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1) && tpl.PidHandle.IsCoexist { //idPath|id_path
		} else if garray.NewStrArrayFrom([]string{`status`, `type`, `method`, `pos`, `position`, `gender`}).Contains(fieldSuffix) && ((gstr.Pos(column[`Type`].String(), `int`) != -1 && gstr.Pos(column[`Type`].String(), `point`) == -1) || gstr.Pos(column[`Type`].String(), `char`) != -1) { //status,type,method,pos,position,gender等后缀
			isStr := true
			if gstr.Pos(column[`Type`].String(), `int`) != -1 && gstr.Pos(column[`Type`].String(), `point`) == -1 {
				isStr = false
			}
			statusList := MyGenStatusList(comment, isStr)
			defaultVal := column[`Default`].String()
			if defaultVal == `` {
				defaultVal = statusList[0][0]
			}
			if isStr {
				viewSaveDataInit += `
        ` + field + `: '` + defaultVal + `',`
			} else {
				viewSaveDataInit += `
        ` + field + `: ` + defaultVal + `,`
			}
			viewSaveRule += `
        ` + field + `: [
            { type: 'enum', enum: (tm('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.status.` + field + `') as any).map((item: any) => item.value), trigger: 'change', message: t('validation.select') },
        ],`
			viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">`
			//超过5个状态用select组件，小于5个用radio组件
			if len(statusList) > 5 {
				viewSaveField += `
                    <el-select-v2 v-model="saveForm.data.` + field + `" :options="tm('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.status.` + field + `')" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :clearable="false" />`
			} else {
				viewSaveField += `
                    <el-radio-group v-model="saveForm.data.` + field + `">
                        <el-radio v-for="(item, index) in (tm('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.status.` + field + `') as any)" :key="index" :label="item.value">
                            {{ item.label }}
                        </el-radio>
                    </el-radio-group>`
			}
			viewSaveField += `
                </el-form-item>`
		} else if (garray.NewStrArrayFrom([]string{`icon`, `cover`, `avatar`, `img`, `image`}).Contains(fieldSuffix) || gstr.SubStr(fieldCaseCamelOfRemove, -7) == `ImgList` || gstr.SubStr(fieldCaseCamelOfRemove, -6) == `ImgArr` || gstr.SubStr(fieldCaseCamelOfRemove, -9) == `ImageList` || gstr.SubStr(fieldCaseCamelOfRemove, -8) == `ImageArr`) && (gstr.Pos(column[`Type`].String(), `varchar`) != -1 || gstr.Pos(column[`Type`].String(), `json`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1) { //icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀
			multipleStr := ``
			if gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
				viewSaveRule += `
        ` + field + `: [
            { type: 'string', max: ` + resultStr[1] + `, trigger: 'blur', message: t('validation.max.string', { max: ` + resultStr[1] + ` }) },
            { type: 'url', trigger: 'change', message: t('validation.upload') },
        ],`
			} else {
				multipleStr = ` :multiple="true"`
				requiredStr := ``
				if column[`Null`].String() == `NO` {
					requiredStr = ` required: true,`
				}
				viewSaveRule += `
        ` + field + `: [
            { type: 'array',` + requiredStr + ` trigger: 'change', message: t('validation.upload'), defaultField: { type: 'url', message: t('validation.url') } },
            // { type: 'array',` + requiredStr + ` max: 10, trigger: 'change', message: t('validation.max.upload', { max: 10 }), defaultField: { type: 'url', message: t('validation.url') } },
        ],`
			}
			viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
                    <my-upload v-model="saveForm.data.` + field + `" accept="image/*"` + multipleStr + ` />
                </el-form-item>`
		} else if (garray.NewStrArrayFrom([]string{`video`}).Contains(fieldSuffix) || gstr.SubStr(fieldCaseCamelOfRemove, -9) == `VideoList` || gstr.SubStr(fieldCaseCamelOfRemove, -8) == `VideoArr`) && (gstr.Pos(column[`Type`].String(), `varchar`) != -1 || gstr.Pos(column[`Type`].String(), `json`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1) { //video,video_list,videoList,video_arr,videoArr等后缀
			multipleStr := ``
			if gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
				viewSaveRule += `
        ` + field + `: [
            { type: 'string', max: ` + resultStr[1] + `, trigger: 'blur', message: t('validation.max.string', { max: ` + resultStr[1] + ` }) },
            { type: 'url', trigger: 'change', message: t('validation.upload') },
        ],`
			} else {
				multipleStr = ` :multiple="true"`
				requiredStr := ``
				if column[`Null`].String() == `NO` {
					requiredStr = ` required: true,`
				}
				viewSaveRule += `
        ` + field + `: [
            { type: 'array',` + requiredStr + ` trigger: 'change', message: t('validation.upload'), defaultField: { type: 'url', message: t('validation.url') } },
            // { type: 'array',` + requiredStr + ` max: 10, trigger: 'change', message: t('validation.max.upload', { max: 10 }), defaultField: { type: 'url', message: t('validation.url') } },
        ],`
			}
			viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
                    <my-upload v-model="saveForm.data.` + field + `" accept="video/*" :isImage="false"` + multipleStr + ` />
                </el-form-item>`
		} else if garray.NewStrArrayFrom([]string{`list`, `arr`}).Contains(fieldSuffix) && (gstr.Pos(column[`Type`].String(), `json`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1) { //list,arr等后缀
			viewSaveDataInit += `
        ` + field + `: [],`
			requiredStr := ``
			if column[`Null`].String() == `NO` {
				requiredStr = ` required: true,`
			}
			viewSaveRule += `
        ` + field + `: [
            { type: 'array',` + requiredStr + ` trigger: 'change', message: t('validation.required') },
            // { type: 'array',` + requiredStr + ` max: 10, trigger: 'change', message: t('validation.max.array', { max: 10 }), defaultField: { type: 'string', message: t('validation.input') } },
        ],`
			viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
                    <el-tag v-for="(item, index) in saveForm.data.` + field + `" :type="` + field + `Handle.tagType[index % ` + field + `Handle.tagType.length]" @close="` + field + `Handle.delValue(item)" :key="index" :closable="true" style="margin-right: 10px;">
                        {{ item }}
                    </el-tag>
                    <!-- <el-input-number v-if="` + field + `Handle.visible" :ref="(el: any) => ` + field + `Handle.ref = el" v-model="` + field + `Handle.value" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" @keyup.enter="` + field + `Handle.addValue" @blur="` + field + `Handle.addValue" size="small" style="width: 100px;" :controls="false" /> -->
                    <el-input v-if="` + field + `Handle.visible" :ref="(el: any) => ` + field + `Handle.ref = el" v-model="` + field + `Handle.value" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" @keyup.enter="` + field + `Handle.addValue" @blur="` + field + `Handle.addValue" size="small" style="width: 100px;" />
                    <el-button v-else type="primary" size="small" @click="` + field + `Handle.visibleChange">
                        <autoicon-ep-plus />{{ t('common.add') }}
                    </el-button>
                </el-form-item>`
			viewFieldHandle += `

const ` + field + `Handle = reactive({
    ref: null as any,
    visible: false,
    value: undefined,
    tagType: tm('config.const.tagType') as string[],
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
		} else if garray.NewStrArrayFrom([]string{`remark`, `desc`, `msg`, `message`, `intro`, `content`}).Contains(fieldSuffix) && (gstr.Pos(column[`Type`].String(), `varchar`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1) { //remark,desc,msg,message,intro,content后缀
			if gstr.Pos(column[`Type`].String(), `text`) != -1 {
				viewSaveRule += `
        ` + field + `: [
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],`
				viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
                    <my-editor v-model="saveForm.data.` + field + `" />
                </el-form-item>`
			} else {
				viewSaveRule += `
        ` + field + `: [
            { type: 'string', max: ` + resultStr[1] + `, trigger: 'blur', message: t('validation.max.string', { max: ` + resultStr[1] + ` }) },
        ],`
				viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
                    <el-input v-model="saveForm.data.` + field + `" type="textarea" :autosize="{ minRows: 3 }" maxlength="` + resultStr[1] + `" :show-word-limit="true" />
                </el-form-item>`
			}
		} else if gstr.Pos(column[`Type`].String(), `varchar`) != -1 { //varchar类型
			ruleStr := ``
			requiredStr := ``
			viewSaveFieldTip := ` />`
			if column[`Key`].String() == `UNI` {
				if column[`Null`].String() == `NO` {
					requiredStr = ` required: true,`
				}
				viewSaveFieldTip = ` style="max-width: 250px" />
                    <label>
                        <el-alert :title="t('common.tip.notDuplicate')" type="info" :show-icon="true" :closable="false" />
                    </label>`
			}
			if garray.NewStrArrayFrom([]string{`name`}).Contains(fieldSuffix) { //name后缀
				if gstr.SubStr(gstr.CaseCamel(tpl.PrimaryKey), 0, -2)+`Name` == fieldCaseCamel {
					requiredStr = ` required: true,`
				}
				ruleStr += `
            { pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') },`
			} else if garray.NewStrArrayFrom([]string{`code`}).Contains(fieldSuffix) { //code后缀
				ruleStr += `
            { pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') },`
			} else if garray.NewStrArrayFrom([]string{`phone`, `mobile`}).Contains(fieldSuffix) { //mobile,phone后缀
				ruleStr += `
            { pattern: /^1[3-9]\d{9}$/, trigger: 'blur', message: t('validation.phone') },`
			} else if garray.NewStrArrayFrom([]string{`url`, `link`}).Contains(fieldSuffix) { //url,link后缀
				ruleStr += `
            { type: 'url', trigger: 'change', message: t('validation.url') },`
			}
			viewSaveRule += `
        ` + field + `: [
            { type: 'string',` + requiredStr + ` max: ` + resultStr[1] + `, trigger: 'blur', message: t('validation.max.string', { max: ` + resultStr[1] + ` }) },` + ruleStr + `
        ],`
			viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
                    <el-input v-model="saveForm.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" maxlength="` + resultStr[1] + `" :show-word-limit="true" :clearable="true"` + viewSaveFieldTip + `
                </el-form-item>`
		} else if gstr.Pos(column[`Type`].String(), `char`) != -1 { //char类型
			if garray.NewStrArrayFrom([]string{`password`, `passwd`}).Contains(fieldSuffix) && column[`Type`].String() == `char(32)` { //password,passwd后缀
				viewSaveImportTmp := `
import md5 from 'js-md5'`
				if gstr.Pos(viewSaveImport, viewSaveImportTmp) == -1 {
					viewSaveImport += viewSaveImportTmp
				}
				viewSaveParamHandle += `
            param.` + field + ` ? param.` + field + ` = md5(param.` + field + `) : delete param.` + field

				viewSaveRule += `
        ` + field + `: [
            { type: 'string', required: computed((): boolean => { return saveForm.data.idArr?.length ? false : true; }), min: 6, max: 20, trigger: 'blur', message: t('validation.between.string', { min: 6, max: 20 }) },
        ],`
				viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
                    <el-input v-model="saveForm.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" minlength="6" maxlength="20" :show-word-limit="true" :clearable="true" :show-password="true" style="max-width: 250px" />
                    <label v-if="saveForm.data.idArr?.length">
                        <el-alert :title="t('common.tip.notRequired')" type="info" :show-icon="true" :closable="false" />
                    </label>
                </el-form-item>`
			} else if garray.NewStrArrayFrom([]string{`salt`}).Contains(fieldSuffix) && tpl.PasswordHandleMap[MyGenPasswordHandleMapKey(field)].IsCoexist { //salt后缀
			} else {
				ruleStr := ``
				requiredStr := ``
				viewSaveFieldTip := ` />`
				if column[`Key`].String() == `UNI` {
					if column[`Null`].String() == `NO` {
						requiredStr = ` required: true,`
					}
					viewSaveFieldTip = ` style="max-width: 250px" />
                    <label>
                        <el-alert :title="t('common.tip.notDuplicate')" type="info" :show-icon="true" :closable="false" />
                    </label>`
				}
				viewSaveRule += `
        ` + field + `: [
            { type: 'string',` + requiredStr + ` len: ` + resultStr[1] + `, trigger: 'blur', message: t('validation.size.string', { size: ` + resultStr[1] + ` }) },` + ruleStr + `
        ],`
				viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
                    <el-input v-model="saveForm.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" minlength="` + resultStr[1] + `" maxlength="` + resultStr[1] + `" :show-word-limit="true" :clearable="true"` + viewSaveFieldTip + `
                </el-form-item>`
			}
		} else if gstr.Pos(column[`Type`].String(), `int`) != -1 && gstr.Pos(column[`Type`].String(), `point`) == -1 { //int等类型
			if field == `pid` { //pid
				viewSaveParamHandle += `
            if (param.` + field + ` === undefined) {
                param.` + field + ` = 0
            }`
				viewSaveRule += `
        ` + field + `: [
            { type: 'integer', min: 0, trigger: 'change', message: t('validation.select') },
        ],`
				viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
                    <my-cascader v-model="saveForm.data.` + field + `" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseCamelLower + `/` + tpl.TableNameCaseCamelLower + `/tree', param: { filter: { excIdArr: saveForm.data.idArr } } }" :props="{ checkStrictly: true, emitPath: false }" />
                </el-form-item>`
			} else if field == `level` && tpl.PidHandle.IsCoexist { //level
			} else if garray.NewStrArrayFrom([]string{`sort`, `weight`}).Contains(fieldSuffix) { //sort,weight等后缀
				defaultVal := column[`Default`].Int()
				if defaultVal != 0 {
					viewSaveDataInit += `
        ` + field + `: ` + gconv.String(defaultVal) + `,`
				}
				viewSaveRule += `
        ` + field + `: [
            { type: 'integer', min: 0, max: 100, trigger: 'change', message: t('validation.between.number', { min: 0, max: 100 }) },
        ],`
				viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
                    <el-input-number v-model="saveForm.data.` + field + `" :precision="0" :min="0" :max="100" :step="1" :step-strictly="true" controls-position="right" :value-on-clear="` + gconv.String(defaultVal) + `" />
                    <label>
                        <el-alert :title="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.tip.` + field + `')" type="info" :show-icon="true" :closable="false" />
                    </label>
                </el-form-item>`
			} else if garray.NewStrArrayFrom([]string{`id`}).Contains(fieldSuffix) { //id后缀
				apiUrl := tpl.ModuleDirCaseCamelLower + `/` + gstr.CaseCamelLower(gstr.SubStr(field, 0, -2))
				if tpl.RelTableMap[field].IsExistRelTableDao {
					relTable := tpl.RelTableMap[field]
					apiUrl = relTable.RelDaoDirCaseCamelLower + `/` + relTable.RelTableNameCaseCamelLower
				}
				viewSaveParamHandle += `
            if (param.` + field + ` === undefined) {
                param.` + field + ` = 0
            }`
				viewSaveRule += `
        ` + field + `: [
            { type: 'integer', min: 1, trigger: 'change', message: t('validation.select') },
        ],`
				viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
                    <my-select v-model="saveForm.data.` + field + `" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + apiUrl + `/list' }" />
                </el-form-item>`
			} else if garray.NewStrArrayFrom([]string{`is`}).Contains(fieldPrefix) { //is_前缀
				defaultVal := column[`Default`].Int()
				if defaultVal != 0 {
					viewSaveDataInit += `
        ` + field + `: ` + gconv.String(defaultVal) + `,`
				}
				viewSaveRule += `
        ` + field + `: [
            { type: 'enum', enum: (tm('common.status.whether') as any).map((item: any) => item.value), trigger: 'change', message: t('validation.select') },
        ],`
				viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
                    <el-switch v-model="saveForm.data.` + field + `" :active-value="1" :inactive-value="0" :inline-prompt="true" :active-text="t('common.yes')" :inactive-text="t('common.no')" style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success);" />
                </el-form-item>`
			} else { //默认处理（int等类型）
				if gstr.Pos(column[`Type`].String(), `unsigned`) != -1 {
					defaultVal := column[`Default`].Uint()
					if defaultVal != 0 {
						viewSaveDataInit += `
        ` + field + `: ` + gconv.String(defaultVal) + `,`
					}
					viewSaveRule += `
        ` + field + `: [
            { type: 'integer', min: 0, trigger: 'change', message: t('validation.min.number', { min: 0 }) },
        ],`
					viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
                    <el-input-number v-model="saveForm.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :min="0" :controls="false" :value-on-clear="` + gconv.String(defaultVal) + `" />
                </el-form-item>`
				} else {
					defaultVal := column[`Default`].Int()
					if defaultVal != 0 {
						viewSaveDataInit += `
        ` + field + `: ` + gconv.String(defaultVal) + `,`
					}
					viewSaveRule += `
        ` + field + `: [
            { type: 'integer', trigger: 'change', message: t('validation.input') },
        ],`
					viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
                    <el-input-number v-model="saveForm.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :controls="false" :value-on-clear="` + gconv.String(defaultVal) + `" />
                </el-form-item>`
				}
			}
		} else if gstr.Pos(column[`Type`].String(), `decimal`) != -1 || gstr.Pos(column[`Type`].String(), `double`) != -1 || gstr.Pos(column[`Type`].String(), `float`) != -1 { //float类型
			defaultVal := column[`Default`].Float64()
			if defaultVal != 0 {
				viewSaveDataInit += `
        ` + field + `: ` + gconv.String(defaultVal) + `,`
			}
			if gstr.Pos(column[`Type`].String(), `unsigned`) != -1 {
				viewSaveRule += `
        ` + field + `: [
            { type: 'number'/* 'float' */, min: 0, trigger: 'change', message: t('validation.min.number', { min: 0 }) },    // 类型float值为0时验证不能通过
        ],`
				viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
                    <el-input-number v-model="saveForm.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :min="0" :precision="` + resultFloat[2] + `" :controls="false" :value-on-clear="` + gconv.String(defaultVal) + `" />
                </el-form-item>`
			} else {
				viewSaveRule += `
        ` + field + `: [
            { type: 'number'/* 'float' */, trigger: 'change', message: t('validation.input') },    // 类型float值为0时验证不能通过
        ],`
				viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
                    <el-input-number v-model="saveForm.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :precision="` + resultFloat[2] + `" :controls="false" :value-on-clear="` + gconv.String(defaultVal) + `" />
                </el-form-item>`
			}
		} else if gstr.Pos(column[`Type`].String(), `timestamp`) != -1 || gstr.Pos(column[`Type`].String(), `date`) != -1 { //timestamp或datetime或date类型
			typeDatePicker := `datetime`
			formatDatePicker := `YYYY-MM-DD HH:mm:ss`
			defaultTimeDatePicker := ``
			if gstr.Pos(column[`Type`].String(), `date`) != -1 && gstr.Pos(column[`Type`].String(), `datetime`) == -1 {
				typeDatePicker = `date`
				formatDatePicker = `YYYY-MM-DD`
			}
			requiredStr := ``
			if column[`Null`].String() == `NO` && column[`Default`].String() == `` {
				requiredStr = ` required: true,`
			}

			if garray.NewStrArrayFrom([]string{`start`}).Contains(fieldPrefix) && formatDatePicker == `YYYY-MM-DD HH:mm:ss` { //start_前缀
				defaultTimeDatePicker = ` :default-time="new Date(2000, 0, 1, 0, 0, 0)"`
			} else if garray.NewStrArrayFrom([]string{`end`}).Contains(fieldPrefix) && formatDatePicker == `YYYY-MM-DD HH:mm:ss` { //end_前缀
				defaultTimeDatePicker = ` :default-time="new Date(2000, 0, 1, 23, 59, 59)"`
			}

			viewSaveRule += `
        ` + field + `: [
            { type: 'string',` + requiredStr + ` trigger: 'change', message: t('validation.select') },
        ],`
			viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
                    <el-date-picker v-model="saveForm.data.` + field + `" type="` + typeDatePicker + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" format="` + formatDatePicker + `" value-format="` + formatDatePicker + `"` + defaultTimeDatePicker + ` />
                </el-form-item>`
		} else if gstr.Pos(column[`Type`].String(), `json`) != -1 { //json类型
			requiredStr := ``
			if column[`Null`].String() == `NO` {
				requiredStr = `
                required: true,`
			}
			viewSaveRule += `
        ` + field + `: [
            {
                type: 'object',` + requiredStr + `
                /* fields: {
                    xxxx: { type: 'string', required: true, message: 'xxxx' + t('validation.required') },
                    xxxx: { type: 'integer', required: true, min: 1, message: 'xxxx' + t('validation.min.number', { min: 1 }) },
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
                message: t('validation.json'),
            },
        ],`
			viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
                    <el-alert :title="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.tip.` + field + `')" type="info" :show-icon="true" :closable="false" />
                    <el-input v-model="saveForm.data.` + field + `" type="textarea" :autosize="{ minRows: 3 }" />
                </el-form-item>`
		} else if gstr.Pos(column[`Type`].String(), `text`) != -1 { //text类型
			viewSaveRule += `
        ` + field + `: [],`
			viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
                    <my-editor v-model="saveForm.data.` + field + `" />
                </el-form-item>`
		} else { //默认处理
			viewSaveRule += `
        ` + field + `: [],`
			viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" prop="` + field + `">
                    <el-input v-model="saveForm.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseCamelLowerReplace + `.` + tpl.TableNameCaseCamelLower + `.name.` + field + `')" :clearable="true" />
                </el-form-item>`
		}
	}

	tplView := `<script setup lang="tsx">` + viewSaveImport + `
const { t, tm } = useI18n()

const saveCommon = inject('saveCommon') as { visible: boolean; title: string; data: { [propName: string]: any } }
const listCommon = inject('listCommon') as { ref: any }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {` + viewSaveDataInit + `
        ...saveCommon.data,
    } as { [propName: string]: any },
    rules: {` + viewSaveRule + `
    } as any,
    submit: () => {
        saveForm.ref.validate(async (valid: boolean) => {
            if (!valid) {
                return false
            }
            saveForm.loading = true
            const param = removeEmptyOfObj(saveForm.data)` + viewSaveParamHandle + `
            try {
                if (param?.idArr?.length > 0) {
                    await request(t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseCamelLower + `/` + tpl.TableNameCaseCamelLower + `/update', param, true)
                } else {
                    await request(t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseCamelLower + `/` + tpl.TableNameCaseCamelLower + `/create', param, true)
                }
                listCommon.ref.getList(true)
                saveCommon.visible = false
            } catch (error) {}
            saveForm.loading = false
        })
    },
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
            })
                .then(() => {
                    done()
                })
                .catch(() => {})
        } else {
            done()
        }
    },
    buttonClose: () => {
        //saveCommon.visible = false
        saveDrawer.ref.handleClose() //会触发beforeClose
    },
})` + viewFieldHandle + `
</script>

<template>
    <el-drawer class="save-drawer" :ref="(el: any) => saveDrawer.ref = el" v-model="saveCommon.visible" :title="saveCommon.title" :size="saveDrawer.size" :before-close="saveDrawer.beforeClose">
        <el-scrollbar>
            <el-form :ref="(el: any) => saveForm.ref = el" :model="saveForm.data" :rules="saveForm.rules" label-width="auto" :status-icon="true" :scroll-to-error="true">` + viewSaveField + `
            </el-form>
        </el-scrollbar>
        <template #footer>
            <el-button @click="saveDrawer.buttonClose">{{ t('common.cancel') }}</el-button>
            <el-button type="primary" @click="saveForm.submit" :loading="saveForm.loading">
                {{ t('common.save') }}
            </el-button>
        </template>
    </el-drawer>
</template>
`

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
		fieldCaseSnake := gstr.CaseSnakeFirstUpper(field)
		fieldCaseSnakeOfRemove := gstr.Split(fieldCaseSnake, `_of_`)[0]
		// fieldCaseCamelOfRemove := gstr.CaseCamel(fieldCaseSnakeOfRemove)
		fieldSplitArr := gstr.Split(fieldCaseSnakeOfRemove, `_`)
		// fieldPrefix := fieldSplitArr[0]
		fieldSuffix := fieldSplitArr[len(fieldSplitArr)-1]
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
		} else if garray.NewStrArrayFrom([]string{`UpdatedAt`, `UpdateAt`, `UpdatedTime`, `UpdateTime`}).Contains(fieldCaseCamel) {
			continue
		} else if garray.NewStrArrayFrom([]string{`CreatedAt`, `CreateAt`, `CreatedTime`, `CreateTime`}).Contains(fieldCaseCamel) {
			continue
		} else if column[`Key`].String() == `PRI` && column[`Extra`].String() == `auto_increment` { //主键
			continue
		} else if garray.NewStrArrayFrom([]string{`salt`}).Contains(fieldSuffix) && tpl.PasswordHandleMap[MyGenPasswordHandleMapKey(field)].IsCoexist { //salt后缀
			continue
		} else if garray.NewStrArrayFrom([]string{`status`, `type`, `method`, `pos`, `position`, `gender`}).Contains(fieldSuffix) && ((gstr.Pos(column[`Type`].String(), `int`) != -1 && gstr.Pos(column[`Type`].String(), `point`) == -1) || gstr.Pos(column[`Type`].String(), `char`) != -1) { //status,type,method,pos,position,gender等后缀
			isStr := true
			if gstr.Pos(column[`Type`].String(), `int`) != -1 && gstr.Pos(column[`Type`].String(), `point`) == -1 {
				isStr = false
			}
			statusList := MyGenStatusList(comment, isStr)
			viewI18nStatus += `
        ` + field + `: [`
			for _, status := range statusList {
				if isStr {
					viewI18nStatus += `
            { value: '` + status[0] + `', label: '` + status[1] + `' },`
				} else {
					viewI18nStatus += `
            { value: ` + status[0] + `, label: '` + status[1] + `' },`
				}
			}
			viewI18nStatus += `
        ],`
		} else if gstr.Pos(column[`Type`].String(), `int`) != -1 && gstr.Pos(column[`Type`].String(), `point`) == -1 { //int等类型
			if field == `pid` { //pid
				fieldName = `父级`
			} else if garray.NewStrArrayFrom([]string{`sort`, `weight`}).Contains(fieldSuffix) { //sort,weight等后缀
				viewI18nTip += `
        ` + field + `: '` + tip + `',`
			} else if garray.NewStrArrayFrom([]string{`id`}).Contains(fieldSuffix) { //id后缀
				if tpl.RelTableMap[field].IsExistRelTableDao && !tpl.RelTableMap[field].IsRedundRelNameField {
					fieldName = tpl.RelTableMap[field].RelTableFieldName
				}
			}
		} else if gstr.Pos(column[`Type`].String(), `json`) != -1 { //json类型
			viewI18nTip += `
        ` + field + `: '` + tip + `',`
		}

		viewI18nName += `
        ` + field + `: '` + fieldName + `',`
	}
	tplView := `export default {
    name: {` + viewI18nName + `
    },
    status: {` + viewI18nStatus + `
    },
    tip: {` + viewI18nTip + `
    },
}
`

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
            /*--------前端路由自动代码生成锚点（不允许修改和删除，否则将不能自动生成路由）--------*/`, 1)
	} else { //路由已存在则替换
		tplViewRouter, _ = gregex.ReplaceString(`\{
                path: '`+path+`',[\s\S]*'`+path+`' \}
            \},`, replaceStr, tplViewRouter)
	}
	gfile.PutContents(saveFile, tplViewRouter)

	MyGenMenu(ctx, tpl.SceneId, path, option.CommonName, tpl.TableNameCaseCamel) // 数据库权限菜单处理
}

// 自动生成操作
func MyGenAction(ctx context.Context, sceneId uint, actionCode string, actionName string) {
	actionName = gstr.Replace(actionName, `/`, `-`)

	idVar, _ := daoAuth.Action.ParseDbCtx(ctx).Where(daoAuth.Action.Columns().ActionCode, actionCode).Value(daoAuth.Action.PrimaryKey())
	id := idVar.Int64()
	if id == 0 {
		id, _ = daoAuth.Action.ParseDbCtx(ctx).Data(map[string]interface{}{
			daoAuth.Action.Columns().ActionCode: actionCode,
			daoAuth.Action.Columns().ActionName: actionName,
		}).InsertAndGetId()
	} else {
		daoAuth.Action.ParseDbCtx(ctx).Where(daoAuth.Action.PrimaryKey(), id).Data(daoAuth.Action.Columns().ActionName, actionName).Update()
	}
	daoAuth.ActionRelToScene.ParseDbCtx(ctx).Data(map[string]interface{}{
		daoAuth.ActionRelToScene.Columns().ActionId: id,
		daoAuth.ActionRelToScene.Columns().SceneId:  sceneId,
	}).Save()
}

// 自动生成菜单
func MyGenMenu(ctx context.Context, sceneId uint, menuUrl string, menuName string, menuNameOfEn string) {
	menuNameArr := gstr.Split(menuName, `/`)

	var pid int64 = 0
	for _, v := range menuNameArr[:len(menuNameArr)-1] {
		pidVar, _ := daoAuth.Menu.ParseDbCtx(ctx).Where(daoAuth.Menu.Columns().SceneId, sceneId).Where(daoAuth.Menu.Columns().MenuName, v).Value(daoAuth.Menu.PrimaryKey())
		if pidVar.Uint() == 0 {
			pid, _ = service.AuthMenu().Create(ctx, g.Map{
				daoAuth.Menu.Columns().SceneId:   sceneId,
				daoAuth.Menu.Columns().Pid:       pid,
				daoAuth.Menu.Columns().MenuName:  v,
				daoAuth.Menu.Columns().MenuIcon:  `autoicon-ep-link`,
				daoAuth.Menu.Columns().MenuUrl:   ``,
				daoAuth.Menu.Columns().ExtraData: `{"i18n": {"title": {"en": "", "zh-cn": "` + v + `"}}}`,
			})
		} else {
			pid = pidVar.Int64()
		}
	}

	menuName = menuNameArr[len(menuNameArr)-1]
	idVar, _ := daoAuth.Menu.ParseDbCtx(ctx).Where(daoAuth.Menu.Columns().SceneId, sceneId).Where(daoAuth.Menu.Columns().MenuUrl, menuUrl).Value(daoAuth.Menu.PrimaryKey())
	id := idVar.Uint()
	if id == 0 {
		service.AuthMenu().Create(ctx, g.Map{
			daoAuth.Menu.Columns().SceneId:   sceneId,
			daoAuth.Menu.Columns().Pid:       pid,
			daoAuth.Menu.Columns().MenuName:  menuName,
			daoAuth.Menu.Columns().MenuIcon:  `autoicon-ep-link`,
			daoAuth.Menu.Columns().MenuUrl:   menuUrl,
			daoAuth.Menu.Columns().ExtraData: `{"i18n": {"title": {"en": "` + menuNameOfEn + `", "zh-cn": "` + menuName + `"}}}`,
		})
	} else {
		service.AuthMenu().Update(ctx, g.Map{daoAuth.Menu.PrimaryKey(): id}, g.Map{
			daoAuth.Menu.Columns().MenuName:  menuName,
			daoAuth.Menu.Columns().Pid:       pid,
			daoAuth.Menu.Columns().ExtraData: `{"i18n": {"title": {"en": "` + menuNameOfEn + `", "zh-cn": "` + menuName + `"}}}`,
		})
	}
}

// status字段注释解析
func MyGenStatusList(comment string, isStrOpt ...bool) (statusList [][2]string) {
	isStr := false
	if len(isStrOpt) > 0 && isStrOpt[0] {
		isStr = true
	}

	var tmp [][]string
	if isStr {
		tmp, _ = gregex.MatchAllString(`([A-Za-z0-9]+)[-=:：]?([^\s,，;；)）]+)`, comment)
	} else {
		// tmp, _ = gregex.MatchAllString(`(-?\d+)[-=:：]?([^\d\s,，;；)）]+)`, comment)
		tmp, _ = gregex.MatchAllString(`(-?\d+)[-=:：]?([^\s,，;；)）]+)`, comment)
	}

	if len(tmp) == 0 {
		statusList = [][2]string{{`0`, `请设置表字段注释后，再生成代码`}}
		return
	}
	statusList = make([][2]string, len(tmp))
	for k, v := range tmp {
		statusList[k] = [2]string{v[1], v[2]}
	}
	return
}

// 获取PasswordHandleMap的Key（以Password为主）
func MyGenPasswordHandleMapKey(passwordOrsalt string) (passwordHandleMapKey string) {
	passwordOrsalt = gstr.Replace(gstr.CaseCamel(passwordOrsalt), `Salt`, `Password`, 1) //替换salt
	passwordOrsalt = gstr.Replace(passwordOrsalt, `Passwd`, `Password`, 1)               //替换passwd
	passwordHandleMapKey = gstr.CaseCamelLower(passwordOrsalt)                           //默认：小驼峰
	if gstr.CaseCamelLower(passwordOrsalt) != passwordOrsalt {                           //判断字段是不是蛇形
		passwordHandleMapKey = gstr.CaseSnake(passwordHandleMapKey)
	}
	return
}
