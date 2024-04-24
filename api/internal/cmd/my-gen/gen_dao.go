package my_gen

import (
	"api/internal/cmd/my-gen/internal"
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
	insertHookBefore  []string
	insertHook        []string
	insertHookAfter   []string

	updateParse      []string
	updateHookBefore []string
	updateHookAfter  []string

	deleteHook []string

	groupParse []string

	orderParse []string

	joinParse []string
}

type myGenDaoField struct {
	importDao []string

	filterParse internal.MyGenDataSliceHandler

	fieldParse internal.MyGenDataSliceHandler
	fieldHook  internal.MyGenDataSliceHandler

	insertParseBefore internal.MyGenDataSliceHandler
	insertParse       internal.MyGenDataSliceHandler
	insertHookBefore  internal.MyGenDataSliceHandler
	insertHook        internal.MyGenDataSliceHandler
	insertHookAfter   internal.MyGenDataSliceHandler

	updateParse      internal.MyGenDataSliceHandler
	updateHookBefore internal.MyGenDataSliceHandler
	updateHookAfter  internal.MyGenDataSliceHandler

	orderParse internal.MyGenDataSliceHandler

	joinParse internal.MyGenDataSliceHandler
}

func (daoThis *myGenDao) Add(daoField myGenDaoField) {
	daoThis.importDao = append(daoThis.importDao, daoField.importDao...)
	daoThis.filterParse = append(daoThis.filterParse, daoField.filterParse.GetData()...)
	daoThis.fieldParse = append(daoThis.fieldParse, daoField.fieldParse.GetData()...)
	daoThis.fieldHook = append(daoThis.fieldHook, daoField.fieldHook.GetData()...)
	daoThis.insertParseBefore = append(daoThis.insertParseBefore, daoField.insertParseBefore.GetData()...)
	daoThis.insertParse = append(daoThis.insertParse, daoField.insertParse.GetData()...)
	daoThis.insertHookBefore = append(daoThis.insertHookBefore, daoField.insertHookBefore.GetData()...)
	daoThis.insertHook = append(daoThis.insertHook, daoField.insertHook.GetData()...)
	daoThis.insertHookAfter = append(daoThis.insertHookAfter, daoField.insertHookAfter.GetData()...)
	daoThis.updateParse = append(daoThis.updateParse, daoField.updateParse.GetData()...)
	daoThis.updateHookBefore = append(daoThis.updateHookBefore, daoField.updateHookBefore.GetData()...)
	daoThis.updateHookAfter = append(daoThis.updateHookAfter, daoField.updateHookAfter.GetData()...)
	// daoThis.groupParse = append(daoThis.groupParse, daoField.groupParse.GetData()...)
	daoThis.orderParse = append(daoThis.orderParse, daoField.orderParse.GetData()...)
	daoThis.joinParse = append(daoThis.joinParse, daoField.joinParse.GetData()...)
	// daoThis.functions = append(daoThis.functions, daoField.functions.GetData()...)
}

func (daoThis *myGenDao) Merge(daoOther myGenDao) {
	daoThis.importDao = append(daoThis.importDao, daoOther.importDao...)
	daoThis.filterParse = append(daoThis.filterParse, daoOther.filterParse...)
	daoThis.fieldParse = append(daoThis.fieldParse, daoOther.fieldParse...)
	daoThis.fieldHook = append(daoThis.fieldHook, daoOther.fieldHook...)
	daoThis.insertParseBefore = append(daoThis.insertParseBefore, daoOther.insertParseBefore...)
	daoThis.insertParse = append(daoThis.insertParse, daoOther.insertParse...)
	daoThis.insertHookBefore = append(daoThis.insertHookBefore, daoOther.insertHookBefore...)
	daoThis.insertHook = append(daoThis.insertHook, daoOther.insertHook...)
	daoThis.insertHookAfter = append(daoThis.insertHookAfter, daoOther.insertHookAfter...)
	daoThis.updateParse = append(daoThis.updateParse, daoOther.updateParse...)
	daoThis.updateHookBefore = append(daoThis.updateHookBefore, daoOther.updateHookBefore...)
	daoThis.updateHookAfter = append(daoThis.updateHookAfter, daoOther.updateHookAfter...)
	daoThis.deleteHook = append(daoThis.deleteHook, daoOther.deleteHook...)
	daoThis.groupParse = append(daoThis.groupParse, daoOther.groupParse...)
	daoThis.orderParse = append(daoThis.orderParse, daoOther.orderParse...)
	daoThis.joinParse = append(daoThis.joinParse, daoOther.joinParse...)
}

func (daoThis *myGenDao) Unique() {
	daoThis.importDao = garray.NewStrArrayFrom(daoThis.importDao).Unique().Slice()
	// daoThis.filterParse = garray.NewStrArrayFrom(daoThis.filterParse).Unique().Slice()
	// daoThis.fieldParse = garray.NewStrArrayFrom(daoThis.fieldParse).Unique().Slice()
	// daoThis.fieldHook = garray.NewStrArrayFrom(daoThis.fieldHook).Unique().Slice()
	// daoThis.insertParseBefore = garray.NewStrArrayFrom(daoThis.insertParseBefore).Unique().Slice()
	// daoThis.insertParse = garray.NewStrArrayFrom(daoThis.insertParse).Unique().Slice()
	// daoThis.insertHookBefore = garray.NewStrArrayFrom(daoThis.insertHookBefore).Unique().Slice()
	// daoThis.insertHook = garray.NewStrArrayFrom(daoThis.insertHook).Unique().Slice()
	// daoThis.insertHookAfter = garray.NewStrArrayFrom(daoThis.insertHookAfter).Unique().Slice()
	// daoThis.updateParse = garray.NewStrArrayFrom(daoThis.updateParse).Unique().Slice()
	// daoThis.updateHookBefore = garray.NewStrArrayFrom(daoThis.updateHookBefore).Unique().Slice()
	// daoThis.updateHookAfter = garray.NewStrArrayFrom(daoThis.updateHookAfter).Unique().Slice()
	// daoThis.groupParse = garray.NewStrArrayFrom(daoThis.groupParse).Unique().Slice()
	// daoThis.orderParse = garray.NewStrArrayFrom(daoThis.orderParse).Unique().Slice()
	// daoThis.joinParse = garray.NewStrArrayFrom(daoThis.joinParse).Unique().Slice()
}

// dao生成
func genDao(tpl myGenTpl) {
	tpl.gfGenDao(true) //dao文件生成
	saveFile := gfile.SelfDir() + `/internal/dao/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseSnake + `.go`
	tplDao := gfile.GetContents(saveFile)

	dao := getDaoIdAndLabel(tpl)
	for _, v := range tpl.FieldListOfDefault {
		dao.Add(getDaoField(tpl, v))
	}
	for _, v := range tpl.FieldListOfAfter {
		dao.Add(getDaoField(tpl, v))
	}
	for _, v := range tpl.Handle.ExtendTableOneList {
		genDao(v.tpl)
		dao.Merge(getDaoExtendMiddleOne(v))
	}
	for _, v := range tpl.Handle.MiddleTableOneList {
		genDao(v.tpl)
		dao.Merge(getDaoExtendMiddleOne(v))
	}
	for _, v := range tpl.Handle.ExtendTableManyList {
		v.tpl.gfGenDao(false)
		dao.Merge(getDaoExtendMiddleMany(v))
	}
	for _, v := range tpl.Handle.MiddleTableManyList {
		v.tpl.gfGenDao(false)
		dao.Merge(getDaoExtendMiddleMany(v))
	}
	dao.Unique()

	if len(dao.importDao) > 0 {
		importDaoPoint := `"api/internal/dao/` + tpl.ModuleDirCaseKebab + `/internal"`
		tplDao = gstr.Replace(tplDao, importDaoPoint, importDaoPoint+gstr.Join(append([]string{``}, dao.importDao...), `
	`), 1)
	}
	tplDao = gstr.Replace(tplDao, `"github.com/gogf/gf/v2/util/gconv"`, `"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/gjson"
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
				if daoModel.IsAutoField && !daoThis.ColumnArr().Contains(k) {
					continue
				}
				insertData[k] = v`
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
		tplDao = gstr.Replace(tplDao, insertHookPoint, `id, _ := result.LastInsertId()

			`+gstr.Join(append(dao.insertHookBefore, ``), `
			`)+`for k, v := range daoModel.AfterInsert {
				switch k {`+gstr.Join(append([]string{``}, dao.insertHook...), `
				`)+`
				}
			}`+gstr.Join(append([]string{``}, dao.insertHookAfter...), `
			`), 1)
	}

	// 解析update
	if len(dao.updateParse) > 0 {
		updateParsePoint := `default:
				if daoModel.IsAutoField && !daoThis.ColumnArr().Contains(k) {
					continue
				}
				updateData[k] = v`
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
			tplDao = gstr.Replace(tplDao, updateHookPoint, `for k, v := range daoModel.AfterUpdate {
				switch k {`+gstr.Join(append([]string{``}, dao.updateHookBefore...), `
				`)+`
				}
			}

			`+updateHookPoint, 1)
		}
		if len(dao.updateHookAfter) > 0 {
			tplDao = gstr.Replace(tplDao, updateHookPoint, `row, _ := result.RowsAffected()
			if row == 0 {
				return
			}

			for k, v := range daoModel.AfterUpdate {
				switch k {`+gstr.Join(append([]string{``}, dao.updateHookAfter...), `
				`)+`
				}
			}`, 1)
		}
	}

	// 解析delete
	if len(dao.deleteHook) > 0 {
		deleteHookPoint := `/* row, _ := result.RowsAffected()
			if row == 0 {
				return
			} */
			return`
		tplDao = gstr.Replace(tplDao, deleteHookPoint, `row, _ := result.RowsAffected()
			if row == 0 {
				return
			}

			`+gstr.Join(append(dao.deleteHook, ``), `
			`)+`return`, 1)
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
		dao.filterParse = append(dao.filterParse, `case `+"`id`, `"+internal.GetStrByFieldStyle(tpl.FieldStyle, `id_arr`)+"`"+`:
				m = m.Where(daoModel.DbTable+`+"`.`"+`+daoThis.PrimaryKey(), v)`)
		dao.filterParse = append(dao.filterParse, `case `+"`"+internal.GetStrByFieldStyle(tpl.FieldStyle, `exc_id`)+"`, `"+internal.GetStrByFieldStyle(tpl.FieldStyle, `exc_id_arr`)+"`"+`:
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
					updateData[daoThis.PrimaryKey()] = v`)
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
		dao.filterParse = append(dao.filterParse, `case `+"`id`, `"+internal.GetStrByFieldStyle(tpl.FieldStyle, `id_arr`)+"`"+`:
				idArr := []string{gconv.String(v)}
				if gvar.New(v).IsSlice() {
					idArr = gconv.SliceStr(v)
				}
				inStrArr := []string{}
				for _, id := range idArr {
					inStrArr = append(inStrArr, `+"`('`+gstr.Replace(id, `"+concatStr+"`, `', '`)+`')`)"+`
				}
				m = m.Where(`+"`(`"+gstr.Join(filterParseStrArr, "`, `")+"`) IN (` + gstr.Join(inStrArr, `, `) + `)`)")
		dao.filterParse = append(dao.filterParse, `case `+"`"+internal.GetStrByFieldStyle(tpl.FieldStyle, `exc_id`)+"`, `"+internal.GetStrByFieldStyle(tpl.FieldStyle, `exc_id_arr`)+"`"+`:
				idArr := []string{gconv.String(v)}
				if gvar.New(v).IsSlice() {
					idArr = gconv.SliceStr(v)
				}
				inStrArr := []string{}
				for _, id := range idArr {
					inStrArr = append(inStrArr, `+"`('`+gstr.Replace(id, `"+concatStr+"`, `', '`)+`')`)"+`
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

	fieldParseStr := `case ` + "`label`" + `:
				m = m.Fields(daoModel.DbTable + ` + "`.`" + ` + daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.LabelList[0]) + ` + ` + "` AS `" + ` + v)`
	filterParseStr := `case ` + "`label`" + `:
				m = m.WhereLike(daoModel.DbTable+` + "`.`" + `+daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.LabelList[0]) + `, ` + "`%`" + `+gconv.String(v)+` + "`%`" + `)`
	labelListLen := len(tpl.Handle.LabelList)
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
	return
}

func getDaoField(tpl myGenTpl, v myGenField) (daoField myGenDaoField) {
	daoPath := `daoThis`
	daoTable := `daoModel.DbTable`

	/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 开始--------*/
	switch v.FieldType {
	case internal.TypeInt: // `int等类型`
	case internal.TypeIntU: // `int等类型（unsigned）`
	case internal.TypeFloat: // `float等类型`
	case internal.TypeFloatU: // `float等类型（unsigned）`
	case internal.TypeVarchar, internal.TypeChar: // `varchar类型`	// `char类型`
		if gconv.Uint(v.FieldLimitStr) <= internal.ConfigMaxLenOfStrFilter {
			daoField.filterParse.Method = internal.ReturnType
		}
		if v.IsNull && v.IsUnique {
			daoField.insertParse.Method = internal.ReturnType
			daoField.insertParse.DataType = append(daoField.insertParse.DataType, `case `+daoPath+`.Columns().`+v.FieldCaseCamel+`:
				if gconv.String(v) == `+"``"+` {
					v = nil
				}
				insertData[k] = v`)

			daoField.updateParse.Method = internal.ReturnType
			daoField.updateParse.DataType = append(daoField.updateParse.DataType, `case `+daoPath+`.Columns().`+v.FieldCaseCamel+`:
				if gconv.String(v) == `+"``"+` {
					v = nil
				}
				updateData[k] = v`)
		}
	case internal.TypeText: // `text类型`
	case internal.TypeJson: // `json类型`
		if v.IsNull {
			daoField.insertParse.Method = internal.ReturnType
			daoField.insertParse.DataType = append(daoField.insertParse.DataType, `case `+daoPath+`.Columns().`+v.FieldCaseCamel+`:
				if gconv.String(v) == `+"``"+` {
					v = nil
				}
				insertData[k] = v`)

			daoField.updateParse.Method = internal.ReturnType
			daoField.updateParse.DataType = append(daoField.updateParse.DataType, `case `+daoPath+`.Columns().`+v.FieldCaseCamel+`:
				if gconv.String(v) == `+"``"+` {
					updateData[k] = nil
					continue
				}
				updateData[k] = v`)
		}
	case internal.TypeDatetime, internal.TypeTimestamp: // `datetime类型`	// `timestamp类型`
	case internal.TypeDate: // `date类型`
		daoField.filterParse.Method = internal.ReturnType
		daoField.orderParse.Method = internal.ReturnType
		daoField.orderParse.DataType = append(daoField.orderParse.DataType, `case `+daoPath+`.Columns().`+v.FieldCaseCamel+`:
				m = m.Order(`+daoTable+` + `+"`.`"+` + v)
				m = m.OrderDesc(daoModel.DbTable + `+"`.`"+` + daoThis.PrimaryKey())`) //追加主键倒序。mysql排序字段有重复值时，分页会导致同一条数据可能在不同页都出现
	case internal.TypeTime: // `time类型`
	default:
		daoField.filterParse.Method = internal.ReturnType
	}
	/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 结束--------*/

	/*--------根据字段主键类型处理 开始--------*/
	switch v.FieldTypePrimary {
	case internal.TypePrimary: // 独立主键
	case internal.TypePrimaryAutoInc: // 独立主键（自增）
		return myGenDaoField{}
	case internal.TypePrimaryMany: // 联合主键
	case internal.TypePrimaryManyAutoInc: // 联合主键（自增）
		return myGenDaoField{}
	}
	/*--------根据字段主键类型处理 结束--------*/

	/*--------根据字段命名类型处理 开始--------*/
	switch v.FieldTypeName {
	case internal.TypeNameDeleted: // 软删除字段
	case internal.TypeNameUpdated: // 更新时间字段
	case internal.TypeNameCreated: // 创建时间字段
		daoField.filterParse.Method = internal.ReturnTypeName
		daoField.filterParse.DataTypeName = append(daoField.filterParse.DataTypeName, `case `+"`"+internal.GetStrByFieldStyle(tpl.FieldStyle, `time_range_start`)+"`"+`:
				m = m.WhereGTE(`+daoTable+`+`+"`.`"+`+`+daoPath+`.Columns().`+v.FieldCaseCamel+`, v)
			case `+"`"+internal.GetStrByFieldStyle(tpl.FieldStyle, `time_range_end`)+"`"+`:
				m = m.WhereLTE(`+daoTable+`+`+"`.`"+`+`+daoPath+`.Columns().`+v.FieldCaseCamel+`, v)`)
	case internal.TypeNamePid: // pid；	类型：int等类型；
		daoField.filterParse.Method = internal.ReturnTypeName

		daoField.fieldParse.Method = internal.ReturnTypeName
		daoField.fieldParse.DataTypeName = append(daoField.fieldParse.DataTypeName, `case `+"`"+internal.GetStrByFieldStyle(tpl.FieldStyle, tpl.Handle.LabelList[0], `p`)+"`"+`:
				tableP := `+"`p_`"+` + `+daoTable+`
				m = m.Fields(tableP + `+"`.`"+` + `+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.LabelList[0])+` + `+"` AS `"+` + v)
				m = m.Handler(daoThis.ParseJoin(tableP, daoModel))`)
		daoField.fieldParse.DataTypeName = append(daoField.fieldParse.DataTypeName, `case `+"`tree`"+`:
				m = m.Fields(`+daoTable+` + `+"`.`"+` + `+daoPath+`.PrimaryKey())
				m = m.Fields(`+daoTable+` + `+"`.`"+` + `+daoPath+`.Columns().`+v.FieldCaseCamel+`)
				m = m.Handler(daoThis.ParseOrder([]string{`+"`tree`"+`}, daoModel))`)

		orderParseStr := `case ` + "`tree`" + `:
				m = m.OrderAsc(` + daoTable + ` + ` + "`.`" + ` + ` + daoPath + `.Columns().` + v.FieldCaseCamel + `)`
		if tpl.Handle.Pid.Sort != `` {
			orderParseStr += `
				m = m.OrderAsc(` + daoTable + ` + ` + "`.`" + ` + ` + daoPath + `.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Sort) + `)`
		}
		orderParseStr += `
				m = m.OrderAsc(daoModel.DbTable + ` + "`.`" + ` + daoThis.PrimaryKey())`
		daoField.orderParse.Method = internal.ReturnTypeName
		daoField.orderParse.DataTypeName = append(daoField.orderParse.DataTypeName, orderParseStr)

		daoField.joinParse.Method = internal.ReturnTypeName
		daoField.joinParse.DataTypeName = append(daoField.joinParse.DataTypeName, `case `+"`p_`"+` + `+daoTable+`:
			m = m.LeftJoin(`+daoTable+`+`+"` AS `"+`+joinTable, joinTable+`+"`.`"+`+`+daoPath+`.PrimaryKey()+`+"` = `"+`+`+daoTable+`+`+"`.`"+`+`+daoPath+`.Columns().`+v.FieldCaseCamel+`)`)

		if tpl.Handle.Pid.IsCoexist {
			daoField.insertParseBefore.Method = internal.ReturnTypeName
			daoField.insertParseBefore.DataTypeName = append(daoField.insertParseBefore.DataTypeName, `if _, ok := insert[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Pid)+`]; !ok {
			insert[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Pid)+`] = 0
		}`)

			daoField.insertParse.Method = internal.ReturnTypeName
			daoField.insertParse.DataTypeName = append(daoField.insertParse.DataTypeName, `case `+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Pid)+`:
				insertData[k] = v
				if gconv.Uint(v) > 0 {
					pInfo, _ := `+daoPath+`.CtxDaoModel(m.GetCtx()).Filter(`+daoPath+`.PrimaryKey(), v).One()
					daoModel.AfterInsert[`+"`pIdPath`"+`] = pInfo[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.IdPath)+`].String()
					daoModel.AfterInsert[`+"`pLevel`"+`] = pInfo[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Level)+`].Uint()
				} else {
					daoModel.AfterInsert[`+"`pIdPath`"+`] = `+"`0`"+`
					daoModel.AfterInsert[`+"`pLevel`"+`] = 0
				}`)

			daoField.insertHookBefore.Method = internal.ReturnTypeName
			daoField.insertHookBefore.DataTypeName = append(daoField.insertHookBefore.DataTypeName, `updateSelfData := map[string]interface{}{}`)

			daoField.insertHook.Method = internal.ReturnTypeName
			daoField.insertHook.DataTypeName = append(daoField.insertHook.DataTypeName,
				`case `+"`pIdPath`"+`:
					updateSelfData[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.IdPath)+`] = gconv.String(v) + `+"`-`"+` + gconv.String(id)`,
				`case `+"`pLevel`"+`:
					updateSelfData[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Level)+`] = gconv.Uint(v) + 1`,
			)
			daoField.insertHookAfter.Method = internal.ReturnTypeName
			daoField.insertHookAfter.DataTypeName = append(daoField.insertHookAfter.DataTypeName, `if len(updateSelfData) > 0 {
				daoModel.CloneNew().Filter(`+daoPath+`.PrimaryKey(), id).HookUpdate(updateSelfData).Update()
			}`)

			daoField.updateParse.Method = internal.ReturnTypeName
			updateChildIdPathAndLevelListVar := internal.GetStrByFieldStyle(tpl.FieldStyle, `update_child_id_path_and_level_list`)
			daoField.updateParse.DataTypeName = append(daoField.updateParse.DataTypeName, `case `+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Pid)+`:
				updateData[k] = v
				pIdPath := `+"`0`"+`
				var pLevel uint = 0
				if gconv.Uint(v) > 0 {
					pInfo, _ := `+daoPath+`.CtxDaoModel(m.GetCtx()).Filter(`+daoPath+`.PrimaryKey(), v).One()
					pIdPath = pInfo[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.IdPath)+`].String()
					pLevel = pInfo[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Level)+`].Uint()
				}
				updateData[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.IdPath)+`] = gdb.Raw(`+"`CONCAT('`"+` + pIdPath + `+"`-', `"+` + `+daoPath+`.PrimaryKey() + `+"`)`"+`)
				updateData[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Level)+`] = pLevel + 1
				//更新所有子孙级的idPath和level
				`+gstr.CaseCamelLower(updateChildIdPathAndLevelListVar)+` := []map[string]interface{}{}
				oldList, _ := `+daoPath+`.CtxDaoModel(m.GetCtx()).Filter(`+daoPath+`.PrimaryKey(), daoModel.IdArr).All()
				for _, oldInfo := range oldList {
					if gconv.Uint(v) != oldInfo[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Pid)+`].Uint() {
						`+gstr.CaseCamelLower(updateChildIdPathAndLevelListVar)+` = append(`+gstr.CaseCamelLower(updateChildIdPathAndLevelListVar)+`, map[string]interface{}{
							`+"`pIdPathOfOld`"+`: oldInfo[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.IdPath)+`],
							`+"`pIdPathOfNew`"+`: pIdPath + `+"`-`"+` + oldInfo[`+daoPath+`.PrimaryKey()].String(),
							`+"`pLevelOfOld`"+`:  oldInfo[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Level)+`],
							`+"`pLevelOfNew`"+`:  pLevel + 1,
						})
					}
				}
				if len(`+gstr.CaseCamelLower(updateChildIdPathAndLevelListVar)+`) > 0 {
					daoModel.AfterUpdate[`+"`"+updateChildIdPathAndLevelListVar+"`"+`] = `+gstr.CaseCamelLower(updateChildIdPathAndLevelListVar)+`
				}
			case `+"`childIdPath`"+`: //更新所有子孙级的idPath。参数：map[string]interface{}{`+"`pIdPathOfOld`"+`: `+"`父级IdPath（旧）`"+`, `+"`pIdPathOfNew`"+`: `+"`父级IdPath（新）`"+`}
				val := gconv.Map(v)
				pIdPathOfOld := gconv.String(val[`+"`pIdPathOfOld`"+`])
				pIdPathOfNew := gconv.String(val[`+"`pIdPathOfNew`"+`])
				updateData[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.IdPath)+`] = gdb.Raw(`+"`REPLACE(`"+` + `+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.IdPath)+` + `+"`, '`"+` + pIdPathOfOld + `+"`', '`"+` + pIdPathOfNew + `+"`')`"+`)
			case `+"`childLevel`"+`: //更新所有子孙级的level。参数：map[string]interface{}{`+"`pLevelOfOld`"+`: `+"`父级Level（旧）`"+`, `+"`pLevelOfNew`"+`: `+"`父级Level（新）`"+`}
				val := gconv.Map(v)
				pLevelOfOld := gconv.Uint(val[`+"`pLevelOfOld`"+`])
				pLevelOfNew := gconv.Uint(val[`+"`pLevelOfNew`"+`])
				updateData[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Level)+`] = gdb.Raw(`+daoTable+` + `+"`.`"+` + `+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Level)+` + `+"` + `"+` + gconv.String(pLevelOfNew-pLevelOfOld))
				if pLevelOfNew < pLevelOfOld {
					updateData[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Level)+`] = gdb.Raw(`+daoTable+` + `+"`.`"+` + `+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Level)+` + `+"` - `"+` + gconv.String(pLevelOfOld-pLevelOfNew))
				}`)

			daoField.updateHookAfter.Method = internal.ReturnTypeName
			daoField.updateHookAfter.DataTypeName = append(daoField.updateHookAfter.DataTypeName, `case `+"`"+updateChildIdPathAndLevelListVar+"`"+`: //修改pid时，更新所有子孙级的idPath和level。参数：[]map[string]interface{}{`+"`pIdPathOfOld`"+`: `+"`父级IdPath（旧）`"+`, `+"`pIdPathOfNew`"+`: `+"`父级IdPath（新）`"+`, `+"`pLevelOfOld`"+`: `+"`父级Level（旧）`"+`, `+"`pLevelOfNew`"+`: `+"`父级Level（新）`"+`}
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
					}`)

			daoField.filterParse.Method = internal.ReturnTypeName
			daoField.filterParse.DataTypeName = append(daoField.filterParse.DataTypeName, `case `+"`pIdPathOfOld`"+`: //父级IdPath（旧）
				m = m.WhereLike(`+daoTable+`+`+"`.`"+`+`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.IdPath)+`, gconv.String(v)+`+"`-%`"+`)`)
		}
	case internal.TypeNameLevel: // level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；
		daoField.filterParse.Method = internal.ReturnTypeName

		daoField.orderParse.Method = internal.ReturnTypeName
		daoField.orderParse.DataTypeName = append(daoField.orderParse.DataTypeName, `case `+daoPath+`.Columns().`+v.FieldCaseCamel+`:
				m = m.Order(`+daoTable+` + `+"`.`"+` + v)
				m = m.OrderDesc(daoModel.DbTable + `+"`.`"+` + daoThis.PrimaryKey())`) //追加主键倒序。mysql排序字段有重复值时，分页会导致同一条数据可能在不同页都出现
	case internal.TypeNameIdPath: // idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；
		return myGenDaoField{}
	case internal.TypeNamePasswordSuffix: // password,passwd后缀；		类型：char(32)；
		insertParseStr := `case ` + daoPath + `.Columns().` + v.FieldCaseCamel + `:
				password := gconv.String(v)
				if len(password) != 32 {
					password = gmd5.MustEncrypt(password)
				}`
		updateParseStr := `case ` + daoPath + `.Columns().` + v.FieldCaseCamel + `:
				password := gconv.String(v)
				if len(password) != 32 {
					password = gmd5.MustEncrypt(password)
				}`
		passwordMapKey := internal.GetHandlePasswordMapKey(v.FieldRaw)
		if tpl.Handle.PasswordMap[passwordMapKey].IsCoexist {
			insertParseStr += `
				salt := grand.S(` + tpl.Handle.PasswordMap[passwordMapKey].SaltLength + `)
				insertData[` + daoPath + `.Columns().` + gstr.CaseCamel(tpl.Handle.PasswordMap[passwordMapKey].SaltField) + `] = salt
				password = gmd5.MustEncrypt(password + salt)`
			updateParseStr += `
				salt := grand.S(` + tpl.Handle.PasswordMap[passwordMapKey].SaltLength + `)
				updateData[` + daoPath + `.Columns().` + gstr.CaseCamel(tpl.Handle.PasswordMap[passwordMapKey].SaltField) + `] = salt
				password = gmd5.MustEncrypt(password + salt)`
		}
		insertParseStr += `
				insertData[k] = password`
		updateParseStr += `
				updateData[k] = password`

		daoField.insertParse.Method = internal.ReturnTypeName
		daoField.insertParse.DataTypeName = append(daoField.insertParse.DataTypeName, insertParseStr)
		daoField.updateParse.Method = internal.ReturnTypeName
		daoField.updateParse.DataTypeName = append(daoField.updateParse.DataTypeName, updateParseStr)
	case internal.TypeNameSaltSuffix: // salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；
		return myGenDaoField{}
	case internal.TypeNameNameSuffix: // name,title后缀；	类型：varchar；
		daoField.filterParse.Method = internal.ReturnTypeName
		daoField.filterParse.DataTypeName = append(daoField.filterParse.DataTypeName, `case `+daoPath+`.Columns().`+v.FieldCaseCamel+`:
				m = m.WhereLike(`+daoTable+`+`+"`.`"+`+k, `+"`%`"+`+gconv.String(v)+`+"`%`"+`)`)
	case internal.TypeNameCodeSuffix: // code后缀；	类型：varchar；
	case internal.TypeNameAccountSuffix: // account后缀；	类型：varchar；
	case internal.TypeNamePhoneSuffix: // phone,mobile后缀；	类型：varchar；
	case internal.TypeNameEmailSuffix: // email后缀；	类型：varchar；
	case internal.TypeNameUrlSuffix: // url,link后缀；	类型：varchar；
	case internal.TypeNameIpSuffix: // IP后缀；	类型：varchar；
	case internal.TypeNameIdSuffix: // id后缀；	类型：int等类型；
		relIdObj := tpl.Handle.RelIdMap[v.FieldRaw]
		daoField.filterParse.Method = internal.ReturnTypeName
		if relIdObj.tpl.Table != `` {
			daoPathRel := relIdObj.tpl.TableCaseCamel
			daoTableRel := `table` + relIdObj.tpl.TableCaseCamel
			if relIdObj.tpl.ModuleDirCaseKebab != tpl.ModuleDirCaseKebab {
				daoField.importDao = append(daoField.importDao, `dao`+relIdObj.tpl.ModuleDirCaseCamel+` "api/internal/dao/`+relIdObj.tpl.ModuleDirCaseKebab+`"`)
				daoPathRel = `dao` + relIdObj.tpl.ModuleDirCaseCamel + `.` + relIdObj.tpl.TableCaseCamel
				if relIdObj.tpl.ModuleDirCaseCamel != relIdObj.tpl.TableCaseCamel {
					daoTableRel = `table` + relIdObj.tpl.ModuleDirCaseCamel + relIdObj.tpl.TableCaseCamel
				}
			}

			if !relIdObj.IsRedundName {
				fieldParseStr := `case ` + daoPathRel + `.Columns().` + gstr.CaseCamel(relIdObj.tpl.Handle.LabelList[0]) + `:` + `
				` + daoTableRel + ` := ` + daoPathRel + `.ParseDbTable(m.GetCtx())
				m = m.Fields(` + daoTableRel + ` + ` + "`.`" + ` + v)
				m = m.Handler(daoThis.ParseJoin(` + daoTableRel + `, daoModel))`
				if relIdObj.Suffix != `` {
					fieldParseStr = `case ` + daoPathRel + `.Columns().` + gstr.CaseCamel(relIdObj.tpl.Handle.LabelList[0]) + " + `" + relIdObj.Suffix + "`:" + `
				` + daoTableRel + relIdObj.SuffixCaseCamel + ` := ` + daoPathRel + `.ParseDbTable(m.GetCtx()) + ` + "`" + relIdObj.SuffixCaseSnake + "`" + `
				m = m.Fields(` + daoTableRel + relIdObj.SuffixCaseCamel + ` + ` + "`.`" + ` + ` + daoPathRel + `.Columns().` + gstr.CaseCamel(relIdObj.tpl.Handle.LabelList[0]) + ` + ` + "` AS `" + ` + v)
				m = m.Handler(daoThis.ParseJoin(` + daoTableRel + relIdObj.SuffixCaseCamel + `, daoModel))`
				}
				daoField.fieldParse.Method = internal.ReturnTypeName
				daoField.fieldParse.DataTypeName = append(daoField.fieldParse.DataTypeName, fieldParseStr)
			}

			joinParseStr := `case ` + daoPathRel + `.ParseDbTable(m.GetCtx()):
			m = m.LeftJoin(joinTable, joinTable+` + "`.`" + `+` + daoPathRel + `.PrimaryKey()+` + "` = `" + `+` + daoTable + `+` + "`.`" + `+` + daoPath + `.Columns().` + v.FieldCaseCamel + `)`
			if relIdObj.Suffix != `` {
				joinParseStr = `case ` + daoPathRel + `.ParseDbTable(m.GetCtx()) + ` + "`" + relIdObj.SuffixCaseSnake + "`" + `:
			m = m.LeftJoin(` + daoPathRel + `.ParseDbTable(m.GetCtx())+` + "` AS `" + `+joinTable, joinTable+` + "`.`" + `+` + daoPathRel + `.PrimaryKey()+` + "` = `" + `+` + daoTable + `+` + "`.`" + `+` + daoPath + `.Columns().` + v.FieldCaseCamel + `)`
			}
			daoField.joinParse.Method = internal.ReturnTypeName
			daoField.joinParse.DataTypeName = append(daoField.joinParse.DataTypeName, joinParseStr)
		}
	case internal.TypeNameSortSuffix, internal.TypeNameSort: // sort,weight等后缀；	类型：int等类型； // sort，且pid,level,idPath|id_path,sort同时存在时（才）有效；	类型：int等类型；
		daoField.orderParse.Method = internal.ReturnTypeName
		daoField.orderParse.DataTypeName = append(daoField.orderParse.DataTypeName, `case `+daoPath+`.Columns().`+v.FieldCaseCamel+`:
				m = m.Order(`+daoTable+` + `+"`.`"+` + v)
				m = m.OrderDesc(daoModel.DbTable + `+"`.`"+` + daoThis.PrimaryKey())`) //追加主键倒序。mysql排序字段有重复值时，分页会导致同一条数据可能在不同页都出现
	case internal.TypeNameStatusSuffix: // status,type,method,pos,position,gender等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
		daoField.filterParse.Method = internal.ReturnTypeName
	case internal.TypeNameIsPrefix: // is_前缀；		类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（停用：0否 1是）
		daoField.filterParse.Method = internal.ReturnTypeName
	case internal.TypeNameStartPrefix: // start_前缀；	类型：datetime或date或timestamp或time；
		filterParseStr := `m = m.WhereLTE(` + daoTable + `+` + "`.`" + `+k, v)`
		if v.IsNull {
			filterParseStr = `m = m.Where(m.Builder().WhereLTE(` + daoTable + `+` + "`.`" + `+k, v).WhereOrNull(` + daoTable + ` + ` + "`.`" + ` + k))`
		}
		daoField.filterParse.Method = internal.ReturnTypeName
		daoField.filterParse.DataTypeName = append(daoField.filterParse.DataTypeName, `case `+daoPath+`.Columns().`+v.FieldCaseCamel+`:
				`+filterParseStr)
	case internal.TypeNameEndPrefix: // end_前缀；	类型：datetime或date或timestamp或time；
		filterParseStr := `m = m.WhereGTE(` + daoTable + `+` + "`.`" + `+k, v)`
		if v.IsNull {
			filterParseStr = `m = m.Where(m.Builder().WhereGTE(` + daoTable + `+` + "`.`" + `+k, v).WhereOrNull(` + daoTable + ` + ` + "`.`" + ` + k))`
		}
		daoField.filterParse.Method = internal.ReturnTypeName
		daoField.filterParse.DataTypeName = append(daoField.filterParse.DataTypeName, `case `+daoPath+`.Columns().`+v.FieldCaseCamel+`:
				`+filterParseStr)
	case internal.TypeNameRemarkSuffix: // remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器
		daoField.filterParse.Method = internal.ReturnEmpty
	case internal.TypeNameImageSuffix, internal.TypeNameVideoSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text	// video,video_list,videoList,video_arr,videoArr等后缀；		类型：单视频varchar，多视频json或text
		if v.FieldType == internal.TypeVarchar {
			daoField.filterParse.Method = internal.ReturnEmpty
		}
	case internal.TypeNameArrSuffix: // list,arr等后缀；	类型：json或text；
	}
	/*--------根据字段命名类型处理 结束--------*/
	return
}

func getDaoExtendMiddleOne(tplEM handleExtendMiddle) (dao myGenDao) {
	tpl := tplEM.tpl
	if tpl.ModuleDirCaseKebab != tplEM.tplOfTop.ModuleDirCaseKebab {
		dao.importDao = append(dao.importDao, `dao`+tpl.ModuleDirCaseCamel+` "api/internal/dao/`+tpl.ModuleDirCaseKebab+`"`)
	}

	dao.fieldParse = append(dao.fieldParse, `case `+gstr.Join(tplEM.FieldColumnArr, `, `)+`:
				`+tplEM.daoTableVar+` := `+tplEM.daoPath+`.ParseDbTable(m.GetCtx())
				m = m.Fields(`+tplEM.daoTableVar+` + `+"`.`"+` + v)
				m = m.Handler(daoThis.ParseJoin(`+tplEM.daoTableVar+`, daoModel))`)

	dao.insertParse = append(dao.insertParse, `case `+gstr.Join(tplEM.FieldColumnArr, `, `)+`:
				if garray.NewStrArrayFrom([]string{`+"``, `0`, `[]`, `{}`"+`}).Contains(gconv.String(v)) { //gvar.New(v).IsEmpty()无法验证指针的值是空的数据
					continue
				}
				insertData, ok := daoModel.AfterInsert[`+"`"+tplEM.FieldVar+"`"+`].(map[string]interface{})
				if !ok {
					insertData = map[string]interface{}{}
				}
				insertData[k] = v
				daoModel.AfterInsert[`+"`"+tplEM.FieldVar+"`"+`] = insertData`)
	insertHookStr := `insertData[` + tplEM.daoPath + `.Columns().` + gstr.CaseCamel(tplEM.RelId) + `] = id
					` + tplEM.daoPath + `.CtxDaoModel(ctx).HookInsert(insertData).Insert()`
	switch tplEM.TableType {
	case internal.TableTypeExtendOne:
		dao.insertHook = append(dao.insertHook, `case `+"`"+tplEM.FieldVar+"`"+`:
					insertData, _ := v.(map[string]interface{})
					`+insertHookStr)
	case internal.TableTypeMiddleOne:
		insertHookIdSuffixArr := []string{}
		insertHookIdSuffixIfArr := []string{}
		for k, v := range tplEM.FieldListOfIdSuffix {
			insertHookIdSuffixArr = append(insertHookIdSuffixArr, `_, ok`+v.FieldCaseCamel+` := insertData[`+tplEM.FieldColumnArrOfIdSuffix[k]+`]`)
			insertHookIdSuffixIfArr = append(insertHookIdSuffixIfArr, `!ok`+v.FieldCaseCamel)
		}
		dao.insertHook = append(dao.insertHook, `case `+"`"+tplEM.FieldVar+"`"+`:
					insertData, _ := v.(map[string]interface{})
					`+gstr.Join(append(insertHookIdSuffixArr, ``), `
					`)+`if `+gstr.Join(insertHookIdSuffixIfArr, ` && `)+` { //多ID时，全部ID都不存在（都等于0已在ParseInsert解析时已过滤，故存在就肯定不等于0）不插入。可根据自己业务修改
						continue
					}
					`+insertHookStr)
	}

	dao.updateParse = append(dao.updateParse, `case `+gstr.Join(tplEM.FieldColumnArr, `, `)+`:
				updateData, ok := daoModel.AfterUpdate[`+"`"+tplEM.FieldVar+"`"+`].(map[string]interface{})
				if !ok {
					updateData = map[string]interface{}{}
				}
				updateData[k] = v
				daoModel.AfterUpdate[`+"`"+tplEM.FieldVar+"`"+`] = updateData`)
	updateHookBeforeStr := `for _, id := range daoModel.IdArr {
						updateData[` + tplEM.daoPath + `.Columns().` + gstr.CaseCamel(tplEM.RelId) + `] = id
						if row, _ := ` + tplEM.daoPath + `.CtxDaoModel(ctx).Filter(` + tplEM.daoPath + `.Columns().` + gstr.CaseCamel(tplEM.RelId) + `, id).HookUpdate(updateData).UpdateAndGetAffected(); row == 0 { //更新失败，有可能记录不存在，这时做插入操作
							` + tplEM.daoPath + `.CtxDaoModel(ctx).HookInsert(updateData).Insert()
						}
					}`
	switch tplEM.TableType {
	case internal.TableTypeExtendOne:
		dao.updateHookBefore = append(dao.updateHookBefore, `case `+"`"+tplEM.FieldVar+"`"+`:
					updateData, _ := v.(map[string]interface{})
					`+updateHookBeforeStr)
	case internal.TableTypeMiddleOne:
		updateHookBeforeIdSuffixArr := []string{}
		updateHookBeforeIdSuffixIfArr := []string{}
		for k, v := range tplEM.FieldListOfIdSuffix {
			updateHookBeforeIdSuffixArr = append(updateHookBeforeIdSuffixArr, gstr.CaseCamelLower(v.FieldRaw)+`, ok`+v.FieldCaseCamel+` := updateData[`+tplEM.FieldColumnArrOfIdSuffix[k]+`]`)
			updateHookBeforeIdSuffixIfArr = append(updateHookBeforeIdSuffixIfArr, `(ok`+v.FieldCaseCamel+` && gconv.Uint(`+gstr.CaseCamelLower(v.FieldRaw)+`) == 0)`)
		}
		dao.updateHookBefore = append(dao.updateHookBefore, `case `+"`"+tplEM.FieldVar+"`"+`:
					updateData, _ := v.(map[string]interface{})
					`+gstr.Join(append(updateHookBeforeIdSuffixArr, ``), `
					`)+`if `+gstr.Join(updateHookBeforeIdSuffixIfArr, ` && `)+` { //多ID时，全部ID存在且等于0就删除。可根据自己业务修改
						for _, id := range daoModel.IdArr {
							`+tplEM.daoPath+`.CtxDaoModel(ctx).Filter(`+tplEM.daoPath+`.Columns().`+gstr.CaseCamel(tplEM.RelId)+`, id).Delete()
						}
						continue
					}
					`+updateHookBeforeStr)
	}

	dao.deleteHook = append(dao.deleteHook, tplEM.daoPath+`.CtxDaoModel(ctx).Filter(`+tplEM.daoPath+`.Columns().`+gstr.CaseCamel(tplEM.RelId)+`, daoModel.IdArr).Delete()`)

	dao.joinParse = append(dao.joinParse, `case `+tplEM.daoPath+`.ParseDbTable(m.GetCtx()):
			m = m.LeftJoin(joinTable, joinTable+`+"`.`"+`+`+tplEM.daoPath+`.Columns().`+gstr.CaseCamel(tplEM.RelId)+`+`+"` = `"+`+daoModel.DbTable+`+"`.`"+`+daoThis.PrimaryKey())`)

	fieldArrOfFilter := []string{}
	daoFieldList := []myGenDaoField{}
	for _, v := range tplEM.FieldList {
		daoField := myGenDaoField{}
		/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 开始--------*/
		switch v.FieldType {
		case internal.TypeInt: // `int等类型`
		case internal.TypeIntU: // `int等类型（unsigned）`
		case internal.TypeFloat: // `float等类型`
		case internal.TypeFloatU: // `float等类型（unsigned）`
		case internal.TypeVarchar, internal.TypeChar: // `varchar类型`	// `char类型`
			if gconv.Uint(v.FieldLimitStr) <= internal.ConfigMaxLenOfStrFilter {
				daoField.filterParse.Method = internal.ReturnType
			}
		case internal.TypeText: // `text类型`
		case internal.TypeJson: // `json类型`
		case internal.TypeDatetime, internal.TypeTimestamp: // `datetime类型`	// `timestamp类型`
		case internal.TypeDate: // `date类型`
			daoField.filterParse.Method = internal.ReturnType
			daoField.orderParse.Method = internal.ReturnType
			daoField.orderParse.DataType = append(daoField.orderParse.DataType, `case `+tplEM.daoPath+`.Columns().`+v.FieldCaseCamel+`:
				`+tplEM.daoTableVar+` := `+tplEM.daoPath+`.ParseDbTable(m.GetCtx())
				m = m.Order(`+tplEM.daoTableVar+` + `+"`.`"+` + v)
				m = m.OrderDesc(daoModel.DbTable + `+"`.`"+` + daoThis.PrimaryKey())
				m = m.Handler(daoThis.ParseJoin(`+tplEM.daoTableVar+`, daoModel))`) //追加主键倒序。mysql排序字段有重复值时，分页会导致同一条数据可能在不同页都出现
		case internal.TypeTime: // `time类型`
		default:
			daoField.filterParse.Method = internal.ReturnType
		}
		/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 结束--------*/

		/*--------根据字段主键类型处理 开始--------*/
		switch v.FieldTypePrimary {
		case internal.TypePrimary: // 独立主键
		case internal.TypePrimaryAutoInc: // 独立主键（自增）
			continue
		case internal.TypePrimaryMany: // 联合主键
		case internal.TypePrimaryManyAutoInc: // 联合主键（自增）
			continue
		}
		/*--------根据字段主键类型处理 结束--------*/

		/*--------根据字段命名类型处理 开始--------*/
		switch v.FieldTypeName {
		case internal.TypeNameDeleted: // 软删除字段
			continue
		case internal.TypeNameUpdated: // 更新时间字段
			continue
		case internal.TypeNameCreated: // 创建时间字段
			continue
		case internal.TypeNamePid: // pid；	类型：int等类型；
			continue
		case internal.TypeNameLevel: // level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；
			continue
		case internal.TypeNameIdPath: // idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；
			continue
		case internal.TypeNamePasswordSuffix: // password,passwd后缀；		类型：char(32)；
			continue
		case internal.TypeNameSaltSuffix: // salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；
			continue
		case internal.TypeNameNameSuffix: // name,title后缀；	类型：varchar；
			daoField.filterParse.Method = internal.ReturnTypeName
			daoField.filterParse.DataTypeName = append(daoField.filterParse.DataTypeName, `case `+tplEM.daoPath+`.Columns().`+v.FieldCaseCamel+`:
				`+tplEM.daoTableVar+` := `+tplEM.daoPath+`.ParseDbTable(m.GetCtx())
				m = m.WhereLike(`+tplEM.daoTableVar+`+`+"`.`"+`+k, `+"`%`"+`+gconv.String(v)+`+"`%`"+`)
				m = m.Handler(daoThis.ParseJoin(`+tplEM.daoTableVar+`, daoModel))`)
		case internal.TypeNameCodeSuffix: // code后缀；	类型：varchar；
		case internal.TypeNameAccountSuffix: // account后缀；	类型：varchar；
		case internal.TypeNamePhoneSuffix: // phone,mobile后缀；	类型：varchar；
		case internal.TypeNameEmailSuffix: // email后缀；	类型：varchar；
		case internal.TypeNameUrlSuffix: // url,link后缀；	类型：varchar；
		case internal.TypeNameIpSuffix: // IP后缀；	类型：varchar；
		case internal.TypeNameIdSuffix: // id后缀；	类型：int等类型；
			relIdObj := tpl.Handle.RelIdMap[v.FieldRaw]
			daoField.filterParse.Method = internal.ReturnTypeName
			if relIdObj.tpl.Table != `` {
				daoPathRel := relIdObj.tpl.TableCaseCamel
				daoTableRel := `table` + relIdObj.tpl.TableCaseCamel
				if relIdObj.tpl.ModuleDirCaseKebab != tplEM.tplOfTop.ModuleDirCaseKebab {
					daoField.importDao = append(daoField.importDao, `dao`+relIdObj.tpl.ModuleDirCaseCamel+` "api/internal/dao/`+relIdObj.tpl.ModuleDirCaseKebab+`"`)
					daoPathRel = `dao` + relIdObj.tpl.ModuleDirCaseCamel + `.` + relIdObj.tpl.TableCaseCamel
					if relIdObj.tpl.ModuleDirCaseCamel != relIdObj.tpl.TableCaseCamel {
						daoTableRel = `table` + relIdObj.tpl.ModuleDirCaseCamel + relIdObj.tpl.TableCaseCamel
					}
				}

				if !relIdObj.IsRedundName {
					fieldParseStr := `case ` + daoPathRel + `.Columns().` + gstr.CaseCamel(relIdObj.tpl.Handle.LabelList[0]) + `:` + `
				` + tplEM.daoTableVar + ` := ` + tplEM.daoPath + `.ParseDbTable(m.GetCtx())
				` + daoTableRel + ` := ` + daoPathRel + `.ParseDbTable(m.GetCtx())
				m = m.Fields(` + daoTableRel + ` + ` + "`.`" + ` + v)
				m = m.Handler(daoThis.ParseJoin(` + tplEM.daoTableVar + `, daoModel))
				m = m.Handler(daoThis.ParseJoin(` + daoTableRel + `, daoModel))`
					if relIdObj.Suffix != `` {
						fieldParseStr = `case ` + daoPathRel + `.Columns().` + gstr.CaseCamel(relIdObj.tpl.Handle.LabelList[0]) + " + `" + relIdObj.Suffix + "`:" + `
				` + tplEM.daoTableVar + ` := ` + tplEM.daoPath + `.ParseDbTable(m.GetCtx())
				` + daoTableRel + relIdObj.SuffixCaseCamel + ` := ` + daoPathRel + `.ParseDbTable(m.GetCtx()) + ` + "`" + relIdObj.SuffixCaseSnake + "`" + `
				m = m.Fields(` + daoTableRel + relIdObj.SuffixCaseCamel + ` + ` + "`.`" + ` + ` + daoPathRel + `.Columns().` + gstr.CaseCamel(relIdObj.tpl.Handle.LabelList[0]) + ` + ` + "` AS `" + ` + v)
				m = m.Handler(daoThis.ParseJoin(` + tplEM.daoTableVar + `, daoModel))
				m = m.Handler(daoThis.ParseJoin(` + daoTableRel + relIdObj.SuffixCaseCamel + `, daoModel))`
					}
					daoField.fieldParse.Method = internal.ReturnTypeName
					daoField.fieldParse.DataTypeName = append(daoField.fieldParse.DataTypeName, fieldParseStr)
				}

				joinParseStr := `case ` + daoPathRel + `.ParseDbTable(m.GetCtx()):
			m = m.LeftJoin(joinTable, joinTable+` + "`.`" + `+` + daoPathRel + `.PrimaryKey()+` + "` = `" + `+` + tplEM.daoTable + `+` + "`.`" + `+` + tplEM.daoPath + `.Columns().` + v.FieldCaseCamel + `)`
				if relIdObj.Suffix != `` {
					joinParseStr = `case ` + daoPathRel + `.ParseDbTable(m.GetCtx()) + ` + "`" + relIdObj.SuffixCaseSnake + "`" + `:
			m = m.LeftJoin(` + daoPathRel + `.ParseDbTable(m.GetCtx())+` + "` AS `" + `+joinTable, joinTable+` + "`.`" + `+` + daoPathRel + `.PrimaryKey()+` + "` = `" + `+` + tplEM.daoTable + `+` + "`.`" + `+` + tplEM.daoPath + `.Columns().` + v.FieldCaseCamel + `)`
				}
				daoField.joinParse.Method = internal.ReturnTypeName
				daoField.joinParse.DataTypeName = append(daoField.joinParse.DataTypeName, joinParseStr)
			}
		case internal.TypeNameSortSuffix, internal.TypeNameSort: // sort,weight等后缀；	类型：int等类型； // sort，且pid,level,idPath|id_path,sort同时存在时（才）有效；	类型：int等类型；
			daoField.orderParse.Method = internal.ReturnTypeName
			daoField.orderParse.DataTypeName = append(daoField.orderParse.DataTypeName, `case `+tplEM.daoPath+`.Columns().`+v.FieldCaseCamel+`:
				`+tplEM.daoTableVar+` := `+tplEM.daoPath+`.ParseDbTable(m.GetCtx())
				m = m.Order(`+tplEM.daoTableVar+` + `+"`.`"+` + v)
				m = m.OrderDesc(daoModel.DbTable + `+"`.`"+` + daoThis.PrimaryKey())
				m = m.Handler(daoThis.ParseJoin(`+tplEM.daoTableVar+`, daoModel))`) //追加主键倒序。mysql排序字段有重复值时，分页会导致同一条数据可能在不同页都出现
		case internal.TypeNameStatusSuffix: // status,type,method,pos,position,gender等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
			daoField.filterParse.Method = internal.ReturnTypeName
		case internal.TypeNameIsPrefix: // is_前缀；		类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（停用：0否 1是）
			daoField.filterParse.Method = internal.ReturnTypeName
		case internal.TypeNameStartPrefix: // start_前缀；	类型：datetime或date或timestamp或time；
			filterParseStr := `m = m.WhereLTE(` + tplEM.daoTable + `+` + "`.`" + `+k, v)`
			if v.IsNull {
				filterParseStr = `m = m.Where(m.Builder().WhereLTE(` + tplEM.daoTable + `+` + "`.`" + `+k, v).WhereOrNull(` + tplEM.daoTable + ` + ` + "`.`" + ` + k))`
			}
			daoField.filterParse.Method = internal.ReturnTypeName
			daoField.filterParse.DataTypeName = append(daoField.filterParse.DataTypeName, `case `+tplEM.daoPath+`.Columns().`+v.FieldCaseCamel+`:
				`+tplEM.daoTableVar+` := `+tplEM.daoPath+`.ParseDbTable(m.GetCtx())
				`+filterParseStr+`
				m = m.Handler(daoThis.ParseJoin(`+tplEM.daoTableVar+`, daoModel))`)
		case internal.TypeNameEndPrefix: // end_前缀；	类型：datetime或date或timestamp或time；
			filterParseStr := `m = m.WhereGTE(` + tplEM.daoTable + `+` + "`.`" + `+k, v)`
			if v.IsNull {
				filterParseStr = `m = m.Where(m.Builder().WhereGTE(` + tplEM.daoTable + `+` + "`.`" + `+k, v).WhereOrNull(` + tplEM.daoTable + ` + ` + "`.`" + ` + k))`
			}
			daoField.filterParse.Method = internal.ReturnTypeName
			daoField.filterParse.DataTypeName = append(daoField.filterParse.DataTypeName, `case `+tplEM.daoPath+`.Columns().`+v.FieldCaseCamel+`:
				`+tplEM.daoTableVar+` := `+tplEM.daoPath+`.ParseDbTable(m.GetCtx())
				`+filterParseStr+`
				m = m.Handler(daoThis.ParseJoin(`+tplEM.daoTableVar+`, daoModel))`)
		case internal.TypeNameRemarkSuffix: // remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器
			daoField.filterParse.Method = internal.ReturnEmpty
		case internal.TypeNameImageSuffix, internal.TypeNameVideoSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text	// video,video_list,videoList,video_arr,videoArr等后缀；		类型：单视频varchar，多视频json或text
			if v.FieldType == internal.TypeVarchar {
				daoField.filterParse.Method = internal.ReturnEmpty
			}
		case internal.TypeNameArrSuffix: // list,arr等后缀；	类型：json或text；
		}
		/*--------根据字段命名类型处理 结束--------*/

		if daoField.filterParse.Method != internal.ReturnEmpty && len(daoField.filterParse.GetData()) == 0 {
			fieldArrOfFilter = append(fieldArrOfFilter, tplEM.daoPath+`.Columns().`+v.FieldCaseCamel)
		}
		daoFieldList = append(daoFieldList, daoField)
	}

	if len(fieldArrOfFilter) > 0 {
		dao.filterParse = append(dao.filterParse, `case `+gstr.Join(fieldArrOfFilter, `, `)+`:
				`+tplEM.daoTableVar+` := `+tplEM.daoPath+`.ParseDbTable(m.GetCtx())
				m = m.Where(`+tplEM.daoTableVar+`+`+"`.`"+`+k, v)
				m = m.Handler(daoThis.ParseJoin(`+tplEM.daoTableVar+`, daoModel))`)
	}

	for _, daoField := range daoFieldList {
		dao.Add(daoField)
	}
	return
}

func getDaoExtendMiddleMany(tplEM handleExtendMiddle) (dao myGenDao) {
	tpl := tplEM.tpl
	if tpl.ModuleDirCaseKebab != tplEM.tplOfTop.ModuleDirCaseKebab {
		dao.importDao = append(dao.importDao, `dao`+tpl.ModuleDirCaseCamel+` "api/internal/dao/`+tpl.ModuleDirCaseKebab+`"`)
	}

	dao.fieldParse = append(dao.fieldParse, `case `+"`"+tplEM.FieldVar+"`"+`:
				m = m.Fields(daoModel.DbTable + `+"`.`"+` + daoThis.PrimaryKey())
				daoModel.AfterField.Add(v)`)
	dao.insertParse = append(dao.insertParse, `case `+"`"+tplEM.FieldVar+"`"+`:
				daoModel.AfterInsert[k] = v`)
	dao.updateParse = append(dao.updateParse, `case `+"`"+tplEM.FieldVar+"`"+`:
				daoModel.AfterUpdate[k] = v`)
	if len(tplEM.FieldList) == 1 {
		dao.fieldHook = append(dao.fieldHook, `case `+"`"+tplEM.FieldVar+"`"+`:
						`+gstr.CaseCamelLower(tplEM.FieldVar)+`, _ := `+tplEM.daoPath+`.CtxDaoModel(ctx).Filter(`+tplEM.daoPath+`.Columns().`+gstr.CaseCamel(tplEM.RelId)+`, record[daoThis.PrimaryKey()]).Array(`+tplEM.daoPath+`.Columns().`+tplEM.FieldList[0].FieldCaseCamel+`)
						record[v] = gvar.New(`+gstr.CaseCamelLower(tplEM.FieldVar)+`)`)
		dao.insertHook = append(dao.insertHook, `case `+"`"+tplEM.FieldVar+"`"+`:
					insertList := []map[string]interface{}{}
					for _, item := range gconv.SliceAny(v) {
						insertList = append(insertList, map[string]interface{}{
							`+tplEM.daoPath+`.Columns().`+gstr.CaseCamel(tplEM.RelId)+`: id,
							`+tplEM.daoPath+`.Columns().`+tplEM.FieldList[0].FieldCaseCamel+`: item,
						})
					}
					`+tplEM.daoPath+`.CtxDaoModel(ctx).Data(insertList).Insert()`)
		dao.updateHookBefore = append(dao.updateHookBefore, `case `+"`"+tplEM.FieldVar+"`"+`:
					// daoIndex.SaveArrRelManyWithSort(ctx, &`+tplEM.daoPath+`, `+tplEM.daoPath+`.Columns().`+gstr.CaseCamel(tplEM.RelId)+`, `+tplEM.daoPath+`.Columns().`+tplEM.FieldList[0].FieldCaseCamel+`, gconv.SliceAny(daoModel.IdArr), gconv.SliceAny(v)) // 有顺序要求时使用，同时注释下面代码
					valArr := gconv.SliceStr(v)
					for _, id := range daoModel.IdArr {
						daoIndex.SaveArrRelMany(ctx, &`+tplEM.daoPath+`, `+tplEM.daoPath+`.Columns().`+gstr.CaseCamel(tplEM.RelId)+`, `+tplEM.daoPath+`.Columns().`+tplEM.FieldList[0].FieldCaseCamel+`, id, valArr )
					}`)
	} else {
		dao.fieldHook = append(dao.fieldHook, `case `+"`"+tplEM.FieldVar+"`"+`:
						`+gstr.CaseCamelLower(tplEM.FieldVar)+`, _ := `+tplEM.daoPath+`.CtxDaoModel(ctx).Filter(`+tplEM.daoPath+`.Columns().`+gstr.CaseCamel(tplEM.RelId)+`, record[daoThis.PrimaryKey()]). /* OrderAsc(`+tplEM.daoPath+`.Columns().CreatedAt). */ All()	// 有顺序要求时使用，自定义OrderAsc
						record[v] = gvar.New(gjson.MustEncodeString(`+gstr.CaseCamelLower(tplEM.FieldVar)+`)) //转成json字符串，控制器中list.Structs(&res.List)和info.Struct(&res.Info)才有效`)
		dao.insertHook = append(dao.insertHook, `case `+"`"+tplEM.FieldVar+"`"+`:
					insertList := []map[string]interface{}{}
					for _, item := range gconv.SliceMap(v) {
						insertItem := gjson.New(gjson.MustEncodeString(item)).Map()
						insertItem[`+tplEM.daoPath+`.Columns().`+gstr.CaseCamel(tplEM.RelId)+`] = id
						insertList = append(insertList, insertItem)
					}
					`+tplEM.daoPath+`.CtxDaoModel(ctx).Data(insertList).Insert()`)
		switch tplEM.TableType {
		case internal.TableTypeExtendMany:
			dao.updateHookBefore = append(dao.updateHookBefore, `case `+"`"+tplEM.FieldVar+"`"+`:
					valList := gconv.SliceMap(v)
					daoIndex.SaveListRelManyWithSort(ctx, &`+tplEM.daoPath+`, `+tplEM.daoPath+`.Columns().`+gstr.CaseCamel(tplEM.RelId)+`, gconv.SliceAny(daoModel.IdArr), valList)`)
		case internal.TableTypeMiddleMany:
			dao.updateHookBefore = append(dao.updateHookBefore, `case `+"`"+tplEM.FieldVar+"`"+`:
					// daoIndex.SaveListRelManyWithSort(ctx, &`+tplEM.daoPath+`, `+tplEM.daoPath+`.Columns().`+gstr.CaseCamel(tplEM.RelId)+`, gconv.SliceAny(daoModel.IdArr), gconv.SliceMap(v)) // 有顺序要求时使用，同时注释下面代码
					valList := gconv.SliceMap(v)
					for _, id := range daoModel.IdArr {
						daoIndex.SaveListRelMany(ctx, &`+tplEM.daoPath+`, `+tplEM.daoPath+`.Columns().`+gstr.CaseCamel(tplEM.RelId)+`, []string{`+gstr.Join(tplEM.FieldColumnArrOfIdSuffix, `, `)+`}, id, valList )
					}`)
		}
	}

	dao.deleteHook = append(dao.deleteHook, tplEM.daoPath+`.CtxDaoModel(ctx).Filter(`+tplEM.daoPath+`.Columns().`+gstr.CaseCamel(tplEM.RelId)+`, daoModel.IdArr).Delete()`)

	dao.joinParse = append(dao.joinParse, `case `+tplEM.daoPath+`.ParseDbTable(m.GetCtx()):
			m = m.LeftJoin(joinTable, joinTable+`+"`.`"+`+`+tplEM.daoPath+`.Columns().`+gstr.CaseCamel(tplEM.RelId)+`+`+"` = `"+`+daoModel.DbTable+`+"`.`"+`+daoThis.PrimaryKey())`)

	fieldArrOfFilter := []string{}
	daoFieldList := []myGenDaoField{}
	for _, v := range tplEM.FieldList {
		daoField := myGenDaoField{}
		/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 开始--------*/
		switch v.FieldType {
		case internal.TypeInt: // `int等类型`
		case internal.TypeIntU: // `int等类型（unsigned）`
		case internal.TypeFloat: // `float等类型`
		case internal.TypeFloatU: // `float等类型（unsigned）`
		case internal.TypeVarchar, internal.TypeChar: // `varchar类型`	// `char类型`
			if gconv.Uint(v.FieldLimitStr) <= internal.ConfigMaxLenOfStrFilter {
				daoField.filterParse.Method = internal.ReturnType
			}
		case internal.TypeText: // `text类型`
		case internal.TypeJson: // `json类型`
		case internal.TypeDatetime, internal.TypeTimestamp: // `datetime类型`	// `timestamp类型`
		case internal.TypeDate: // `date类型`
			daoField.filterParse.Method = internal.ReturnType
			daoField.orderParse.Method = internal.ReturnType
			daoField.orderParse.DataType = append(daoField.orderParse.DataType, `case `+tplEM.daoPath+`.Columns().`+v.FieldCaseCamel+`:
				`+tplEM.daoTableVar+` := `+tplEM.daoPath+`.ParseDbTable(m.GetCtx())
				m = m.Order(`+tplEM.daoTableVar+` + `+"`.`"+` + v)
				m = m.OrderDesc(daoModel.DbTable + `+"`.`"+` + daoThis.PrimaryKey())
				m = m.Handler(daoThis.ParseJoin(`+tplEM.daoTableVar+`, daoModel))`) //追加主键倒序。mysql排序字段有重复值时，分页会导致同一条数据可能在不同页都出现
		case internal.TypeTime: // `time类型`
		default:
			daoField.filterParse.Method = internal.ReturnType
		}
		/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 结束--------*/

		/*--------根据字段主键类型处理 开始--------*/
		switch v.FieldTypePrimary {
		case internal.TypePrimary: // 独立主键
		case internal.TypePrimaryAutoInc: // 独立主键（自增）
			continue
		case internal.TypePrimaryMany: // 联合主键
		case internal.TypePrimaryManyAutoInc: // 联合主键（自增）
			continue
		}
		/*--------根据字段主键类型处理 结束--------*/

		/*--------根据字段命名类型处理 开始--------*/
		switch v.FieldTypeName {
		case internal.TypeNameDeleted: // 软删除字段
			continue
		case internal.TypeNameUpdated: // 更新时间字段
			continue
		case internal.TypeNameCreated: // 创建时间字段
			continue
		case internal.TypeNamePid: // pid；	类型：int等类型；
			continue
		case internal.TypeNameLevel: // level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；
			continue
		case internal.TypeNameIdPath: // idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；
			continue
		case internal.TypeNamePasswordSuffix: // password,passwd后缀；		类型：char(32)；
			continue
		case internal.TypeNameSaltSuffix: // salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；
			continue
		case internal.TypeNameNameSuffix: // name,title后缀；	类型：varchar；
			daoField.filterParse.Method = internal.ReturnTypeName
			daoField.filterParse.DataTypeName = append(daoField.filterParse.DataTypeName, `case `+tplEM.daoPath+`.Columns().`+v.FieldCaseCamel+`:
				`+tplEM.daoTableVar+` := `+tplEM.daoPath+`.ParseDbTable(m.GetCtx())
				m = m.WhereLike(`+tplEM.daoTableVar+`+`+"`.`"+`+k, `+"`%`"+`+gconv.String(v)+`+"`%`"+`)
				m = m.Handler(daoThis.ParseJoin(`+tplEM.daoTableVar+`, daoModel))`)
		case internal.TypeNameCodeSuffix: // code后缀；	类型：varchar；
		case internal.TypeNameAccountSuffix: // account后缀；	类型：varchar；
		case internal.TypeNamePhoneSuffix: // phone,mobile后缀；	类型：varchar；
		case internal.TypeNameEmailSuffix: // email后缀；	类型：varchar；
		case internal.TypeNameUrlSuffix: // url,link后缀；	类型：varchar；
		case internal.TypeNameIpSuffix: // IP后缀；	类型：varchar；
		case internal.TypeNameIdSuffix: // id后缀；	类型：int等类型；
			daoField.filterParse.Method = internal.ReturnTypeName
		case internal.TypeNameSortSuffix, internal.TypeNameSort: // sort,weight等后缀；	类型：int等类型； // sort，且pid,level,idPath|id_path,sort同时存在时（才）有效；	类型：int等类型；
		case internal.TypeNameStatusSuffix: // status,type,method,pos,position,gender等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
			daoField.filterParse.Method = internal.ReturnTypeName
		case internal.TypeNameIsPrefix: // is_前缀；		类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（停用：0否 1是）
			daoField.filterParse.Method = internal.ReturnTypeName
		case internal.TypeNameStartPrefix: // start_前缀；	类型：datetime或date或timestamp或time；
			filterParseStr := `m = m.WhereLTE(` + tplEM.daoTable + `+` + "`.`" + `+k, v)`
			if v.IsNull {
				filterParseStr = `m = m.Where(m.Builder().WhereLTE(` + tplEM.daoTable + `+` + "`.`" + `+k, v).WhereOrNull(` + tplEM.daoTable + ` + ` + "`.`" + ` + k))`
			}
			daoField.filterParse.Method = internal.ReturnTypeName
			daoField.filterParse.DataTypeName = append(daoField.filterParse.DataTypeName, `case `+tplEM.daoPath+`.Columns().`+v.FieldCaseCamel+`:
				`+tplEM.daoTableVar+` := `+tplEM.daoPath+`.ParseDbTable(m.GetCtx())
				`+filterParseStr+`
				m = m.Handler(daoThis.ParseJoin(`+tplEM.daoTableVar+`, daoModel))`)
		case internal.TypeNameEndPrefix: // end_前缀；	类型：datetime或date或timestamp或time；
			filterParseStr := `m = m.WhereGTE(` + tplEM.daoTable + `+` + "`.`" + `+k, v)`
			if v.IsNull {
				filterParseStr = `m = m.Where(m.Builder().WhereGTE(` + tplEM.daoTable + `+` + "`.`" + `+k, v).WhereOrNull(` + tplEM.daoTable + ` + ` + "`.`" + ` + k))`
			}
			daoField.filterParse.Method = internal.ReturnTypeName
			daoField.filterParse.DataTypeName = append(daoField.filterParse.DataTypeName, `case `+tplEM.daoPath+`.Columns().`+v.FieldCaseCamel+`:
				`+tplEM.daoTableVar+` := `+tplEM.daoPath+`.ParseDbTable(m.GetCtx())
				`+filterParseStr+`
				m = m.Handler(daoThis.ParseJoin(`+tplEM.daoTableVar+`, daoModel))`)
		case internal.TypeNameRemarkSuffix: // remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器
			daoField.filterParse.Method = internal.ReturnEmpty
		case internal.TypeNameImageSuffix, internal.TypeNameVideoSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text	// video,video_list,videoList,video_arr,videoArr等后缀；		类型：单视频varchar，多视频json或text
			if v.FieldType == internal.TypeVarchar {
				daoField.filterParse.Method = internal.ReturnEmpty
			}
		case internal.TypeNameArrSuffix: // list,arr等后缀；	类型：json或text；
		}
		/*--------根据字段命名类型处理 结束--------*/

		if daoField.filterParse.Method != internal.ReturnEmpty && len(daoField.filterParse.GetData()) == 0 {
			fieldArrOfFilter = append(fieldArrOfFilter, tplEM.daoPath+`.Columns().`+v.FieldCaseCamel)
		}
		daoFieldList = append(daoFieldList, daoField)
	}

	if len(fieldArrOfFilter) > 0 {
		dao.filterParse = append(dao.filterParse, `case `+gstr.Join(fieldArrOfFilter, `, `)+`:
				`+tplEM.daoTableVar+` := `+tplEM.daoPath+`.ParseDbTable(m.GetCtx())
				m = m.Where(`+tplEM.daoTableVar+`+`+"`.`"+`+k, v)
				m = m.Handler(daoThis.ParseJoin(`+tplEM.daoTableVar+`, daoModel))`)
	}

	for _, daoField := range daoFieldList {
		dao.Add(daoField)
	}
	return
}
