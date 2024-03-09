package my_gen

import (
	daoAuth "api/internal/dao/auth"
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type myGen struct {
	ctx       context.Context
	sceneId   uint     //场景ID
	sceneName string   //场景名称
	dbLink    string   //当前数据库连接配置（gf gen dao命令生成dao时需要）
	db        gdb.DB   //当前数据库连接
	tableArr  []string //当前db全部数据表
	option    myGenOption
	tpl       myGenTpl
}

// 命令参数解析后的数据
type myGenOption struct {
	SceneCode          string `json:"sceneCode"`          //场景标识，必须在数据库表auth_scene已存在。示例：platform
	DbGroup            string `json:"dbGroup"`            //db分组。示例：default
	DbTable            string `json:"dbTable"`            //db表。示例：auth_test
	RemovePrefixCommon string `json:"removePrefixCommon"` //要删除的共有前缀，没有可为空。removePrefixCommon + removePrefixAlone必须和hack/config.yaml内removePrefix保持一致
	RemovePrefixAlone  string `json:"removePrefixAlone"`  //要删除的独有前缀。removePrefixCommon + removePrefixAlone必须和hack/config.yaml内removePrefix保持一致，示例：auth_
	CommonName         string `json:"commonName"`         //公共名称，将同时在swagger文档Tag标签，权限菜单和权限操作中使用。示例：用户，权限管理/测试
	IsList             bool   `json:"isList" `            //是否生成列表接口(0和no为false，1和yes为true)
	IsCount            bool   `json:"isCount" `           //列表接口是否返回总数
	IsInfo             bool   `json:"isInfo" `            //是否生成详情接口
	IsCreate           bool   `json:"isCreate"`           //是否生成创建接口
	IsUpdate           bool   `json:"isUpdate"`           //是否生成更新接口
	IsDelete           bool   `json:"isDelete"`           //是否生成删除接口
	IsApi              bool   `json:"isApi"`              //是否生成后端接口文件
	IsAuthAction       bool   `json:"isAuthAction"`       //是否判断操作权限，如是，则同时会生成操作权限
	IsView             bool   `json:"isView"`             //是否生成前端视图文件
}

func NewMyGen(ctx context.Context, parser *gcmd.Parser) (myGenObj myGen) {
	defer func() {
		// TODO
		// myGenObj.tpl = myGenObj.createTpl(myGenObj.option.DbTable, myGenObj.option.RemovePrefixCommon, myGenObj.option.RemovePrefixAlone)
	}()

	myGenObj.ctx = ctx
	optionMap := parser.GetOptAll()
	gconv.Struct(optionMap, &myGenObj.option)

	// 命令执行前提示搭配Git使用
	gcmd.Scan(
		color.HiYellowString(`重要提示：强烈建议搭配Git使用，防止代码覆盖风险。`)+"\n",
		color.HiYellowString(`    Git库已创建或忽略风险，请按`)+color.HiGreenString(`[Enter]`)+color.HiYellowString(`继续执行`)+"\n",
		color.HiYellowString(`    Git库未创建，请按`)+color.HiRedString(`[Ctrl + C]`)+color.HiYellowString(`中断执行`)+"\n",
	)

	// 场景标识
	if myGenObj.option.SceneCode == `` {
		myGenObj.option.SceneCode = gcmd.Scan(color.BlueString(`> 请输入场景标识：`))
	}
	for {
		if myGenObj.option.SceneCode != `` {
			sceneInfo, _ := daoAuth.Scene.CtxDaoModel(myGenObj.ctx).Filter(daoAuth.Scene.Columns().SceneCode, myGenObj.option.SceneCode).One()
			if !sceneInfo.IsEmpty() {
				myGenObj.sceneId = sceneInfo[daoAuth.Scene.Columns().SceneId].Uint()
				myGenObj.sceneName = sceneInfo[daoAuth.Scene.Columns().SceneName].String()
				break
			}
		}
		myGenObj.option.SceneCode = gcmd.Scan(color.RedString(`    场景标识不存在，请重新输入：`))
	}
	// db分组
	if myGenObj.option.DbGroup == `` {
		myGenObj.option.DbGroup = gcmd.Scan(color.BlueString(`> 请输入db分组，默认(default)：`))
		if myGenObj.option.DbGroup == `` {
			myGenObj.option.DbGroup = `default`
		}
	}
	for {
		err := g.Try(myGenObj.ctx, func(ctx context.Context) {
			myGenObj.db = g.DB(myGenObj.option.DbGroup)
			myGenObj.dbLink = gconv.String(gconv.SliceMap(g.Cfg().MustGet(myGenObj.ctx, `database`).MapStrAny()[myGenObj.option.DbGroup])[0][`link`])
		})
		if err == nil {
			break
		}
		myGenObj.option.DbGroup = gcmd.Scan(color.RedString(`    db分组不存在，请重新输入，默认(default)：`))
		if myGenObj.option.DbGroup == `` {
			myGenObj.option.DbGroup = `default`
		}
	}
	// db表
	myGenObj.tableArr, _ = myGenObj.db.Tables(myGenObj.ctx)
	if myGenObj.option.DbTable == `` {
		myGenObj.option.DbTable = gcmd.Scan(color.BlueString(`> 请输入db表：`))
	}
	for {
		if myGenObj.option.DbTable != `` && garray.NewStrArrayFrom(myGenObj.tableArr).Contains(myGenObj.option.DbTable) {
			break
		}
		myGenObj.option.DbTable = gcmd.Scan(color.RedString(`    db表不存在，请重新输入：`))
	}
	// 要删除的共有前缀
	if _, ok := optionMap[`removePrefixCommon`]; !ok {
		myGenObj.option.RemovePrefixCommon = gcmd.Scan(color.BlueString(`> 请输入要删除的共有前缀，默认(空)：`))
	}
	for {
		if myGenObj.option.RemovePrefixCommon == `` || gstr.Pos(myGenObj.option.DbTable, myGenObj.option.RemovePrefixCommon) == 0 {
			break
		}
		myGenObj.option.RemovePrefixCommon = gcmd.Scan(color.RedString(`    要删除的共有前缀不存在，请重新输入，默认(空)：`))
	}
	// 要删除的独有前缀
	if _, ok := optionMap[`removePrefixAlone`]; !ok {
		myGenObj.option.RemovePrefixAlone = gcmd.Scan(color.BlueString(`> 请输入要删除的独有前缀，默认(空)：`))
	}
	for {
		if myGenObj.option.RemovePrefixAlone == `` || gstr.Pos(myGenObj.option.DbTable, myGenObj.option.RemovePrefixCommon+myGenObj.option.RemovePrefixAlone) == 0 {
			break
		}
		myGenObj.option.RemovePrefixAlone = gcmd.Scan(color.RedString(`    要删除的独有前缀不存在，请重新输入，默认(空)：`))
	}
	// 公共名称，将同时在swagger文档Tag标签，权限菜单和权限操作中使用。示例：场景
	for {
		if myGenObj.option.CommonName != `` {
			break
		}
		myGenObj.option.CommonName = gcmd.Scan(color.BlueString(`> 请输入公共名称，将同时在swagger文档Tag标签，权限菜单和权限操作中使用：`))
	}
noAllRestart:
	// 是否生成列表接口
	isList, ok := optionMap[`isList`]
	if !ok {
		isList = gcmd.Scan(color.BlueString(`> 是否生成列表接口，默认(yes)：`))
	}
isListEnd:
	for {
		switch isList {
		case ``, `1`, `yes`:
			myGenObj.option.IsList = true
			break isListEnd
		case `0`, `no`:
			myGenObj.option.IsList = false
			break isListEnd
		default:
			isList = gcmd.Scan(color.RedString(`    输入错误，请重新输入，是否生成列表接口，默认(yes)：`))
		}
	}
	// 列表接口是否返回总数
	isCount, ok := optionMap[`isCount`]
	if !ok {
		isCount = gcmd.Scan(color.BlueString(`> 列表接口是否返回总数，默认(yes)：`))
	}
isCountEnd:
	for {
		switch isCount {
		case ``, `1`, `yes`:
			myGenObj.option.IsCount = true
			break isCountEnd
		case `0`, `no`:
			myGenObj.option.IsCount = false
			break isCountEnd
		default:
			isCount = gcmd.Scan(color.RedString(`    输入错误，请重新输入，列表接口是否返回总数，默认(yes)：`))
		}
	}
	// 是否生成详情接口
	isInfo, ok := optionMap[`isInfo`]
	if !ok {
		isInfo = gcmd.Scan(color.BlueString(`> 是否生成详情接口，默认(yes)：`))
	}
isInfoEnd:
	for {
		switch isInfo {
		case ``, `1`, `yes`:
			myGenObj.option.IsInfo = true
			break isInfoEnd
		case `0`, `no`:
			myGenObj.option.IsInfo = false
			break isInfoEnd
		default:
			isInfo = gcmd.Scan(color.RedString(`    输入错误，请重新输入，是否生成详情接口，默认(yes)：`))
		}
	}
	// 是否生成创建接口
	isCreate, ok := optionMap[`isCreate`]
	if !ok {
		isCreate = gcmd.Scan(color.BlueString(`> 是否生成创建接口，默认(yes)：`))
	}
isCreateEnd:
	for {
		switch isCreate {
		case ``, `1`, `yes`:
			myGenObj.option.IsCreate = true
			break isCreateEnd
		case `0`, `no`:
			myGenObj.option.IsCreate = false
			break isCreateEnd
		default:
			isCreate = gcmd.Scan(color.RedString(`    输入错误，请重新输入，是否生成创建接口，默认(yes)：`))
		}
	}
	// 是否生成更新接口
	isUpdate, ok := optionMap[`isUpdate`]
	if !ok {
		isUpdate = gcmd.Scan(color.BlueString(`> 是否生成更新接口，默认(yes)：`))
	}
isUpdateEnd:
	for {
		switch isUpdate {
		case ``, `1`, `yes`:
			myGenObj.option.IsUpdate = true
			break isUpdateEnd
		case `0`, `no`:
			myGenObj.option.IsUpdate = false
			break isUpdateEnd
		default:
			isUpdate = gcmd.Scan(color.RedString(`    输入错误，请重新输入，是否生成更新接口，默认(yes)：`))
		}
	}
	// 是否生成删除接口
	isDelete, ok := optionMap[`isDelete`]
	if !ok {
		isDelete = gcmd.Scan(color.BlueString(`> 是否生成删除接口，默认(yes)：`))
	}
isDeleteEnd:
	for {
		switch isDelete {
		case ``, `1`, `yes`:
			myGenObj.option.IsDelete = true
			break isDeleteEnd
		case `0`, `no`:
			myGenObj.option.IsDelete = false
			break isDeleteEnd
		default:
			isDelete = gcmd.Scan(color.RedString(`    输入错误，请重新输入，是否生成删除接口，默认(yes)：`))
		}
	}
	if !(myGenObj.option.IsList || myGenObj.option.IsInfo || myGenObj.option.IsCreate || myGenObj.option.IsUpdate || myGenObj.option.IsDelete) {
		fmt.Println(`请重新选择生成哪些接口，不能全是no！`)
		goto noAllRestart
	}
	// 是否生成后端接口文件
	isApi, ok := optionMap[`isApi`]
	if !ok {
		isApi = gcmd.Scan(color.BlueString(`> 是否生成后端接口文件，默认(yes)：`))
	}
isApiEnd:
	for {
		switch isApi {
		case ``, `1`, `yes`:
			myGenObj.option.IsApi = true
			break isApiEnd
		case `0`, `no`:
			myGenObj.option.IsApi = false
			break isApiEnd
		default:
			isApi = gcmd.Scan(color.RedString(`    输入错误，请重新输入，是否生成后端接口文件，默认(yes)：`))
		}
	}
	if myGenObj.option.IsApi {
		// 是否判断操作权限，如是，则同时会生成操作权限
		isAuthAction, ok := optionMap[`isAuthAction`]
		if !ok {
			isAuthAction = gcmd.Scan(color.BlueString(`> 是否判断操作权限，如是，则同时会生成操作权限，默认(yes)：`))
		}
	isAuthActionEnd:
		for {
			switch isAuthAction {
			case ``, `1`, `yes`:
				myGenObj.option.IsAuthAction = true
				break isAuthActionEnd
			case `0`, `no`:
				myGenObj.option.IsAuthAction = false
				break isAuthActionEnd
			default:
				isAuthAction = gcmd.Scan(color.RedString(`    输入错误，请重新输入，是否判断操作权限，如是，则同时会生成操作权限，默认(yes)：`))
			}
		}
	}
	// 是否生成前端视图文件
	isView, ok := optionMap[`isView`]
	if !ok {
		isView = gcmd.Scan(color.BlueString(`> 是否生成前端视图文件，默认(yes)：`))
	}
isViewEnd:
	for {
		switch isView {
		case ``, `1`, `yes`:
			myGenObj.option.IsView = true
			break isViewEnd
		case `0`, `no`:
			myGenObj.option.IsView = false
			break isViewEnd
		default:
			isView = gcmd.Scan(color.RedString(`    输入错误，请重新输入，是否生成前端视图文件，默认(yes)：`))
		}
	}

	return myGenObj
}

// 初始化
func (myGenThis *myGen) handle() {
	/* myGenHandlerObj.genDao()   // dao模板生成
	myGenHandlerObj.genLogic() // logic模板生成

	if myGenHandlerObj.option.IsApi {
		myGenHandlerObj.genApi()        // api模板生成
		myGenHandlerObj.genController() // controller模板生成
		myGenHandlerObj.genRouter()     // 后端路由生成
	}

	if myGenHandlerObj.option.IsView {
		myGenHandlerObj.genViewIndex()  // 视图模板Index生成
		myGenHandlerObj.genViewList()   // 视图模板List生成
		myGenHandlerObj.genViewQuery()  // 视图模板Query生成
		myGenHandlerObj.genViewSave()   // 视图模板Save生成
		myGenHandlerObj.genViewI18n()   // 视图模板I18n生成
		myGenHandlerObj.genViewRouter() // 前端路由生成

		myGenHandlerObj.command(`前端代码格式化`, false, gfile.SelfDir()+`/../view/`+myGenHandlerObj.option.SceneCode, `npm`, `run`, `format`) // 前端代码格式化
	} */
}
