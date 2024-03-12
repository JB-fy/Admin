package my_gen

import (
	daoAuth "api/internal/dao/auth"
	"api/internal/utils"
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type myGen struct {
	ctx       context.Context
	sceneInfo gdb.Record //场景信息
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
		myGenObj.tpl = createTpl(ctx, myGenObj.option.DbGroup, myGenObj.option.DbTable, myGenObj.option.RemovePrefixCommon, myGenObj.option.RemovePrefixAlone)
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
			myGenObj.sceneInfo, _ = daoAuth.Scene.CtxDaoModel(myGenObj.ctx).Filter(daoAuth.Scene.Columns().SceneCode, myGenObj.option.SceneCode).One()
			if !myGenObj.sceneInfo.IsEmpty() {
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
		/* err := g.Try(myGenObj.ctx, func(ctx context.Context) {
			myGenObj.dbLink = gconv.String(gconv.SliceMap(g.Cfg().MustGet(myGenObj.ctx, `database`).MapStrAny()[myGenObj.option.DbGroup])[0][`link`])
		})
		if err == nil {
			break
		} */
		err := g.Try(ctx, func(ctx context.Context) {
			g.DB(myGenObj.option.DbGroup)
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
	tableArr, _ := g.DB(myGenObj.option.DbGroup).Tables(myGenObj.ctx)
	if myGenObj.option.DbTable == `` {
		myGenObj.option.DbTable = gcmd.Scan(color.BlueString(`> 请输入db表：`))
	}
	for {
		if myGenObj.option.DbTable != `` && garray.NewStrArrayFrom(tableArr).Contains(myGenObj.option.DbTable) {
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

// 生成代码
func (myGenThis *myGen) Handle() {
	myGenThis.genDao()   // dao模板生成
	myGenThis.genLogic() // logic模板生成

	if myGenThis.option.IsApi {
		myGenThis.genApi()        // api模板生成
		myGenThis.genController() // controller模板生成
		myGenThis.genRouter()     // 后端路由生成
	}

	if myGenThis.option.IsView {
		myGenThis.genViewIndex()  // 视图模板Index生成
		myGenThis.genViewList()   // 视图模板List生成
		myGenThis.genViewQuery()  // 视图模板Query生成
		myGenThis.genViewSave()   // 视图模板Save生成
		myGenThis.genViewI18n()   // 视图模板I18n生成
		myGenThis.genViewRouter() // 前端路由生成

		command(`前端代码格式化`, false, gfile.SelfDir()+`/../view/`+myGenThis.option.SceneCode, `npm`, `run`, `format`) // 前端代码格式化
	}
}

// dao模板生成
func (myGenThis *myGen) genDao() {
	tpl := myGenThis.tpl

	tpl.gfGenDao(true) //dao文件生成
	saveFile := gfile.SelfDir() + `/internal/dao/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseSnake + `.go`
	tplDao := gfile.GetContents(saveFile)

	type dao struct {
		ImportDao []string

		InsertParseBefore []string
		InsertParse       []string
		InsertHook        []string

		UpdateParse      []string
		UpdateHookBefore []string
		UpdateHookAfter  []string

		FieldParse []string
		FieldHook  []string

		FilterParse []string

		OrderParse []string

		JoinParse []string
	}
	daoObj := dao{}

	daoFieldList := getDaoFieldList(tpl)
	for _, v := range daoFieldList {
		daoObj.ImportDao = garray.NewStrArrayFrom(append(daoObj.ImportDao, v.ImportDao...)).Unique().Slice()

		daoObj.InsertParseBefore = garray.NewStrArrayFrom(append(daoObj.InsertParseBefore, v.InsertParseBefore.getData()...)).Unique().Slice()
		daoObj.InsertParse = garray.NewStrArrayFrom(append(daoObj.InsertParse, v.InsertParse.getData()...)).Unique().Slice()
		daoObj.InsertHook = garray.NewStrArrayFrom(append(daoObj.InsertHook, v.InsertHook.getData()...)).Unique().Slice()

		daoObj.UpdateParse = garray.NewStrArrayFrom(append(daoObj.UpdateParse, v.UpdateParse.getData()...)).Unique().Slice()
		daoObj.UpdateHookBefore = garray.NewStrArrayFrom(append(daoObj.UpdateHookBefore, v.UpdateHookBefore.getData()...)).Unique().Slice()
		daoObj.UpdateHookAfter = garray.NewStrArrayFrom(append(daoObj.UpdateHookAfter, v.UpdateHookAfter.getData()...)).Unique().Slice()

		daoObj.FieldParse = garray.NewStrArrayFrom(append(daoObj.FieldParse, v.FieldParse.getData()...)).Unique().Slice()
		daoObj.FieldHook = garray.NewStrArrayFrom(append(daoObj.FieldHook, v.FieldHook.getData()...)).Unique().Slice()

		daoObj.FilterParse = garray.NewStrArrayFrom(append(daoObj.FilterParse, v.FilterParse.getData()...)).Unique().Slice()

		daoObj.OrderParse = garray.NewStrArrayFrom(append(daoObj.OrderParse, v.OrderParse.getData()...)).Unique().Slice()

		daoObj.JoinParse = garray.NewStrArrayFrom(append(daoObj.JoinParse, v.JoinParse.getData()...)).Unique().Slice()
	}

	if len(daoObj.InsertParseBefore) > 0 {
		pointOfInsertParseBefore := `insertData := map[string]interface{}{}`
		tplDao = gstr.Replace(tplDao, pointOfInsertParseBefore, gstr.Join(append(daoObj.InsertParseBefore, ``), `
		`)+pointOfInsertParseBefore, 1)
	}
	if len(daoObj.InsertParse) > 0 {
		pointOfInsertParse := `case ` + "`id`" + `:
				insertData[daoThis.PrimaryKey()] = v`
		tplDao = gstr.Replace(tplDao, pointOfInsertParse, pointOfInsertParse+gstr.Join(append([]string{``}, daoObj.InsertParse...), `
			`), 1)
	}
	if len(daoObj.InsertHook) > 0 {
		pointOfInsertHook := `// id, _ := result.LastInsertId()

			/* for k, v := range daoModel.AfterInsert {
				switch k {
				case ` + "`xxxx`" + `:
					daoModel.CloneNew().Filter(daoThis.PrimaryKey(), id).HookUpdate(g.Map{k: v}).Update()
				}
			} */`
		tplDao = gstr.Replace(tplDao, pointOfInsertHook, `id, _ := result.LastInsertId()`+gstr.Join(append([]string{``}, daoObj.InsertHook...), `

			`), 1)
	}
	if len(daoObj.UpdateParse) > 0 {
		pointOfUpdateParse := `case ` + "`id`" + `:
				updateData[daoModel.DbTable+` + "`.`" + `+daoThis.PrimaryKey()] = v`
		tplDao = gstr.Replace(tplDao, pointOfUpdateParse, pointOfUpdateParse+gstr.Join(append([]string{``}, daoObj.UpdateParse...), `
			`), 1)
	}
	if len(daoObj.UpdateHookBefore) > 0 || len(daoObj.UpdateHookAfter) > 0 {
		pointOfUpdateHook := `/* row, _ := result.RowsAffected()
			if row == 0 {
				return
			} */

			/* for k, v := range daoModel.AfterUpdate {
				switch k {
				case ` + "`xxxx`" + `:
					for _, id := range daoModel.IdArr {
						daoModel.CloneNew().Filter(daoThis.PrimaryKey(), id).HookUpdate(g.Map{k: v}).Update()
					}
				}
			} */`
		if len(daoObj.UpdateHookBefore) > 0 {
			tplDao = gstr.Replace(tplDao, pointOfUpdateHook, gstr.Join(append(daoObj.UpdateHookBefore, ``), `

			`)+pointOfUpdateHook, 1)
		}
		if len(daoObj.UpdateHookAfter) > 0 {
			tplDao = gstr.Replace(tplDao, pointOfUpdateHook, `row, _ := result.RowsAffected()
			if row == 0 {
				return
			}`+gstr.Join(append([]string{``}, daoObj.UpdateHookAfter...), `

			`), 1)
		}
	}
	if len(daoObj.FieldParse) > 0 {
		pointOfFieldParse := `case ` + "`id`" + `:
				m = m.Fields(daoModel.DbTable + ` + "`.`" + ` + daoThis.PrimaryKey() + ` + "` AS `" + ` + v)`
		tplDao = gstr.Replace(tplDao, pointOfFieldParse, pointOfFieldParse+gstr.Join(append([]string{``}, daoObj.FieldParse...), `
			`), 1)
	}
	if len(daoObj.FieldHook) > 0 {
		pointOfFieldHook := `default:
						record[v] = gvar.New(nil)`
		tplDao = gstr.Replace(tplDao, pointOfFieldHook, gstr.Join(append(daoObj.FieldHook, ``), `
					`)+pointOfFieldHook, 1)
	}
	if len(daoObj.FilterParse) > 0 {
		pointOfFilterParse := `case ` + "`id`, `idArr`" + `:
				m = m.Where(daoModel.DbTable+` + "`.`" + `+daoThis.PrimaryKey(), v)`
		tplDao = gstr.Replace(tplDao, pointOfFilterParse, pointOfFilterParse+gstr.Join(append([]string{``}, daoObj.FilterParse...), `
			`), 1)
	}
	if len(daoObj.OrderParse) > 0 {
		pointOfOrderParse := `case ` + "`id`" + `:
				m = m.Order(daoModel.DbTable + ` + "`.`" + ` + gstr.Replace(v, k, daoThis.PrimaryKey(), 1))`
		tplDao = gstr.Replace(tplDao, pointOfOrderParse, pointOfOrderParse+gstr.Join(append([]string{``}, daoObj.OrderParse...), `
			`), 1)
	}
	if len(daoObj.JoinParse) > 0 {
		pointOfJoinParse := `/* case Xxxx.ParseDbTable(m.GetCtx()):
		m = m.LeftJoin(joinTable, joinTable+` + "`.`" + `+Xxxx.Columns().XxxxId+` + "` = `" + `+daoModel.DbTable+` + "`.`" + `+daoThis.PrimaryKey())
		// m = m.LeftJoin(Xxxx.ParseDbTable(m.GetCtx())+` + "` AS `" + `+joinTable, joinTable+` + "`.`" + `+Xxxx.Columns().XxxxId+` + "` = `" + `+daoModel.DbTable+` + "`.`" + `+daoThis.PrimaryKey()) */`
		tplDao = gstr.Replace(tplDao, pointOfJoinParse, pointOfJoinParse+gstr.Join(append([]string{``}, daoObj.JoinParse...), `
		`), 1)
	}
	if len(daoObj.ImportDao) > 0 {
		pointOfImportDao := `"api/internal/dao/` + tpl.ModuleDirCaseKebab + `/internal"`
		tplDao = gstr.Replace(tplDao, pointOfImportDao, pointOfImportDao+gstr.Join(append([]string{``}, daoObj.ImportDao...), `
	`), 1)
	}

	tplDao = gstr.Replace(tplDao, `"github.com/gogf/gf/v2/util/gconv"`, `"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"`, 1)

	gfile.PutContents(saveFile, tplDao)
	utils.GoFileFmt(saveFile)
}

// logic模板生成
func (myGenThis *myGen) genLogic() {
	tpl := myGenThis.tpl

	tplLogic := `package logic

import (
	dao` + tpl.ModuleDirCaseCamel + ` "api/internal/dao/` + tpl.ModuleDirCaseKebab + `"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type s` + myGenThis.tpl.LogicStructName + ` struct{}

func New` + myGenThis.tpl.LogicStructName + `() *s` + myGenThis.tpl.LogicStructName + ` {
	return &s` + myGenThis.tpl.LogicStructName + `{}
}

func init() {
	service.Register` + myGenThis.tpl.LogicStructName + `(New` + myGenThis.tpl.LogicStructName + `())
}

// 新增
func (logicThis *s` + myGenThis.tpl.LogicStructName + `) Create(ctx context.Context, data map[string]interface{}) (id int64, err error) {
	daoThis := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `
	daoModelThis := daoThis.CtxDaoModel(ctx)
`
	if tpl.Handle.Pid.Pid != `` {
		tplLogic += `
	if _, ok := data[daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Pid) + `]; ok {
		pid := gconv.Uint(data[daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Pid) + `])
		if pid > 0 {
			pInfo, _ := daoModelThis.CloneNew().Filter(daoThis.PrimaryKey(), pid).One()
			if pInfo.IsEmpty() {
				err = utils.NewErrorCode(ctx, 29999997, ` + "``" + `)
				return
			}
		}
	}
`
	}
	tplLogic += `
	id, err = daoModelThis.HookInsert(data).InsertAndGetId()
	return
}

// 修改
func (logicThis *s` + myGenThis.tpl.LogicStructName + `) Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error) {
	daoThis := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `
	daoModelThis := daoThis.CtxDaoModel(ctx)

	daoModelThis.Filters(filter).SetIdArr()
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ` + "``" + `)
		return
	}
`
	if tpl.Handle.Pid.Pid != `` {
		tplLogic += `
	if _, ok := data[daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Pid) + `]; ok {`
		if tpl.Handle.Pid.IsCoexist {
			tplLogic += `
		pid := gconv.Uint(data[daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Pid) + `])
		if pid > 0 {
			pInfo, _ := daoModelThis.CloneNew().Filter(daoThis.PrimaryKey(), pid).One()
			if pInfo.IsEmpty() {
				err = utils.NewErrorCode(ctx, 29999997, ` + "``" + `)
				return
			}
			oldList, _ := daoModelThis.CloneNew().Filter(daoThis.PrimaryKey(), daoModelThis.IdArr).All()
			for _, oldInfo := range oldList {
				if pid == oldInfo[daoThis.PrimaryKey()].Uint() { //父级不能是自身
					err = utils.NewErrorCode(ctx, 29999996, ` + "``" + `)
					return
				}
				if pid != oldInfo[daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Pid) + `].Uint() {
					if garray.NewStrArrayFrom(gstr.Split(pInfo[daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.IdPath) + `].String(), ` + "`-`" + `)).Contains(oldInfo[daoThis.PrimaryKey()].String()) { //父级不能是自身的子孙级
						err = utils.NewErrorCode(ctx, 29999995, ` + "``" + `)
						return
					}
				}
			}
		}`
		} else {
			tplLogic += `
		pid := gconv.Uint(data[daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Pid) + `])
		if pid > 0 {
			pInfo, _ := daoModelThis.CloneNew().Filter(daoThis.PrimaryKey(), pid).One()
			if pInfo.IsEmpty() {
				err = utils.NewErrorCode(ctx, 29999997, ` + "``" + `)
				return
			}
		}
		for _, id := range daoModelThis.IdArr {
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
	row, err = daoModelThis.HookUpdate(data).UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *s` + myGenThis.tpl.LogicStructName + `) Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error) {
	daoThis := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `
	daoModelThis := daoThis.CtxDaoModel(ctx)

	daoModelThis.Filters(filter).SetIdArr()
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ` + "``" + `)
		return
	}
`
	if tpl.Handle.Pid.Pid != `` {
		tplLogic += `
	count, _ := daoModelThis.CloneNew().Filter(daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Pid) + `, daoModelThis.IdArr).Count()
	if count > 0 {
		err = utils.NewErrorCode(ctx, 29999994, ` + "``" + `)
		return
	}
`
	}
	tplLogic += `
	row, err = daoModelThis.HookDelete().DeleteAndGetAffected()
	return
}
`

	saveFile := gfile.SelfDir() + `/internal/logic/` + gstr.Replace(tpl.ModuleDirCaseKebab, `/`, `-`) + `/` + tpl.TableCaseSnake + `.go`
	gfile.PutContents(saveFile, tplLogic)
	utils.GoFileFmt(saveFile)
	command(`service生成`, true, ``, `gf`, `gen`, `service`)
}

// api模板生成
func (myGenThis *myGen) genApi() {
	tpl := myGenThis.tpl

	type api struct {
		filter   []string
		create   []string
		update   []string
		res      []string
		resOfAdd []string
	}
	apiObj := api{
		filter:   []string{},
		create:   []string{},
		update:   []string{},
		res:      []string{},
		resOfAdd: []string{},
	}
	if len(tpl.Handle.LabelList) > 0 {
		apiObj.filter = append(apiObj.filter, `Label string `+"`"+`json:"label,omitempty" v:"max-length:30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"标签。常用于前端组件"`+"`")
		apiObj.res = append(apiObj.res, `Label *string `+"`"+`json:"label,omitempty" dc:"标签。常用于前端组件"`+"`")
	}

	type apiItem struct {
		isSkip bool //字段命名类型处理后，不再需要根据字段数据类型对filter,create,update,res四个属性再做处理时，设置为true
		filter bool
		create bool
		update bool
		res    bool

		isSkipType bool //字段命名类型处理后，不再需要根据字段数据类型对filterType,createType,updateType,resType四个属性再做处理时，设置为true
		filterType string
		createType string
		updateType string
		resType    string

		filterRule []string
		createRule []string
		updateRule []string
		isRequired bool //用于方便将required规则放首位
	}

	for _, v := range tpl.FieldList {
		apiItemObj := apiItem{
			filterRule: []string{},
			createRule: []string{},
			updateRule: []string{},
		}

		/*--------根据字段命名类型处理 开始--------*/
		switch v.FieldTypeName {
		case TypeNameDeleted: // 软删除字段
			continue
		case TypeNameUpdated: // 更新时间字段
			apiItemObj.isSkip = true
			apiItemObj.res = true
		case TypeNameCreated: // 创建时间字段
			apiObj.filter = append(apiObj.filter,
				`TimeRangeStart *gtime.Time `+"`"+`json:"timeRangeStart,omitempty" v:"date-format:Y-m-d H:i:s" dc:"开始时间：YYYY-mm-dd HH:ii:ss"`+"`",
				`TimeRangeEnd   *gtime.Time `+"`"+`json:"timeRangeEnd,omitempty" v:"date-format:Y-m-d H:i:s|after-equal:TimeRangeStart" dc:"结束时间：YYYY-mm-dd HH:ii:ss"`+"`",
			)
			apiItemObj.isSkip = true
			apiItemObj.res = true
		case TypeNamePri: // 主键
		case TypeNamePriAutoInc: // 主键（自增）
			if v.FieldRaw == `id` {
				continue
			}
			apiItemObj.isSkip = true
			apiItemObj.filter = true
			apiItemObj.res = true

			apiItemObj.filterRule = append(apiItemObj.filterRule, `min:1`)
		case TypeNamePid: // pid；	类型：int等类型；
			apiItemObj.isSkip = true
			apiItemObj.filter = true
			apiItemObj.create = true
			apiItemObj.update = true
			apiItemObj.res = true

			if len(tpl.Handle.LabelList) > 0 {
				apiObj.resOfAdd = append(apiObj.resOfAdd, `P`+gstr.CaseCamel(tpl.Handle.LabelList[0])+` *string `+"`"+`json:"p`+gstr.CaseCamel(tpl.Handle.LabelList[0])+`,omitempty" dc:"父级"`+"`")
			}
		case TypeNameLevel: // level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；
			apiItemObj.isSkip = true
			apiItemObj.filter = true
			apiItemObj.res = true

			apiItemObj.filterRule = append(apiItemObj.filterRule, `min:1`)
		case TypeNameIdPath: // idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；
			apiItemObj.isSkip = true
			apiItemObj.res = true
		case TypeNamePasswordSuffix: // password,passwd后缀；		类型：char(32)；
			apiItemObj.isSkip = true
			apiItemObj.create = true
			apiItemObj.update = true

			apiItemObj.isRequired = true
		case TypeNameSaltSuffix: // salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；
			continue
		case TypeNameNameSuffix: // name,title后缀；	类型：varchar；
			// 去掉该验证规则。有时会用到特殊符号
			// apiItemObj.filterRule = append(apiItemObj.filterRule, `regex:^[\\p{L}\\p{M}\\p{N}_-]+$`)
			// apiItemObj.createRule = append(apiItemObj.createRule, `regex:^[\\p{L}\\p{M}\\p{N}_-]+$`)
			// apiItemObj.updateRule = append(apiItemObj.updateRule, `regex:^[\\p{L}\\p{M}\\p{N}_-]+$`)
			if len(tpl.Handle.LabelList) > 0 && gstr.CaseCamel(tpl.Handle.LabelList[0]) == v.FieldCaseCamel {
				apiItemObj.isRequired = true
			}
		case TypeNameCodeSuffix: // code后缀；	类型：varchar；
			apiItemObj.filterRule = append(apiItemObj.filterRule, `regex:^[\\p{L}\\p{M}\\p{N}_-]+$`)
			apiItemObj.createRule = append(apiItemObj.createRule, `regex:^[\\p{L}\\p{M}\\p{N}_-]+$`)
			apiItemObj.updateRule = append(apiItemObj.updateRule, `regex:^[\\p{L}\\p{M}\\p{N}_-]+$`)
		case TypeNameAccountSuffix: // account后缀；	类型：varchar；
			// apiItemObj.filterRule = append(apiItemObj.filterRule, `passport`)
			// apiItemObj.createRule = append(apiItemObj.createRule, `passport`)
			// apiItemObj.updateRule = append(apiItemObj.updateRule, `passport`)
			apiItemObj.filterRule = append(apiItemObj.filterRule, `regex:^(?!\\d*$)[\\p{L}\\p{M}\\p{N}_]+$`)
			apiItemObj.createRule = append(apiItemObj.createRule, `regex:^(?!\\d*$)[\\p{L}\\p{M}\\p{N}_]+$`)
			apiItemObj.updateRule = append(apiItemObj.updateRule, `regex:^(?!\\d*$)[\\p{L}\\p{M}\\p{N}_]+$`)
		case TypeNamePhoneSuffix: // phone,mobile后缀；	类型：varchar；
			apiItemObj.filterRule = append(apiItemObj.filterRule, `phone`)
			apiItemObj.createRule = append(apiItemObj.createRule, `phone`)
			apiItemObj.updateRule = append(apiItemObj.updateRule, `phone`)
		case TypeNameEmailSuffix: // email后缀；	类型：varchar；
			apiItemObj.filterRule = append(apiItemObj.filterRule, `email`)
			apiItemObj.createRule = append(apiItemObj.createRule, `email`)
			apiItemObj.updateRule = append(apiItemObj.updateRule, `email`)
		case TypeNameUrlSuffix: // url,link后缀；	类型：varchar；
			apiItemObj.filterRule = append(apiItemObj.filterRule, `url`)
			apiItemObj.createRule = append(apiItemObj.createRule, `url`)
			apiItemObj.updateRule = append(apiItemObj.updateRule, `url`)
		case TypeNameIpSuffix: // IP后缀；	类型：varchar；
			apiItemObj.filterRule = append(apiItemObj.filterRule, `ip`)
			apiItemObj.createRule = append(apiItemObj.createRule, `ip`)
			apiItemObj.updateRule = append(apiItemObj.updateRule, `ip`)
		case TypeNameIdSuffix: // id后缀；	类型：int等类型；
			apiItemObj.isSkip = true
			apiItemObj.filter = true
			apiItemObj.create = true
			apiItemObj.update = true
			apiItemObj.res = true

			apiItemObj.filterRule = append(apiItemObj.filterRule, `min:1`)
			apiItemObj.createRule = append(apiItemObj.createRule, `min:1`)
			apiItemObj.updateRule = append(apiItemObj.updateRule, `min:1`)

			if tpl.Handle.RelIdMap[v.FieldRaw].tpl.Table != `` && !tpl.Handle.RelIdMap[v.FieldRaw].IsRedundName {
				relIdObj := tpl.Handle.RelIdMap[v.FieldRaw]
				apiObj.resOfAdd = append(apiObj.resOfAdd, gstr.CaseCamel(relIdObj.tpl.Handle.LabelList[0])+gstr.CaseCamel(relIdObj.Suffix)+` *string `+"`"+`json:"`+relIdObj.tpl.Handle.LabelList[0]+relIdObj.Suffix+`,omitempty" dc:"`+relIdObj.FieldName+`"`+"`")
			}
		case TypeNameSortSuffix, TypeNameSort: // sort,weight等后缀；	类型：int等类型； // sort，且pid,level,idPath|id_path,sort同时存在时（才）有效；	类型：int等类型；
			apiItemObj.createRule = append(apiItemObj.createRule, `between:0,100`)
			apiItemObj.updateRule = append(apiItemObj.updateRule, `between:0,100`)
		case TypeNameStatusSuffix: // status,type,method,pos,position,gender等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
			apiItemObj.isSkip = true
			apiItemObj.filter = true
			apiItemObj.create = true
			apiItemObj.update = true
			apiItemObj.res = true

			statusArr := make([]string, len(v.StatusList))
			for index, item := range v.StatusList {
				statusArr[index] = item[0]
			}
			statusStr := gstr.Join(statusArr, `,`)
			apiItemObj.filterRule = append(apiItemObj.filterRule, `in:`+statusStr)
			apiItemObj.createRule = append(apiItemObj.createRule, `in:`+statusStr)
			apiItemObj.updateRule = append(apiItemObj.updateRule, `in:`+statusStr)
		case TypeNameIsPrefix: // is_前缀；		类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（停用：0否 1是）
			apiItemObj.isSkip = true
			apiItemObj.filter = true
			apiItemObj.create = true
			apiItemObj.update = true
			apiItemObj.res = true

			/* TODO 可改成状态一样处理，同时需要修改前端开关组件属性设置（暂时不改）
			statusArr := make([]string, len(v.StatusList))
			for index, item := range v.StatusList {
				statusArr[index] = item[0]
			}
			statusStr := gstr.Join(statusArr, `,`)
			apiItemObj.filterRule = append(apiItemObj.filterRule, `in:`+statusStr)
			apiItemObj.createRule = append(apiItemObj.createRule, `in:`+statusStr)
			apiItemObj.updateRule = append(apiItemObj.updateRule, `in:`+statusStr) */
			apiItemObj.filterRule = append(apiItemObj.filterRule, `in:0,1`)
			apiItemObj.createRule = append(apiItemObj.createRule, `in:0,1`)
			apiItemObj.updateRule = append(apiItemObj.updateRule, `in:0,1`)
		case TypeNameStartPrefix: // start_前缀；	类型：timestamp或datetime或date；
			apiItemObj.isSkip = true
			apiItemObj.filter = true
			apiItemObj.create = true
			apiItemObj.update = true
			apiItemObj.res = true
		case TypeNameEndPrefix: // end_前缀；	类型：timestamp或datetime或date；
			apiItemObj.isSkip = true
			apiItemObj.filter = true
			apiItemObj.create = true
			apiItemObj.update = true
			apiItemObj.res = true
		case TypeNameRemarkSuffix: // remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器
			apiItemObj.isSkip = true
			apiItemObj.create = true
			apiItemObj.update = true
			apiItemObj.res = true
		case TypeNameImageSuffix, TypeNameVideoSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text	// video,video_list,videoList,video_arr,videoArr等后缀；		类型：单视频varchar，多视频json或text
			apiItemObj.isSkip = true
			apiItemObj.create = true
			apiItemObj.update = true
			apiItemObj.res = true
			if v.FieldType == TypeVarchar {
				apiItemObj.createRule = append(apiItemObj.createRule, `url`)
				apiItemObj.updateRule = append(apiItemObj.updateRule, `url`)
			} else {
				apiItemObj.isSkipType = true
				apiItemObj.createType = `*[]string`
				apiItemObj.updateType = `*[]string`
				apiItemObj.resType = `[]string`

				apiItemObj.createRule = append(apiItemObj.createRule, `distinct`, `foreach`, `url`, `foreach`, `min-length:1`)
				apiItemObj.updateRule = append(apiItemObj.updateRule, `distinct`, `foreach`, `url`, `foreach`, `min-length:1`)
				if !v.IsNull {
					apiItemObj.isRequired = true
				}
			}
		case TypeNameArrSuffix: // list,arr等后缀；	类型：json或text；
			apiItemObj.isSkip = true
			apiItemObj.create = true
			apiItemObj.update = true
			apiItemObj.res = true

			apiItemObj.isSkipType = true
			apiItemObj.createType = `*[]interface{}`
			apiItemObj.updateType = `*[]interface{}`
			apiItemObj.resType = `[]interface{}`

			apiItemObj.createRule = append(apiItemObj.createRule, `distinct`)
			apiItemObj.updateRule = append(apiItemObj.updateRule, `distinct`)
			if !v.IsNull {
				apiItemObj.isRequired = true
			}
		}
		/*--------根据字段命名类型处理 结束--------*/

		/*--------根据字段数据类型处理（注意：这里是字段命名类型处理的后续操作，改动需考虑兼容） 开始--------*/
		switch v.FieldType {
		case TypeInt: // `int等类型`
			if !apiItemObj.isSkip {
				// apiItemObj.filter = true
				apiItemObj.create = true
				apiItemObj.update = true
				apiItemObj.res = true
			}
			if !apiItemObj.isSkipType {
				apiItemObj.filterType = `*int`
				apiItemObj.createType = `*int`
				apiItemObj.updateType = `*int`
				apiItemObj.resType = `*int`
			}
		case TypeIntU: // `int等类型（unsigned）`
			if !apiItemObj.isSkip {
				// apiItemObj.filter = true
				apiItemObj.create = true
				apiItemObj.update = true
				apiItemObj.res = true
			}
			if !apiItemObj.isSkipType {
				apiItemObj.filterType = `*uint`
				apiItemObj.createType = `*uint`
				apiItemObj.updateType = `*uint`
				apiItemObj.resType = `*uint`
			}
		case TypeFloat: // `float等类型`
			if !apiItemObj.isSkip {
				// apiItemObj.filter = true
				apiItemObj.create = true
				apiItemObj.update = true
				apiItemObj.res = true
			}
			if !apiItemObj.isSkipType {
				apiItemObj.filterType = `*float64`
				apiItemObj.createType = `*float64`
				apiItemObj.updateType = `*float64`
				apiItemObj.resType = `*float64`
			}
		case TypeFloatU: // // `float等类型（unsigned）`
			if !apiItemObj.isSkip {
				// apiItemObj.filter = true
				apiItemObj.create = true
				apiItemObj.update = true
				apiItemObj.res = true
			}
			if !apiItemObj.isSkipType {
				apiItemObj.filterType = `*float64`
				apiItemObj.createType = `*float64`
				apiItemObj.updateType = `*float64`
				apiItemObj.resType = `*float64`
			}

			apiItemObj.filterRule = append([]string{`min:0`}, apiItemObj.filterRule...)
			apiItemObj.createRule = append([]string{`min:0`}, apiItemObj.createRule...)
			apiItemObj.updateRule = append([]string{`min:0`}, apiItemObj.updateRule...)
		case TypeVarchar: // `varchar类型`
			if !apiItemObj.isSkip {
				apiItemObj.filter = true
				apiItemObj.create = true
				apiItemObj.update = true
				apiItemObj.res = true
			}
			if !apiItemObj.isSkipType {
				apiItemObj.filterType = `string`
				apiItemObj.createType = `*string`
				apiItemObj.updateType = `*string`
				apiItemObj.resType = `*string`
			}

			apiItemObj.filterRule = append([]string{`max-length:` + v.FieldLimitStr}, apiItemObj.filterRule...)
			apiItemObj.createRule = append([]string{`max-length:` + v.FieldLimitStr}, apiItemObj.createRule...)
			apiItemObj.updateRule = append([]string{`max-length:` + v.FieldLimitStr}, apiItemObj.updateRule...)
			if v.IndexRaw == `UNI` && !v.IsNull {
				apiItemObj.isRequired = true
			}
		case TypeChar: // `char类型`
			if !apiItemObj.isSkip {
				apiItemObj.filter = true
				apiItemObj.create = true
				apiItemObj.update = true
				apiItemObj.res = true
			}
			if !apiItemObj.isSkipType {
				apiItemObj.filterType = `string`
				apiItemObj.createType = `*string`
				apiItemObj.updateType = `*string`
				apiItemObj.resType = `*string`
			}

			apiItemObj.filterRule = append([]string{`max-length:` + v.FieldLimitStr}, apiItemObj.filterRule...)
			apiItemObj.createRule = append([]string{`size:` + v.FieldLimitStr}, apiItemObj.createRule...)
			apiItemObj.updateRule = append([]string{`size:` + v.FieldLimitStr}, apiItemObj.updateRule...)
			if v.IndexRaw == `UNI` && !v.IsNull {
				apiItemObj.isRequired = true
			}
		case TypeText: // `text类型`
			if !apiItemObj.isSkip {
				// apiItemObj.filter = true
				apiItemObj.create = true
				apiItemObj.update = true
				apiItemObj.res = true
			}
			if !apiItemObj.isSkipType {
				apiItemObj.filterType = `string`
				apiItemObj.createType = `*string`
				apiItemObj.updateType = `*string`
				apiItemObj.resType = `*string`
			}
		case TypeJson: // `json类型`
			if !apiItemObj.isSkip {
				// apiItemObj.filter = true
				apiItemObj.create = true
				apiItemObj.update = true
				apiItemObj.res = true
			}
			if !apiItemObj.isSkipType {
				apiItemObj.filterType = `string`
				apiItemObj.createType = `*string`
				apiItemObj.updateType = `*string`
				apiItemObj.resType = `*string`
			}

			apiItemObj.filterRule = append([]string{`json`}, apiItemObj.filterRule...)
			apiItemObj.createRule = append([]string{`json`}, apiItemObj.createRule...)
			apiItemObj.updateRule = append([]string{`json`}, apiItemObj.updateRule...)
			if !v.IsNull {
				apiItemObj.isRequired = true
			}
		case TypeTimestamp, TypeDatetime: // `timestamp类型` // `datetime类型`
			if !apiItemObj.isSkip {
				// apiItemObj.filter = true
				apiItemObj.create = true
				apiItemObj.update = true
				apiItemObj.res = true
			}
			if !apiItemObj.isSkipType {
				apiItemObj.filterType = `*gtime.Time`
				apiItemObj.createType = `*gtime.Time`
				apiItemObj.updateType = `*gtime.Time`
				apiItemObj.resType = `*gtime.Time`
			}

			apiItemObj.filterRule = append([]string{`date-format:Y-m-d H:i:s`}, apiItemObj.filterRule...)
			apiItemObj.createRule = append([]string{`date-format:Y-m-d H:i:s`}, apiItemObj.createRule...)
			apiItemObj.updateRule = append([]string{`date-format:Y-m-d H:i:s`}, apiItemObj.updateRule...)
			if !v.IsNull && gconv.String(v.Default) == `` {
				apiItemObj.isRequired = true
			}
		case TypeDate: // `date类型`
			if !apiItemObj.isSkip {
				apiItemObj.filter = true
				apiItemObj.create = true
				apiItemObj.update = true
				apiItemObj.res = true
			}
			if !apiItemObj.isSkipType {
				apiItemObj.filterType = `*gtime.Time`
				apiItemObj.createType = `*gtime.Time`
				apiItemObj.updateType = `*gtime.Time`
				apiItemObj.resType = `*string`
			}

			apiItemObj.filterRule = append([]string{`date-format:Y-m-d`}, apiItemObj.filterRule...)
			apiItemObj.createRule = append([]string{`date-format:Y-m-d`}, apiItemObj.createRule...)
			apiItemObj.updateRule = append([]string{`date-format:Y-m-d`}, apiItemObj.updateRule...)
			if !v.IsNull && gconv.String(v.Default) == `` {
				apiItemObj.isRequired = true
			}
		default:
			if !apiItemObj.isSkip {
				apiItemObj.filter = true
				apiItemObj.create = true
				apiItemObj.update = true
				apiItemObj.res = true
			}
			if !apiItemObj.isSkipType {
				apiItemObj.filterType = `string`
				apiItemObj.createType = `*string`
				apiItemObj.updateType = `*string`
				apiItemObj.resType = `*string`
			}
		}
		/*--------根据字段数据类型处理（注意：这里是字段命名类型处理的后续操作，改动需考虑兼容） 结束--------*/

		if apiItemObj.filter {
			apiObj.filter = append(apiObj.filter, v.FieldCaseCamel+` `+apiItemObj.filterType+` `+"`"+`json:"`+v.FieldRaw+`,omitempty" v:"`+gstr.Join(apiItemObj.filterRule, `|`)+`" dc:"`+v.FieldDesc+`"`+"`")
		}
		if apiItemObj.create {
			if apiItemObj.isRequired {
				apiItemObj.createRule = append([]string{`required`}, apiItemObj.createRule...)
			}
			apiObj.create = append(apiObj.create, v.FieldCaseCamel+` `+apiItemObj.createType+` `+"`"+`json:"`+v.FieldRaw+`,omitempty" v:"`+gstr.Join(apiItemObj.createRule, `|`)+`" dc:"`+v.FieldDesc+`"`+"`")
		}
		if apiItemObj.update {
			apiObj.update = append(apiObj.update, v.FieldCaseCamel+` `+apiItemObj.updateType+` `+"`"+`json:"`+v.FieldRaw+`,omitempty" v:"`+gstr.Join(apiItemObj.updateRule, `|`)+`" dc:"`+v.FieldDesc+`"`+"`")
		}
		if apiItemObj.res {
			apiObj.res = append(apiObj.res, v.FieldCaseCamel+` `+apiItemObj.resType+` `+"`"+`json:"`+v.FieldRaw+`,omitempty" dc:"`+v.FieldDesc+`"`+"`")
		}
	}

	tplApi := `package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

`
	if myGenThis.option.IsList {
		tplApi += `
/*--------列表 开始--------*/
type ` + tpl.TableCaseCamel + `ListReq struct {
	g.Meta ` + "`" + `path:"/` + tpl.TableCaseKebab + `/list" method:"post" tags:"` + myGenThis.sceneInfo[daoAuth.Scene.Columns().SceneName].String() + `/` + myGenThis.option.CommonName + `" sm:"列表"` + "`" + `
	Filter ` + tpl.TableCaseCamel + `ListFilter ` + "`" + `json:"filter" dc:"过滤条件"` + "`" + `
	Field  []string        ` + "`" + `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"` + "`" + `
	Sort   string          ` + "`" + `json:"sort" default:"id DESC" dc:"排序"` + "`" + `
	Page   int             ` + "`" + `json:"page" v:"min:1" default:"1" dc:"页码"` + "`" + `
	Limit  int             ` + "`" + `json:"limit" v:"min:0" default:"10" dc:"每页数量。可传0取全部"` + "`" + `
}

type ` + tpl.TableCaseCamel + `ListFilter struct {
	Id             *uint       ` + "`" + `json:"id,omitempty" v:"min:1" dc:"ID"` + "`" + `
	IdArr          []uint      ` + "`" + `json:"idArr,omitempty" v:"distinct|foreach|min:1" dc:"ID数组"` + "`" + `
	ExcId          *uint       ` + "`" + `json:"excId,omitempty" v:"min:1" dc:"排除ID"` + "`" + `
	ExcIdArr       []uint      ` + "`" + `json:"excIdArr,omitempty" v:"distinct|foreach|min:1" dc:"排除ID数组"` + "`" + gstr.Join(append([]string{``}, apiObj.filter...), `
	`) + `
}

type ` + tpl.TableCaseCamel + `ListRes struct {`
		if myGenThis.option.IsCount {
			tplApi += `
	Count int         ` + "`" + `json:"count" dc:"总数"` + "`"
		}
		tplApi += `
	List  []` + tpl.TableCaseCamel + `ListItem ` + "`" + `json:"list" dc:"列表"` + "`" + `
}

type ` + tpl.TableCaseCamel + `ListItem struct {
	Id          *uint       ` + "`" + `json:"id,omitempty" dc:"ID"` + "`" + gstr.Join(append([]string{``}, apiObj.res...), `
	`) + gstr.Join(append([]string{``}, apiObj.resOfAdd...), `
	`) + `
}

/*--------列表 结束--------*/

`
	}
	if myGenThis.option.IsInfo {
		tplApi += `/*--------详情 开始--------*/
type ` + tpl.TableCaseCamel + `InfoReq struct {
	g.Meta ` + "`" + `path:"/` + tpl.TableCaseKebab + `/info" method:"post" tags:"` + myGenThis.sceneInfo[daoAuth.Scene.Columns().SceneName].String() + `/` + myGenThis.option.CommonName + `" sm:"详情"` + "`" + `
	Id     uint     ` + "`" + `json:"id" v:"required|min:1" dc:"ID"` + "`" + `
	Field  []string ` + "`" + `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"` + "`" + `
}

type ` + tpl.TableCaseCamel + `InfoRes struct {
	Info ` + tpl.TableCaseCamel + `Info ` + "`" + `json:"info" dc:"详情"` + "`" + `
}

type ` + tpl.TableCaseCamel + `Info struct {
	Id          *uint       ` + "`" + `json:"id,omitempty" dc:"ID"` + "`" + gstr.Join(append([]string{``}, apiObj.res...), `
	`) + `
}

/*--------详情 结束--------*/

`
	}
	if myGenThis.option.IsCreate {
		tplApi += `/*--------新增 开始--------*/
type ` + tpl.TableCaseCamel + `CreateReq struct {
	g.Meta      ` + "`" + `path:"/` + tpl.TableCaseKebab + `/create" method:"post" tags:"` + myGenThis.sceneInfo[daoAuth.Scene.Columns().SceneName].String() + `/` + myGenThis.option.CommonName + `" sm:"新增"` + "`" + gstr.Join(append([]string{``}, apiObj.create...), `
	`) + `
}

/*--------新增 结束--------*/

`
	}

	if myGenThis.option.IsUpdate {
		tplApi += `/*--------修改 开始--------*/
type ` + tpl.TableCaseCamel + `UpdateReq struct {
	g.Meta      ` + "`" + `path:"/` + tpl.TableCaseKebab + `/update" method:"post" tags:"` + myGenThis.sceneInfo[daoAuth.Scene.Columns().SceneName].String() + `/` + myGenThis.option.CommonName + `" sm:"修改"` + "`" + `
	IdArr       []uint  ` + "`" + `json:"idArr,omitempty" v:"required|distinct|foreach|min:1" dc:"ID数组"` + "`" + gstr.Join(append([]string{``}, apiObj.update...), `
	`) + `
}

/*--------修改 结束--------*/

`
	}

	if myGenThis.option.IsDelete {
		tplApi += `/*--------删除 开始--------*/
type ` + tpl.TableCaseCamel + `DeleteReq struct {
	g.Meta ` + "`" + `path:"/` + tpl.TableCaseKebab + `/del" method:"post" tags:"` + myGenThis.sceneInfo[daoAuth.Scene.Columns().SceneName].String() + `/` + myGenThis.option.CommonName + `" sm:"删除"` + "`" + `
	IdArr  []uint ` + "`" + `json:"idArr,omitempty" v:"required|distinct|foreach|min:1" dc:"ID数组"` + "`" + `
}

/*--------删除 结束--------*/
`
	}

	if myGenThis.option.IsList && tpl.Handle.Pid.Pid != `` {
		tplApi += `
/*--------列表（树状） 开始--------*/
type ` + tpl.TableCaseCamel + `TreeReq struct {
	g.Meta ` + "`" + `path:"/` + tpl.TableCaseKebab + `/tree" method:"post" tags:"` + myGenThis.sceneInfo[daoAuth.Scene.Columns().SceneName].String() + `/` + myGenThis.option.CommonName + `" sm:"列表（树状）"` + "`" + `
	Field  []string       ` + "`" + `json:"field" v:"foreach|min-length:1"` + "`" + `
	Filter ` + tpl.TableCaseCamel + `ListFilter ` + "`" + `json:"filter" dc:"过滤条件"` + "`" + `
}

type ` + tpl.TableCaseCamel + `TreeRes struct {
	Tree []` + tpl.TableCaseCamel + `TreeItem ` + "`" + `json:"tree" dc:"列表（树状）"` + "`" + `
}

type ` + tpl.TableCaseCamel + `TreeItem struct {
	Id       *uint       ` + "`" + `json:"id,omitempty" dc:"ID"` + "`" + gstr.Join(append([]string{``}, apiObj.res...), `
	`) + `
	Children []` + tpl.TableCaseCamel + `TreeItem ` + "`" + `json:"children" dc:"子级列表"` + "`" + `
}

/*--------列表（树状） 结束--------*/
`
	}

	saveFile := gfile.SelfDir() + `/api/` + myGenThis.option.SceneCode + `/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseSnake + `.go`
	gfile.PutContents(saveFile, tplApi)
	utils.GoFileFmt(saveFile)
}

// controller模板生成
func (myGenThis *myGen) genController() {
	tpl := myGenThis.tpl

	type controller struct {
		importDao []string
		list      []string
		info      []string
		tree      []string
		noAuth    []string
		// diff      []string // 可以不要。数据返回时，会根据API文件中的结构体做过滤
	}
	controllerObj := controller{
		list:   []string{"`id`"},
		info:   []string{"`id`"},
		tree:   []string{"`id`"},
		noAuth: []string{"`id`"},
	}

	if len(tpl.Handle.LabelList) > 0 {
		controllerObj.list = append(controllerObj.list, "`label`")
		controllerObj.info = append(controllerObj.info, "`label`")
		controllerObj.tree = append(controllerObj.tree, "`label`")
		if tpl.Handle.Pid.Pid != `` {
			controllerObj.list = append(controllerObj.list, "`p"+gstr.CaseCamel(tpl.Handle.LabelList[0])+"`")
			// controllerObj.info = append(controllerObj.info, "`p"+gstr.CaseCamel(tpl.Handle.LabelList[0])+"`")
		}
		controllerObj.noAuth = append(controllerObj.noAuth, "`label`")
		if tpl.FieldPrimary != `` && tpl.FieldPrimary != `id` {
			controllerObj.noAuth = append(controllerObj.noAuth, `dao`+tpl.ModuleDirCaseCamel+`.`+tpl.TableCaseCamel+`.Columns().`+gstr.CaseCamel(tpl.FieldPrimary))
		}
		controllerObj.noAuth = append(controllerObj.noAuth, `dao`+tpl.ModuleDirCaseCamel+`.`+tpl.TableCaseCamel+`.Columns().`+gstr.CaseCamel(tpl.Handle.LabelList[0]))
		/* for _, v := range tpl.Handle.LabelList {
			controllerObj.noAuth = append(controllerObj.noAuth, `dao`+tpl.ModuleDirCaseCamel+`.`+tpl.TableCaseCamel+`.Columns().`+gstr.CaseCamel(v))
		} */
	}

	for _, v := range tpl.FieldList {
		/*--------根据字段命名类型处理 开始--------*/
		switch v.FieldTypeName {
		case TypeNameDeleted: // 软删除字段
		case TypeNameUpdated: // 更新时间字段
		case TypeNameCreated: // 创建时间字段
		case TypeNamePri: // 主键
		case TypeNamePriAutoInc: // 主键（自增）
		case TypeNamePid: // pid；	类型：int等类型；
		case TypeNameLevel: // level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；
		case TypeNameIdPath: // idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；
		case TypeNamePasswordSuffix: // password,passwd后缀；		类型：char(32)；
			// controllerObj.diff = append(controllerObj.diff, `dao`+tpl.ModuleDirCaseCamel+`.`+tpl.TableCaseCamel+`.Columns().`+v.FieldCaseCamel)
		case TypeNameSaltSuffix: // salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；
			// controllerObj.diff = append(controllerObj.diff, `dao`+tpl.ModuleDirCaseCamel+`.`+tpl.TableCaseCamel+`.Columns().`+v.FieldCaseCamel)
		case TypeNameNameSuffix: // name,title后缀；	类型：varchar；
		case TypeNameCodeSuffix: // code后缀；	类型：varchar；
		case TypeNameAccountSuffix: // account后缀；	类型：varchar；
		case TypeNamePhoneSuffix: // phone,mobile后缀；	类型：varchar；
		case TypeNameEmailSuffix: // email后缀；	类型：varchar；
		case TypeNameUrlSuffix: // url,link后缀；	类型：varchar；
		case TypeNameIpSuffix: // IP后缀；	类型：varchar；
		case TypeNameIdSuffix: // id后缀；	类型：int等类型；
			relIdObj := tpl.Handle.RelIdMap[v.FieldRaw]
			if relIdObj.tpl.Table != `` && !relIdObj.IsRedundName {
				daoPath := `dao` + relIdObj.tpl.ModuleDirCaseCamel + `.` + relIdObj.tpl.TableCaseCamel
				importDaoStr := `dao` + relIdObj.tpl.ModuleDirCaseCamel + ` "api/internal/dao/` + relIdObj.tpl.ModuleDirCaseKebab + `"`
				if !garray.NewStrArrayFrom(controllerObj.importDao).Contains(importDaoStr) {
					controllerObj.importDao = append(controllerObj.importDao, importDaoStr)
				}
				fieldTmp := daoPath + `.Columns().` + gstr.CaseCamel(relIdObj.tpl.Handle.LabelList[0])
				if relIdObj.Suffix != `` {
					fieldTmp += "+`" + relIdObj.Suffix + "`"
				}
				controllerObj.list = append(controllerObj.list, fieldTmp)
			}
		case TypeNameSortSuffix, TypeNameSort: // sort,weight等后缀；	类型：int等类型； // sort，且pid,level,idPath|id_path,sort同时存在时（才）有效；	类型：int等类型；
		case TypeNameStatusSuffix: // status,type,method,pos,position,gender等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
		case TypeNameIsPrefix: // is_前缀；		类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（停用：0否 1是）
		case TypeNameStartPrefix: // start_前缀；	类型：timestamp或datetime或date；
		case TypeNameEndPrefix: // end_前缀；	类型：timestamp或datetime或date；
		case TypeNameRemarkSuffix: // remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器
		case TypeNameImageSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text
		case TypeNameVideoSuffix: // video,video_list,videoList,video_arr,videoArr等后缀；		类型：单视频varchar，多视频json或text
		case TypeNameArrSuffix: // list,arr等后缀；	类型：json或text；
		}
		/*--------根据字段命名类型处理 结束--------*/
	}

	tplController := `package controller

import (
	"api/api"
	api` + tpl.ModuleDirCaseCamel + ` "api/api/` + myGenThis.option.SceneCode + `/` + tpl.ModuleDirCaseKebab + `"
	dao` + tpl.ModuleDirCaseCamel + ` "api/internal/dao/` + tpl.ModuleDirCaseKebab + `"` + gstr.Join(append([]string{``}, controllerObj.importDao...), `
	`) + `
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/util/gconv"
)

type ` + tpl.TableCaseCamel + ` struct{}

func New` + tpl.TableCaseCamel + `() *` + tpl.TableCaseCamel + ` {
	return &` + tpl.TableCaseCamel + `{}
}
`
	if myGenThis.option.IsList {
		tplController += `
// 列表
func (controllerThis *` + tpl.TableCaseCamel + `) List(ctx context.Context, req *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `ListReq) (res *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `ListRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.Map(req.Filter, gconv.MapOption{Deep: true, OmitEmpty: true})
	if filter == nil {
		filter = map[string]interface{}{}
	}
`
		tplController += `
	allowField := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.ColumnArr().Slice()
	allowField = append(allowField` + gstr.Join(append([]string{``}, controllerObj.list...), `, `) + `)`
		/* if len(controllerObj.diff) > 0 {
			tplController += `
		allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{` + gstr.Join(controllerObj.diff, `, `) + `})).Slice() //移除敏感字段`
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
		if myGenThis.option.IsAuthAction {
			actionCode := gstr.CaseCamelLower(myGenThis.tpl.LogicStructName) + `Look`
			actionName := myGenThis.option.CommonName + `-查看`
			myGenThis.genAction(myGenThis.sceneInfo[daoAuth.Scene.Columns().SceneId].Uint(), actionCode, actionName) // 数据库权限操作处理
			tplController += `
	/**--------权限验证 开始--------**/
	isAuth, _ := service.AuthAction().CheckAuth(ctx, ` + "`" + actionCode + "`" + `)
	if !isAuth {
		field = []string{` + gstr.Join(controllerObj.noAuth, `, `) + `}
	}
	/**--------权限验证 结束--------**/
`
		}
		tplController += `
	daoModelThis := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.CtxDaoModel(ctx).Filters(filter)`
		if myGenThis.option.IsCount {
			tplController += `
	count, err := daoModelThis.CountPri()
	if err != nil {
		return
	}`
		}
		tplController += `
	list, err := daoModelThis.Fields(field).HookSelect().Order(req.Sort).Page(req.Page, req.Limit).ListPri()
	if err != nil {
		return
	}

	res = &api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `ListRes{`
		if myGenThis.option.IsCount {
			tplController += `Count: count, `
		}
		tplController += `List:  []api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `ListItem{}}
	list.Structs(&res.List)
	return
}
`
	}
	if myGenThis.option.IsInfo {
		tplController += `
// 详情
func (controllerThis *` + tpl.TableCaseCamel + `) Info(ctx context.Context, req *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `InfoReq) (res *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `InfoRes, err error) {
	/**--------参数处理 开始--------**/
	allowField := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.ColumnArr().Slice()
	allowField = append(allowField` + gstr.Join(append([]string{``}, controllerObj.info...), `, `) + `)`
		/* if len(controllerObj.diff) > 0 {
			tplController += `
		allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{` + gstr.Join(controllerObj.diff, `, `) + `})).Slice() //移除敏感字段`
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
		if myGenThis.option.IsAuthAction {
			actionCode := gstr.CaseCamelLower(myGenThis.tpl.LogicStructName) + `Look`
			actionName := myGenThis.option.CommonName + `-查看`
			myGenThis.genAction(myGenThis.sceneInfo[daoAuth.Scene.Columns().SceneId].Uint(), actionCode, actionName) // 数据库权限操作处理
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
	info, err := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.CtxDaoModel(ctx).Filters(filter).Fields(field).HookSelect().InfoPri()
	if err != nil {
		return
	}
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 29999998, ` + "``" + `)
		return
	}

	res = &api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `InfoRes{}
	info.Struct(&res.Info)
	return
}
`
	}
	if myGenThis.option.IsCreate {
		tplController += `
// 新增
func (controllerThis *` + tpl.TableCaseCamel + `) Create(ctx context.Context, req *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `CreateReq) (res *api.CommonCreateRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true})
	/**--------参数处理 结束--------**/
`
		if myGenThis.option.IsAuthAction {
			actionCode := gstr.CaseCamelLower(myGenThis.tpl.LogicStructName) + `Create`
			actionName := myGenThis.option.CommonName + `-新增`
			myGenThis.genAction(myGenThis.sceneInfo[daoAuth.Scene.Columns().SceneId].Uint(), actionCode, actionName) // 数据库权限操作处理
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
	id, err := service.` + myGenThis.tpl.LogicStructName + `().Create(ctx, data)
	if err != nil {
		return
	}
	res = &api.CommonCreateRes{Id: id}
	return
}
`
	}

	if myGenThis.option.IsUpdate {
		tplController += `
// 修改
func (controllerThis *` + tpl.TableCaseCamel + `) Update(ctx context.Context, req *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `UpdateReq) (res *api.CommonNoDataRes, err error) {
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
		if myGenThis.option.IsAuthAction {
			actionCode := gstr.CaseCamelLower(myGenThis.tpl.LogicStructName) + `Update`
			actionName := myGenThis.option.CommonName + `-编辑`
			myGenThis.genAction(myGenThis.sceneInfo[daoAuth.Scene.Columns().SceneId].Uint(), actionCode, actionName) // 数据库权限操作处理
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
	_, err = service.` + myGenThis.tpl.LogicStructName + `().Update(ctx, filter, data)
	return
}
`
	}

	if myGenThis.option.IsDelete {
		tplController += `
// 删除
func (controllerThis *` + tpl.TableCaseCamel + `) Delete(ctx context.Context, req *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `DeleteReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	filter := map[string]interface{}{` + "`id`" + `: req.IdArr}
	/**--------参数处理 结束--------**/
`
		if myGenThis.option.IsAuthAction {
			actionCode := gstr.CaseCamelLower(myGenThis.tpl.LogicStructName) + `Delete`
			actionName := myGenThis.option.CommonName + `-删除`
			myGenThis.genAction(myGenThis.sceneInfo[daoAuth.Scene.Columns().SceneId].Uint(), actionCode, actionName) // 数据库权限操作处理
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
	_, err = service.` + myGenThis.tpl.LogicStructName + `().Delete(ctx, filter)
	return
}
`
	}

	if myGenThis.option.IsList && tpl.Handle.Pid.Pid != `` {
		tplController += `
// 列表（树状）
func (controllerThis *` + tpl.TableCaseCamel + `) Tree(ctx context.Context, req *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `TreeReq) (res *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `TreeRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.Map(req.Filter, gconv.MapOption{Deep: true, OmitEmpty: true})
	if filter == nil {
		filter = map[string]interface{}{}
	}

	allowField := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.ColumnArr().Slice()
	allowField = append(allowField` + gstr.Join(append([]string{``}, controllerObj.tree...), `, `) + `)`
		/* if len(controllerObj.diff) > 0 {
			tplController += `
		allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{` + gstr.Join(controllerObj.diff, `, `) + `})).Slice() //移除敏感字段`
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
		if myGenThis.option.IsAuthAction {
			actionCode := gstr.CaseCamelLower(myGenThis.tpl.LogicStructName) + `Look`
			actionName := myGenThis.option.CommonName + `-查看`
			myGenThis.genAction(myGenThis.sceneInfo[daoAuth.Scene.Columns().SceneId].Uint(), actionCode, actionName) // 数据库权限操作处理
			tplController += `
	/**--------权限验证 开始--------**/
	isAuth, _ := service.AuthAction().CheckAuth(ctx, ` + "`" + actionCode + "`" + `)
	if !isAuth {
		field = []string{` + gstr.Join(controllerObj.noAuth, `, `) + `}
	}
	/**--------权限验证 结束--------**/
`
		}
		tplController += `
	field = append(field, ` + "`tree`" + `)

	list, err :=dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.CtxDaoModel(ctx).Filters(filter).Fields(field).HookSelect().ListPri()
	if err != nil {
		return
	}
	tree := utils.Tree(list.List(), 0, dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.Columns().` + gstr.CaseCamel(tpl.FieldPrimary) + `, dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Pid) + `)

	res = &api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `TreeRes{}
	gconv.Structs(tree, &res.Tree)
	return
}
`
	}

	saveFile := gfile.SelfDir() + `/internal/controller/` + myGenThis.option.SceneCode + `/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseSnake + `.go`
	gfile.PutContents(saveFile, tplController)
	utils.GoFileFmt(saveFile)
}

// 后端路由生成
func (myGenThis *myGen) genRouter() {
	tpl := myGenThis.tpl

	saveFile := gfile.SelfDir() + `/internal/router/` + myGenThis.option.SceneCode + `.go`
	tplRouter := gfile.GetContents(saveFile)

	//控制器不存在时导入
	importControllerStr := `controller` + tpl.ModuleDirCaseCamel + ` "api/internal/controller/` + myGenThis.option.SceneCode + `/` + tpl.ModuleDirCaseKebab + `"`
	if gstr.Pos(tplRouter, importControllerStr) == -1 {
		tplRouter = gstr.Replace(tplRouter, `"api/internal/middleware"`, importControllerStr+`
	"api/internal/middleware"`, 1)
		//路由生成
		tplRouter = gstr.Replace(tplRouter, `/*--------后端路由自动代码生成锚点（不允许修改和删除，否则将不能自动生成路由）--------*/`, `group.Group(`+"`"+`/`+tpl.ModuleDirCaseKebab+"`"+`, func(group *ghttp.RouterGroup) {
				group.Bind(controller`+tpl.ModuleDirCaseCamel+`.New`+tpl.TableCaseCamel+`())
			})

			/*--------后端路由自动代码生成锚点（不允许修改和删除，否则将不能自动生成路由）--------*/`, 1)
		gfile.PutContents(saveFile, tplRouter)
	} else {
		//路由不存在时需生成
		if gstr.Pos(tplRouter, `group.Bind(controller`+tpl.ModuleDirCaseCamel+`.New`+tpl.TableCaseCamel+`())`) == -1 {
			//路由生成
			tplRouter = gstr.Replace(tplRouter, `group.Group(`+"`"+`/`+tpl.ModuleDirCaseKebab+"`"+`, func(group *ghttp.RouterGroup) {`, `group.Group(`+"`"+`/`+tpl.ModuleDirCaseKebab+"`"+`, func(group *ghttp.RouterGroup) {
				group.Bind(controller`+tpl.ModuleDirCaseCamel+`.New`+tpl.TableCaseCamel+`())`, 1)
			gfile.PutContents(saveFile, tplRouter)
		}
	}

	utils.GoFileFmt(saveFile)
}

// 视图模板Index生成
func (myGenThis *myGen) genViewIndex() {
	tpl := myGenThis.tpl

	tplView := `<script setup lang="tsx">
import List from './List.vue'
import Query from './Query.vue'`
	if myGenThis.option.IsCreate || myGenThis.option.IsUpdate {
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
	if myGenThis.option.IsCreate || myGenThis.option.IsUpdate {
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
	if myGenThis.option.IsCreate || myGenThis.option.IsUpdate {
		tplView += `

        <!-- 加上v-if每次都重新生成组件。可防止不同操作之间的影响；新增操作数据的默认值也能写在save组件内 -->
        <save v-if="saveCommon.visible" />`
	}
	tplView += `
    </el-container>
</template>
`

	saveFile := gfile.SelfDir() + `/../view/` + myGenThis.option.SceneCode + `/src/views/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/Index.vue`
	gfile.PutContents(saveFile, tplView)
}

// 视图模板List生成
func (myGenThis *myGen) genViewList() {
	tpl := myGenThis.tpl

	type viewList struct {
		rowHeight uint
		columns   []string
	}
	viewListObj := viewList{
		rowHeight: 50,
		columns:   []string{},
	}
	type columnAttr struct {
		dataKey      string
		title        string
		key          string
		align        string
		width        string
		sortable     string
		hidden       string
		cellRenderer string
	}
	for _, v := range tpl.FieldList {
		columnAttrObj := columnAttr{
			dataKey: `'` + v.FieldRaw + `'`,
			title:   `t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')`,
			key:     `'` + v.FieldRaw + `'`,
			align:   `'center'`,
			width:   `150`,
		}

		/*--------根据字段命名类型处理 开始--------*/
		switch v.FieldTypeName {
		case TypeNameDeleted: // 软删除字段
			continue
		case TypeNameUpdated: // 更新时间字段
			columnAttrObj.title = `t('common.name.updatedAt')`
		case TypeNameCreated: // 创建时间字段
			columnAttrObj.title = `t('common.name.createdAt')`
		case TypeNamePri: // 主键
		case TypeNamePriAutoInc: // 主键（自增）
			continue
		case TypeNamePid: // pid；	类型：int等类型；
			if len(tpl.Handle.LabelList) > 0 {
				columnAttrObj.dataKey = `'p` + gstr.CaseCamel(tpl.Handle.LabelList[0]) + `'`
			}
		case TypeNameLevel: // level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；
			columnAttrObj.sortable = `true`
		case TypeNameIdPath: // idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；
			columnAttrObj.hidden = `true`
		case TypeNamePasswordSuffix: // password,passwd后缀；		类型：char(32)；
			continue
		case TypeNameSaltSuffix: // salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；
			continue
		case TypeNameNameSuffix: // name,title后缀；	类型：varchar；
		case TypeNameCodeSuffix: // code后缀；	类型：varchar；
		case TypeNameAccountSuffix: // account后缀；	类型：varchar；
		case TypeNamePhoneSuffix: // phone,mobile后缀；	类型：varchar；
		case TypeNameEmailSuffix: // email后缀；	类型：varchar；
		case TypeNameUrlSuffix: // url,link后缀；	类型：varchar；
		case TypeNameIpSuffix: // IP后缀；	类型：varchar；
		case TypeNameIdSuffix: // id后缀；	类型：int等类型；
			if tpl.Handle.RelIdMap[v.FieldRaw].tpl.Table != `` && !tpl.Handle.RelIdMap[v.FieldRaw].IsRedundName {
				columnAttrObj.dataKey = `'` + tpl.Handle.RelIdMap[v.FieldRaw].tpl.Handle.LabelList[0] + tpl.Handle.RelIdMap[v.FieldRaw].Suffix + `'`
			}
		case TypeNameSortSuffix, TypeNameSort: // sort,weight等后缀；	类型：int等类型； // sort，且pid,level,idPath|id_path,sort同时存在时（才）有效；	类型：int等类型；
			columnAttrObj.sortable = `true`
			if myGenThis.option.IsUpdate {
				columnAttrObj.cellRenderer = `(props: any): any => {
                if (props.rowData.edit` + gstr.CaseCamel(v.FieldRaw) + `) {
                    let currentRef: any
                    let currentVal = props.rowData.` + v.FieldRaw + `
                    return [
                        <el-input-number
                            ref={(el: any) => {
                                currentRef = el
                                el?.focus()
                            }}
                            model-value={currentVal}
                            placeholder={t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.tip.` + v.FieldRaw + `')}
                            precision={0}
                            min={0}
                            max={100}
                            step={1}
                            step-strictly={true}
                            controls={false} //控制按钮会导致诸多问题。如：焦点丢失；` + v.FieldRaw + `是0或100时，只一个按钮可点击
                            controls-position="right"
                            onChange={(val: number) => (currentVal = val)}
                            onBlur={() => {
                                props.rowData.edit` + gstr.CaseCamel(v.FieldRaw) + ` = false
                                if ((currentVal || currentVal === 0) && currentVal != props.rowData.` + v.FieldRaw + `) {
                                    handleUpdate({
                                        idArr: [props.rowData.id],
                                        ` + v.FieldRaw + `: currentVal,
                                    })
                                        .then((res) => {
                                            props.rowData.` + v.FieldRaw + ` = currentVal
                                        })
                                        .catch((error) => {})
                                }
                            }}
                            onKeydown={(event: any) => {
                                switch (event.keyCode) {
                                    // case 27:    //Esc键：Escape
                                    // case 32:    //空格键：" "
                                    case 13: //Enter键：Enter
                                        // props.rowData.edit` + gstr.CaseCamel(v.FieldRaw) + ` = false    //也会触发onBlur事件
                                        currentRef?.blur()
                                        break
                                }
                            }}
                        />,
                    ]
                }
                return [
                    <div class="inline-edit" onClick={() => (props.rowData.edit` + gstr.CaseCamel(v.FieldRaw) + ` = true)}>
                        {props.rowData.` + v.FieldRaw + `}
                    </div>,
                ]
            }`
			}
		case TypeNameStatusSuffix: // status,type,method,pos,position,gender等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
			columnAttrObj.cellRenderer = `(props: any): any => {
                let tagType = tm('config.const.tagType') as string[]
                let obj = tm('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.status.` + v.FieldRaw + `') as { value: any, label: string }[]
                let index = obj.findIndex((item) => { return item.value == props.rowData.` + v.FieldRaw + ` })
                return <el-tag type={tagType[index % tagType.length]}>{obj[index]?.label}</el-tag>
            }`
		case TypeNameIsPrefix: // is_前缀；		类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（停用：0否 1是）
			cellRendererTmp := `disabled={true}`
			if myGenThis.option.IsUpdate {
				cellRendererTmp = `onChange={(val: number) => {
                            handleUpdate({
                                idArr: [props.rowData.id],
                                ` + v.FieldRaw + `: val,
                            })
                                .then((res) => {
                                    props.rowData.` + v.FieldRaw + ` = val
                                })
                                .catch((error) => {})
                        }}`
			}
			columnAttrObj.cellRenderer = `(props: any): any => {
                return [
                    <el-switch
                        model-value={props.rowData.` + v.FieldRaw + `}
                        active-value={1}
                        inactive-value={0}
                        inline-prompt={true}
                        active-text={t('common.yes')}
                        inactive-text={t('common.no')}
                        style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success);"
                        ` + cellRendererTmp + `
                    />,
                ]
            }`
		case TypeNameStartPrefix: // start_前缀；	类型：timestamp或datetime或date；
		case TypeNameEndPrefix: // end_前缀；	类型：timestamp或datetime或date；
		case TypeNameRemarkSuffix: // remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器
		case TypeNameImageSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text
			columnAttrObj.width = `100`
			cellRendererTmp := `
                const imageList = [props.rowData.` + v.FieldRaw + `]`
			if v.FieldType != TypeVarchar {
				cellRendererTmp = `
                let imageList: string[]
                if (Array.isArray(props.rowData.` + v.FieldRaw + `)) {
                    imageList = props.rowData.` + v.FieldRaw + `
                } else {
                    imageList = JSON.parse(props.rowData.` + v.FieldRaw + `)
                }`
			}
			columnAttrObj.cellRenderer = `(props: any): any => {
                if (!props.rowData.` + v.FieldRaw + `) {
                    return
                }` + cellRendererTmp + `
                return [
                    <el-scrollbar wrap-style="display: flex; align-items: center;" view-style="margin: auto;">
                        {imageList.map((item) => {
                            //修改宽高时，可同时修改table属性row-height增加行高，则不会显示滚动条
                            return <el-image style="width: 45px;" src={item} lazy={true} hide-on-click-modal={true} preview-teleported={true} preview-src-list={imageList} />
                        })}
                    </el-scrollbar>
                ]
            }`
			goto skipFieldTypeOfViewList
		case TypeNameVideoSuffix: // video,video_list,videoList,video_arr,videoArr等后缀；		类型：单视频varchar，多视频json或text
			if viewListObj.rowHeight < 100 {
				viewListObj.rowHeight = 100
			}
			cellRendererTmp := `
                const videoList = [props.rowData.` + v.FieldRaw + `]`
			if v.FieldType != TypeVarchar {
				cellRendererTmp = `
                let videoList: string[]
                if (Array.isArray(props.rowData.` + v.FieldRaw + `)) {
                    videoList = props.rowData.` + v.FieldRaw + `
                } else {
                    videoList = JSON.parse(props.rowData.` + v.FieldRaw + `)
                }`
			}
			columnAttrObj.cellRenderer = `(props: any): any => {
                if (!props.rowData.` + v.FieldRaw + `) {
                    return
                }` + cellRendererTmp + `
                return [
                    <el-scrollbar wrap-style="display: flex; align-items: center;" view-style="margin: auto;">
                        {videoList.map((item) => {
                            //修改宽高时，可同时修改table属性row-height增加行高，则不会显示滚动条
                            return <video style="width: 120px; height: 80px;" preload="none" controls={true} src={item} />
                        })}
                    </el-scrollbar>,
                ]
            }`
			goto skipFieldTypeOfViewList
		case TypeNameArrSuffix: // list,arr等后缀；	类型：json或text；
			columnAttrObj.width = `100`
			columnAttrObj.cellRenderer = `(props: any): any => {
                if (!props.rowData.` + v.FieldRaw + `) {
                    return
                }
                let arrList: any[]
                if (Array.isArray(props.rowData.` + v.FieldRaw + `)) {
                    arrList = props.rowData.` + v.FieldRaw + `
                } else {
                    arrList = JSON.parse(props.rowData.` + v.FieldRaw + `)
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
            }`
			goto skipFieldTypeOfViewList
		}
		/*--------根据字段命名类型处理 结束--------*/

		/*--------根据字段数据类型处理（注意：这里是字段命名类型处理的后续操作，改动需考虑兼容） 开始--------*/
		switch v.FieldType {
		case TypeInt, TypeIntU: // `int等类型` // `int等类型（unsigned）`
			if gstr.Pos(v.FieldTypeRaw, `tinyint`) != -1 || gstr.Pos(v.FieldTypeRaw, `smallint`) != -1 {
				columnAttrObj.width = `100`
			}
		case TypeFloat: // `float等类型`
		case TypeFloatU: // `float等类型（unsigned）`
		case TypeVarchar, TypeChar: // `varchar类型` // `char类型`
			if gconv.Uint(v.FieldLimitStr) >= 120 {
				columnAttrObj.width = `200`
				columnAttrObj.hidden = `true`
			}
		case TypeText, TypeJson: // `text类型` // `json类型`
			columnAttrObj.width = `200`
			columnAttrObj.hidden = `true`
		case TypeTimestamp, TypeDatetime: // `timestamp类型` // `datetime类型`
			columnAttrObj.sortable = `true`
		case TypeDate: // `date类型`
			columnAttrObj.width = `100`
			columnAttrObj.sortable = `true`
		}
		/*--------根据字段数据类型处理（注意：这里是字段命名类型处理的后续操作，改动需考虑兼容） 结束--------*/

	skipFieldTypeOfViewList: //跳过字段数据类型处理标签
		columnAttrStr := []string{
			`dataKey: ` + columnAttrObj.dataKey + `,`,
			`title: ` + columnAttrObj.title + `,`,
			`key: ` + columnAttrObj.key + `,`,
			`align: ` + columnAttrObj.align + `,`,
			`width: ` + columnAttrObj.width + `,`,
		}
		if columnAttrObj.sortable != `` {
			columnAttrStr = append(columnAttrStr, `sortable: `+columnAttrObj.sortable+`,`)
		}
		if columnAttrObj.hidden != `` {
			columnAttrStr = append(columnAttrStr, `hidden: `+columnAttrObj.hidden+`,`)
		}
		if columnAttrObj.cellRenderer != `` {
			columnAttrStr = append(columnAttrStr, `cellRenderer: `+columnAttrObj.cellRenderer+`,`)
		}
		viewListObj.columns = append(viewListObj.columns, `{`+gstr.Join(append([]string{``}, columnAttrStr...), `
            `)+`
        },`)
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
	if myGenThis.option.IsUpdate || myGenThis.option.IsDelete {
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
        },` + gstr.Join(append([]string{``}, viewListObj.columns...), `
        `)
	if myGenThis.option.IsCreate || myGenThis.option.IsUpdate || myGenThis.option.IsDelete {
		tplView += `
        {
            title: t('common.name.action'),
            key: 'action',
            align: 'center',
            width: 250,
            fixed: 'right',
            cellRenderer: (props: any): any => {
                return [`
		if myGenThis.option.IsUpdate {
			tplView += `
                    <el-button type="primary" size="small" onClick={() => handleEditCopy(props.rowData.id)}>
                        <autoicon-ep-edit />
                        {t('common.edit')}
                    </el-button>,`
		}
		if myGenThis.option.IsDelete {
			tplView += `
                    <el-button type="danger" size="small" onClick={() => handleDelete(props.rowData.id)}>
                        <autoicon-ep-delete />
                        {t('common.delete')}
                    </el-button>,`
		}
		if myGenThis.option.IsCreate {
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
	if myGenThis.option.IsCreate || myGenThis.option.IsUpdate {
		tplView += `

const saveCommon = inject('saveCommon') as { visible: boolean; title: string; data: { [propName: string]: any } }`
	}
	if myGenThis.option.IsCreate {
		tplView += `
//新增
const handleAdd = () => {
    saveCommon.data = {}
    saveCommon.title = t('common.add')
    saveCommon.visible = true
}`
	}
	if myGenThis.option.IsDelete {
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
	if myGenThis.option.IsCreate || myGenThis.option.IsUpdate {
		tplView += `
//编辑|复制
const handleEditCopy = (id: number, type: string = 'edit') => {
    request(t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/info', { id: id })
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
	if myGenThis.option.IsDelete {
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
            request(t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/del', { idArr: idArr }, true)
                .then((res) => {
                    getList()
                })
                .catch(() => {})
        })
        .catch(() => {})
}`
	}
	if myGenThis.option.IsUpdate {
		tplView += `
//更新
const handleUpdate = async (param: { idArr: number[]; [propName: string]: any }) => {
    await request(t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/update', param, true)
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
        const res = await request(t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/list', param)
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
	if myGenThis.option.IsCreate {
		tplView += `
                <el-button type="primary" @click="handleAdd"> <autoicon-ep-edit-pen />{{ t('common.add') }} </el-button>`
	}
	if myGenThis.option.IsDelete {
		tplView += `
                <el-button type="danger" @click="handleBatchDelete"> <autoicon-ep-delete-filled />{{ t('common.batchDelete') }} </el-button>`
	}
	tplView += `
            </el-space>
        </el-col>
        <el-col :span="8" style="text-align: right">
            <el-space :size="10" style="height: 100%">
                <my-export-button i18nPrefix="` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `" :headerList="table.columns" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/list', param: { filter: queryCommon.data, sort: table.sort.key + ' ' + table.sort.order } }" />
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
                <el-table-v2 class="main-table" :columns="table.columns" :data="table.data" :sort-by="table.sort" @column-sort="table.handleSort" :width="width" :height="height" :fixed="true" :row-height="` + gconv.String(viewListObj.rowHeight) + `">
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

	saveFile := gfile.SelfDir() + `/../view/` + myGenThis.option.SceneCode + `/src/views/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/List.vue`
	gfile.PutContents(saveFile, tplView)
}

// 视图模板Query生成
func (myGenThis *myGen) genViewQuery() {
	tpl := myGenThis.tpl

	type viewQuery struct {
		dataInit []string
		form     []string
	}
	viewQueryObj := viewQuery{
		dataInit: []string{},
		form:     []string{},
	}

	type viewQueryItem struct {
		elDatePicker struct {
			defaultTime string
		}
	}
	for _, v := range tpl.FieldList {
		viewQueryItemObj := viewQueryItem{}
		/*--------根据字段命名类型处理 开始--------*/
		switch v.FieldTypeName {
		case TypeNameDeleted: // 软删除字段
			continue
		case TypeNameUpdated: // 更新时间字段
			continue
		case TypeNameCreated: // 创建时间字段
			viewQueryObj.dataInit = append(viewQueryObj.dataInit,
				`timeRange: (() => {
        return undefined
        /* const date = new Date()
        return [
            new Date(date.getFullYear(), date.getMonth(), date.getDate(), 0, 0, 0),
            new Date(date.getFullYear(), date.getMonth(), date.getDate(), 23, 59, 59),
        ] */
    })(),`,
				`timeRangeStart: computed(() => {
        if (queryCommon.data.timeRange?.length) {
            return dayjs(queryCommon.data.timeRange[0]).format('YYYY-MM-DD HH:mm:ss')
        }
        return ''
    }),`,
				`timeRangeEnd: computed(() => {
        if (queryCommon.data.timeRange?.length) {
            return dayjs(queryCommon.data.timeRange[1]).format('YYYY-MM-DD HH:mm:ss')
        }
        return ''
    }),`,
			)
			viewQueryObj.form = append(viewQueryObj.form, `<el-form-item prop="timeRange">
            <el-date-picker v-model="queryCommon.data.timeRange" type="datetimerange" range-separator="-" :default-time="[new Date(2000, 0, 1, 0, 0, 0), new Date(2000, 0, 1, 23, 59, 59)]" :start-placeholder="t('common.name.timeRangeStart')" :end-placeholder="t('common.name.timeRangeEnd')" />
        </el-form-item>`)
			continue
		case TypeNamePri: // 主键
		case TypeNamePriAutoInc: // 主键（自增）
			continue
		case TypeNamePid: // pid；	类型：int等类型；
			viewQueryObj.form = append(viewQueryObj.form, `<el-form-item prop="`+v.FieldRaw+`">
            <my-cascader v-model="queryCommon.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/`+tpl.ModuleDirCaseKebab+`/`+tpl.TableCaseKebab+`/tree' }" :defaultOptions="[{ id: 0, label: t('common.name.allTopLevel') }]" :props="{ checkStrictly: true, emitPath: false }" />
        </el-form-item>`)
		case TypeNameLevel: // level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；
			viewQueryObj.form = append(viewQueryObj.form, `<el-form-item prop="`+v.FieldRaw+`">
            <el-input-number v-model="queryCommon.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" :min="1" :controls="false" />
        </el-form-item>`)
		case TypeNameIdPath: // idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；
			continue
		case TypeNamePasswordSuffix: // password,passwd后缀；		类型：char(32)；
			continue
		case TypeNameSaltSuffix: // salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；
			continue
		case TypeNameNameSuffix: // name,title后缀；	类型：varchar；
		case TypeNameCodeSuffix: // code后缀；	类型：varchar；
		case TypeNameAccountSuffix: // account后缀；	类型：varchar；
		case TypeNamePhoneSuffix: // phone,mobile后缀；	类型：varchar；
		case TypeNameEmailSuffix: // email后缀；	类型：varchar；
		case TypeNameUrlSuffix: // url,link后缀；	类型：varchar；
		case TypeNameIpSuffix: // IP后缀；	类型：varchar；
		case TypeNameIdSuffix: // id后缀；	类型：int等类型；
			apiUrl := tpl.ModuleDirCaseKebab + `/` + gstr.CaseKebab(gstr.SubStr(v.FieldCaseCamelRemove, 0, -2))
			if tpl.Handle.RelIdMap[v.FieldRaw].tpl.Table != `` {
				relIdObj := tpl.Handle.RelIdMap[v.FieldRaw]
				apiUrl = relIdObj.tpl.ModuleDirCaseKebab + `/` + relIdObj.tpl.TableCaseKebab
			}
			if tpl.Handle.RelIdMap[v.FieldRaw].tpl.Handle.Pid.Pid != `` {
				viewQueryObj.form = append(viewQueryObj.form, `<el-form-item prop="`+v.FieldRaw+`">
            <my-cascader v-model="queryCommon.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/`+apiUrl+`/tree' }" :props="{ emitPath: false }" />
        </el-form-item>`)
			} else {
				viewQueryObj.form = append(viewQueryObj.form, `<el-form-item prop="`+v.FieldRaw+`">
            <my-select v-model="queryCommon.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/`+apiUrl+`/list' }" />
        </el-form-item>`)
			}
		case TypeNameSortSuffix, TypeNameSort: // sort,weight等后缀；	类型：int等类型； // sort，且pid,level,idPath|id_path,sort同时存在时（才）有效；	类型：int等类型；
			continue
		case TypeNameStatusSuffix: // status,type,method,pos,position,gender等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
			viewQueryObj.form = append(viewQueryObj.form, `<el-form-item prop="`+v.FieldRaw+`" style="width: 120px">
            <el-select-v2 v-model="queryCommon.data.`+v.FieldRaw+`" :options="tm('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.status.`+v.FieldRaw+`')" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" :clearable="true" />
        </el-form-item>`)
			continue
		case TypeNameIsPrefix: // is_前缀；		类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（停用：0否 1是）
			viewQueryObj.form = append(viewQueryObj.form, `<el-form-item prop="`+v.FieldRaw+`" style="width: 120px">
            <el-select-v2 v-model="queryCommon.data.`+v.FieldRaw+`" :options="tm('common.status.whether')" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" :clearable="true" />
        </el-form-item>`)
		case TypeNameStartPrefix: // start_前缀；	类型：timestamp或datetime或date；
		case TypeNameEndPrefix: // end_前缀；	类型：timestamp或datetime或date；
			if v.FieldType != TypeDate {
				viewQueryItemObj.elDatePicker.defaultTime = ` :default-time="new Date(2000, 0, 1, 23, 59, 59)"`
			}
		case TypeNameRemarkSuffix: // remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器
			continue
		case TypeNameImageSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text
			continue
		case TypeNameVideoSuffix: // video,video_list,videoList,video_arr,videoArr等后缀；		类型：单视频varchar，多视频json或text
			continue
		case TypeNameArrSuffix: // list,arr等后缀；	类型：json或text；
			continue
		}
		/*--------根据字段命名类型处理 结束--------*/

		/*--------根据字段数据类型处理（注意：这里是字段命名类型处理的后续操作，改动需考虑兼容） 开始--------*/
		switch v.FieldType {
		case TypeInt: // `int等类型`
			/* viewQueryObj.form += `
			<el-form-item prop="` + v.FieldRaw + `">
				<el-input-number v-model="queryCommon.data.` + v.FieldRaw + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')" :controls="false" />
			</el-form-item>` */
		case TypeIntU: // `int等类型（unsigned）`
			/* viewQueryObj.form += `
			<el-form-item prop="` + v.FieldRaw + `">
				<el-input-number v-model="queryCommon.data.` + v.FieldRaw + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')" :min="0" :controls="false" />
			</el-form-item>` */
		case TypeFloat: // `float等类型`
			/* viewQueryObj.form += `
			<el-form-item prop="` + v.FieldRaw + `">
				<el-input-number v-model="queryCommon.data.` + v.FieldRaw + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')" :precision="` + v.FieldLimitFloat[1] + `" :controls="false" />
			</el-form-item>` */
		case TypeFloatU: // `float等类型（unsigned）`
			/* viewQueryObj.form += `
			<el-form-item prop="` + v.FieldRaw + `">
				<el-input-number v-model="queryCommon.data.` + v.FieldRaw + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')" :min="0" :precision="` + v.FieldLimitFloat[1] + `" :controls="false" />
			</el-form-item>` */
		case TypeVarchar: // `varchar类型`
			viewQueryObj.form = append(viewQueryObj.form, `<el-form-item prop="`+v.FieldRaw+`">
            <el-input v-model="queryCommon.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" maxlength="`+v.FieldLimitStr+`" :clearable="true" />
        </el-form-item>`)
		case TypeChar: // `char类型`
			viewQueryObj.form = append(viewQueryObj.form, `<el-form-item prop="`+v.FieldRaw+`">
            <el-input v-model="queryCommon.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" minlength="`+v.FieldLimitStr+`" maxlength="`+v.FieldLimitStr+`" :clearable="true" />
        </el-form-item>`)
		case TypeText: // `text类型`
		case TypeJson: // `json类型`
		case TypeTimestamp, TypeDatetime: // `timestamp类型` // `datetime类型`
			viewQueryObj.form = append(viewQueryObj.form, `<el-form-item prop="`+v.FieldRaw+`">
            <el-date-picker v-model="queryCommon.data.`+v.FieldRaw+`" type="datetime" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" format="YYYY-MM-DD HH:mm:ss" value-format="YYYY-MM-DD HH:mm:ss"`+viewQueryItemObj.elDatePicker.defaultTime+` />
        </el-form-item>`)
		case TypeDate: // `date类型`
			viewQueryObj.form = append(viewQueryObj.form, `<el-form-item prop="`+v.FieldRaw+`">
            <el-date-picker v-model="queryCommon.data.`+v.FieldRaw+`" type="date" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" format="YYYY-MM-DD" value-format="YYYY-MM-DD" />
        </el-form-item>`)
		default:
			viewQueryObj.form = append(viewQueryObj.form, `<el-form-item prop="`+v.FieldRaw+`">
            <el-input v-model="queryCommon.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" :clearable="true" />
        </el-form-item>`)
		}
		/*--------根据字段数据类型处理（注意：这里是字段命名类型处理的后续操作，改动需考虑兼容） 结束--------*/
	}

	tplView := `<script setup lang="tsx">
import dayjs from 'dayjs'

const { t, tm } = useI18n()

const queryCommon = inject('queryCommon') as { data: { [propName: string]: any } }
queryCommon.data = {
    ...queryCommon.data,` + gstr.Join(append([]string{``}, viewQueryObj.dataInit...), `
    `) + `
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
        </el-form-item>` + gstr.Join(append([]string{``}, viewQueryObj.form...), `
        `) + `
        <el-form-item>
            <el-button type="primary" @click="queryForm.submit" :loading="queryForm.loading"> <autoicon-ep-search />{{ t('common.query') }} </el-button>
            <el-button type="info" @click="queryForm.reset"> <autoicon-ep-circle-close />{{ t('common.reset') }} </el-button>
        </el-form-item>
    </el-form>
</template>
`

	saveFile := gfile.SelfDir() + `/../view/` + myGenThis.option.SceneCode + `/src/views/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/Query.vue`
	gfile.PutContents(saveFile, tplView)
}

// 视图模板Save生成
func (myGenThis *myGen) genViewSave() {
	tpl := myGenThis.tpl

	if !(myGenThis.option.IsCreate || myGenThis.option.IsUpdate) {
		return
	}

	type viewSave struct {
		importModule   []string
		paramHandle    []string
		dataInitBefore []string
		dataInitAfter  []string
		rule           []string
		form           []string
		formHandle     []string
	}
	viewSaveObj := viewSave{
		importModule:   []string{},
		paramHandle:    []string{},
		dataInitBefore: []string{},
		dataInitAfter:  []string{},
		rule:           []string{},
		form:           []string{},
		formHandle:     []string{},
	}

	type viewSaveItem struct {
		rule         []string
		required     string
		elDatePicker struct {
			defaultTime string
		}
	}
	for _, v := range tpl.FieldList {
		viewSaveItemObj := viewSaveItem{
			rule: []string{},
		}

		/*--------根据字段命名类型处理 开始--------*/
		switch v.FieldTypeName {
		case TypeNameDeleted: // 软删除字段
			continue
		case TypeNameUpdated: // 更新时间字段
			continue
		case TypeNameCreated: // 创建时间字段
			continue
		case TypeNamePri: // 主键
		case TypeNamePriAutoInc: // 主键（自增）
			continue
		case TypeNamePid: // pid；	类型：int等类型；
			viewSaveObj.paramHandle = append(viewSaveObj.paramHandle, `param.`+v.FieldRaw+` === undefined ? param.`+v.FieldRaw+` = 0 : null`)
			viewSaveObj.rule = append(viewSaveObj.rule, v.FieldRaw+`: [
            { type: 'integer', min: 0, trigger: 'change', message: t('validation.select') },
        ],`)
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <my-cascader v-model="saveForm.data.`+v.FieldRaw+`" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/`+tpl.ModuleDirCaseKebab+`/`+tpl.TableCaseKebab+`/tree', param: { filter: { excIdArr: saveForm.data.idArr } } }" :props="{ checkStrictly: true, emitPath: false }" />
                </el-form-item>`)
			continue
		case TypeNameLevel: // level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；
			continue
		case TypeNameIdPath: // idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；
			continue
		case TypeNamePasswordSuffix: // password,passwd后缀；		类型：char(32)；
			if !garray.NewStrArrayFrom(viewSaveObj.importModule).Contains(`import md5 from 'js-md5'`) {
				viewSaveObj.importModule = append(viewSaveObj.importModule, `import md5 from 'js-md5'`)
			}
			viewSaveObj.paramHandle = append(viewSaveObj.paramHandle, `param.`+v.FieldRaw+` ? param.`+v.FieldRaw+` = md5(param.`+v.FieldRaw+`) : delete param.`+v.FieldRaw)
			viewSaveObj.rule = append(viewSaveObj.rule, v.FieldRaw+`: [
            { type: 'string', required: computed((): boolean => { return saveForm.data.idArr?.length ? false : true; }), min: 6, max: 20, trigger: 'blur', message: t('validation.between.string', { min: 6, max: 20 }) },
        ],`)
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-input v-model="saveForm.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" minlength="6" maxlength="20" :show-word-limit="true" :clearable="true" :show-password="true" style="max-width: 250px" />
                    <label v-if="saveForm.data.idArr?.length">
                        <el-alert :title="t('common.tip.notRequired')" type="info" :show-icon="true" :closable="false" />
                    </label>
                </el-form-item>`)
			continue
		case TypeNameSaltSuffix: // salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；
			continue
		case TypeNameNameSuffix: // name,title后缀；	类型：varchar；
			if len(tpl.Handle.LabelList) > 0 && gstr.CaseCamel(tpl.Handle.LabelList[0]) == v.FieldCaseCamel {
				viewSaveItemObj.required = ` required: true,`
			}
			// 去掉该验证规则。有时会用到特殊符号
			// viewSaveItemObj.rule = append(viewSaveItemObj.rule, `{ pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') },`)
		case TypeNameCodeSuffix: // code后缀；	类型：varchar；
			viewSaveItemObj.rule = append(viewSaveItemObj.rule, `{ pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') },`)
		case TypeNameAccountSuffix: // account后缀；	类型：varchar；
			viewSaveItemObj.rule = append(viewSaveItemObj.rule, `{ pattern: /^(?!\d*$)[\p{L}\p{M}\p{N}_]+$/u, trigger: 'blur', message: t('validation.account') },`)
		case TypeNamePhoneSuffix: // phone,mobile后缀；	类型：varchar；
			viewSaveItemObj.rule = append(viewSaveItemObj.rule, `{ pattern: /^1[3-9]\d{9}$/, trigger: 'blur', message: t('validation.phone') },`)
		case TypeNameEmailSuffix: // email后缀；	类型：varchar；
			viewSaveItemObj.rule = append(viewSaveItemObj.rule, `{ type: 'email', trigger: 'blur', message: t('validation.email') },`)
		case TypeNameUrlSuffix: // url,link后缀；	类型：varchar；
			viewSaveItemObj.rule = append(viewSaveItemObj.rule, `{ type: 'url', trigger: 'blur', message: t('validation.url') },`)
		case TypeNameIpSuffix: // IP后缀；	类型：varchar；
		case TypeNameIdSuffix: // id后缀；	类型：int等类型；
			apiUrl := tpl.ModuleDirCaseKebab + `/` + gstr.CaseKebab(gstr.SubStr(v.FieldCaseCamelRemove, 0, -2))
			if tpl.Handle.RelIdMap[v.FieldRaw].tpl.Table != `` {
				relIdObj := tpl.Handle.RelIdMap[v.FieldRaw]
				apiUrl = relIdObj.tpl.ModuleDirCaseKebab + `/` + relIdObj.tpl.TableCaseKebab
			}
			viewSaveObj.paramHandle = append(viewSaveObj.paramHandle, `param.`+v.FieldRaw+` === undefined ? param.`+v.FieldRaw+` = 0 : null`)
			viewSaveObj.rule = append(viewSaveObj.rule, v.FieldRaw+`: [
            { type: 'integer', /* required: true, */ min: 1, trigger: 'change', message: t('validation.select') },
        ],`)
			if tpl.Handle.RelIdMap[v.FieldRaw].tpl.Handle.Pid.Pid != `` {
				viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <my-cascader v-model="saveForm.data.`+v.FieldRaw+`" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/`+apiUrl+`/tree' }" :props="{ emitPath: false }" />
                </el-form-item>`)
			} else {
				viewSaveObj.dataInitAfter = append(viewSaveObj.dataInitAfter, v.FieldRaw+`: saveCommon.data.`+v.FieldRaw+` ? saveCommon.data.`+v.FieldRaw+` : undefined,`)
				viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <my-select v-model="saveForm.data.`+v.FieldRaw+`" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/`+apiUrl+`/list' }" />
                </el-form-item>`)
			}
			continue
		case TypeNameSortSuffix, TypeNameSort: // sort,weight等后缀；	类型：int等类型； // sort，且pid,level,idPath|id_path,sort同时存在时（才）有效；	类型：int等类型；
			defaultVal := gconv.Int(v.Default)
			if defaultVal != 0 {
				viewSaveObj.dataInitBefore = append(viewSaveObj.dataInitBefore, v.FieldRaw+`: `+gconv.String(defaultVal)+`,`)
			}
			viewSaveObj.rule = append(viewSaveObj.rule, v.FieldRaw+`: [
            { type: 'integer', min: 0, max: 100, trigger: 'change', message: t('validation.between.number', { min: 0, max: 100 }) },
        ],`)
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-input-number v-model="saveForm.data.`+v.FieldRaw+`" :precision="0" :min="0" :max="100" :step="1" :step-strictly="true" controls-position="right" :value-on-clear="`+gconv.String(defaultVal)+`" />
                    <label>
                        <el-alert :title="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.tip.`+v.FieldRaw+`')" type="info" :show-icon="true" :closable="false" />
                    </label>
                </el-form-item>`)
			continue
		case TypeNameStatusSuffix: // status,type,method,pos,position,gender等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
			defaultVal := gconv.String(v.Default)
			if defaultVal == `` {
				defaultVal = v.StatusList[0][0]
			}
			if garray.NewFrom([]interface{}{TypeVarchar, TypeChar}).Contains(v.FieldType) {
				viewSaveObj.dataInitBefore = append(viewSaveObj.dataInitBefore, v.FieldRaw+`: '`+defaultVal+`',`)
			} else {
				viewSaveObj.dataInitBefore = append(viewSaveObj.dataInitBefore, v.FieldRaw+`: `+defaultVal+`,`)
			}

			viewSaveObj.rule = append(viewSaveObj.rule, v.FieldRaw+`: [
            { type: 'enum', enum: (tm('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.status.`+v.FieldRaw+`') as any).map((item: any) => item.value), trigger: 'change', message: t('validation.select') },
        ],`)

			//超过5个状态用select组件，小于5个用radio组件
			if len(v.StatusList) > 5 {
				viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-select-v2 v-model="saveForm.data.`+v.FieldRaw+`" :options="tm('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.status.`+v.FieldRaw+`')" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" :clearable="false" />
                </el-form-item>`)
			} else {
				viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-radio-group v-model="saveForm.data.`+v.FieldRaw+`">
                        <el-radio v-for="(item, index) in (tm('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.status.`+v.FieldRaw+`') as any)" :key="index" :label="item.value">
                            {{ item.label }}
                        </el-radio>
                    </el-radio-group>
                </el-form-item>`)
			}
			continue
		case TypeNameIsPrefix: // is_前缀；		类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（停用：0否 1是）
			defaultVal := gconv.Int(v.Default)
			if defaultVal != 0 {
				viewSaveObj.dataInitBefore = append(viewSaveObj.dataInitBefore, v.FieldRaw+`: `+gconv.String(defaultVal)+`,`)
			}
			viewSaveObj.rule = append(viewSaveObj.rule, v.FieldRaw+`: [
            { type: 'enum', enum: (tm('common.status.whether') as any).map((item: any) => item.value), trigger: 'change', message: t('validation.select') },
        ],`)
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-switch v-model="saveForm.data.`+v.FieldRaw+`" :active-value="1" :inactive-value="0" :inline-prompt="true" :active-text="t('common.yes')" :inactive-text="t('common.no')" style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success);" />
                </el-form-item>`)
			continue
		case TypeNameStartPrefix: // start_前缀；	类型：timestamp或datetime或date；
		case TypeNameEndPrefix: // end_前缀；	类型：timestamp或datetime或date；
			if v.FieldType != TypeDate {
				viewSaveItemObj.elDatePicker.defaultTime = ` :default-time="new Date(2000, 0, 1, 23, 59, 59)"`
			}
		case TypeNameRemarkSuffix: // remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器
			if v.FieldType == TypeVarchar {
				viewSaveObj.rule = append(viewSaveObj.rule, v.FieldRaw+`: [
            { type: 'string', max: `+v.FieldLimitStr+`, trigger: 'blur', message: t('validation.max.string', { max: `+v.FieldLimitStr+` }) },
        ],`)
				viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-input v-model="saveForm.data.`+v.FieldRaw+`" type="textarea" :autosize="{ minRows: 3 }" maxlength="`+v.FieldLimitStr+`" :show-word-limit="true" />
                </el-form-item>`)
				continue
			}
		case TypeNameImageSuffix, TypeNameVideoSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text // video,video_list,videoList,video_arr,videoArr等后缀；		类型：单视频varchar，多视频json或text
			multipleStr := ``
			if v.FieldType == TypeVarchar {
				viewSaveObj.rule = append(viewSaveObj.rule, v.FieldRaw+`: [
            { type: 'string', max: `+v.FieldLimitStr+`, trigger: 'blur', message: t('validation.max.string', { max: `+v.FieldLimitStr+` }) },
            { type: 'url', trigger: 'change', message: t('validation.upload') },
        ],`)
			} else {
				multipleStr = ` :multiple="true"`
				requiredStr := ``
				if !v.IsNull {
					requiredStr = ` required: true,`
				}
				viewSaveObj.rule = append(viewSaveObj.rule, v.FieldRaw+`: [
            { type: 'array',`+requiredStr+` trigger: 'change', message: t('validation.upload'), defaultField: { type: 'url', message: t('validation.url') } },
            // { type: 'array',`+requiredStr+` max: 10, trigger: 'change', message: t('validation.max.upload', { max: 10 }), defaultField: { type: 'url', message: t('validation.url') } },
        ],`)
			}

			acceptStr := ` accept="image/*"`
			if v.FieldTypeName == TypeNameVideoSuffix {
				acceptStr = ` accept="video/*" :isImage="false"`
			}
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <my-upload v-model="saveForm.data.`+v.FieldRaw+`"`+multipleStr+acceptStr+` />
                </el-form-item>`)
			continue
		case TypeNameArrSuffix: // list,arr等后缀；	类型：json或text；
			viewSaveObj.dataInitBefore = append(viewSaveObj.dataInitBefore, v.FieldRaw+`: [],`)
			requiredStr := ``
			if !v.IsNull {
				requiredStr = ` required: true,`
			}
			viewSaveObj.rule = append(viewSaveObj.rule, v.FieldRaw+`: [
            { type: 'array',`+requiredStr+` trigger: 'change', message: t('validation.required') },
            // { type: 'array',`+requiredStr+` max: 10, trigger: 'change', message: t('validation.max.array', { max: 10 }), defaultField: { type: 'string', message: t('validation.input') } },
        ],`)
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-tag v-for="(item, index) in saveForm.data.`+v.FieldRaw+`" :type="`+v.FieldRaw+`Handle.tagType[index % `+v.FieldRaw+`Handle.tagType.length]" @close="`+v.FieldRaw+`Handle.delValue(item)" :key="index" :closable="true" style="margin-right: 10px;">
                        {{ item }}
                    </el-tag>
                    <!-- <el-input-number v-if="`+v.FieldRaw+`Handle.visible" :ref="(el: any) => `+v.FieldRaw+`Handle.ref = el" v-model="`+v.FieldRaw+`Handle.value" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" @keyup.enter="`+v.FieldRaw+`Handle.addValue" @blur="`+v.FieldRaw+`Handle.addValue" size="small" style="width: 100px;" :controls="false" /> -->
                    <el-input v-if="`+v.FieldRaw+`Handle.visible" :ref="(el: any) => `+v.FieldRaw+`Handle.ref = el" v-model="`+v.FieldRaw+`Handle.value" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" @keyup.enter="`+v.FieldRaw+`Handle.addValue" @blur="`+v.FieldRaw+`Handle.addValue" size="small" style="width: 100px;" />
                    <el-button v-else type="primary" size="small" @click="`+v.FieldRaw+`Handle.visibleChange">
                        <autoicon-ep-plus />{{ t('common.add') }}
                    </el-button>
                </el-form-item>`)
			viewSaveObj.formHandle = append(viewSaveObj.formHandle, `const `+v.FieldRaw+`Handle = reactive({
    ref: null as any,
    visible: false,
    value: undefined,
    tagType: tm('config.const.tagType') as string[],
    visibleChange: () => {
        `+v.FieldRaw+`Handle.visible = true
        nextTick(() => {
            `+v.FieldRaw+`Handle.ref?.focus()
        })
    },
    addValue: () => {
        if (`+v.FieldRaw+`Handle.value) {
            saveForm.data.`+v.FieldRaw+`.push(`+v.FieldRaw+`Handle.value)
        }
        `+v.FieldRaw+`Handle.visible = false
        `+v.FieldRaw+`Handle.value = undefined
    },
    delValue: (item: any) => {
        saveForm.data.`+v.FieldRaw+`.splice(saveForm.data.`+v.FieldRaw+`.indexOf(item), 1)
    },
})`)
			continue
		}
		/*--------根据字段命名类型处理 结束--------*/

		/*--------根据字段数据类型处理（注意：这里是字段命名类型处理的后续操作，改动需考虑兼容） 开始--------*/
		switch v.FieldType {
		case TypeInt: // `int等类型`
			defaultVal := gconv.Int(v.Default)
			if defaultVal != 0 {
				viewSaveObj.dataInitBefore = append(viewSaveObj.dataInitBefore, v.FieldRaw+`: `+gconv.String(defaultVal)+`,`)
			}
			viewSaveItemObj.rule = append(viewSaveItemObj.rule, `{ type: 'integer', trigger: 'change', message: t('validation.input') },`)
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-input-number v-model="saveForm.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" :controls="false" :value-on-clear="`+gconv.String(defaultVal)+`" />
                </el-form-item>`)
		case TypeIntU: // `int等类型（unsigned）`
			defaultVal := gconv.Uint(v.Default)
			if defaultVal != 0 {
				viewSaveObj.dataInitBefore = append(viewSaveObj.dataInitBefore, v.FieldRaw+`: `+gconv.String(defaultVal)+`,`)
			}
			viewSaveItemObj.rule = append(viewSaveItemObj.rule, `{ type: 'integer', min: 0, trigger: 'change', message: t('validation.min.number', { min: 0 }) },`)
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-input-number v-model="saveForm.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" :min="0" :controls="false" :value-on-clear="`+gconv.String(defaultVal)+`" />
                </el-form-item>`)
		case TypeFloat: // `float等类型`
			defaultVal := gconv.Float64(v.Default)
			if defaultVal != 0 {
				viewSaveObj.dataInitBefore = append(viewSaveObj.dataInitBefore, v.FieldRaw+`: `+gconv.String(defaultVal)+`,`)
			}
			viewSaveItemObj.rule = append(viewSaveItemObj.rule, `{ type: 'number'/* 'float' */, trigger: 'change', message: t('validation.input') },    // 类型float值为0时验证不能通过`)
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-input-number v-model="saveForm.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" :precision="`+v.FieldLimitFloat[1]+`" :controls="false" :value-on-clear="`+gconv.String(defaultVal)+`" />
                </el-form-item>`)
		case TypeFloatU: // `float等类型（unsigned）`
			defaultVal := gconv.Float64(v.Default)
			if defaultVal != 0 {
				viewSaveObj.dataInitBefore = append(viewSaveObj.dataInitBefore, v.FieldRaw+`: `+gconv.String(defaultVal)+`,`)
			}
			viewSaveItemObj.rule = append(viewSaveItemObj.rule, `{ type: 'number'/* 'float' */, min: 0, trigger: 'change', message: t('validation.min.number', { min: 0 }) },    // 类型float值为0时验证不能通过`)
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-input-number v-model="saveForm.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" :min="0" :precision="`+v.FieldLimitFloat[1]+`" :controls="false" :value-on-clear="`+gconv.String(defaultVal)+`" />
                </el-form-item>`)
		case TypeVarchar: // `varchar类型`
			if v.IndexRaw == `UNI` && !v.IsNull {
				viewSaveItemObj.required = ` required: true,`
			}
			viewSaveItemObj.rule = append([]string{`{ type: 'string',` + viewSaveItemObj.required + ` max: ` + v.FieldLimitStr + `, trigger: 'blur', message: t('validation.max.string', { max: ` + v.FieldLimitStr + ` }) },`}, viewSaveItemObj.rule...)

			if v.IndexRaw == `UNI` {
				viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-input v-model="saveForm.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" maxlength="`+v.FieldLimitStr+`" :show-word-limit="true" :clearable="true" style="max-width: 250px" />
                    <label>
                        <el-alert :title="t('common.tip.notDuplicate')" type="info" :show-icon="true" :closable="false" />
                    </label>
                </el-form-item>`)
			} else {
				viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-input v-model="saveForm.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" maxlength="`+v.FieldLimitStr+`" :show-word-limit="true" :clearable="true" />
                </el-form-item>`)
			}
		case TypeChar: // `char类型`
			if v.IndexRaw == `UNI` && !v.IsNull {
				viewSaveItemObj.required = ` required: true,`
			}
			viewSaveItemObj.rule = append([]string{`{ type: 'string',` + viewSaveItemObj.required + ` len: ` + v.FieldLimitStr + `, trigger: 'blur', message: t('validation.size.string', { size: ` + v.FieldLimitStr + ` }) },`}, viewSaveItemObj.rule...)

			if v.IndexRaw == `UNI` {
				viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-input v-model="saveForm.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" minlength="`+v.FieldLimitStr+`" maxlength="`+v.FieldLimitStr+`" :show-word-limit="true" :clearable="true" />
                </el-form-item>`)
			} else {
				viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-input v-model="saveForm.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" minlength="`+v.FieldLimitStr+`" maxlength="`+v.FieldLimitStr+`" :show-word-limit="true" :clearable="true" style="max-width: 250px" />
                    <label>
                        <el-alert :title="t('common.tip.notDuplicate')" type="info" :show-icon="true" :closable="false" />
                    </label>
                </el-form-item>`)
			}
		case TypeText: // `text类型`
			viewSaveItemObj.rule = append([]string{`{ type: 'string', trigger: 'blur', message: t('validation.input') },`}, viewSaveItemObj.rule...)
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <my-editor v-model="saveForm.data.`+v.FieldRaw+`" />
                </el-form-item>`)
		case TypeJson: // `json类型`
			if !v.IsNull {
				viewSaveItemObj.required = `
                required: true,`
			}
			viewSaveItemObj.rule = append(viewSaveItemObj.rule, `{
                type: 'object',`+viewSaveItemObj.required+`
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
            },`)
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-alert :title="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.tip.`+v.FieldRaw+`')" type="info" :show-icon="true" :closable="false" />
                    <el-input v-model="saveForm.data.`+v.FieldRaw+`" type="textarea" :autosize="{ minRows: 3 }" />
                </el-form-item>`)
		case TypeTimestamp, TypeDatetime: // `timestamp类型` // `datetime类型`
			if !v.IsNull && gconv.String(v.Default) == `` {
				viewSaveItemObj.required = ` required: true,`
			}
			viewSaveItemObj.rule = append([]string{`{ type: 'string',` + viewSaveItemObj.required + ` trigger: 'change', message: t('validation.select') },`}, viewSaveItemObj.rule...)
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-date-picker v-model="saveForm.data.`+v.FieldRaw+`" type="datetime" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" format="YYYY-MM-DD HH:mm:ss" value-format="YYYY-MM-DD HH:mm:ss"`+viewSaveItemObj.elDatePicker.defaultTime+` />
                </el-form-item>`)
		case TypeDate: // `date类型`
			if !v.IsNull && gconv.String(v.Default) == `` {
				viewSaveItemObj.required = ` required: true,`
			}
			viewSaveItemObj.rule = append([]string{`{ type: 'string',` + viewSaveItemObj.required + ` trigger: 'change', message: t('validation.select') },`}, viewSaveItemObj.rule...)
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-date-picker v-model="saveForm.data.`+v.FieldRaw+`" type="date" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" format="YYYY-MM-DD" value-format="YYYY-MM-DD" />
                </el-form-item>`)
		default:
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-input v-model="saveForm.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" :clearable="true" />
                </el-form-item>`)
		}
		/*--------根据字段数据类型处理（注意：这里是字段命名类型处理的后续操作，改动需考虑兼容） 结束--------*/

		if len(viewSaveItemObj.rule) > 0 {
			viewSaveObj.rule = append(viewSaveObj.rule, v.FieldRaw+`: [`+gstr.Join(append([]string{``}, viewSaveItemObj.rule...), `
			`)+`
        ],`)
		} else {
			viewSaveObj.rule = append(viewSaveObj.rule, v.FieldRaw+`: [],`)
		}
	}

	tplView := `<script setup lang="tsx">` + gstr.Join(append([]string{``}, viewSaveObj.importModule...), `
`) + `
const { t, tm } = useI18n()

const saveCommon = inject('saveCommon') as { visible: boolean; title: string; data: { [propName: string]: any } }
const listCommon = inject('listCommon') as { ref: any }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {` + gstr.Join(append([]string{``}, viewSaveObj.dataInitBefore...), `
        `) + `
        ...saveCommon.data,` + gstr.Join(append([]string{``}, viewSaveObj.dataInitAfter...), `
        `) + `
    } as { [propName: string]: any },
    rules: {` + gstr.Join(append([]string{``}, viewSaveObj.rule...), `
        `) + `
    } as any,
    submit: () => {
        saveForm.ref.validate(async (valid: boolean) => {
            if (!valid) {
                return false
            }
            saveForm.loading = true
            const param = removeEmptyOfObj(saveForm.data)` + gstr.Join(append([]string{``}, viewSaveObj.paramHandle...), `
            `) + `
            try {
                if (param?.idArr?.length > 0) {
                    await request(t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/update', param, true)
                } else {
                    await request(t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/create', param, true)
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
})` + gstr.Join(append([]string{``}, viewSaveObj.formHandle...), `

`) + `
</script>

<template>
    <el-drawer class="save-drawer" :ref="(el: any) => saveDrawer.ref = el" v-model="saveCommon.visible" :title="saveCommon.title" :size="saveDrawer.size" :before-close="saveDrawer.beforeClose">
        <el-scrollbar>
            <el-form :ref="(el: any) => saveForm.ref = el" :model="saveForm.data" :rules="saveForm.rules" label-width="auto" :status-icon="true" :scroll-to-error="true">` + gstr.Join(append([]string{``}, viewSaveObj.form...), `
                `) + `
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

	saveFile := gfile.SelfDir() + `/../view/` + myGenThis.option.SceneCode + `/src/views/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/Save.vue`
	gfile.PutContents(saveFile, tplView)
}

// 视图模板I18n生成
func (myGenThis *myGen) genViewI18n() {
	tpl := myGenThis.tpl

	type viewI18n struct {
		name   []string
		status []string
		tip    []string
	}
	viewI18nObj := viewI18n{
		name:   []string{},
		status: []string{},
		tip:    []string{},
	}
	for _, v := range tpl.FieldList {
		/*--------根据字段命名类型处理 开始--------*/
		switch v.FieldTypeName {
		case TypeNameDeleted: // 软删除字段
			continue
		case TypeNameUpdated: // 更新时间字段
			continue
		case TypeNameCreated: // 创建时间字段
			continue
		case TypeNamePri: // 主键
		case TypeNamePriAutoInc: // 主键（自增）
			continue
		case TypeNamePid: // pid；	类型：int等类型；
			viewI18nObj.name = append(viewI18nObj.name, v.FieldRaw+`: '父级',`)
			continue
		case TypeNameLevel: // level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；
		case TypeNameIdPath: // idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；
		case TypeNamePasswordSuffix: // password,passwd后缀；		类型：char(32)；
		case TypeNameSaltSuffix: // salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；
			continue
		case TypeNameNameSuffix: // name,title后缀；	类型：varchar；
		case TypeNameCodeSuffix: // code后缀；	类型：varchar；
		case TypeNameAccountSuffix: // account后缀；	类型：varchar；
		case TypeNamePhoneSuffix: // phone,mobile后缀；	类型：varchar；
		case TypeNameEmailSuffix: // email后缀；	类型：varchar；
		case TypeNameUrlSuffix: // url,link后缀；	类型：varchar；
		case TypeNameIpSuffix: // IP后缀；	类型：varchar；
		case TypeNameIdSuffix: // id后缀；	类型：int等类型；
			relIdObj := tpl.Handle.RelIdMap[v.FieldRaw]
			if relIdObj.tpl.Table != `` && !relIdObj.IsRedundName {
				viewI18nObj.name = append(viewI18nObj.name, v.FieldRaw+`: '`+relIdObj.FieldName+`',`)
				continue
			}
		case TypeNameSortSuffix, TypeNameSort: // sort,weight等后缀；	类型：int等类型； // sort，且pid,level,idPath|id_path,sort同时存在时（才）有效；	类型：int等类型；
			viewI18nObj.tip = append(viewI18nObj.tip, v.FieldRaw+`: '`+v.FieldTip+`',`)
		case TypeNameStatusSuffix: // status,type,method,pos,position,gender等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
			statusItem := []string{}
			if garray.NewFrom([]interface{}{TypeVarchar, TypeChar}).Contains(v.FieldType) {
				for _, status := range v.StatusList {
					statusItem = append(statusItem, `{ value: '`+status[0]+`', label: '`+status[1]+`' },`)
				}
			} else {
				for _, status := range v.StatusList {
					statusItem = append(statusItem, `{ value: `+status[0]+`, label: '`+status[1]+`' },`)
				}
			}
			viewI18nObj.status = append(viewI18nObj.status, v.FieldRaw+`: [`+gstr.Join(append([]string{``}, statusItem...), `
            `)+`
        ],`)
		case TypeNameIsPrefix: // is_前缀；		类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（停用：0否 1是）
		case TypeNameStartPrefix: // start_前缀；	类型：timestamp或datetime或date；
		case TypeNameEndPrefix: // end_前缀；	类型：timestamp或datetime或date；
		case TypeNameRemarkSuffix: // remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器
		case TypeNameImageSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text
		case TypeNameVideoSuffix: // video,video_list,videoList,video_arr,videoArr等后缀；		类型：单视频varchar，多视频json或text
		case TypeNameArrSuffix: // list,arr等后缀；	类型：json或text；
		}
		/*--------根据字段命名类型处理 结束--------*/

		/*--------根据字段数据类型处理（注意：这里是字段命名类型处理的后续操作，改动需考虑兼容） 开始--------*/
		switch v.FieldType {
		case TypeInt: // `int等类型`
		case TypeIntU: // `int等类型（unsigned）`
		case TypeFloat: // `float等类型`
		case TypeFloatU: // `float等类型（unsigned）`
		case TypeVarchar: // `varchar类型`
		case TypeChar: // `char类型`
		case TypeText: // `text类型`
		case TypeJson: // `json类型`
			viewI18nObj.tip = append(viewI18nObj.tip, v.FieldRaw+`: '`+v.FieldTip+`',`)
		case TypeTimestamp: // `timestamp类型`
		case TypeDatetime: // `datetime类型`
		case TypeDate: // `date类型`
		default:
		}
		/*--------根据字段数据类型处理（注意：这里是字段命名类型处理的后续操作，改动需考虑兼容） 结束--------*/
		viewI18nObj.name = append(viewI18nObj.name, v.FieldRaw+`: '`+v.FieldName+`',`)
	}

	tplView := `export default {
    name: {` + gstr.Join(append([]string{``}, viewI18nObj.name...), `
        `) + `
    },
    status: {` + gstr.Join(append([]string{``}, viewI18nObj.status...), `
        `) + `
    },
    tip: {` + gstr.Join(append([]string{``}, viewI18nObj.tip...), `
        `) + `
    },
}
`

	saveFile := gfile.SelfDir() + `/../view/` + myGenThis.option.SceneCode + `/src/i18n/language/zh-cn/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `.ts`
	gfile.PutContents(saveFile, tplView)
}

// 前端路由生成
func (myGenThis *myGen) genViewRouter() {
	tpl := myGenThis.tpl

	saveFile := gfile.SelfDir() + `/../view/` + myGenThis.option.SceneCode + `/src/router/index.ts`
	tplViewRouter := gfile.GetContents(saveFile)

	path := `/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab
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

	myGenThis.genMenu(myGenThis.sceneInfo[daoAuth.Scene.Columns().SceneId].Uint(), path, myGenThis.option.CommonName, tpl.TableCaseCamel) // 数据库权限菜单处理
}

// 自动生成操作权限
func (myGenThis *myGen) genAction(sceneId uint, actionCode string, actionName string) {
	ctx := myGenThis.ctx

	actionName = gstr.Replace(actionName, `/`, `-`)

	idVar, _ := daoAuth.Action.CtxDaoModel(ctx).Filter(daoAuth.Action.Columns().ActionCode, actionCode).Value(daoAuth.Action.PrimaryKey())
	id := idVar.Int64()
	if id == 0 {
		id, _ = daoAuth.Action.CtxDaoModel(ctx).HookInsert(map[string]interface{}{
			daoAuth.Action.Columns().ActionCode: actionCode,
			daoAuth.Action.Columns().ActionName: actionName,
		}).InsertAndGetId()
	} else {
		daoAuth.Action.CtxDaoModel(ctx).Filter(daoAuth.Action.PrimaryKey(), id).HookUpdate(g.Map{daoAuth.Action.Columns().ActionName: actionName}).Update()
	}
	daoAuth.ActionRelToScene.CtxDaoModel(ctx).Data(map[string]interface{}{
		daoAuth.ActionRelToScene.Columns().ActionId: id,
		daoAuth.ActionRelToScene.Columns().SceneId:  sceneId,
	}).Save()
}

// 自动生成菜单
func (myGenThis *myGen) genMenu(sceneId uint, menuUrl string, menuName string, menuNameOfEn string) {
	ctx := myGenThis.ctx

	menuNameArr := gstr.Split(menuName, `/`)

	var pid int64 = 0
	for _, v := range menuNameArr[:len(menuNameArr)-1] {
		pidVar, _ := daoAuth.Menu.CtxDaoModel(ctx).Filters(g.Map{
			daoAuth.Menu.Columns().SceneId:  sceneId,
			daoAuth.Menu.Columns().MenuName: v,
		}).Value(daoAuth.Menu.PrimaryKey())
		if pidVar.Uint() == 0 {
			pid, _ = daoAuth.Menu.CtxDaoModel(ctx).HookInsert(g.Map{
				daoAuth.Menu.Columns().SceneId:   sceneId,
				daoAuth.Menu.Columns().Pid:       pid,
				daoAuth.Menu.Columns().MenuName:  v,
				daoAuth.Menu.Columns().MenuIcon:  `autoicon-ep-link`,
				daoAuth.Menu.Columns().MenuUrl:   ``,
				daoAuth.Menu.Columns().ExtraData: `{"i18n": {"title": {"en": "", "zh-cn": "` + v + `"}}}`,
			}).InsertAndGetId()
		} else {
			pid = pidVar.Int64()
		}
	}

	menuName = menuNameArr[len(menuNameArr)-1]
	idVar, _ := daoAuth.Menu.CtxDaoModel(ctx).Filters(g.Map{
		daoAuth.Menu.Columns().SceneId: sceneId,
		daoAuth.Menu.Columns().MenuUrl: menuUrl,
	}).Value(daoAuth.Menu.PrimaryKey())
	id := idVar.Uint()
	if id == 0 {
		daoAuth.Menu.CtxDaoModel(ctx).HookInsert(g.Map{
			daoAuth.Menu.Columns().SceneId:   sceneId,
			daoAuth.Menu.Columns().Pid:       pid,
			daoAuth.Menu.Columns().MenuName:  menuName,
			daoAuth.Menu.Columns().MenuIcon:  `autoicon-ep-link`,
			daoAuth.Menu.Columns().MenuUrl:   menuUrl,
			daoAuth.Menu.Columns().ExtraData: `{"i18n": {"title": {"en": "` + menuNameOfEn + `", "zh-cn": "` + menuName + `"}}}`,
		}).Insert()
	} else {
		daoAuth.Menu.CtxDaoModel(ctx).Filter(daoAuth.Menu.PrimaryKey(), id).
			SetIdArr().
			HookUpdate(g.Map{
				daoAuth.Menu.Columns().MenuName:  menuName,
				daoAuth.Menu.Columns().Pid:       pid,
				daoAuth.Menu.Columns().ExtraData: `{"i18n": {"title": {"en": "` + menuNameOfEn + `", "zh-cn": "` + menuName + `"}}}`,
			}).Update()
	}
}
