package internal

import (
	"context"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type pgsql struct {
	common
}

func (dbHandler pgsql) GetFieldList(ctx context.Context, group, table string) (fieldList []MyGenField) {
	/* fieldListTmp, _ := g.DB(group).GetAll(ctx, `SELECT *,col_description ('`+table+`' :: REGCLASS, ordinal_position) AS column_comment FROM information_schema.COLUMNS WHERE TABLE_NAME = '`+table+`'`)
	fieldList = make([]MyGenField, len(fieldListTmp))
	for _, v := range fieldListTmp {
		field := MyGenField{
			FieldRaw:     v[`column_name`].String(),
			FieldTypeRaw: v[`udt_name`].String(),	//v[`data_type`].String()当前框架使用的与udt_name一致
			IsNull:       v[`is_nullable`].Bool(),
			Default:      v[`column_default`].Interface(),
			Comment:      v[`column_comment`].String(),
		}
		if !v[`column_default`].IsNil() {
			defaultStr := v[`column_default`].String()
			if gstr.Pos(defaultStr, `nextval`) == 0 {
				field.IsAutoInc = true
				// field.Default = 0
			} else {
				if gstr.Pos(defaultStr, `::`) != -1 {
					switch gstr.Split(defaultStr, `::`)[0] {
					case `''`:
						field.Default = ``
					case `NULL`:
						field.Default = nil
					}
				}
			}
		}
		fieldList[v[`ordinal_position`].Int()-1] = field
	} */
	fieldListTmp, _ := g.DB(group).TableFields(ctx, table)
	fieldList = make([]MyGenField, len(fieldListTmp))
	for _, v := range fieldListTmp {
		field := MyGenField{
			FieldRaw:     v.Name,
			FieldTypeRaw: v.Type,
			IsNull:       v.Null,
			Default:      v.Default,
			Comment:      v.Comment,
		}
		if !gvar.New(v.Default).IsNil() {
			defaultStr := gconv.String(v.Default)
			if gstr.Pos(defaultStr, `nextval`) == 0 {
				field.IsAutoInc = true
				// field.Default = 0
			} else {
				if gstr.Pos(defaultStr, `::`) != -1 {
					switch gstr.Split(defaultStr, `::`)[0] {
					case `''`:
						field.Default = ``
					case `NULL`:
						field.Default = nil
					}
				}
			}
		}
		fieldList[v.Index] = field
	}
	return
}

func (dbHandler pgsql) GetKeyList(ctx context.Context, group, table string) (keyList []MyGenKey) {
	indrelid, _ := g.DB(group).GetValue(ctx, `SELECT oid FROM pg_class WHERE relname = '`+table+`'`)
	keyListTmp, _ := g.DB(group).GetAll(ctx, `SELECT * FROM pg_index WHERE indrelid = '`+indrelid.String()+`'`)
	fieldList := dbHandler.GetFieldList(ctx, group, table)
	for _, v := range keyListTmp {
		if !v[`indisvalid`].Bool() {
			continue
		}
		key := MyGenKey{
			IsPrimary: v[`indisprimary`].Bool(),
			IsUnique:  v[`indisunique`].Bool(),
		}
		// g.DB(group).GetValue(ctx, `SELECT indkey FROM pg_index WHERE indexrelid = `+v[`indexrelid`].String())
		fieldIndex := v[`indkey`].Int() //TODO indkey返回值有BUG。联合索引应该返回4 5，但这里却是0。
		if fieldIndex != 0 {
			key.FieldArr = append(key.FieldArr, fieldList[fieldIndex-1].FieldRaw)
			if fieldList[fieldIndex-1].IsAutoInc {
				key.IsAutoInc = true
			}
		}
		keyList = append(keyList, key)
	}
	return
}

func (dbHandler pgsql) GetFieldLimitStr(ctx context.Context, field MyGenField, group, table string) (fieldLimitStr string) {
	fieldInfo, _ := g.DB(group).GetOne(ctx, `SELECT * FROM information_schema.COLUMNS WHERE TABLE_NAME = '`+table+`' AND column_name = '`+field.FieldRaw+`'`)
	fieldLimitStr = fieldInfo[`character_maximum_length`].String()
	return
}

func (dbHandler pgsql) GetFieldLimitFloat(ctx context.Context, field MyGenField, group, table string) (fieldLimitFloat [2]string) {
	fieldInfo, _ := g.DB(group).GetOne(ctx, `SELECT * FROM information_schema.COLUMNS WHERE TABLE_NAME = '`+table+`' AND column_name = '`+field.FieldRaw+`'`)
	fieldLimitFloat = [2]string{fieldInfo[`numeric_precision_radix`].String(), fieldInfo[`numeric_scale`].String()}
	return
}

func (dbHandler pgsql) GetFieldLimitInt(ctx context.Context, field MyGenField, group, table string) (fieldLimitInt int) {
	fieldLimitInt = 4
	switch field.FieldTypeRaw {
	case `int2`:
		fieldLimitInt = 2
	case `int4`:
		fieldLimitInt = 4
	case `int8`:
		fieldLimitInt = 8
	}
	return
}

func (dbHandler pgsql) GetFieldType(ctx context.Context, field MyGenField, group, table string) (fieldType MyGenFieldType) {
	switch field.FieldTypeRaw {
	case `int2`, `int4`, `int8`: //int等类型
		fieldType = TypeInt
		// pgsql数字类型不分正负。故只判断id后缀为非0参数
		if field.IsAutoInc {
			fieldType = TypeIntU
		} else {
			fieldCaseSnake := gstr.CaseSnake(field.FieldRaw)
			fieldCaseSnakeRemove := gstr.Split(fieldCaseSnake, `_of_`)[0]
			fieldSplitArr := gstr.Split(fieldCaseSnakeRemove, `_`)
			fieldSuffix := fieldSplitArr[len(fieldSplitArr)-1]
			if fieldSuffix != `id` {
				fieldType = TypeIntU
			}
		}
	case `numeric`, `float4`, `float8`: //float类型
		fieldType = TypeFloat
		// pgsql数字类型不分正负
	case `varchar`: //varchar类型
		fieldType = TypeVarchar
	case `bpchar`: //char类型
		fieldType = TypeChar
	case `text`: //text类型
		fieldType = TypeText
	case `json`: //json类型
		fieldType = TypeJson
	case `timestamp`: //datetime类型（在pgsql中，timestamp类型就是datetime类型）
		fieldType = TypeDatetime
	case `date`: //date类型
		fieldType = TypeDate
	case `time`: //time类型
		fieldType = TypeTime
	}
	return
}
