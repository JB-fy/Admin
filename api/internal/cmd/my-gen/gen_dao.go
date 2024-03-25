package my_gen

import (
	"api/internal/utils"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type myGenDao struct {
	primaryKeyFunction string

	importDao []string

	filterParse []string

	fieldParse []string
	fieldHook  []string

	insertParseBefore []string
	insertParse       []string
	insertHook        []string

	updateParse      []string
	updateHookBefore []string
	updateHookAfter  []string

	groupParse []string

	orderParse []string

	joinParse []string
}

type myGenDaoField struct {
	importDao []string

	filterParse myGenDataSliceHandler

	fieldParse myGenDataSliceHandler
	fieldHook  myGenDataSliceHandler

	insertParseBefore myGenDataSliceHandler
	insertParse       myGenDataSliceHandler
	insertHook        myGenDataSliceHandler

	updateParse      myGenDataSliceHandler
	updateHookBefore myGenDataSliceHandler
	updateHookAfter  myGenDataSliceHandler

	orderParse myGenDataSliceHandler

	joinParse myGenDataSliceHandler
}

func (daoThis *myGenDao) Add(daoField myGenDaoField) {
	daoThis.importDao = append(daoThis.importDao, daoField.importDao...)
	daoThis.insertParseBefore = append(daoThis.insertParseBefore, daoField.insertParseBefore.getData()...)
	daoThis.insertParse = append(daoThis.insertParse, daoField.insertParse.getData()...)
	daoThis.insertHook = append(daoThis.insertHook, daoField.insertHook.getData()...)
	daoThis.updateParse = append(daoThis.updateParse, daoField.updateParse.getData()...)
	daoThis.updateHookBefore = append(daoThis.updateHookBefore, daoField.updateHookBefore.getData()...)
	daoThis.updateHookAfter = append(daoThis.updateHookAfter, daoField.updateHookAfter.getData()...)
	daoThis.fieldParse = append(daoThis.fieldParse, daoField.fieldParse.getData()...)
	daoThis.fieldHook = append(daoThis.fieldHook, daoField.fieldHook.getData()...)
	daoThis.filterParse = append(daoThis.filterParse, daoField.filterParse.getData()...)
	daoThis.orderParse = append(daoThis.orderParse, daoField.orderParse.getData()...)
	daoThis.joinParse = append(daoThis.joinParse, daoField.joinParse.getData()...)
}

func (daoThis *myGenDao) Merge(daoOther myGenDao) {
	daoThis.importDao = append(daoThis.importDao, daoOther.importDao...)
	daoThis.filterParse = append(daoThis.filterParse, daoOther.filterParse...)
	daoThis.fieldParse = append(daoThis.fieldParse, daoOther.fieldParse...)
	daoThis.fieldHook = append(daoThis.fieldHook, daoOther.fieldHook...)
	daoThis.insertParseBefore = append(daoThis.insertParseBefore, daoOther.insertParseBefore...)
	daoThis.insertParse = append(daoThis.insertParse, daoOther.insertParse...)
	daoThis.insertHook = append(daoThis.insertHook, daoOther.insertHook...)
	daoThis.updateParse = append(daoThis.updateParse, daoOther.updateParse...)
	daoThis.updateHookBefore = append(daoThis.updateHookBefore, daoOther.updateHookBefore...)
	daoThis.updateHookAfter = append(daoThis.updateHookAfter, daoOther.updateHookAfter...)
	daoThis.groupParse = append(daoThis.groupParse, daoOther.groupParse...)
	daoThis.orderParse = append(daoThis.orderParse, daoOther.orderParse...)
	daoThis.joinParse = append(daoThis.joinParse, daoOther.joinParse...)
}

func (daoThis *myGenDao) Unique() {
	daoThis.importDao = garray.NewStrArrayFrom(daoThis.importDao).Unique().Slice()
	daoThis.insertParseBefore = garray.NewStrArrayFrom(daoThis.insertParseBefore).Unique().Slice()
	daoThis.insertParse = garray.NewStrArrayFrom(daoThis.insertParse).Unique().Slice()
	daoThis.insertHook = garray.NewStrArrayFrom(daoThis.insertHook).Unique().Slice()
	daoThis.updateParse = garray.NewStrArrayFrom(daoThis.updateParse).Unique().Slice()
	daoThis.updateHookBefore = garray.NewStrArrayFrom(daoThis.updateHookBefore).Unique().Slice()
	daoThis.updateHookAfter = garray.NewStrArrayFrom(daoThis.updateHookAfter).Unique().Slice()
	daoThis.fieldParse = garray.NewStrArrayFrom(daoThis.fieldParse).Unique().Slice()
	daoThis.fieldHook = garray.NewStrArrayFrom(daoThis.fieldHook).Unique().Slice()
	daoThis.filterParse = garray.NewStrArrayFrom(daoThis.filterParse).Unique().Slice()
	daoThis.orderParse = garray.NewStrArrayFrom(daoThis.orderParse).Unique().Slice()
	daoThis.joinParse = garray.NewStrArrayFrom(daoThis.joinParse).Unique().Slice()
}

// dao生成
func genDao(tpl myGenTpl) {
	tpl.gfGenDao(true) //dao文件生成
	saveFile := gfile.SelfDir() + `/internal/dao/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseSnake + `.go`
	tplDao := gfile.GetContents(saveFile)

	dao := getDaoIdAndLabel(tpl)
	dao.Merge(getDaoFieldList(tpl))
	dao.Unique()

	if len(dao.importDao) > 0 {
		importDaoPoint := `"api/internal/dao/` + tpl.ModuleDirCaseKebab + `/internal"`
		tplDao = gstr.Replace(tplDao, importDaoPoint, importDaoPoint+gstr.Join(append([]string{``}, dao.importDao...), `
	`), 1)
	}
	tplDao = gstr.Replace(tplDao, `"github.com/gogf/gf/v2/util/gconv"`, `"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"`, 1)

	if dao.primaryKeyFunction != `` {
		primaryKeyFunctionPoint := `// 解析filter`
		tplDao = gstr.Replace(tplDao, primaryKeyFunctionPoint, dao.primaryKeyFunction+`

`+primaryKeyFunctionPoint, 1)
	}

	// 解析filter
	if len(dao.filterParse) > 0 {
		filterParsePoint := `/* case ` + "`xxxx`" + `:
			tableXxxx := Xxxx.ParseDbTable(m.GetCtx())
			m = m.Where(tableXxxx+` + "`.`" + `+k, v)
			m = m.Handler(daoThis.ParseJoin(tableXxxx, daoModel)) */`
		tplDao = gstr.Replace(tplDao, filterParsePoint, filterParsePoint+gstr.Join(append([]string{``}, dao.filterParse...), `
			`), 1)
	}

	// 解析field
	if len(dao.fieldParse) > 0 {
		fieldParsePoint := `/* case ` + "`xxxx`" + `:
			tableXxxx := Xxxx.ParseDbTable(m.GetCtx())
			m = m.Fields(tableXxxx + ` + "`.`" + ` + v)
			m = m.Handler(daoThis.ParseJoin(tableXxxx, daoModel))
			daoModel.AfterField.Add(v) */`
		tplDao = gstr.Replace(tplDao, fieldParsePoint, fieldParsePoint+gstr.Join(append([]string{``}, dao.fieldParse...), `
			`), 1)
	}
	if len(dao.fieldHook) > 0 {
		fieldHookPoint := `default:
						record[v] = gvar.New(nil)`
		tplDao = gstr.Replace(tplDao, fieldHookPoint, gstr.Join(append(dao.fieldHook, ``), `
					`)+fieldHookPoint, 1)
	}

	// 解析insert
	if len(dao.insertParseBefore) > 0 {
		insertParseBeforePoint := `insertData := map[string]interface{}{}`
		tplDao = gstr.Replace(tplDao, insertParseBeforePoint, gstr.Join(append(dao.insertParseBefore, ``), `
		`)+insertParseBeforePoint, 1)
	}
	if len(dao.insertParse) > 0 {
		insertParsePoint := `default:
				if daoThis.ColumnArr().Contains(k) {
					insertData[k] = v
				}`
		tplDao = gstr.Replace(tplDao, insertParsePoint, gstr.Join(append(dao.insertParse, ``), `
			`)+insertParsePoint, 1)
	}
	if len(dao.insertHook) > 0 {
		insertHookPoint := `// id, _ := result.LastInsertId()

			/* for k, v := range daoModel.AfterInsert {
				switch k {
				case ` + "`xxxx`" + `:
					daoModel.CloneNew().Filter(daoThis.PrimaryKey(), id).HookUpdate(g.Map{k: v}).Update()
				}
			} */`
		tplDao = gstr.Replace(tplDao, insertHookPoint, `id, _ := result.LastInsertId()`+gstr.Join(append([]string{``}, dao.insertHook...), `

			`), 1)
	}

	// 解析update
	if len(dao.updateParse) > 0 {
		updateParsePoint := `default:
				if daoThis.ColumnArr().Contains(k) {
					updateData[daoModel.DbTable+` + "`.`" + `+k] = gvar.New(v) //因下面bug处理方式，json类型字段传参必须是gvar变量，否则不会自动生成json格式
				}`
		tplDao = gstr.Replace(tplDao, updateParsePoint, gstr.Join(append(dao.updateParse, ``), `
			`)+updateParsePoint, 1)
	}
	if len(dao.updateHookBefore) > 0 || len(dao.updateHookAfter) > 0 {
		updateHookPoint := `/* row, _ := result.RowsAffected()
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
		if len(dao.updateHookBefore) > 0 {
			tplDao = gstr.Replace(tplDao, updateHookPoint, gstr.Join(append(dao.updateHookBefore, ``), `

			`)+updateHookPoint, 1)
		}
		if len(dao.updateHookAfter) > 0 {
			tplDao = gstr.Replace(tplDao, updateHookPoint, `row, _ := result.RowsAffected()
			if row == 0 {
				return
			}`+gstr.Join(append([]string{``}, dao.updateHookAfter...), `

			`), 1)
		}
	}

	// 解析order
	if len(dao.groupParse) > 0 {
		groupParsePoint := `default:
				if daoThis.ColumnArr().Contains(v) {
					m = m.Group(daoModel.DbTable + ` + "`.`" + ` + v)
				} else {
					m = m.Group(v)
				}`
		tplDao = gstr.Replace(tplDao, groupParsePoint, gstr.Join(append(dao.groupParse, ``), `
			`)+groupParsePoint, 1)
	}

	// 解析order
	if len(dao.orderParse) > 0 {
		orderParsePoint := `default:
				if daoThis.ColumnArr().Contains(k) {
					m = m.Order(daoModel.DbTable + ` + "`.`" + ` + v)
				} else {
					m = m.Order(v)
				}`
		tplDao = gstr.Replace(tplDao, orderParsePoint, gstr.Join(append(dao.orderParse, ``), `
			`)+orderParsePoint, 1)
	}

	// 解析join
	if len(dao.joinParse) > 0 {
		joinParsePoint := `/* case Xxxx.ParseDbTable(m.GetCtx()):
		m = m.LeftJoin(joinTable, joinTable+` + "`.`" + `+Xxxx.Columns().XxxxId+` + "` = `" + `+daoModel.DbTable+` + "`.`" + `+daoThis.PrimaryKey())
		// m = m.LeftJoin(Xxxx.ParseDbTable(m.GetCtx())+` + "` AS `" + `+joinTable, joinTable+` + "`.`" + `+Xxxx.Columns().XxxxId+` + "` = `" + `+daoModel.DbTable+` + "`.`" + `+daoThis.PrimaryKey()) */`
		tplDao = gstr.Replace(tplDao, joinParsePoint, joinParsePoint+gstr.Join(append([]string{``}, dao.joinParse...), `
		`), 1)
	}

	gfile.PutContents(saveFile, tplDao)
	utils.GoFileFmt(saveFile)
}

func getDaoIdAndLabel(tpl myGenTpl) (dao myGenDao) {
	if tpl.Handle.Id.List[0].FieldRaw != tpl.FieldList[0].FieldRaw {
		dao.primaryKeyFunction = `// 主键ID
func (daoThis *` + gstr.CaseCamelLower(tpl.TableCaseCamel) + `Dao) PrimaryKey() string {
	return ` + "`" + tpl.Handle.Id.List[0].FieldRaw + "`" + `
}`
	}
	if len(tpl.Handle.Id.List) == 1 {
		dao.filterParse = append(dao.filterParse, `case `+"`id`, `idArr`"+`:
				m = m.Where(daoModel.DbTable+`+"`.`"+`+daoThis.PrimaryKey(), v)`)
		dao.filterParse = append(dao.filterParse, `case `+"`excId`, `excIdArr`"+`:
				if gvar.New(v).IsSlice() {
					m = m.WhereNotIn(daoModel.DbTable+`+"`.`"+`+daoThis.PrimaryKey(), v)
				} else {
					m = m.WhereNot(daoModel.DbTable+`+"`.`"+`+daoThis.PrimaryKey(), v)
				}`)
		dao.fieldParse = append(dao.fieldParse, `case `+"`id`"+`:
				m = m.Fields(daoModel.DbTable + `+"`.`"+` + daoThis.PrimaryKey() + `+"` AS `"+` + v)`)
		if !tpl.Handle.Id.List[0].IsAutoInc {
			dao.insertParse = append(dao.insertParse, `case `+"`id`"+`:
					insertData[daoThis.PrimaryKey()] = v`)
			dao.updateParse = append(dao.updateParse, `case `+"`id`"+`:
					updateData[daoModel.DbTable+`+"`.`"+`+daoThis.PrimaryKey()] = v`)
		}
		dao.groupParse = append(dao.groupParse, `case `+"`id`"+`:
				m = m.Group(daoModel.DbTable + `+"`.`"+` + daoThis.PrimaryKey())`)
		dao.orderParse = append(dao.orderParse, `case `+"`id`"+`:
				m = m.Order(daoModel.DbTable + `+"`.`"+` + gstr.Replace(v, k, daoThis.PrimaryKey(), 1))`)
	} else {
		concatStr := `|`
		filterParseStrArr := []string{}
		fieldParseStrArr := []string{}
		groupParseStrArr := []string{}
		orderParseStrArr := []string{}
		for _, v := range tpl.Handle.Id.List {
			filterParseStrArr = append(filterParseStrArr, ` + daoModel.DbTable + `+"`.`"+` + daoThis.Columns().`+v.FieldCaseCamel+` + `)
			fieldParseStrArr = append(fieldParseStrArr, "IFNULL(` + daoModel.DbTable + `.` + daoThis.Columns()."+v.FieldCaseCamel+" + `, '')")
			groupParseStrArr = append(groupParseStrArr, `m = m.Group(daoModel.DbTable + `+"`.`"+` + daoThis.Columns().`+v.FieldCaseCamel+`)`)
			orderParseStrArr = append(orderParseStrArr, `m = m.Order(daoModel.DbTable + `+"`.`"+` + daoThis.Columns().`+v.FieldCaseCamel+` + suffix)`)
		}
		dao.filterParse = append(dao.filterParse, `case `+"`id`, `idArr`"+`:
				idArr := []string{gconv.String(v)}
				if gvar.New(v).IsSlice() {
					idArr = gconv.SliceStr(v)
				}
				inStrArr := []string{}
				for _, id := range idArr {
					gstr.Replace(gconv.String(id), `+"`"+concatStr+"`, `', '`)"+`
					inStrArr = append(inStrArr, `+"`('`+gstr.Replace(gconv.String(id), `"+concatStr+"`, `', '`)+`')`)"+`
				}
				m = m.Where(`+"`(`"+gstr.Join(filterParseStrArr, "`, `")+"`) IN (` + gstr.Join(inStrArr, `, `) + `)`)")
		dao.filterParse = append(dao.filterParse, `case `+"`excId`, `excIdArr`"+`:
				idArr := []string{gconv.String(v)}
				if gvar.New(v).IsSlice() {
					idArr = gconv.SliceStr(v)
				}
				inStrArr := []string{}
				for _, id := range idArr {
					gstr.Replace(gconv.String(id), `+"`"+concatStr+"`, `', '`)"+`
					inStrArr = append(inStrArr, `+"`('`+gstr.Replace(gconv.String(id), `"+concatStr+"`, `', '`)+`')`)"+`
				}
				m = m.Where(`+"`(`"+gstr.Join(filterParseStrArr, "`, `")+"`) NOT IN (` + gstr.Join(inStrArr, `, `) + `)`)")
		dao.fieldParse = append(dao.fieldParse, `case `+"`id`"+`:
				m = m.Fields(`+"`"+`CONCAT_WS('`+concatStr+`', `+gstr.Join(fieldParseStrArr, `, `)+")` + ` AS ` + v)")
		dao.groupParse = append(dao.groupParse, `case `+"`id`"+`:`+gstr.Join(append([]string{``}, groupParseStrArr...), `
				`))
		dao.orderParse = append(dao.orderParse, `case `+"`id`"+`:
				suffix := gstr.TrimLeftStr(kArr[0], k, 1)
				`+gstr.Join(append(orderParseStrArr, ``), `
				`)+`remain := gstr.TrimLeftStr(gstr.TrimLeftStr(v, k+suffix, 1), `+"`,`"+`, 1)
				if remain != `+"``"+` {
					m = m.Order(remain)
				}`)
	}

	labelListLen := len(tpl.Handle.LabelList)
	if labelListLen > 0 {
		fieldParseStr := `case ` + "`label`" + `:
				m = m.Fields(daoModel.DbTable + ` + "`.`" + ` + daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.LabelList[0]) + ` + ` + "` AS `" + ` + v)`
		filterParseStr := `case ` + "`label`" + `:
				m = m.WhereLike(daoModel.DbTable+` + "`.`" + `+daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.LabelList[0]) + `, ` + "`%`" + `+gconv.String(v)+` + "`%`" + `)`
		if labelListLen > 1 {
			fieldParseStrTmp := "` + daoModel.DbTable + `.` + daoThis.Columns()." + gstr.CaseCamel(tpl.Handle.LabelList[labelListLen-1]) + " + `"
			parseFilterStr := "WhereOrLike(daoModel.DbTable+`.`+daoThis.Columns()." + gstr.CaseCamel(tpl.Handle.LabelList[labelListLen-1]) + ", `%`+gconv.String(v)+`%`)"
			for i := labelListLen - 2; i >= 0; i-- {
				fieldParseStrTmp = "IF(IFNULL(` + daoModel.DbTable + `.` + daoThis.Columns()." + gstr.CaseCamel(tpl.Handle.LabelList[i]) + " + `, '') != '', ` + daoModel.DbTable + `.` + daoThis.Columns()." + gstr.CaseCamel(tpl.Handle.LabelList[i]) + " + `, " + fieldParseStrTmp + ")"
				if i == 0 {
					parseFilterStr = "WhereLike(daoModel.DbTable+`.`+daoThis.Columns()." + gstr.CaseCamel(tpl.Handle.LabelList[i]) + ", `%`+gconv.String(v)+`%`)." + parseFilterStr
				} else {
					parseFilterStr = "WhereOrLike(daoModel.DbTable+`.`+daoThis.Columns()." + gstr.CaseCamel(tpl.Handle.LabelList[i]) + ", `%`+gconv.String(v)+`%`)." + parseFilterStr
				}
			}
			fieldParseStr = `case ` + "`label`" + `:
				m = m.Fields(` + "`" + fieldParseStrTmp + " AS ` + v)"
			filterParseStr = `case ` + "`label`" + `:
				m = m.Where(m.Builder().` + parseFilterStr + `)`
		}
		dao.fieldParse = append(dao.fieldParse, fieldParseStr)
		dao.filterParse = append(dao.filterParse, filterParseStr)
	}
	return
}

func getDaoFieldList(tpl myGenTpl) (dao myGenDao) {
	for _, v := range tpl.FieldList {
		daoField := myGenDaoField{}
		/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 开始--------*/
		switch v.FieldType {
		case TypeInt: // `int等类型`
		case TypeIntU: // `int等类型（unsigned）`
		case TypeFloat: // `float等类型`
		case TypeFloatU: // `float等类型（unsigned）`
		case TypeVarchar, TypeChar: // `varchar类型`	// `char类型`
			if v.IsUnique && v.IsNull {
				daoField.insertParse.Method = ReturnType
				daoField.insertParse.DataType = append(daoField.insertParse.DataType, `case daoThis.Columns().`+v.FieldCaseCamel+`:
				insertData[k] = v
				if gconv.String(v) == `+"``"+` {
					insertData[k] = nil
				}`)

				daoField.updateParse.Method = ReturnType
				daoField.updateParse.DataType = append(daoField.updateParse.DataType, `case daoThis.Columns().`+v.FieldCaseCamel+`:
				updateData[daoModel.DbTable+`+"`.`"+`+k] = v
				if gconv.String(v) == `+"``"+` {
					updateData[daoModel.DbTable+`+"`.`"+`+k] = nil
				}`)
			}
		case TypeText: // `text类型`
		case TypeJson: // `json类型`
			if v.IsNull {
				daoField.insertParse.Method = ReturnType
				daoField.insertParse.DataType = append(daoField.insertParse.DataType, `case daoThis.Columns().`+v.FieldCaseCamel+`:
				insertData[k] = v
				if gconv.String(v) == `+"``"+` {
					insertData[k] = nil
				}`)

				daoField.updateParse.Method = ReturnType
				daoField.updateParse.DataType = append(daoField.updateParse.DataType, `case daoThis.Columns().`+v.FieldCaseCamel+`:
				updateData[daoModel.DbTable+`+"`.`"+`+k] = gvar.New(v)
				if gconv.String(v) == `+"``"+` {
					updateData[daoModel.DbTable+`+"`.`"+`+k] = nil
				}`)
			}
		case TypeTimestamp: // `timestamp类型`
		case TypeDatetime: // `datetime类型`
		case TypeDate: // `date类型`
			daoField.orderParse.Method = ReturnType
			daoField.orderParse.DataType = append(daoField.orderParse.DataType, `case daoThis.Columns().`+v.FieldCaseCamel+`:
				m = m.Order(daoModel.DbTable + `+"`.`"+` + v)
				m = m.OrderDesc(daoModel.DbTable + `+"`.`"+` + daoThis.PrimaryKey())`) //追加主键倒序。mysql排序字段有重复值时，分页会导致同一条数据可能在不同页都出现
		}
		/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 结束--------*/

		/*--------根据字段主键类型处理 开始--------*/
		switch v.FieldTypePrimary {
		case TypePrimary: // 独立主键
		case TypePrimaryAutoInc: // 独立主键（自增）
			continue
		case TypePrimaryMany: // 联合主键
		case TypePrimaryManyAutoInc: // 联合主键（自增）
			continue
		}
		/*--------根据字段主键类型处理 结束--------*/

		/*--------根据字段命名类型处理 开始--------*/
		switch v.FieldTypeName {
		case TypeNameDeleted: // 软删除字段
		case TypeNameUpdated: // 更新时间字段
		case TypeNameCreated: // 创建时间字段
			daoField.filterParse.Method = ReturnTypeName
			daoField.filterParse.DataTypeName = append(daoField.filterParse.DataTypeName, `case `+"`timeRangeStart`"+`:
				m = m.WhereGTE(daoModel.DbTable+`+"`.`"+`+daoThis.Columns().`+v.FieldCaseCamel+`, v)
			case `+"`timeRangeEnd`"+`:
				m = m.WhereLTE(daoModel.DbTable+`+"`.`"+`+daoThis.Columns().`+v.FieldCaseCamel+`, v)`)
		case TypeNamePid: // pid；	类型：int等类型；
			daoField.fieldParse.Method = ReturnTypeName
			if len(tpl.Handle.LabelList) > 0 {
				daoField.fieldParse.DataTypeName = append(daoField.fieldParse.DataTypeName, `case `+"`p"+gstr.CaseCamel(tpl.Handle.LabelList[0])+"`"+`:
					tableP := `+"`p_`"+` + daoModel.DbTable
					m = m.Fields(tableP + `+"`.`"+` + daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.LabelList[0])+` + `+"` AS `"+` + v)
					m = m.Handler(daoThis.ParseJoin(tableP, daoModel))`)
			}
			daoField.fieldParse.DataTypeName = append(daoField.fieldParse.DataTypeName, `case `+"`tree`"+`:
				m = m.Fields(daoModel.DbTable + `+"`.`"+` + daoThis.PrimaryKey())
				m = m.Fields(daoModel.DbTable + `+"`.`"+` + daoThis.Columns().`+v.FieldCaseCamel+`)
				m = m.Handler(daoThis.ParseOrder([]string{`+"`tree`"+`}, daoModel))`)

			orderParseStr := `case ` + "`tree`" + `:
				m = m.OrderAsc(daoModel.DbTable + ` + "`.`" + ` + daoThis.Columns().` + v.FieldCaseCamel + `)`
			if tpl.Handle.Pid.Sort != `` {
				orderParseStr += `
				m = m.OrderAsc(daoModel.DbTable + ` + "`.`" + ` + daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Sort) + `)`
			}
			orderParseStr += `
				m = m.OrderAsc(daoModel.DbTable + ` + "`.`" + ` + daoThis.PrimaryKey())`
			daoField.orderParse.Method = ReturnTypeName
			daoField.orderParse.DataTypeName = append(daoField.orderParse.DataTypeName, orderParseStr)

			daoField.joinParse.Method = ReturnTypeName
			daoField.joinParse.DataTypeName = append(daoField.joinParse.DataTypeName, `case `+"`p_`"+` + daoModel.DbTable:
			m = m.LeftJoin(daoModel.DbTable+`+"` AS `"+`+joinTable, joinTable+`+"`.`"+`+daoThis.PrimaryKey()+`+"` = `"+`+daoModel.DbTable+`+"`.`"+`+daoThis.Columns().`+v.FieldCaseCamel+`)`)

			if tpl.Handle.Pid.IsCoexist {
				daoField.insertParseBefore.Method = ReturnTypeName
				daoField.insertParseBefore.DataTypeName = append(daoField.insertParseBefore.DataTypeName, `if _, ok := insert[daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Pid)+`]; !ok {
			insert[daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Pid)+`] = 0
		}`)

				daoField.insertParse.Method = ReturnTypeName
				daoField.insertParse.DataTypeName = append(daoField.insertParse.DataTypeName, `case daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Pid)+`:
				insertData[k] = v
				if gconv.Uint(v) > 0 {
					pInfo, _ := daoThis.CtxDaoModel(m.GetCtx()).Filter(daoThis.PrimaryKey(), v).One()
					daoModel.AfterInsert[`+"`pIdPath`"+`] = pInfo[daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.IdPath)+`].String()
					daoModel.AfterInsert[`+"`pLevel`"+`] = pInfo[daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Level)+`].Uint()
				} else {
					daoModel.AfterInsert[`+"`pIdPath`"+`] = `+"`0`"+`
					daoModel.AfterInsert[`+"`pLevel`"+`] = 0
				}`)

				daoField.insertHook.Method = ReturnTypeName
				daoField.insertHook.DataTypeName = append(daoField.insertHook.DataTypeName, `updateSelfData := map[string]interface{}{}
			for k, v := range daoModel.AfterInsert {
				switch k {
				case `+"`pIdPath`"+`:
					updateSelfData[daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.IdPath)+`] = gconv.String(v) + `+"`-`"+` + gconv.String(id)
				case `+"`pLevel`"+`:
					updateSelfData[daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Level)+`] = gconv.Uint(v) + 1
				}
			}
			if len(updateSelfData) > 0 {
				daoModel.CloneNew().Filter(daoThis.PrimaryKey(), id).HookUpdate(updateSelfData).Update()
			}`)

				daoField.updateParse.Method = ReturnTypeName
				daoField.updateParse.DataTypeName = append(daoField.updateParse.DataTypeName, `case daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Pid)+`:
				updateData[daoModel.DbTable+`+"`.`"+`+k] = v
				pIdPath := `+"`0`"+`
				var pLevel uint = 0
				if gconv.Uint(v) > 0 {
					pInfo, _ := daoThis.CtxDaoModel(m.GetCtx()).Filter(daoThis.PrimaryKey(), v).One()
					pIdPath = pInfo[daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.IdPath)+`].String()
					pLevel = pInfo[daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Level)+`].Uint()
				}
				updateData[daoModel.DbTable+`+"`.`"+`+daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.IdPath)+`] = gdb.Raw(`+"`CONCAT('`"+` + pIdPath + `+"`-', `"+` + daoThis.PrimaryKey() + `+"`)`"+`)
				updateData[daoModel.DbTable+`+"`.`"+`+daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Level)+`] = pLevel + 1
				//更新所有子孙级的idPath和level
				updateChildIdPathAndLevelList := []map[string]interface{}{}
				oldList, _ := daoThis.CtxDaoModel(m.GetCtx()).Filter(daoThis.PrimaryKey(), daoModel.IdArr).All()
				for _, oldInfo := range oldList {
					if gconv.Uint(v) != oldInfo[daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Pid)+`].Uint() {
						updateChildIdPathAndLevelList = append(updateChildIdPathAndLevelList, map[string]interface{}{
							`+"`pIdPathOfOld`"+`: oldInfo[daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.IdPath)+`],
							`+"`pIdPathOfNew`"+`: pIdPath + `+"`-`"+` + oldInfo[daoThis.PrimaryKey()].String(),
							`+"`pLevelOfOld`"+`:  oldInfo[daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Level)+`],
							`+"`pLevelOfNew`"+`:  pLevel + 1,
						})
					}
				}
				if len(updateChildIdPathAndLevelList) > 0 {
					daoModel.AfterUpdate[`+"`updateChildIdPathAndLevelList`"+`] = updateChildIdPathAndLevelList
				}
			case `+"`childIdPath`"+`: //更新所有子孙级的idPath。参数：map[string]interface{}{`+"`pIdPathOfOld`"+`: `+"`父级IdPath（旧）`"+`, `+"`pIdPathOfNew`"+`: `+"`父级IdPath（新）`"+`}
				val := gconv.Map(v)
				pIdPathOfOld := gconv.String(val[`+"`pIdPathOfOld`"+`])
				pIdPathOfNew := gconv.String(val[`+"`pIdPathOfNew`"+`])
				updateData[daoModel.DbTable+`+"`.`"+`+daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.IdPath)+`] = gdb.Raw(`+"`REPLACE(`"+` + daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.IdPath)+` + `+"`, '`"+` + pIdPathOfOld + `+"`', '`"+` + pIdPathOfNew + `+"`')`"+`)
			case `+"`childLevel`"+`: //更新所有子孙级的level。参数：map[string]interface{}{`+"`pLevelOfOld`"+`: `+"`父级Level（旧）`"+`, `+"`pLevelOfNew`"+`: `+"`父级Level（新）`"+`}
				val := gconv.Map(v)
				pLevelOfOld := gconv.Uint(val[`+"`pLevelOfOld`"+`])
				pLevelOfNew := gconv.Uint(val[`+"`pLevelOfNew`"+`])
				updateData[daoModel.DbTable+`+"`.`"+`+daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Level)+`] = gdb.Raw(daoModel.DbTable + `+"`.`"+` + daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Level)+` + `+"` + `"+` + gconv.String(pLevelOfNew-pLevelOfOld))
				if pLevelOfNew < pLevelOfOld {
					updateData[daoModel.DbTable+`+"`.`"+`+daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Level)+`] = gdb.Raw(daoModel.DbTable + `+"`.`"+` + daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Level)+` + `+"` - `"+` + gconv.String(pLevelOfOld-pLevelOfNew))
				}`)

				daoField.updateHookAfter.Method = ReturnTypeName
				daoField.updateHookAfter.DataTypeName = append(daoField.updateHookAfter.DataTypeName, `for k, v := range daoModel.AfterUpdate {
				switch k {
				case `+"`updateChildIdPathAndLevelList`"+`: //修改pid时，更新所有子孙级的idPath和level。参数：[]map[string]interface{}{`+"`pIdPathOfOld`"+`: `+"`父级IdPath（旧）`"+`, `+"`pIdPathOfNew`"+`: `+"`父级IdPath（新）`"+`, `+"`pLevelOfOld`"+`: `+"`父级Level（旧）`"+`, `+"`pLevelOfNew`"+`: `+"`父级Level（新）`"+`}
					val := v.([]map[string]interface{})
					for _, v1 := range val {
						daoModel.CloneNew().Filter(`+"`pIdPathOfOld`"+`, v1[`+"`pIdPathOfOld`"+`]).HookUpdate(g.Map{
							`+"`childIdPath`"+`: g.Map{
								`+"`pIdPathOfOld`"+`: v1[`+"`pIdPathOfOld`"+`],
								`+"`pIdPathOfNew`"+`: v1[`+"`pIdPathOfNew`"+`],
							},
							`+"`childLevel`"+`: g.Map{
								`+"`pLevelOfOld`"+`: v1[`+"`pLevelOfOld`"+`],
								`+"`pLevelOfNew`"+`: v1[`+"`pLevelOfNew`"+`],
							},
						}).Update()
					}
				}
			}`)

				daoField.filterParse.Method = ReturnTypeName
				daoField.filterParse.DataTypeName = append(daoField.filterParse.DataTypeName, `case `+"`pIdPathOfOld`"+`: //父级IdPath（旧）
				m = m.WhereLike(daoModel.DbTable+`+"`.`"+`+daoThis.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.IdPath)+`, gconv.String(v)+`+"`-%`"+`)`)
			}
		case TypeNameLevel: // level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；
			daoField.orderParse.Method = ReturnTypeName
			daoField.orderParse.DataTypeName = append(daoField.orderParse.DataTypeName, `case daoThis.Columns().`+v.FieldCaseCamel+`:
				m = m.Order(daoModel.DbTable + `+"`.`"+` + v)
				m = m.OrderDesc(daoModel.DbTable + `+"`.`"+` + daoThis.PrimaryKey())`) //追加主键倒序。mysql排序字段有重复值时，分页会导致同一条数据可能在不同页都出现
		case TypeNameIdPath: // idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；
			continue
		case TypeNamePasswordSuffix: // password,passwd后缀；		类型：char(32)；
			insertParseStr := `case daoThis.Columns().` + v.FieldCaseCamel + `:
				password := gconv.String(v)
				if len(password) != 32 {
					password = gmd5.MustEncrypt(password)
				}`
			updateParseStr := `case daoThis.Columns().` + v.FieldCaseCamel + `:
				password := gconv.String(v)
				if len(password) != 32 {
					password = gmd5.MustEncrypt(password)
				}`
			passwordMapKey := tpl.getHandlePasswordMapKey(v.FieldRaw)
			if tpl.Handle.PasswordMap[passwordMapKey].IsCoexist {
				insertParseStr += `
				salt := grand.S(` + tpl.Handle.PasswordMap[passwordMapKey].SaltLength + `)
				insertData[daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.PasswordMap[passwordMapKey].SaltField) + `] = salt
				password = gmd5.MustEncrypt(password + salt)`
				updateParseStr += `
				salt := grand.S(` + tpl.Handle.PasswordMap[passwordMapKey].SaltLength + `)
				updateData[daoModel.DbTable+` + "`.`" + `+daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.PasswordMap[passwordMapKey].SaltField) + `] = salt
				password = gmd5.MustEncrypt(password + salt)`
			}
			insertParseStr += `
				insertData[k] = password`
			updateParseStr += `
				updateData[daoModel.DbTable+` + "`.`" + `+k] = password`

			daoField.insertParse.Method = ReturnTypeName
			daoField.insertParse.DataTypeName = append(daoField.insertParse.DataTypeName, insertParseStr)
			daoField.updateParse.Method = ReturnTypeName
			daoField.updateParse.DataTypeName = append(daoField.updateParse.DataTypeName, updateParseStr)
		case TypeNameSaltSuffix: // salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；
			continue
		case TypeNameNameSuffix: // name,title后缀；	类型：varchar；
			daoField.filterParse.Method = ReturnTypeName
			daoField.filterParse.DataTypeName = append(daoField.filterParse.DataTypeName, `case daoThis.Columns().`+v.FieldCaseCamel+`:
				m = m.WhereLike(daoModel.DbTable+`+"`.`"+`+k, `+"`%`"+`+gconv.String(v)+`+"`%`"+`)`)
		case TypeNameCodeSuffix: // code后缀；	类型：varchar；
		case TypeNameAccountSuffix: // account后缀；	类型：varchar；
		case TypeNamePhoneSuffix: // phone,mobile后缀；	类型：varchar；
		case TypeNameEmailSuffix: // email后缀；	类型：varchar；
		case TypeNameUrlSuffix: // url,link后缀；	类型：varchar；
		case TypeNameIpSuffix: // IP后缀；	类型：varchar；
		case TypeNameIdSuffix: // id后缀；	类型：int等类型；
			if tpl.Handle.RelIdMap[v.FieldRaw].tpl.Table != `` {
				relIdObj := tpl.Handle.RelIdMap[v.FieldRaw]
				daoPath := relIdObj.tpl.TableCaseCamel
				if relIdObj.tpl.ModuleDirCaseKebab != tpl.ModuleDirCaseKebab {
					daoPath = `dao` + relIdObj.tpl.ModuleDirCaseCamel + `.` + relIdObj.tpl.TableCaseCamel
					daoField.importDao = append(daoField.importDao, `dao`+relIdObj.tpl.ModuleDirCaseCamel+` "api/internal/dao/`+relIdObj.tpl.ModuleDirCaseKebab+`"`)
				}

				if !tpl.Handle.RelIdMap[v.FieldRaw].IsRedundName {
					fieldParseStr := `case ` + daoPath + `.Columns().` + gstr.CaseCamel(relIdObj.tpl.Handle.LabelList[0]) + `:` + `
				table` + relIdObj.tpl.TableCaseCamel + ` := ` + daoPath + `.ParseDbTable(m.GetCtx())
				m = m.Fields(table` + relIdObj.tpl.TableCaseCamel + ` + ` + "`.`" + ` + v)
				m = m.Handler(daoThis.ParseJoin(table` + relIdObj.tpl.TableCaseCamel + `, daoModel))`
					if relIdObj.Suffix != `` {
						fieldParseStr = `case ` + daoPath + `.Columns().` + gstr.CaseCamel(relIdObj.tpl.Handle.LabelList[0]) + " + `" + relIdObj.Suffix + "`:" + `
				table` + relIdObj.tpl.TableCaseCamel + gstr.CaseCamel(relIdObj.Suffix) + ` := ` + daoPath + `.ParseDbTable(m.GetCtx()) + ` + "`" + gstr.CaseSnake(relIdObj.Suffix) + "`" + `
				m = m.Fields(table` + relIdObj.tpl.TableCaseCamel + gstr.CaseCamel(relIdObj.Suffix) + ` + ` + "`.`" + ` + ` + daoPath + `.Columns().` + gstr.CaseCamel(relIdObj.tpl.Handle.LabelList[0]) + ` + ` + "` AS `" + ` + v)
				m = m.Handler(daoThis.ParseJoin(table` + relIdObj.tpl.TableCaseCamel + gstr.CaseCamel(relIdObj.Suffix) + `, daoModel))`
					}
					daoField.fieldParse.Method = ReturnTypeName
					daoField.fieldParse.DataTypeName = append(daoField.fieldParse.DataTypeName, fieldParseStr)
				}

				joinParseStr := `case ` + daoPath + `.ParseDbTable(m.GetCtx()):
			m = m.LeftJoin(joinTable, joinTable+` + "`.`" + `+` + daoPath + `.PrimaryKey()+` + "` = `" + `+daoModel.DbTable+` + "`.`" + `+daoThis.Columns().` + v.FieldCaseCamel + `)`
				if relIdObj.Suffix != `` {
					joinParseStr = `case ` + daoPath + `.ParseDbTable(m.GetCtx()) + ` + "`" + gstr.CaseSnake(relIdObj.Suffix) + "`" + `:
			m = m.LeftJoin(` + daoPath + `.ParseDbTable(m.GetCtx())+` + "` AS `" + `+joinTable, joinTable+` + "`.`" + `+` + daoPath + `.PrimaryKey()+` + "` = `" + `+daoModel.DbTable+` + "`.`" + `+daoThis.Columns().` + v.FieldCaseCamel + `)`
				}
				daoField.joinParse.Method = ReturnTypeName
				daoField.joinParse.DataTypeName = append(daoField.joinParse.DataTypeName, joinParseStr)
			}
		case TypeNameSortSuffix, TypeNameSort: // sort,weight等后缀；	类型：int等类型； // sort，且pid,level,idPath|id_path,sort同时存在时（才）有效；	类型：int等类型；
			daoField.orderParse.Method = ReturnTypeName
			daoField.orderParse.DataTypeName = append(daoField.orderParse.DataTypeName, `case daoThis.Columns().`+v.FieldCaseCamel+`:
				m = m.Order(daoModel.DbTable + `+"`.`"+` + v)
				m = m.OrderDesc(daoModel.DbTable + `+"`.`"+` + daoThis.PrimaryKey())`) //追加主键倒序。mysql排序字段有重复值时，分页会导致同一条数据可能在不同页都出现
		case TypeNameStatusSuffix: // status,type,method,pos,position,gender等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
		case TypeNameIsPrefix: // is_前缀；		类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（停用：0否 1是）
		case TypeNameStartPrefix: // start_前缀；	类型：timestamp或datetime或date；
			filterParseStr := `case daoThis.Columns().` + v.FieldCaseCamel + `:
				m = m.WhereLTE(daoModel.DbTable+` + "`.`" + `+k, v)`
			if !v.IsNull && gconv.String(v.Default) == `` {
				filterParseStr = `case daoThis.Columns().` + v.FieldCaseCamel + `:
				m = m.Where(m.Builder().WhereLTE(daoModel.DbTable+` + "`.`" + `+k, v).WhereOrNull(daoModel.DbTable + ` + "`.`" + ` + k))`
			}
			daoField.filterParse.Method = ReturnTypeName
			daoField.filterParse.DataTypeName = append(daoField.filterParse.DataTypeName, filterParseStr)
		case TypeNameEndPrefix: // end_前缀；	类型：timestamp或datetime或date；
			filterParseStr := `case daoThis.Columns().` + v.FieldCaseCamel + `:
				m = m.WhereGTE(daoModel.DbTable+` + "`.`" + `+k, v)`
			if !v.IsNull && gconv.String(v.Default) == `` {
				filterParseStr = `case daoThis.Columns().` + v.FieldCaseCamel + `:
				m = m.Where(m.Builder().WhereGTE(daoModel.DbTable+` + "`.`" + `+k, v).WhereOrNull(daoModel.DbTable + ` + "`.`" + ` + k))`
			}
			daoField.filterParse.Method = ReturnTypeName
			daoField.filterParse.DataTypeName = append(daoField.filterParse.DataTypeName, filterParseStr)
		case TypeNameRemarkSuffix: // remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器
		case TypeNameImageSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text
		case TypeNameVideoSuffix: // video,video_list,videoList,video_arr,videoArr等后缀；		类型：单视频varchar，多视频json或text
		case TypeNameArrSuffix: // list,arr等后缀；	类型：json或text；
		}
		/*--------根据字段命名类型处理 结束--------*/

		dao.Add(daoField)
	}
	return
}
