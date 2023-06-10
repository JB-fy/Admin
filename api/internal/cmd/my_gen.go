package cmd

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gvar"
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
	TablePrefix                string
	TableNameCaseCamel         string //去除前缀表名（小驼峰）
	ApiFilterColumn            string //api列表过滤
	ApiSaveColumn              string //api创建更新
}

func MyGenFunc(ctx context.Context, parser *gcmd.Parser) (err error) {
	optMap := parser.GetOptAll()
	option := MyGenOption{}
	gconv.Struct(optMap, &option)
	var db gdb.DB
	option.DbGroup = `default`
	if option.DbGroup == `` {
		option.DbGroup = gcmd.Scan("> 请输入db分组,默认(default):\n")
		if option.DbGroup == `` {
			option.DbGroup = `default`
		}
	}
	for {
		err = g.Try(ctx, func(ctx context.Context) {
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
	option.DbTable = `auth_test`
	if option.DbTable == `` {
		option.DbTable = gcmd.Scan("> 请输入db表:\n")
	}
	for {
		if option.DbTable != `` && garray.NewStrArrayFrom(tableArr).Contains(option.DbTable) {
			break
		}
		option.DbTable = gcmd.Scan("> db表不存在，请重新输入:\n")
	}
	/* _, ok := optMap[`removePrefix`]
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
	noList, ok := optMap[`noList`]
	if !ok {
		noList = gcmd.Scan("> 是否生成列表接口,默认(yes):\n")
	}
noListEnd:
	for {
		switch noList {
		case ``, `yes`:
			option.NoList = true
			break noListEnd
		case `no`:
			option.NoList = true
			break noListEnd
		default:
			noList = gcmd.Scan("> 输入错误，请重新输入，是否生成列表接口,默认(yes):\n")
		}
	}
	noCreate, ok := optMap[`noCreate`]
	if !ok {
		noCreate = gcmd.Scan("> 是否生成创建接口,默认(yes):\n")
	}
noCreateEnd:
	for {
		switch noCreate {
		case ``, `yes`:
			option.NoCreate = true
			break noCreateEnd
		case `no`:
			option.NoCreate = true
			break noCreateEnd
		default:
			noCreate = gcmd.Scan("> 输入错误，请重新输入，是否生成创建接口,默认(yes):\n")
		}
	}
	noUpdate, ok := optMap[`noUpdate`]
	if !ok {
		noUpdate = gcmd.Scan("> 是否生成更新接口,默认(yes):\n")
	}
noUpdateEnd:
	for {
		switch noUpdate {
		case ``, `yes`:
			option.NoUpdate = true
			break noUpdateEnd
		case `no`:
			option.NoUpdate = true
			break noUpdateEnd
		default:
			noUpdate = gcmd.Scan("> 输入错误，请重新输入，是否生成更新接口,默认(yes):\n")
		}
	}
	noDelete, ok := optMap[`noDelete`]
	if !ok {
		noDelete = gcmd.Scan("> 是否生成删除接口,默认(yes):\n")
	}
noDeleteEnd:
	for {
		switch noDelete {
		case ``, `yes`:
			option.NoDelete = true
			break noDeleteEnd
		case `no`:
			option.NoDelete = true
			break noDeleteEnd
		default:
			noDelete = gcmd.Scan("> 输入错误，请重新输入，是否生成删除接口,默认(yes):\n")
		}
	}

	// AdminId                       *uint  `c:"adminId,omitempty" p:"adminId" v:"integer|min:1"`
	// Account                       string `c:"account,omitempty" p:"account" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	// Phone                         string `c:"phone,omitempty" p:"phone" v:"phone"`
	// RoleId                        *uint  `c:"roleId,omitempty" p:"roleId" v:"integer|min:1"`

	// Account   *string `c:"account,omitempty" p:"account" v:"required-without:Phone|length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	// Phone     *string `c:"phone,omitempty" p:"phone" v:"required-without:Account|phone"`
	// Password  *string `c:"password,omitempty" p:"password" v:"required|size:32|regex:^[\\p{L}\\p{N}_-]+$"`
	// RoleIdArr *[]uint `c:"roleIdArr,omitempty" p:"roleIdArr" v:"required|distinct|foreach|integer|foreach|min:1"`
	// Nickname  *string `c:"nickname,omitempty" p:"nickname" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	// Avatar    *string `c:"avatar,omitempty" p:"avatar" v:"url|length:1,120"`
	// IsStop    *uint   `c:"isStop,omitempty" p:"isStop" v:"integer|in:0,1"`

	tpl := MyGenTpl{
		RawTableNameCaseCamelLower: gstr.CaseCamelLower(option.DbTable),
		TablePrefix:                gstr.CaseCamel(option.RemovePrefix),
		TableNameCaseCamel:         gstr.CaseCamel(gstr.Replace(option.DbTable, option.RemovePrefix, ``, 1)),
	}
	tableColumnList, _ := db.GetAll(ctx, `SHOW FULL COLUMNS FROM `+option.DbTable)

	for _, column := range tableColumnList {
		field := column[`Field`].String()
		fieldCaseCamel := gstr.CaseCamel(field)
		switch fieldCaseCamel {
		case `UpdatedAt`, `CreatedAt`, `DeletedAt`:
		default:
			//主键
			if column[`Key`].String() == `PRI` && column[`Extra`].String() == `auto_increment` {
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
			if gstr.ToLower(gstr.SubStr(field, 2)) == `is` {
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
				tpl.ApiFilterColumn += fieldCaseCamel + ` string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"ip"` + "` //" + column[`Comment`].String() + "\n"
				tpl.ApiSaveColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"ip"` + "` //" + column[`Comment`].String() + "\n"
				continue
			}
			//含有mobile和phone字符
			if gstr.Pos(gstr.ToLower(field), `mobile`) != -1 || gstr.Pos(gstr.ToLower(field), `phone`) != -1 {
				tpl.ApiFilterColumn += fieldCaseCamel + ` string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"phone"` + "` //" + column[`Comment`].String() + "\n"
				tpl.ApiSaveColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"phone"` + "` //" + column[`Comment`].String() + "\n"
				continue
			}
			//含有url和link字符
			if gstr.Pos(gstr.ToLower(field), `url`) != -1 || gstr.Pos(gstr.ToLower(field), `link`) != -1 {
				tpl.ApiFilterColumn += fieldCaseCamel + ` string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"url"` + "` //" + column[`Comment`].String() + "\n"
				tpl.ApiSaveColumn += fieldCaseCamel + ` *string ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"url"` + "` //" + column[`Comment`].String() + "\n"
				continue
			}
			if gstr.Pos(column[`Type`].String(), `int`) != -1 {
				if gstr.Pos(column[`Type`].String(), `unsigned`) != -1 {
					tpl.ApiFilterColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"integer|min:1"` + "` //" + column[`Comment`].String() + "\n"
					tpl.ApiSaveColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"integer|min:1"` + "` //" + column[`Comment`].String() + "\n"
				} else {
					tpl.ApiFilterColumn += fieldCaseCamel + ` *int ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"integer"` + "` //" + column[`Comment`].String() + "\n"
					tpl.ApiSaveColumn += fieldCaseCamel + ` *int ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"integer"` + "` //" + column[`Comment`].String() + "\n"
				}
			} else if gstr.Pos(column[`Type`].String(), `decimal`) != -1 || gstr.Pos(column[`Type`].String(), `double`) != -1 || gstr.Pos(column[`Type`].String(), `float`) != -1 {
				if gstr.Pos(column[`Type`].String(), `unsigned`) != -1 {
					tpl.ApiFilterColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"integer|min:1"` + "` //" + column[`Comment`].String() + "\n"
					tpl.ApiSaveColumn += fieldCaseCamel + ` *uint ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"integer|min:1"` + "` //" + column[`Comment`].String() + "\n"
				} else {
					tpl.ApiFilterColumn += fieldCaseCamel + ` *int ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"integer"` + "` //" + column[`Comment`].String() + "\n"
					tpl.ApiSaveColumn += fieldCaseCamel + ` *int ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"integer"` + "` //" + column[`Comment`].String() + "\n"
				}
				tpl.ApiFilterColumn += fieldCaseCamel + ` *float ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"integer"` + "` //" + column[`Comment`].String() + "\n"
				tpl.ApiSaveColumn += fieldCaseCamel + ` *float ` + "`" + `c:"` + field + `,omitempty" p:"` + field + `" v:"integer"` + "` //" + column[`Comment`].String() + "\n"
			} else if gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
				column[`relType`] = gvar.New(`varchar`)
				result, _ := gregex.MatchString(`varchar\((\d*)\)`, column[`Type`].String())
				column[`maxLength`] = gvar.New(result[1])
			} else if gstr.Pos(column[`Type`].String(), `char`) != -1 {
				column[`relType`] = gvar.New(`char`)
				result, _ := gregex.MatchString(`char\((\d*)\)`, column[`Type`].String())
				column[`maxLength`] = gvar.New(result[1])
			} else if gstr.Pos(column[`Type`].String(), `json`) != -1 {
				column[`relType`] = gvar.New(`json`)
			} else {
				column[`relType`] = gvar.New(`string`)
			}
		}
	}

	fmt.Println(tableColumnList)
	path := gfile.SelfDir()
	tplApi := gfile.GetContents(path + `/resource/template/gen_template_api.txt`)
	tplApi = gstr.Replace(tplApi, `{TplTableNameCaseCamel}`, tpl.TableNameCaseCamel)
	fmt.Println(tplApi)
	return
}
