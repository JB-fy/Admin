package cmd

import (
	"context"
	"fmt"

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
	RawTableNameCaseCamelLower string //原始表名（小驼峰）
	TableNameCaseCamel         string //去除前缀表名（大驼峰）
	TableNameCaseSnake         string //去除前缀表名（蛇形）
	PathSuffixCaseCamelLower   string //路径后缀（小驼峰）
	PathSuffixCaseCamel        string //路径后缀（大驼峰）
	ApiFilterColumn            string //api列表过滤
	ApiSaveColumn              string //api创建更新
}

func MyGenFunc(ctx context.Context, parser *gcmd.Parser) (err error) {
	option := MyGenOptionHandle(ctx, parser)
	tpl := MyGenTplHandle(ctx, option)

	MyGenTplApi(ctx, option, tpl)
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

// 模板变量获取
func MyGenTplHandle(ctx context.Context, option *MyGenOption) (tpl *MyGenTpl) {
	tpl = &MyGenTpl{
		RawTableNameCaseCamelLower: gstr.CaseCamelLower(option.DbTable),
		TableNameCaseCamel:         gstr.CaseCamel(gstr.Replace(option.DbTable, option.RemovePrefix, ``, 1)),
		PathSuffixCaseCamelLower:   gstr.CaseCamelLower(option.RemovePrefix),
		PathSuffixCaseCamel:        gstr.CaseCamel(option.RemovePrefix),
	}
	tpl.TableNameCaseSnake = gstr.CaseSnakeFirstUpper(tpl.TableNameCaseCamel)

	ApiFilterColumn := [][]string{}
	tableColumnList, _ := g.DB(option.DbGroup).GetAll(ctx, `SHOW FULL COLUMNS FROM `+option.DbTable)
	for _, column := range tableColumnList {
		field := column[`Field`].String()
		fieldCaseCamel := gstr.CaseCamel(field)
		switch fieldCaseCamel {
		case `UpdatedAt`, `CreatedAt`, `DeletedAt`: //不处理的字段
		default:
			//主键
			if column[`Key`].String() == `PRI` && column[`Extra`].String() == `auto_increment` {
				ApiFilterColumn = append(ApiFilterColumn, []string{
					"    #" + fieldCaseCamel,
					" # " + "*uint",
					" #" + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"integer|min:1"` + "`",
					" #" + fmt.Sprintf(`// %s`, column[`Comment`].String()),
				})
				tpl.ApiFilterColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"integer|min:1"` + "` //" + column[`Comment`].String() + "\n"
				continue
			}
			//id后缀
			if gstr.ToLower(gstr.SubStr(field, -2)) == `id` {
				tpl.ApiFilterColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"integer|min:1"` + "` //" + column[`Comment`].String() + "\n"
				tpl.ApiSaveColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"integer|min:1"` + "` //" + column[`Comment`].String() + "\n"
				continue
			}
			//status后缀
			if gstr.ToLower(gstr.SubStr(field, -6)) == `status` {
				tpl.ApiFilterColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"integer|in:0,1,2"` + "` //" + column[`Comment`].String() + "\n"
				tpl.ApiSaveColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"integer|in:0,1,2"` + "` //" + column[`Comment`].String() + "\n"
				continue
			}
			//is前缀
			if gstr.ToLower(gstr.SubStr(field, 0, 2)) == `is` {
				tpl.ApiFilterColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"integer|in:0,1"` + "` //" + column[`Comment`].String() + "\n"
				tpl.ApiSaveColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"integer|in:0,1"` + "` //" + column[`Comment`].String() + "\n"
				continue
			}
			//name和code后缀
			if gstr.ToLower(gstr.SubStr(field, -4)) == `name` || gstr.ToLower(gstr.SubStr(field, -4)) == `code` {
				result, _ := gregex.MatchString(`.*\((\d*)\)`, column[`Type`].String())
				tpl.ApiFilterColumn += fieldCaseCamel + ` string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"length:1,` + result[1] + `|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"` + "` //" + column[`Comment`].String() + "\n"
				tpl.ApiSaveColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"length:1,` + result[1] + `|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"` + "` //" + column[`Comment`].String() + "\n"
				continue
			}
			//ip后缀
			if gstr.ToLower(gstr.SubStr(field, -2)) == `ip` {
				result, _ := gregex.MatchString(`.*\((\d*)\)`, column[`Type`].String())
				tpl.ApiFilterColumn += fieldCaseCamel + ` string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"ip|length:1,` + result[1] + `"` + "` //" + column[`Comment`].String() + "\n"
				tpl.ApiSaveColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"ip|length:1,` + result[1] + `"` + "` //" + column[`Comment`].String() + "\n"
				continue
			}
			//含有mobile和phone字符
			if gstr.Pos(gstr.ToLower(field), `mobile`) != -1 || gstr.Pos(gstr.ToLower(field), `phone`) != -1 {
				result, _ := gregex.MatchString(`.*\((\d*)\)`, column[`Type`].String())
				tpl.ApiFilterColumn += fieldCaseCamel + ` string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"phone|length:1,` + result[1] + `"` + "` //" + column[`Comment`].String() + "\n"
				tpl.ApiSaveColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"phone|length:1,` + result[1] + `"` + "` //" + column[`Comment`].String() + "\n"
				continue
			}
			//含有url和link字符
			if gstr.Pos(gstr.ToLower(field), `url`) != -1 || gstr.Pos(gstr.ToLower(field), `link`) != -1 {
				result, _ := gregex.MatchString(`.*\((\d*)\)`, column[`Type`].String())
				tpl.ApiFilterColumn += fieldCaseCamel + ` string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"url|length:1,` + result[1] + `"` + "` //" + column[`Comment`].String() + "\n"
				tpl.ApiSaveColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"url|length:1,` + result[1] + `"` + "` //" + column[`Comment`].String() + "\n"
				continue
			}
			//int类型
			if gstr.Pos(column[`Type`].String(), `int`) != -1 {
				if gstr.Pos(column[`Type`].String(), `unsigned`) != -1 {
					tpl.ApiFilterColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"integer"` + "` //" + column[`Comment`].String() + "\n"
					tpl.ApiSaveColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"integer"` + "` //" + column[`Comment`].String() + "\n"
				} else {
					tpl.ApiFilterColumn += fieldCaseCamel + ` *int ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"integer"` + "` //" + column[`Comment`].String() + "\n"
					tpl.ApiSaveColumn += fieldCaseCamel + ` *int ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"integer"` + "` //" + column[`Comment`].String() + "\n"
				}
				continue
			}
			//float类型
			if gstr.Pos(column[`Type`].String(), `decimal`) != -1 || gstr.Pos(column[`Type`].String(), `double`) != -1 || gstr.Pos(column[`Type`].String(), `float`) != -1 {
				tpl.ApiFilterColumn += fieldCaseCamel + ` *float64 ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"float"` + "` //" + column[`Comment`].String() + "\n"
				tpl.ApiSaveColumn += fieldCaseCamel + ` *float64 ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"float"` + "` //" + column[`Comment`].String() + "\n"
				continue
			}
			//varchar类型
			if gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
				result, _ := gregex.MatchString(`.*\((\d*)\)`, column[`Type`].String())
				tpl.ApiFilterColumn += fieldCaseCamel + ` string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"length:1,` + result[1] + `"` + "` //" + column[`Comment`].String() + "\n"
				tpl.ApiSaveColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"length:1,` + result[1] + `"` + "` //" + column[`Comment`].String() + "\n"
				continue
			}
			//char类型
			if gstr.Pos(column[`Type`].String(), `char`) != -1 {
				result, _ := gregex.MatchString(`.*\((\d*)\)`, column[`Type`].String())
				tpl.ApiFilterColumn += fieldCaseCamel + ` string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"length:1,` + result[1] + `"` + "` //" + column[`Comment`].String() + "\n"
				tpl.ApiSaveColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"size:` + result[1] + `"` + "` //" + column[`Comment`].String() + "\n"
				continue
			}
			//json类型
			if gstr.Pos(column[`Type`].String(), `json`) != -1 {
				tpl.ApiFilterColumn += fieldCaseCamel + ` string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"json"` + "` //" + column[`Comment`].String() + "\n"
				tpl.ApiSaveColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"json"` + "` //" + column[`Comment`].String() + "\n"
				continue
			}
			//默认处理
			tpl.ApiFilterColumn += fieldCaseCamel + ` string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:""` + "` //" + column[`Comment`].String() + "\n"
			tpl.ApiSaveColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:""` + "` //" + column[`Comment`].String() + "\n"
		}
	}
	tpl.ApiFilterColumn = gstr.SubStr(tpl.ApiFilterColumn, 0, -len("\n"))
	tpl.ApiSaveColumn = gstr.SubStr(tpl.ApiSaveColumn, 0, -len("\n"))
	return
}

// api模板生成
func MyGenTplApi(ctx context.Context, option *MyGenOption, tpl *MyGenTpl) {
	path := gfile.SelfDir()
	//tplApiPath := path + `/resource/template/gen_template_api.txt`
	//tplApi := gfile.GetContents(tplApiPath)
	fmt.Println(option)
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

	tplApi = gstr.Replace(tplApi, `{TplTableNameCaseCamel}`, tpl.TableNameCaseCamel)
	tplApi = gstr.Replace(tplApi, `{TplApiFilterColumn}`, tpl.ApiFilterColumn)
	tplApi = gstr.Replace(tplApi, `{TplApiSaveColumn}`, tpl.ApiSaveColumn)

	saveApiPath := path + `/api/` + tpl.PathSuffixCaseCamelLower + `/` + tpl.TableNameCaseSnake + `.go`
	//saveApiPath := path + `/internal/logic/` + tpl.PathSuffixCaseCamelLower + `/` + tpl.TableNameCaseSnake + `.go`
	gfile.PutContents(saveApiPath, tplApi)
}
