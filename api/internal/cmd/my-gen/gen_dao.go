package my_gen

import (
	"api/internal/cmd/my-gen/internal"
	"api/internal/utils"
	"slices"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type myGenDao struct {
	idParse    string
	labelParse string

	importDao          []string
	filterParse        []string
	fieldParse         []string
	fieldHook          []string
	insertParseBefore  []string
	insertParse        []string
	insertHook         []string
	updateParse        []string
	updateHookBefore   []string
	updateHookAfter    []string
	deleteHookBefore   []string
	deleteHookAfter    []string
	deleteHookOtherRel []string
	groupParse         []string
	orderParse         []string
	joinParse          []string
}

type myGenDaoField struct {
	importDao         []string
	filterParse       internal.MyGenDataSliceHandler
	fieldParse        internal.MyGenDataSliceHandler
	fieldHook         internal.MyGenDataSliceHandler
	insertParseBefore internal.MyGenDataSliceHandler
	insertParse       internal.MyGenDataSliceHandler
	insertHook        internal.MyGenDataSliceHandler
	updateParse       internal.MyGenDataSliceHandler
	updateHookBefore  internal.MyGenDataSliceHandler
	updateHookAfter   internal.MyGenDataSliceHandler
	deleteHookBefore  internal.MyGenDataSliceHandler
	deleteHookAfter   internal.MyGenDataSliceHandler
	orderParse        internal.MyGenDataSliceHandler
	joinParse         internal.MyGenDataSliceHandler
}

func (daoThis *myGenDao) Add(daoField myGenDaoField) {
	daoThis.importDao = append(daoThis.importDao, daoField.importDao...)
	daoThis.filterParse = append(daoThis.filterParse, daoField.filterParse.GetData()...)
	daoThis.fieldParse = append(daoThis.fieldParse, daoField.fieldParse.GetData()...)
	daoThis.fieldHook = append(daoThis.fieldHook, daoField.fieldHook.GetData()...)
	daoThis.insertParseBefore = append(daoThis.insertParseBefore, daoField.insertParseBefore.GetData()...)
	daoThis.insertParse = append(daoThis.insertParse, daoField.insertParse.GetData()...)
	daoThis.insertHook = append(daoThis.insertHook, daoField.insertHook.GetData()...)
	daoThis.updateParse = append(daoThis.updateParse, daoField.updateParse.GetData()...)
	daoThis.updateHookBefore = append(daoThis.updateHookBefore, daoField.updateHookBefore.GetData()...)
	daoThis.updateHookAfter = append(daoThis.updateHookAfter, daoField.updateHookAfter.GetData()...)
	daoThis.deleteHookBefore = append(daoThis.deleteHookBefore, daoField.deleteHookBefore.GetData()...)
	daoThis.deleteHookAfter = append(daoThis.deleteHookAfter, daoField.deleteHookAfter.GetData()...)
	// daoThis.groupParse = append(daoThis.groupParse, daoField.groupParse.GetData()...)
	daoThis.orderParse = append(daoThis.orderParse, daoField.orderParse.GetData()...)
	daoThis.joinParse = append(daoThis.joinParse, daoField.joinParse.GetData()...)
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
	daoThis.deleteHookBefore = append(daoThis.deleteHookBefore, daoOther.deleteHookBefore...)
	daoThis.deleteHookAfter = append(daoThis.deleteHookAfter, daoOther.deleteHookAfter...)
	daoThis.deleteHookOtherRel = append(daoThis.deleteHookOtherRel, daoOther.deleteHookOtherRel...)
	daoThis.groupParse = append(daoThis.groupParse, daoOther.groupParse...)
	daoThis.orderParse = append(daoThis.orderParse, daoOther.orderParse...)
	daoThis.joinParse = append(daoThis.joinParse, daoOther.joinParse...)
}

func (daoThis *myGenDao) Unique() {
	daoThis.importDao = garray.NewStrArrayFrom(daoThis.importDao).Unique().Slice()
	// daoThis.joinParse = garray.NewStrArrayFrom(daoThis.joinParse).Unique().Slice()
}

// dao生成
func genDao(tpl *myGenTpl) {
	tpl.gfGenDao(true) //dao文件生成
	saveFile := gfile.SelfDir() + `/internal/dao/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseSnake + `.go`
	tplDao := gfile.GetContents(saveFile)

	dao := getDaoIdAndLabel(tpl)
	for _, v := range tpl.FieldList {
		if v.FieldTypeName == internal.TypeNameDeleted { //存在软删除字段时，HookDelete内的事件需改成Update
			tplDao = gstr.Replace(tplDao, `Delete: func(ctx context.Context, in *gdb.HookDeleteInput) (result sql.Result, err error) {`, `Update: func(ctx context.Context, in *gdb.HookUpdateInput) (result sql.Result, err error) {`, 1)
			break
		}
	}
	for _, v := range tpl.FieldListOfDefault {
		dao.Add(getDaoField(tpl, v))
	}
	for _, v := range tpl.FieldListOfAfter1 {
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
	for _, v := range tpl.FieldListOfAfter2 {
		dao.Add(getDaoField(tpl, v))
	}
	for _, v := range tpl.Handle.OtherRelTableList {
		v.tpl.gfGenDao(false)
		dao.Merge(getDaoOtherRel(v))
	}
	dao.Unique()

	if len(dao.importDao) > 0 {
		importDaoPoint := `"api/internal/dao/` + tpl.ModuleDirCaseKebab + `/internal"`
		tplDao = gstr.Replace(tplDao, importDaoPoint, importDaoPoint+gstr.Join(append([]string{``}, dao.importDao...), `
	`), 1)
	}
	tplDao = gstr.Replace(tplDao, `"github.com/gogf/gf/v2/text/gstr"`, `"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"`, 1)

	if dao.idParse != `` {
		idParsePoint := `daoModel.DbTable + ` + "`.`" + ` + reflect.ValueOf(*daoThis.Columns()).Field(0).String()`
		tplDao = gstr.Replace(tplDao, idParsePoint, dao.idParse, 1)
	}
	if dao.labelParse != `` {
		labelParsePoint := `daoModel.DbTable + ` + "`.`" + ` + reflect.ValueOf(*daoThis.Columns()).Field(1).String()`
		tplDao = gstr.Replace(tplDao, labelParsePoint, dao.labelParse, 1)
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
			daoModel.AfterField[v] = struct{}{} */`
		tplDao = gstr.Replace(tplDao, fieldParsePoint, fieldParsePoint+gstr.Join(append([]string{``}, dao.fieldParse...), `
			`), 1)
	}
	if len(dao.fieldHook) > 0 {
		fieldHookPoint := `default:
			if v == struct{}{} {
				record[k] = gvar.New(nil)
			} else {
				record[k] = gvar.New(v)
			}`
		tplDao = gstr.Replace(tplDao, fieldHookPoint, gstr.Join(append(dao.fieldHook, ``), `
		`)+fieldHookPoint, 1)
	}

	// 解析insert
	if len(dao.insertParseBefore) > 0 {
		insertParseBeforePoint := `for k, v := range insert {`
		tplDao = gstr.Replace(tplDao, insertParseBeforePoint, gstr.Join(append(dao.insertParseBefore, ``), `
		`)+insertParseBeforePoint, 1)
	}
	if len(dao.insertParse) > 0 {
		insertParsePoint := `default:
				if daoThis.Contains(k) {
					daoModel.SaveData[k] = v
				}
			}
		}
		m = m.Data(daoModel.SaveData)
		if len(daoModel.AfterInsert) > 0 {`
		tplDao = gstr.Replace(tplDao, insertParsePoint, gstr.Join(append(dao.insertParse, ``), `
			`)+insertParsePoint, 1)
	}
	if len(dao.insertHook) > 0 {
		insertHookPoint := `// id, _ := result.LastInsertId()

			/* for k, v := range daoModel.AfterInsert {
				switch k {
				case ` + "`xxxx`" + `:
					daoModel.CloneNew().FilterPri(id).HookUpdateOne(k, v).Update()
				}
			} */`
		insertHookReplaceOfId := `id, _ := result.LastInsertId()`
		if tpl.Handle.Id.IsPrimary && len(tpl.Handle.Id.List) == 1 && !tpl.Handle.Id.List[0].IsAutoInc {
			insertHookReplaceOfId = `id := daoModel.IdArr[0]`
		}
		tplDao = gstr.Replace(tplDao, insertHookPoint, insertHookReplaceOfId+`

			for k, v := range daoModel.AfterInsert {
				switch k {`+gstr.Join(append([]string{``}, dao.insertHook...), `
				`)+`
				}
			}`, 1)
	}

	// 解析update
	if len(dao.updateParse) > 0 {
		updateParsePoint := `default:
				if daoThis.Contains(k) {
					daoModel.SaveData[k] = v
				}
			}
		}
		m = m.Data(daoModel.SaveData)
		if len(daoModel.AfterUpdate) > 0 {`
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
						daoModel.CloneNew().FilterPri(id).HookUpdateOne(k, v).Update()
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
	if len(dao.deleteHookBefore) > 0 {
		deleteHookBeforePoint := `//有软删除字段时需改成Update事件`
		tplDao = gstr.Replace(tplDao, deleteHookBeforePoint, deleteHookBeforePoint+gstr.Join(append([]string{``}, dao.deleteHookBefore...), `
			`)+`
			
`, 1)
	}
	if len(dao.deleteHookAfter) > 0 || len(dao.deleteHookOtherRel) > 0 {
		deleteHookPoint := `/* row, _ := result.RowsAffected()
			if row == 0 {
				return
			} */

			return`
		deleteHookPointReplace := deleteHookPoint
		if len(dao.deleteHookAfter) > 0 {
			deleteHookPointReplace = `row, _ := result.RowsAffected()
			if row == 0 {
				return
			}

			` + gstr.Join(append(dao.deleteHookAfter, ``), `
			`) + `return`
		}
		if len(dao.deleteHookOtherRel) > 0 {
			deleteHookPointReplace = gstr.Replace(deleteHookPointReplace, `
			return`, gstr.Join(append([]string{``, `/* // 对并发有要求时，可使用以下代码解决情形1。并发说明请参考：api/internal/dao/auth/scene.go中HookDelete方法内的注释`}, dao.deleteHookOtherRel...), `
			`)+` */
			return`, 1)
		}
		tplDao = gstr.Replace(tplDao, deleteHookPoint, deleteHookPointReplace, 1)
	}

	// 解析order
	if len(dao.groupParse) > 0 {
		groupParsePoint := `default:
				if daoThis.Contains(v) {
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
				if daoThis.Contains(k) {
					m = m.Order(daoModel.DbTable + ` + "`.`" + ` + v)
				} else {
					m = m.Order(v)
				}`
		tplDao = gstr.Replace(tplDao, orderParsePoint, gstr.Join(append(dao.orderParse, ``), `
			`)+orderParsePoint, 1)
	}

	// 解析join
	if tpl.Handle.Id.IsPrimary && len(tpl.Handle.Id.List) == 1 {
		dao.joinParse = append(dao.joinParse, `default:
			m = m.LeftJoin(joinTable, joinTable+`+"`.`"+`+daoThis.Columns().`+tpl.Handle.Id.List[0].FieldCaseCamel+`+`+"` = `"+`+daoModel.DbTable+`+"`.`"+`+daoThis.Columns().`+tpl.Handle.Id.List[0].FieldCaseCamel+`)`)
	}
	if len(dao.joinParse) > 0 {
		joinParsePoint := `/* case Xxxx.ParseDbTable(m.GetCtx()):
		m = m.LeftJoin(joinTable, joinTable+` + "`.`" + `+Xxxx.Columns().XxxxId+` + "` = `" + `+daoModel.DbTable+` + "`.`" + `+daoThis.Columns().XxxxId)
		// m = m.LeftJoin(Xxxx.ParseDbTable(m.GetCtx())+` + "` AS `" + `+joinTable, joinTable+` + "`.`" + `+Xxxx.Columns().XxxxId+` + "` = `" + `+daoModel.DbTable+` + "`.`" + `+daoThis.Columns().XxxxId) */`
		tplDao = gstr.Replace(tplDao, joinParsePoint, joinParsePoint+gstr.Join(append([]string{``}, dao.joinParse...), `
		`), 1)
	}

	utils.FilePutFormat(saveFile, []byte(tplDao)...)
}

func getDaoIdAndLabel(tpl *myGenTpl) (dao myGenDao) {
	if len(tpl.Handle.Id.List) == 1 {
		dao.idParse = `daoModel.DbTable + ` + "`.`" + ` + daoThis.Columns().` + tpl.Handle.Id.List[0].FieldCaseCamel
		dao.filterParse = append(dao.filterParse, `case `+"`id`, `"+internal.GetStrByFieldStyle(tpl.FieldStyle, `id_arr`)+"`"+`:
				m = m.Where(daoModel.DbTable+`+"`.`"+`+daoThis.Columns().`+tpl.Handle.Id.List[0].FieldCaseCamel+`, v)`)
		dao.filterParse = append(dao.filterParse, `case `+"`"+internal.GetStrByFieldStyle(tpl.FieldStyle, `exc_id`)+"`, `"+internal.GetStrByFieldStyle(tpl.FieldStyle, `exc_id_arr`)+"`"+`:
				if gvar.New(v).IsSlice() {
					m = m.WhereNotIn(daoModel.DbTable+`+"`.`"+`+daoThis.Columns().`+tpl.Handle.Id.List[0].FieldCaseCamel+`, v)
				} else {
					m = m.WhereNot(daoModel.DbTable+`+"`.`"+`+daoThis.Columns().`+tpl.Handle.Id.List[0].FieldCaseCamel+`, v)
				}`)
		if !tpl.Handle.Id.List[0].IsAutoInc {
			if tpl.Handle.Id.IsPrimary {
				dao.insertParse = append(dao.insertParse, `case `+"`id`, daoThis.Columns()."+tpl.Handle.Id.List[0].FieldCaseCamel+`:
					daoModel.SaveData[daoThis.Columns().`+tpl.Handle.Id.List[0].FieldCaseCamel+`] = v
					daoModel.IdArr = []*gvar.Var{gvar.New(v)}`)
			} else {
				dao.insertParse = append(dao.insertParse, `case `+"`id`"+`:
					daoModel.SaveData[daoThis.Columns().`+tpl.Handle.Id.List[0].FieldCaseCamel+`] = v`)
			}
			dao.updateParse = append(dao.updateParse, `case `+"`id`"+`:
					daoModel.SaveData[daoThis.Columns().`+tpl.Handle.Id.List[0].FieldCaseCamel+`] = v`)
		}
		dao.groupParse = append(dao.groupParse, `case `+"`id`"+`:
				m = m.Group(daoModel.DbTable + `+"`.`"+` + daoThis.Columns().`+tpl.Handle.Id.List[0].FieldCaseCamel+`)`)
		dao.orderParse = append(dao.orderParse, `case `+"`id`"+`:
				m = m.Order(daoModel.DbTable + `+"`.`"+` + gstr.Replace(v, k, daoThis.Columns().`+tpl.Handle.Id.List[0].FieldCaseCamel+`, 1))`)
	} else {
		concatStr := `|`
		idArrStrArr := []string{}
		for _, v := range tpl.Handle.Id.List {
			idArrStrArr = append(idArrStrArr, `daoModel.DbTable+`+"`.`"+`+daoThis.Columns().`+v.FieldCaseCamel)
		}

		dao.idParse = `fmt.Sprintf(` + "`" + `CONCAT_WS( '` + concatStr + `'` + gstr.Repeat(`, COALESCE( %s, '' )`, len(tpl.Handle.Id.List)) + ` )` + "`" + `, ` + gstr.Join(idArrStrArr, `, `) + `)`
		dao.filterParse = append(dao.filterParse, `case `+"`id`, `"+internal.GetStrByFieldStyle(tpl.FieldStyle, `id_arr`)+"`"+`:
				idArr := []string{gconv.String(v)}
				if gvar.New(v).IsSlice() {
					idArr = gconv.Strings(v)
				}
				inStrArr := make([]string, len(idArr))
				for index, id := range idArr {
					inStrArr[index] = `+"`('`+gstr.Replace(id, `"+concatStr+"`, `', '`)+`')`"+`
				}
				operator := `+"`IN`"+`
				if len(inStrArr) == 1 {
					operator = `+"`=`"+`
				}
				m = m.Where(fmt.Sprintf(`+"`(%s, %s) %s (%s)`, "+gstr.Join(idArrStrArr, `, `)+", operator, gstr.Join(inStrArr, `, `)))")
		dao.filterParse = append(dao.filterParse, `case `+"`"+internal.GetStrByFieldStyle(tpl.FieldStyle, `exc_id`)+"`, `"+internal.GetStrByFieldStyle(tpl.FieldStyle, `exc_id_arr`)+"`"+`:
				idArr := []string{gconv.String(v)}
				if gvar.New(v).IsSlice() {
					idArr = gconv.Strings(v)
				}
				inStrArr := make([]string, len(idArr))
				for index, id := range idArr {
					inStrArr[index] = `+"`('`+gstr.Replace(id, `"+concatStr+"`, `', '`)+`')`"+`
				}
				operator := `+"`NOT IN`"+`
				if len(inStrArr) == 1 {
					operator = `+"`!=`"+`
				}
				m = m.Where(fmt.Sprintf(`+"`(%s, %s) %s (%s)`, "+gstr.Join(idArrStrArr, `, `)+", operator, gstr.Join(inStrArr, `, `)))")
		dao.groupParse = append(dao.groupParse, `case `+"`id`"+`:
				m = m.Group(`+gstr.Join(idArrStrArr, `, `)+`)`)
		dao.orderParse = append(dao.orderParse, `case `+"`id`"+`:
				suffix := gstr.TrimLeftStr(kArr[0], k, 1)
				m = m.Order(`+gstr.TrimRightStr(gstr.Join(append(idArrStrArr, ``), `+suffix, `), `, `)+`)
				remain := gstr.TrimLeftStr(gstr.TrimLeftStr(v, k+suffix, 1), `+"`,`"+`, 1)
				if remain != `+"``"+` {
					m = m.Order(remain)
				}`)
	}
	dao.fieldParse = append(dao.fieldParse, `case `+"`id`"+`:
				m = m.Fields(daoThis.ParseId(daoModel) + `+"` AS `"+` + v)`)

	dao.labelParse = `daoModel.DbTable + ` + "`.`" + ` + daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.LabelList[0])
	filterParseStr := `case ` + "`label`" + `:
				m = m.WhereLike(daoModel.DbTable+` + "`.`" + `+daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.LabelList[0]) + `, ` + "`%`" + `+gconv.String(v)+` + "`%`" + `)`
	labelListLen := len(tpl.Handle.LabelList)
	if labelListLen > 1 {
		labelParseStrArr := []string{`daoModel.DbTable+` + "`.`" + `+daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.LabelList[labelListLen-1])}
		parseFilterStr := "WhereOrLike(daoModel.DbTable+`.`+daoThis.Columns()." + gstr.CaseCamel(tpl.Handle.LabelList[labelListLen-1]) + ", `%`+gconv.String(v)+`%`)"
		for i := labelListLen - 2; i >= 0; i-- {
			labelParseStrArr = append([]string{`daoModel.DbTable+` + "`.`" + `+daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.LabelList[i])}, labelParseStrArr...)
			if i == 0 {
				parseFilterStr = "WhereLike(daoModel.DbTable+`.`+daoThis.Columns()." + gstr.CaseCamel(tpl.Handle.LabelList[i]) + ", `%`+gconv.String(v)+`%`)." + parseFilterStr
			} else {
				parseFilterStr = "WhereOrLike(daoModel.DbTable+`.`+daoThis.Columns()." + gstr.CaseCamel(tpl.Handle.LabelList[i]) + ", `%`+gconv.String(v)+`%`)." + parseFilterStr
			}
		}
		dao.labelParse = `fmt.Sprintf(` + "`" + `COALESCE( ` + gstr.TrimLeftStr(gstr.Repeat(`, NULLIF( %s, '' )`, labelListLen), `, `, 1) + ` )` + "`" + `, ` + gstr.Join(labelParseStrArr, `, `) + `)`

		filterParseStr = `case ` + "`label`" + `:
				m = m.Where(m.Builder().` + parseFilterStr + `)`
	}
	dao.fieldParse = append(dao.fieldParse, `case `+"`label`"+`:
				m = m.Fields(daoThis.ParseLabel(daoModel) + `+"` AS `"+` + v)`)
	dao.filterParse = append(dao.filterParse, filterParseStr)
	return
}

func getDaoField(tpl *myGenTpl, v myGenField) (daoField myGenDaoField) {
	daoPath := `daoThis`
	daoTable := `daoModel.DbTable`

	/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 开始--------*/
	switch v.FieldType {
	case internal.TypeInt, internal.TypeIntU: // `int等类型` // `int等类型（unsigned）`
	case internal.TypeFloat, internal.TypeFloatU: // `float等类型`  // `float等类型（unsigned）`
	case internal.TypeVarchar, internal.TypeChar: // `varchar类型`	// `char类型`
		if v.IsUnique || gconv.Uint(v.FieldLimitStr) <= internal.ConfigMaxLenOfStrFilter {
			daoField.filterParse.Method = internal.ReturnType
		}
		if v.IsNull && v.IsUnique {
			daoField.insertParse.Method = internal.ReturnType
			daoField.insertParse.DataType = append(daoField.insertParse.DataType, `case `+daoPath+`.Columns().`+v.FieldCaseCamel+`:
				if gconv.String(v) == `+"``"+` {
					v = nil
				}
				daoModel.SaveData[k] = v`)

			daoField.updateParse.Method = internal.ReturnType
			daoField.updateParse.DataType = append(daoField.updateParse.DataType, `case `+daoPath+`.Columns().`+v.FieldCaseCamel+`:
				if gconv.String(v) == `+"``"+` {
					v = nil
				}
				daoModel.SaveData[k] = v`)
		}
	case internal.TypeText: // `text类型`
	case internal.TypeJson: // `json类型`
		if v.IsNull {
			daoField.insertParse.Method = internal.ReturnType
			daoField.insertParse.DataType = append(daoField.insertParse.DataType, `case `+daoPath+`.Columns().`+v.FieldCaseCamel+`:
				if gconv.String(v) == `+"``"+` {
					v = nil
				}
				daoModel.SaveData[k] = v`)

			daoField.updateParse.Method = internal.ReturnType
			daoField.updateParse.DataType = append(daoField.updateParse.DataType, `case `+daoPath+`.Columns().`+v.FieldCaseCamel+`:
				if gconv.String(v) == `+"``"+` {
					v = nil
				}
				daoModel.SaveData[k] = v`)
		}
	case internal.TypeDatetime, internal.TypeTimestamp, internal.TypeDate, internal.TypeTime: // `datetime类型`	// `timestamp类型`	 // `date类型`	// `time类型`
		if v.IsNull {
			daoField.insertParse.Method = internal.ReturnType
			daoField.insertParse.DataType = append(daoField.insertParse.DataType, `case `+daoPath+`.Columns().`+v.FieldCaseCamel+`:
				if gconv.String(v) == `+"``"+` {
					v = nil
				}
				daoModel.SaveData[k] = v`)

			daoField.updateParse.Method = internal.ReturnType
			daoField.updateParse.DataType = append(daoField.updateParse.DataType, `case `+daoPath+`.Columns().`+v.FieldCaseCamel+`:
				if gconv.String(v) == `+"``"+` {
					v = nil
				}
				daoModel.SaveData[k] = v`)
		}
		if slices.Contains([]internal.MyGenFieldType{internal.TypeDate /* , internal.TypeTime */}, v.FieldType) {
			daoField.filterParse.Method = internal.ReturnType
			daoField.orderParse.Method = internal.ReturnType
			daoField.orderParse.DataType = append(daoField.orderParse.DataType, `case `+daoPath+`.Columns().`+v.FieldCaseCamel+`:
				m = m.Order(`+daoTable+` + `+"`.`"+` + v)
				`+getAddOrder(tpl.Handle.Id.List, tpl.Handle.DefSort.Field, tpl.Handle.DefSort.Order))
		}
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
	case internal.TypeNamePid: // pid，且与主键类型相同时（才）有效；	类型：int等类型或varchar或char；
		daoField.filterParse.Method = internal.ReturnTypeName

		daoField.fieldParse.Method = internal.ReturnTypeName
		daoField.fieldParse.DataTypeName = append(daoField.fieldParse.DataTypeName, `case `+"`"+internal.GetStrByFieldStyle(tpl.FieldStyle, tpl.Handle.LabelList[0], `p`)+"`"+`:
				tableP := `+"`p_`"+` + `+daoTable+`
				m = m.Fields(tableP + `+"`.`"+` + `+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.LabelList[0])+` + `+"` AS `"+` + v)
				m = m.Handler(`+daoPath+`.ParseJoin(tableP, daoModel))`)
		if tpl.Handle.Pid.IsLeaf == `` {
			daoField.fieldParse.DataTypeName = append(daoField.fieldParse.DataTypeName, `case `+"`"+internal.GetStrByFieldStyle(tpl.FieldStyle, `is_leaf`)+"`"+`:
				m = m.Fields(`+daoTable+` + `+"`.`"+` + `+daoPath+`.Columns().`+tpl.Handle.Id.List[0].FieldCaseCamel+`)
				daoModel.AfterField[v] = struct{}{}`)

			daoField.fieldHook.Method = internal.ReturnTypeName
			daoField.fieldHook.DataTypeName = append(daoField.fieldHook.DataTypeName, `case `+"`"+internal.GetStrByFieldStyle(tpl.FieldStyle, `is_leaf`)+"`"+`:
			isLeaf := 0
			if count, _ := daoModel.CloneNew().Filter(`+daoPath+`.Columns().`+v.FieldCaseCamel+`, record[`+daoPath+`.Columns().`+tpl.Handle.Id.List[0].FieldCaseCamel+`]).Count(); count > 0 {
				isLeaf = 1
			}
			record[k] = gvar.New(isLeaf)`)
		}
		daoField.fieldParse.DataTypeName = append(daoField.fieldParse.DataTypeName, `case `+"`tree`"+`:
				m = m.Fields(`+daoTable+` + `+"`.`"+` + `+daoPath+`.Columns().`+tpl.Handle.Id.List[0].FieldCaseCamel+`)
				m = m.Fields(`+daoTable+` + `+"`.`"+` + `+daoPath+`.Columns().`+v.FieldCaseCamel+`)
				m = m.Handler(`+daoPath+`.ParseOrder([]string{`+"`tree`"+`}, daoModel))`)

		orderParseStr := `case ` + "`tree`" + `:`
		if tpl.Handle.Pid.Level != `` {
			orderParseStr += `
				m = m.OrderAsc(` + daoTable + ` + ` + "`.`" + ` + ` + daoPath + `.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Level) + `)`
		} else {
			orderParseStr += `
				m = m.OrderAsc(` + daoTable + ` + ` + "`.`" + ` + ` + daoPath + `.Columns().` + v.FieldCaseCamel + `)`
		}
		for _, sort := range tpl.Handle.Pid.Sort {
			orderParseStr += `
				m = m.OrderDesc(` + daoTable + ` + ` + "`.`" + ` + ` + daoPath + `.Columns().` + gstr.CaseCamel(sort) + `)`
		}
		orderParseStr += `
				` + getAddOrder(tpl.Handle.Id.List, tpl.Handle.DefSort.Field, `ASC`)
		daoField.orderParse.Method = internal.ReturnTypeName
		daoField.orderParse.DataTypeName = append(daoField.orderParse.DataTypeName, orderParseStr)

		daoField.joinParse.Method = internal.ReturnTypeName
		daoField.joinParse.DataTypeName = append(daoField.joinParse.DataTypeName, `case `+"`p_`"+` + `+daoTable+`:
			m = m.LeftJoin(`+daoTable+`+`+"` AS `"+`+joinTable, joinTable+`+"`.`"+`+`+daoPath+`.Columns().`+tpl.Handle.Id.List[0].FieldCaseCamel+`+`+"` = `"+`+`+daoTable+`+`+"`.`"+`+`+daoPath+`.Columns().`+v.FieldCaseCamel+`)`)

		if tpl.Handle.Pid.IsCoexist {
			daoField.insertParseBefore.Method = internal.ReturnTypeName
			daoField.insertParseBefore.DataTypeName = append(daoField.insertParseBefore.DataTypeName, `if _, ok := insert[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Pid)+`]; !ok {
			insert[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Pid)+`] = `+tpl.Handle.Pid.Tpl.PidDefVal+`
		}`)

			daoField.filterParse.Method = internal.ReturnTypeName
			daoField.insertParse.Method = internal.ReturnTypeName
			daoField.insertHook.Method = internal.ReturnTypeName
			daoField.updateParse.Method = internal.ReturnTypeName
			daoField.updateHookAfter.Method = internal.ReturnTypeName

			selfUpdateStr := internal.GetStrByFieldStyle(tpl.FieldStyle, `self_update`)
			pIdPathStr := internal.GetStrByFieldStyle(tpl.FieldStyle, `p_id_path`)
			pNamePathStr := internal.GetStrByFieldStyle(tpl.FieldStyle, `p_name_path`)
			pLevelStr := internal.GetStrByFieldStyle(tpl.FieldStyle, `p_level`)

			childUpdateStr := internal.GetStrByFieldStyle(tpl.FieldStyle, `child_update_list`)
			pIdPathOfOldStr := internal.GetStrByFieldStyle(tpl.FieldStyle, `p_id_path_of_old`)
			pIdPathOfNewStr := internal.GetStrByFieldStyle(tpl.FieldStyle, `p_id_path_of_new`)
			pNamePathOfOldStr := internal.GetStrByFieldStyle(tpl.FieldStyle, `p_name_path_of_old`)
			pNamePathOfNewStr := internal.GetStrByFieldStyle(tpl.FieldStyle, `p_name_path_of_new`)
			pLevelOfOldStr := internal.GetStrByFieldStyle(tpl.FieldStyle, `p_level_of_old`)
			pLevelOfNewStr := internal.GetStrByFieldStyle(tpl.FieldStyle, `p_level_of_new`)
			childIdPathStr := internal.GetStrByFieldStyle(tpl.FieldStyle, `child_id_path`)
			childNamePathStr := internal.GetStrByFieldStyle(tpl.FieldStyle, `child_name_path`)
			childLevelStr := internal.GetStrByFieldStyle(tpl.FieldStyle, `child_level`)

			daoField.filterParse.DataTypeName = append(daoField.filterParse.DataTypeName, `case `+"`"+pIdPathOfOldStr+"`"+`: //父级ID路径（旧）
				m = m.WhereLike(`+daoTable+`+`+"`.`"+`+`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.IdPath)+`, gconv.String(v)+`+"`-%`"+`)`)

			afterInsertStrArr := []string{}
			afterInsertStrArrOfEmpty := []string{}
			afterInsertMapKeyArr := []string{"`" + pIdPathStr + "`" + `: pInfo[` + daoPath + `.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.IdPath) + `],`}
			afterInsertMapKeyArrOfEmpty := []string{"`" + pIdPathStr + "`" + `: ` + tpl.Handle.Pid.Tpl.PIdPathDefVal + `,`}
			insertHookMapKeyArr := []string{daoPath + `.Columns().IdPath: gconv.String(val[` + "`" + pIdPathStr + "`" + `]) + ` + "`-`" + ` + gconv.String(id),`}

			afterUpdateStrArr1 := []string{gstr.CaseCamelLower(pIdPathStr) + ` := ` + tpl.Handle.Pid.Tpl.PIdPathDefVal}
			afterUpdateStrArr2 := []string{gstr.CaseCamelLower(pIdPathStr) + ` = pInfo[` + daoPath + `.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.IdPath) + `].String()`}
			afterUpdateStrArr3 := []string{`daoModel.SaveData[` + daoPath + `.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.IdPath) + `] = gdb.Raw(fmt.Sprintf(` + "`" + `CONCAT( '%s-', %s )` + "`" + `, ` + gstr.CaseCamelLower(pIdPathStr) + `, ` + daoPath + `.Columns().` + tpl.Handle.Id.List[0].FieldCaseCamel + `))`}
			afterUpdateStrArr4 := []string{}
			afterUpdateStrArr5 := []string{}
			afterUpdateStrArr6 := []string{}
			childUpdateNamePathStr := ``
			childUpdateMapKeyArr := []string{
				"`" + pIdPathOfOldStr + "`" + `: oldInfo[` + daoPath + `.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.IdPath) + `],`,
				"`" + childIdPathStr + "`" + `: map[string]any{
								` + "`" + pIdPathOfOldStr + "`" + `: oldInfo[` + daoPath + `.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.IdPath) + `],
								` + "`" + pIdPathOfNewStr + "`" + `: ` + gstr.CaseCamelLower(pIdPathStr) + ` + ` + "`-`" + ` + oldInfo[` + daoPath + `.Columns().` + tpl.Handle.Id.List[0].FieldCaseCamel + `].String(),
							},`,
			}
			updateParseArr := []string{`case ` + "`" + childIdPathStr + "`" + `: //更新所有子孙级的ID路径。参数：map[string]any{` + "`" + pIdPathOfOldStr + "`" + ": `父级ID路径（旧）`" + `, ` + "`" + pIdPathOfNewStr + "`" + ": `父级ID路径（新）`" + `}
				val := gconv.Map(v)
				daoModel.SaveData[` + daoPath + `.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.IdPath) + `] = gdb.Raw(fmt.Sprintf(` + "`" + `REPLACE( %s, '%s', '%s' )` + "`" + `, ` + daoPath + `.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.IdPath) + `, gconv.String(val[` + "`" + pIdPathOfOldStr + "`" + `]), gconv.String(val[` + "`" + pIdPathOfNewStr + "`" + `])))`}
			if tpl.Handle.Pid.NamePath != `` {
				afterInsertMapKeyArr = append(afterInsertMapKeyArr, "`"+pNamePathStr+"`"+`: pInfo[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.NamePath)+`],`, "`name`"+`: insert[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.LabelList[0])+`],`)
				afterInsertMapKeyArrOfEmpty = append(afterInsertMapKeyArrOfEmpty, "`"+pNamePathStr+"`"+": ``"+`,`, "`name`"+`: insert[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.LabelList[0])+`],`)
				insertHookMapKeyArr = append(insertHookMapKeyArr, daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.NamePath)+`: gconv.String(val[`+"`"+pNamePathStr+"`"+`]) + `+"`-`"+` + gconv.String(val[`+"`name`"+`]),`)

				afterUpdateStrArr1 = append(afterUpdateStrArr1, gstr.CaseCamelLower(pNamePathStr)+` := `+"``")
				afterUpdateStrArr2 = append(afterUpdateStrArr2, gstr.CaseCamelLower(pNamePathStr)+` = pInfo[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.NamePath)+`].String()`)
				afterUpdateStrArr3 = append(afterUpdateStrArr3, `daoModel.SaveData[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.NamePath)+`] = gdb.Raw(fmt.Sprintf(`+"`"+`CONCAT( '%s-', %s )`+"`"+`, `+gstr.CaseCamelLower(pNamePathStr)+`, `+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.LabelList[0])+`))`)
				afterUpdateStrArr3 = append(afterUpdateStrArr3, `_, ok`+gstr.CaseCamel(tpl.Handle.LabelList[0])+` := update[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.LabelList[0])+`]
				if ok`+gstr.CaseCamel(tpl.Handle.LabelList[0])+` {
					daoModel.SaveData[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.NamePath)+`] = gdb.Raw(fmt.Sprintf(`+"`"+`CONCAT('%s-', '%s')`+"`"+`, `+gstr.CaseCamelLower(pNamePathStr)+`, gconv.String(update[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.LabelList[0])+`])))
				}`)
				childUpdateMapKeyArr = append(childUpdateMapKeyArr, "`"+childNamePathStr+"`"+`: map[string]any{
								`+"`"+pNamePathOfOldStr+"`"+`: oldInfo[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.NamePath)+`],
								`+"`"+pNamePathOfNewStr+"`"+`: `+gstr.CaseCamelLower(pNamePathStr)+` + `+"`-`"+` + oldInfo[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.LabelList[0])+`].String(),
							},`)
				updateParseArr = append(updateParseArr, `case `+"`"+childNamePathStr+"`"+`: //更新所有子孙级的名称路径。参数：map[string]any{`+"`"+pNamePathOfOldStr+"`"+": `父级名称路径（旧）`"+`, `+"`"+pNamePathOfNewStr+"`"+": `父级名称路径（新）`"+`}
				val := gconv.Map(v)
				daoModel.SaveData[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.NamePath)+`] = gdb.Raw(fmt.Sprintf(`+"`"+`REGEXP_REPLACE( %s, CONCAT( '^', '%s' ), '%s' )`+"`"+`, `+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.NamePath)+`, gconv.String(val[`+"`"+pNamePathOfOldStr+"`"+`]), gconv.String(val[`+"`"+pNamePathOfNewStr+"`"+`])))`)
			}
			if tpl.Handle.Pid.Level != `` {
				afterInsertMapKeyArr = append(afterInsertMapKeyArr, "`"+pLevelStr+"`"+`:   pInfo[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Level)+`],`)
				afterInsertMapKeyArrOfEmpty = append(afterInsertMapKeyArrOfEmpty, "`"+pLevelStr+"`"+`:   0,`)
				insertHookMapKeyArr = append(insertHookMapKeyArr, daoPath+`.Columns().Level:  gconv.Uint(val[`+"`"+pLevelStr+"`"+`]) + 1,`)

				afterUpdateStrArr1 = append(afterUpdateStrArr1, `var pLevel uint = 0`)
				afterUpdateStrArr2 = append(afterUpdateStrArr2, `pLevel = pInfo[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Level)+`].Uint()`)
				afterUpdateStrArr3 = append(afterUpdateStrArr3, `daoModel.SaveData[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Level)+`] = pLevel + 1`)
				childUpdateMapKeyArr = append(childUpdateMapKeyArr, "`"+childLevelStr+"`"+`: map[string]any{
								`+"`"+pLevelOfOldStr+"`"+`: oldInfo[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Level)+`],
								`+"`"+pLevelOfNewStr+"`"+`: pLevel + 1,
							},`)
				updateParseArr = append(updateParseArr, `case `+"`"+childLevelStr+"`"+`: //更新所有子孙级的层级。参数：map[string]any{`+"`"+pLevelOfOldStr+"`"+": `父级层级（旧）`"+`, `+"`"+pLevelOfNewStr+"`"+": `父级层级（新）`"+`}
				val := gconv.Map(v)
				pLevelOfOld := gconv.Uint(val[`+"`"+pLevelOfOldStr+"`"+`])
				pLevelOfNew := gconv.Uint(val[`+"`"+pLevelOfNewStr+"`"+`])
				daoModel.SaveData[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Level)+`] = gdb.Raw(`+daoTable+` + `+"`.`"+` + `+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Level)+` + `+"` + `"+` + gconv.String(pLevelOfNew-pLevelOfOld))
				if pLevelOfNew < pLevelOfOld {
					daoModel.SaveData[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Level)+`] = gdb.Raw(`+daoTable+` + `+"`.`"+` + `+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Level)+` + `+"` - `"+` + gconv.String(pLevelOfOld-pLevelOfNew))
				}`)
			}

			afterInsertStrArr = append(afterInsertStrArr, `daoModel.AfterInsert[`+"`"+selfUpdateStr+"`"+`] = map[string]any{`+gstr.Join(append([]string{``}, afterInsertMapKeyArr...), `
						`)+`
					}`)
			afterInsertStrArrOfEmpty = append(afterInsertStrArrOfEmpty, `daoModel.AfterInsert[`+"`"+selfUpdateStr+"`"+`] = map[string]any{`+gstr.Join(append([]string{``}, afterInsertMapKeyArrOfEmpty...), `
						`)+`
					}`)
			daoField.insertHook.DataTypeName = append(daoField.insertHook.DataTypeName, `case `+"`"+selfUpdateStr+"`"+`: //更新自身的ID路径和层级。参数：map[string]any{`+"`"+pIdPathStr+"`"+": `父级ID路径`"+`, `+"`"+pNamePathStr+"`"+": `父级名称路径`"+`, `+"`name`: `当前名称`"+`, `+"`"+pLevelStr+"`"+": `父级层级`"+`}
					val := v.(map[string]any)
					daoModel.CloneNew().FilterPri(id).HookUpdate(map[string]any{`+gstr.Join(append([]string{``}, insertHookMapKeyArr...), `
						`)+`
					}).Update()`)

			afterUpdateStrArr4 = append(afterUpdateStrArr4, gstr.CaseCamelLower(childUpdateStr)+` := []map[string]any{} //更新所有子孙级的ID路径，名称路径和层级`)
			afterUpdateStrArr5 = append(afterUpdateStrArr5, gstr.CaseCamelLower(childUpdateStr)+` = append(`+gstr.CaseCamelLower(childUpdateStr)+`, map[string]any{`+gstr.Join(append([]string{``}, childUpdateMapKeyArr...), `
							`)+`
						})`)
			if tpl.Handle.Pid.NamePath != `` {
				afterUpdateStrArr5 = append(afterUpdateStrArr5, `if ok`+gstr.CaseCamel(tpl.Handle.LabelList[0])+` {
							`+gstr.CaseCamelLower(childUpdateStr)+`[len(`+gstr.CaseCamelLower(childUpdateStr)+`)-1][`+"`"+childNamePathStr+"`"+`].(map[string]any)[`+"`"+pNamePathOfNewStr+"`"+`] = `+gstr.CaseCamelLower(pNamePathStr)+` + `+"`-`"+` + gconv.String(update[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.LabelList[0])+`])
						}`)
				childUpdateNamePathStr = ` else if ok` + gstr.CaseCamel(tpl.Handle.LabelList[0]) + ` {
						if name := gconv.String(update[` + daoPath + `.Columns().` + gstr.CaseCamel(tpl.Handle.LabelList[0]) + `]); name != oldInfo[` + daoPath + `.Columns().` + gstr.CaseCamel(tpl.Handle.LabelList[0]) + `].String() {
							` + gstr.CaseCamelLower(childUpdateStr) + ` = append(` + gstr.CaseCamelLower(childUpdateStr) + `, map[string]any{
								` + "`" + pIdPathOfOldStr + "`" + `: oldInfo[` + daoPath + `.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.IdPath) + `],
								` + "`" + childNamePathStr + "`" + `: map[string]any{
									` + "`" + pNamePathOfOldStr + "`" + `: oldInfo[` + daoPath + `.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.NamePath) + `],
									` + "`" + pNamePathOfNewStr + "`" + `: ` + gstr.CaseCamelLower(pNamePathStr) + ` + ` + "`-`" + ` + name,
								},
							})
						}
					}`
			}
			afterUpdateStrArr6 = append(afterUpdateStrArr6, `if len(`+gstr.CaseCamelLower(childUpdateStr)+`) > 0 {
					daoModel.AfterUpdate[`+"`"+childUpdateStr+"`"+`] = `+gstr.CaseCamelLower(childUpdateStr)+`
				}`)

			daoField.updateHookAfter.DataTypeName = append(daoField.updateHookAfter.DataTypeName, `case `+"`"+childUpdateStr+"`"+`: //修改pid时，更新所有子孙级的ID路径，名称路径和层级。参数：[]map[string]any{`+"`"+pIdPathOfOldStr+"`"+": `父级ID路径（旧）`"+`, `+"`"+childIdPathStr+"`"+`: map[string]any{`+"`"+pIdPathOfOldStr+"`"+": `父级ID路径（旧）`"+`, `+"`"+pIdPathOfNewStr+"`"+": `父级ID路径（新）`"+`}, `+"`"+childNamePathStr+"`"+`: map[string]any{`+"`"+pNamePathOfOldStr+"`"+": `父级名称路径（旧）`"+`, `+"`"+pNamePathOfNewStr+"`"+": `父级名称路径（新）`"+`}, `+"`"+childLevelStr+"`"+`: map[string]any{`+"`"+pLevelOfOldStr+"`"+": `父级层级（旧）`"+`, `+"`"+pLevelOfNewStr+"`"+": `父级层级（新）`"+`}}
					val := v.([]map[string]any)
					for _, v1 := range val {
						pIdPathOfOld := gconv.String(v1[`+"`"+pIdPathOfOldStr+"`"+`])
						delete(v1, `+"`"+pIdPathOfOldStr+"`"+`)
						daoModel.CloneNew().Filter(`+"`"+pIdPathOfOldStr+"`"+`, pIdPathOfOld).HookUpdate(v1).Update()
					}`)

			if tpl.Handle.Pid.IsLeaf != `` {
				daoField.insertParseBefore.DataTypeName = append(daoField.insertParseBefore.DataTypeName, `if _, ok := insert[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.IsLeaf)+`]; !ok {
			insert[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.IsLeaf)+`] = 1
		}`)

				pIsLeafStr := internal.GetStrByFieldStyle(tpl.FieldStyle, `p_is_leaf`)
				pIsLeafCheckStr := internal.GetStrByFieldStyle(tpl.FieldStyle, `p_is_leaf_check`)
				pIsLeafCheckStrCaseCamelLower := gstr.CaseCamelLower(pIsLeafCheckStr)

				afterInsertStrArr = append(afterInsertStrArr, `if pInfo[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.IsLeaf)+`].Uint() == 1 {
						daoModel.AfterInsert[`+"`"+pIsLeafStr+"`"+`] = v
					}`)
				daoField.insertHook.DataTypeName = append(daoField.insertHook.DataTypeName, `case `+"`"+pIsLeafStr+"`"+`: //更新父级叶子。参数：父级ID
					daoModel.CloneNew().FilterPri(v).HookUpdateOne(`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.IsLeaf)+`, 0).Update()`)

				afterUpdateStrArr2 = append(afterUpdateStrArr2, `if pInfo[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.IsLeaf)+`].Uint() == 1 {
						daoModel.AfterUpdate[`+"`"+pIsLeafStr+"`"+`] = v
					}`)
				afterUpdateStrArr4 = append(afterUpdateStrArr4, pIsLeafCheckStrCaseCamelLower+` := []`+tpl.Handle.Pid.Tpl.PidType+`{} //更新原父级叶子`)
				afterUpdateStrArr5 = append(afterUpdateStrArr5, `if `+pIsLeafCheckStrCaseCamelLower+`Tmp := gconv.`+tpl.Handle.Pid.Tpl.PidGconvMethod+`s(gstr.Split(oldInfo[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.IdPath)+`].String(), `+"`-`"+`)); len(`+pIsLeafCheckStrCaseCamelLower+`Tmp) > 2 {
							`+pIsLeafCheckStrCaseCamelLower+` = append(`+pIsLeafCheckStrCaseCamelLower+`, `+pIsLeafCheckStrCaseCamelLower+`Tmp[len(`+pIsLeafCheckStrCaseCamelLower+`Tmp)-2])
						}`)
				afterUpdateStrArr6 = append(afterUpdateStrArr6, `if len(`+pIsLeafCheckStrCaseCamelLower+`) > 0 {
					daoModel.AfterUpdate[`+"`"+pIsLeafCheckStr+"`"+`] = `+pIsLeafCheckStrCaseCamelLower+`
				}`)
				daoField.updateHookAfter.DataTypeName = append(daoField.updateHookAfter.DataTypeName, `case `+"`"+pIsLeafStr+"`"+`: //更新父级叶子。参数：父级ID
					daoModel.CloneNew().FilterPri(v).HookUpdateOne(`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.IsLeaf)+`, 0).Update()`)
				daoField.updateHookAfter.DataTypeName = append(daoField.updateHookAfter.DataTypeName, `case `+"`"+pIsLeafCheckStr+"`"+`: //更新原父级叶子。参数：[]{父级ID,...}
					pidArr, _ := daoModel.CloneNew().Filter(`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Pid)+`, v).Distinct().Array(`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Pid)+`)
					if idArr := gset.NewFrom(v).Diff(gset.NewFrom(gconv.`+tpl.Handle.Pid.Tpl.PidGconvMethod+`s(pidArr))).Slice(); len(idArr) > 0 {
						daoModel.CloneNew().FilterPri(idArr).HookUpdateOne(`+daoPath+`.Columns().IsLeaf, 1).Update()
					}`)

				daoField.deleteHookBefore.Method = internal.ReturnTypeName
				daoField.deleteHookBefore.DataTypeName = append(daoField.deleteHookBefore.DataTypeName, pIsLeafCheckStrCaseCamelLower+` := []`+tpl.Handle.Pid.Tpl.PidType+`{} //更新原父级叶子
			idPathArr, _ := daoModel.CloneNew().FilterPri(daoModel.IdArr).ArrayStr(`+daoPath+`.Columns().IdPath)
			for _, idPath := range idPathArr {
				if `+pIsLeafCheckStrCaseCamelLower+`Tmp := gconv.`+tpl.Handle.Pid.Tpl.PidGconvMethod+`s(gstr.Split(idPath, `+"`-`"+`)); len(`+pIsLeafCheckStrCaseCamelLower+`Tmp) > 2 {
					`+pIsLeafCheckStrCaseCamelLower+` = append(`+pIsLeafCheckStrCaseCamelLower+`, `+pIsLeafCheckStrCaseCamelLower+`Tmp[len(`+pIsLeafCheckStrCaseCamelLower+`Tmp)-2])
				}
			}`)
				daoField.deleteHookAfter.Method = internal.ReturnTypeName
				daoField.deleteHookAfter.DataTypeName = append(daoField.deleteHookAfter.DataTypeName, `pidArr, _ := daoModel.CloneNew().Filter(`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Pid)+`, `+pIsLeafCheckStrCaseCamelLower+`).Distinct().Array(`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Pid)+`)
			if idArr := gset.NewFrom(`+pIsLeafCheckStrCaseCamelLower+`).Diff(gset.NewFrom(gconv.`+tpl.Handle.Pid.Tpl.PidGconvMethod+`s(pidArr))).Slice(); len(idArr) > 0 {
				daoModel.CloneNew().FilterPri(idArr).HookUpdateOne(`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.IsLeaf)+`, 1).Update()
			}`)
			}

			afterInsertStr := `case ` + daoPath + `.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Pid) + `:
				daoModel.SaveData[k] = v
				if gconv.` + tpl.Handle.Pid.Tpl.PidGconvMethod + `(v) ` + tpl.Handle.Pid.Tpl.PidJudge + ` {
					pInfo, _ := daoModel.CloneNew().FilterPri(v).One()` + gstr.Join(append([]string{``}, afterInsertStrArr...), `
					`) + `
				}`
			if len(afterInsertStrArrOfEmpty) > 0 {
				afterInsertStr += ` else {` + gstr.Join(append([]string{``}, afterInsertStrArrOfEmpty...), `
					`) + `
				}`
			}
			daoField.insertParse.DataTypeName = append(daoField.insertParse.DataTypeName, afterInsertStr)

			daoField.updateParse.DataTypeName = append(daoField.updateParse.DataTypeName, `case `+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Pid)+`:
				daoModel.SaveData[k] = v`+gstr.Join(append([]string{``}, afterUpdateStrArr1...), `
				`)+`
				if gconv.`+tpl.Handle.Pid.Tpl.PidGconvMethod+`(v) `+tpl.Handle.Pid.Tpl.PidJudge+` {
					pInfo, _ := daoModel.CloneNew().FilterPri(v).One()`+gstr.Join(append([]string{``}, afterUpdateStrArr2...), `
					`)+`
				}`+gstr.Join(append([]string{``}, afterUpdateStrArr3...), `
				`)+gstr.Join(append([]string{``}, afterUpdateStrArr4...), `
				`)+`
				oldList, _ := daoModel.CloneNew().FilterPri(daoModel.IdArr).All()
				for _, oldInfo := range oldList {
					if gconv.`+tpl.Handle.Pid.Tpl.PidGconvMethod+`(v) != oldInfo[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Pid)+`].`+tpl.Handle.Pid.Tpl.PidGconvMethod+`() {`+gstr.Join(append([]string{``}, afterUpdateStrArr5...), `
						`)+`
					}`+childUpdateNamePathStr+`
				}`+gstr.Join(append([]string{``}, afterUpdateStrArr6...), `
				`))
			if tpl.Handle.Pid.NamePath != `` {
				daoField.updateParse.DataTypeName = append(daoField.updateParse.DataTypeName, `case `+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.LabelList[0])+`:
				if _, ok := update[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.Pid)+`]; ok {
					daoModel.SaveData[k] = v
				} else {
					nameOfNew := gconv.String(v)
					daoModel.SaveData[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.NamePath)+`] = gdb.Raw(fmt.Sprintf(`+"`"+`REGEXP_REPLACE( %s, CONCAT( %s, '$' ), '%s' ),%s = '%s'`+"`"+`, `+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.NamePath)+`, `+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.LabelList[0])+`, nameOfNew, `+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.LabelList[0])+`, nameOfNew))
					`+gstr.CaseCamelLower(childUpdateStr)+` := []map[string]any{} //更新所有子孙级的名称路径
					oldList, _ := daoModel.CloneNew().FilterPri(daoModel.IdArr).All()
					for _, oldInfo := range oldList {
						if nameOfOld := oldInfo[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.LabelList[0])+`].String(); nameOfNew != nameOfOld {
							namePath := oldInfo[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.NamePath)+`].String()
							`+gstr.CaseCamelLower(childUpdateStr)+` = append(`+gstr.CaseCamelLower(childUpdateStr)+`, map[string]any{
								`+"`"+pIdPathOfOldStr+"`"+`: oldInfo[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.IdPath)+`],
								`+"`"+childNamePathStr+"`"+`: map[string]any{
									`+"`"+pNamePathOfOldStr+"`"+`: oldInfo[`+daoPath+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.NamePath)+`],
									`+"`"+pNamePathOfNewStr+"`"+`: namePath[:len(namePath)-len(nameOfOld)] + nameOfNew,
								},
							})
						}
					}
					if len(`+gstr.CaseCamelLower(childUpdateStr)+`) > 0 {
						daoModel.AfterUpdate[`+"`"+childUpdateStr+"`"+`] = `+gstr.CaseCamelLower(childUpdateStr)+`
					}
				}`)
			}
			daoField.updateParse.DataTypeName = append(daoField.updateParse.DataTypeName, updateParseArr...)
		}
	case internal.TypeNameIdPath, internal.TypeNameNamePath: // id_path|idPath，且pid同时存在时（才）有效；	类型：varchar或text；	// name_path|namePath，且pid，id_path|idPath同时存在时（才）有效；	类型：varchar或text；
		return myGenDaoField{}
	case internal.TypeNameLevel: // level，且pid，id_path|idPath同时存在时（才）有效；	类型：int等类型；
		daoField.filterParse.Method = internal.ReturnTypeName

		daoField.orderParse.Method = internal.ReturnTypeName
		daoField.orderParse.DataTypeName = append(daoField.orderParse.DataTypeName, `case `+daoPath+`.Columns().`+v.FieldCaseCamel+`:
				m = m.Order(`+daoTable+` + `+"`.`"+` + v)
				`+getAddOrder(tpl.Handle.Id.List, tpl.Handle.DefSort.Field, tpl.Handle.DefSort.Order))
	case internal.TypeNamePasswordSuffix: // password,passwd后缀；	类型：char(32)；
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
				daoModel.SaveData[` + daoPath + `.Columns().` + gstr.CaseCamel(tpl.Handle.PasswordMap[passwordMapKey].SaltField) + `] = salt
				password = gmd5.MustEncrypt(password + salt)`
			updateParseStr += `
				salt := grand.S(` + tpl.Handle.PasswordMap[passwordMapKey].SaltLength + `)
				daoModel.SaveData[` + daoPath + `.Columns().` + gstr.CaseCamel(tpl.Handle.PasswordMap[passwordMapKey].SaltField) + `] = salt
				password = gmd5.MustEncrypt(password + salt)`
		}
		insertParseStr += `
				daoModel.SaveData[k] = password`
		updateParseStr += `
				daoModel.SaveData[k] = password`

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
	case internal.TypeNameColorSuffix: // color后缀；	类型：varchar；
	case internal.TypeNameIdSuffix: // id后缀；	类型：int等类型或varchar或char；
		relIdObj := tpl.Handle.RelIdMap[v.FieldRaw]
		daoField.filterParse.Method = internal.ReturnTypeName
		if relIdObj.tpl != nil {
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
				m = m.Handler(` + daoPath + `.ParseJoin(` + daoTableRel + `, daoModel))`
				if relIdObj.Suffix != `` {
					fieldParseStr = `case ` + daoPathRel + `.Columns().` + gstr.CaseCamel(relIdObj.tpl.Handle.LabelList[0]) + " + `" + relIdObj.Suffix + "`:" + `
				` + daoTableRel + relIdObj.SuffixCaseCamel + ` := ` + daoPathRel + `.ParseDbTable(m.GetCtx()) + ` + "`" + relIdObj.SuffixCaseSnake + "`" + `
				m = m.Fields(` + daoTableRel + relIdObj.SuffixCaseCamel + ` + ` + "`.`" + ` + ` + daoPathRel + `.Columns().` + gstr.CaseCamel(relIdObj.tpl.Handle.LabelList[0]) + ` + ` + "` AS `" + ` + v)
				m = m.Handler(` + daoPath + `.ParseJoin(` + daoTableRel + relIdObj.SuffixCaseCamel + `, daoModel))`
				}
				daoField.fieldParse.Method = internal.ReturnTypeName
				daoField.fieldParse.DataTypeName = append(daoField.fieldParse.DataTypeName, fieldParseStr)
			}

			joinParseStr := `case ` + daoPathRel + `.ParseDbTable(m.GetCtx()):
			m = m.LeftJoin(joinTable, joinTable+` + "`.`" + `+` + daoPathRel + `.Columns().` + relIdObj.tpl.Handle.Id.List[0].FieldCaseCamel + `+` + "` = `" + `+` + daoTable + `+` + "`.`" + `+` + daoPath + `.Columns().` + v.FieldCaseCamel + `)`
			if relIdObj.Suffix != `` {
				joinParseStr = `case ` + daoPathRel + `.ParseDbTable(m.GetCtx()) + ` + "`" + relIdObj.SuffixCaseSnake + "`" + `:
			m = m.LeftJoin(` + daoPathRel + `.ParseDbTable(m.GetCtx())+` + "` AS `" + `+joinTable, joinTable+` + "`.`" + `+` + daoPathRel + `.Columns().` + relIdObj.tpl.Handle.Id.List[0].FieldCaseCamel + `+` + "` = `" + `+` + daoTable + `+` + "`.`" + `+` + daoPath + `.Columns().` + v.FieldCaseCamel + `)`
			}
			daoField.joinParse.Method = internal.ReturnTypeName
			daoField.joinParse.DataTypeName = append(daoField.joinParse.DataTypeName, joinParseStr)
		}
	case internal.TypeNameStatusSuffix, internal.TypeNameIsPrefix, internal.TypeNameIsLeaf: // status,type,scene,method,pos,position,gender,currency等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，.。;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）	// is_前缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，.。;；]等字符分隔。示例（停用：0否 1是）	// is_leaf|isLeaf，且pid，id_path|idPath同时存在时（才）有效；	类型：int等类型；
		daoField.filterParse.Method = internal.ReturnTypeName
	case internal.TypeNameSortSuffix, internal.TypeNameNoSuffix: // sort,num,number,weight等后缀；	类型：int等类型；	// no,level,rank等后缀；	类型：int等类型；
		daoField.orderParse.Method = internal.ReturnTypeName
		daoField.orderParse.DataTypeName = append(daoField.orderParse.DataTypeName, `case `+daoPath+`.Columns().`+v.FieldCaseCamel+`:
				m = m.Order(`+daoTable+` + `+"`.`"+` + v)
				`+getAddOrder(tpl.Handle.Id.List, tpl.Handle.DefSort.Field, tpl.Handle.DefSort.Order))
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
	case internal.TypeNameImageSuffix, internal.TypeNameVideoSuffix, internal.TypeNameAudioSuffix, internal.TypeNameFileSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text	// video,video_list,videoList,video_arr,videoArr等后缀；	类型：单视频varchar，多视频json或text	// audio,audio_list,audioList,audio_arr,audioArr等后缀；	类型：单音频varchar，多音频json或text	// file,file_list,fileList,file_arr,fileArr等后缀；	类型：单文件varchar，多文件json或text
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
				if slices.Contains([]string{`+"``, `0`, `[]`, `{}`"+`}, gconv.String(v)) { //gvar.New(v).IsEmpty()无法验证指针的值是空的数据
					continue
				}
				insertData, ok := daoModel.AfterInsert[`+"`"+tplEM.FieldVar+"`"+`].(map[string]any)
				if !ok {
					insertData = map[string]any{}
				}
				insertData[k] = v
				daoModel.AfterInsert[`+"`"+tplEM.FieldVar+"`"+`] = insertData`)
	insertHookStr := `insertData[` + tplEM.daoPath + `.Columns().` + gstr.CaseCamel(tplEM.RelId) + `] = id
					` + tplEM.daoPath + `.CtxDaoModel(ctx).HookInsert(insertData).Insert()`
	switch tplEM.TableType {
	case internal.TableTypeExtendOne:
		dao.insertHook = append(dao.insertHook, `case `+"`"+tplEM.FieldVar+"`"+`:
					insertData, _ := v.(map[string]any)
					`+insertHookStr)
	case internal.TableTypeMiddleOne:
		insertHookIdSuffixArr := []string{}
		insertHookIdSuffixIfArr := []string{}
		for k, v := range tplEM.FieldListOfIdSuffix {
			insertHookIdSuffixArr = append(insertHookIdSuffixArr, `_, ok`+v.FieldCaseCamel+` := insertData[`+tplEM.FieldColumnArrOfIdSuffix[k]+`]`)
			insertHookIdSuffixIfArr = append(insertHookIdSuffixIfArr, `!ok`+v.FieldCaseCamel)
		}
		dao.insertHook = append(dao.insertHook, `case `+"`"+tplEM.FieldVar+"`"+`:
					insertData, _ := v.(map[string]any)
					`+gstr.Join(append(insertHookIdSuffixArr, ``), `
					`)+`if `+gstr.Join(insertHookIdSuffixIfArr, ` && `)+` { //多ID时，全部ID都不存在（都等于0已在ParseInsert解析时已过滤，故存在就肯定不等于0）不插入。可根据自己业务修改
						continue
					}
					`+insertHookStr)
	}

	dao.updateParse = append(dao.updateParse, `case `+gstr.Join(tplEM.FieldColumnArr, `, `)+`:
				updateData, ok := daoModel.AfterUpdate[`+"`"+tplEM.FieldVar+"`"+`].(map[string]any)
				if !ok {
					updateData = map[string]any{}
				}
				updateData[k] = v
				daoModel.AfterUpdate[`+"`"+tplEM.FieldVar+"`"+`] = updateData`)
	updateHookBeforeStr := `for _, id := range daoModel.IdArr {
						updateData[` + tplEM.daoPath + `.Columns().` + gstr.CaseCamel(tplEM.RelId) + `] = id
						` + tplEM.daoPath + `.CtxDaoModel(ctx).HookInsert(updateData).OnConflict(` + tplEM.daoPath + `.Columns().` + gstr.CaseCamel(tplEM.RelId) + `).Save() // Save()只能触发HookInsert()，但因扩展表（一对一）或中间表（一对一）可能没有自增ID，HookInsert()一般无实际作用！要触发HookUpdate()时使用下方代码，同时注释该行
						/* if row, _ := ` + tplEM.daoPath + `.CtxDaoModel(ctx).Filter(` + tplEM.daoPath + `.Columns().` + gstr.CaseCamel(tplEM.RelId) + `, id).HookUpdate(updateData).UpdateAndGetAffected(); row == 0 { //更新失败，有可能记录不存在，这时做插入操作
							` + tplEM.daoPath + `.CtxDaoModel(ctx).HookInsert(updateData).Insert()
						} */
					}`
	switch tplEM.TableType {
	case internal.TableTypeExtendOne:
		dao.updateHookBefore = append(dao.updateHookBefore, `case `+"`"+tplEM.FieldVar+"`"+`:
					updateData, _ := v.(map[string]any)
					`+updateHookBeforeStr)
	case internal.TableTypeMiddleOne:
		updateHookBeforeIdSuffixArr := []string{}
		updateHookBeforeIdSuffixIfArr := []string{}
		for k, v := range tplEM.FieldListOfIdSuffix {
			updateHookBeforeIdSuffixArr = append(updateHookBeforeIdSuffixArr, gstr.CaseCamelLower(v.FieldRaw)+`, ok`+v.FieldCaseCamel+` := updateData[`+tplEM.FieldColumnArrOfIdSuffix[k]+`]`)
			updateHookBeforeIdSuffixIfArr = append(updateHookBeforeIdSuffixIfArr, `(ok`+v.FieldCaseCamel+` && gconv.Uint(`+gstr.CaseCamelLower(v.FieldRaw)+`) == 0)`)
		}
		dao.updateHookBefore = append(dao.updateHookBefore, `case `+"`"+tplEM.FieldVar+"`"+`:
					updateData, _ := v.(map[string]any)
					`+gstr.Join(append(updateHookBeforeIdSuffixArr, ``), `
					`)+`if `+gstr.Join(updateHookBeforeIdSuffixIfArr, ` && `)+` { //多ID时，全部ID存在且等于0就删除。可根据自己业务修改
						for _, id := range daoModel.IdArr {
							`+tplEM.daoPath+`.CtxDaoModel(ctx).Filter(`+tplEM.daoPath+`.Columns().`+gstr.CaseCamel(tplEM.RelId)+`, id).Delete()
						}
						continue
					}
					`+updateHookBeforeStr)
	}

	dao.deleteHookAfter = append(dao.deleteHookAfter, tplEM.daoPath+`.CtxDaoModel(ctx).Filter(`+tplEM.daoPath+`.Columns().`+gstr.CaseCamel(tplEM.RelId)+`, daoModel.IdArr).Delete()`)

	dao.joinParse = append(dao.joinParse, `case `+tplEM.daoPath+`.ParseDbTable(m.GetCtx()):
			m = m.LeftJoin(joinTable, joinTable+`+"`.`"+`+`+tplEM.daoPath+`.Columns().`+gstr.CaseCamel(tplEM.RelId)+`+`+"` = `"+`+daoModel.DbTable+`+"`.`"+`+daoThis.Columns().`+tplEM.tplOfTop.Handle.Id.List[0].FieldCaseCamel+`)`)

	fieldArrOfFilter := []string{}
	daoFieldList := []myGenDaoField{}
	for _, v := range tplEM.FieldList {
		daoField := myGenDaoField{}
		/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 开始--------*/
		switch v.FieldType {
		case internal.TypeInt, internal.TypeIntU: // `int等类型` // `int等类型（unsigned）`
		case internal.TypeFloat, internal.TypeFloatU: // `float等类型`  // `float等类型（unsigned）`
		case internal.TypeVarchar, internal.TypeChar: // `varchar类型`	// `char类型`
			if v.IsUnique || gconv.Uint(v.FieldLimitStr) <= internal.ConfigMaxLenOfStrFilter {
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
				`+getAddOrder(tplEM.tplOfTop.Handle.Id.List, tplEM.tplOfTop.Handle.DefSort.Field, tplEM.tplOfTop.Handle.DefSort.Order)+`
				m = m.Handler(daoThis.ParseJoin(`+tplEM.daoTableVar+`, daoModel))`)
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
		case internal.TypeNamePid: // pid，且与主键类型相同时（才）有效；	类型：int等类型或varchar或char；
			continue
		case internal.TypeNameIdPath, internal.TypeNameNamePath: // id_path|idPath，且pid同时存在时（才）有效；	类型：varchar或text；	// name_path|namePath，且pid，id_path|idPath同时存在时（才）有效；	类型：varchar或text；
			continue
		case internal.TypeNameLevel, internal.TypeNameIsLeaf: // level，且pid，id_path|idPath同时存在时（才）有效；	类型：int等类型；	// is_leaf|isLeaf，且pid，id_path|idPath同时存在时（才）有效；	类型：int等类型；
			continue
		case internal.TypeNamePasswordSuffix: // password,passwd后缀；	类型：char(32)；
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
		case internal.TypeNameColorSuffix: // color后缀；	类型：varchar；
		case internal.TypeNameIdSuffix: // id后缀；	类型：int等类型或varchar或char；
			relIdObj := tpl.Handle.RelIdMap[v.FieldRaw]
			daoField.filterParse.Method = internal.ReturnTypeName
			if relIdObj.tpl != nil {
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
			m = m.LeftJoin(joinTable, joinTable+` + "`.`" + `+` + daoPathRel + `.Columns().` + relIdObj.tpl.Handle.Id.List[0].FieldCaseCamel + `+` + "` = `" + `+` + tplEM.daoTable + `+` + "`.`" + `+` + tplEM.daoPath + `.Columns().` + v.FieldCaseCamel + `)`
				if relIdObj.Suffix != `` {
					joinParseStr = `case ` + daoPathRel + `.ParseDbTable(m.GetCtx()) + ` + "`" + relIdObj.SuffixCaseSnake + "`" + `:
			m = m.LeftJoin(` + daoPathRel + `.ParseDbTable(m.GetCtx())+` + "` AS `" + `+joinTable, joinTable+` + "`.`" + `+` + daoPathRel + `.Columns().` + relIdObj.tpl.Handle.Id.List[0].FieldCaseCamel + `+` + "` = `" + `+` + tplEM.daoTable + `+` + "`.`" + `+` + tplEM.daoPath + `.Columns().` + v.FieldCaseCamel + `)`
				}
				daoField.joinParse.Method = internal.ReturnTypeName
				daoField.joinParse.DataTypeName = append(daoField.joinParse.DataTypeName, joinParseStr)
			}
		case internal.TypeNameStatusSuffix, internal.TypeNameIsPrefix: // status,type,scene,method,pos,position,gender,currency等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，.。;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）	// is_前缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，.。;；]等字符分隔。示例（停用：0否 1是）
			daoField.filterParse.Method = internal.ReturnTypeName
		case internal.TypeNameSortSuffix, internal.TypeNameNoSuffix: // sort,num,number,weight等后缀；	类型：int等类型；	// no,level,rank等后缀；	类型：int等类型；
			daoField.orderParse.Method = internal.ReturnTypeName
			daoField.orderParse.DataTypeName = append(daoField.orderParse.DataTypeName, `case `+tplEM.daoPath+`.Columns().`+v.FieldCaseCamel+`:
				`+tplEM.daoTableVar+` := `+tplEM.daoPath+`.ParseDbTable(m.GetCtx())
				m = m.Order(`+tplEM.daoTableVar+` + `+"`.`"+` + v)
				`+getAddOrder(tplEM.tplOfTop.Handle.Id.List, tplEM.tplOfTop.Handle.DefSort.Field, tplEM.tplOfTop.Handle.DefSort.Order)+`
				m = m.Handler(daoThis.ParseJoin(`+tplEM.daoTableVar+`, daoModel))`)
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
		case internal.TypeNameImageSuffix, internal.TypeNameVideoSuffix, internal.TypeNameAudioSuffix, internal.TypeNameFileSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text	// video,video_list,videoList,video_arr,videoArr等后缀；	类型：单视频varchar，多视频json或text	// audio,audio_list,audioList,audio_arr,audioArr等后缀；	类型：单音频varchar，多音频json或text	// file,file_list,fileList,file_arr,fileArr等后缀；	类型：单文件varchar，多文件json或text
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
				m = m.Fields(daoModel.DbTable + `+"`.`"+` + daoThis.Columns().`+tplEM.tplOfTop.Handle.Id.List[0].FieldCaseCamel+`)
				daoModel.AfterField[v] = struct{}{}`)
	dao.insertParse = append(dao.insertParse, `case `+"`"+tplEM.FieldVar+"`"+`:
				daoModel.AfterInsert[k] = v`)
	dao.updateParse = append(dao.updateParse, `case `+"`"+tplEM.FieldVar+"`"+`:
				daoModel.AfterUpdate[k] = v`)
	if len(tplEM.FieldList) == 1 {
		dao.fieldHook = append(dao.fieldHook, `case `+"`"+tplEM.FieldVar+"`"+`:
			`+gstr.CaseCamelLower(tplEM.FieldVar)+`, _ := `+tplEM.daoPath+`.CtxDaoModel(ctx).Filter(`+tplEM.daoPath+`.Columns().`+gstr.CaseCamel(tplEM.RelId)+`, record[daoThis.Columns().`+tplEM.tplOfTop.Handle.Id.List[0].FieldCaseCamel+`]).Array(`+tplEM.daoPath+`.Columns().`+tplEM.FieldList[0].FieldCaseCamel+`)
			record[k] = gvar.New(`+gstr.CaseCamelLower(tplEM.FieldVar)+`)`)
		dao.insertHook = append(dao.insertHook, `case `+"`"+tplEM.FieldVar+"`"+`:
					vArr := gconv.SliceAny(v)
					insertList := make([]map[string]any, len(vArr))
					for index, item := range vArr {
						insertList[index] = map[string]any{`+tplEM.daoPath+`.Columns().`+gstr.CaseCamel(tplEM.RelId)+`: id, `+tplEM.daoPath+`.Columns().`+tplEM.FieldList[0].FieldCaseCamel+`: item}
					}
					`+tplEM.daoPath+`.CtxDaoModel(ctx).Data(insertList).Insert()`)
		dao.updateHookBefore = append(dao.updateHookBefore, `case `+"`"+tplEM.FieldVar+"`"+`:
					// daoIndex.SaveArrRelManyWithSort(ctx, &`+tplEM.daoPath+`, `+tplEM.daoPath+`.Columns().`+gstr.CaseCamel(tplEM.RelId)+`, `+tplEM.daoPath+`.Columns().`+tplEM.FieldList[0].FieldCaseCamel+`, gconv.SliceAny(daoModel.IdArr), gconv.SliceAny(v)) // 有顺序要求时使用，同时注释下面代码
					valArr := gconv.Strings(v)
					for _, id := range daoModel.IdArr {
						daoIndex.SaveArrRelMany(ctx, &`+tplEM.daoPath+`, `+tplEM.daoPath+`.Columns().`+gstr.CaseCamel(tplEM.RelId)+`, `+tplEM.daoPath+`.Columns().`+tplEM.FieldList[0].FieldCaseCamel+`, id, valArr )
					}`)
	} else {
		dao.fieldHook = append(dao.fieldHook, `case `+"`"+tplEM.FieldVar+"`"+`:
			`+gstr.CaseCamelLower(tplEM.FieldVar)+`, _ := `+tplEM.daoPath+`.CtxDaoModel(ctx).Filter(`+tplEM.daoPath+`.Columns().`+gstr.CaseCamel(tplEM.RelId)+`, record[daoThis.Columns().`+tplEM.tplOfTop.Handle.Id.List[0].FieldCaseCamel+`]). /* OrderAsc(`+tplEM.daoPath+`.Columns().CreatedAt). */ All()	// 有顺序要求时使用，自定义OrderAsc
			record[k] = gvar.New(gjson.MustEncodeString(`+gstr.CaseCamelLower(tplEM.FieldVar)+`)) //转成json字符串，控制器中list.Structs(&res.List)和info.Struct(&res.Info)才有效`)
		dao.insertHook = append(dao.insertHook, `case `+"`"+tplEM.FieldVar+"`"+`:
					vList := gconv.Maps(v)
					insertList := make([]map[string]any, len(vList))
					for index, item := range vList {
						insertItem := gjson.New(gjson.MustEncodeString(item)).Map()
						insertItem[`+tplEM.daoPath+`.Columns().`+gstr.CaseCamel(tplEM.RelId)+`] = id
						insertList[index] = insertItem
					}
					`+tplEM.daoPath+`.CtxDaoModel(ctx).Data(insertList).Insert()`)
		switch tplEM.TableType {
		case internal.TableTypeExtendMany:
			dao.updateHookBefore = append(dao.updateHookBefore, `case `+"`"+tplEM.FieldVar+"`"+`:
					valList := gconv.Maps(v)
					daoIndex.SaveListRelManyWithSort(ctx, &`+tplEM.daoPath+`, `+tplEM.daoPath+`.Columns().`+gstr.CaseCamel(tplEM.RelId)+`, gconv.SliceAny(daoModel.IdArr), valList)`)
		case internal.TableTypeMiddleMany:
			dao.updateHookBefore = append(dao.updateHookBefore, `case `+"`"+tplEM.FieldVar+"`"+`:
					// daoIndex.SaveListRelManyWithSort(ctx, &`+tplEM.daoPath+`, `+tplEM.daoPath+`.Columns().`+gstr.CaseCamel(tplEM.RelId)+`, gconv.SliceAny(daoModel.IdArr), gconv.Maps(v)) // 有顺序要求时使用，同时注释下面代码
					valList := gconv.Maps(v)
					for _, id := range daoModel.IdArr {
						daoIndex.SaveListRelMany(ctx, &`+tplEM.daoPath+`, `+tplEM.daoPath+`.Columns().`+gstr.CaseCamel(tplEM.RelId)+`, []string{`+gstr.Join(tplEM.FieldColumnArrOfIdSuffix, `, `)+`}, id, valList )
					}`)
		}
	}

	dao.deleteHookAfter = append(dao.deleteHookAfter, tplEM.daoPath+`.CtxDaoModel(ctx).Filter(`+tplEM.daoPath+`.Columns().`+gstr.CaseCamel(tplEM.RelId)+`, daoModel.IdArr).Delete()`)

	dao.joinParse = append(dao.joinParse, `case `+tplEM.daoPath+`.ParseDbTable(m.GetCtx()):
			m = m.LeftJoin(joinTable, joinTable+`+"`.`"+`+`+tplEM.daoPath+`.Columns().`+gstr.CaseCamel(tplEM.RelId)+`+`+"` = `"+`+daoModel.DbTable+`+"`.`"+`+daoThis.Columns().`+tplEM.tplOfTop.Handle.Id.List[0].FieldCaseCamel+`)`)

	fieldArrOfFilter := []string{}
	daoFieldList := []myGenDaoField{}
	for _, v := range tplEM.FieldList {
		daoField := myGenDaoField{}
		/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 开始--------*/
		switch v.FieldType {
		case internal.TypeInt, internal.TypeIntU: // `int等类型` // `int等类型（unsigned）`
		case internal.TypeFloat, internal.TypeFloatU: // `float等类型`  // `float等类型（unsigned）`
		case internal.TypeVarchar, internal.TypeChar: // `varchar类型`	// `char类型`
			if v.IsUnique || gconv.Uint(v.FieldLimitStr) <= internal.ConfigMaxLenOfStrFilter {
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
				`+getAddOrder(tplEM.tplOfTop.Handle.Id.List, tplEM.tplOfTop.Handle.DefSort.Field, tplEM.tplOfTop.Handle.DefSort.Order)+`
				m = m.Handler(daoThis.ParseJoin(`+tplEM.daoTableVar+`, daoModel))`)
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
		case internal.TypeNamePid: // pid，且与主键类型相同时（才）有效；	类型：int等类型或varchar或char；
			continue
		case internal.TypeNameIdPath, internal.TypeNameNamePath: // id_path|idPath，且pid同时存在时（才）有效；	类型：varchar或text；	// name_path|namePath，且pid，id_path|idPath同时存在时（才）有效；	类型：varchar或text；
			continue
		case internal.TypeNameLevel, internal.TypeNameIsLeaf: // level，且pid，id_path|idPath同时存在时（才）有效；	类型：int等类型；	// is_leaf|isLeaf，且pid，id_path|idPath同时存在时（才）有效；	类型：int等类型；
			continue
		case internal.TypeNamePasswordSuffix: // password,passwd后缀；	类型：char(32)；
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
		case internal.TypeNameColorSuffix: // color后缀；	类型：varchar；
		case internal.TypeNameIdSuffix: // id后缀；	类型：int等类型或varchar或char；
			daoField.filterParse.Method = internal.ReturnTypeName
		case internal.TypeNameStatusSuffix, internal.TypeNameIsPrefix: // status,type,scene,method,pos,position,gender,currency等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，.。;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）	// is_前缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，.。;；]等字符分隔。示例（停用：0否 1是）
			daoField.filterParse.Method = internal.ReturnTypeName
		case internal.TypeNameSortSuffix, internal.TypeNameNoSuffix: // sort,num,number,weight等后缀；	类型：int等类型；	// no,level,rank等后缀；	类型：int等类型；
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
		case internal.TypeNameImageSuffix, internal.TypeNameVideoSuffix, internal.TypeNameAudioSuffix, internal.TypeNameFileSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text	// video,video_list,videoList,video_arr,videoArr等后缀；	类型：单视频varchar，多视频json或text	// audio,audio_list,audioList,audio_arr,audioArr等后缀；	类型：单音频varchar，多音频json或text	// file,file_list,fileList,file_arr,fileArr等后缀；	类型：单文件varchar，多文件json或text
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

func getDaoOtherRel(tplOR handleOtherRel) (dao myGenDao) {
	if tplOR.tpl.ModuleDirCaseKebab != tplOR.tplOfTop.ModuleDirCaseKebab {
		dao.importDao = append(dao.importDao, `dao`+tplOR.tpl.ModuleDirCaseCamel+` "api/internal/dao/`+tplOR.tpl.ModuleDirCaseKebab+`"`)
	}
	dao.deleteHookOtherRel = append(dao.deleteHookOtherRel, tplOR.daoPath+`.CtxDaoModel(ctx).Filter(`+tplOR.daoPath+`.Columns().`+gstr.CaseCamel(tplOR.RelId)+`, daoModel.IdArr).Delete()`)
	return
}

// 追加排序。mysql排序字段有重复值时，分页会导致同一条数据可能在不同页都出现
func getAddOrder(idList []myGenField, defSortField string, defSortOrder string) (order string) {
	orderMethod := `OrderDesc`
	if gstr.ToLower(defSortOrder) == `asc` {
		orderMethod = `OrderAsc`
	}
	orderArr := []string{}
	if defSortField != `id` {
		orderArr = append(orderArr, `m = m.`+orderMethod+`(daoModel.DbTable + `+"`.`"+` + daoThis.Columns().`+gstr.CaseCamel(defSortField)+`)`)
	}
	for _, v := range idList {
		orderArr = append(orderArr, `m = m.`+orderMethod+`(daoModel.DbTable + `+"`.`"+` + daoThis.Columns().`+v.FieldCaseCamel+`)`)
	}
	order = gstr.Join(orderArr, `
				`)
	return
}
