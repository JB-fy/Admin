package cmd

import (
	"context"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type MyGenOption struct {
	DbGroup      string `c:"dbGroup"`              //db分组
	DbTable      string `c:"dbTable"`              //db表
	RemovePrefix string `c:"removePrefix"`         //要删除的db表前缀。和hcak/config.yaml内的removePrefix一致
	NoList       bool   `c:"noList,default:true" ` //不生成列表接口(0,false,off,no,""为false，其他都为true)
	NoCreate     bool   `c:"noCreate"`             //不生成创建接口(0,false,off,no,""为false，其他都为true)
	NoUpdate     bool   `c:"noUpdate"`             //不生成更新接口(0,false,off,no,""为false，其他都为true)
	NoDelete     bool   `c:"noDelete"`             //不生成删除接口(0,false,off,no,""为false，其他都为true)
}

type MyGenTpl struct {
	RawTableNameCaseCamelLower  string //原始表名（小驼峰）
	TableNameCaseCamelLower     string //去除前缀表名（小驼峰）
	TableNameCaseCamel          string //去除前缀表名（大驼峰）
	TableNameCaseSnake          string //去除前缀表名（蛇形）
	PathSuffixCaseCamelLower    string //路径后缀（小驼峰）
	PathSuffixCaseCamel         string //路径后缀（大驼峰）
	ApiFilterColumn             string //api列表过滤
	ApiSaveColumn               string //api创建更新
	ControllerAlloweFieldAppend string //controller追加字段
	ControllerAlloweFieldDiff   string //controller移除字段
	ViewListColumn              string //view列表字段
	ViewQueryField              string //view查询字段
	ViewSaveField               string //view创建更新字段
}

func MyGenFunc(ctx context.Context, parser *gcmd.Parser) (err error) {
	option := MyGenOptionHandle(ctx, parser)
	tpl := MyGenTplHandle(ctx, option)

	MyGenTplApi(ctx, option, tpl)
	MyGenTplLogic(ctx, option, tpl)
	MyGenTplController(ctx, option, tpl)
	MyGenTplView(ctx, option, tpl)
	return
}

// 参数处理
func MyGenOptionHandle(ctx context.Context, parser *gcmd.Parser) (option *MyGenOption) {
	optionMap := parser.GetOptAll()
	option = &MyGenOption{}
	gconv.Struct(optionMap, option)

	var db gdb.DB
	option.DbGroup = `default`
	if option.DbGroup == `` {
		option.DbGroup = gcmd.Scan("> 请输入db分组,默认(default):\n")
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
		option.DbGroup = gcmd.Scan("> db分组不存在，请重新输入,默认(default):\n")
		if option.DbGroup == `` {
			option.DbGroup = `default`
		}
	}
	tableArrTmp, _ := db.GetArray(ctx, `SHOW TABLES`)
	tableArr := gconv.SliceStr(tableArrTmp)
	option.DbTable = `auth_test_scene`
	if option.DbTable == `` {
		option.DbTable = gcmd.Scan("> 请输入db表:\n")
	}
	for {
		if option.DbTable != `` && garray.NewStrArrayFrom(tableArr).Contains(option.DbTable) {
			break
		}
		option.DbTable = gcmd.Scan("> db表不存在，请重新输入:\n")
	}
	/* _, ok := optionMap[`removePrefix`]
	if !ok {
		option.RemovePrefix = gcmd.Scan("> 请输入要删除的db表前缀,默认(空):\n")
	} */
	option.RemovePrefix = `auth_`
	for {
		if option.RemovePrefix == `` || gstr.Pos(option.DbTable, option.RemovePrefix) == 0 {
			break
		}
		option.RemovePrefix = gcmd.Scan("> 要删除的db表前缀不存在，请重新输入,默认(空):\n")
	}
	noList, ok := optionMap[`noList`]
	if !ok {
		noList = gcmd.Scan("> 是否生成列表接口,默认(yes):\n")
	}
noListEnd:
	for {
		switch noList {
		case ``, `yes`:
			option.NoList = false
			break noListEnd
		case `no`:
			option.NoList = true
			break noListEnd
		default:
			noList = gcmd.Scan("> 输入错误，请重新输入，是否生成列表接口,默认(yes):\n")
		}
	}
	noCreate, ok := optionMap[`noCreate`]
	if !ok {
		noCreate = gcmd.Scan("> 是否生成创建接口,默认(yes):\n")
	}
noCreateEnd:
	for {
		switch noCreate {
		case ``, `yes`:
			option.NoCreate = false
			break noCreateEnd
		case `no`:
			option.NoCreate = true
			break noCreateEnd
		default:
			noCreate = gcmd.Scan("> 输入错误，请重新输入，是否生成创建接口,默认(yes):\n")
		}
	}
	noUpdate, ok := optionMap[`noUpdate`]
	if !ok {
		noUpdate = gcmd.Scan("> 是否生成更新接口,默认(yes):\n")
	}
noUpdateEnd:
	for {
		switch noUpdate {
		case ``, `yes`:
			option.NoUpdate = false
			break noUpdateEnd
		case `no`:
			option.NoUpdate = true
			break noUpdateEnd
		default:
			noUpdate = gcmd.Scan("> 输入错误，请重新输入，是否生成更新接口,默认(yes):\n")
		}
	}
	noDelete, ok := optionMap[`noDelete`]
	if !ok {
		noDelete = gcmd.Scan("> 是否生成删除接口,默认(yes):\n")
	}
noDeleteEnd:
	for {
		switch noDelete {
		case ``, `yes`:
			option.NoDelete = false
			break noDeleteEnd
		case `no`:
			option.NoDelete = true
			break noDeleteEnd
		default:
			noDelete = gcmd.Scan("> 输入错误，请重新输入，是否生成删除接口,默认(yes):\n")
		}
	}
	return
}

// 模板变量处理
func MyGenTplHandle(ctx context.Context, option *MyGenOption) (tpl *MyGenTpl) {
	tpl = &MyGenTpl{
		RawTableNameCaseCamelLower: gstr.CaseCamelLower(option.DbTable),
		TableNameCaseCamel:         gstr.CaseCamel(gstr.Replace(option.DbTable, option.RemovePrefix, ``, 1)),
		PathSuffixCaseCamelLower:   gstr.CaseCamelLower(option.RemovePrefix),
		PathSuffixCaseCamel:        gstr.CaseCamel(option.RemovePrefix),
	}
	tpl.TableNameCaseCamelLower = gstr.CaseCamelLower(tpl.TableNameCaseCamel)
	tpl.TableNameCaseSnake = gstr.CaseSnakeFirstUpper(tpl.TableNameCaseCamel)

	tableColumnList, _ := g.DB(option.DbGroup).GetAll(ctx, `SHOW FULL COLUMNS FROM `+option.DbTable)
	for _, column := range tableColumnList {
		field := column[`Field`].String()
		fieldCaseCamel := gstr.CaseCamel(field)
		switch fieldCaseCamel {
		case `UpdatedAt`, `CreatedAt`, `DeletedAt`: //不处理的字段
		default:
			//允许字段
			if (column[`Key`].String() == `PRI` && column[`Extra`].String() == `auto_increment`) || gstr.ToLower(gstr.CaseCamelLower(field)) == gstr.ToLower(tpl.TableNameCaseCamel+`name`) {
				tpl.ControllerAlloweFieldAppend += "`" + field + "`, "
			}
			//排除字段
			if gstr.ToLower(field) == `password` || gstr.ToLower(field) == `passwd` {
				tpl.ControllerAlloweFieldDiff += "`" + field + "`, "
			}
			//主键
			if column[`Key`].String() == `PRI` && column[`Extra`].String() == `auto_increment` {
				tpl.ApiFilterColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"integer|min:1"` + "` // " + column[`Comment`].String() + "\n"
				continue
			}
			//id后缀
			if gstr.ToLower(gstr.SubStr(field, -2)) == `id` {
				tpl.ApiFilterColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"integer|min:1"` + "` // " + column[`Comment`].String() + "\n"
				tpl.ApiSaveColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"integer|min:1"` + "` // " + column[`Comment`].String() + "\n"
				tpl.ViewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('common.name.{TplPathSuffixCaseCamelLower}.{TplTableNameCaseSnake}.` + field + `'),
		key: '` + field + `',
		align: 'center',
		width: 150,
	},`
				if gstr.ToLower(field) == `pid` {

				} else {
					tpl.ViewQueryField += `
		<ElFormItem prop="` + field + `">
			<MySelect v-model="queryCommon.data.` + field + `" :placeholder="t('common.name.rel.` + field + `')" :api="{ code: '{TplPathSuffixCaseCamelLower}/{TplTableNameCaseCamelLower}/list', param: { field: ['id', 'keyword'] } }" />
		</ElFormItem>`
				}
				continue
			}
			//status后缀
			if gstr.ToLower(gstr.SubStr(field, -6)) == `status` {
				tpl.ApiFilterColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"integer|in:0,1,2"` + "` // " + column[`Comment`].String() + "\n"
				tpl.ApiSaveColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"integer|in:0,1,2"` + "` // " + column[`Comment`].String() + "\n"
				tpl.ViewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('common.name.{TplPathSuffixCaseCamelLower}.{TplTableNameCaseSnake}.` + field + `'),
		key: '` + field + `',
		align: 'center',
		width: 100,
	},`
				continue
			}
			//is前缀
			if gstr.ToLower(gstr.SubStr(field, 0, 2)) == `is` {
				tpl.ApiFilterColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"integer|in:0,1"` + "` // " + column[`Comment`].String() + "\n"
				tpl.ApiSaveColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"integer|in:0,1"` + "` // " + column[`Comment`].String() + "\n"
				tpl.ViewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('common.name.` + field + `'),
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
				if !option.NoUpdate {
					tpl.ViewListColumn += `
					onChange: (val: number) => {
						handleUpdate({
							idArr: [props.rowData.id],
							` + field + `: val
						}).then((res) => {
							props.rowData.` + field + ` = val
						}).catch((error) => { })
					}`
				}
				tpl.ViewListColumn += `
				})
			]
		}
	},`
				continue
			}
			//含有sort和weight字符
			if gstr.Pos(gstr.ToLower(field), `sort`) != -1 || gstr.Pos(gstr.ToLower(field), `weight`) != -1 {
				tpl.ApiFilterColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"integer|between:0,100"` + "` // " + column[`Comment`].String() + "\n"
				tpl.ApiSaveColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"integer|between:0,100"` + "` // " + column[`Comment`].String() + "\n"
				tpl.ViewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('common.name.` + field + `'),
		key: '` + field + `',
		align: 'center',
		width: 100,
		sortable: true,`
				if !option.NoUpdate {
					tpl.ViewListColumn += `
		cellRenderer: (props: any): any => {
			if (props.rowData.edit` + gstr.CaseCamel(field) + `) {
				let currentRef: any
				let currentVal = props.rowData.` + field + `
				return [
					h(ElInputNumber as any, {
						'ref': (el: any) => { currentRef = el; el?.focus() },
						'model-value': currentVal,
						'placeholder': t('common.tip.` + field + `'),
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
				tpl.ViewListColumn += `
	},`
				continue
			}
			//name和code后缀
			if gstr.ToLower(gstr.SubStr(field, -4)) == `name` || gstr.ToLower(gstr.SubStr(field, -4)) == `code` {
				result, _ := gregex.MatchString(`.*\((\d*)\)`, column[`Type`].String())
				tpl.ApiFilterColumn += fieldCaseCamel + ` string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"length:1,` + result[1] + `|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"` + "` // " + column[`Comment`].String() + "\n"
				tpl.ApiSaveColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"length:1,` + result[1] + `|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"` + "` // " + column[`Comment`].String() + "\n"
				tpl.ViewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('common.name.{TplPathSuffixCaseCamelLower}.{TplTableNameCaseSnake}.` + field + `'),
		key: '` + field + `',
		width: 150,
		align: 'center',
	},`
				tpl.ViewQueryField += `
		<ElFormItem prop="` + field + `">
			<ElInput v-model="queryCommon.data.` + field + `" :placeholder="t('common.name.{TplPathSuffixCaseCamelLower}.{TplTableNameCaseSnake}.` + field + `')" :clearable="true" />
		</ElFormItem>,`
				continue
			}
			//ip后缀
			if gstr.ToLower(gstr.SubStr(field, -2)) == `ip` {
				result, _ := gregex.MatchString(`.*\((\d*)\)`, column[`Type`].String())
				tpl.ApiFilterColumn += fieldCaseCamel + ` string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"ip|length:1,` + result[1] + `"` + "` // " + column[`Comment`].String() + "\n"
				tpl.ApiSaveColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"ip|length:1,` + result[1] + `"` + "` // " + column[`Comment`].String() + "\n"
				tpl.ViewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('common.name.{TplPathSuffixCaseCamelLower}.{TplTableNameCaseSnake}.` + field + `'),
		key: '` + field + `',
		width: 150,
		align: 'center',
	},`
				continue
			}
			//含有mobile和phone字符
			if gstr.Pos(gstr.ToLower(field), `mobile`) != -1 || gstr.Pos(gstr.ToLower(field), `phone`) != -1 {
				result, _ := gregex.MatchString(`.*\((\d*)\)`, column[`Type`].String())
				tpl.ApiFilterColumn += fieldCaseCamel + ` string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"phone|length:1,` + result[1] + `"` + "` // " + column[`Comment`].String() + "\n"
				tpl.ApiSaveColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"phone|length:1,` + result[1] + `"` + "` // " + column[`Comment`].String() + "\n"
				tpl.ViewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('common.name.{TplPathSuffixCaseCamelLower}.{TplTableNameCaseSnake}.` + field + `'),
		key: '` + field + `',
		width: 150,
		align: 'center',
	},`
				continue
			}
			//含有url和link字符
			if gstr.Pos(gstr.ToLower(field), `url`) != -1 || gstr.Pos(gstr.ToLower(field), `link`) != -1 {
				result, _ := gregex.MatchString(`.*\((\d*)\)`, column[`Type`].String())
				tpl.ApiFilterColumn += fieldCaseCamel + ` string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"url|length:1,` + result[1] + `"` + "` // " + column[`Comment`].String() + "\n"
				tpl.ApiSaveColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"url|length:1,` + result[1] + `"` + "` // " + column[`Comment`].String() + "\n"
				tpl.ViewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('common.name.{TplPathSuffixCaseCamelLower}.{TplTableNameCaseSnake}.` + field + `'),
		key: '` + field + `',
		width: 300,
		align: 'center',
	},`
				continue
			}
			//int类型
			if gstr.Pos(column[`Type`].String(), `int`) != -1 {
				if gstr.Pos(column[`Type`].String(), `unsigned`) != -1 {
					tpl.ApiFilterColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"integer"` + "` // " + column[`Comment`].String() + "\n"
					tpl.ApiSaveColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"integer"` + "` // " + column[`Comment`].String() + "\n"
				} else {
					tpl.ApiFilterColumn += fieldCaseCamel + ` *int ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"integer"` + "` // " + column[`Comment`].String() + "\n"
					tpl.ApiSaveColumn += fieldCaseCamel + ` *int ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"integer"` + "` // " + column[`Comment`].String() + "\n"
				}
				tpl.ViewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('common.name.{TplPathSuffixCaseCamelLower}.{TplTableNameCaseSnake}.` + field + `'),
		key: '` + field + `',
		width: 150,
		align: 'center',
	},`
				continue
			}
			//float类型
			if gstr.Pos(column[`Type`].String(), `decimal`) != -1 || gstr.Pos(column[`Type`].String(), `double`) != -1 || gstr.Pos(column[`Type`].String(), `float`) != -1 {
				tpl.ApiFilterColumn += fieldCaseCamel + ` *float64 ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"float"` + "` // " + column[`Comment`].String() + "\n"
				tpl.ApiSaveColumn += fieldCaseCamel + ` *float64 ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"float"` + "` // " + column[`Comment`].String() + "\n"
				tpl.ViewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('common.name.{TplPathSuffixCaseCamelLower}.{TplTableNameCaseSnake}.` + field + `'),
		key: '` + field + `',
		width: 150,
		align: 'center',
	},`
				continue
			}
			//varchar类型
			if gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
				result, _ := gregex.MatchString(`.*\((\d*)\)`, column[`Type`].String())
				tpl.ApiFilterColumn += fieldCaseCamel + ` string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"length:1,` + result[1] + `"` + "` // " + column[`Comment`].String() + "\n"
				tpl.ApiSaveColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"length:1,` + result[1] + `"` + "` // " + column[`Comment`].String() + "\n"
				tpl.ViewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('common.name.{TplPathSuffixCaseCamelLower}.{TplTableNameCaseSnake}.` + field + `'),
		key: '` + field + `',
		width: 150,
		align: 'center',
	},`
				continue
			}
			//char类型
			if gstr.Pos(column[`Type`].String(), `char`) != -1 {
				result, _ := gregex.MatchString(`.*\((\d*)\)`, column[`Type`].String())
				tpl.ApiFilterColumn += fieldCaseCamel + ` string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"length:1,` + result[1] + `"` + "` // " + column[`Comment`].String() + "\n"
				tpl.ApiSaveColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"size:` + result[1] + `"` + "` // " + column[`Comment`].String() + "\n"
				tpl.ViewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('common.name.{TplPathSuffixCaseCamelLower}.{TplTableNameCaseSnake}.` + field + `'),
		key: '` + field + `',
		width: 150,
		align: 'center',
	},`
				continue
			}
			//json类型
			if gstr.Pos(column[`Type`].String(), `json`) != -1 {
				tpl.ApiFilterColumn += fieldCaseCamel + ` string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"json"` + "` // " + column[`Comment`].String() + "\n"
				tpl.ApiSaveColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"json"` + "` // " + column[`Comment`].String() + "\n"
				tpl.ViewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('common.name.{TplPathSuffixCaseCamelLower}.{TplTableNameCaseSnake}.` + field + `'),
		key: '` + field + `',
		width: 200,
		align: 'center',
        hidden: true
	},`
				continue
			}
			//date类型
			if gstr.Pos(column[`Type`].String(), `datetime`) != -1 || gstr.Pos(column[`Type`].String(), `timestamp`) != -1 {
				tpl.ApiFilterColumn += fieldCaseCamel + ` string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"date-format:Y-m-d H:i:s"` + "` // " + column[`Comment`].String() + "\n"
				tpl.ApiSaveColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"date-format:Y-m-d H:i:s"` + "` // " + column[`Comment`].String() + "\n"
				tpl.ViewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('common.name.{TplPathSuffixCaseCamelLower}.{TplTableNameCaseSnake}.` + field + `'),
		key: '` + field + `',
		width: 150,
		align: 'center',
        sortable: true
	},`
				continue
			}
			//date类型
			if gstr.Pos(column[`Type`].String(), `date`) != -1 {
				tpl.ApiFilterColumn += fieldCaseCamel + ` string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"date-format:Y-m-d"` + "` // " + column[`Comment`].String() + "\n"
				tpl.ApiSaveColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"date-format:Y-m-d"` + "` // " + column[`Comment`].String() + "\n"
				tpl.ViewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('common.name.{TplPathSuffixCaseCamelLower}.{TplTableNameCaseSnake}.` + field + `'),
		key: '` + field + `',
		width: 150,
		align: 'center',
        sortable: true
	},`
				continue
			}
			//默认处理
			tpl.ApiFilterColumn += fieldCaseCamel + ` string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:""` + "` // " + column[`Comment`].String() + "\n"
			tpl.ApiSaveColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:""` + "` // " + column[`Comment`].String() + "\n"
			tpl.ViewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('common.name.{TplPathSuffixCaseCamelLower}.{TplTableNameCaseSnake}.` + field + `'),
		key: '` + field + `',
		width: 150,
		align: 'center',
	},`
		}
	}
	tpl.ApiFilterColumn = gstr.SubStr(tpl.ApiFilterColumn, 0, -len("\n"))
	tpl.ApiSaveColumn = gstr.SubStr(tpl.ApiSaveColumn, 0, -len("\n"))
	tpl.ControllerAlloweFieldAppend = gstr.SubStr(tpl.ControllerAlloweFieldAppend, 0, -len(`, `))
	tpl.ControllerAlloweFieldDiff = gstr.SubStr(tpl.ControllerAlloweFieldDiff, 0, -len(`, `))
	return
}

// api模板生成
func MyGenTplApi(ctx context.Context, option *MyGenOption, tpl *MyGenTpl) {
	saveFile := gfile.SelfDir() + `/api/` + tpl.PathSuffixCaseCamelLower + `/` + tpl.TableNameCaseSnake + `.go`
	if gfile.IsFile(saveFile) {
		return
	}
	tplApi := `package api

import (
	apiCommon "api/api"
)

`
	if !option.NoList {
		tplApi += `type {TplTableNameCaseCamel}ListReq struct {
	apiCommon.CommonListReq
	Filter {TplTableNameCaseCamel}ListFilterReq ` + "`" + `p:"filter"` + "`" + `
}

type {TplTableNameCaseCamel}ListFilterReq struct {
	apiCommon.CommonListFilterReq ` + "`" + `c:",omitempty"` + "`" + `
	{TplApiFilterColumn}
}

`
	}
	if !option.NoUpdate {
		tplApi += `type {TplTableNameCaseCamel}InfoReq struct {
	apiCommon.CommonInfoReq
}

`
	}
	if !option.NoCreate {
		tplApi += `type {TplTableNameCaseCamel}CreateReq struct {
	{TplApiSaveColumn}
}

`
	}

	if !option.NoUpdate {
		tplApi += `type {TplTableNameCaseCamel}UpdateReq struct {
	apiCommon.CommonUpdateDeleteIdArrReq ` + "`" + `c:",omitempty"` + "`" + `
	{TplApiSaveColumn}
}

`
	}

	if !option.NoDelete {
		tplApi += `type {TplTableNameCaseCamel}DeleteReq struct {
	apiCommon.CommonUpdateDeleteIdArrReq
}

`
	}

	tplApi = gstr.ReplaceByMap(tplApi, map[string]string{
		`{TplTableNameCaseCamel}`: tpl.TableNameCaseCamel,
		`{TplApiFilterColumn}`:    tpl.ApiFilterColumn,
		`{TplApiSaveColumn}`:      tpl.ApiSaveColumn,
	})

	gfile.PutContents(saveFile, tplApi)
}

// logic模板生成
func MyGenTplLogic(ctx context.Context, option *MyGenOption, tpl *MyGenTpl) {
	saveFile := gfile.SelfDir() + `/internal/logic/` + tpl.PathSuffixCaseCamelLower + `/` + tpl.TableNameCaseSnake + `.go`
	if gfile.IsFile(saveFile) {
		return
	}
	tplLogic := `package logic

import (
	dao{TplPathSuffixCaseCamel} "api/internal/model/dao/{TplPathSuffixCaseCamelLower}"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/text/gregex"
)

type s{TplTableNameCaseCamel} struct{}

func New{TplTableNameCaseCamel}() *s{TplTableNameCaseCamel} {
	return &s{TplTableNameCaseCamel}{}
}

func init() {
	service.Register{TplTableNameCaseCamel}(New{TplTableNameCaseCamel}())
}

`
	if !option.NoList {
		tplLogic += `// 总数
func (logicThis *s{TplTableNameCaseCamel}) Count(ctx context.Context, filter map[string]interface{}) (count int, err error) {
	daoThis := dao{TplPathSuffixCaseCamel}.{TplTableNameCaseCamel}
	joinTableArr := []string{}
	model := daoThis.ParseDbCtx(ctx)
	if len(filter) > 0 {
		model = model.Handler(daoThis.ParseFilter(filter, &joinTableArr))
	}
	if len(joinTableArr) > 0 {
		count, err = model.Handler(daoThis.ParseGroup([]string{` + "`" + `id` + "`" + `}, &joinTableArr)).Distinct().Count(daoThis.PrimaryKey())
	} else {
		count, err = model.Count()
	}
	return
}

// 列表
func (logicThis *s{TplTableNameCaseCamel}) List(ctx context.Context, filter map[string]interface{}, field []string, order [][2]string, page int, limit int) (list gdb.Result, err error) {
	daoThis := dao{TplPathSuffixCaseCamel}.{TplTableNameCaseCamel}
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
		model = model.Handler(daoThis.ParseGroup([]string{` + "`" + `id` + "`" + `}, &joinTableArr))
	}
	if limit > 0 {
		model = model.Offset((page-1)*limit).Limit(limit)
	}
	list, err = model.All()
	return
}

`
	}
	if !option.NoUpdate {
		tplLogic += `// 详情
func (logicThis *s{TplTableNameCaseCamel}) Info(ctx context.Context, filter map[string]interface{}, field ...[]string) (info gdb.Record, err error) {
	daoThis := dao{TplPathSuffixCaseCamel}.{TplTableNameCaseCamel}
	joinTableArr := []string{}
	model := daoThis.ParseDbCtx(ctx)
	model = model.Handler(daoThis.ParseFilter(filter, &joinTableArr))
	if len(field) > 0 && len(field[0]) > 0 {
		model = model.Handler(daoThis.ParseField(field[0], &joinTableArr))
	}
	if len(joinTableArr) > 0 {
		model = model.Handler(daoThis.ParseGroup([]string{` + "`" + `id` + "`" + `}, &joinTableArr))
	}
	info, err = model.One()
	if err != nil {
		return
	}
	if len(info) == 0 {
		err = utils.NewErrorCode(ctx, 29999999, ` + "`" + "`" + `)
		return
	}
	return
}

`
	}
	if !option.NoCreate {
		tplLogic += `// 创建
func (logicThis *s{TplTableNameCaseCamel}) Create(ctx context.Context, data map[string]interface{}) (id int64, err error) {
	daoThis := dao{TplPathSuffixCaseCamel}.{TplTableNameCaseCamel}
	id, err = daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseInsert([]map[string]interface{}{data})).InsertAndGetId()
	if err != nil {
		match, _ := gregex.MatchString(` + "`" + `1062.*Duplicate.*\.([^']*)'` + "`" + `, err.Error())
		if len(match) > 0 {
			err = utils.NewErrorCode(ctx, 29991062, ` + "`" + "`" + `, map[string]interface{}{` + "`" + `errField` + "`" + `: match[1]})
			return
		}
		return
	}
	return
}

`
	}

	if !option.NoUpdate {
		tplLogic += `// 更新
func (logicThis *s{TplTableNameCaseCamel}) Update(ctx context.Context, data map[string]interface{}, filter map[string]interface{}) (row int64, err error) {
	daoThis := dao{TplPathSuffixCaseCamel}.{TplTableNameCaseCamel}
	result, err := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseUpdate(data), daoThis.ParseFilter(filter, &[]string{})).Update()
	if err != nil {
		match, _ := gregex.MatchString(` + "`" + `1062.*Duplicate.*\.([^']*)'` + "`" + `, err.Error())
		if len(match) > 0 {
			err = utils.NewErrorCode(ctx, 29991062, ` + "`" + "`" + `, map[string]interface{}{` + "`" + `errField` + "`" + `: match[1]})
			return
		}
		return
	}
	row, err = result.RowsAffected()
	return
}

`
	}

	if !option.NoDelete {
		tplLogic += `// 删除
func (logicThis *s{TplTableNameCaseCamel}) Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error) {
	daoThis := dao{TplPathSuffixCaseCamel}.{TplTableNameCaseCamel}
	result, err := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Delete()
	if err != nil {
		return
	}
	row, err = result.RowsAffected()
	return
}

`
	}

	tplLogic = gstr.ReplaceByMap(tplLogic, map[string]string{
		`{TplPathSuffixCaseCamel}`:      tpl.PathSuffixCaseCamel,
		`{TplPathSuffixCaseCamelLower}`: tpl.PathSuffixCaseCamelLower,
		`{TplTableNameCaseCamel}`:       tpl.TableNameCaseCamel,
	})

	gfile.PutContents(saveFile, tplLogic)
}

// controller模板生成
func MyGenTplController(ctx context.Context, option *MyGenOption, tpl *MyGenTpl) {
	saveFile := gfile.SelfDir() + `/internal/controller/` + tpl.PathSuffixCaseCamelLower + `/` + tpl.TableNameCaseSnake + `.go`
	if gfile.IsFile(saveFile) {
		return
	}
	tplController := `package controller

import (
	api{TplPathSuffixCaseCamel} "api/api/{TplPathSuffixCaseCamelLower}"
	dao{TplPathSuffixCaseCamel} "api/internal/model/dao/{TplPathSuffixCaseCamelLower}"
	"api/internal/service"
	"api/internal/utils"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

type {TplTableNameCaseCamel} struct{}

func New{TplTableNameCaseCamel}() *{TplTableNameCaseCamel} {
	return &{TplTableNameCaseCamel}{}
}

`
	if !option.NoList {
		tplController += `// 列表
func (controllerThis *{TplTableNameCaseCamel}) List(r *ghttp.Request) {
	/**--------参数处理 开始--------**/
	var param *api{TplPathSuffixCaseCamel}.{TplTableNameCaseCamel}ListReq
	err := r.Parse(&param)
		if err != nil {
			utils.HttpFailJson(r, utils.NewErrorCode(r.GetCtx(), 89999999, err.Error()))
			return
		}
	filter := gconv.Map(param.Filter)
	order := [][2]string{{` + "`" + `id` + "`" + `, ` + "`" + `DESC` + "`" + `}}
	if param.Sort.Key != ` + "`" + "`" + ` {
		order[0][0] = param.Sort.Key
	}
	if param.Sort.Order != ` + "`" + "`" + ` {
		order[0][1] = param.Sort.Order
	}
	if param.Page <= 0 {
		param.Page = 1
	}
	if param.Limit <= 0 {
		param.Limit = 10
	}
	/**--------参数处理 结束--------**/

	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case ` + "`" + `platformAdmin` + "`" + `:
		/**--------权限验证 开始--------**/
		isAuth, _ := service.Action().CheckAuth(r.GetCtx(), ` + "`" + `{TplRawTableNameCaseCamelLower}Look` + "`" + `)
		allowField := []string{{TplControllerAlloweFieldAppend}` + `, ` + "`" + `id` + "`" + `}
		if isAuth {
			allowField = dao{TplPathSuffixCaseCamel}.{TplTableNameCaseCamel}.ColumnArr()
			allowField = append(allowField, ` + "`" + `id` + "`" + `)`
		if tpl.ControllerAlloweFieldDiff != `` {
			tplController += `
			allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{{TplControllerAlloweFieldDiff}})).Slice() //移除敏感字段`
		}
		tplController += `
		}
		field := allowField
		if len(param.Field) > 0 {
			field = gset.NewStrSetFrom(param.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
			if len(field) == 0 {
				field = allowField
			}
		}
		/**--------权限验证 结束--------**/

		count, err := service.{TplTableNameCaseCamel}().Count(r.GetCtx(), filter)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		list, err := service.{TplTableNameCaseCamel}().List(r.GetCtx(), filter, field, order, param.Page, param.Limit)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		utils.HttpSuccessJson(r, map[string]interface{}{` + "`" + `count` + "`" + `: count, ` + "`" + `list` + "`" + `: list}, 0)
	}
}

`
	}
	if !option.NoUpdate {
		tplController += `// 详情
func (controllerThis *{TplTableNameCaseCamel}) Info(r *ghttp.Request) {
	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case ` + "`" + `platformAdmin` + "`" + `:
		/**--------参数处理 开始--------**/
		var param *api{TplPathSuffixCaseCamel}.{TplTableNameCaseCamel}InfoReq
		err := r.Parse(&param)
		if err != nil {
			utils.HttpFailJson(r, utils.NewErrorCode(r.GetCtx(), 89999999, err.Error()))
			return
		}

		allowField := dao{TplPathSuffixCaseCamel}.{TplTableNameCaseCamel}.ColumnArr()
		allowField = append(allowField, ` + "`" + `id` + "`" + `)`
		if tpl.ControllerAlloweFieldDiff != `` {
			tplController += `
			allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{{TplControllerAlloweFieldDiff}})).Slice() //移除敏感字段`
		}
		tplController += `
		field := allowField
		if len(param.Field) > 0 {
			field = gset.NewStrSetFrom(param.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
			if len(field) == 0 {
				field = allowField
			}
		}
		filter := map[string]interface{}{` + "`" + `id` + "`" + `: param.Id}
		/**--------参数处理 结束--------**/

		/**--------权限验证 开始--------**/
		_, err = service.Action().CheckAuth(r.GetCtx(), ` + "`" + `{TplRawTableNameCaseCamelLower}Look` + "`" + `)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		/**--------权限验证 结束--------**/

		info, err := service.{TplTableNameCaseCamel}().Info(r.GetCtx(), filter, field)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		utils.HttpSuccessJson(r, map[string]interface{}{` + "`" + `info` + "`" + `: info}, 0)
	}
}

`
	}
	if !option.NoCreate {
		tplController += `// 创建
func (controllerThis *{TplTableNameCaseCamel}) Create(r *ghttp.Request) {
	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case ` + "`" + `platformAdmin` + "`" + `:
		/**--------参数处理 开始--------**/
		var param *api{TplPathSuffixCaseCamel}.{TplTableNameCaseCamel}CreateReq
		err := r.Parse(&param)
		if err != nil {
			utils.HttpFailJson(r, utils.NewErrorCode(r.GetCtx(), 89999999, err.Error()))
			return
		}
		data := gconv.Map(param)
		/**--------参数处理 结束--------**/

		/**--------权限验证 开始--------**/
		_, err = service.Action().CheckAuth(r.GetCtx(), ` + "`" + `{TplRawTableNameCaseCamelLower}Create` + "`" + `)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		/**--------权限验证 结束--------**/

		_, err = service.{TplTableNameCaseCamel}().Create(r.GetCtx(), data)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		utils.HttpSuccessJson(r, map[string]interface{}{}, 0)
	}
}

`
	}

	if !option.NoUpdate {
		tplController += `// 更新
func (controllerThis *{TplTableNameCaseCamel}) Update(r *ghttp.Request) {
	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case ` + "`" + `platformAdmin` + "`" + `:
		/**--------参数处理 开始--------**/
		var param *api{TplPathSuffixCaseCamel}.{TplTableNameCaseCamel}UpdateReq
		err := r.Parse(&param)
		if err != nil {
			utils.HttpFailJson(r, utils.NewErrorCode(r.GetCtx(), 89999999, err.Error()))
			return
		}
		data := gconv.Map(param)
		delete(data, ` + "`" + `idArr` + "`" + `)
		if len(data) == 0 {
			utils.HttpFailJson(r, utils.NewErrorCode(r.GetCtx(), 89999999, ` + "`" + "`" + `))
			return
		}
		filter := map[string]interface{}{` + "`" + `id` + "`" + `: param.IdArr}
		/**--------参数处理 结束--------**/

		/**--------权限验证 开始--------**/
		_, err = service.Action().CheckAuth(r.GetCtx(), ` + "`" + `{TplRawTableNameCaseCamelLower}Update` + "`" + `)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		/**--------权限验证 结束--------**/

		_, err = service.{TplTableNameCaseCamel}().Update(r.GetCtx(), data, filter)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		utils.HttpSuccessJson(r, map[string]interface{}{}, 0)
	}
}

`
	}

	if !option.NoDelete {
		tplController += `// 删除
func (controllerThis *{TplTableNameCaseCamel}) Delete(r *ghttp.Request) {
	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case ` + "`" + `platformAdmin` + "`" + `:
		/**--------参数处理 开始--------**/
		var param *api{TplPathSuffixCaseCamel}.{TplTableNameCaseCamel}DeleteReq
		err := r.Parse(&param)
		if err != nil {
			utils.HttpFailJson(r, utils.NewErrorCode(r.GetCtx(), 89999999, err.Error()))
			return
		}
		filter := map[string]interface{}{` + "`" + `id` + "`" + `: param.IdArr}
		/**--------参数处理 结束--------**/

		/**--------权限验证 开始--------**/
		_, err = service.Action().CheckAuth(r.GetCtx(), ` + "`" + `{TplRawTableNameCaseCamelLower}Delete` + "`" + `)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		/**--------权限验证 结束--------**/

		_, err = service.{TplTableNameCaseCamel}().Delete(r.GetCtx(), filter)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		utils.HttpSuccessJson(r, map[string]interface{}{}, 0)
	}
}

`
	}

	tplController = gstr.ReplaceByMap(tplController, map[string]string{
		`{TplRawTableNameCaseCamelLower}`:  tpl.RawTableNameCaseCamelLower,
		`{TplPathSuffixCaseCamel}`:         tpl.PathSuffixCaseCamel,
		`{TplPathSuffixCaseCamelLower}`:    tpl.PathSuffixCaseCamelLower,
		`{TplTableNameCaseCamel}`:          tpl.TableNameCaseCamel,
		`{TplControllerAlloweFieldAppend}`: tpl.ControllerAlloweFieldAppend,
		`{TplControllerAlloweFieldDiff}`:   tpl.ControllerAlloweFieldDiff,
	})
	gfile.PutContents(saveFile, tplController)
}

// view模板生成
func MyGenTplView(ctx context.Context, option *MyGenOption, tpl *MyGenTpl) {
	path := gfile.SelfDir()
	saveViewPathPrefix := path + `/../dev/admin/platform/src/views/` + tpl.PathSuffixCaseCamelLower + `/` + tpl.TableNameCaseSnake
	MyGenTplViewIndex(ctx, option, tpl, saveViewPathPrefix)
	MyGenTplViewList(ctx, option, tpl, saveViewPathPrefix)
	MyGenTplViewQuery(ctx, option, tpl, saveViewPathPrefix)
	MyGenTplViewSave(ctx, option, tpl, saveViewPathPrefix)
}

// view模板生成Index
func MyGenTplViewIndex(ctx context.Context, option *MyGenOption, tpl *MyGenTpl, saveViewPathPrefix string) {
	saveFile := saveViewPathPrefix + `/Index.vue`
	if gfile.IsFile(saveFile) {
		return
	}
	tplView := `<script setup lang="ts">
import List from './List.vue'
import Query from './Query.vue'`
	if !(option.NoCreate && option.NoUpdate) {
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
	if !(option.NoCreate && option.NoUpdate) {
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
	if !(option.NoCreate && option.NoUpdate) {
		tplView += `

		<!-- 加上v-if每次都重新生成组件。可防止不同操作之间的影响；新增操作数据的默认值也能写在save组件内 -->
		<Save v-if="saveCommon.visible" />`
	}
	tplView += `
	</ElContainer>
</template>`

	gfile.PutContents(saveViewPathPrefix+`/Index.vue`, tplView)
}

// view模板生成List
func MyGenTplViewList(ctx context.Context, option *MyGenOption, tpl *MyGenTpl, saveViewPathPrefix string) {
	saveFile := saveViewPathPrefix + `/List.vue`
	if gfile.IsFile(saveFile) {
		return
	}
	tplView := `<script setup lang="ts">
const { t } = useI18n()

const table = reactive({
	columns: [{
		dataKey: 'id',
		title: t('common.name.id'),
		key: 'id',
		width: 200,
		align: 'center',
		fixed: 'left',
		sortable: true,`
	if !(option.NoUpdate && option.NoDelete) {
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
	},{TplViewListColumn}
	{
		dataKey: 'updatedAt',
		title: t('common.name.updatedAt'),
		key: 'updatedAt',
		align: 'center',
		width: 150,
		sortable: true,
	},
	{
		dataKey: 'createdAt',
		title: t('common.name.createdAt'),
		key: 'createdAt',
		align: 'center',
		width: 150,
		sortable: true
	},`
	if !(option.NoUpdate && option.NoDelete) {
		tplView += `
	{
		title: t('common.name.action'),
		key: 'action',
		align: 'center',
		fixed: 'right',
		width: 250,
		cellRenderer: (props: any): any => {
			return [
				h(ElButton, {
					type: 'primary',
					size: 'small',
					onClick: () => handleEditCopy(props.rowData.id)
				}, {
					default: () => [h(AutoiconEpEdit), t('common.edit')]
				}),
				h(ElButton, {
					type: 'danger',
					size: 'small',
					onClick: () => handleDelete([props.rowData.id])
				}, {
					default: () => [h(AutoiconEpDelete), t('common.delete')]
				}),
				h(ElButton, {
					type: 'warning',
					size: 'small',
					onClick: () => handleEditCopy(props.rowData.id, 'copy')
				}, {
					default: () => [h(AutoiconEpDocumentCopy), t('common.copy')]
				})
			]
		},
	}] as any,`
	}
	tplView += `
	data: [],
	loading: false,
	sort: { key: 'id', order: 'desc' } as any,
	handleSort: (sort: any) => {
		table.sort.key = sort.key
		table.sort.order = sort.order
		getList()
	},
})`
	if !(option.NoCreate && option.NoUpdate) {
		tplView += `

const saveCommon = inject('saveCommon') as { visible: boolean, title: string, data: { [propName: string]: any } }`
	}
	if !option.NoCreate {
		tplView += `
//新增
const handleAdd = () => {
	saveCommon.data = {}
	saveCommon.title = t('common.add')
	saveCommon.visible = true
}`
	}
	if !option.NoDelete {
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
	if !(option.NoCreate && option.NoUpdate) {
		tplView += `
//编辑|复制
const handleEditCopy = (id: number, type: string = 'edit') => {
	request('/{TplPathSuffixCaseCamelLower}/{TplTableNameCaseCamelLower}/info', { id: id }).then((res) => {
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
	if !option.NoDelete {
		tplView += `
//删除
const handleDelete = (idArr: number[]) => {
	ElMessageBox.confirm('', {
		type: 'warning',
		title: t('common.tip.configDelete'),
		center: true,
		showClose: false,
	}).then(() => {
		request('/{TplPathSuffixCaseCamelLower}/{TplTableNameCaseCamelLower}/del', { idArr: idArr }, true).then((res) => {
			getList()
		}).catch(() => { })
	}).catch(() => { })
}`
	}
	if !option.NoUpdate {
		tplView += `
//更新
const handleUpdate = async (param: { idArr: number[], [propName: string]: any }) => {
	await request('/{TplPathSuffixCaseCamelLower}/{TplTableNameCaseCamelLower}/update', param, true)
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
		sort: table.sort,
		page: pagination.page,
		limit: pagination.size
	}
	table.loading = true
	try {
		const res = await request('/{TplPathSuffixCaseCamelLower}/{TplTableNameCaseCamelLower}/list', param)
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
	if !option.NoCreate {
		tplView += `
				<ElButton type="primary" @click="handleAdd">
					<AutoiconEpEditPen />{{ t('common.add') }}
				</ElButton>`
	}
	if !option.NoDelete {
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
					@column-sort="table.handleSort" :width="width" :height="height" :fixed="true" :row-height="50">
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

	tplView = gstr.ReplaceByMap(tplView, map[string]string{
		`{TplViewListColumn}`: tpl.ViewListColumn, //县替换这个！内部还有变量要替换
	})
	tplView = gstr.ReplaceByMap(tplView, map[string]string{
		`{TplRawTableNameCaseCamelLower}`: tpl.RawTableNameCaseCamelLower,
		`{TplTableNameCaseCamelLower}`:    tpl.TableNameCaseCamelLower,
		`{TplTableNameCaseCamel}`:         tpl.TableNameCaseCamel,
		`{TplTableNameCaseSnake}`:         tpl.TableNameCaseSnake,
		`{TplPathSuffixCaseCamelLower}`:   tpl.PathSuffixCaseCamelLower,
		`{TplPathSuffixCaseCamel}`:        tpl.PathSuffixCaseCamel,
	})
	gfile.PutContents(saveFile, tplView)
}

// view模板生成Query
func MyGenTplViewQuery(ctx context.Context, option *MyGenOption, tpl *MyGenTpl, saveViewPathPrefix string) {
	saveFile := saveViewPathPrefix + `/Query.vue`
	if gfile.IsFile(saveFile) {
		return
	}
	tplView := `<script setup lang="ts">
const { t, tm } = useI18n()

const queryCommon = inject('queryCommon') as { data: { [propName: string]: any } }
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
			<ElInputNumber v-model="queryCommon.data.id" :placeholder="t('common.name.id')" :min="1"
				:controls="false" />
		</ElFormItem>
		<ElFormItem prop="sceneName">
			<ElInput v-model="queryCommon.data.sceneName" :placeholder="t('common.name.auth.scene.sceneName')"
				:clearable="true" />
		</ElFormItem>
		<ElFormItem prop="sceneCode">
			<ElInput v-model="queryCommon.data.sceneCode" :placeholder="t('common.name.auth.scene.sceneCode')"
				:clearable="true" />
		</ElFormItem>
		<ElFormItem prop="isStop" style="width: 100px;">
			<ElSelectV2 v-model="queryCommon.data.isStop" :options="tm('common.status.whether')"
				:placeholder="t('common.name.isStop')" :clearable="true" />
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

	tplView = gstr.ReplaceByMap(tplView, map[string]string{
		`{TplViewListColumn}`: tpl.ViewListColumn, //县替换这个！内部还有变量要替换
	})
	tplView = gstr.ReplaceByMap(tplView, map[string]string{
		`{TplRawTableNameCaseCamelLower}`: tpl.RawTableNameCaseCamelLower,
		`{TplTableNameCaseCamelLower}`:    tpl.TableNameCaseCamelLower,
		`{TplTableNameCaseCamel}`:         tpl.TableNameCaseCamel,
		`{TplTableNameCaseSnake}`:         tpl.TableNameCaseSnake,
		`{TplPathSuffixCaseCamelLower}`:   tpl.PathSuffixCaseCamelLower,
		`{TplPathSuffixCaseCamel}`:        tpl.PathSuffixCaseCamel,
	})
	gfile.PutContents(saveFile, tplView)
}

// view模板生成Save
func MyGenTplViewSave(ctx context.Context, option *MyGenOption, tpl *MyGenTpl, saveViewPathPrefix string) {
	saveFile := saveViewPathPrefix + `/Save.vue`
	if gfile.IsFile(saveFile) {
		return
	}
	tplView := `<script setup lang="ts">
const { t, tm } = useI18n()

const queryCommon = inject('queryCommon') as { data: { [propName: string]: any } }
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
			<ElInputNumber v-model="queryCommon.data.id" :placeholder="t('common.name.id')" :min="1"
				:controls="false" />
		</ElFormItem>
		<ElFormItem prop="sceneName">
			<ElInput v-model="queryCommon.data.sceneName" :placeholder="t('common.name.auth.scene.sceneName')"
				:clearable="true" />
		</ElFormItem>
		<ElFormItem prop="sceneCode">
			<ElInput v-model="queryCommon.data.sceneCode" :placeholder="t('common.name.auth.scene.sceneCode')"
				:clearable="true" />
		</ElFormItem>
		<ElFormItem prop="isStop" style="width: 100px;">
			<ElSelectV2 v-model="queryCommon.data.isStop" :options="tm('common.status.whether')"
				:placeholder="t('common.name.isStop')" :clearable="true" />
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

	tplView = gstr.ReplaceByMap(tplView, map[string]string{
		`{TplViewListColumn}`: tpl.ViewListColumn, //县替换这个！内部还有变量要替换
	})
	tplView = gstr.ReplaceByMap(tplView, map[string]string{
		`{TplRawTableNameCaseCamelLower}`: tpl.RawTableNameCaseCamelLower,
		`{TplTableNameCaseCamelLower}`:    tpl.TableNameCaseCamelLower,
		`{TplTableNameCaseCamel}`:         tpl.TableNameCaseCamel,
		`{TplTableNameCaseSnake}`:         tpl.TableNameCaseSnake,
		`{TplPathSuffixCaseCamelLower}`:   tpl.PathSuffixCaseCamelLower,
		`{TplPathSuffixCaseCamel}`:        tpl.PathSuffixCaseCamel,
	})
	gfile.PutContents(saveFile, tplView)
}
