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
	keyListTmp, _ := g.DB(group).GetAll(ctx, `SELECT * FROM pg_index WHERE indrelid = `+indrelid.String()) //联合索引时，indkey返回值有BUG。正确应该返回4 5，实际却返回0
	fieldList := dbHandler.GetFieldList(ctx, group, table)
	for _, v := range keyListTmp {
		if !v[`indisvalid`].Bool() {
			continue
		}
		key := MyGenKey{
			IsPrimary: v[`indisprimary`].Bool(),
			IsUnique:  v[`indisunique`].Bool(),
		}
		/* keyFieldIndexArr := gstr.Split(v[`indkey`].String(), ` `) //联合索引时，indkey返回值有BUG。正确应该返回4 5，实际却返回0
		for _, keyFieldIndex := range keyFieldIndexArr {
			fieldIndex := gconv.Int(keyFieldIndex) - 1
			key.FieldArr = append(key.FieldArr, fieldList[fieldIndex].FieldRaw)
			if fieldList[fieldIndex].IsAutoInc {
				key.IsAutoInc = true
			}
		} */
		keyFieldList, _ := g.DB(group).GetAll(ctx, `SELECT * FROM pg_attribute WHERE attrelid = `+v[`indexrelid`].String()+` ORDER BY attnum`)
		for _, keyField := range keyFieldList {
			for _, field := range fieldList {
				if keyField[`attname`].String() != field.FieldRaw {
					continue
				}
				key.FieldArr = append(key.FieldArr, field.FieldRaw)
				if field.IsAutoInc {
					key.IsAutoInc = field.IsAutoInc
					key.FieldTypeRaw = field.FieldTypeRaw
				}
			}
		}
		keyList = append(keyList, key)
	}
	return
}

func (dbHandler pgsql) GetFieldType(ctx context.Context, field MyGenField, group, table string) (fieldType MyGenFieldType) {
	switch field.FieldTypeRaw {
	case `int2`, `int4`, `int8`: //int等类型
		fieldType = TypeInt
		// pgsql数字类型不分正负
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

func (dbHandler pgsql) GetFieldLimitStr(ctx context.Context, field MyGenField, group, table string) (fieldLimitStr string) {
	fieldInfo, _ := g.DB(group).GetOne(ctx, `SELECT * FROM information_schema.COLUMNS WHERE TABLE_NAME = '`+table+`' AND column_name = '`+field.FieldRaw+`'`)
	fieldLimitStr = fieldInfo[`character_maximum_length`].String()
	return
}

func (dbHandler pgsql) GetFieldLimitInt(ctx context.Context, field MyGenField, group, table string) (fieldLimitInt MyGenFieldLimitInt) {
	switch field.FieldTypeRaw {
	case `int2`:
		fieldLimitInt.Size = 2
		fieldLimitInt.Min = `-32768`
		fieldLimitInt.Max = `32767`
	case `int4`:
		fieldLimitInt.Size = 4
		fieldLimitInt.Min = `-2147483648`
		fieldLimitInt.Max = `2147483647`
	case `int8`:
		fieldLimitInt.Size = 8
		fieldLimitInt.Min = `-9223372036854775808`
		fieldLimitInt.Max = `9223372036854775807`
	}
	if field.IsAutoInc {
		fieldLimitInt.Min = `1`
	}
	return
}

func (dbHandler pgsql) GetFieldLimitFloat(ctx context.Context, field MyGenField, group, table string) (fieldLimitFloat MyGenFieldLimitFloat) {
	fieldInfo, _ := g.DB(group).GetOne(ctx, `SELECT * FROM information_schema.COLUMNS WHERE TABLE_NAME = '`+table+`' AND column_name = '`+field.FieldRaw+`'`)
	switch field.FieldTypeRaw {
	case `numeric`:
		fieldLimitFloat.Size = fieldInfo[`numeric_precision`].Int()
		fieldLimitFloat.Precision = fieldInfo[`numeric_scale`].Int()
		maxInt := `0`
		if fieldLimitFloat.Size-fieldLimitFloat.Precision > 0 {
			maxInt = gstr.Repeat(`9`, fieldLimitFloat.Size-fieldLimitFloat.Precision)
		}
		fieldLimitFloat.Max = maxInt + `.` + gstr.Repeat(`9`, fieldLimitFloat.Precision)
		fieldLimitFloat.Min = `-` + fieldLimitFloat.Max
	case `float4`, `float8`:
		fieldLimitFloat.Size = 10
		fieldLimitFloat.Precision = 2
	}
	return
}

func (dbHandler pgsql) GetFuncFieldFormat(dbFuncCode MyGenDbFuncCode, field string) (fieldFormat string) {
	fieldFormat = field //默认值
	switch dbFuncCode {
	case DbFuncCodeNULLIF, DbFuncCodeCOALESCE, DbFuncCodeREPLACE: //Mysql和Postgresql通用（差别：Postgresql数字字段需加::TEXT转成字符串）
		fieldFormat = field + `::TEXT`
	case DbFuncCodeCONCAT: //Mysql和Postgresql通用
	case DbFuncCodeCONCAT_WS: //Mysql和Postgresql通用
	}
	return
}
