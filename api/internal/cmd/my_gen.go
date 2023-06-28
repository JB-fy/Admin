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
	DbGroup       string `c:"dbGroup"`              //db分组
	DbTable       string `c:"dbTable"`              //db表
	RemovePrefix  string `c:"removePrefix"`         //要删除的db表前缀
	DirPathSuffix string `c:"pathSuffix"`           //路径后缀。即为模块文件夹名称
	NoList        bool   `c:"noList,default:true" ` //不生成列表接口(0,false,off,no,""为false，其他都为true)
	NoCreate      bool   `c:"noCreate"`             //不生成创建接口(0,false,off,no,""为false，其他都为true)
	NoUpdate      bool   `c:"noUpdate"`             //不生成更新接口(0,false,off,no,""为false，其他都为true)
	NoDelete      bool   `c:"noDelete"`             //不生成删除接口(0,false,off,no,""为false，其他都为true)
	IsCover       bool   `c:"isCover"`              //如果生成的文件已存在，是否覆盖
}

type MyGenTpl struct {
	TableColumnList             gdb.Result //表字段详情
	RawTableNameCaseCamelLower  string     //原始表名（小驼峰）
	TableNameCaseCamelLower     string     //去除前缀表名（小驼峰）
	TableNameCaseCamel          string     //去除前缀表名（大驼峰）
	TableNameCaseSnake          string     //去除前缀表名（蛇形）
	PathSuffixCaseCamelLower    string     //路径后缀（小驼峰）
	PathSuffixCaseCamel         string     //路径后缀（大驼峰）
	ApiFilterColumn             string     //api列表字段
	ApiCreateColumn             string     //api创建字段
	ApiUpdateColumn             string     //api更新字段
	ControllerAlloweFieldAppend string     //controller追加字段
	ControllerAlloweFieldDiff   string     //controller移除字段
	ViewListColumn              string     //view列表字段
	ViewQueryField              string     //view查询字段
	ViewSaveRule                string     //view创建更新字段验证规则
	ViewSaveField               string     //view创建更新字段
	ViewI18nField               string     //view多语言字段
}

func MyGenFunc(ctx context.Context, parser *gcmd.Parser) (err error) {
	option := MyGenOptionHandle(ctx, parser)

	tableColumnList, _ := g.DB(option.DbGroup).GetAll(ctx, `SHOW FULL COLUMNS FROM `+option.DbTable)
	tpl := &MyGenTpl{
		TableColumnList:            tableColumnList,
		RawTableNameCaseCamelLower: gstr.CaseCamelLower(option.DbTable),
		TableNameCaseCamel:         gstr.CaseCamel(gstr.Replace(option.DbTable, option.RemovePrefix, ``, 1)),
		PathSuffixCaseCamelLower:   gstr.CaseCamelLower(option.RemovePrefix),
		PathSuffixCaseCamel:        gstr.CaseCamel(option.RemovePrefix),
	}
	tpl.TableNameCaseCamelLower = gstr.CaseCamelLower(tpl.TableNameCaseCamel)
	tpl.TableNameCaseSnake = gstr.CaseSnakeFirstUpper(tpl.TableNameCaseCamel)

	MyGenTplApi(ctx, option, tpl)
	MyGenTplLogic(ctx, option, tpl)
	MyGenTplController(ctx, option, tpl)
	MyGenTplRouter(ctx, option, tpl)

	MyGenTplViewIndex(ctx, option, tpl)
	MyGenTplViewList(ctx, option, tpl)
	MyGenTplViewQuery(ctx, option, tpl)
	MyGenTplViewSave(ctx, option, tpl)
	MyGenTplViewI18n(ctx, option, tpl)
	MyGenTplViewRouter(ctx, option, tpl)
	return
}

// 参数处理
func MyGenOptionHandle(ctx context.Context, parser *gcmd.Parser) (option *MyGenOption) {
	optionMap := parser.GetOptAll()
	option = &MyGenOption{}
	gconv.Struct(optionMap, option)

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
	_, ok := optionMap[`removePrefix`]
	if !ok {
		option.RemovePrefix = gcmd.Scan("> 请输入要删除的db表前缀,默认(空):\n")
	}
	for {
		if option.RemovePrefix == `` || gstr.Pos(option.DbTable, option.RemovePrefix) == 0 {
			break
		}
		option.RemovePrefix = gcmd.Scan("> 要删除的db表前缀不存在，请重新输入，默认(空):\n")
	}
noAllRestart:
	noList, ok := optionMap[`noList`]
	if !ok {
		noList = gcmd.Scan("> 是否生成列表接口，默认(yes):\n")
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
			noList = gcmd.Scan("> 输入错误，请重新输入，是否生成列表接口，默认(yes):\n")
		}
	}
	noCreate, ok := optionMap[`noCreate`]
	if !ok {
		noCreate = gcmd.Scan("> 是否生成创建接口，默认(yes):\n")
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
			noCreate = gcmd.Scan("> 输入错误，请重新输入，是否生成创建接口，默认(yes):\n")
		}
	}
	noUpdate, ok := optionMap[`noUpdate`]
	if !ok {
		noUpdate = gcmd.Scan("> 是否生成更新接口，默认(yes):\n")
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
			noUpdate = gcmd.Scan("> 输入错误，请重新输入，是否生成更新接口，默认(yes):\n")
		}
	}
	noDelete, ok := optionMap[`noDelete`]
	if !ok {
		noDelete = gcmd.Scan("> 是否生成删除接口，默认(yes):\n")
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
			noDelete = gcmd.Scan("> 输入错误，请重新输入，是否生成删除接口，默认(yes):\n")
		}
	}
	if option.NoList && option.NoCreate && option.NoUpdate && option.NoDelete {
		fmt.Println("请重新选择生成哪些接口，不能全是no！")
		goto noAllRestart
	}
	isCover, ok := optionMap[`isCover`]
	if !ok {
		isCover = gcmd.Scan("> 如果文件已存在，是否覆盖原文件，默认(no):\n")
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
			isCover = gcmd.Scan("> 输入错误，请重新输入，是否覆盖原文件，默认(no):\n")
		}
	}
	return
}

// api模板生成
func MyGenTplApi(ctx context.Context, option *MyGenOption, tpl *MyGenTpl) {
	saveFile := gfile.SelfDir() + `/api/` + tpl.PathSuffixCaseCamelLower + `/` + tpl.TableNameCaseSnake + `.go`
	if !option.IsCover && gfile.IsFile(saveFile) {
		return
	}

	for _, column := range tpl.TableColumnList {
		field := column[`Field`].String()
		fieldCaseCamel := gstr.CaseCamel(field)
		fieldCaseCamelLower := gstr.CaseCamelLower(field)
		fieldCaseSnake := gstr.CaseSnakeFirstUpper(field)
		comment := gstr.Trim(gstr.ReplaceByArray(column[`Comment`].String(), g.SliceStr{
			"\n", " ",
			"\r", " ",
		}))
		result, _ := gregex.MatchString(`.*\((\d*)\)`, column[`Type`].String())
		switch fieldCaseCamel {
		case `CreatedAt`, `UpdatedAt`, `DeletedAt`: //不处理的字段
		default:
			//主键
			if column[`Key`].String() == `PRI` && column[`Extra`].String() == `auto_increment` && field != `id` {
				tpl.ApiFilterColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer|min:1"` + "` // " + comment + "\n"
				continue
			}
			//pid字段
			if field == `pid` {
				tpl.ApiFilterColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer|min:1"` + "` // " + comment + "\n"
				tpl.ApiCreateColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer|min:1"` + "` // " + comment + "\n"
				tpl.ApiUpdateColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer|min:1"` + "` // " + comment + "\n"
				continue
			}
			//is_stop或isStop字段
			if fieldCaseCamelLower == `isStop` {
				tpl.ApiFilterColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer|in:0,1"` + "` // " + comment + "\n"
				tpl.ApiCreateColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer|in:0,1"` + "` // " + comment + "\n"
				tpl.ApiUpdateColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer|in:0,1"` + "` // " + comment + "\n"
				continue
			}
			//gender字段
			if field == `gender` {
				tpl.ApiFilterColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer|in:0,1,2"` + "` // " + comment + "\n"
				tpl.ApiCreateColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer|in:0,1,2"` + "` // " + comment + "\n"
				tpl.ApiUpdateColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer|in:0,1,2"` + "` // " + comment + "\n"
				continue
			}
			//avator字段
			if field == `avator` {
				tpl.ApiCreateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"url|length:1,` + result[1] + `"` + "` // " + comment + "\n"
				tpl.ApiUpdateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"url|length:1,` + result[1] + `"` + "` // " + comment + "\n"
				continue
			}
			//id后缀
			if gstr.SubStr(fieldCaseCamel, -2) == `Id` {
				tpl.ApiFilterColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer|min:1"` + "` // " + comment + "\n"
				tpl.ApiCreateColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer|min:1"` + "` // " + comment + "\n"
				tpl.ApiUpdateColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer|min:1"` + "` // " + comment + "\n"
				continue
			}
			//name或code后缀
			if gstr.SubStr(fieldCaseCamel, -4) == `Name` || gstr.SubStr(fieldCaseCamel, -4) == `Code` {
				tpl.ApiFilterColumn += fieldCaseCamel + ` string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"length:1,` + result[1] + `|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"` + "` // " + comment + "\n"
				tpl.ApiCreateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"required|length:1,` + result[1] + `|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"` + "` // " + comment + "\n"
				tpl.ApiUpdateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"length:1,` + result[1] + `|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"` + "` // " + comment + "\n"
				continue
			}
			//mobile或phone后缀
			if gstr.SubStr(fieldCaseCamel, -5) == `Phone` || gstr.SubStr(fieldCaseCamel, -6) == `Mobile` {
				tpl.ApiFilterColumn += fieldCaseCamel + ` string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"phone|length:1,` + result[1] + `"` + "` // " + comment + "\n"
				tpl.ApiCreateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"phone|length:1,` + result[1] + `"` + "` // " + comment + "\n"
				tpl.ApiUpdateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"phone|length:1,` + result[1] + `"` + "` // " + comment + "\n"
				continue
			}
			//url或link后缀
			if gstr.SubStr(fieldCaseCamel, -3) == `Url` || gstr.SubStr(fieldCaseCamel, -4) == `Link` {
				tpl.ApiFilterColumn += fieldCaseCamel + ` string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"url|length:1,` + result[1] + `"` + "` // " + comment + "\n"
				tpl.ApiCreateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"url|length:1,` + result[1] + `"` + "` // " + comment + "\n"
				tpl.ApiUpdateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"url|length:1,` + result[1] + `"` + "` // " + comment + "\n"
				continue
			}
			//img或image或cover后缀
			if gstr.SubStr(fieldCaseCamel, -3) == `Img` || gstr.SubStr(fieldCaseCamel, -5) == `Image` || gstr.SubStr(fieldCaseCamel, -5) == `Cover` {
				if column[`Type`].String() == `json` {
					tpl.ApiCreateColumn += fieldCaseCamel + ` *[]string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"distinct|foreach|url|foreach|min-length:1"` + "` // " + comment + "\n"
					tpl.ApiUpdateColumn += fieldCaseCamel + ` *[]string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"distinct|foreach|url|foreach|min-length:1"` + "` // " + comment + "\n"
				} else {
					tpl.ApiCreateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"url|length:1,` + result[1] + `"` + "` // " + comment + "\n"
					tpl.ApiUpdateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"url|length:1,` + result[1] + `"` + "` // " + comment + "\n"
				}
				continue
			}
			//video后缀
			if gstr.SubStr(fieldCaseCamel, -5) == `Video` {
				if column[`Type`].String() == `json` {
					tpl.ApiCreateColumn += fieldCaseCamel + ` *[]string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"distinct|foreach|url|foreach|min-length:1"` + "` // " + comment + "\n"
					tpl.ApiUpdateColumn += fieldCaseCamel + ` *[]string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"distinct|foreach|url|foreach|min-length:1"` + "` // " + comment + "\n"
				} else {
					tpl.ApiCreateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"url|length:1,` + result[1] + `"` + "` // " + comment + "\n"
					tpl.ApiUpdateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"url|length:1,` + result[1] + `"` + "` // " + comment + "\n"
				}
				continue
			}
			//Ip后缀
			if gstr.SubStr(fieldCaseCamel, -2) == `Ip` {
				tpl.ApiFilterColumn += fieldCaseCamel + ` string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"ip|length:1,` + result[1] + `"` + "` // " + comment + "\n"
				tpl.ApiCreateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"ip|length:1,` + result[1] + `"` + "` // " + comment + "\n"
				tpl.ApiUpdateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"ip|length:1,` + result[1] + `"` + "` // " + comment + "\n"
				continue
			}
			//sort或weight后缀
			if gstr.SubStr(fieldCaseCamel, -4) == `Sort` || gstr.SubStr(fieldCaseCamel, -6) == `Weight` {
				tpl.ApiFilterColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer|between:0,100"` + "` // " + comment + "\n"
				tpl.ApiCreateColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer|between:0,100"` + "` // " + comment + "\n"
				tpl.ApiUpdateColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer|between:0,100"` + "` // " + comment + "\n"
				continue
			}
			//status后缀
			if gstr.SubStr(fieldCaseCamel, -6) == `Status` {
				tpl.ApiFilterColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer|in:0,1,2"` + "` // " + comment + "\n"
				tpl.ApiCreateColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer|in:0,1,2"` + "` // " + comment + "\n"
				tpl.ApiUpdateColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer|in:0,1,2"` + "` // " + comment + "\n"
				continue
			}
			//is_前缀
			if gstr.SubStr(fieldCaseSnake, 0, 3) == `is_` {
				tpl.ApiFilterColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer|in:0,1"` + "` // " + comment + "\n"
				tpl.ApiCreateColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer|in:0,1"` + "` // " + comment + "\n"
				tpl.ApiUpdateColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer|in:0,1"` + "` // " + comment + "\n"
				continue
			}
			//int类型
			if gstr.Pos(column[`Type`].String(), `int`) != -1 {
				if gstr.Pos(column[`Type`].String(), `unsigned`) != -1 {
					tpl.ApiFilterColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer"` + "` // " + comment + "\n"
					tpl.ApiCreateColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer"` + "` // " + comment + "\n"
					tpl.ApiUpdateColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer"` + "` // " + comment + "\n"
				} else {
					tpl.ApiFilterColumn += fieldCaseCamel + ` *int ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer"` + "` // " + comment + "\n"
					tpl.ApiCreateColumn += fieldCaseCamel + ` *int ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer"` + "` // " + comment + "\n"
					tpl.ApiUpdateColumn += fieldCaseCamel + ` *int ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"integer"` + "` // " + comment + "\n"
				}
				continue
			}
			//float类型
			if gstr.Pos(column[`Type`].String(), `decimal`) != -1 || gstr.Pos(column[`Type`].String(), `double`) != -1 || gstr.Pos(column[`Type`].String(), `float`) != -1 {
				tpl.ApiFilterColumn += fieldCaseCamel + ` *float64 ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"float"` + "` // " + comment + "\n"
				tpl.ApiCreateColumn += fieldCaseCamel + ` *float64 ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"float"` + "` // " + comment + "\n"
				tpl.ApiUpdateColumn += fieldCaseCamel + ` *float64 ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"float"` + "` // " + comment + "\n"
				continue
			}
			//varchar类型
			if gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
				tpl.ApiFilterColumn += fieldCaseCamel + ` string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"length:1,` + result[1] + `"` + "` // " + comment + "\n"
				tpl.ApiCreateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"length:1,` + result[1] + `"` + "` // " + comment + "\n"
				tpl.ApiUpdateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"length:1,` + result[1] + `"` + "` // " + comment + "\n"
				continue
			}
			//char类型
			if gstr.Pos(column[`Type`].String(), `char`) != -1 {
				tpl.ApiFilterColumn += fieldCaseCamel + ` string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"length:1,` + result[1] + `"` + "` // " + comment + "\n"
				tpl.ApiCreateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"size:` + result[1] + `"` + "` // " + comment + "\n"
				tpl.ApiUpdateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"size:` + result[1] + `"` + "` // " + comment + "\n"
				continue
			}
			//datetime和timestamp类型
			if gstr.Pos(column[`Type`].String(), `datetime`) != -1 || gstr.Pos(column[`Type`].String(), `timestamp`) != -1 {
				tpl.ApiFilterColumn += fieldCaseCamel + ` string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"date-format:Y-m-d H:i:s"` + "` // " + comment + "\n"
				if column[`Null`].String() == `NO` && column[`Default`].String() == `` {
					tpl.ApiCreateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"required|date-format:Y-m-d H:i:s"` + "` // " + comment + "\n"
				} else {
					tpl.ApiCreateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"date-format:Y-m-d H:i:s"` + "` // " + comment + "\n"
				}
				tpl.ApiUpdateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"date-format:Y-m-d H:i:s"` + "` // " + comment + "\n"
				continue
			}
			//date类型
			if gstr.Pos(column[`Type`].String(), `date`) != -1 {
				tpl.ApiFilterColumn += fieldCaseCamel + ` string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"date-format:Y-m-d"` + "` // " + comment + "\n"
				if column[`Null`].String() == `NO` && column[`Default`].String() == `` {
					tpl.ApiCreateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"required|date-format:Y-m-d"` + "` // " + comment + "\n"
				} else {
					tpl.ApiCreateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"date-format:Y-m-d"` + "` // " + comment + "\n"
				}
				tpl.ApiUpdateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"date-format:Y-m-d"` + "` // " + comment + "\n"
				continue
			}
			//json类型
			if gstr.Pos(column[`Type`].String(), `json`) != -1 {
				tpl.ApiFilterColumn += fieldCaseCamel + ` string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"json"` + "` // " + comment + "\n"
				if column[`Null`].String() == `NO` {
					tpl.ApiCreateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"required|json"` + "` // " + comment + "\n"
				} else {
					tpl.ApiCreateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"json"` + "` // " + comment + "\n"
				}
				tpl.ApiUpdateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"json"` + "` // " + comment + "\n"
				continue
			}
			//text类型
			if gstr.Pos(column[`Type`].String(), `text`) != -1 {
				tpl.ApiFilterColumn += fieldCaseCamel + ` string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:""` + "` // " + comment + "\n"
				if column[`Null`].String() == `NO` {
					tpl.ApiCreateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:"required"` + "` // " + comment + "\n"
				} else {
					tpl.ApiCreateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:""` + "` // " + comment + "\n"
				}
				tpl.ApiUpdateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:""` + "` // " + comment + "\n"
				continue
			}
			//默认处理
			tpl.ApiFilterColumn += fieldCaseCamel + ` string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:""` + "` // " + comment + "\n"
			tpl.ApiCreateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:""` + "` // " + comment + "\n"
			tpl.ApiUpdateColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" json:"` + field + `" v:""` + "` // " + comment + "\n"
		}
	}
	tpl.ApiFilterColumn = gstr.SubStr(tpl.ApiFilterColumn, 0, -len("\n"))
	tpl.ApiCreateColumn = gstr.SubStr(tpl.ApiCreateColumn, 0, -len("\n"))
	tpl.ApiUpdateColumn = gstr.SubStr(tpl.ApiUpdateColumn, 0, -len("\n"))

	tplApi := `package api

import (
	apiCommon "api/api"
)

`
	if !option.NoList {
		tplApi += `type {TplTableNameCaseCamel}ListReq struct {
	Filter {TplTableNameCaseCamel}ListFilterReq ` + "`" + `json:"filter"` + "`" + `
	Field []string ` + "`" + `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段。默认会返回全部查询字段。如果需要的字段较少，建议指定字段，传值参考默认返回的字段"` + "`" + `
	Sort  string   ` + "`" + `json:"sort" default:"id DESC" dc:"排序"` + "`" + `
	Page  int      ` + "`" + `json:"page" v:"integer|min:1" default:"1" dc:"页码"` + "`" + `
	Limit int      ` + "`" + `json:"limit" v:"integer|min:0" default:"10" dc:"每页数量。可传0取全部"` + "`" + `
}

type {TplTableNameCaseCamel}ListFilterReq struct {
	Id        *uint       ` + "`" + `c:"id,omitempty" json:"id" v:"integer|min:1" dc:"ID"` + "`" + `
	IdArr     []uint      ` + "`" + `c:"idArr,omitempty" json:"idArr" v:"distinct|foreach|integer|foreach|min:1" dc:"ID数组"` + "`" + `
	ExcId     *uint       ` + "`" + `c:"excId,omitempty" json:"excId" v:"integer|min:1" dc:"排除ID"` + "`" + `
	ExcIdArr  []uint      ` + "`" + `c:"excIdArr,omitempty" json:"excIdArr" v:"distinct|foreach|integer|foreach|min:1" dc:"排除ID数组"` + "`" + `
	StartTime *gtime.Time ` + "`" + `c:"startTime,omitempty" json:"startTime" v:"date-format:Y-m-d H:i:s" dc:"开始时间。示例：2000-01-01 00:00:00"` + "`" + `
	EndTime   *gtime.Time ` + "`" + `c:"endTime,omitempty" json:"endTime" v:"date-format:Y-m-d H:i:s|after-equal:StartTime" dc:"结束时间。示例：2000-01-01 00:00:00"` + "`" + `
	Name      string      ` + "`" + `c:"name,omitempty" json:"name" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"名称。后台公共列表常用"` + "`" + `
	{TplApiFilterColumn}
}

`
	}
	if !option.NoUpdate {
		tplApi += `type {TplTableNameCaseCamel}InfoReq struct {
	Id    uint     ` + "`" + `json:"id" v:"required|integer|min:1" dc:"ID"` + "`" + `
	Field []string ` + "`" + `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段。默认会返回全部查询字段。如果需要的字段较少，建议指定字段，传值参考默认返回的字段"` + "`" + `
}

`
	}
	if !option.NoCreate {
		tplApi += `type {TplTableNameCaseCamel}CreateReq struct {
	{TplApiCreateColumn}
}

`
	}

	if !option.NoUpdate {
		tplApi += `type {TplTableNameCaseCamel}UpdateReq struct {
	IdArr []uint ` + "`" + `c:"idArr,omitempty" json:"idArr" v:"required|distinct|foreach|integer|foreach|min:1" dc:"ID数组"` + "`" + `
	{TplApiUpdateColumn}
}

`
	}

	if !option.NoDelete {
		tplApi += `type {TplTableNameCaseCamel}DeleteReq struct {
	IdArr []uint ` + "`" + `c:"idArr,omitempty" json:"idArr" v:"required|distinct|foreach|integer|foreach|min:1" dc:"ID数组"` + "`" + `
}
`
	}

	tplApi = gstr.ReplaceByMap(tplApi, map[string]string{
		`{TplTableNameCaseCamel}`: tpl.TableNameCaseCamel,
		`{TplApiFilterColumn}`:    tpl.ApiFilterColumn,
		`{TplApiCreateColumn}`:    tpl.ApiCreateColumn,
		`{TplApiUpdateColumn}`:    tpl.ApiUpdateColumn,
	})

	gfile.PutContents(saveFile, tplApi)
}

// logic模板生成
func MyGenTplLogic(ctx context.Context, option *MyGenOption, tpl *MyGenTpl) {
	saveFile := gfile.SelfDir() + `/internal/logic/` + tpl.PathSuffixCaseCamelLower + `/` + tpl.TableNameCaseSnake + `.go`
	if !option.IsCover && gfile.IsFile(saveFile) {
		return
	}

	tplLogic := `package logic

import (
	dao{TplPathSuffixCaseCamel} "api/internal/dao/{TplPathSuffixCaseCamelLower}"
	"api/internal/service"`
	if !(option.NoCreate && option.NoUpdate && option.NoDelete) {
		tplLogic += `
	"api/internal/utils"`
	}
	tplLogic += `
	"context"

	"github.com/gogf/gf/v2/database/gdb"`
	if !(option.NoCreate && option.NoUpdate) {
		tplLogic += `
	"github.com/gogf/gf/v2/text/gregex"`
	}
	tplLogic += `
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
		count, err = model.Handler(daoThis.ParseGroup([]string{` + "`id`" + `}, &joinTableArr)).Distinct().Count(daoThis.PrimaryKey())
	} else {
		count, err = model.Count()
	}
	return
}

// 列表
func (logicThis *s{TplTableNameCaseCamel}) List(ctx context.Context, filter map[string]interface{}, field []string, order []string, page int, limit int) (list gdb.Result, err error) {
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
		model = model.Handler(daoThis.ParseGroup([]string{` + "`id`" + `}, &joinTableArr))
	}
	if limit > 0 {
		model = model.Offset((page - 1) * limit).Limit(limit)
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
		model = model.Handler(daoThis.ParseGroup([]string{` + "`id`" + `}, &joinTableArr))
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
		tplLogic += `// 新增
func (logicThis *s{TplTableNameCaseCamel}) Create(ctx context.Context, data map[string]interface{}) (id int64, err error) {
	daoThis := dao{TplPathSuffixCaseCamel}.{TplTableNameCaseCamel}
	id, err = daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseInsert(data)).InsertAndGetId()
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
		tplLogic += `// 修改
func (logicThis *s{TplTableNameCaseCamel}) Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (err error) {
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
	row, _ := result.RowsAffected()
	if row == 0 {
		err = utils.NewErrorCode(ctx, 99999999, ` + "``" + `)
		return
	}
	return
}

`
	}

	if !option.NoDelete {
		tplLogic += `// 删除
func (logicThis *s{TplTableNameCaseCamel}) Delete(ctx context.Context, filter map[string]interface{}) (err error) {
	daoThis := dao{TplPathSuffixCaseCamel}.{TplTableNameCaseCamel}
	result, err := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Delete()
	if err != nil {
		return
	}
	row, _ := result.RowsAffected()
	if row == 0 {
		err = utils.NewErrorCode(ctx, 99999999, ` + "``" + `)
		return
	}
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
	if !option.IsCover && gfile.IsFile(saveFile) {
		return
	}

	for _, column := range tpl.TableColumnList {
		field := column[`Field`].String()
		fieldCaseCamel := gstr.CaseCamel(field)
		switch fieldCaseCamel {
		case `CreatedAt`, `UpdatedAt`, `DeletedAt`: //不处理的字段
		default:
			if (column[`Key`].String() == `PRI` && column[`Extra`].String() == `auto_increment` && field != `id`) || field == `name` || gstr.CaseCamel(field) == tpl.TableNameCaseCamel+`Name` {
				tpl.ControllerAlloweFieldAppend += "`" + field + "`, "
				continue
			}
			//password或passwd后缀
			if gstr.SubStr(fieldCaseCamel, -8) == `Password` || gstr.SubStr(fieldCaseCamel, -6) == `Passwd` {
				tpl.ControllerAlloweFieldDiff += "`" + field + "`, "
				continue
			}
		}
	}
	tpl.ControllerAlloweFieldAppend = gstr.SubStr(tpl.ControllerAlloweFieldAppend, 0, -len(`, `))
	tpl.ControllerAlloweFieldDiff = gstr.SubStr(tpl.ControllerAlloweFieldDiff, 0, -len(`, `))

	tplController := `package controller

import (
	api{TplPathSuffixCaseCamel} "api/api/{TplPathSuffixCaseCamelLower}"
	dao{TplPathSuffixCaseCamel} "api/internal/dao/{TplPathSuffixCaseCamelLower}"
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
	if filter == nil {
		filter = map[string]interface{}{}
	}
	order := [][2]string{{` + "`id`" + `, ` + "`" + `DESC` + "`" + `}}
	if param.Sort.Key != ` + "`" + "`" + ` {
		order[0][0] = param.Sort.Key
	}
	if param.Sort.Order != ` + "`" + "`" + ` {
		order[0][1] = param.Sort.Order
	}
	page := param.Page
	limit := param.Limit
	/**--------参数处理 结束--------**/

	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case ` + "`" + `platform` + "`" + `:
		/**--------权限验证 开始--------**/
		isAuth, _ := service.Action().CheckAuth(r.GetCtx(), ` + "`" + `{TplRawTableNameCaseCamelLower}Look` + "`" + `)
		allowField := []string{` + "`id`, `name`, " + `{TplControllerAlloweFieldAppend}}
		if isAuth {
			allowField = dao{TplPathSuffixCaseCamel}.{TplTableNameCaseCamel}.ColumnArr()
			allowField = append(allowField, ` + "`id`, `name`" + `)`
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
		list, err := service.{TplTableNameCaseCamel}().List(r.GetCtx(), filter, field, order, page, limit)
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
	case ` + "`" + `platform` + "`" + `:
		/**--------参数处理 开始--------**/
		var param *api{TplPathSuffixCaseCamel}.{TplTableNameCaseCamel}InfoReq
		err := r.Parse(&param)
		if err != nil {
			utils.HttpFailJson(r, utils.NewErrorCode(r.GetCtx(), 89999999, err.Error()))
			return
		}

		allowField := dao{TplPathSuffixCaseCamel}.{TplTableNameCaseCamel}.ColumnArr()
		allowField = append(allowField, ` + "`id`, `name`" + `)`
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
		filter := map[string]interface{}{` + "`id`" + `: param.Id}
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
		tplController += `// 新增
func (controllerThis *{TplTableNameCaseCamel}) Create(r *ghttp.Request) {
	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case ` + "`" + `platform` + "`" + `:
		/**--------参数处理 开始--------**/
		var param *api{TplPathSuffixCaseCamel}.{TplTableNameCaseCamel}CreateReq
		err := r.Parse(&param)
		if err != nil {
			utils.HttpFailJson(r, utils.NewErrorCode(r.GetCtx(), 89999999, err.Error()))
			return
		}
		data := gconv.MapDeep(param)
		/**--------参数处理 结束--------**/

		/**--------权限验证 开始--------**/
		_, err = service.Action().CheckAuth(r.GetCtx(), ` + "`" + `{TplRawTableNameCaseCamelLower}Create` + "`" + `)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		/**--------权限验证 结束--------**/

		id, err := service.{TplTableNameCaseCamel}().Create(r.GetCtx(), data)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		utils.HttpSuccessJson(r, map[string]interface{}{` + "`id`" + `: id}, 0)
	}
}

`
	}

	if !option.NoUpdate {
		tplController += `// 修改
func (controllerThis *{TplTableNameCaseCamel}) Update(r *ghttp.Request) {
	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case ` + "`" + `platform` + "`" + `:
		/**--------参数处理 开始--------**/
		var param *api{TplPathSuffixCaseCamel}.{TplTableNameCaseCamel}UpdateReq
		err := r.Parse(&param)
		if err != nil {
			utils.HttpFailJson(r, utils.NewErrorCode(r.GetCtx(), 89999999, err.Error()))
			return
		}
		data := gconv.MapDeep(param)
		delete(data, ` + "`" + `idArr` + "`" + `)
		if len(data) == 0 {
			utils.HttpFailJson(r, utils.NewErrorCode(r.GetCtx(), 89999999, ` + "`" + "`" + `))
			return
		}
		filter := map[string]interface{}{` + "`id`" + `: param.IdArr}
		/**--------参数处理 结束--------**/

		/**--------权限验证 开始--------**/
		_, err = service.Action().CheckAuth(r.GetCtx(), ` + "`" + `{TplRawTableNameCaseCamelLower}Update` + "`" + `)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		/**--------权限验证 结束--------**/

		err = service.{TplTableNameCaseCamel}().Update(r.GetCtx(), filter, data)
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
	case ` + "`" + `platform` + "`" + `:
		/**--------参数处理 开始--------**/
		var param *api{TplPathSuffixCaseCamel}.{TplTableNameCaseCamel}DeleteReq
		err := r.Parse(&param)
		if err != nil {
			utils.HttpFailJson(r, utils.NewErrorCode(r.GetCtx(), 89999999, err.Error()))
			return
		}
		filter := map[string]interface{}{` + "`id`" + `: param.IdArr}
		/**--------参数处理 结束--------**/

		/**--------权限验证 开始--------**/
		_, err = service.Action().CheckAuth(r.GetCtx(), ` + "`" + `{TplRawTableNameCaseCamelLower}Delete` + "`" + `)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		/**--------权限验证 结束--------**/

		err = service.{TplTableNameCaseCamel}().Delete(r.GetCtx(), filter)
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

// 添加路由
func MyGenTplRouter(ctx context.Context, option *MyGenOption, tpl *MyGenTpl) {
	saveFile := gfile.SelfDir() + `/internal/router/platform_admin.go`

	tplView := gfile.GetContents(saveFile)

	replaceStr := ``
	if !option.NoList {
		replaceStr += `
						"/list":   controllerThis.List,`
	}
	if !option.NoUpdate {
		replaceStr += `
						"/info":   controllerThis.Info,`
	}
	if !option.NoCreate {
		replaceStr += `
						"/create":   controllerThis.Create,`
	}
	if !option.NoUpdate {
		replaceStr += `
						"/update":   controllerThis.Update,`
	}
	if !option.NoDelete {
		replaceStr += `
						"/del":   controllerThis.Delete,`
	}
	replaceStr = `group.Group("/` + tpl.PathSuffixCaseCamelLower + `/` + tpl.TableNameCaseCamelLower + `", func(group *ghttp.RouterGroup) {
					controllerThis := controller` + tpl.PathSuffixCaseCamel + `.New` + tpl.TableNameCaseCamel + `()
					group.ALLMap(g.Map{` + replaceStr + `
					})
				})`

	if gstr.Pos(tplView, `"/`+tpl.PathSuffixCaseCamelLower+`/`+tpl.TableNameCaseCamelLower+`"`) == -1 { //路由不存在时新增
		tplView = gstr.Replace(tplView, `/*--------自动代码生成锚点（不允许修改和删除，否则将不能自动生成路由）--------*/`, replaceStr+`
	
				/*--------自动代码生成锚点（不允许修改和删除，否则将不能自动生成路由）--------*/`)
	} else { //路由已存在则替换
		tplView, _ = gregex.ReplaceString(`group.Group\("/`+tpl.PathSuffixCaseCamelLower+`/`+tpl.TableNameCaseCamelLower+`",[\s\S]*
					}\)
				}\)`, replaceStr, tplView)
	}
	gfile.PutContents(saveFile, tplView)
}

// view模板生成Index
func MyGenTplViewIndex(ctx context.Context, option *MyGenOption, tpl *MyGenTpl) {
	saveFile := gfile.SelfDir() + `/../dev/platform/src/views/` + tpl.PathSuffixCaseCamelLower + `/` + tpl.TableNameCaseCamelLower + `/Index.vue`
	if !option.IsCover && gfile.IsFile(saveFile) {
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

	gfile.PutContents(saveFile, tplView)
}

// view模板生成List
func MyGenTplViewList(ctx context.Context, option *MyGenOption, tpl *MyGenTpl) {
	saveFile := gfile.SelfDir() + `/../dev/platform/src/views/` + tpl.PathSuffixCaseCamelLower + `/` + tpl.TableNameCaseCamelLower + `/List.vue`
	if !option.IsCover && gfile.IsFile(saveFile) {
		return
	}

	rawCreatedAtField := ``
	rawUpdatedAtField := ``
	//rawDeletedAtField := ``
	for _, column := range tpl.TableColumnList {
		field := column[`Field`].String()
		fieldCaseCamel := gstr.CaseCamel(field)
		fieldCaseCamelLower := gstr.CaseCamelLower(field)
		fieldCaseSnake := gstr.CaseSnakeFirstUpper(field)
		switch fieldCaseCamel {
		case `CreatedAt`: //不处理的字段
			rawCreatedAtField = field
		case `UpdatedAt`: //不处理的字段
			rawUpdatedAtField = field
		case `DeletedAt`: //不处理的字段
			//rawDeletedAtField = field
		default:
			//主键
			if column[`Key`].String() == `PRI` && column[`Extra`].String() == `auto_increment` {
				continue
			}
			//password或passwd后缀
			if gstr.SubStr(fieldCaseCamel, -8) == `Password` || gstr.SubStr(fieldCaseCamel, -6) == `Passwd` {
				continue
			}
			//pid字段
			if field == `pid` {
				tpl.ViewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('common.name.` + field + `'),
		key: '` + field + `',
		align: 'center',
		width: 150,
	},`
				continue
			}
			//is_stop或isStop字段
			if fieldCaseCamelLower == `isStop` {
				tpl.ViewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('common.name.` + fieldCaseCamelLower + `'),
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
			//gender字段
			if field == `gender` {
				tpl.ViewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('common.name.` + field + `'),
		key: '` + field + `',
		align: 'center',
		width: 100,
		cellRenderer: (props: any): any => {
			let typeObj: any = { 0: 'warning', 1: 'success', 2: 'danger' }
			return [
				h(ElTag as any, {
					type: typeObj[props.rowData.` + field + `]
				}, {
					default: () => (tm('common.status.` + field + `') as any).find((item: any) => { return item.value == props.rowData.` + field + ` })?.label
				})
			]
		}
	},`
				continue
			}
			//avator字段
			if field == `avator` {
				tpl.ViewListColumn += `
	{
        dataKey: '` + field + `',
        title: t('common.name.` + field + `'),
        key: '` + field + `',
        width: 100,
        align: 'center',
        cellRenderer: (props: any): any => {
            if (!props.rowData.` + field + `) {
                return
            }
            //const imageList= JSON.parse(props.rowData.` + field + `)
            const imageList = [props.rowData.` + field + `]
            return [
                h(ElScrollbar, {
                    'wrap-style': 'display: flex; align-items: center;',
                    'view-style': 'margin: auto;',
                }, {
                    default: () => {
                        const content = imageList.map((item) => {
                            return h(ElImage as any, {
                                'style': 'width: 80px;',    //不想显示滚动条，需设置table属性row-height增加行高
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
			//id后缀
			if gstr.SubStr(fieldCaseCamel, -2) == `Id` {
				tpl.ViewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `'),
		key: '` + field + `',
		align: 'center',
		width: 150,
	},`
				continue
			}
			//name或code后缀
			if gstr.SubStr(fieldCaseCamel, -4) == `Name` || gstr.SubStr(fieldCaseCamel, -4) == `Code` {
				tpl.ViewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `'),
		key: '` + field + `',
		width: 150,
		align: 'center',
	},`
				continue
			}
			//mobile或phone后缀
			if gstr.SubStr(fieldCaseCamel, -5) == `Phone` || gstr.SubStr(fieldCaseCamel, -6) == `Mobile` {
				tpl.ViewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `'),
		key: '` + field + `',
		width: 150,
		align: 'center',
	},`
				continue
			}
			//url或link后缀
			if gstr.SubStr(fieldCaseCamel, -3) == `Url` || gstr.SubStr(fieldCaseCamel, -4) == `Link` {
				tpl.ViewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `'),
		key: '` + field + `',
		width: 300,
		align: 'center',
	},`
				continue
			}
			//img或image或cover后缀
			if gstr.SubStr(fieldCaseCamel, -3) == `Img` || gstr.SubStr(fieldCaseCamel, -5) == `Image` || gstr.SubStr(fieldCaseCamel, -5) == `Cover` {
				tpl.ViewListColumn += `
	{
        dataKey: '` + field + `',
        title: t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `'),
        key: '` + field + `',
        width: 100,
        align: 'center',
        cellRenderer: (props: any): any => {
            if (!props.rowData.` + field + `) {
                return
            }`
				if column[`Type`].String() == `json` {
					tpl.ViewListColumn += `
			let imageList: string[]
			if (Array.isArray(props.rowData.` + field + `)) {
				imageList = props.rowData.` + field + `
			} else {
				imageList = JSON.parse(props.rowData.` + field + `)
			}`
				} else {
					tpl.ViewListColumn += `
			const imageList = [props.rowData.` + field + `]`
				}
				tpl.ViewListColumn += `
            return [
                h(ElScrollbar, {
                    'wrap-style': 'display: flex; align-items: center;',
                    'view-style': 'margin: auto;',
                }, {
                    default: () => {
                        const content = imageList.map((item) => {
                            return h(ElImage as any, {
                                'style': 'width: 80px;',    //不想显示滚动条，需设置table属性row-height增加行高
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
			//video后缀
			if gstr.SubStr(fieldCaseCamel, -5) == `Video` {
				tpl.ViewListColumn += `
	{
        dataKey: '` + field + `',
        title: t('common.name.{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.` + field + `'),
        key: '` + field + `',
        width: 100,
        align: 'center',
        cellRenderer: (props: any): any => {
            if (!props.rowData.` + field + `) {
                return
            }`
				if column[`Type`].String() == `json` {
					tpl.ViewListColumn += `
			let videoList: string[]
			if (Array.isArray(props.rowData.` + field + `)) {
				videoList = props.rowData.` + field + `
			} else {
				videoList = JSON.parse(props.rowData.` + field + `)
			}`
				} else {
					tpl.ViewListColumn += `
			const videoList = [props.rowData.` + field + `]`
				}
				tpl.ViewListColumn += `
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
				tpl.ViewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `'),
		key: '` + field + `',
		width: 150,
		align: 'center',
	},`
				continue
			}
			//sort或weight后缀
			if gstr.SubStr(fieldCaseCamel, -4) == `Sort` || gstr.SubStr(fieldCaseCamel, -6) == `Weight` {
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
			//remark后缀
			if gstr.SubStr(fieldCaseCamel, -6) == `Remark` {
				tpl.ViewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('common.name.` + field + `'),
		key: '` + field + `',
		width: 150,
		align: 'center',
		hidden: true
	},`
				continue
			}
			//status后缀
			if gstr.SubStr(fieldCaseCamel, -6) == `Status` {
				tpl.ViewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `'),
		key: '` + field + `',
		align: 'center',
		width: 100,
		cellRenderer: (props: any): any => {
			let typeObj: any = { 0: '', 1: 'success', 2: 'danger', 3: 'info', 4: 'warning' }
			return [
				h(ElTag as any, {
					type: typeObj[props.rowData.` + field + `]
				}, {
					default: () => (tm('common.status.gender') as any).find((item: any) => { return item.value == props.rowData.` + field + ` })?.label
				})
			]
		}
	},`
				continue
			}
			//is_前缀
			if gstr.SubStr(fieldCaseSnake, 0, 3) == `is_` {
				tpl.ViewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `'),
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
			//int类型
			if gstr.Pos(column[`Type`].String(), `int`) != -1 {
				tpl.ViewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `'),
		key: '` + field + `',
		width: 150,
		align: 'center',
	},`
				continue
			}
			//float类型
			if gstr.Pos(column[`Type`].String(), `decimal`) != -1 || gstr.Pos(column[`Type`].String(), `double`) != -1 || gstr.Pos(column[`Type`].String(), `float`) != -1 {
				tpl.ViewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `'),
		key: '` + field + `',
		width: 150,
		align: 'center',
	},`
				continue
			}
			//varchar类型
			if gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
				tpl.ViewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `'),
		key: '` + field + `',
		width: 150,
		align: 'center',
	},`
				continue
			}
			//char类型
			if gstr.Pos(column[`Type`].String(), `char`) != -1 {
				tpl.ViewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `'),
		key: '` + field + `',
		width: 150,
		align: 'center',
	},`
				continue
			}
			//json类型
			if column[`Type`].String() == `json` {
				tpl.ViewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `'),
		key: '` + field + `',
		width: 200,
		align: 'center',
        hidden: true
	},`
				continue
			}
			//datetime和timestamp类型
			if gstr.Pos(column[`Type`].String(), `datetime`) != -1 || gstr.Pos(column[`Type`].String(), `timestamp`) != -1 {
				tpl.ViewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `'),
		key: '` + field + `',
		width: 150,
		align: 'center',
        sortable: true
	},`
				continue
			}
			//date类型
			if gstr.Pos(column[`Type`].String(), `date`) != -1 {
				tpl.ViewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `'),
		key: '` + field + `',
		width: 150,
		align: 'center',
        sortable: true
	},`
				continue
			}
			//默认处理
			tpl.ViewListColumn += `
	{
		dataKey: '` + field + `',
		title: t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `'),
		key: '` + field + `',
		width: 150,
		align: 'center',
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
	if !(option.NoCreate && option.NoUpdate && option.NoDelete) {
		tplView += `
	{
		title: t('common.name.action'),
		key: 'action',
		align: 'center',
		fixed: 'right',
		width: 250,
		cellRenderer: (props: any): any => {
			return [`
		if !option.NoUpdate {
			tplView += `
				h(ElButton, {
					type: 'primary',
					size: 'small',
					onClick: () => handleEditCopy(props.rowData.id)
				}, {
					default: () => [h(AutoiconEpEdit), t('common.edit')]
				}),`
		}
		if !option.NoDelete {
			tplView += `
				h(ElButton, {
					type: 'danger',
					size: 'small',
					onClick: () => handleDelete([props.rowData.id])
				}, {
					default: () => [h(AutoiconEpDelete), t('common.delete')]
				}),`
		}
		if !option.NoCreate {
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
		sort: table.sort.key + ' ' + table.sort.order,
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
		`{TplViewListColumn}`: tpl.ViewListColumn, //先替换这个！内部还有变量要替换
	})
	tplView = gstr.ReplaceByMap(tplView, map[string]string{
		`{TplRawTableNameCaseCamelLower}`: tpl.RawTableNameCaseCamelLower,
		`{TplTableNameCaseCamelLower}`:    tpl.TableNameCaseCamelLower,
		`{TplTableNameCaseCamel}`:         tpl.TableNameCaseCamel,
		`{TplPathSuffixCaseCamelLower}`:   tpl.PathSuffixCaseCamelLower,
		`{TplPathSuffixCaseCamel}`:        tpl.PathSuffixCaseCamel,
	})
	gfile.PutContents(saveFile, tplView)
}

// view模板生成Query
func MyGenTplViewQuery(ctx context.Context, option *MyGenOption, tpl *MyGenTpl) {
	saveFile := gfile.SelfDir() + `/../dev/platform/src/views/` + tpl.PathSuffixCaseCamelLower + `/` + tpl.TableNameCaseCamelLower + `/Query.vue`
	if !option.IsCover && gfile.IsFile(saveFile) {
		return
	}
	for _, column := range tpl.TableColumnList {
		field := column[`Field`].String()
		fieldCaseCamel := gstr.CaseCamel(field)
		fieldCaseCamelLower := gstr.CaseCamelLower(field)
		fieldCaseSnake := gstr.CaseSnakeFirstUpper(field)
		switch fieldCaseCamel {
		case `CreatedAt`, `UpdatedAt`, `DeletedAt`: //不处理的字段
		default:
			//password或passwd后缀
			if gstr.SubStr(fieldCaseCamel, -8) == `Password` || gstr.SubStr(fieldCaseCamel, -6) == `Passwd` {
				continue
			}
			//主键
			if column[`Key`].String() == `PRI` && column[`Extra`].String() == `auto_increment` {
				continue
			}
			//pid字段
			if field == `pid` {
				tpl.ViewQueryField += `
		<ElFormItem prop="` + field + `">
			<MyCascader v-model="queryCommon.data.` + field + `" :placeholder="t('common.name.` + field + `')" :api="{ code: '/{TplPathSuffixCaseCamelLower}/{TplTableNameCaseCamelLower}/tree' }" :defaultOptions="[{ id: 0, name: t('common.name.allTopLevel') }]" />
		</ElFormItem>`
				continue
			}
			//is_stop或isStop字段
			if fieldCaseCamelLower == `isStop` {
				tpl.ViewQueryField += `
		<ElFormItem prop="` + field + `" style="width: 100px;">
			<ElSelectV2 v-model="queryCommon.data.` + field + `" :options="tm('common.status.whether')" :placeholder="t('common.name.` + fieldCaseCamelLower + `')" :clearable="true" />
		</ElFormItem>`
				continue
			}
			//gender字段
			if field == `gender` {
				tpl.ViewQueryField += `
		<ElFormItem prop="` + field + `" style="width: 100px;">
			<ElSelectV2 v-model="queryCommon.data.` + field + `" :options="tm('common.status.` + field + `')" :placeholder="t('common.name.` + field + `')" :clearable="true" />
		</ElFormItem>`
				continue
			}
			//avator字段
			if field == `avator` {
				continue
			}
			//id后缀
			if gstr.SubStr(fieldCaseCamel, -2) == `Id` {
				tpl.ViewQueryField += `
		<ElFormItem prop="` + field + `">
			<MySelect v-model="queryCommon.data.` + field + `" :placeholder="t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `')" :api="{ code: '/{TplPathSuffixCaseCamelLower}/` + gstr.CaseCamelLower(gstr.SubStr(field, 0, -2)) + `/list' }" />
		</ElFormItem>`
				continue
			}
			//name或code后缀
			if gstr.SubStr(fieldCaseCamel, -4) == `Name` || gstr.SubStr(fieldCaseCamel, -4) == `Code` {
				tpl.ViewQueryField += `
		<ElFormItem prop="` + field + `">
			<ElInput v-model="queryCommon.data.` + field + `" :placeholder="t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `')" :clearable="true" />
		</ElFormItem>`
				continue
			}
			//mobile或phone后缀
			if gstr.SubStr(fieldCaseCamel, -5) == `Phone` || gstr.SubStr(fieldCaseCamel, -6) == `Mobile` {
				tpl.ViewQueryField += `
		<ElFormItem prop="` + field + `">
			<ElInput v-model="queryCommon.data.` + field + `" :placeholder="t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `')" :clearable="true" />
		</ElFormItem>`
				continue
			}
			//url或link后缀
			if gstr.SubStr(fieldCaseCamel, -3) == `Url` || gstr.SubStr(fieldCaseCamel, -4) == `Link` {
				tpl.ViewQueryField += `
		<ElFormItem prop="` + field + `">
			<ElInput v-model="queryCommon.data.` + field + `" :placeholder="t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `')" :clearable="true" />
		</ElFormItem>`
				continue
			}
			//img或image或cover后缀
			if gstr.SubStr(fieldCaseCamel, -3) == `Img` || gstr.SubStr(fieldCaseCamel, -5) == `Image` || gstr.SubStr(fieldCaseCamel, -5) == `Cover` {
				continue
			}
			//video后缀
			if gstr.SubStr(fieldCaseCamel, -5) == `Video` {
				continue
			}
			//Ip后缀
			if gstr.SubStr(fieldCaseCamel, -2) == `Ip` {
				tpl.ViewQueryField += `
		<ElFormItem prop="` + field + `">
			<ElInput v-model="queryCommon.data.` + field + `" :placeholder="t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `')" :clearable="true" />
		</ElFormItem>`
				continue
			}
			//sort或weight后缀
			if gstr.SubStr(fieldCaseCamel, -4) == `Sort` || gstr.SubStr(fieldCaseCamel, -6) == `Weight` {
				continue
			}
			//status后缀
			if gstr.SubStr(fieldCaseCamel, -6) == `Status` {
				tpl.ViewQueryField += `
		<ElFormItem prop="` + field + `" style="width: 100px;">
			<ElSelectV2 v-model="queryCommon.data.` + field + `" :options="tm('common.status.gender')" :placeholder="t('common.name.` + field + `')" :clearable="true" />
		</ElFormItem>`
				continue
			}
			//remark后缀
			if gstr.SubStr(fieldCaseCamel, -6) == `Remark` {
				continue
			}
			//is_前缀
			if gstr.SubStr(fieldCaseSnake, 0, 3) == `is_` {
				tpl.ViewQueryField += `
		<ElFormItem prop="` + field + `" style="width: 100px;">
			<ElSelectV2 v-model="queryCommon.data.` + field + `" :options="tm('common.status.whether')" :placeholder="t('common.name.` + field + `')" :clearable="true" />
		</ElFormItem>`
				continue
			}
			//int类型
			if gstr.Pos(column[`Type`].String(), `int`) != -1 {
				tpl.ViewQueryField += `
		<ElFormItem prop="` + field + `">
			<ElInputNumber v-model="queryCommon.data.` + field + `" :placeholder="t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `')" :controls="false" />
		</ElFormItem>`
				continue
			}
			//float类型
			if gstr.Pos(column[`Type`].String(), `decimal`) != -1 || gstr.Pos(column[`Type`].String(), `double`) != -1 || gstr.Pos(column[`Type`].String(), `float`) != -1 {
				tpl.ViewQueryField += `
		<ElFormItem prop="` + field + `">
			<ElInputNumber v-model="queryCommon.data.` + field + `" :placeholder="t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `')" :precision="2" :controls="false" />
		</ElFormItem>`
				continue
			}
			//varchar类型
			if gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
				tpl.ViewQueryField += `
		<ElFormItem prop="` + field + `">
			<ElInput v-model="queryCommon.data.` + field + `" :placeholder="t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `')" :clearable="true" />
		</ElFormItem>`
				continue
			}
			//char类型
			if gstr.Pos(column[`Type`].String(), `char`) != -1 {
				tpl.ViewQueryField += `
		<ElFormItem prop="` + field + `">
			<ElInput v-model="queryCommon.data.` + field + `" :placeholder="t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `')" :clearable="true" />
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
			tpl.ViewQueryField += `
		<ElFormItem prop="` + field + `">
			<ElInput v-model="queryCommon.data.` + field + `" :placeholder="t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `')" :clearable="true" />
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
			new Date(date.getFullYear(), date.getMonth(), date.getDate(), 0, 0, 0),
			new Date(date.getFullYear(), date.getMonth(), date.getDate(), 23, 59, 59),
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
		</ElFormItem>{TplViewQueryField}
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

	tplView = gstr.ReplaceByMap(tplView, map[string]string{
		`{TplViewQueryField}`: tpl.ViewQueryField, //先替换这个！内部还有变量要替换
	})
	tplView = gstr.ReplaceByMap(tplView, map[string]string{
		`{TplRawTableNameCaseCamelLower}`: tpl.RawTableNameCaseCamelLower,
		`{TplTableNameCaseCamelLower}`:    tpl.TableNameCaseCamelLower,
		`{TplTableNameCaseCamel}`:         tpl.TableNameCaseCamel,
		`{TplPathSuffixCaseCamelLower}`:   tpl.PathSuffixCaseCamelLower,
		`{TplPathSuffixCaseCamel}`:        tpl.PathSuffixCaseCamel,
	})
	gfile.PutContents(saveFile, tplView)
}

// view模板生成Save
func MyGenTplViewSave(ctx context.Context, option *MyGenOption, tpl *MyGenTpl) {
	if option.NoCreate && option.NoUpdate {
		return
	}
	saveFile := gfile.SelfDir() + `/../dev/platform/src/views/` + tpl.PathSuffixCaseCamelLower + `/` + tpl.TableNameCaseCamelLower + `/Save.vue`
	if !option.IsCover && gfile.IsFile(saveFile) {
		return
	}

	for _, column := range tpl.TableColumnList {
		field := column[`Field`].String()
		fieldCaseCamel := gstr.CaseCamel(field)
		fieldCaseCamelLower := gstr.CaseCamelLower(field)
		fieldCaseSnake := gstr.CaseSnakeFirstUpper(field)
		result, _ := gregex.MatchString(`.*\((\d*)\)`, column[`Type`].String())
		switch fieldCaseCamel {
		case `CreatedAt`, `UpdatedAt`, `DeletedAt`: //不处理的字段
		default:
			//主键
			if column[`Key`].String() == `PRI` && column[`Extra`].String() == `auto_increment` {
				continue
			}
			//password或passwd后缀
			if gstr.SubStr(fieldCaseCamel, -8) == `Password` || gstr.SubStr(fieldCaseCamel, -6) == `Passwd` {
				tpl.ViewSaveRule += `
		` + field + `: [
			{ type: 'string', required: computed((): boolean => { return saveForm.data.idArr?.length ? false : true; }), min: 1, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 30 }) }
		],`
				tpl.ViewSaveField += `
				<ElFormItem :label="t('common.name.` + field + `')" prop="` + field + `">
                    <ElInput v-model="saveForm.data.` + field + `" :placeholder="t('common.name.` + field + `')" minlength="1"
                        maxlength="30" :show-word-limit="true" :clearable="true" :show-password="true"
                        style="max-width: 250px;" />
                    <label v-if="saveForm.data.id">
                        <ElAlert :title="t('common.tip.notRequired')" type="info" :show-icon="true" :closable="false" />
                    </label>
                </ElFormItem>`
				continue
			}
			//pid字段
			if field == `pid` {
				tpl.ViewSaveRule += `
		` + field + `: [
			{ type: 'integer', min: 0, trigger: 'change', message: t('validation.select') }
		],`
				tpl.ViewSaveField += `
				<ElFormItem :label="t('common.name.` + field + `')" prop="` + field + `">
                    <MyCascader v-model="saveForm.data.` + field + `" :api="{ code: '/{TplPathSuffixCaseCamelLower}/{TplTableNameCaseCamelLower}/tree', param: { filter: { excId: saveForm.data.id } } }" :defaultOptions="[{ id: 0, name: t('common.name.without') }]" :clearable="false" />
                </ElFormItem>`
				continue
			}
			//is_stop或isStop字段
			if fieldCaseCamelLower == `isStop` {
				tpl.ViewSaveRule += `
		` + field + `: [
            { type: 'enum', enum: [0, 1], trigger: 'change', message: t('validation.select') }
        ],`
				tpl.ViewSaveField += `
				<ElFormItem :label="t('common.name.` + fieldCaseCamelLower + `')" prop="` + field + `">
                    <ElSwitch v-model="saveForm.data.` + field + `" :active-value="1" :inactive-value="0" :inline-prompt="true" :active-text="t('common.yes')" :inactive-text="t('common.no')"
                        style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success);" />
                </ElFormItem>`
				continue
			}
			//gender字段
			if field == `gender` {
				tpl.ViewSaveRule += `
		` + field + `: [
			{ type: 'enum', enum: [0, 1, 2], trigger: 'change', message: t('validation.select') }
		],`
				tpl.ViewSaveField += `
				<ElFormItem :label="t('common.name.` + field + `')" prop="` + field + `">
					<ElSelectV2 v-model="saveForm.data.` + field + `" :options="tm('common.status.` + field + `')" :placeholder="t('common.name.` + field + `')" :clearable="true" />
				</ElFormItem>`
				continue
			}
			//avator字段
			if field == `avator` {
				tpl.ViewSaveRule += `
		` + field + `: [
			{ type: 'url', trigger: 'change', message: t('validation.upload') },
			{ type: 'string', min: 1, max: ` + result[1] + `, trigger: 'blur', message: t('validation.between.string', { min: 1, max: ` + result[1] + ` }) }
        ],`
				tpl.ViewSaveField += `
				<ElFormItem :label="t('common.name.` + field + `')" prop="` + field + `">
                    <MyUpload v-model="saveForm.data.` + field + `" accept="image/*" />
                </ElFormItem>`
				continue
			}
			//id后缀
			if gstr.SubStr(fieldCaseCamel, -2) == `Id` {
				tpl.ViewSaveRule += `
		` + field + `: [
			{ type: 'integer', min: 1, trigger: 'change', message: t('validation.select') }
		],`
				tpl.ViewSaveField += `
				<ElFormItem :label="t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `')" prop="` + field + `">
                    <MySelect v-model="saveForm.data.` + field + `" :api="{ code: '/{TplPathSuffixCaseCamelLower}/` + gstr.CaseCamelLower(gstr.SubStr(field, 0, -2)) + `/list' }" />
                </ElFormItem>`
				continue
			}
			//name或code后缀
			if gstr.SubStr(fieldCaseCamel, -4) == `Name` || gstr.SubStr(fieldCaseCamel, -4) == `Code` {
				tpl.ViewSaveRule += `
		` + field + `: [
			{ type: 'string', required: true, min: 1, max: ` + result[1] + `, trigger: 'blur', message: t('validation.between.string', { min: 1, max: ` + result[1] + ` }) },
			{ pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') }
		],`
				tpl.ViewSaveField += `
				<ElFormItem :label="t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `')" prop="` + field + `">
					<ElInput v-model="saveForm.data.` + field + `" :placeholder="t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `')" minlength="1" maxlength="` + result[1] + `" :show-word-limit="true" :clearable="true" />
				</ElFormItem>`
				continue
			}
			//mobile或phone后缀
			if gstr.SubStr(fieldCaseCamel, -5) == `Phone` || gstr.SubStr(fieldCaseCamel, -6) == `Mobile` {
				tpl.ViewSaveRule += `
		` + field + `: [
			{ type: 'string', min: 1, max: ` + result[1] + `, trigger: 'blur', message: t('validation.between.string', { min: 1, max: ` + result[1] + ` }) },
			{ pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') }
		],`
				tpl.ViewSaveField += `
				<ElFormItem :label="t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `')" prop="` + field + `">
					<ElInput v-model="saveForm.data.` + field + `" :placeholder="t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `')" minlength="1" maxlength="` + result[1] + `" :show-word-limit="true" :clearable="true" />
				</ElFormItem>`
				continue
			}
			//url或link后缀
			if gstr.SubStr(fieldCaseCamel, -3) == `Url` || gstr.SubStr(fieldCaseCamel, -4) == `Link` {
				tpl.ViewSaveRule += `
		` + field + `: [
			{ type: 'url', trigger: 'change', message: t('validation.url') },
			{ type: 'string', min: 1, max: ` + result[1] + `, trigger: 'change', message: t('validation.between.string', { min: 1, max: ` + result[1] + ` }) }
		],`
				tpl.ViewSaveField += `
				<ElFormItem :label="t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `')" prop="` + field + `">
					<ElInput v-model="saveForm.data.` + field + `" :placeholder="t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `')" minlength="1" maxlength="` + result[1] + `" :show-word-limit="true" :clearable="true" />
				</ElFormItem>`
				continue
			}
			//img或image或cover后缀
			if gstr.SubStr(fieldCaseCamel, -3) == `Img` || gstr.SubStr(fieldCaseCamel, -5) == `Image` || gstr.SubStr(fieldCaseCamel, -5) == `Cover` {
				if column[`Type`].String() == `json` {
					tpl.ViewSaveRule += `
		` + field + `: [
            { type: 'array', trigger: 'change', defaultField: { type: 'url', message: t('validation.url') }, message: t('validation.upload') },
            { type: 'array', min: 1, trigger: 'change', message: t('validation.min.upload', { min: 1 }) },
            { type: 'array', max: 10, trigger: 'change', message: t('validation.max.upload', { max: 10 }) }
        ],`
					tpl.ViewSaveField += `
				<ElFormItem :label="t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `')" prop="` + field + `">
					<MyUpload v-model="saveForm.data.` + field + `" accept="image/*" :multiple="true" />
				</ElFormItem>`
				} else {
					tpl.ViewSaveRule += `
		` + field + `: [
			{ type: 'url', trigger: 'change', message: t('validation.upload') },
			{ type: 'string', min: 1, max: ` + result[1] + `, trigger: 'blur', message: t('validation.between.string', { min: 1, max: ` + result[1] + ` }) }
        ],`
					tpl.ViewSaveField += `
				<ElFormItem :label="t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `')" prop="` + field + `">
                    <MyUpload v-model="saveForm.data.` + field + `" accept="image/*" />
                </ElFormItem>`
				}
				continue
			}
			//video后缀
			if gstr.SubStr(fieldCaseCamel, -5) == `Video` {
				if column[`Type`].String() == `json` {
					tpl.ViewSaveRule += `
		` + field + `: [
            { type: 'array', trigger: 'change', defaultField: { type: 'url', message: t('validation.url') }, message: t('validation.upload') },
            { type: 'array', min: 1, trigger: 'change', message: t('validation.min.upload', { min: 1 }) },
            { type: 'array', max: 10, trigger: 'change', message: t('validation.max.upload', { max: 10 }) }
        ],`
					tpl.ViewSaveField += `
				<ElFormItem :label="t('common.name.{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.` + field + `')" prop="` + field + `">
					<MyUpload v-model="saveForm.data.` + field + `" accept="video/*" :isImage="false" :multiple="true" />
				</ElFormItem>`
				} else {
					tpl.ViewSaveRule += `
		` + field + `: [
			{ type: 'url', trigger: 'change', message: t('validation.upload') },
			{ type: 'string', min: 1, max: ` + result[1] + `, trigger: 'blur', message: t('validation.between.string', { min: 1, max: ` + result[1] + ` }) }
        ],`
					tpl.ViewSaveField += `
				<ElFormItem :label="t('common.name.{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.` + field + `')" prop="` + field + `">
                    <MyUpload v-model="saveForm.data.` + field + `" accept="video/*" :isImage="false" />
                </ElFormItem>`
				}
				continue
			}
			//Ip后缀
			if gstr.SubStr(fieldCaseCamel, -2) == `Ip` {
				tpl.ViewSaveRule += `
		` + field + `: [
			{ type: 'string', min: 1, max: ` + result[1] + `, trigger: 'blur', message: t('validation.between.string', { min: 1, max: ` + result[1] + ` }) }
		],`
				tpl.ViewSaveField += `
				<ElFormItem :label="t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `')" prop="` + field + `">
					<ElInput v-model="saveForm.data.` + field + `" :placeholder="t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `')" minlength="1" maxlength="` + result[1] + `" :show-word-limit="true" :clearable="true" />
				</ElFormItem>`
				continue
			}
			//sort或weight后缀
			if gstr.SubStr(fieldCaseCamel, -4) == `Sort` || gstr.SubStr(fieldCaseCamel, -6) == `Weight` {
				tpl.ViewSaveRule += `
		` + field + `: [
			{ type: 'integer', min: 0, max: 100, trigger: 'change', message: t('validation.between.number', { min: 0, max: 100 }) }
		],`
				tpl.ViewSaveField += `
				<ElFormItem :label="t('common.name.` + field + `')" prop="` + field + `">
                    <ElInputNumber v-model="saveForm.data.` + field + `" :precision="0" :min="0" :max="100" :step="1"
                        :step-strictly="true" controls-position="right" :value-on-clear="50" />
                    <label>
                        <ElAlert :title="t('common.tip.` + field + `')" type="info" :show-icon="true" :closable="false" />
                    </label>
                </ElFormItem>`
				continue
			}
			//status后缀
			if gstr.SubStr(fieldCaseCamel, -6) == `Status` {
				tpl.ViewSaveRule += `
		` + field + `: [
			{ type: 'enum', enum: [0, 1, 2], trigger: 'change', message: t('validation.select') }
		],`
				tpl.ViewSaveField += `
				<ElFormItem :label="t('common.name.` + field + `')" prop="` + field + `">
					<ElSelectV2 v-model="saveForm.data.` + field + `" :options="tm('common.status.gender')" :placeholder="t('common.name.` + field + `')" :clearable="true" />
				</ElFormItem>`
				continue
			}
			//remark后缀
			if gstr.SubStr(fieldCaseCamel, -6) == `Remark` {
				tpl.ViewSaveRule += `
		` + field + `: [
			{ type: 'string', min: 1, max: ` + result[1] + `, trigger: 'blur', message: t('validation.between.string', { min: 1, max: ` + result[1] + ` }) },
		],`
				tpl.ViewSaveField += `
				<ElFormItem :label="t('common.name.` + field + `')" prop="` + field + `">
					<ElInput v-model="saveForm.data.` + field + `" type="textarea" :autosize="{ minRows: 3 }" />
				</ElFormItem>`
				continue
			}
			//is_前缀
			if gstr.SubStr(fieldCaseSnake, 0, 3) == `is_` {
				tpl.ViewSaveRule += `
		` + field + `: [
            { type: 'enum', enum: [0, 1], trigger: 'change', message: t('validation.select') }
        ],`
				tpl.ViewSaveField += `
				<ElFormItem :label="t('common.name.` + field + `')" prop="` + field + `">
                    <ElSwitch v-model="saveForm.data.` + field + `" :active-value="1" :inactive-value="0" :inline-prompt="true" :active-text="t('common.yes')" :inactive-text="t('common.no')"
                        style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success);" />
                </ElFormItem>`
				continue
			}
			//int类型
			if gstr.Pos(column[`Type`].String(), `int`) != -1 {
				tpl.ViewSaveRule += `
		` + field + `: [
			{ type: 'integer', trigger: 'change', message: '' }
		],`
				tpl.ViewSaveField += `
				<ElFormItem :label="t('common.name.` + field + `')" prop="` + field + `">
					<ElInputNumber v-model="saveForm.data.` + field + `" :placeholder="t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `')" :controls="false"/>
				</ElFormItem>`
				continue
			}
			//float类型
			if gstr.Pos(column[`Type`].String(), `decimal`) != -1 || gstr.Pos(column[`Type`].String(), `double`) != -1 || gstr.Pos(column[`Type`].String(), `float`) != -1 {
				tpl.ViewSaveRule += `
		` + field + `: [
			{ type: 'float', trigger: 'change', message: '' }
		],`
				tpl.ViewSaveField += `
				<ElFormItem :label="t('common.name.` + field + `')" prop="` + field + `">
					<ElInputNumber v-model="saveForm.data.` + field + `" :placeholder="t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `')" :precision="2" :controls="false"/>
				</ElFormItem>`
				continue
			}
			//varchar类型
			if gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
				tpl.ViewSaveRule += `
		` + field + `: [
			{ type: 'string', min: 1, max: ` + result[1] + `, trigger: 'blur', message: t('validation.between.string', { min: 1, max: ` + result[1] + ` }) },
		],`
				tpl.ViewSaveField += `
				<ElFormItem :label="t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `')" prop="` + field + `">
					<ElInput v-model="saveForm.data.` + field + `" :placeholder="t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `')" minlength="1" maxlength="` + result[1] + `" :show-word-limit="true" :clearable="true" />
				</ElFormItem>`
				continue
			}
			//char类型
			if gstr.Pos(column[`Type`].String(), `char`) != -1 {
				tpl.ViewSaveRule += `
		` + field + `: [
			{ type: 'string', min: ` + result[1] + `, max: ` + result[1] + `, trigger: 'blur', message: t('validation.between.string', { min: ` + result[1] + `, max: ` + result[1] + ` }) },
		],`
				tpl.ViewSaveField += `
				<ElFormItem :label="t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `')" prop="` + field + `">
					<ElInput v-model="saveForm.data.` + field + `" :placeholder="t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `')" minlength="` + result[1] + `" maxlength="` + result[1] + `" :show-word-limit="true" :clearable="true" />
				</ElFormItem>`
				continue
			}
			//json类型
			if column[`Type`].String() == `json` {
				tpl.ViewSaveRule += `
		` + field + `: [
			{
				type: 'object',
				fields: {
					xxxx: { type: 'string', min: 1, message: 'xxxx' + t('validation.min.string', { min: 1 }) }
				},
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
				tpl.ViewSaveField += `
				<ElFormItem :label="t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `')" prop="` + field + `">
					<ElInput v-model="saveForm.data.` + field + `" type="textarea" :autosize="{ minRows: 3 }" />
				</ElFormItem>`
				continue
			}
			//datetime和timestamp类型
			if gstr.Pos(column[`Type`].String(), `datetime`) != -1 || gstr.Pos(column[`Type`].String(), `timestamp`) != -1 {
				tpl.ViewSaveRule += `
		` + field + `: [],`
				tpl.ViewSaveField += `
				<ElFormItem :label="t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `')" prop="` + field + `">
					<ElDatePicker v-model="saveForm.data.` + field + `" type="datetime" :placeholder="t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `')" format="YYYY-MM-DD HH:mm:ss" value-format="YYYY-MM-DD HH:mm:ss" />
				</ElFormItem>`
				continue
			}
			//date类型
			if gstr.Pos(column[`Type`].String(), `date`) != -1 {
				tpl.ViewSaveRule += `
		` + field + `: [],`
				tpl.ViewSaveField += `
				<ElFormItem :label="t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `')" prop="` + field + `">
					<ElDatePicker v-model="saveForm.data.` + field + `" type="date" :placeholder="t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `')" format="YYYY-MM-DD" value-format="YYYY-MM-DD" />
				</ElFormItem>`
				continue
			}
			//默认处理
			tpl.ViewSaveRule += `
		` + field + `: [],`
			tpl.ViewSaveField += `
				<ElFormItem :label="t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `')" prop="` + field + `">
					<ElInput v-model="saveForm.data.` + field + `" :placeholder="t('{TplPathSuffixCaseCamelLower}.{TplTableNameCaseCamelLower}.name.` + field + `')" :show-word-limit="true" :clearable="true" />
				</ElFormItem>`
		}
	}

	tplView := `<script setup lang="ts">
const { t, tm } = useI18n()

const saveCommon = inject('saveCommon') as { visible: boolean, title: string, data: { [propName: string]: any } }
const listCommon = inject('listCommon') as { ref: any }

const saveForm = reactive({
	ref: null as any,
	loading: false,
	data: {
		...saveCommon.data
	} as { [propName: string]: any },
	rules: {{TplViewSaveRule}
	} as any,
	submit: () => {
		saveForm.ref.validate(async (valid: boolean) => {
			if (!valid) {
				return false
			}
			saveForm.loading = true
			const param = removeEmptyOfObj(saveForm.data, false)
			try {
				if (param?.idArr?.length > 0) {
					await request('/{TplPathSuffixCaseCamelLower}/{TplTableNameCaseCamelLower}/update', param, true)
				} else {
					await request('/{TplPathSuffixCaseCamelLower}/{TplTableNameCaseCamelLower}/create', param, true)
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
				label-width="auto" :status-icon="true" :scroll-to-error="true">{TplViewSaveField}
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

	tplView = gstr.ReplaceByMap(tplView, map[string]string{
		`{TplViewSaveRule}`:  tpl.ViewSaveRule,  //先替换这个！内部还有变量要替换
		`{TplViewSaveField}`: tpl.ViewSaveField, //先替换这个！内部还有变量要替换
	})
	tplView = gstr.ReplaceByMap(tplView, map[string]string{
		`{TplRawTableNameCaseCamelLower}`: tpl.RawTableNameCaseCamelLower,
		`{TplTableNameCaseCamelLower}`:    tpl.TableNameCaseCamelLower,
		`{TplTableNameCaseCamel}`:         tpl.TableNameCaseCamel,
		`{TplPathSuffixCaseCamelLower}`:   tpl.PathSuffixCaseCamelLower,
		`{TplPathSuffixCaseCamel}`:        tpl.PathSuffixCaseCamel,
	})
	gfile.PutContents(saveFile, tplView)
}

// view模板生成I18n
func MyGenTplViewI18n(ctx context.Context, option *MyGenOption, tpl *MyGenTpl) {
	saveFile := gfile.SelfDir() + `/../dev/platform/src/i18n/language/zh-cn/` + tpl.PathSuffixCaseCamelLower + `/` + tpl.TableNameCaseCamelLower + `.ts`
	if !option.IsCover && gfile.IsFile(saveFile) {
		return
	}

	for _, column := range tpl.TableColumnList {
		field := column[`Field`].String()
		fieldCaseCamel := gstr.CaseCamel(field)
		fieldCaseCamelLower := gstr.CaseCamelLower(field)
		comment := gstr.Trim(gstr.ReplaceByArray(column[`Comment`].String(), g.SliceStr{
			"\n", " ",
			"\r", " ",
		}))
		switch fieldCaseCamel {
		case `CreatedAt`, `UpdatedAt`, `DeletedAt`: //不处理的字段
		default:
			//主键
			if column[`Key`].String() == `PRI` && column[`Extra`].String() == `auto_increment` {
				continue
			}
			if !garray.NewStrArrayFrom([]string{`remark`, `isStop`, `sort`, `pid`, `account`, `password`, `phone`}).Contains(fieldCaseCamelLower) {
				tpl.ViewI18nField += `
	` + field + `: '` + comment + `',`
			}
		}
	}
	tplView := `export default {
    name:{{TplViewI18nField}
    },
}`
	tplView = gstr.ReplaceByMap(tplView, map[string]string{
		`{TplViewI18nField}`: tpl.ViewI18nField, //先替换这个！内部还有变量要替换
	})
	gfile.PutContents(saveFile, tplView)
}

// view模板生成路由
func MyGenTplViewRouter(ctx context.Context, option *MyGenOption, tpl *MyGenTpl) {
	saveFile := gfile.SelfDir() + `/../dev/platform/src/router/index.ts`

	tplView := gfile.GetContents(saveFile)

	replaceStr := `{
                path: '/` + tpl.PathSuffixCaseCamelLower + `/` + tpl.TableNameCaseCamelLower + `',
                component: async () => {
                    const component = await import('@/views/` + tpl.PathSuffixCaseCamelLower + `/` + tpl.TableNameCaseCamelLower + `/Index.vue')
                    component.default.name = '/` + tpl.PathSuffixCaseCamelLower + `/` + tpl.TableNameCaseCamelLower + `'
                    return component
                },
                meta: { isAuth: true, keepAlive: true, componentName: '/` + tpl.PathSuffixCaseCamelLower + `/` + tpl.TableNameCaseCamelLower + `' }
            },`

	if gstr.Pos(tplView, `'/`+tpl.PathSuffixCaseCamelLower+`/`+tpl.TableNameCaseCamelLower+`'`) == -1 { //路由不存在时新增
		tplView = gstr.Replace(tplView, `/*--------自动代码生成锚点（不允许修改和删除，否则将不能自动生成路由）--------*/`, replaceStr+`
            /*--------自动代码生成锚点（不允许修改和删除，否则将不能自动生成路由）--------*/`)
	} else { //路由已存在则替换
		tplView, _ = gregex.ReplaceString(`\{
                path: '/`+tpl.PathSuffixCaseCamelLower+`/`+tpl.TableNameCaseCamelLower+`',[\s\S]*'/`+tpl.PathSuffixCaseCamelLower+`/`+tpl.TableNameCaseCamelLower+`' \}
            \},`, replaceStr, tplView)
	}
	gfile.PutContents(saveFile, tplView)
}
